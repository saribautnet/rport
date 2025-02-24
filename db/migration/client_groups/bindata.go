// Code generated for package client_groups by go-bindata DO NOT EDIT. (@generated)
// sources:
// 001_init.down.sql
// 001_init.up.sql
// 002_add_allowed_user_groups.down.sql
// 002_add_allowed_user_groups.up.sql
package client_groups

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

var __001_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\xce\xc9\x4c\xcd\x2b\x89\x4f\x2f\xca\x2f\x2d\x28\xb6\xe6\x02\x04\x00\x00\xff\xff\xee\xde\xdd\xb3\x1a\x00\x00\x00")

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

	info := bindataFileInfo{name: "001_init.down.sql", size: 26, mode: os.FileMode(420), modTime: time.Unix(1656320425, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __001_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x48\xce\xc9\x4c\xcd\x2b\x89\x4f\x2f\xca\x2f\x2d\x28\x56\xd0\xe0\x52\x50\x50\x50\xc8\x4c\x51\x08\x71\x8d\x08\x51\x08\x08\xf2\xf4\x75\x0c\x8a\x54\xf0\x76\x8d\x54\xf0\xf3\x0f\x51\xf0\x0b\xf5\xf1\xd1\xe1\xe2\x4c\x49\x2d\x4e\x2e\xca\x2c\x28\xc9\xcc\xcf\x83\xa8\x43\x92\x2b\x48\x2c\x4a\xcc\x2d\x46\x15\xe6\xd2\x54\x08\xf7\x0c\xf1\xf0\x0f\x0d\x51\x08\xf2\x0f\xf7\x74\xb1\xe6\x02\x04\x00\x00\xff\xff\xa5\xc7\xf9\xc7\x82\x00\x00\x00")

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

	info := bindataFileInfo{name: "001_init.up.sql", size: 130, mode: os.FileMode(420), modTime: time.Unix(1656320425, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __002_add_allowed_user_groupsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _002_add_allowed_user_groupsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__002_add_allowed_user_groupsDownSql,
		"002_add_allowed_user_groups.down.sql",
	)
}

func _002_add_allowed_user_groupsDownSql() (*asset, error) {
	bytes, err := _002_add_allowed_user_groupsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "002_add_allowed_user_groups.down.sql", size: 0, mode: os.FileMode(420), modTime: time.Unix(1663914692, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __002_add_allowed_user_groupsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\xcc\x29\x49\x2d\x52\x28\x49\x4c\xca\x49\x55\x50\x4a\xce\xc9\x4c\xcd\x2b\x89\x4f\x2f\xca\x2f\x2d\x28\x56\x52\x48\x4c\x49\x51\x48\xcc\xc9\xc9\x2f\x4f\x4d\x89\x2f\x2d\x4e\x2d\x82\x4a\x28\x84\xb8\x46\x84\x28\xf8\xf9\x87\x28\xf8\x85\xfa\xf8\x28\xb8\xb8\xba\x39\x86\xfa\x84\x28\xa8\x47\xc7\xaa\x5b\x03\x02\x00\x00\xff\xff\x73\xd5\xd2\x17\x4f\x00\x00\x00")

func _002_add_allowed_user_groupsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__002_add_allowed_user_groupsUpSql,
		"002_add_allowed_user_groups.up.sql",
	)
}

func _002_add_allowed_user_groupsUpSql() (*asset, error) {
	bytes, err := _002_add_allowed_user_groupsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "002_add_allowed_user_groups.up.sql", size: 79, mode: os.FileMode(420), modTime: time.Unix(1663914692, 0)}
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
	"001_init.down.sql":                    _001_initDownSql,
	"001_init.up.sql":                      _001_initUpSql,
	"002_add_allowed_user_groups.down.sql": _002_add_allowed_user_groupsDownSql,
	"002_add_allowed_user_groups.up.sql":   _002_add_allowed_user_groupsUpSql,
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
	"001_init.down.sql":                    &bintree{_001_initDownSql, map[string]*bintree{}},
	"001_init.up.sql":                      &bintree{_001_initUpSql, map[string]*bintree{}},
	"002_add_allowed_user_groups.down.sql": &bintree{_002_add_allowed_user_groupsDownSql, map[string]*bintree{}},
	"002_add_allowed_user_groups.up.sql":   &bintree{_002_add_allowed_user_groupsUpSql, map[string]*bintree{}},
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
