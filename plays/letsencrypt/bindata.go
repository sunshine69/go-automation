// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// inventory-letsencrypt/group_vars/all.yaml
// inventory-letsencrypt/group_vars/prod.yaml
// inventory-letsencrypt/group_vars/uat.yaml
// inventory-letsencrypt/hosts.ini
package main

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _inventoryLetsencryptGroup_varsAllYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xc1\x8a\xe3\x30\x0c\x86\xef\x7e\x0a\xd1\x1e\x7a\xd9\x4d\xd2\x65\x59\xb6\x86\x79\x80\x79\x83\xde\x8c\x6a\x2b\x8d\xa8\x63\x1b\x59\x49\xa7\x94\xbe\xfb\x60\xda\x99\xe3\x9c\x24\xfe\xff\x43\xf0\x69\x0b\x45\xf2\xca\x81\x02\xa0\x82\xcf\xf3\x8c\x29\x40\xe4\x44\x80\x5e\x39\xa7\xb7\xe3\xf1\x08\xd7\x89\x84\xa0\x6d\x1e\x13\x9c\x08\xbc\x10\x2a\xb9\xa5\x92\x98\x27\x67\x61\xb7\x33\xe8\x7d\x5e\x92\x3a\x9a\x91\xa3\x35\xad\x76\x17\xba\xb9\x82\x3a\x59\xe3\xd1\x05\x16\xb7\x48\xb4\xb0\x99\x54\x8b\xed\xfb\xfd\xe1\x4f\xb7\xff\xf7\xbf\x3b\x1c\xba\xfd\x30\xd8\xbf\xc3\x30\xf4\x81\x85\xbc\x66\xb9\x6d\xcc\x8c\x82\x21\x55\xe7\x73\x1a\xf9\xec\x46\x8e\x64\xa1\x27\xf5\xfd\xab\xf9\x9a\x5d\x96\xb3\xd9\x36\x01\x84\x4a\x05\x22\x57\x85\x3c\x42\xc8\x33\x72\xfa\x05\x23\x4b\x0b\x12\x01\x57\xd0\x89\xa0\x08\x27\xcf\x25\x52\x0b\xcd\x13\xb3\x66\x0b\xef\x23\xa4\xac\x40\x1f\x5c\xb5\x02\x2b\x5c\x39\xc6\xa6\x7c\xa6\x44\x82\x4a\xe1\x75\x4b\x79\x26\x53\x84\xd7\xf6\x88\x6f\x49\xa8\x24\x2b\xc9\xef\xfb\x1d\x28\xad\xf0\x78\x74\x17\xba\x99\xb2\x9c\x22\xfb\x9f\x28\x2f\xfa\x19\x00\x00\xff\xff\x76\xb9\xc4\x06\x8b\x01\x00\x00"

func inventoryLetsencryptGroup_varsAllYamlBytes() ([]byte, error) {
	return bindataRead(
		_inventoryLetsencryptGroup_varsAllYaml,
		"inventory-letsencrypt/group_vars/all.yaml",
	)
}

