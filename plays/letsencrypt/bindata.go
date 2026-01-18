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

var _inventoryLetsencryptGroup_varsAllYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x55\x4d\x8f\xe4\x34\x10\xbd\xe7\x57\x3c\x4d\x1f\xb2\x2b\xcd\xa6\xd3\xbb\x03\xda\x8d\x04\x42\x2c\x42\xac\x40\x5a\x04\x23\x76\x6f\x91\xdb\xae\x74\xac\x71\xec\x60\x57\x7a\x26\x0c\xf3\xdf\x51\x39\xfd\xc5\xc0\x01\x71\x9a\x9e\x72\xb9\x52\xf5\xea\xbd\xe7\x15\xc6\x18\xf6\xd6\x90\x81\x62\xe8\x30\x0c\xca\x1b\x38\xeb\x09\x4a\xb3\x0d\xfe\xab\xcf\x9f\x3f\xe3\xbe\xa7\x48\x90\x5f\x5a\x79\x6c\x09\x3a\x92\x62\x6a\xa7\x44\xb1\xc2\xf7\x36\x26\x06\xdb\x81\x10\x27\x7f\x0d\x65\x0c\x54\xdc\xa1\x3c\x54\xb8\x48\x2e\x8b\x15\x38\x1c\xae\x23\x78\x2a\x96\x9c\x06\x65\x59\x14\x2b\xdc\x7e\xfc\xee\x23\xde\xf7\xca\xef\x08\xdc\xdb\x54\x28\xad\xc3\xe4\xb9\xa5\x41\x59\xd7\xa0\x9c\xc3\xf4\xcd\x1c\xa6\x68\xc2\xa0\xac\x97\x72\xb6\x83\x65\x98\x40\x09\x3e\x30\xe8\xc1\x26\x96\xc8\xbd\x75\x4e\x5a\xdd\x91\x8f\x8a\xc9\xa0\x3b\xb5\x59\x48\x2b\xed\x1d\xcd\xed\xa8\xb8\x6f\x50\x3e\x3e\xc2\xfa\x3d\x79\x0e\x71\x6e\x8d\x8d\x78\x7a\x5a\x77\xd6\x51\x5a\xe7\x09\xef\x68\xce\xdd\x7d\xe8\x90\x88\x65\x00\x1a\x46\x9e\x71\x4f\xcb\x57\xd4\xc4\x01\xa3\xd5\x77\x98\x46\x74\x31\x0c\x70\xb4\x0b\x70\x76\x5b\xe1\xb6\x27\xd0\x83\x1a\x46\x47\xd8\x92\x0b\xf7\xb0\x09\x5d\x88\x30\xb4\x17\xb8\xbc\xf5\x3b\x98\xa0\xef\x28\x42\x07\xcf\xca\x7a\x8a\xf8\x36\x4c\xce\x50\x2c\xb4\x92\x76\xda\x29\xba\x06\x57\x3d\xf3\xd8\xac\xd7\x9b\x77\xaf\xab\xcd\x97\x6f\xab\x77\xef\xaa\x4d\x5d\x37\x37\x75\x5d\xaf\x8d\x8d\xa4\xa5\xfb\xab\x13\x8a\x7a\x41\xd1\xe6\x7e\xd3\x64\x59\x90\x9a\xc3\x84\x29\x11\x8c\x4f\xf5\x06\xaf\x90\xe8\xd0\x55\x31\xa8\xa8\x8c\x4f\xad\x0e\xbe\xb3\xbb\x56\xa6\x6f\x50\xae\x89\xf5\xfa\x70\x74\xfc\x5b\x85\xb8\x2b\x4f\xf9\x91\x5c\x50\xa6\xd5\x83\x69\x50\xa6\x39\x31\x0d\x9a\x1d\x96\x30\x0e\x59\x19\x3d\x1d\x06\x85\x44\x23\x9c\x6c\x28\x74\x58\x76\x78\x7d\x58\x4c\xf0\x24\xd0\x70\x4f\x18\xa3\xf5\xda\x0a\x62\x12\xbc\xef\xad\xee\xd1\xab\x94\xe9\x19\x3c\xbc\x1a\xa8\xc2\x2f\x34\xd0\xb0\xa5\x78\x39\xd6\xbd\x75\x46\xab\x68\xa0\x29\x32\x5e\x58\x9f\x98\x94\x41\xe8\x5e\x16\x2b\xa4\x30\x10\xf7\x02\x77\x9e\xf8\x5a\x3e\xe5\xf3\xd5\x5e\xed\x49\x50\x3a\x23\x13\x29\x05\xb7\xa7\x08\x95\xe0\x88\x13\x79\x1d\xe7\x51\x9a\x74\x33\x94\xd6\x34\x0a\xa3\x74\x1f\xac\xa6\x67\xb4\x3d\x03\x4e\xc5\x32\x61\x83\xf2\xc0\x80\x4a\x87\xe1\xfa\xd4\x47\x75\x11\xcd\xc2\xe8\x09\x53\x74\xb0\x5e\x08\x32\x28\xc6\x45\x42\x73\x73\xf3\x26\x4b\xa7\x27\x7d\x97\x53\xe9\x61\xb4\x71\x86\x51\x4c\xd5\x91\x9a\x47\xde\x2f\x59\x5b\xea\x42\x24\x28\x66\xe1\xab\x4c\xce\x01\x91\x7e\x9f\x28\x71\x86\xa8\x2a\x72\x62\x2b\xbf\x17\x92\x95\x65\x61\xd4\x9c\x5a\x0e\x6d\xae\x4f\x47\x65\x7e\xe8\x90\x15\x71\x16\x59\x7a\xa6\x32\xfa\x87\xcc\xc6\x68\xf7\xa2\xfc\xff\xa4\xb4\x44\x71\x4f\xf1\xd5\xe3\x23\xc8\xef\xf1\xf4\xb4\xa8\x6e\x9c\xb6\xce\xea\xff\x59\x41\x47\xce\xbd\xf3\x3c\x92\x10\x4e\xf7\xca\x39\xf2\x3b\xaa\xf0\x9b\x72\x13\xa5\x66\x59\xf7\xb5\x28\xab\xde\x54\xf8\x24\x8c\x98\x92\x20\xb5\x84\x64\x5a\x71\x22\xc5\xc5\x0a\x9f\x08\x69\x24\x6d\xbb\x19\xca\xe3\x87\xdb\xdb\x9f\x31\x86\x98\xa9\xfc\x45\x5d\xbf\x86\x38\xa7\xf2\xb8\xfd\xe9\xd7\xcb\xf8\x06\xc1\x43\x39\x59\x2b\x53\xec\x94\xa6\x54\xac\xb0\x25\xad\x32\x63\x09\x2a\x92\x2f\xf9\xe4\x05\x2a\x21\x86\xc0\xb9\x9a\x56\x72\xb2\xb5\x52\x38\xcb\x46\x50\x96\x25\xe6\xfa\x6f\xeb\x9c\x74\x73\xf3\xa6\x58\xe1\xc5\x94\xc8\xc0\x29\xa6\x28\x76\xed\x73\xe5\x65\xf1\xf9\x82\x4a\xe9\x3c\x7f\x7a\x59\xe1\x47\xa2\x51\xa8\x36\x48\x75\x99\x30\x4b\x21\xb1\x75\xae\x58\xc1\x13\x99\x7c\x2f\x86\x87\xf9\x7c\x11\x1c\x55\xd7\x59\xbd\xb8\xdc\xdb\x1a\xaf\xbe\x3e\xcf\x9e\x63\x42\xd3\x25\xb8\x59\x02\xe2\xd6\x88\x61\x92\xc6\x8e\x53\x5f\x48\xaa\x58\x9d\xbd\x5b\x39\xf1\x47\xe5\x67\x04\xee\x29\x2e\x53\x66\x95\xfe\x6d\xd4\x45\xa2\x2a\x09\xc0\x5e\x0c\x95\xc5\x38\x7a\x15\x8d\x0e\x66\x69\xfb\xe0\x3b\xc7\xc7\x2d\x9e\x5e\xb9\xea\xdf\x1e\x99\xd3\x78\xed\x31\xbf\x41\x99\xbf\x22\xb2\x7c\x3f\x25\x0e\x83\xfd\x83\xcc\xd2\x90\x98\xf7\x81\x1d\x21\x82\x5d\xaa\x37\x85\xfc\xdf\xca\x69\x83\x2b\xc1\xe3\x2a\x47\xd2\x45\x68\x93\x9d\x39\x4e\xfe\xf4\xc8\xaa\x4e\x20\x11\x66\x18\x79\x08\xc7\x90\xb8\x3d\x9c\x35\xf8\xb3\x00\x48\xf7\x01\x57\x72\x58\xc1\x04\xc1\x65\xb1\xaf\xfc\x18\x3b\x7b\x47\xd0\x61\x9c\x17\xaf\xb3\x5e\x4c\x27\x0b\xe0\xfa\x68\xbe\xc4\xfa\xaa\x00\x56\xd0\x23\x1e\x1f\xf1\x5c\x8d\x78\x7a\x42\xf6\x77\xbf\xb3\xfe\x61\x9d\x92\x3b\x28\x48\x94\xf7\x57\x00\x00\x00\xff\xff\x18\x82\xef\x76\x18\x08\x00\x00"

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

	info := bindataFileInfo{name: "inventory-letsencrypt/group_vars/all.yaml", size: 2072, mode: os.FileMode(436), modTime: time.Unix(1768715856, 0)}
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
