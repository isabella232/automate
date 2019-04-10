// Code generated by go-bindata.
// sources:
// policy/authz.rego
// policy/authz_v2.rego
// policy/common.rego
// policy/introspection.rego
// policy/introspection_v2.rego
// policy/rule_mappings.rego
// DO NOT EDIT!

package opa

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

var _policyAuthzRego = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\xc1\x6e\xdb\x3a\x10\x3c\x9b\x5f\x31\xb0\x2f\x2f\x0f\xb2\x82\xa6\x2d\x02\x04\xf0\xa1\x40\xfe\xa0\xbd\x05\x81\xb1\x11\x57\x16\x6b\x9a\x2b\x90\x94\x5b\xa7\xc8\xbf\x17\xa4\x44\xc3\x76\x0c\xe4\xd2\x9b\x34\xb3\x9c\x19\x0e\xc9\x9e\x9a\x2d\x6d\x18\x34\xc4\xee\x55\x29\xb3\xeb\xc5\x47\x68\x8a\x54\x37\xb2\xdb\x89\x3b\x83\x7a\xb1\xa6\x31\x1c\xb0\x80\x0c\x1e\xe5\x57\x29\xcd\x2d\x0d\x36\x66\x19\xf1\xe6\x95\x35\x56\x68\xc9\x06\x56\x6a\x81\xef\xc3\xcb\x4f\x6e\x62\xa8\x30\x38\x6b\xb6\x0c\xcf\x41\x06\xdf\x70\x00\x39\x0d\x6a\xa2\x11\x17\x2a\x90\x67\xf4\x5e\xf6\x46\xb3\x06\x25\x12\xe4\x3d\x1d\x6a\xb5\xc0\x8f\xce\x04\x34\x1d\x37\xdb\x80\xd8\x51\x04\x45\x58\xa6\x10\x21\x8e\x21\x2d\x62\xc7\x30\xae\x1f\x22\xc2\x64\x87\x1d\xc5\xa6\x4b\x26\x05\x82\x71\x79\x2e\x07\x3f\xd4\xaa\xa3\xb0\x9e\xa8\xa7\x5e\xec\xda\xe8\x67\xfc\x51\xb3\xb2\xaf\x82\xd5\x45\xf1\x69\xfd\x8c\x55\x5a\x9d\x56\xa9\x59\xb6\xbb\x20\x33\x36\xd2\x63\x81\x85\x5f\x4f\x69\xfe\x3b\x4e\x54\x45\xe9\x46\xbd\xa9\x9c\xa5\xf4\x72\x1a\x66\x52\x29\xd4\xb9\xcc\x11\xae\xf0\x2e\x74\xa1\x8e\xea\x63\xcf\xa7\xda\x23\x72\x21\x39\x82\x57\x04\x47\x22\xcb\xbd\x5b\x58\x21\x44\xf1\xac\x6f\x92\xac\x71\x58\x4d\xff\x1f\x0f\x8f\x9f\x58\x61\xfe\xff\x3c\x4f\x5b\x2b\xbf\xae\x9e\x02\xb7\x6d\x3a\xc3\x15\xe6\x79\x66\xae\x66\x57\xce\x6f\x04\x2f\x8b\x1c\xd1\xf3\x02\x92\x99\x66\x77\xf8\xc0\x2b\x8d\xfc\x0b\xab\xf3\x27\x93\xef\xfa\xa3\x09\xe2\x35\x7b\x7b\xa8\xb0\x63\x72\xc6\x6d\xf0\x98\x12\xc9\x9e\xbd\x37\x9a\x03\xbe\xa5\x9d\xaa\x05\x5a\xf1\x68\x07\x1f\x3b\xf6\xe0\xdf\xbd\x25\x47\xc9\xe0\x01\x5d\x8c\x7d\x78\xb8\xbd\x7d\xb1\xb2\xa9\xa5\x67\x37\xde\x6d\xda\xb0\x8b\xb5\xf8\xcd\xed\xe4\xb0\xdc\xb3\x0f\x43\x58\xea\xa3\xe7\xb2\x64\x59\xde\x7f\xba\xff\x72\xff\xb5\xb9\xfb\xac\xef\x5a\x75\xf2\x84\xd3\x0d\xc9\xfe\x33\x27\x11\xa9\x08\xf5\xa6\xfe\x06\x00\x00\xff\xff\xe1\x34\x20\xc8\x32\x04\x00\x00")

