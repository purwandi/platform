// Code generated for package schema by go-bindata DO NOT EDIT. (@generated)
// sources:
// gateway/schema/module.graphql
// gateway/schema/product.graphql
// gateway/schema/schema.graphql
// gateway/schema/user.graphql
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

var _gatewaySchemaModuleGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x8e\xb1\xca\x02\x31\x10\x84\xfb\x7d\x8a\xf9\xb9\xb7\xf8\x5b\xab\x2b\xac\xb4\x97\x90\x2c\x12\x39\xb3\x61\xb3\x29\x44\xee\xdd\x25\x91\x83\x28\x16\x82\xdd\xce\x07\x3b\xdf\xc4\x94\xab\x61\xa7\xec\x8c\xf7\x12\xea\xc2\x73\x27\x77\x02\x44\x03\xeb\x29\x86\x7f\xcc\xc9\x08\xc8\x2a\x17\xf6\xb6\x91\x3f\x02\xec\x96\xb9\xe7\x83\x69\x4c\xe7\x86\x16\xf1\xce\xa2\xa4\x8d\x11\x10\xb8\x78\x8d\xf9\x85\xae\x44\xed\x19\x4f\x69\xf7\x0d\xbd\x5f\xa8\xa7\x4f\xf2\x09\xc5\x9c\xd5\xf2\xf3\xa0\x63\x3b\xdf\x46\x79\x09\x3c\x16\x27\x77\x1d\xf2\x4a\x8f\x00\x00\x00\xff\xff\x84\xa9\x69\x25\x4c\x01\x00\x00")

func gatewaySchemaModuleGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_gatewaySchemaModuleGraphql,
		"gateway/schema/module.graphql",
	)
}

func gatewaySchemaModuleGraphql() (*asset, error) {
	bytes, err := gatewaySchemaModuleGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gateway/schema/module.graphql", size: 332, mode: os.FileMode(420), modTime: time.Unix(1606552341, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gatewaySchemaProductGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x28\xca\x4f\x29\x4d\x2e\x51\xa8\xe6\x52\x50\xc8\x4c\xb1\x52\xf0\xcc\x2b\x51\xe4\x52\x50\x48\xce\x4f\x49\xb5\x52\x08\x2e\x29\xca\xcc\x4b\xe7\x52\x50\xc8\x4b\xcc\x85\x73\x41\xd2\x29\xa9\xc5\xc9\x45\x99\x05\x25\x99\xf9\x79\x70\x55\xb5\x5c\x5c\x99\x79\x05\xa5\x25\x0a\xce\x45\xa9\x89\x25\xa9\x50\x83\x3d\xc1\x42\xd5\x14\x1a\x19\x5a\x90\x82\xcd\x48\x4a\x1d\x0c\x08\x00\x00\xff\xff\x89\x1c\x20\x04\x02\x01\x00\x00")

func gatewaySchemaProductGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_gatewaySchemaProductGraphql,
		"gateway/schema/product.graphql",
	)
}

func gatewaySchemaProductGraphql() (*asset, error) {
	bytes, err := gatewaySchemaProductGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gateway/schema/product.graphql", size: 258, mode: os.FileMode(420), modTime: time.Unix(1606548274, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gatewaySchemaSchemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\xb1\xae\x83\x20\x14\x86\x77\x9e\xe2\xb0\xe9\x2b\xb0\xde\xc9\xc1\xc4\xdb\xd6\xc9\x38\x18\x3c\x49\x6d\xaa\x50\x3c\xb4\x21\x8d\xef\xde\x00\x16\x6d\xda\x09\xce\x17\xbe\xff\x07\xc8\x69\x84\xa3\x3c\xe3\xd8\xc1\x93\x01\xdc\x2c\x1a\x27\xe0\xdf\x2f\x0c\x60\xb4\xd4\xd1\xa0\x26\x01\xe5\xba\x63\x0b\x63\x41\x0a\x47\x82\x73\x1f\xf0\x81\x26\xcb\x05\xd4\x33\x1a\xee\x35\xd5\xdb\x2b\x9e\x9c\xc6\xd9\xe3\xa6\x4c\x33\x6f\x79\x4a\x78\x47\x86\x10\x69\xb0\x23\xac\x8c\xea\xad\xa4\x6c\x98\xb4\x25\x01\x7f\x7b\x58\x78\xc6\x73\x01\xeb\xec\x8b\xac\xee\xbf\xad\x7a\x0f\x7f\x59\xb1\x2b\x5e\xea\xb3\x2a\xb2\xe4\xc4\x71\x2b\xf2\xef\xab\x8c\xba\xa0\xa4\x83\xda\xdc\xa6\x98\x88\xb7\x3c\x7d\xc0\xc2\x5e\x01\x00\x00\xff\xff\x0e\xe7\x71\x71\x57\x01\x00\x00")

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

	info := bindataFileInfo{name: "gateway/schema/schema.graphql", size: 343, mode: os.FileMode(420), modTime: time.Unix(1606552413, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gatewaySchemaUserGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x2d\x4e\x2d\x52\xa8\xe6\x52\x50\xc8\x4c\xb1\x52\x08\x2e\x29\xca\xcc\x4b\x57\xe4\x52\x50\x28\x2d\x4e\x2d\xca\x4b\xcc\x4d\x45\x16\x4b\xcd\x4d\xcc\xcc\x41\x16\x48\x2b\xcd\xc9\x41\x56\xc4\xa5\xa0\x90\x9c\x88\x26\x54\xcb\x05\x08\x00\x00\xff\xff\x17\xb9\x5b\xab\x67\x00\x00\x00")

func gatewaySchemaUserGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_gatewaySchemaUserGraphql,
		"gateway/schema/user.graphql",
	)
}

func gatewaySchemaUserGraphql() (*asset, error) {
	bytes, err := gatewaySchemaUserGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gateway/schema/user.graphql", size: 103, mode: os.FileMode(420), modTime: time.Unix(1606548231, 0)}
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
	"gateway/schema/module.graphql":  gatewaySchemaModuleGraphql,
	"gateway/schema/product.graphql": gatewaySchemaProductGraphql,
	"gateway/schema/schema.graphql":  gatewaySchemaSchemaGraphql,
	"gateway/schema/user.graphql":    gatewaySchemaUserGraphql,
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
			"module.graphql":  &bintree{gatewaySchemaModuleGraphql, map[string]*bintree{}},
			"product.graphql": &bintree{gatewaySchemaProductGraphql, map[string]*bintree{}},
			"schema.graphql":  &bintree{gatewaySchemaSchemaGraphql, map[string]*bintree{}},
			"user.graphql":    &bintree{gatewaySchemaUserGraphql, map[string]*bintree{}},
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
