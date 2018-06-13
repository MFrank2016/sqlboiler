// Code generated by go-bindata.
// sources:
// override/templates/17_upsert.tpl
// override/templates/singleton/mssql_upsert.tpl
// override/templates_test/singleton/mssql_main_test.tpl
// override/templates_test/singleton/mssql_suites_test.tpl
// override/templates_test/upsert.tpl
// DO NOT EDIT!

package driver

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

var _templates17_upsertTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\xdd\x6e\xdb\xc6\x12\xbe\x26\x9f\x62\x62\x1c\xc4\xe4\x39\x32\x7d\x7a\xeb\x42\x17\x76\x9c\xa4\x41\x62\x57\x89\xe2\x06\xa8\x61\x18\x2b\x72\x28\x2d\xbc\xda\x65\x96\x4b\x39\xaa\xca\x77\x2f\x66\x77\x29\x52\xb2\x64\xc9\x49\x5d\xf4\xc2\xb0\xc9\x1d\xce\xcc\x7e\xdf\xfc\x7a\xb1\x38\x82\xff\x30\xc1\x59\x09\x27\x7d\x48\x4e\xe9\x2f\x2c\x93\xcf\x6c\x24\x10\xdc\xaf\xe4\x92\x4d\xb1\xae\x43\x2b\x5a\xa6\x13\x9c\x32\x77\x4c\x1f\xb4\x12\xf0\x27\x24\xc3\xf6\xd4\x7e\xc0\x73\x48\x4e\xb3\xec\xad\x50\x23\x26\xe0\xa8\xae\xc3\xe3\x63\xb8\x2a\x4a\xd4\xe6\x2d\x30\x63\x70\x5a\x98\x12\x98\x04\x2e\xe9\x5d\x0f\x98\xcc\x20\x53\x68\xdf\x55\x45\xc6\x0c\x82\xd2\xc0\xc7\x52\x69\x04\x25\x21\x55\x32\x17\x3c\x35\x49\x98\x57\x32\x85\x48\xc1\x7f\x17\x0b\xe7\x7f\x72\x55\x0c\xb9\x1c\x57\x82\xe9\xba\x8e\x1b\x2b\x91\x75\x42\x2a\x03\xc9\xa5\x7a\xa5\xa4\xc1\x6f\xa6\xae\x53\xf3\x8d\x54\xd1\x43\xe2\x5f\xf6\x60\xb1\x40\x99\x91\x93\xde\xf2\x2b\x25\xaa\xa9\x2c\x7b\xde\x39\xff\x08\x23\xc5\x45\xe2\x1f\x62\x40\xad\x95\x86\x45\x18\x68\x34\x95\x96\xa0\x12\x67\xd8\xd9\xed\xda\xb4\xdf\xbd\x45\x73\x7e\x16\xc5\x8b\x05\x8a\x12\xad\x1f\x3d\x68\x0e\xbc\xa4\x3f\x97\x59\x5d\xf7\x1e\xf5\x24\x0e\xeb\x30\x5c\x3a\x1d\x3a\xb8\x09\xc0\x0e\xe4\xf4\xe7\x80\x49\x9e\xae\x81\x3f\xf8\x31\xf4\xc1\xea\x2c\xe9\x9d\x05\x60\x6f\x3a\x06\xcf\xcd\xc7\x22\x0c\x78\x4e\x4e\x51\x74\xfe\x93\x64\xfc\x6c\x8d\xbe\xe8\x83\xe4\x82\xbc\x08\x0a\x82\x28\xb2\xfa\xbe\x68\x56\xbc\xd6\x3a\x42\xad\xe3\x38\x0c\xea\x4d\xc4\x6d\x61\x6a\x13\x51\x50\x95\x5c\x8e\xe9\x19\xbf\x61\x5a\x19\xa5\x9f\x92\x38\x1d\xd5\xc5\xf7\xb1\x38\x78\x88\x27\x39\xe2\xb0\x7b\xed\x5d\xea\xa0\xfa\x90\xda\x56\xdc\xbf\xea\x7c\xb5\x1b\xeb\xfd\x29\xdf\x10\x67\xdd\xb8\x22\x37\x9e\x8f\xd6\x25\xd0\x7f\x3b\x85\xfb\xd1\xf4\xef\x62\x69\x59\x28\x79\x0e\x0a\xfa\x2d\xa0\xbe\x70\xda\xf3\x32\xb9\xc4\xfb\xe8\x60\xb1\x48\x06\x77\x63\xd7\x76\x4e\x40\x2a\x58\x2c\x56\x5a\x11\x14\x5a\xcd\x78\x86\x19\xe4\x4a\x43\x65\x6f\x7b\x60\x19\x08\x03\xea\x52\x84\xb6\x20\xfc\x0e\x0c\x9f\x62\x69\xd8\xb4\xb8\x75\x52\xb7\x13\x14\x05\xea\x03\x48\xa0\x76\xd2\x6d\x94\xfc\xa2\xd4\x5d\x69\xa9\x5b\x89\xa7\x4c\x9d\x61\xae\x34\x3a\x50\xad\xd0\xde\xc1\xf5\x30\x7c\xda\xdb\x92\xbb\xd6\x5b\x8b\x65\x18\x06\xf2\x8f\x73\xcc\x59\x25\x8c\x6d\xc5\x5f\x2b\xd4\x1c\xcb\xe4\x52\xc9\xdf\x51\x2b\x7f\x34\x44\xa2\xd5\x93\x7e\xae\xee\x65\x4b\xbb\x47\xfa\x0b\x37\x13\x2f\xdc\x03\x15\x87\x61\x70\x7c\x0c\x67\x15\x17\x19\xa4\x2c\x9d\x20\xdc\xe1\x1c\xb8\x3c\x12\x5c\x22\x54\x63\xc1\xc5\x1c\x8e\x60\x3a\x2f\xbf\x0a\x98\x95\x50\xd0\xef\x42\xab\x91\xc0\x69\x19\x06\xa3\x2a\x27\x67\x4a\xa3\xa7\x4c\x8e\x05\x52\x69\x3c\xab\xf2\x1c\x75\x14\xdb\xd3\xe4\x8b\xe6\x06\x87\x46\x73\x39\x8e\x4a\xa3\x53\x25\x67\xc9\x3b\xa3\x58\xb4\x12\x1b\xc9\x7b\x2e\x33\x4a\x12\x22\xec\xb6\x07\x29\x69\xd5\x4c\x8e\x71\x35\x86\x28\x5e\x4a\x0b\xd4\xba\xee\xd4\xf2\xdb\xbe\x3e\x9b\x1b\x8c\x0e\x93\xc3\x5d\x6e\xac\xc4\xe4\x23\x6e\xac\xca\x7d\x8f\x1b\x0f\x75\x76\x18\x7d\x44\x17\x11\x72\xd2\x07\x3a\xf5\x07\x71\x18\xb4\x88\x0f\xaa\x06\xf1\x51\x95\xc7\x36\x66\x37\xf2\xef\xe2\xf3\x15\x71\x7c\x51\x99\xe4\xd3\x07\x95\xde\x91\x26\xcb\x7a\xcf\x91\x9f\x91\xa1\xdd\xdf\x5f\xdf\xe1\xfc\x66\x6f\x43\x57\x52\x38\x53\x61\x30\x63\xda\x06\xbc\x4d\xe6\xd0\xe6\xd1\x0b\x6f\x98\x00\x68\xe6\x0c\x8d\x86\x1c\x59\x85\xfc\x5d\xe7\x89\xc2\x3c\x0c\x82\x6d\x1e\x34\x05\x67\xb7\x48\x37\x1b\xf6\x93\x56\x95\xe9\x7e\xd0\x52\x48\x8f\xf1\xf2\x12\xd0\x4d\x8a\x21\xcd\x0b\xd3\x42\xe0\x14\xa5\x89\x9a\x5b\xee\xb6\x75\x5a\x19\x45\x2a\x29\x72\x78\x0f\x66\xeb\xd1\x68\x41\x23\x10\x5b\x53\x54\x6d\x18\x97\xe5\xa9\x9c\x6f\x2b\x04\x03\xcd\xa7\x4c\xcf\xdf\xe3\x7c\x59\x98\x67\x31\xbc\x7c\xf9\x34\x2d\x9b\xca\xc9\x2c\x76\x1e\xb5\x18\xb0\xa2\x40\x99\xf9\x2b\x5f\x9f\xf0\x9b\xa6\x09\x5c\xf3\xff\xfd\x74\x72\x93\x24\x09\xdd\x8f\xa2\xdc\xfe\xf0\x1c\x04\x4a\x2f\x1e\x53\x17\xf8\xbf\xd3\xb8\xb3\x09\x54\xd2\xae\x1c\x46\xf9\x72\xbf\xde\x12\x7a\x90\xaa\x4a\x64\xb6\x28\x8f\x6c\xb5\xf3\x3e\xa6\xf6\x1e\x20\x78\x69\x5b\x84\xed\x11\x64\x6e\x9d\xc0\x0b\xd4\x63\x8c\x34\x3e\x89\xb8\x1f\xd5\xe3\x91\xa5\xd4\x09\x7c\xcb\x3f\xe9\xaf\x55\xc4\xab\xce\xd3\x8f\xe7\xc5\xc3\xe0\xf0\x61\xed\xcd\x6f\x0f\x6b\x27\xb0\x3f\x3a\x2d\xdd\xee\xcb\x67\xa6\xdb\xfb\xbf\x91\x6e\x5b\x82\x12\xea\xa8\x73\x68\x3b\xab\x6d\x8a\xae\x98\x7d\xa4\xa3\x8b\xe1\xf0\xe3\x87\x28\xe3\x4c\x60\x6a\x7a\x70\xb0\x66\xf2\x60\xeb\xd5\x37\x24\x5c\x03\x56\xa7\xe2\x59\x44\xee\x27\xdc\x20\x39\xd7\xe6\xce\x53\x95\xfa\x9c\xda\xa0\x69\xf9\xaa\xb1\x6b\x45\x97\xf7\x9f\x31\x51\xe1\x05\x2b\x0a\x2e\xc7\x3d\x5b\xa6\x3b\x58\x70\x99\xf9\xa3\x6d\x0e\x7d\x9e\x17\xdb\xd9\x5f\xaa\x5d\xfa\x10\x3b\xfe\xd7\xc6\x9f\x95\xf9\xa7\x5b\x12\x08\x20\x12\xf4\x01\xe2\x3c\xd6\x68\x9e\xdb\x5f\x4b\x4c\xb0\xd1\xd5\x55\x5f\x9b\x1a\x56\xdb\x36\x27\x2a\x9b\xa9\x1a\x73\x0a\x96\xe4\x9d\xcc\xb8\xc6\xd4\x44\xcd\x8b\xdf\x48\xe2\xd7\x3c\x52\x34\x6c\xcc\x98\x58\x19\xe9\xec\x61\xf9\x46\xab\x69\x73\x05\xab\xd0\xf7\xe8\x15\x9e\x62\xd7\x53\x9d\x27\x25\x5c\xdf\x70\x69\x50\xe7\x2c\xc5\x85\x1b\x53\x09\xbb\x75\xb0\x3a\x40\x36\x1f\xb6\xc6\x07\x46\x6f\x37\xdd\xd1\xe1\x6e\xca\x73\x37\xc6\x9f\xe3\xa8\x1a\x5f\xa8\x0c\xad\xd6\x7c\x6a\x92\x37\x85\xe6\xd2\x08\x19\xb5\xe7\x76\xb6\xd1\x8d\x2e\x9b\x6e\xf1\x6e\x69\x42\xa7\xb5\xb6\xe3\x3e\x6b\x2b\x8d\x9b\xd6\x03\x17\x1b\x34\x70\x27\x36\x93\x3f\xa9\xfb\xa8\xe3\x84\xb3\x41\xc9\x90\x0c\x53\x66\x63\x8d\x40\xf1\x89\xe4\x56\xa0\xed\x9a\xbc\xa9\xc8\x8e\xf7\x4f\xd1\xea\xf7\xc0\x65\x6c\xf5\xfb\x50\x7e\x15\xc9\x6b\xad\x2f\xd5\x27\x75\xef\x66\x42\x6f\x91\x82\xee\xf8\x18\x06\xaa\x34\x63\x8d\xa5\xdd\x03\xe5\xa1\xf1\xc4\x03\x93\x73\x33\xa1\x85\xf1\x7e\x82\x12\xcc\x04\x35\x1e\x96\xb4\x18\xb9\x92\xe0\x23\x13\xec\x45\xb6\xc3\x74\xdb\x64\x91\xbd\x1f\x2d\x73\x9b\x51\x5a\x07\xe5\xe1\x77\xbb\x31\x59\x85\xa0\x5d\xa9\x36\xae\x42\xd4\x04\x68\x99\xa6\x4d\xda\x15\xde\x27\xb4\x82\x83\x36\x78\xba\xb3\xe6\x7e\xc3\x6b\x33\x24\xef\x21\x6e\x87\x62\xe8\xbb\xeb\xee\x6d\x60\x39\x1c\x07\x8f\xac\x9b\xcb\xff\x1d\x66\xea\x34\x37\xa8\xf7\x59\x35\x1f\xfb\x7f\x99\x5d\x3d\xfd\x72\xb9\xa4\xd1\x1b\x91\x5c\x74\xd7\xce\x3a\xfc\x2b\x00\x00\xff\xff\x9c\xc2\xea\x1c\x02\x16\x00\x00")

