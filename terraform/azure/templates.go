// Code generated by go-bindata.
// sources:
// templates/cf_lb.tf
// templates/network.tf
// templates/network_security_group.tf
// templates/output.tf
// templates/resource_group.tf
// templates/storage.tf
// templates/tls.tf
// templates/vars.tf
// DO NOT EDIT!

package azure

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

var _templatesCf_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x56\xc1\x8e\xe3\x28\x10\xbd\xe7\x2b\x10\xda\xab\x9d\x64\xa6\xd5\xda\x8b\x2f\xab\x3d\xec\x71\xa5\xf9\x00\x54\xb6\x71\x8c\x42\x80\x81\x22\x3d\xd9\x96\xff\x7d\x85\xb1\x93\x98\xd8\xee\x64\x72\x8b\xeb\x55\x51\xf5\x5e\x55\xc1\x19\xac\x80\x52\x72\x42\xdd\xc5\x21\x3f\xb1\x5a\x9f\x40\x28\x4a\x3e\xbb\xcd\xe6\x66\x34\xcd\x2f\x56\x71\x8b\xac\x04\xc7\xdf\xdf\xe6\xcc\x06\x9c\xfb\xd0\xb6\x8e\x36\xcb\x9d\xf6\xb6\xe2\x84\xc2\x7f\xde\x72\x7b\x62\xce\x97\x8a\x23\x25\xb4\x6a\x32\x17\x0e\xd8\x10\xa2\xe0\xc4\x49\xfa\x2b\x08\xfd\xe3\xf3\x0c\x36\xe7\xea\xcc\x44\xdd\x65\xd1\x61\x43\x08\xd4\xb5\xe5\xce\x31\x63\x79\x23\x7e\xdd\xe0\xfb\x5d\xbe\xcb\xf7\xef\xf9\x6e\xfb\x6d\x17\x70\xe3\xe1\xec\x60\xb5\x37\x2c\x9e\xd2\x87\x1d\x93\x99\x22\xf2\x52\xbb\x36\x0f\xb0\x2e\xb8\x9f\x85\x45\x0f\x92\x29\x8e\x1f\xda\x1e\xa3\xff\xc4\x3d\x41\x4c\xfc\x67\x8b\x37\xbe\x94\xa2\x62\xc2\x50\x42\x65\xb9\x52\xfc\x1a\x09\xb2\xcc\x84\x09\x09\x4a\x5d\x01\x0a\xad\xd6\x3d\x2d\x3f\x08\xad\xba\x45\x42\x26\x0e\x4f\x11\x73\xad\x82\x8d\x4a\x80\xbc\xe6\x52\x10\x5a\x5f\x14\x9c\x44\xb5\xc0\x01\x18\x23\x45\x04\xb3\x03\x20\xff\x80\x0b\x25\x74\xa0\x70\x99\x92\x07\x26\xc0\x98\x6c\xf4\x5f\xa8\xed\xf9\x92\xe6\xa8\x7c\x64\x70\x43\x88\x3b\xfa\x3e\xc5\xbb\x24\x0b\x42\x7f\x20\xa8\x1a\x6c\xcd\x7e\x9c\x40\x4a\xda\xdb\x51\x70\x9b\xda\xa3\xa5\x02\x03\x95\xc0\x0b\x29\xc8\xb7\x0d\x21\x5d\x88\x6b\xac\x2e\x79\x1a\x79\x9a\xcc\xbf\x01\xb2\xdb\xc7\x18\xc6\x6a\xd4\x95\x96\x09\xe6\x1f\x44\x33\x00\x00\xdb\x99\x20\x5b\xa9\x0f\x42\x45\x48\xab\x1d\xce\x40\x7a\x44\x1e\x4b\x9f\x6c\x83\x2e\xba\x09\x85\xdc\x9e\x21\x39\xfa\x7d\x37\x54\x7d\xe2\xda\x23\x99\x35\x7a\xd5\x72\x90\xd8\x5e\x18\xb6\x96\xbb\x56\xcb\x9a\x14\xe4\xfb\xc8\xc1\xa0\x66\x68\xac\x4a\xab\x46\x1c\xbc\x8d\xa2\xa4\xb4\xcc\x4d\xc5\xe0\x9c\x09\x93\x4d\x9c\x63\xce\x71\xeb\x30\x51\x3f\x31\xc0\xa2\xee\xb6\x11\xef\xb6\x37\x68\xfc\x92\xf7\x4b\xe8\xd6\x37\x7d\xde\x8d\xd5\x0a\xb9\xaa\x99\xd1\x16\xef\x93\x2d\x08\x1d\x6d\xc1\xd4\x22\x1a\x37\xa8\x13\x90\x05\x79\x7b\xfb\xfe\x6a\x10\xa9\x0f\x69\x8c\x99\x20\x5f\x52\xb8\x36\x59\x55\x93\x8d\x81\x16\xe8\x7c\xdc\x00\x29\xb3\x57\x44\x2e\xcb\xc0\xe8\x95\xac\x12\xaa\x63\xc8\xf0\xba\xc4\xb5\x96\x49\xb9\x0f\xd9\x0c\x3e\xd9\xe0\x93\x05\x9f\x87\x80\x81\x5d\xe6\x38\xa2\x50\x07\xb7\x56\xef\xb3\x3b\x3c\x2b\x79\xd6\xa2\xc3\x61\x68\xb5\x3e\x0a\xde\x5f\x7c\x35\x83\xa6\x11\x2a\x4e\x30\xfd\x5b\xb8\x70\xfb\xd5\x77\xa2\xcc\x9c\xf8\xe7\x6e\x71\x6c\x93\xc1\xb5\xfc\xa7\xe7\x0e\xd9\x74\x92\x0a\xb2\x1f\x03\x94\x9c\xa5\x75\x4d\xb7\x43\x4f\x8b\x73\xb2\xbf\xaa\x45\x13\x96\xed\xc3\x6a\x29\x08\x75\x4e\x66\x01\x11\x8f\xad\x01\x61\xda\x0e\xc9\x65\xdf\x8d\x7b\x25\xde\xef\x53\xdc\xf8\xf5\xa6\x73\x2f\x87\x14\x0e\xb9\xe2\x76\x55\x8e\x97\x75\x09\xa1\xa5\xc3\xa1\x17\x17\x7b\x9e\x2d\xf6\xd3\x17\xdd\x3d\x19\xc5\x07\xae\xd7\xa6\x7a\x56\xdd\x54\xe6\x01\x9c\x08\x94\x9e\x93\x08\xd4\x73\x3a\xb6\x86\xd5\x3e\x74\x39\xb3\x5e\xae\xdd\x19\x2f\xd2\x6a\x7f\x8e\xbd\x10\xe2\x32\xbc\x98\xf9\xd9\xf9\x0b\x5c\xb8\xdc\xc3\xbf\x89\xc8\xd3\x0a\x7e\x4f\xcc\xb9\xed\x70\xff\x70\x7b\x72\x31\x2c\x6c\x85\x97\x9e\x70\xf7\xe3\xdf\x85\xa7\x8c\xf6\x68\x3c\x12\x3a\xfb\x80\x09\x2a\x9c\x41\xfa\x24\xfc\x0c\x36\x1f\x8f\xba\x3e\x14\xff\x0f\x00\x00\xff\xff\x2d\x03\x11\x34\x7d\x0b\x00\x00")

