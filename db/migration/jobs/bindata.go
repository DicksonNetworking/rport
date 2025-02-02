// Code generated for package jobs by go-bindata DO NOT EDIT. (@generated)
// sources:
// 001_init.down.sql
// 001_init.up.sql
// 002_schedules.down.sql
// 002_schedules.up.sql
// 003_multi_job_schedule_id.down.sql
// 003_multi_job_schedule_id.up.sql
package jobs

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __001_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\xf0\xf4\x73\x71\x8d\x50\xc8\x4c\xa9\x88\xcf\xca\x4f\x2a\x8e\x4f\xce\xc9\x4c\xcd\x2b\x89\xcf\x4c\x89\x2f\xc9\xcc\x4d\xb5\xe6\xe2\xc2\xa6\x28\xb7\x34\xa7\x24\x33\x3e\x33\x05\x26\x1d\xe2\xe8\xe4\xe3\xaa\x00\x92\x42\x15\x81\xa8\x83\x88\x03\x02\x00\x00\xff\xff\x32\x12\x92\x70\x6d\x00\x00\x00")

func _001_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_initDownSql,
		"001_init.down.sql",
	)
}

func _001_initDownSql() (*asset, error) {
	bytes, err := _001_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_init.down.sql", size: 109, mode: os.FileMode(420), modTime: time.Unix(1656320425, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __001_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\xcd\x6e\xab\x30\x10\x85\xf7\x3c\xc5\x2c\x41\xca\x1b\x64\xc5\x85\xe1\xd6\x2a\x31\x95\x99\x28\x64\x65\x91\xd8\x55\x07\x91\x54\x8a\x1d\xa9\x7d\xfb\xaa\xa6\x41\x71\xd5\xf4\x67\xfd\x1d\xcf\x7c\x73\x5c\x28\xcc\x09\x81\xf2\x7f\x35\x82\xa8\x40\x36\x04\xd8\x89\x96\x5a\x38\x9c\x47\xcf\x7a\x78\xde\x39\x48\x13\x00\x80\x81\x0d\x10\x76\x04\x0f\x4a\xac\x72\xb5\x85\x7b\xdc\x86\x07\x72\x5d\xd7\x8b\x10\x71\xbe\x3f\x79\x6b\x74\xef\xa1\xcc\x09\x49\xac\xf0\x53\x62\x7f\xb2\xfd\x7b\x62\xf7\x3a\xcd\x8a\xa9\xb1\xbe\xe7\xd1\xc5\x28\xc9\x60\x23\xe8\xae\x59\x13\xa8\x66\x23\xca\x65\x92\x44\xda\x7f\x56\xf4\x67\xf7\xd5\xf2\x9f\xe5\x1f\xf9\xc8\xee\x29\x8e\xfc\xe6\xac\xfd\xc8\xf6\xe8\xf5\x45\x2e\x86\x73\xcf\x17\xfe\x4d\x15\x13\xaa\x1a\x85\xe2\xbf\x0c\xc7\xa5\xd7\xcf\x33\x50\x58\xa1\x42\x59\xe0\xf5\xff\xa5\x03\x9b\xec\x76\x8b\x42\x96\xd8\x01\x9b\x97\x10\xd6\xb3\xac\xf6\x7c\xb0\x61\x61\x23\x3f\x4a\x9e\xd9\x22\xee\x02\xdb\x22\xbb\x39\x70\x12\x61\x13\x8f\x8a\xbc\x97\xc9\x5b\x00\x00\x00\xff\xff\x97\x9b\x70\x8a\x89\x02\x00\x00")

func _001_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_initUpSql,
		"001_init.up.sql",
	)
}

func _001_initUpSql() (*asset, error) {
	bytes, err := _001_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_init.up.sql", size: 649, mode: os.FileMode(420), modTime: time.Unix(1656320425, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __002_schedulesDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x4e\xce\x48\x4d\x29\xcd\x49\x2d\xb6\xe6\x02\x04\x00\x00\xff\xff\x0b\xb6\x9b\xfb\x16\x00\x00\x00")