func templates17_upsertTplBytes() ([]byte, error) {
	return bindataRead(
		_templates17_upsertTpl,
		"templates/17_upsert.tpl",
	)
}

func templates17_upsertTpl() (*asset, error) {
	bytes, err := templates17_upsertTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/17_upsert.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSingletonMssql_upsertTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x53\x4d\x6f\xda\x40\x10\x3d\x7b\x7f\xc5\xd4\x52\x24\xaf\xb2\x72\x9a\x6b\x23\x2a\xd1\xe0\x26\x54\xc4\x40\x6c\xda\x03\xe1\xb0\xe0\x31\x59\xc9\x2c\x68\x3f\x50\xa3\x2a\xff\xbd\x1a\xdb\x94\xcf\x4a\x55\x2f\xe0\x9d\x8f\xa7\x37\xef\xcd\xdc\xdc\xc0\xdc\xab\xaa\x98\x6c\x2c\x1a\x37\xf6\x68\xde\x9e\xb2\x6c\x3c\x68\xa2\x16\x24\xd0\xc3\x3a\xe9\x70\x85\xda\x81\x75\x46\xe9\x25\x78\x4b\xbf\xee\x15\xc1\xd7\x8d\x3d\xe9\x24\x6c\xcc\x7a\xab\x0a\x2c\x62\x56\x7a\xbd\xb8\x8c\x1b\x15\x4a\x42\x61\xd4\x16\x8d\x8d\x7b\x4a\x56\xb8\x70\x02\x9c\x9c\x57\x98\xca\x15\xb6\xf8\x02\x36\x46\xad\xa4\x79\x13\xe0\x37\x85\x74\x28\x40\x69\x02\x82\xe9\x6c\x57\xb1\xf6\x6e\xe3\xf7\x01\xbe\xa3\xf6\x8b\x05\x6d\x6d\x87\x42\x2b\xa9\x97\x15\xc6\xfd\x02\xb5\x1b\xfb\xb5\xc3\xac\x52\x0b\x24\x1a\xf1\x60\x2c\x80\xfe\x9f\xc7\x3b\x78\xce\x58\x30\xf7\x25\x7c\x3a\x6c\x7d\x40\xf7\xc5\x97\x25\x9a\x88\xb3\xa0\xc0\x12\xcd\x41\x72\xe4\x77\xc9\xb9\x2f\xa9\xdd\x3a\x69\x5c\x5f\x17\xf8\x93\x50\x6e\x19\x0b\xca\x95\x8b\xbf\x6e\x8c\xd2\xae\xa4\x22\x01\xe1\x53\xf2\xfc\x90\x40\x3f\xcd\x87\x70\x65\x41\x5a\x98\xba\xd9\x8b\x0e\x0f\x74\xe0\x97\xda\x26\x59\x3f\x7d\x80\x28\x4b\x06\xc9\x7d\x0e\x57\x96\xd7\xad\x76\x06\xd1\xf4\xca\xce\x38\x21\xb0\x20\x38\xe0\x56\xc9\x05\xbe\xae\xab\x02\x8d\xad\x07\x9e\x58\xac\x99\x1d\x26\x04\x54\xa8\xa3\x56\x6e\x2e\x60\xcf\x5f\xc0\x2d\x6f\x01\x95\x5e\xda\xf8\xdb\x5a\xfd\x29\x14\xad\xda\x51\xa3\x1f\xbf\x0e\x45\x78\x7d\x10\x1a\x8c\x39\x3f\x9a\xa1\x1d\x61\x98\x42\x14\x52\x62\x6d\x40\x09\xd8\x92\x46\x46\xea\x25\xee\x0c\x27\xfb\x02\x55\x82\x82\x0f\x1d\xf8\x58\xbf\xce\x51\xa0\x9b\xf6\x80\x60\x82\x77\x16\x5c\x10\x6a\x6a\x67\x31\x49\x02\x1d\x52\xb6\xfe\x0c\x05\x6c\x05\x6c\x39\xa3\x96\x33\x40\xd2\xee\xc4\xbc\xeb\xce\x91\x30\xec\x42\xd7\x8f\xc7\x24\x85\xa7\x6e\x7e\xff\x98\xf4\x20\xa7\x47\x78\xd9\xb7\x51\xaf\x9b\x27\x90\x25\x64\x5a\xed\xf3\xde\xa3\x0c\xdd\x48\x1a\xb9\x22\xd3\x6d\x74\xac\xe0\xa9\xc8\xc7\xe6\x34\x87\xc1\x2f\xd3\x6e\x93\x7f\x65\x9d\x0e\xf3\x7f\x61\xde\x4f\xb3\xe4\x39\x87\x88\x76\xed\x7b\x77\x30\x49\xb2\xfa\x3b\x3c\x5b\x8b\xe6\x7c\x04\x84\x24\xe6\x7f\x6f\x61\x7b\x84\xa7\x4b\x48\x63\xa8\xb2\xae\x68\x8e\x9e\xc3\xe7\x76\x37\xce\x29\xbf\xe8\xe1\x24\x1f\x4d\x72\x68\xb8\x27\xbd\xda\xfe\xbb\x70\x27\x66\x4b\xb8\x01\x12\x10\xce\xc4\xbe\x30\xa4\x9d\x7d\x07\xac\x2c\x9e\xa0\xb7\xe0\x77\x61\xbd\x40\x2c\x30\xe8\xbc\xd1\x30\xf7\x65\x9c\x35\x1e\x71\xf6\xce\x7e\x07\x00\x00\xff\xff\xa0\xc3\x9b\xd6\x4d\x05\x00\x00")

func templatesSingletonMssql_upsertTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSingletonMssql_upsertTpl,
		"templates/singleton/mssql_upsert.tpl",
	)
}