func templatesCf_lbTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesCf_lbTf,
		"templates/cf_lb.tf",
	)
}

func templatesCf_lbTf() (*asset, error) {
	bytes, err := templatesCf_lbTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/cf_lb.tf", size: 2941, mode: os.FileMode(480), modTime: time.Unix(1515707288, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesNetworkTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xcf\x4a\xc4\x30\x10\xc6\xef\x79\x8a\x61\xf0\xa0\xb0\x5b\x3c\x7a\xf1\x49\x44\x42\x36\x19\xd7\xe0\x36\x29\x93\x3f\x8a\x25\xef\x2e\xa9\xad\xd8\xd8\xe2\xe6\xfc\x9b\x99\xef\xf7\x85\x29\xf8\xc4\x9a\x00\xd5\x67\x62\xe2\x5e\x66\xcb\x31\xa9\x8b\x74\x14\xdf\x3d\xbf\x21\xe0\xc9\x87\x57\x84\x51\x00\x38\xd5\x13\x34\xef\x11\xf0\x66\xcc\x8a\x3b\x72\x59\x5a\x53\x8e\x15\x3f\x66\x87\x02\x40\x19\xc3\x14\x82\x0c\x83\xd2\xf4\xc3\x3f\xcd\x03\xf3\x05\xa9\xad\xe1\x82\xcf\x02\xe0\xe2\xb5\x8a\xd6\xbb\xcd\xfd\x4c\x67\xeb\x5d\xa9\x7b\x97\xd4\xf2\xcc\x3e\x0d\x72\x8a\x35\x71\x8b\xc4\x1a\xe8\x6a\xa4\xae\x52\x05\x45\x11\xe2\xaf\x74\x48\x27\x47\xf1\x5f\xd7\x1d\xd9\xb0\x92\x1d\x98\x5e\xec\xc7\xef\x81\x2a\xf8\x7d\xe1\xb6\xf5\x3e\xc0\xc3\x01\xee\xef\x76\xad\xae\xd6\x02\x68\x3e\x6e\xa3\x95\x86\x68\x6a\xf9\x0a\x00\x00\xff\xff\xdb\x9a\xf5\x1a\x0b\x02\x00\x00")

func templatesNetworkTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesNetworkTf,
		"templates/network.tf",
	)
}

func templatesNetworkTf() (*asset, error) {
	bytes, err := templatesNetworkTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/network.tf", size: 523, mode: os.FileMode(480), modTime: time.Unix(1512421079, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesNetwork_security_groupTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x96\xcd\x6e\x9b\x40\x10\x80\xef\x3c\xc5\x68\xd5\x53\x24\x5b\x29\x81\xc8\x97\x1c\x7a\xec\xbd\x77\xb4\xde\x1d\xf0\xaa\x78\x07\xcd\x2e\x4e\xdb\x88\x77\xaf\x16\x6a\x27\x38\x38\xc2\xa8\x55\x85\x05\xe7\x6f\xfe\x76\x3e\x8d\x60\x74\x54\xb3\x42\x10\xf2\x57\xcd\xc8\xfb\xcc\xa2\x7f\x26\xfe\x9e\x39\x54\x35\x1b\xff\x33\x2b\x98\xea\x4a\x80\xd8\x92\xdb\x09\x78\x89\x00\xac\xdc\x23\x9c\x7d\x4f\x20\x3e\xbd\x1c\x24\xaf\xd1\x1e\x32\xa3\x9b\x55\x8b\x47\x00\x25\x29\xe9\x0d\xd9\x41\x98\xb1\x30\x64\x9b\xc0\x1d\x3b\xe9\xea\x65\x6d\x8d\x96\x3b\x36\xd6\x07\xd6\x21\xff\x3a\x50\x8d\x88\x22\x00\x2f\x0b\xd7\x36\x07\x80\xf6\x60\x98\xec\x1e\xad\x7f\xd7\x56\xa8\xd4\x44\x4d\x14\x5d\x31\xb8\xca\xaf\x18\x5b\xe5\xf3\x1e\x9a\xeb\x12\x05\x08\xf7\xd1\xae\x2f\x0e\xef\xba\x95\x57\x6c\x28\x24\x1b\x8e\x89\xef\xef\x23\x00\x6d\x18\xd5\xf9\x13\xbd\xe6\xfd\x6a\xb7\x54\x5b\x1d\xb2\x49\xa5\xd0\xb9\x8b\x1d\x7c\x29\x4b\x7a\xee\xaa\x92\x27\x45\xe5\x05\xee\x9b\xaa\x02\xf5\xe7\x39\x2b\x62\x9f\xb1\xb4\x05\xf6\xa9\xbb\xc0\x68\x74\xde\xd8\x76\x81\xef\xc0\x27\x10\x71\xfc\x26\x91\xd4\x9a\xd1\xb9\xac\x62\xcc\xcd\x8f\x0f\x12\x9d\x83\x47\x66\x48\x81\xde\x03\x8f\x50\x01\x60\x58\xde\x01\xa1\x86\xc1\x5e\xb6\xab\x44\x09\x81\x2b\x59\xa0\xf5\x13\x7c\x79\x13\x3c\x42\x9b\xcf\xf3\xd6\xe6\x71\xf3\xb8\x59\xc4\xe9\x8b\xd3\xad\x93\x78\xaa\x3b\xa7\xf8\x11\xfa\xc4\xf3\xd6\x27\x4e\xd3\x34\x5d\xfc\x39\xf9\xa3\xad\x9b\x60\x4d\x88\x1a\xe1\xca\xc3\xff\x70\xe5\xee\x2f\x99\x92\x3e\x2c\x9a\x9c\x34\x51\x8c\x7a\x57\x6f\x27\xa8\x72\x8c\x1c\xa1\x4b\x32\xef\xd3\xb2\xd9\x24\xc9\xa2\xcc\xab\x32\xf9\x6a\xe7\x7d\xf5\x0f\xcf\xcb\xcc\xff\x64\x92\xe4\x06\x2f\x8c\xca\xa7\xca\x52\x52\x31\xe5\xbc\x74\x81\xb7\xff\xe3\x92\xdc\xbc\x2e\xbf\x03\x00\x00\xff\xff\x77\x27\x82\x78\x45\x11\x00\x00")