func inventoryLetsencryptGroup_varsAllYaml() (*asset, error) {
	bytes, err := inventoryLetsencryptGroup_varsAllYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "inventory-letsencrypt/group_vars/all.yaml", size: 395, mode: os.FileMode(436), modTime: time.Unix(1768634940, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _inventoryLetsencryptGroup_varsProdYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x4e\x8c\x4f\xc9\x2c\x8a\x2f\x2d\xca\xb1\x02\x04\x00\x00\xff\xff\x12\xe2\x89\x83\x0b\x00\x00\x00"

func inventoryLetsencryptGroup_varsProdYamlBytes() ([]byte, error) {
	return bindataRead(
		_inventoryLetsencryptGroup_varsProdYaml,
		"inventory-letsencrypt/group_vars/prod.yaml",
	)
}

func inventoryLetsencryptGroup_varsProdYaml() (*asset, error) {
	bytes, err := inventoryLetsencryptGroup_varsProdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "inventory-letsencrypt/group_vars/prod.yaml", size: 11, mode: os.FileMode(436), modTime: time.Unix(1768634863, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _inventoryLetsencryptGroup_varsUatYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x4e\x8c\x4f\xc9\x2c\x8a\x2f\x2d\xca\xb1\x02\x04\x00\x00\xff\xff\x12\xe2\x89\x83\x0b\x00\x00\x00"

func inventoryLetsencryptGroup_varsUatYamlBytes() ([]byte, error) {
	return bindataRead(
		_inventoryLetsencryptGroup_varsUatYaml,
		"inventory-letsencrypt/group_vars/uat.yaml",
	)
}

func inventoryLetsencryptGroup_varsUatYaml() (*asset, error) {
	bytes, err := inventoryLetsencryptGroup_varsUatYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "inventory-letsencrypt/group_vars/uat.yaml", size: 11, mode: os.FileMode(436), modTime: time.Unix(1768634856, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _inventoryLetsencryptHostsIni = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\x41\xaa\xc3\x30\x0c\x44\xf7\x3a\x85\x20\xb7\xf8\xf0\x37\xbd\x86\x09\xc1\x44\x43\x09\x18\xd9\xd8\xb2\xa1\xb7\x2f\x4a\xbb\x08\xc5\x59\x6a\x1e\x4f\x33\xbd\x48\x34\x6c\x09\xd6\xa0\x7b\x7d\x15\xdb\x04\x83\x17\x36\x34\xe3\xde\x0e\x7d\x72\xca\x7b\x4c\xfc\xc8\x3d\x09\x2a\x4d\x8c\x1e\xed\xc7\x80\x35\xfe\xd2\x4f\x0e\x95\x92\x0f\xb5\x99\x5e\x6a\x16\x5e\xb8\x22\x26\xce\x0a\xa2\x70\xa1\xeb\xcc\x10\x8c\x9b\x1d\x77\xff\x89\x82\x60\xdc\x3e\x73\xf8\x37\x62\x6d\x2b\x41\xc7\xbf\x47\x14\x7a\x9c\xb7\x7b\x8d\xc3\x8b\xe0\x11\x05\x2f\x9a\x1a\xe7\x82\x13\x5f\x1c\x3f\xdf\x01\x00\x00\xff\xff\x0f\xdf\x61\x79\x7f\x01\x00\x00"

func inventoryLetsencryptHostsIniBytes() ([]byte, error) {
	return bindataRead(
		_inventoryLetsencryptHostsIni,
		"inventory-letsencrypt/hosts.ini",
	)
}

func inventoryLetsencryptHostsIni() (*asset, error) {
	bytes, err := inventoryLetsencryptHostsIniBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "inventory-letsencrypt/hosts.ini", size: 383, mode: os.FileMode(436), modTime: time.Unix(1768635067, 0)}
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
	"inventory-letsencrypt/group_vars/all.yaml":  inventoryLetsencryptGroup_varsAllYaml,
	"inventory-letsencrypt/group_vars/prod.yaml": inventoryLetsencryptGroup_varsProdYaml,
	"inventory-letsencrypt/group_vars/uat.yaml":  inventoryLetsencryptGroup_varsUatYaml,
	"inventory-letsencrypt/hosts.ini":            inventoryLetsencryptHostsIni,
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
	"inventory-letsencrypt": &bintree{nil, map[string]*bintree{
		"group_vars": &bintree{nil, map[string]*bintree{
			"all.yaml":  &bintree{inventoryLetsencryptGroup_varsAllYaml, map[string]*bintree{}},
			"prod.yaml": &bintree{inventoryLetsencryptGroup_varsProdYaml, map[string]*bintree{}},
			"uat.yaml":  &bintree{inventoryLetsencryptGroup_varsUatYaml, map[string]*bintree{}},
		}},
		"hosts.ini": &bintree{inventoryLetsencryptHostsIni, map[string]*bintree{}},
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
