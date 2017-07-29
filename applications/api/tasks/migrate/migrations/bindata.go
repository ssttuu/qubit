// Code generated by go-bindata.
// sources:
// sql/1497233022_initial_schema.down.sql
// sql/1497233022_initial_schema.up.sql
// DO NOT EDIT!

package migrations

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __1497233022_initial_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xcc\x4d\x0a\xc2\x40\x0c\xc5\xf1\x7d\x4f\xd1\x7b\xcc\x4a\x69\x17\x82\xa8\x88\x0b\x77\x61\x68\x1f\x25\x42\x93\x9a\x8c\x45\x3c\xbd\x30\xc5\xaf\x71\x17\xf2\xfe\xfc\x9a\xe3\xfe\x50\x6f\x76\x4d\x7b\xae\xb9\xbf\x93\x41\x7a\x18\xe9\x04\x8b\x49\xcd\x89\xbc\x83\x80\xb8\x7f\x5d\x33\xcc\x59\x85\xde\x4d\xde\x12\x8f\x08\x55\xc6\x4e\xab\xf5\xb6\xad\x17\xc8\x3f\x52\xa8\xbe\xe7\xc5\xc2\x0c\x49\x1e\xfe\x86\x22\xe6\x31\x0e\xf8\xcd\xf2\x8b\x1c\xd7\x1b\xa4\x2b\xfb\xc9\xf4\x82\xae\x80\xd5\x86\x28\xfc\x88\x89\x55\x3c\x54\xcf\x00\x00\x00\xff\xff\x1c\x68\x3a\xf2\xf8\x00\x00\x00")

func _1497233022_initial_schemaDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1497233022_initial_schemaDownSql,
		"1497233022_initial_schema.down.sql",
	)
}