func templatesSingletonMssql_upsertTpl() (*asset, error) {
	bytes, err := templatesSingletonMssql_upsertTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/singleton/mssql_upsert.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonMssql_main_testTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x57\xdf\x53\xdb\xc6\x13\x7f\x96\xfe\x8a\xfd\x7a\x26\x89\xc4\x57\x3d\x92\xa6\xd3\x07\x32\x9e\x8c\xb1\x45\xc2\x04\xdb\xc4\x52\x9b\x76\x08\x85\xb3\xb4\x82\x9b\x48\x27\xfb\xee\x04\x71\x29\xff\x7b\xe7\xee\x24\x59\x26\xd8\x0d\x7d\x29\x2f\xd6\xed\x7e\xf6\xf7\xed\xee\x71\x43\x05\x88\xab\xaf\xe3\x28\xfa\x78\xf2\x05\x57\xd0\x07\x81\x57\xf8\x75\x41\xc6\x95\x54\xc3\xb2\x58\xb0\x1c\xbd\x4b\xef\x6d\xe1\xff\x31\x38\x89\xc3\x19\xc4\x83\xc3\x93\x10\xc8\xde\x60\x34\xfa\x2c\xff\x3f\x9c\x4e\xa2\x78\x36\x38\x9e\xc4\x40\xf6\xe0\x68\x3a\x0b\x8f\xdf\x4d\xe0\x43\xf8\x3b\xd9\x7b\x4b\xf6\x3e\xf3\xb7\xb3\xf0\x28\x9c\x85\x93\x61\x18\x91\xbd\x4b\xdf\x75\xd5\x6a\x81\x50\x48\xb9\xcc\x63\x94\x0a\x05\x48\x25\xaa\x44\xc1\x9d\xeb\xa4\xf3\x61\xc9\x39\xe8\xbf\x3d\xb9\xcc\xc9\xe8\x50\xd3\x26\xb4\x40\x43\x93\x4a\x30\x7e\xe5\x3a\xd7\xa5\x54\x00\x1b\xa4\x4a\xa2\x78\x40\x5a\x50\x29\x1f\x90\xa4\xcc\x8b\x32\xc5\x0d\x54\x29\x1a\x5d\x8c\x2b\xd7\x51\x28\xd5\xe8\xd0\x98\xac\x21\xf7\xae\x9b\x55\x3c\x01\xc6\x99\xf2\x7c\xeb\xe6\x98\x32\x0e\x7d\x78\xde\x09\xe3\xee\xbe\x45\x7a\x05\xec\x75\x38\x3e\x48\x54\xd5\xc2\xf3\x01\x85\x28\x85\xd6\xa0\x73\x8e\x42\x58\x82\xeb\x3a\x37\x6c\x81\x82\x44\xa8\x46\x98\xd1\x2a\x57\x5e\xcf\xc8\x13\x99\x5c\x63\x41\x7b\x01\xf4\xd2\x79\xd9\xf3\x77\x00\x6d\x64\x1a\xa9\x44\x85\xbb\xa0\x3a\xe2\x5e\x00\xaf\x7e\x7a\xfd\xda\x77\x5d\xa7\x20\x75\x86\xfb\x60\x25\xde\xa1\x8a\x4c\xe4\x8d\x40\x3a\xe7\xb4\x30\x2a\x0b\x62\x52\xbf\x15\xa9\xb9\x16\x67\xea\xb1\x15\xa7\xb9\x16\x67\x8a\xb4\x15\xa7\xb9\x35\x4e\x97\xa9\x83\x3b\xe6\x9b\xf1\x18\x50\x53\xde\xad\xfa\x9a\x2c\xe9\xb8\x75\xfa\xfb\x70\x43\x73\x4a\x0e\xf1\x8a\xf1\x5f\x69\xce\x52\xaa\x58\xc9\x3d\x9f\xd4\x07\xf4\x5c\xc7\x31\x10\xab\x67\x52\xaa\xb0\x58\xa8\x95\xb7\x33\xae\x00\x36\x8f\x4f\xd3\x61\x73\xd8\xea\xa8\x8f\x8d\x8e\x49\xa9\x3c\xf3\x11\x2e\x2b\x9a\x4b\x6f\x7b\x42\x02\x78\xd9\x2a\xb1\x94\xa7\x7a\xd2\xd4\xbd\x55\xd3\x12\x9e\xa6\xa7\xcd\x7a\xab\x68\x4d\x71\x1d\x9f\x0c\xaf\x31\xf9\xe2\xe9\x9a\xb0\xcc\x74\xc5\xff\xfa\xc0\x59\xae\xfb\xc4\x11\xa8\x2a\xc1\x35\xd5\x75\xee\x5d\xd7\xd9\xdf\x87\xa1\x40\xaa\x10\x28\x08\xca\xd3\xb2\x60\x7f\x62\x0a\xe9\x1c\xb4\x6b\x44\xdf\x82\x4e\x07\xf7\xd7\x18\x12\x29\x3a\xcf\xd1\x32\xbc\xe6\xd2\x77\x8c\xf6\xa1\x20\xa9\x28\x17\xb1\x11\xf7\xfc\x37\xbb\x5d\xe9\x8a\x25\xc6\xa3\xef\x14\x74\x1d\x0b\x1f\x16\x29\x1c\xf4\x01\xbf\x62\x42\x86\x65\x51\x50\x9e\x7a\x3d\xb9\xcc\x93\x22\xd5\x5d\xfc\x43\xd4\x0b\xc0\xf6\x9b\x3e\xfd\x62\x4e\xfa\x42\xe9\xd3\xa9\x39\xe9\xee\xd0\xa7\xd4\x9c\xd6\x61\xeb\xa0\xb2\xc0\x78\x71\xd0\x87\x52\x92\xe9\x02\xb9\xd7\x33\xf1\xcb\x0b\x3b\x52\x88\x5c\xe6\xba\x69\x76\x65\xbc\x14\x92\x7c\x12\x74\xe1\xa1\xd0\x56\x33\xca\x72\x4c\x41\x95\x50\x2e\x90\xc3\x37\xea\x20\x63\xb9\x99\x12\x3a\xc6\x14\x33\x14\x90\x91\x61\x5e\x4a\x34\xb5\x6d\xa3\x26\x91\x4a\xcd\xf4\xe4\x78\x7b\xf4\x01\x57\x23\x94\x4a\x94\x2b\x14\x5e\x67\x0b\x05\x90\x6d\xd4\xa6\x2b\x4d\x85\xda\x99\xe5\xad\x7e\x4b\x2d\x0a\x36\xc7\x90\xd8\x9c\x37\x0e\x3f\x62\xe9\x13\x65\x8f\x19\xca\x0a\x45\x4e\x05\xe3\x2a\xe7\xda\x82\xff\x7d\xb6\x6f\x29\x53\x90\x95\x62\x8b\xf9\x5a\x03\x67\xf9\x8e\x2d\x22\xf3\x71\x99\xa2\x67\x66\x9c\x5d\x4e\x7e\xfd\xab\x1d\x93\xb7\x4c\x25\xd7\x60\xb8\x77\xae\x93\x50\x89\xf5\x32\x38\x58\xbb\x68\x09\x0d\x37\xa3\xb9\xdc\x64\x5b\x8a\xa9\x9f\xde\x19\x5d\x56\xca\xa4\x2e\x79\x4f\x3b\xbc\xd5\xc7\xcd\x56\x58\x2f\x3c\x7d\x3f\x0e\xfa\xa0\x73\x17\x2d\x74\xf2\x32\xef\xd2\x75\x86\xb3\x70\x10\x87\x30\x1a\xc4\x83\xc3\x41\x14\xc2\x33\xf9\xc6\x75\xde\x4d\x5d\xc7\x3e\x34\xd6\xf4\xb3\x57\xe7\xd2\x75\xa2\x30\x86\x59\x38\x18\x5d\x0c\xa7\xe3\xf1\x71\x1c\x87\xa3\x8b\x68\x32\x38\x8d\xde\x4f\x63\x98\x4e\x8c\xe8\xe5\xc3\x56\x68\xdc\x2f\x88\xa8\xf8\xb0\x48\x3d\xb9\xcc\x03\x78\x7a\xa3\xf9\xdb\x63\xee\x4e\x8d\x75\xc4\xfb\xfb\x10\x31\x9e\x20\x8c\x23\x88\x3e\x9e\xc0\x8f\x2f\x5f\xfd\x0c\x4c\x41\x42\x39\xcc\x11\xd2\x92\x23\xdc\x32\x75\x6d\x90\xa3\xd9\xf4\x74\x1d\xee\x19\x1c\x1f\x41\xf8\xdb\x71\x14\x47\x70\x0e\x77\x90\x52\x45\xe7\x54\xe2\x85\x9e\x71\xf0\xd7\xfa\x2c\x39\x5d\xc8\xeb\x52\x59\xc6\x3d\x9c\x41\x40\x08\xe1\x70\x0e\x67\x6f\xce\xb7\x25\xbd\xd5\xed\x45\xe1\x49\x38\x8c\xcd\xe4\x84\xa3\xd9\x74\x0c\x72\x25\x49\xa3\x5c\x82\xeb\x38\x9f\xde\x87\xb3\xd0\x02\xfa\xf0\xe2\x99\x7c\xa1\x6f\xfb\xa6\xb3\xcf\xe4\x23\x79\xff\x0f\xaa\xa0\x90\x8a\xb4\xbc\xe5\xdd\x1a\xb0\x4c\x0f\x75\xfb\xa8\xec\x74\x70\x43\x6b\xa7\x53\x77\x02\x1c\x3c\x71\x11\x7c\x5f\xef\x36\x61\xeb\xd1\x17\x34\xcd\x5f\x37\x6f\x00\x54\x5c\x49\x20\x84\x34\x4d\xdd\x06\x90\x3c\xb2\x24\x6a\x61\x2b\x45\x08\xf1\x0d\xac\x9d\xaa\x56\x87\x24\x13\xbc\x9d\x21\x4d\x51\x58\xa3\x7a\x9a\x4a\x95\x96\x95\xd2\x0a\x9f\xcf\x57\x0a\x25\x39\xac\xb2\xcc\xbc\x5c\x35\xab\x0e\xfe\x1b\x56\xad\x5c\x4b\x6a\xed\xfa\xa3\x25\xda\x91\x69\x85\xbb\x19\xd4\xec\x59\xc5\xff\x61\x7a\x36\xe3\x51\x54\x9c\x33\x7e\x75\xd0\x6b\x33\x63\x83\xf3\x1f\xe0\xad\xf1\xfa\xd5\xe1\xf9\x8f\xb0\x51\x88\x0d\xf6\xbf\x29\x55\x52\x72\x7d\x89\xbc\xfa\x3f\x90\xc0\x56\xc3\xdf\x71\x9f\xda\xcb\x6d\x59\x81\xd1\x6f\xec\x6d\xbe\xf3\x9d\x35\xa2\x4e\xdc\x32\xaf\x97\xb3\xf1\xa0\x17\x40\x2a\xd8\x0d\x0a\x62\xd6\xe0\x61\xc5\xf2\xf4\x63\x85\x62\x55\x87\xd4\x74\x44\xb3\xfa\x1f\x76\x9c\xed\x1e\xfb\x58\xd6\xbf\xf5\x33\xcb\xdf\xb5\xe7\x39\xcb\x83\x6f\xf2\xb3\x19\xc9\xbd\xfb\x77\x00\x00\x00\xff\xff\x53\x20\x2b\xd4\x23\x0e\x00\x00")

func templates_testSingletonMssql_main_testTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMssql_main_testTpl,
		"templates_test/singleton/mssql_main_test.tpl",
	)
}

