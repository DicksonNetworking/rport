package sessions

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// GetInitStateFromFile returns an initial Client Session Repository state populated with sessions from the file.
func GetInitStateFromFile(fileName string, expiration *time.Duration) ([]*ClientSession, error) {
	log.Println("Start to get init Client Session Repository state from file.")

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSR file: %v", err)
	}
	log.Println("CSR file opened. Reading...")
	defer file.Close()

	return getInitState(file, expiration)
}

func getInitState(r io.Reader, expiration *time.Duration) ([]*ClientSession, error) {
	decoder := json.NewDecoder(r)
	// read array open bracket
	if _, err := decoder.Token(); err != nil {
		if err == io.EOF {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to parse CSR data: %v", err)
	}

	var sessions []*ClientSession
	var err error
	var obsolete int
	for decoder.More() {
		var session ClientSession
		if err = decoder.Decode(&session); err != nil {
			err = fmt.Errorf("failed to parse client session: %v", err)
			break
		}

		if session.Disconnected == nil || !session.Obsolete(expiration) {
			sessions = append(sessions, &session)
		} else {
			obsolete++
		}
	}

	log.Printf("Got %d and skipped %d obsolete client session(s).\n", len(sessions), obsolete)

	// mark previously connected client sessions as disconnected with current time
	now := now()
	for _, session := range sessions {
		if session.Disconnected == nil {
			session.Disconnected = &now
		}
	}

	return sessions, err
}
