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

var _gatewaySchemaSchemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x51\xcb\x4e\xc3\x30\x10\xbc\xfb\x2b\x36\xea\xa5\xfd\x05\x5f\x39\xf5\x50\xa9\x50\xf5\x84\x50\x65\xd9\x0b\x18\x25\x71\x58\xaf\x41\x15\xea\xbf\x23\xdb\x21\x75\x1e\x42\xa8\xe2\x94\x78\xb4\x33\xb3\x33\xcb\xe7\x0e\xe1\xa0\x5f\xb1\x51\xf0\x25\x00\xde\x03\xd2\x59\xc2\x7d\xfc\x08\x80\x26\xb0\x62\xeb\x5a\x09\xbb\xfe\x4f\x5c\x84\x48\xa4\x34\x92\x38\x1f\x16\x3f\x91\xd6\x1b\x09\x47\x8f\x54\x0d\x13\x3f\x94\x34\xa4\x09\x15\xe3\x9e\x9c\x09\x9a\xd7\xb6\xed\x02\x4b\xb8\x2b\xc1\x6d\xc4\xaa\x8d\x84\xfe\x5d\x09\x80\xd0\x99\x39\xeb\x58\x82\x4b\xac\xec\xb5\x73\x26\xd4\x38\xb6\xca\xd8\xc0\xc9\xcf\xab\x51\xdc\x7f\x4f\xee\x0d\x35\x3f\xb8\x2b\xf7\x71\xdb\x72\xf5\x54\xcd\x02\xc6\x57\x0a\x67\x8d\x84\x03\x93\x6d\x5f\x92\x96\x47\x6a\x55\x83\x25\x86\x8d\xb2\x75\x09\x3c\x87\xba\x2e\x87\xe2\xda\x6a\x02\x5d\x84\x48\x1b\x2c\xf4\x94\x3b\x75\xa6\xe4\x4f\x3d\x0d\x7a\x4d\xb6\xcb\xf7\x9b\x4a\xce\x4b\x1c\x92\xc4\xb4\x37\xab\xa7\x62\x7a\xd9\xff\x51\x2c\x2b\x28\xee\x97\xc4\x1d\x19\xa4\x53\x6f\x21\x00\xba\x7c\xbc\x53\x61\x1a\x17\x3a\x8d\x0f\x54\x3b\xad\x46\x2e\xbf\xa7\xc9\xa6\xd3\x30\x7f\xb0\x5e\x2d\x99\xaf\xc0\xb3\xe2\xe0\x6f\x5a\xe8\x3b\x00\x00\xff\xff\xaa\x17\xfc\x8b\xaf\x03\x00\x00")

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

	info := bindataFileInfo{name: "gateway/schema/schema.graphql", size: 943, mode: os.FileMode(420), modTime: time.Unix(1606268730, 0)}
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
