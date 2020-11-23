// Code generated for package schema by go-bindata DO NOT EDIT. (@generated)
// sources:
// gateway/schema/schema.graphql
package schema

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

var _gatewaySchemaSchemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x51\x31\x4e\x03\x31\x10\xec\xf7\x15\xeb\x2e\xf9\x82\x5b\xaa\x14\x48\x81\x28\x15\xa2\xb0\xec\x05\x8c\xee\xec\x63\x6f\x0d\x8a\x50\xfe\x8e\xbc\xb9\x9c\xac\x3b\x68\x50\x2a\x7b\x47\x33\xe3\x99\xb5\x9c\x06\xc2\x83\x7f\xa3\xde\xe1\x37\x20\x7e\x14\xe2\x93\xc5\x87\x7a\x00\x62\x5f\xc4\x49\xcc\xc9\xe2\xfd\x74\x83\x33\x80\x8a\x94\xa2\x9a\xcf\x48\x5f\xc4\x9b\xad\xc5\xe3\x48\x6c\x66\xc6\x55\xa2\x24\xcf\xe4\x84\xf6\x9c\x43\xf1\xb2\x89\x69\x28\x62\xf1\xae\x05\x77\x15\x33\x5b\x8b\xd3\x6c\x00\xb1\x0c\x61\xad\x3a\xb6\xe0\xdf\xaa\x1a\x66\xcf\xf9\x9d\xbc\x3c\xe6\x8e\xae\xea\xa7\x5d\x12\xf3\x6c\x56\x69\xeb\xa4\x49\x63\xb0\x78\x10\x8e\xe9\x55\xbd\x46\xe2\xe4\x7a\x6a\x31\xea\x5d\xec\x5a\xe0\xa5\x74\x5d\x4b\xaa\x7d\xdd\x02\x3a\x03\x68\x82\x5f\x4a\x5f\x16\x94\x43\xab\x5f\xbe\x19\x68\xf4\x1c\x87\xcb\x67\x2c\x2d\xd7\x1b\x99\x9b\xd4\xb6\xff\x76\xd7\xc5\x4c\xb6\x37\x71\xfc\x09\x00\x00\xff\xff\xf2\xb8\x49\x4a\x6f\x02\x00\x00")

func gatewaySchemaSchemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_gatewaySchemaSchemaGraphql,
		"gateway/schema/schema.graphql",
	)
}

func gatewaySchemaSchemaGraphql() (*asset, error) {
	bytes, err := gatewaySchemaSchemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gateway/schema/schema.graphql", size: 623, mode: os.FileMode(420), modTime: time.Unix(1606143751, 0)}
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
	"gateway/schema/schema.graphql": gatewaySchemaSchemaGraphql,
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
	"gateway": &bintree{nil, map[string]*bintree{
		"schema": &bintree{nil, map[string]*bintree{
			"schema.graphql": &bintree{gatewaySchemaSchemaGraphql, map[string]*bintree{}},
		}},
	}},
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