func templatesNetwork_security_groupTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesNetwork_security_groupTf,
		"templates/network_security_group.tf",
	)
}

func templatesNetwork_security_groupTf() (*asset, error) {
	bytes, err := templatesNetwork_security_groupTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/network_security_group.tf", size: 4421, mode: os.FileMode(480), modTime: time.Unix(1512163867, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesOutputTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\x51\x6f\xdb\x20\x14\x85\xdf\xfd\x2b\x90\xb5\x87\x55\xda\x52\x2d\x92\xf7\x10\x69\xbf\x05\x11\xfb\xd6\x61\xc5\x80\xee\xbd\x38\xed\xaa\xfc\xf7\x09\x13\x67\x26\x35\x4b\x5a\x3f\xc2\x39\xdf\x3d\x1c\xb0\x0b\xec\x03\x8b\x7a\xb4\xc0\xd2\xaa\x01\x6a\xf1\x56\x09\x31\x2a\x13\x40\xfc\x12\xf5\x97\x37\xf5\x27\x20\xe0\x20\x47\x8d\x1c\x94\x91\x16\xf8\xe8\xf0\x79\xb3\x77\x74\xd8\x44\xc7\xa9\xae\x4e\x55\x35\x83\x28\xec\x6f\xa2\x92\xa6\x44\x40\x20\x17\xb0\x05\xd9\xa3\x0b\xfe\xff\xa4\x5c\x5b\xcc\xc4\x0e\x55\x0f\x52\xb5\xad\x0b\xf6\x56\xb8\x5c\x5c\x62\x76\xf0\xa4\x82\x61\x49\xd0\x06\xd4\xfc\x9a\x12\x14\xa9\xe7\xd6\xae\xe4\x25\x38\xbc\x30\xa0\x55\x46\xea\x32\xd1\x87\xbd\xd1\xad\xd4\x67\x88\xf6\x52\x75\x1d\x02\xd1\x55\x4e\x8d\xd0\xb2\xc3\x79\xf7\x8a\x77\x60\xf6\xb4\x7b\x7c\xbc\x87\xbb\xdb\x36\x4d\xd3\x64\x74\x8f\x7a\x54\x0c\xf2\x19\x5e\x97\xe0\xf8\x4d\x61\xd9\x90\x5c\x68\x26\xa4\x1c\x07\xda\x2c\x16\xa5\x87\xe1\x54\x57\x42\x10\x58\xd2\xac\xc7\x18\x8c\x31\x40\x36\x28\xa5\xfa\xf8\x9c\x8b\x4f\x3a\x0f\x96\xe8\xf0\x6e\xd4\x93\x32\x94\xcd\xfa\x1d\x06\xbf\x77\x2f\x32\xa0\xf9\x44\xfb\xbb\xed\x36\xab\x68\xbe\xf9\x56\x77\xf8\x0e\x37\x2a\xdc\x2c\x05\xe9\xee\x8c\x6b\x95\xa1\x49\x7b\xb9\xbe\xf8\x48\xa2\x27\x8e\xfb\x9e\x8c\x60\x47\xa9\xbb\xe9\x3c\xda\x9e\x1f\x4c\x84\xfc\x43\x67\xcb\xb9\xb0\x3f\x26\x59\xdc\x39\x38\xe2\xaf\xd3\xd0\xdc\xf1\x4d\xfc\x78\x98\x5c\x73\x23\x97\x5d\xed\xef\x71\x37\xc9\x7d\x39\xc3\x07\xed\x3f\x1f\x0a\x4f\x79\xf5\xff\x4d\x88\x4c\x93\xdb\x33\x7a\xc1\x7e\x5d\xd8\x9a\xbd\x3f\xde\x32\xf7\xc7\xdc\x3a\xd7\xb7\x2c\xa0\xc0\x58\x69\xba\x50\xc2\x1d\xb0\xb5\xe2\x13\xed\x6f\x00\x00\x00\xff\xff\x2b\xd8\x75\x99\xf7\x05\x00\x00")

func templatesOutputTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesOutputTf,
		"templates/output.tf",
	)
}

func templatesOutputTf() (*asset, error) {
	bytes, err := templatesOutputTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/output.tf", size: 1527, mode: os.FileMode(480), modTime: time.Unix(1515706590, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResource_groupTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\x41\x8a\xc4\x20\x10\x45\xf7\x9e\xe2\x23\xb3\x9d\xdc\x60\xce\x22\x15\x53\x64\x84\x44\x43\xa9\x59\x4c\xf0\xee\x83\x42\xa7\x3b\x24\xdd\x0d\x5d\x4b\xf9\xf5\xeb\xf9\x84\x63\xc8\x62\x19\x9a\xfe\xb2\xb0\xcc\xe6\xf6\x62\x46\x09\x79\xd1\xd0\x7d\x88\xbf\x1a\x9b\x02\x3c\xcd\x8c\x3a\x3f\xd0\x5f\xdb\x4a\xd2\xb1\x5f\x8d\x1b\xca\x77\xcb\x28\x60\x0a\x96\x92\x0b\xfe\x9e\x10\x1e\x5d\xf0\x45\x2b\x05\x24\x1a\x63\x2b\x02\xd8\xaf\x4e\x82\x9f\xd9\xa7\x53\x5b\x2d\x2a\xaa\x28\x75\x86\x5b\x72\x3f\x39\x6b\xdc\x13\xae\xab\x79\xcf\xfa\x72\x6b\xe7\x07\x8e\x66\xcc\xf1\x6a\x5b\xb8\x76\xd8\xd5\x8b\x5d\x8d\xb7\x9a\xfd\x0f\x86\x86\x41\x38\x46\x43\xd3\xa3\xb7\x98\x28\x39\xfb\x89\xb0\xff\x00\x00\x00\xff\xff\x58\xec\xbb\xb4\xcd\x01\x00\x00")

func templatesResource_groupTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesResource_groupTf,
		"templates/resource_group.tf",
	)
}

func templatesResource_groupTf() (*asset, error) {
	bytes, err := templatesResource_groupTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/resource_group.tf", size: 461, mode: os.FileMode(480), modTime: time.Unix(1512609663, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesStorageTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x91\xcd\x6a\xc3\x30\x0c\xc7\xef\x7e\x0a\x61\x76\xee\x1b\xec\xbc\xfb\xfa\x00\x46\x71\x44\x66\x70\xa4\x20\x2b\x81\xad\xe4\xdd\x47\xbc\x36\x6d\x0d\xa5\x3d\x2e\xd7\xfc\xf4\xff\xb2\x52\x91\x59\x23\x81\xc7\x9f\x59\x49\xc7\x50\x4c\x14\x07\x0a\x18\xa3\xcc\x6c\x1e\x7c\x27\xe5\xcb\xc3\xc9\x01\x30\x8e\x04\xcd\xf7\x0e\xfe\xed\xb4\xa0\x1e\x4a\x1a\xa7\x4c\x81\x78\x09\xa9\x5f\xbd\x03\xb8\x88\x87\x41\x65\x9e\x42\xbd\xae\xf8\xc5\xeb\x1e\x38\x6c\x46\x87\x8d\x5a\xbd\x73\x00\x59\x22\x5a\x12\x6e\x1d\xaf\x96\x4a\x43\x12\xae\x5e\xe7\xb8\xc1\x12\x69\x0b\x1f\x0d\xb9\x47\xed\x6f\x39\xa5\x29\xa7\x3f\xfd\x60\xdf\x53\x0d\xf6\xf1\x79\xac\xc6\x86\x43\xa9\x7d\x01\x88\x97\xa4\xc2\x23\xb1\x5d\x6d\x6f\x2a\xae\x6e\x75\xee\xf1\x88\x51\xd8\x30\x31\xe9\xd3\x19\x6b\xd0\x8a\x3c\x18\x0e\x5e\x9e\x0e\xa0\x79\xc3\xb3\xc0\xdd\x7d\x83\x34\x02\x7b\xee\xed\x3f\x95\xb2\x4f\x34\x69\x5a\xd0\xc8\xbf\x5e\xbb\x18\x8d\x91\x72\x7e\x52\x7d\xc7\xfe\x75\xfd\x2e\x4b\xb7\x75\xff\x0d\x00\x00\xff\xff\x23\xe4\x2a\x58\x37\x03\x00\x00")

func templatesStorageTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesStorageTf,
		"templates/storage.tf",
	)
}