func _1497233022_initial_schemaDownSql() (*asset, error) {
	bytes, err := _1497233022_initial_schemaDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1497233022_initial_schema.down.sql", size: 248, mode: os.FileMode(436), modTime: time.Unix(1499468127, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1497233022_initial_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x95\x5f\x6f\xd3\x30\x14\xc5\xdf\xf7\x29\xae\x2a\x1e\x5a\xa9\x43\x1d\xe2\x01\x81\x10\x4a\x5b\x6f\x0b\x94\x04\x9c\x14\xd4\xa7\xca\x73\xae\x1a\x43\x6b\x87\xd8\x59\xc7\x3e\x3d\x4a\xd2\x34\x75\xd3\x7f\x62\x79\x8d\x7d\xce\x3d\xf7\x77\x1d\x7b\x44\x89\x13\x12\x08\x9d\xe1\x84\x80\x4a\x17\x4c\x8a\x67\x66\x84\x92\x1a\xba\x57\x00\x00\x22\x82\xa1\x7b\x17\x10\xea\x3a\x13\xf8\x46\xdd\xaf\x0e\x9d\xc1\x17\x32\xeb\x17\xab\x1d\xc9\x56\xd8\x81\x1f\x0e\x1d\xdd\x3b\xb4\x7b\x33\x18\xf4\xc0\xf3\x43\xf0\xa6\x93\xc9\x55\xef\xc3\x95\x65\x9f\xa4\xea\x17\x72\x73\x99\xf3\x6e\x96\x79\xb9\xd5\xf5\xc2\xad\xf9\xf9\xf2\xe5\x8e\x5b\x9f\x12\xf7\xce\xcb\x7d\xa1\xbb\xe7\xd9\x03\x4a\x6e\x09\x25\xde\x88\x04\x76\xef\x5d\x11\xf5\xf2\xf8\x76\x7e\xb1\x62\x0b\x9c\x6b\xfc\x93\xa1\xe4\x78\x59\x1b\x9b\x9e\xdb\xea\xa0\xb6\xb3\xc2\x57\x64\xab\xdc\xcd\xd8\x97\xa5\xb5\x3b\x6c\x2b\x74\xc3\xd5\xca\xbe\x47\x75\x8b\xde\x6e\x42\x73\x94\x38\xc7\x47\x94\x17\x9e\x9f\x52\x70\xac\x83\x48\xad\xe5\xfc\x11\x53\x2d\x94\x04\xd7\x0b\x77\xbe\xf2\x98\xc9\x05\xce\x23\x66\x18\x7c\x0e\x7c\x6f\x4f\x99\x25\xbb\xba\xe6\xe2\x19\xb9\x05\xa6\xca\x68\xf1\x28\x3e\x16\x18\x4a\xc5\xf5\x35\x84\xfe\xd8\x7f\x0f\x41\xac\xb2\x65\x04\x65\x01\x28\x0a\x3c\x20\x08\xc9\x97\x59\x84\x11\x08\x09\x26\x46\x9d\x7f\x89\x04\x47\xfd\x09\x5c\x30\xb1\x90\xbf\x41\x18\x58\x8b\xe5\x12\x1e\x96\x6a\x6d\x5b\xaa\xcc\xe4\xa2\x5c\x82\x4f\xa0\xc5\x33\x6a\x30\x4a\xc1\x2a\xe3\x31\x18\x95\xfb\xaf\x55\x6a\x62\x10\xe6\x75\xa5\x0c\xc8\x84\x8c\x42\x8b\x60\xbf\x49\xee\xe7\x3d\xa1\x64\x3b\x84\x8f\xaf\x6e\xc0\xa7\x63\x42\x61\x38\xb3\xd9\x8f\x49\x30\x2a\x9c\xa7\x9e\xfb\x7d\x4a\x6a\x26\x7d\x6b\x5f\xcd\x62\x53\xbe\x1e\x43\x7f\x9f\xfa\x89\xd2\x3b\xc3\x73\x8e\xd6\xad\x37\x35\xff\xa5\x72\x97\x96\x2c\xd1\xb1\x6a\xe5\x24\x1e\x3f\x4c\xe7\xff\x32\x95\x60\xca\x8c\x4a\x75\x3b\x27\xad\x89\x62\x97\xc3\x01\x10\xed\x5c\x81\xff\x73\xc1\xd9\x61\x6a\x0e\x2f\x9f\x07\x57\xd2\xe0\x93\x39\x85\x5d\xc8\x24\xcb\x13\xea\xca\xc1\xa1\xd4\x39\xfc\x20\xbe\x19\xbc\x7d\xb7\x2f\xef\x98\xbf\xc9\xe9\xb9\x26\x2c\x65\x2b\x34\xf8\xf2\xc1\x16\xac\x78\x8a\xcc\xe0\x86\x15\x57\x52\x22\xbf\xfc\x85\xaf\x9a\x3d\x4c\x6b\xb3\x5a\x5c\x1f\xcd\x55\x95\x99\x53\xe2\x6a\xf9\x88\xda\x7e\x43\x36\x31\xec\x37\xbb\x9a\x7b\x7d\x86\xed\xf7\xbe\xaa\x7f\x5c\x95\x03\xfa\x17\x00\x00\xff\xff\xb1\x32\x43\x14\x04\x09\x00\x00")

func _1497233022_initial_schemaUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1497233022_initial_schemaUpSql,
		"1497233022_initial_schema.up.sql",
	)
}

func _1497233022_initial_schemaUpSql() (*asset, error) {
	bytes, err := _1497233022_initial_schemaUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1497233022_initial_schema.up.sql", size: 2308, mode: os.FileMode(436), modTime: time.Unix(1500833845, 0)}
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
	"1497233022_initial_schema.down.sql": _1497233022_initial_schemaDownSql,
	"1497233022_initial_schema.up.sql":   _1497233022_initial_schemaUpSql,
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
	"1497233022_initial_schema.down.sql": &bintree{_1497233022_initial_schemaDownSql, map[string]*bintree{}},
	"1497233022_initial_schema.up.sql":   &bintree{_1497233022_initial_schemaUpSql, map[string]*bintree{}},
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