func policyAuthzRegoBytes() ([]byte, error) {
	return bindataRead(
		_policyAuthzRego,
		"policy/authz.rego",
	)
}

func policyAuthzRego() (*asset, error) {
	bytes, err := policyAuthzRegoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "policy/authz.rego", size: 1074, mode: os.FileMode(420), modTime: time.Unix(1540690498, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _policyAuthz_v2Rego = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x55\xb1\x6e\xdb\x30\x10\x9d\xcd\xaf\x38\x70\xb2\x03\xd5\x28\x3a\x06\xd0\x50\x14\x5d\x8a\x00\x2d\xda\x6e\x82\x40\xd0\x12\xdd\xb0\x95\x48\x81\xa4\x12\x24\x81\xfd\xed\x05\x79\x94\x4c\x29\x76\x6c\x4f\x99\xa4\x7b\x7c\xef\xee\xe9\x44\x1e\x3b\x5e\xfd\xe3\x7f\x04\xf0\xde\xdd\x3f\xb3\x87\x4f\x84\xc8\xb6\xd3\xc6\x41\xcd\x1d\x5f\x57\xba\x6d\xb5\x9a\x40\x9d\x6e\x64\x25\x85\x9d\x80\x46\x37\xc2\x12\x52\x8b\x2d\xef\x1b\x17\x92\x69\x23\x9f\x45\x0d\x39\x6c\x79\x63\x05\x21\xf7\xdc\xb2\x56\xb4\x1b\x61\x8a\x4e\x37\x4c\xd6\x25\xbc\x90\x85\x7f\xb5\xfd\x06\x6e\x73\x18\x12\x0f\xcb\x6b\x64\xdb\x82\x95\x64\x21\x55\xd7\xbb\x81\x19\x82\xb5\xed\x37\x7f\x45\xe5\x70\x1d\x8d\x0e\x18\x6b\xb9\xab\xee\x85\x5d\x8e\xb2\x0c\x62\xa5\x15\xd9\xa1\x17\x23\xac\xee\x4d\x25\x8a\x58\x2f\x03\xeb\xb8\x13\xad\x50\xce\x57\x0f\xee\x0e\xc8\xc0\x3e\x6a\x74\xa4\xd9\x62\x92\x63\x3d\xa8\x26\x1e\x07\x70\x6a\x72\x84\x53\x1f\x03\x16\x4c\x2b\xcd\x1e\x65\x53\x57\xdc\xd4\x4b\xbe\xf2\xf6\x2a\xad\x1c\x97\xca\x2e\x79\x06\xf4\x86\xae\x20\x1f\xba\xbd\x23\x84\x57\x4e\x6a\x95\x14\xf1\x89\xb5\x11\x75\x90\xa6\xc9\x22\xec\x9b\xec\x33\x60\x78\x26\x45\x0e\xe9\xe2\xd2\x76\x8d\x74\x31\x51\x06\xf4\x96\xae\x32\x40\xcc\x8b\x7c\xbc\x9a\xa6\x5b\x16\x56\x98\x07\xe9\x3f\x97\xde\xd0\x32\x83\x43\xcc\x32\x60\xa5\xaf\xe0\x4c\x2f\x4e\xaa\xdc\x53\x77\x44\x8b\xe8\x05\x72\x7a\x43\x33\x78\x10\x66\x33\x2f\x1d\xb0\x53\xf2\x89\x8a\x5d\xc4\x2f\x33\x60\x87\x65\xbf\xf1\x90\x72\xd9\xb6\x43\xee\x75\x9b\x0e\x35\xb8\xe5\x5e\xfd\x41\xbf\xd3\x10\x4c\x0b\x23\x32\x1e\x8d\xf3\x0e\xaf\x38\x02\xba\x11\x90\x83\x7f\x30\x59\x93\x45\x98\x14\x45\x0c\x53\xb3\x03\x09\x91\x33\xd6\x13\xe6\xe8\xba\x33\xda\x9f\xfd\xa2\x88\x2f\x78\xe4\x4f\xf8\x47\xca\x75\x8d\x8d\x22\xec\x6c\x0c\x66\xfe\x12\x4a\x06\x31\x78\x27\x83\x95\xee\x95\x9b\xd9\x0a\x13\xe2\xe3\xfb\xf8\x51\xda\xc1\xd4\x8d\xf7\x31\x6f\x23\x9b\x4c\x29\x7c\xf5\xa6\xe9\x7e\xff\xf9\xee\xee\xc3\x8f\x9f\xdf\xbf\x7d\xfd\xf2\xfb\xd7\x7e\x4f\x8f\xa9\x67\x43\xee\x20\x97\xea\x14\x9d\xc5\x49\x8a\xcd\x3a\x34\x28\x70\x8a\x42\x6c\xb7\x67\x3a\x83\x8c\xeb\x1a\x83\x1a\xb2\x78\x7d\x27\x22\x76\xee\x6e\x42\xd6\xdb\xc7\x34\x0c\xf0\xa6\xd1\x8f\x90\x43\xfc\x18\x1a\x62\x1a\x47\x6c\xe9\xef\x6b\xf5\x94\x2c\xfb\x30\x59\x4d\x6e\xf1\x17\xb2\x08\x5a\xfc\x8d\x9e\x37\xa6\x17\xf5\xb8\x91\xe2\x33\xb4\x65\x5e\xf2\xad\x0f\xb9\x74\x23\x86\x9a\xa3\xa9\xa3\x65\x4f\x59\x22\x3b\xf2\x3f\x00\x00\xff\xff\x17\xb5\x82\x86\xea\x08\x00\x00")

func policyAuthz_v2RegoBytes() ([]byte, error) {
	return bindataRead(
		_policyAuthz_v2Rego,
		"policy/authz_v2.rego",
	)
}

func policyAuthz_v2Rego() (*asset, error) {
	bytes, err := policyAuthz_v2RegoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "policy/authz_v2.rego", size: 2282, mode: os.FileMode(420), modTime: time.Unix(1554324388, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _policyCommonRego = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x54\xdd\x8e\xdb\x36\x13\xbd\x36\x9f\xe2\x7c\x52\x00\xff\x40\xb1\xbf\xe6\xa6\xa8\x50\x17\x08\x9a\xeb\xa0\x68\x83\xde\x04\x81\x41\x51\x63\x8b\x5d\x99\x14\xc8\xa1\xbd\xee\x62\xdf\xbd\x20\xf5\xb3\x5e\x6f\xda\x24\xad\xaf\x4c\x6a\xe6\x1c\xce\x39\x33\x93\xe3\x17\xa9\xee\xe4\x81\xa0\xec\xf1\x68\x0d\x94\x35\x2c\xb5\xf1\xd8\x07\xa3\x58\x5b\xe3\x21\x4d\x0d\x17\x5a\xf2\xe0\x46\x32\xa4\x23\xf8\x46\x3a\xaa\x51\x11\x9f\x89\x0c\x4e\xdf\xa5\xa0\xd3\x1b\x91\xc3\xee\x61\x83\x83\xa3\x83\x45\x6b\x0f\x5a\x89\xee\x19\x83\x10\xb9\xc8\xf1\xbb\x74\x5a\x56\x2d\x81\xee\x3b\x69\xbc\xb6\x46\xe4\x42\x18\xbb\x3b\x0d\x1f\xfc\x42\x2e\xf1\x20\x66\xe3\x7b\x16\xb2\x40\xf6\xea\x21\x5b\x62\xbb\xc5\x5e\xb6\x9e\xc4\xa3\x10\xb7\xd1\xda\xd4\x74\x6f\xf7\x4f\xc1\x3f\xe2\xfa\xea\x31\x5b\xc6\xac\x1c\xef\x68\xaf\x4d\x2a\x88\xa6\x42\x31\x4f\x6f\xa9\xe7\x38\x37\x5a\x35\x70\xc4\xc1\x19\x0f\xcd\x1e\x27\xd9\x06\xc2\x49\xcb\x94\x61\x03\x77\x81\x31\x92\x8b\x7c\x4c\xa5\x7a\xbe\x16\x39\xde\x5b\xa6\x12\x2a\x38\x47\x86\xdb\x4b\x01\x6b\xda\x4b\x5f\x69\xdd\x73\x5a\x43\x53\x3a\xce\x84\x3b\x63\xcf\x25\x5e\x3d\xc8\x37\x65\xf0\xe4\x8c\x3c\xd2\xe3\x5a\xf4\x19\x0b\xeb\xf4\x61\x89\x2d\x46\x8e\x58\x68\x8e\x31\x0e\xda\xc3\x58\x46\x9d\x4a\xaa\xb1\xfd\x29\x1e\x1b\x6d\x0e\x60\x3b\xe4\x88\x59\x8c\x18\x33\xc4\x6c\x42\x2a\xb7\x88\xe8\x51\x94\x7f\x20\xbb\x0e\x77\xd4\xb5\x52\x51\x8a\x4b\x22\x5f\x3f\x39\x2b\x26\x92\x24\xf4\xf4\xc6\x2d\x42\x6f\x4f\x17\x78\xed\x43\xf5\x07\x29\xf6\x1f\x77\x9f\xb0\x45\xba\xdb\xf9\x50\x89\x99\xef\x5a\xcd\x8b\xe9\xa2\x40\x56\x66\x05\x3e\x66\x11\x26\x2b\xb0\x2b\x10\x3e\x25\xdc\xb3\x6e\x6b\x25\x5d\x3d\xb8\x4e\xa6\xf6\x67\xcd\x4d\xf2\xb8\x5c\x8d\x26\xff\xdc\x90\xba\xeb\x7b\x56\x33\x6a\x4b\xbd\x50\x64\x6a\xc4\xe8\x14\x8a\xb7\xef\xdf\x4d\x21\x83\x92\x12\xde\xb6\x9a\xa5\xbb\x20\x5b\x65\x4f\x86\x7e\x68\x08\xad\x64\x26\x17\x23\xeb\xe8\xa1\xb7\x7d\xf2\x99\xe2\x79\xde\x83\x87\xae\xc7\x9f\xf7\xe5\x3a\xf2\x36\x38\x15\x45\x58\xcd\x45\x3e\x8c\x92\x36\x69\x4c\x3a\xe9\x58\xcb\x16\x8e\x7c\x68\xd9\x8f\x64\x13\xaa\x3c\x59\x5d\x23\x33\x96\xb3\x62\xe8\xcb\x26\x76\xb4\xf3\x37\xa9\xb0\x1d\xeb\xa3\xfe\x53\xa6\x99\x2d\xe0\x29\xf6\x65\xc3\xdc\xf9\x72\xb3\x39\x68\x6e\x42\xb5\x56\xf6\xb8\xb1\x1d\x99\xd7\x9d\x6d\xb5\xba\xbc\x96\x07\x32\xbc\xb1\x9d\xdc\x68\xef\x03\xf9\xcd\xf7\xff\xff\x61\x2d\x8c\xe5\xdd\x97\x14\x7e\x9a\xc2\x99\xc4\xff\xb6\x51\xa8\x5e\xf4\x0f\x8d\xf6\xf0\xa1\xeb\xac\xe3\xd4\xeb\x9e\x50\x05\x1f\xa7\xcd\xf7\xa5\x97\x22\x47\x84\x7d\x8b\x91\x04\x47\x79\xe9\x67\xc4\x2a\x15\x5c\xd4\x86\x93\xd6\x9e\xe1\x29\x0d\x67\x14\x66\x51\xbd\x48\x8a\x76\x55\x69\xb3\x54\xa9\xf9\x93\xee\x12\x9d\xa3\xbd\xbe\xc7\x82\xd6\x87\x35\x94\x34\x31\xcc\xcb\x0b\xb2\xfb\xf2\x52\xee\xad\x5d\x65\xcb\x04\xa8\x9e\x01\xca\xae\x6b\x75\x5c\x0a\x36\xd1\x0f\xf3\x3b\xbe\x20\x6d\x38\x69\x2e\xa8\x89\x3a\x72\xe3\xb5\x17\x39\xe2\xaf\xe7\xca\x64\xec\xa9\xa3\x64\xd5\x90\x8f\xa7\x2a\x4b\x79\xf1\x5f\xa9\xb2\x02\xc4\x6a\xbd\x5c\x4f\x0d\xbc\x4b\xa1\x51\xd5\x2a\xe9\xec\x59\x3a\x9e\x94\x66\xa7\x8f\x8b\x38\x06\xab\x6c\xd9\xf7\xb4\xc8\xf1\xeb\xd8\x4e\x29\x55\x9b\x83\xc8\xc5\xd8\x62\xbb\x81\x79\xa1\x4d\x01\xcf\xd6\x51\x9d\x60\x9f\xed\xd5\xe1\x3e\xed\x84\x27\x9f\xa7\x5b\x6d\xa2\xb7\xfd\x31\x72\xfe\x3b\xec\x97\xb8\x37\x25\x5f\xa1\x7c\x0d\xcb\x37\x3e\x7f\x58\x66\xff\x8d\xe0\x25\xf8\xb4\x08\x5f\x30\x7c\xae\xbc\x31\xfa\xf3\xfc\xbb\xde\x57\x6c\xc1\x2e\x50\xef\xed\x6f\xfd\x66\xbc\xb6\x76\x58\x96\x7f\xaf\xfe\xd7\x78\xf8\x05\x90\x6f\x35\xeb\x16\xee\xa6\x94\xbf\x02\x00\x00\xff\xff\xd6\xfb\x29\x9b\x54\x08\x00\x00")

func policyCommonRegoBytes() ([]byte, error) {
	return bindataRead(
		_policyCommonRego,
		"policy/common.rego",
	)
}

func policyCommonRego() (*asset, error) {
	bytes, err := policyCommonRegoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "policy/common.rego", size: 2132, mode: os.FileMode(420), modTime: time.Unix(1552666033, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _policyIntrospectionRego = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x51\x4d\x6a\xf3\x30\x10\x5d\x7b\x4e\x31\x38\x8b\x7c\x1f\x04\x1f\xa0\x90\x2b\xf4\x02\x26\x18\x45\x9a\xd4\x6a\x6d\x8d\x90\x46\x2d\x4d\xc9\xdd\x8b\xe4\x38\xc5\x49\x0a\xed\xa6\x4b\xbd\x79\xa3\xf7\x33\x5e\xe9\x17\xf5\x44\xa8\x92\xf4\xc7\xc6\x3a\x09\x1c\x3d\x69\xb1\xec\x00\xec\xe8\x39\x08\x1a\x25\xaa\x29\x84\x05\xa2\x79\x1c\xd9\x2d\x20\xcf\x83\xd5\x96\x22\x80\x57\x36\x74\xa3\x12\xdd\x53\xec\x02\x45\x4e\x41\x53\xdb\x7a\x1e\x3a\x6b\x36\x98\xc7\xbb\x1d\x7e\x40\x35\xaf\x9c\x47\x3b\xdc\xa2\xe7\x01\x2a\xeb\x7c\x92\x26\xf3\x62\xdb\x15\x54\xd9\x00\xd5\x24\xda\xcc\x3f\xce\x12\xff\xf2\xf4\x82\x6e\xf2\x17\x97\xd7\x7f\x38\x5d\xf9\x51\x25\xdf\x3d\x37\xdf\xa8\x4e\xed\x4c\x6b\x4b\xc9\x09\x2b\x82\x8b\x18\xe7\x41\xd1\x5e\xe1\x23\x0b\x3d\xa0\x30\x06\x92\x14\x1c\x4a\x4f\x18\xd3\x3e\x92\x20\x1f\xca\x2b\x2b\x70\xb0\x47\x32\x45\x32\xce\xb8\x0f\xfc\x6a\x0d\x19\x2c\xce\x36\xb0\x42\x4e\x01\x43\x1a\x28\xe2\x98\xa2\x60\x3d\x7d\x59\x17\xf6\x3a\xef\xae\xa7\x5b\x80\x1a\x06\x7e\x23\xd3\x65\xac\x2d\x11\xef\xf5\xdd\xd0\xe1\x40\x5a\x70\x8b\x75\x59\xa8\xe7\xb4\xbd\x8a\x5d\x4c\xfb\x67\xd2\x32\x73\xa1\xfa\xd1\x59\xaf\x68\xf7\xdb\xce\xc5\x18\x72\xf6\x37\x06\x0d\xb9\xf7\xbf\xf4\xf7\x75\x94\x2b\x8f\xb7\xd5\x42\xe5\x58\xf0\x26\x10\x9c\xe0\x33\x00\x00\xff\xff\x8b\x33\xc3\xfe\x62\x03\x00\x00")

func policyIntrospectionRegoBytes() ([]byte, error) {
	return bindataRead(
		_policyIntrospectionRego,
		"policy/introspection.rego",
	)
}

func policyIntrospectionRego() (*asset, error) {
	bytes, err := policyIntrospectionRegoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "policy/introspection.rego", size: 866, mode: os.FileMode(420), modTime: time.Unix(1540690498, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _policyIntrospection_v2Rego = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x93\x4d\x6e\xdb\x30\x10\x85\xd7\x9a\x53\x0c\x9c\x45\x5a\x40\xd0\xa2\xcb\x00\xbe\x42\x2f\x20\x08\x04\x23\x8e\x6b\xb6\x12\x47\x20\x47\x29\x9a\x20\x77\x2f\x48\x8a\xae\x6c\x39\x81\xfb\x83\xae\x24\x0e\xdf\x7b\x9a\x8f\x1c\x4d\xba\xff\xa6\xbf\x10\xea\x59\x8e\xcf\xea\xe9\x53\x63\x9d\x78\x0e\x13\xf5\x62\xd9\x01\xd8\x71\x62\x2f\x68\xb4\xe8\xa6\x68\x50\x87\xac\x3f\xdb\xed\x79\x1c\xd9\x9d\x95\x26\x1e\x6c\x6f\x29\x9c\x15\x3d\x0f\x14\x00\x26\x6d\xbd\x1a\xb5\xf4\x47\x0a\xca\x53\xe0\xd9\xf7\xd4\xb6\x13\x0f\xca\x9a\x1a\x83\x68\xa1\x91\x9c\xa4\x55\x14\x77\x1d\xbe\x40\x55\x22\x17\x61\xd7\x9c\x84\xa1\x5d\x7b\xba\xa6\x64\x86\x56\x75\xb8\x5f\x05\x96\x0d\xa8\xac\x9b\x66\x69\x62\xf8\x22\x8a\xaf\x50\x65\x92\x53\x40\xe9\xf2\x43\xdc\x3d\x55\xeb\x2b\x89\x1f\xe1\xf5\x02\x4c\xa7\x63\xfc\x97\x58\x39\x71\x03\x95\xcb\x6f\x22\xa5\xeb\x5a\xbc\xe7\x3c\xb9\x56\x6f\xa2\xfe\x07\x4b\x1c\x05\xdc\x63\x7c\x28\x6b\xa0\x4a\xa3\xd1\x2e\xcb\x0b\xd4\x54\xfd\x6b\xca\x55\x4a\x02\x4c\x22\x15\x25\x6d\x4b\x87\x03\xf5\x92\x49\x6a\xbc\x46\x99\xf9\xb2\x0e\x1f\xf6\x78\x3b\x69\xf6\x94\x0e\x8f\x3a\xa8\x91\xc6\x47\xf2\xc5\x09\xd5\x1f\xfc\x0f\x17\xa6\x5b\xee\x27\x32\xdf\xe1\x67\x16\x7a\x40\x61\xf4\x24\xb3\x77\x28\x47\xc2\x30\x3f\x06\x12\xe4\x43\x5a\xc5\x3e\xd9\xdb\x67\x32\xc9\x18\x4a\x7d\xf2\xfc\x64\x0d\x19\x4c\x37\x50\xc3\x1d\xf2\xec\xd1\xcf\x03\x05\x1c\xe7\x20\xb8\xcb\x91\xbb\xa4\xbe\x8f\xde\xfb\xfc\xdb\x83\x1e\x06\xfe\x4e\x26\x1f\x76\xea\x06\xf7\xb8\x3e\xff\x5d\x52\xec\xca\x05\xa8\x1a\x55\xd7\x01\x18\x72\xf6\x5d\x9b\x21\xf7\x63\xe3\xfa\x05\xb0\x76\xbe\x40\xb5\x6d\x03\x2a\xc7\x82\x9b\xcf\xc4\xa3\x3a\x89\x3d\x7f\xa5\x5e\xda\xe5\x99\xc7\x3c\xbf\xff\xde\x1c\x2c\xa6\x38\xb5\x50\xbd\x03\x7f\x7d\xfa\x52\x4b\x2b\xb0\x2b\x5d\xbd\xd5\x31\xbc\xc2\xcf\x00\x00\x00\xff\xff\x3c\xe8\x62\xec\xeb\x05\x00\x00")

func policyIntrospection_v2RegoBytes() ([]byte, error) {
	return bindataRead(
		_policyIntrospection_v2Rego,
		"policy/introspection_v2.rego",
	)
}

func policyIntrospection_v2Rego() (*asset, error) {
	bytes, err := policyIntrospection_v2RegoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "policy/introspection_v2.rego", size: 1515, mode: os.FileMode(420), modTime: time.Unix(1553727342, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _policyRule_mappingsRego = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\x2a\xcd\x49\x8d\xcf\x4d\x2c\x28\xc8\xcc\x4b\x2f\xe6\xe2\xca\xcc\x2d\xc8\x2f\x2a\x51\x48\x49\x2c\x49\xd4\x03\x49\x15\x73\x71\x81\xa9\xf8\xb4\xfc\xa2\xf8\x82\xa2\xfc\xac\xd4\xe4\x92\xe8\xa2\xd4\xe2\xd2\x9c\x92\x58\x85\x6a\x2e\x4e\x08\x53\xc1\xca\x16\x49\x4b\x74\x66\x5e\x41\x69\x89\x1e\x54\x75\x7c\x66\x4a\x6c\x74\x7c\x2c\x57\x2d\xb2\x49\x89\x39\x39\x30\xd3\x8a\x09\x19\xc7\x55\xcb\x05\x08\x00\x00\xff\xff\xb0\xed\xb0\x1b\xad\x00\x00\x00")

func policyRule_mappingsRegoBytes() ([]byte, error) {
	return bindataRead(
		_policyRule_mappingsRego,
		"policy/rule_mappings.rego",
	)
}

func policyRule_mappingsRego() (*asset, error) {
	bytes, err := policyRule_mappingsRegoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "policy/rule_mappings.rego", size: 173, mode: os.FileMode(420), modTime: time.Unix(1552666033, 0)}
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
	"policy/authz.rego": policyAuthzRego,
	"policy/authz_v2.rego": policyAuthz_v2Rego,
	"policy/common.rego": policyCommonRego,
	"policy/introspection.rego": policyIntrospectionRego,
	"policy/introspection_v2.rego": policyIntrospection_v2Rego,
	"policy/rule_mappings.rego": policyRule_mappingsRego,
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
	"policy": &bintree{nil, map[string]*bintree{
		"authz.rego": &bintree{policyAuthzRego, map[string]*bintree{}},
		"authz_v2.rego": &bintree{policyAuthz_v2Rego, map[string]*bintree{}},
		"common.rego": &bintree{policyCommonRego, map[string]*bintree{}},
		"introspection.rego": &bintree{policyIntrospectionRego, map[string]*bintree{}},
		"introspection_v2.rego": &bintree{policyIntrospection_v2Rego, map[string]*bintree{}},
		"rule_mappings.rego": &bintree{policyRule_mappingsRego, map[string]*bintree{}},
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