func templates_testSingletonMssql_main_testTpl() (*asset, error) {
	bytes, err := templates_testSingletonMssql_main_testTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mssql_main_test.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonMssql_suites_testTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xc1\x0a\xc2\x30\x10\x44\xef\xfd\x8a\xa5\xe4\xd0\x4a\x9b\x0f\x10\x3c\x78\xd4\x83\x88\xb4\x1f\x10\xed\xb6\x04\xe2\x5a\xba\x5b\x10\x42\xfe\x5d\xd2\x46\xe9\xc1\xdb\x0c\x6f\x32\x99\xed\x67\x7a\x40\x83\x2c\xed\xc8\x38\x49\x21\xb0\x13\x64\xb1\x34\xe8\xa6\x04\x9f\x01\x78\x5f\xc3\x64\x68\x40\x50\x96\x3a\x7c\x57\xa0\xc4\xdc\x1d\xc2\xfe\x00\xba\x89\x8a\x43\x48\x39\xdb\x27\xa8\x4f\x7c\x7e\x59\x5a\x30\xd4\x3f\x8e\x8e\xb7\x56\x19\x67\x0d\xc7\x22\xa5\x8f\x51\x22\xaf\x8d\xdf\x96\x8b\x79\xe2\x92\x16\x7d\x9b\xa9\xc8\xbd\x5f\x9f\xe8\x76\xbc\xba\x79\x32\x2e\x84\xbc\x82\x38\xf8\x0f\x59\x2f\x2a\x97\xbf\x90\xba\xed\x8c\xe4\x42\xf6\x09\x00\x00\xff\xff\x11\x5d\x4c\xce\xff\x00\x00\x00")

func templates_testSingletonMssql_suites_testTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMssql_suites_testTpl,
		"templates_test/singleton/mssql_suites_test.tpl",
	)
}

