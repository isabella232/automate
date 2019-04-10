// Code generated by go-bindata. DO NOT EDIT.
// Code generated by go-bindata.
// sources:
// data/keys.json
// DO NOT EDIT!

package keys

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

var _dataKeysJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xcf\x4d\x6f\x82\x30\x00\x80\xe1\xbf\x62\xb8\x92\xa5\x01\x1b\x90\x25\x3b\xb4\xd0\x16\xd4\xf2\x51\x30\xca\xd6\x1d\x26\x7e\xa0\x98\x0e\x32\x94\xe2\xb2\xff\xbe\xcc\xeb\xde\xe3\x7b\x7b\xbe\x8d\x66\x3f\x7e\x19\xcf\x6f\xc6\xd3\x5f\x98\xb0\x28\x9e\xa4\x2b\xbc\x8c\xfc\xc9\x82\x94\x8f\x29\x15\x8f\xd8\x96\x63\xc4\xf0\xd8\xb1\x9c\x43\x0f\x1d\x09\xc3\xf9\x15\x63\x84\xce\x08\x32\x86\x50\x86\xe9\x89\x58\xa5\xcb\x3e\x2d\x51\xbd\x2e\xe1\x09\xd9\xf7\x34\xba\xac\x47\x70\xe4\xdd\x20\x95\x47\x43\xdf\x3c\x6f\xdb\xdd\xcd\x86\xda\x24\x49\xdc\x0d\xb5\x55\x52\x67\xd0\xd5\x66\x1d\x53\xad\x8b\x53\x0b\x9c\xad\xeb\xd7\xd3\x65\x60\xd5\x01\x77\x33\x67\xae\x03\x17\x79\x04\x1d\xfa\x43\x2a\x55\x38\x85\xae\x39\xe3\x8b\xf8\x8e\xeb\x7d\xd3\x27\xf3\xe6\x03\xdc\x0e\x45\x79\x39\xce\x12\xea\xf1\x2b\x4d\xc4\xbe\x07\x74\xd3\xb6\x36\xd4\x61\xb1\x2b\xa6\x73\x24\xc4\x08\xf2\x2c\x2c\xc5\xe0\x53\xa9\x02\x67\x65\x2e\x44\x4e\x53\xaf\x02\xc5\x26\xcf\x06\xbb\x7a\x91\xea\x41\x24\x71\xf0\x5f\x2d\x95\xf1\xfe\xf3\x1b\x00\x00\xff\xff\xec\x71\xae\x56\x21\x01\x00\x00")

func dataKeysJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataKeysJson,
		"data/keys.json",
	)
}

func dataKeysJson() (*asset, error) {
	bytes, err := dataKeysJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/keys.json", size: 289, mode: os.FileMode(420), modTime: time.Unix(1518742631, 0)}
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
	"data/keys.json": dataKeysJson,
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
	"data": &bintree{nil, map[string]*bintree{
		"keys.json": &bintree{dataKeysJson, map[string]*bintree{}},
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
