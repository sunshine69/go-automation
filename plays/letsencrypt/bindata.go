// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// inventory-letsencrypt/files/README.md
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

var _inventoryLetsencryptFilesReadmeMd = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x4f\xcd\x4b\x2d\x4a\x2c\x49\x4d\xd1\x51\x28\x4a\x2d\xce\x2f\x2d\x4a\x4e\x2d\x56\x48\xcb\xcc\x49\x2d\x56\x48\x2c\x4a\x55\xc8\xcc\x53\xc8\x48\x2d\x4a\x05\x04\x00\x00\xff\xff\x5e\x91\xa2\x87\x26\x00\x00\x00"

func inventoryLetsencryptFilesReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_inventoryLetsencryptFilesReadmeMd,
		"inventory-letsencrypt/files/README.md",
	)
}

func inventoryLetsencryptFilesReadmeMd() (*asset, error) {
	bytes, err := inventoryLetsencryptFilesReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "inventory-letsencrypt/files/README.md", size: 38, mode: os.FileMode(436), modTime: time.Unix(1768695790, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _inventoryLetsencryptGroup_varsAllYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\x51\x8b\xe3\x36\x10\x7e\xf7\xaf\xf8\xd8\x3c\xb8\x85\x5d\x27\xb9\xdb\x96\x5b\x43\x4b\xa1\xa5\xb4\xb4\x70\x85\x2e\xbd\x7d\x33\x8a\x34\x8e\xc5\xca\x92\x91\xc6\x9b\x35\x69\xfe\x7b\x19\xd9\x49\xae\xdb\x97\x72\x4f\x36\xa3\xd1\x68\xbe\xf9\xbe\x6f\x56\x18\x62\x78\xb1\x86\x0c\x14\x43\x87\xbe\x57\xde\xc0\x59\x4f\x50\x9a\x6d\xf0\xdf\x3d\x3d\x3d\xe1\xd0\x51\x24\xc8\x9f\x56\x1e\x3b\x82\x8e\xa4\x98\x9a\x31\x51\xac\xf0\xb3\x8d\x89\xc1\xb6\x27\xc4\xd1\xdf\x42\x19\x03\x15\xf7\x28\x97\x0a\x9f\x25\x97\xc5\x0a\x1c\x96\xeb\x08\x9e\x8a\x39\xa7\x46\x29\x47\x8f\x1f\x7f\xfa\x88\x1f\x3b\xe5\xf7\x04\xee\x6c\x2a\x94\xd6\x61\xf4\xdc\x50\xaf\xac\xab\x51\x4e\x61\xfc\x61\x0a\x63\x34\xa1\x57\xd6\xcb\x15\xdb\xc2\x32\x4c\xa0\x04\x1f\x18\xf4\x6a\x13\x4b\xe4\x60\x9d\x93\x4e\xf7\xe4\xa3\x62\x32\x68\x2f\x5d\x16\xd2\x49\xf3\x4c\x53\x33\x28\xee\x6a\x94\xc7\x23\xac\x7f\x21\xcf\x21\x4e\x8d\xb1\x11\xa7\xd3\xba\xb5\x8e\xd2\x3a\x03\x7c\xa6\x49\x5e\xfa\x95\x61\x13\xd4\xc8\x48\xc4\xb0\x1e\x8e\xf6\x01\xce\xee\x30\x50\x04\xf9\x97\xa5\x17\x3b\x77\x22\x49\x21\xe6\x0f\x07\x50\x3f\xf0\x54\x68\x25\xe5\x9b\x31\xba\x1a\x37\x1d\xf3\x50\xaf\xd7\xdb\x87\x77\xd5\xf6\xdb\x0f\xd5\xc3\x43\xb5\xdd\x6c\xea\xfb\xcd\x66\xb3\x36\x36\x92\x96\x6e\x6e\xce\x43\xd1\xf3\x50\x6c\x2e\x96\x46\x79\xa6\xc5\x14\x46\x8c\x89\x60\x7c\xda\x6c\x71\x87\x44\x84\x1d\xb9\x70\x28\x7a\x15\x95\xf1\xa9\xd1\xc1\xb7\x76\xdf\x08\x98\x1a\xe5\x9a\x58\xaf\x97\xa3\xf3\xb7\x0a\x71\x5f\x5e\xf2\x23\xb9\xa0\x4c\xa3\x7b\x53\xa3\x4c\x53\x62\xea\x35\x3b\xcc\x61\x2c\x59\x32\x0c\x1d\x7a\x85\x44\x03\x9c\xcc\x3b\xb4\x98\x19\xb9\x5d\xc6\x1c\x3c\xc9\x1c\xb8\x23\x0c\xd1\x7a\x6d\x07\x97\xe9\xc6\xa1\xb3\xba\x43\xa7\x52\xd6\x5a\xf0\xf0\xaa\xa7\xea\x0d\xf5\x57\x94\x54\xcc\x75\x6b\x94\xf4\xaa\xfa\xc1\x51\xa5\x43\x9f\xd9\x68\xaf\x7c\xa7\x37\x84\xd3\x7f\x18\x1f\xa2\x7d\x11\x0d\xfe\x2f\xd2\x13\xc5\x17\x8a\x77\xc7\x63\x26\xf5\x74\x9a\x05\x30\x8c\x3b\x67\xf5\x17\x56\xd0\x91\xb3\xf4\xa7\x81\x64\x58\xba\x53\xce\x91\xdf\x53\x85\xbf\x94\x1b\x29\xd5\x33\x87\xb7\xa2\x89\xcd\xb6\xc2\xa7\x8e\x3c\xc6\x64\xfd\x1e\x73\x48\xc0\x8a\x27\x14\x17\x2b\x7c\x22\xa4\x81\xb4\x6d\x27\x28\x8f\x5f\x1e\x1f\xff\xc0\x10\x62\xa6\xe1\x9b\xcd\xe6\x1d\xc4\xc2\xca\xe3\xf1\xf7\x3f\x3f\x8f\x6f\x11\x3c\x94\x73\xb0\x9e\x29\xb6\x4a\x53\x2a\x56\xd8\x91\x56\xa2\xa1\x03\x41\x45\xf2\x25\x8b\x87\xbd\x3c\xac\x12\x62\x08\x9c\xab\x69\x25\x27\x3b\x2b\x85\x33\xe5\x32\x64\x21\x29\xd7\xff\xb0\xc9\x49\xf7\xf7\xef\x8b\x15\xbe\x1a\x13\x19\x38\xc5\x14\x65\x6f\xf8\x5c\x99\x59\x0c\x90\x2f\xa8\x94\xae\xf8\xd3\xd7\x15\x7e\x23\x1a\xc4\x4e\xbd\x54\x17\x84\x59\xd7\x89\xad\x73\xc5\x0a\x9e\xc8\xe4\x7b\x31\xbc\x4e\xd7\x8b\xe0\xa8\xda\xd6\x6a\xb4\x31\xf4\xd2\xc0\xdd\xf7\x57\xec\x39\x76\x7f\xff\x7e\x09\x6e\xe7\x80\xec\x0d\xc4\x30\x4a\x63\x67\xd4\x8e\x38\x91\xd7\x71\x1a\x64\xae\x97\x2d\xa2\x9c\x0b\x07\x28\x3f\x21\x70\x47\x71\x46\xc9\x02\xe6\x5f\x50\x67\xdf\xa9\x24\x03\xf6\xe1\xb0\x98\xbf\x53\xd1\xe8\x60\xe6\xb6\x17\xcb\x9c\xb7\x6c\xbc\xac\xdb\xb7\x9a\xcf\xeb\xee\x02\xaf\x39\xe7\xd7\x28\xf3\x2b\xa2\x9e\x38\xfa\xcb\x86\x56\xad\xc0\x10\x36\x8d\x6c\xd1\x21\x24\x6e\x96\xb3\x54\xe3\xef\x02\x20\xdd\x05\xdc\xc8\x69\x05\x13\x04\x0c\x77\x59\x4f\xb2\xca\x9d\x7d\x26\xe8\x30\x4c\xd0\x14\x65\x99\x89\xdf\xb2\x68\x6f\xcf\x66\x27\xd6\x37\x05\xb0\x82\x1e\x70\x3c\xe2\xad\x83\x70\x3a\x21\xef\x13\xbf\xb7\xfe\x75\x9d\x92\x5b\x54\x2f\x6e\xf9\x27\x00\x00\xff\xff\xf2\x99\x23\xb9\x56\x06\x00\x00"

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

	info := bindataFileInfo{name: "inventory-letsencrypt/group_vars/all.yaml", size: 1622, mode: os.FileMode(436), modTime: time.Unix(1768709102, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _inventoryLetsencryptGroup_varsProdYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x4e\x8c\x4f\xc9\x2c\x8a\x2f\x2d\xca\xb1\x52\x50\x57\x07\x04\x00\x00\xff\xff\x8b\x45\x4d\x5f\x0e\x00\x00\x00"

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

	info := bindataFileInfo{name: "inventory-letsencrypt/group_vars/prod.yaml", size: 14, mode: os.FileMode(436), modTime: time.Unix(1768692753, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _inventoryLetsencryptGroup_varsUatYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x4e\x8c\x4f\xc9\x2c\x8a\x2f\x2d\xca\xb1\x52\x50\x57\x07\x04\x00\x00\xff\xff\x8b\x45\x4d\x5f\x0e\x00\x00\x00"

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

	info := bindataFileInfo{name: "inventory-letsencrypt/group_vars/uat.yaml", size: 14, mode: os.FileMode(436), modTime: time.Unix(1768692757, 0)}
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
	"inventory-letsencrypt/files/README.md":      inventoryLetsencryptFilesReadmeMd,
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
		"files": &bintree{nil, map[string]*bintree{
			"README.md": &bintree{inventoryLetsencryptFilesReadmeMd, map[string]*bintree{}},
		}},
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