func templates_testSingletonMssql_suites_testTpl() (*asset, error) {
	bytes, err := templates_testSingletonMssql_suites_testTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mssql_suites_test.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testUpsertTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\xcd\x6e\xdb\x3c\x10\x3c\x8b\x4f\xb1\x9f\xf1\xb5\xa0\x0a\x85\x41\xaf\x29\x7c\xc8\xdf\x21\x28\x6a\x18\xb1\xfc\x00\x8c\xb4\x72\x08\xd3\xa4\x40\xae\x6a\xb9\x02\xdf\xbd\x20\xe5\x24\x4e\x9c\x14\x39\xb4\x87\x1c\x6c\x89\xc4\xec\xcc\xee\xec\xae\x86\xe1\x04\xfe\x97\x5a\x49\x0f\x67\x53\x10\xe7\xf1\x0d\xbd\x28\xe5\x9d\x46\x18\x1f\x62\x26\x37\x18\x02\x6b\x3a\x53\x01\xa1\xa7\x61\x18\x23\xc4\xb2\x9d\xeb\xce\x49\x1d\xc2\xb2\xf5\xe8\x88\x13\x7c\x89\x00\x65\x56\xa2\xcc\x61\x60\x19\x89\xb9\x74\x52\x6b\xd4\x3c\x67\x2c\x53\x0d\x68\x34\xfc\x91\xe0\xca\x6e\xcd\x42\x99\x55\xa7\xa5\x0b\xe1\xd2\xea\x6e\x63\x7c\x0e\xd3\xe9\x9f\x60\x73\xa7\x36\xd2\xed\xbe\xe3\xee\x31\x60\x60\x59\x46\x62\xb1\x56\x2d\x9f\xc4\xff\x56\x99\x15\x50\xaa\x61\xab\xe8\x1e\xac\xd1\x3b\x68\xc7\x38\x58\xe3\x0e\xaa\x31\x72\x92\xb3\x2c\x30\x96\x79\xc4\x3a\xd6\xef\xa4\xa9\xed\x46\xfd\x42\x31\xc3\xed\x02\xb1\xe6\x39\xcb\x7e\x4a\x07\xe8\xd2\xcf\x3a\x96\x9d\x9e\xc2\x39\x11\x6e\x5a\x02\xba\x47\xb8\x99\x2d\xae\x6f\x4b\xf0\xaa\x46\xb0\x0d\x48\x03\xcb\x79\xbc\x61\x99\x8d\x8c\x07\x5e\x3d\x55\x30\x84\x64\x45\x24\x3d\xd4\x5c\x90\xeb\x2a\xe2\x31\x99\x02\x3e\xdb\x02\xde\x30\xe0\xea\xa2\xdc\xb5\xe8\x0b\x20\xd7\x61\xfe\x2d\xf1\xfc\x37\x05\xa3\xf4\xde\x88\xeb\x98\x69\xc3\x27\x4b\x93\x2c\x20\xfb\x24\xf2\x7a\x42\xe0\x93\xf4\x19\x7c\xf2\x93\x22\xf2\xed\x7d\x19\x06\xd5\x80\xb1\x04\x62\x66\x2f\xad\x21\xec\x29\x84\x8a\xfa\x58\x59\x35\x9e\xc5\x85\xac\xd6\x2b\x67\x3b\x53\xf3\x7c\x18\xd0\xd4\x21\xb0\x6c\x84\xfc\xe8\x3c\x95\x3d\x4f\x2c\x87\x0c\x47\x17\x77\x56\x69\x71\x81\x2b\x65\x12\x87\xf6\x78\x78\x57\xf6\xbc\xa2\xbe\x88\x05\x3e\x28\xbc\x0b\x94\xb3\xac\xc6\x06\x1d\x50\x2f\x6e\xad\xd6\x77\xb2\x5a\xc7\x86\x3e\x5a\x6f\xc5\x7e\x6e\xdf\xaa\x33\xb6\x00\x4d\x0d\x27\x21\x40\x3c\x25\xb9\x1b\xd3\xa0\xe3\xf9\xf3\xd3\xfb\xda\xd0\x25\xb9\xd7\x7b\x70\x64\x7e\x65\x3b\x43\xe9\xe2\xc5\x20\x3d\x2c\x1d\xcf\xc5\x65\xc4\xbc\x33\xfd\xa7\xca\x8f\xb3\xe4\x0f\xb2\x11\x92\x84\x23\xe8\xeb\x33\xc8\x64\x2b\x0d\x81\x35\x08\x0e\x2b\xeb\xea\x02\x56\x96\xce\x26\xc5\x88\xdf\x27\xfd\x62\x3b\x96\xf3\xab\xf3\xf2\xfa\xb5\xed\xf8\x1b\xf3\xdf\x48\xed\xf1\x4d\xd8\xd1\x77\x42\x08\xf1\x4f\xb7\xe5\xe3\xcd\xd5\x07\x19\xab\xc0\x7e\x07\x00\x00\xff\xff\xdf\xb2\x0c\xd1\xa7\x06\x00\x00")

func templates_testUpsertTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testUpsertTpl,
		"templates_test/upsert.tpl",
	)
}

func templates_testUpsertTpl() (*asset, error) {
	bytes, err := templates_testUpsertTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/upsert.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"templates/17_upsert.tpl": templates17_upsertTpl,
	"templates/singleton/mssql_upsert.tpl": templatesSingletonMssql_upsertTpl,
	"templates_test/singleton/mssql_main_test.tpl": templates_testSingletonMssql_main_testTpl,
	"templates_test/singleton/mssql_suites_test.tpl": templates_testSingletonMssql_suites_testTpl,
	"templates_test/upsert.tpl": templates_testUpsertTpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"17_upsert.tpl": &bintree{templates17_upsertTpl, map[string]*bintree{}},
		"singleton": &bintree{nil, map[string]*bintree{
			"mssql_upsert.tpl": &bintree{templatesSingletonMssql_upsertTpl, map[string]*bintree{}},
		}},
	}},
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"mssql_main_test.tpl": &bintree{templates_testSingletonMssql_main_testTpl, map[string]*bintree{}},
			"mssql_suites_test.tpl": &bintree{templates_testSingletonMssql_suites_testTpl, map[string]*bintree{}},
		}},
		"upsert.tpl": &bintree{templates_testUpsertTpl, map[string]*bintree{}},
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