func templatesStorageTf() (*asset, error) {
	bytes, err := templatesStorageTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/storage.tf", size: 823, mode: os.FileMode(480), modTime: time.Unix(1512163867, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTlsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x41\x0a\x02\x31\x0c\x05\xd0\x7d\x4e\xf1\xc9\x09\x5c\x88\xe0\xa2\x0b\xaf\xa0\x07\x08\xad\x04\x5b\x6c\xa9\x24\xb1\x30\x0c\x73\xf7\x79\xa6\x3e\xff\xf6\x56\x70\x74\x97\x9f\xb5\x95\x43\xe5\xab\x1b\x83\xcb\xf4\x2a\x6b\x38\x63\x27\x20\xf7\xcf\xb4\x16\x75\x20\x81\x9f\xaf\x07\x13\x60\x9e\xa5\xb4\x70\x20\xe1\x7a\xb9\xdf\xe8\xa0\x33\x00\x00\xff\xff\x52\x4d\xac\xad\x51\x00\x00\x00")

func templatesTlsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesTlsTf,
		"templates/tls.tf",
	)
}

func templatesTlsTf() (*asset, error) {
	bytes, err := templatesTlsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/tls.tf", size: 81, mode: os.FileMode(480), modTime: time.Unix(1512163867, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVarsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xd1\x4a\xc5\x30\x0c\x40\xdf\xfb\x15\xa1\xf8\x3c\x9d\x88\x6f\x7e\xcb\xe8\xda\x28\xc1\x2e\x1d\x69\x57\xc1\xb1\x7f\x97\xad\x50\xb5\x94\x7b\xb7\xb7\x9c\x73\xa0\x49\x36\x42\x66\xf6\x08\x1a\x39\x4f\xe4\x34\xec\x87\x52\xbf\x53\xc1\x0f\x0a\xdc\x4e\x23\x2d\xab\xc7\xa9\x9f\xc4\x6d\x8e\x56\x68\x4d\x14\xb8\x83\x13\xb2\xe1\xd4\x01\xd6\x13\xde\x02\x11\xad\x60\x6a\x21\x63\xfa\x0a\xf2\x39\x59\x72\xa2\x61\x57\x00\x0e\xdf\xcd\xe6\x13\xbc\x81\x1e\x9f\x86\xeb\x7f\x1c\x5f\xb5\xfa\x97\x11\x27\x14\x36\xfe\x4e\xf7\xfc\x72\x75\xab\x84\x4c\x0e\x05\xb4\xf9\xde\x04\x65\x29\x45\xb3\xe9\x59\x3e\xec\xd9\xc8\xd0\x80\x43\x2b\x80\xba\x37\x94\xaf\xca\x15\x5c\x5a\xbd\x42\xab\x55\xf0\x57\x2b\x37\xe9\x68\x05\x1c\xe7\xeb\x7f\x02\x00\x00\xff\xff\xf5\xa0\x00\x27\xe3\x01\x00\x00")

func templatesVarsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesVarsTf,
		"templates/vars.tf",
	)
}

func templatesVarsTf() (*asset, error) {
	bytes, err := templatesVarsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/vars.tf", size: 483, mode: os.FileMode(480), modTime: time.Unix(1512163867, 0)}
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
	"templates/cf_lb.tf": templatesCf_lbTf,
	"templates/network.tf": templatesNetworkTf,
	"templates/network_security_group.tf": templatesNetwork_security_groupTf,
	"templates/output.tf": templatesOutputTf,
	"templates/resource_group.tf": templatesResource_groupTf,
	"templates/storage.tf": templatesStorageTf,
	"templates/tls.tf": templatesTlsTf,
	"templates/vars.tf": templatesVarsTf,
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
		"cf_lb.tf": &bintree{templatesCf_lbTf, map[string]*bintree{}},
		"network.tf": &bintree{templatesNetworkTf, map[string]*bintree{}},
		"network_security_group.tf": &bintree{templatesNetwork_security_groupTf, map[string]*bintree{}},
		"output.tf": &bintree{templatesOutputTf, map[string]*bintree{}},
		"resource_group.tf": &bintree{templatesResource_groupTf, map[string]*bintree{}},
		"storage.tf": &bintree{templatesStorageTf, map[string]*bintree{}},
		"tls.tf": &bintree{templatesTlsTf, map[string]*bintree{}},
		"vars.tf": &bintree{templatesVarsTf, map[string]*bintree{}},
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

