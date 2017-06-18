// Code generated by go-bindata.
// sources:
// sql/1497233022_initial_schema.down.sql
// sql/1497233022_initial_schema.up.sql
// DO NOT EDIT!

package migrate

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

var __1497233022_initial_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xc8\x2f\x48\x2d\x4a\x2c\xc9\x2f\x8a\xcf\xcc\x2b\x28\x2d\x29\xb6\xe6\xe2\xc2\x22\x89\x26\x9c\x99\x9b\x98\x9e\x8a\x4d\x2c\xbe\x38\xb5\xb0\x34\x35\x2f\x19\x5d\xb2\x38\x39\x35\x0f\x5d\xac\xa0\x28\x3f\x2b\x35\x19\xc3\xc6\xa2\xf4\xc4\xbc\xcc\xaa\xc4\x92\xcc\xfc\xbc\x62\x6b\x2e\x40\x00\x00\x00\xff\xff\x34\x65\x70\xe8\xa9\x00\x00\x00")

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

	info := bindataFileInfo{name: "1497233022_initial_schema.down.sql", size: 169, mode: os.FileMode(436), modTime: time.Unix(1497717787, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1497233022_initial_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x94\xbd\x6e\xc2\x30\x14\x85\xf7\x3c\xc5\x15\x53\x22\x75\xa0\x43\xa7\x4e\x06\x19\x9a\x36\x35\x95\x49\x2b\x31\x45\x56\xb8\x42\xae\x84\x9d\x3a\x46\x6a\xfb\xf4\x55\x12\x12\x70\x7e\x20\x03\x4b\xe7\x6b\x7f\xfe\xce\x89\x9d\x39\xa7\x24\xa6\x10\x93\x59\x44\x41\x9b\x9d\x50\xf2\x57\x58\xa9\x55\x0e\xbe\x07\x00\x20\xb7\x30\x0b\x97\x6b\xca\x43\x12\xc1\x1b\x0f\x5f\x09\xdf\xc0\x0b\xdd\xdc\x95\xd3\x89\x12\x7b\x9c\xc0\x07\xe1\xf3\x27\xc2\xfd\xfb\xe9\x34\x00\xb6\x8a\x81\xbd\x47\x91\x17\x3c\x7a\x9e\xc3\xcf\x8c\xfe\xc4\xd4\x8e\x43\x9f\xcb\x24\xd5\xd2\x90\xc5\x0d\xfd\xfa\xf9\xd5\x8a\xc5\x8a\xd3\x70\xc9\x0a\x2e\xf8\x2d\x66\x00\x9c\x2e\x28\xa7\x6c\x4e\xd7\x6e\x78\x5f\x6e\x83\xae\x7f\x9e\xa2\xc2\x71\xf6\xc7\xa8\xb7\x12\x3f\xe1\x1c\xe7\xba\xd0\x7e\x5d\xb9\x17\x3b\x4c\x72\xfc\x3a\xa0\x4a\xff\x9b\xf7\x38\x5d\x37\xe2\xad\xac\x3b\x54\x47\xbe\x55\x6b\x7f\x06\x9d\xa1\x11\x56\x9b\x71\x31\xca\x8b\x35\x68\x9f\x6a\x65\xf1\xdb\x9e\xe9\xb7\xed\xed\x4f\x86\xcd\xf8\xa1\x33\xbe\x1e\x3f\x13\x46\xec\xd1\xa2\x49\x8c\xd6\x16\x9e\xd7\x2b\x76\xa9\xa0\xda\xd7\xe9\xa5\x7a\x1d\x97\xeb\x48\xa4\xca\x0e\x63\x7f\x00\xcd\x9e\x81\x5e\x4a\xd6\xe0\xd4\x7d\xf8\x27\x96\xfb\xe8\xeb\xcf\x54\x68\xf7\xdd\x84\xe3\x11\xc3\x9b\x8a\xac\x7f\x01\x00\x00\xff\xff\xfd\x17\x7d\xd0\x45\x05\x00\x00")

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

	info := bindataFileInfo{name: "1497233022_initial_schema.up.sql", size: 1349, mode: os.FileMode(436), modTime: time.Unix(1497952893, 0)}
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