func _002_schedulesDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__002_schedulesDownSql,
		"002_schedules.down.sql",
	)
}

func _002_schedulesDownSql() (*asset, error) {
	bytes, err := _002_schedulesDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "002_schedules.down.sql", size: 22, mode: os.FileMode(420), modTime: time.Unix(1656320425, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __002_schedulesUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x28\x4e\xce\x48\x4d\x29\xcd\x49\x2d\x56\xd0\xe0\x52\x50\x50\x50\xc8\x4c\x51\x08\x71\x8d\x08\x51\x08\x08\xf2\xf4\x75\x0c\x8a\x54\xf0\x76\x8d\x54\xf0\xf3\x0f\x51\xf0\x0b\xf5\xf1\xd1\x01\xab\x48\x2e\x4a\x4d\x2c\x49\x4d\x89\x4f\x2c\x51\x70\x71\x0c\x71\x0d\xf1\xf4\x75\xc5\xa1\x22\xa9\x12\x62\x16\xaa\x6c\x5e\x62\x6e\x2a\x36\x71\x98\x4b\xb0\xc9\x95\x54\x16\x60\x15\x4f\x49\x2d\x49\xcc\xcc\x29\x46\x95\xe2\xd2\xb4\xe6\x02\x04\x00\x00\xff\xff\x2a\xa0\x1c\xb3\xe4\x00\x00\x00")

func _002_schedulesUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__002_schedulesUpSql,
		"002_schedules.up.sql",
	)
}

func _002_schedulesUpSql() (*asset, error) {
	bytes, err := _002_schedulesUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "002_schedules.up.sql", size: 228, mode: os.FileMode(420), modTime: time.Unix(1656320425, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __003_multi_job_schedule_idDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _003_multi_job_schedule_idDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__003_multi_job_schedule_idDownSql,
		"003_multi_job_schedule_id.down.sql",
	)
}

func _003_multi_job_schedule_idDownSql() (*asset, error) {
	bytes, err := _003_multi_job_schedule_idDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "003_multi_job_schedule_id.down.sql", size: 0, mode: os.FileMode(420), modTime: time.Unix(1656320425, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __003_multi_job_schedule_idUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xc8\x2d\xcd\x29\xc9\x8c\xcf\xca\x4f\x2a\x56\x70\x74\x71\x51\x28\x4e\xce\x48\x4d\x29\xcd\x49\x8d\xcf\x4c\x51\x08\x71\x8d\x08\x51\xf0\x0b\xf5\xf1\xb1\xe6\x02\x04\x00\x00\xff\xff\x79\xff\xa0\x5d\x32\x00\x00\x00")

func _003_multi_job_schedule_idUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__003_multi_job_schedule_idUpSql,
		"003_multi_job_schedule_id.up.sql",
	)
}

func _003_multi_job_schedule_idUpSql() (*asset, error) {
	bytes, err := _003_multi_job_schedule_idUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "003_multi_job_schedule_id.up.sql", size: 50, mode: os.FileMode(420), modTime: time.Unix(1656320425, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"001_init.down.sql":                  _001_initDownSql,
	"001_init.up.sql":                    _001_initUpSql,
	"002_schedules.down.sql":             _002_schedulesDownSql,
	"002_schedules.up.sql":               _002_schedulesUpSql,
	"003_multi_job_schedule_id.down.sql": _003_multi_job_schedule_idDownSql,
	"003_multi_job_schedule_id.up.sql":   _003_multi_job_schedule_idUpSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"001_init.down.sql":                  &bintree{_001_initDownSql, map[string]*bintree{}},
	"001_init.up.sql":                    &bintree{_001_initUpSql, map[string]*bintree{}},
	"002_schedules.down.sql":             &bintree{_002_schedulesDownSql, map[string]*bintree{}},
	"002_schedules.up.sql":               &bintree{_002_schedulesUpSql, map[string]*bintree{}},
	"003_multi_job_schedule_id.down.sql": &bintree{_003_multi_job_schedule_idDownSql, map[string]*bintree{}},
	"003_multi_job_schedule_id.up.sql":   &bintree{_003_multi_job_schedule_idUpSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
