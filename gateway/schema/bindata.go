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

var _gatewaySchemaModuleGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x8f\xb1\xae\xc2\x30\x0c\x45\xf7\x7c\xc5\x7d\xea\x5f\xbc\x95\xa9\x03\x0b\xb0\xa3\x28\xb1\x50\x50\x6b\x07\xd7\x19\x10\xe2\xdf\x51\x8a\x2a\xa5\xa2\x03\x12\x5b\x7c\x14\xdf\x73\x9d\x38\x17\xc3\x4e\xc9\x1b\xed\x25\x96\x81\xfa\x99\x3c\x1c\x20\x1a\x49\xcf\x29\xfe\xa3\x67\x73\x40\x56\xb9\x52\xb0\x85\xfc\x39\xc0\xee\x99\xe6\xf9\x68\x9a\xf8\x52\xd1\x20\xc1\x5b\x12\x5e\x98\x03\x22\x4d\x41\x53\x5e\xd1\xa7\x73\x75\x19\x6f\xe9\xec\x6b\x72\xbf\x50\x77\x5b\xf2\x0e\x93\x79\x2b\xd3\xcf\x85\x4e\xf5\x59\x4b\x05\x89\xd4\xa6\xb1\x1f\x9b\x79\x59\x3a\xc8\xe7\x0d\xeb\x9f\x80\xd2\xad\x24\xa5\x91\xd8\x5a\xbc\xd1\xa5\xe6\xbe\x02\x00\x00\xff\xff\x87\x25\xf3\x03\x99\x01\x00\x00")

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

	info := bindataFileInfo{name: "gateway/schema/module.graphql", size: 409, mode: os.FileMode(420), modTime: time.Unix(1606633605, 0)}
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

var _gatewaySchemaSchemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\xb1\x8e\x83\x30\x0c\x86\xf7\x3c\x85\xb3\xc1\x2b\x64\xbd\x89\x01\x89\xbb\x2b\x13\x62\x40\xc1\x52\xa9\x80\xa4\xc1\x69\x85\x2a\xde\xbd\x72\x02\x94\xaa\x9d\xc0\x9f\xf3\xf9\x8f\x43\xb3\x45\xf8\xd7\x67\x1c\x1a\x78\x08\x80\xab\x47\x37\x2b\xf8\xe5\x8f\x00\x18\x3c\x35\xd4\x99\x51\x41\xbe\xfe\x89\x45\x88\x20\x85\x23\xc1\xb9\x75\x78\x47\x97\xa4\x0a\xca\x09\x9d\x64\xcd\xb4\xbe\xc7\xd3\x6c\x71\x62\x5c\xe5\x7b\x2d\x6b\xee\x3b\xd3\xaf\x9d\x3f\xd3\x07\xb6\x4d\xdd\x62\xc2\x60\xed\xb0\x21\x2c\x9c\x69\xbd\xa6\xa4\x1b\xad\x27\x05\x3f\x47\x98\x31\x93\xa9\x82\xb5\xe6\xe1\xde\xb6\x9f\x56\x79\x84\xdf\xac\x98\x15\x2f\xfa\x1e\x15\xd9\xee\xc4\xf2\x15\xc4\x3b\x17\xce\x5c\x50\x13\x2f\xb3\xb9\x55\x36\x92\xac\xe5\xfe\x28\x8b\x78\x06\x00\x00\xff\xff\x8b\xf1\x3c\xbd\x6b\x01\x00\x00")

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

	info := bindataFileInfo{name: "gateway/schema/schema.graphql", size: 363, mode: os.FileMode(420), modTime: time.Unix(1606631379, 0)}
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
