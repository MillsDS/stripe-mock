// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package main generated by go-bindata.// sources:
// cert/cert.pem
// cert/key.pem
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

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _certCertPem = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x95\xb7\x0e\xab\x48\x18\x46\x7b\x9e\x62\x7b\x6b\x05\x18\x13\x5c\xdc\x62\x60\x30\x06\x4c\x18\x32\x74\x60\x6c\x72\x30\x69\xc0\x4f\xbf\xba\x77\xb5\x41\xbb\x7f\x79\x7e\xe9\x54\x9f\x74\x7e\xff\x79\xa2\xac\xa8\xe6\x6f\x92\xec\x78\xea\x4d\x95\x80\x27\xff\xa2\x84\xa1\xaa\xf2\x08\x25\x09\x0c\x58\x92\x10\x5c\xec\x77\xdb\x51\x2f\x07\x4b\x1e\x30\xc5\xa2\xf9\x94\x4d\xa5\x5c\x31\x25\x02\x34\xdf\x00\x04\xbe\xe1\xa8\x58\x06\x31\x0c\x10\x82\x10\xb4\x33\x91\x9d\xcd\x32\x53\xca\xed\xc9\x20\x7c\x2f\x9f\xa6\xe1\x15\xd8\xfc\xca\x8c\x51\x03\xca\xfc\x82\x6f\xf8\x93\xd5\xbf\x18\xf5\x37\xab\xff\x11\x11\x7f\x99\xfe\x2d\x2a\x0a\xb5\x32\x00\xa5\x48\xee\x47\x71\xd5\x8c\x81\x48\x16\x01\xf2\x01\xb8\xa8\x12\xc4\xe0\xe7\x5f\x07\x83\x2a\x01\x04\x45\x22\xf2\x4d\xe1\x1a\x81\x35\x88\x87\xb5\xe2\x39\xb7\xbd\x7d\x54\xef\xa2\xb3\x0f\x55\x9b\x8c\xd9\xb8\x7d\x19\xc7\xe7\x1a\x3b\xac\xcf\x77\x57\x12\x39\xfb\x82\xa7\xee\x56\x06\xb8\xfb\x56\xf7\x46\xa6\x3c\x42\xf7\xb3\x63\x8d\xae\xda\x45\xed\x1c\x26\x14\x8b\x47\x1e\x55\x62\x44\xae\x83\x20\x52\x29\x68\x1d\xe4\x4d\x9f\x5b\x73\x0b\xbd\x78\x0d\xb5\xed\x21\x2b\xc3\xf8\x90\x51\x70\xe8\x63\x71\x68\xcb\x2a\x11\x07\xa5\x73\x83\x41\x97\x67\xd7\xf8\x8c\x63\x29\x8f\xef\x98\xdc\x2f\x4a\xe7\x7a\xcd\xea\xfa\x38\x7e\xa2\x59\x1e\x26\x9d\x73\xa8\xb4\xaf\xed\x45\x13\x37\xc7\xec\x48\x9a\x27\xf3\xaa\xad\xda\x90\x23\xae\xe3\x92\x82\x41\xe6\x69\x33\x18\x52\xbd\x57\xb1\xf5\x39\x9a\xf7\xeb\xa5\x44\xfc\x6c\x22\x2b\x15\xad\xb8\x75\xe8\xc7\x54\x9c\x3a\x5c\x2c\xd1\x40\x55\xc5\x89\x6d\x9a\xac\xe9\x16\x4c\x17\x5b\x45\x68\xb6\xa6\x2e\x9d\xda\x36\x29\xc8\xc7\x03\x18\xb0\x33\xc0\x50\x83\x38\xea\x65\xae\x57\x87\xf9\x8a\x77\x92\xc7\x21\x4a\xdb\x23\xd8\x82\x3e\xe0\xda\x9b\x40\xbd\x27\x7b\x03\xcc\x35\x0a\x7b\x9e\x10\x83\x21\x87\x5c\x45\xdf\xae\x9d\x95\xfb\xcc\x5e\xb1\xac\xf7\x09\x64\xdb\xb0\xbe\xf9\xa8\x19\x80\x89\x50\x5f\x0f\x55\xa4\x69\xf9\x8b\x6d\x83\xaa\xe3\x8b\x2a\x2c\x83\xaf\xcc\x72\x07\x6d\xf5\x2b\x91\x14\xe7\x84\x54\x2e\x8b\xc4\x27\x6b\xab\x45\xba\x73\xd7\xc7\xb3\xb5\x87\xb8\x98\x62\xf6\x25\xa6\x43\x65\x9f\x11\xe3\xd3\x0d\xbd\x25\x51\x75\x49\xe8\x31\x6c\x8d\xb3\x1c\x80\x51\x4d\x62\x7d\x6a\x88\x83\x45\x9c\x9f\xcd\xca\x07\x41\xdd\x77\x0b\xa6\xfd\xf0\x07\x3f\x4a\x9b\xd6\x3a\xde\xa3\x80\x37\xe1\x99\x5b\x6f\x34\x0e\xb9\x56\xee\x43\x65\x30\xee\x6d\xea\x3e\x63\x8a\x6a\x30\x3f\x8a\xca\x72\x08\xbe\x2d\xf2\xb6\xa9\x9d\xf7\x6e\x27\xbc\xc0\x93\xfc\xd3\xb1\xe5\xeb\xf3\x9d\x0a\x08\xdb\xdb\x15\x9f\x9c\x59\xb8\x5d\x0c\x52\xcb\x1c\xd8\xed\x4b\xd5\x78\x05\xdd\x68\xaa\xf8\x18\xe1\x14\x4e\x6d\x4f\xc0\x4e\x10\xee\x01\xf0\xb0\x7d\xd1\xbe\xab\x46\xdf\xd1\x7d\x8f\xee\x8b\xf9\x78\xd6\x4b\x62\x59\x73\x24\xe3\x5e\x79\x8a\x90\x44\xec\xdb\xba\x9c\x55\x6e\xeb\xb2\xb5\x6b\x8e\xbe\x31\x26\x77\x32\x27\x82\x2e\x98\x04\x49\x13\xfd\x45\xf6\xf6\xa6\x2b\x72\xef\xa1\xb7\x1d\x6d\xff\x3e\x1c\x1c\x80\x18\xdd\xc7\x2f\x1d\x4a\xa0\x57\x91\x0a\x01\x02\xe2\x7f\x27\x4e\x48\xf8\xcf\x8d\x03\x24\xc5\x76\x37\x75\x67\x58\x8b\x17\x71\x02\xa5\xec\xb0\x51\xa1\x6d\x65\x5e\x39\xfc\x93\xd3\xf2\x6a\x79\xd5\xa3\xda\x59\x59\x39\x2d\x4e\x9a\x3c\x66\x79\x42\x86\x7c\x23\x8e\xb1\x9e\xb2\x94\x1e\xdf\x19\x1b\x37\xa7\xe1\xf9\x1c\x93\xf2\x3a\x31\xb3\x6d\x0a\x0a\xb9\x49\xc3\xf0\xa0\xcb\x24\x21\xc3\xd4\xd6\x8f\xcc\xf8\x7e\x9b\x71\x8a\x3c\x46\x73\x75\xf5\xb0\x64\xbe\xa3\x08\x45\x48\xe7\xdb\x46\x76\x08\x81\xc6\x7a\x9b\x3a\x7b\x46\x9a\xb7\x7e\x46\xa8\x90\xeb\x18\x3d\x31\x28\x56\xc5\xe2\xae\x07\x90\x38\x67\xee\xac\x3a\x96\x75\xe3\x13\x89\x6a\xa0\x25\x5a\xd5\xa0\x2f\x51\xb9\x51\x2a\xb4\x28\x91\x0c\x57\xad\xe1\xd3\x8a\x2c\xfd\xb4\xf6\xb2\x07\x9b\x47\x6e\x86\x46\x90\x0d\xa2\xd9\x0f\x05\x12\x8c\x43\x0b\x4f\x8c\x09\x5d\x2a\x86\xcb\x74\xba\x9c\x85\x2d\x21\x83\x82\xa8\x76\x7b\xe7\x5a\xd8\x74\xf9\x71\x78\x8f\x5d\x78\x65\x8f\x7b\xac\x9e\x00\x79\x5f\x59\x98\x5b\xcc\xd5\x5d\xeb\x3d\xc8\x33\xd6\x3f\x7b\x24\xde\xd6\x9b\xe8\x00\x7a\x61\x73\x18\x1e\xd9\x85\x6e\x15\x22\xaa\xc8\xbb\x28\xed\x43\x1d\x7b\x13\xaf\xc6\x62\x2f\x43\x3e\xba\x38\x8b\xc4\x16\xe7\x7b\x12\xe9\x45\xc9\x73\xc6\x89\x82\x49\x1e\xea\x8c\xcf\xcc\x0b\x2d\x74\x36\x9b\xb0\x9a\xc9\x19\x6e\x31\x5e\x88\x0a\x52\x55\xb5\x23\x94\xb4\xcc\x68\x21\xb1\xab\x0d\x77\x45\xd6\x77\x58\xc5\x33\xea\x24\x20\xf7\x94\xcb\xf0\x4c\x39\x48\x0a\xdb\x57\x35\xf2\x79\x45\x2b\xef\xd5\x57\xaf\xd6\x81\xa2\x62\x85\x25\x98\x9e\xd2\x40\xf8\xe1\xc8\x86\xe2\xd3\x78\xd7\x72\x9b\x86\xa7\xb5\x37\x66\x35\x64\x91\xed\xc5\x61\x8f\x48\x17\x61\xd7\x0e\x73\x4f\x6d\xf0\xb0\xb1\x5b\xbe\x64\x76\xce\x4d\xcb\x23\xe6\xa6\x98\xb8\x5f\xa4\xe7\xcb\x41\x67\x3f\xc6\x6f\xb9\x5e\xc2\x14\x07\xe9\xe7\xcc\x84\x8a\x32\x6f\x34\x09\xc1\x88\x17\x7f\xa2\x4f\x25\x7a\x61\x53\x09\x8f\x3e\xc0\x8d\x58\x59\x2c\x49\x9a\xd4\xe7\x2a\x7f\x1e\x84\xce\xcf\x2b\x63\x64\x35\x7c\xa5\x6e\xfe\x3d\xec\xec\xa1\xcf\x72\x69\x27\x9b\x2e\x1d\x43\x9c\x90\xf7\xfa\xc6\x08\x09\x7e\x71\x8d\x95\xb1\x6f\xb2\x28\x6d\xd0\x38\xc3\x3c\x3b\xb2\xef\xa0\xe3\x33\x11\x9b\x03\x6a\x6a\xc6\xf0\xd5\xbc\xc2\x68\x97\x26\x75\x7d\x0e\x2e\x43\x31\xb4\x1f\xb3\x63\x61\x8a\x2f\xd9\xc7\xb2\x99\xd0\x2c\x16\x4e\xc9\xe0\x52\xf7\x5a\xe4\x1d\xfc\xe3\x07\xf1\x2b\x2b\xb2\x09\xff\x9f\x9a\x3f\x02\x00\x00\xff\xff\xe6\x7b\xbf\xd8\x87\x06\x00\x00")

func certCertPemBytes() ([]byte, error) {
	return bindataRead(
		_certCertPem,
		"cert/cert.pem",
	)
}

func certCertPem() (*asset, error) {
	bytes, err := certCertPemBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cert/cert.pem", size: 1671, mode: os.FileMode(444), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _certKeyPem = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x97\xa5\x12\xec\x0a\x02\x05\x7d\xbe\xe2\xf9\xa9\xad\x30\x89\x15\x61\xce\x84\xc9\x05\x26\xcc\x9c\x7c\xfd\xd6\xde\x2b\xdf\x91\x47\x76\xb5\xe9\xff\xfc\x7f\xac\x20\x29\xe6\x3f\x96\xa3\x04\x8c\x27\xfc\xa3\x09\xf1\x9f\x17\x30\x14\x45\xb5\x2f\x85\x65\x78\xc6\x64\xab\x6e\xa9\xbb\x46\xa2\x2f\x88\x65\x6c\x41\x64\x18\x97\xe3\x5c\xe8\xaa\xaa\x6e\x66\x2a\x81\x61\x26\x85\x63\x6c\x9e\x8d\x7c\x93\xa2\x23\xe6\x08\xe2\xe9\x00\x1a\x92\x70\x7b\x71\x51\x3c\x4c\xc3\x75\x45\x5d\x8d\xcd\x10\x5f\xd4\xf1\x89\xce\x0a\x5b\x44\x76\x39\x96\xb0\xb0\x6b\x1d\xc4\x3a\xb8\x86\xb7\x91\x3b\x01\xf2\x34\x3f\x7b\x8e\x88\x56\x31\x65\x70\x00\x34\x64\x2b\xbd\x88\x1a\x36\x02\x8f\x89\x62\xa1\x94\xe9\x1d\xdb\x5b\x17\xb1\x13\x43\x2f\x3e\x42\xf5\xd4\x05\x69\x9a\x75\xc1\x0e\x1e\x6d\xae\x1e\x75\x3f\xb8\x07\xd2\x88\xc9\x80\x6b\xc4\x35\x16\x60\x9e\x6b\x61\x2e\x63\xf0\xc6\xa4\xc1\xf5\xba\xc3\xf5\xaf\x38\xb7\x37\x61\x5a\x35\xc2\x81\xd2\xb1\xb5\x76\x95\x3d\x1d\x73\x00\x61\x12\x2c\x9a\xbe\xe9\x43\x82\x9e\xf7\x94\x99\x04\x12\x36\x83\x09\x48\xb5\x51\xb9\xbe\xcb\xd3\x95\xbf\x9f\x14\x91\x9b\x69\x7f\x53\xf6\x1b\xf7\x0e\xac\xaf\xd5\x67\xb8\xaa\x3d\x9a\xa0\xa6\xfa\xe0\x5d\x97\x75\xc3\x7e\xc1\xd5\xd9\xa8\x96\xaa\xec\x83\xd2\x77\x29\x53\x00\xf3\xc3\x18\xfc\x60\x30\x53\xcb\xc4\xd1\x28\x10\xa3\x32\x6d\xf4\x75\x83\xe4\x15\xda\x69\xff\x04\x67\x30\x06\x44\x2f\x52\x50\xb9\x5a\x27\x83\xd2\x51\x38\x92\x6c\x30\x15\x3c\xd1\xc0\x22\x3d\x7c\x81\xc2\x47\xef\x06\xc7\xbd\x25\x10\x2c\xe3\xfb\x16\xb3\x6a\x30\x68\x64\x8f\xed\xd4\x44\xaa\x5a\xfc\xf0\x3e\x68\x06\xb2\x6a\xc2\x3a\x78\x05\x9c\x78\xe0\xef\x78\x24\x15\x92\x80\x12\xb6\x73\x64\x72\x00\xbd\x1a\x69\x8e\xac\xcd\xc8\xf7\x0e\xaf\x6a\x8d\xf1\x1f\x9b\x4e\x8d\x85\xd8\xa8\x0f\x77\xf0\x99\x44\x0d\x96\xc0\x73\xd8\x1b\x88\x10\x30\xb3\x92\xc4\xda\xda\x3d\xb8\x4d\xf8\xd9\x26\x2d\x36\xaf\x01\xbe\x5b\xa1\xfd\x42\x3e\xe4\xcc\x9d\x6a\xef\x78\x7a\xc5\x8b\x54\x5e\x7c\x4b\x7b\x9e\x0a\xb5\xbe\xa7\xc6\x40\x5d\x71\x1d\x96\x39\xb5\x5b\x66\xd3\xab\xe6\xeb\x90\x7d\x55\xf4\x5d\xeb\x94\xb7\x95\x00\x24\x45\x82\x64\xee\x58\x02\x9d\x97\x29\x65\x5f\xd6\x49\x5f\x1f\x67\xa3\x44\xcc\x00\xd5\xcc\xe1\x87\x7b\x6f\x3a\xaf\x82\x3b\x55\x61\xf5\x99\x5f\xc3\xb5\x1f\xf9\x81\xa2\xe4\x80\xf1\x2e\x0b\x53\x81\xf7\x50\x61\xd9\x96\xef\x48\xde\x4d\x3d\x6f\xf7\xe4\xfb\xdd\x22\xe1\x1a\xa5\x9c\xe5\x41\x1b\x2f\xbf\x18\xa2\x10\xe7\x90\x1d\x43\xf7\x8c\x9d\xb1\xba\xab\xb9\xc2\x15\x9a\xd8\xdc\x0a\xbf\xb6\x75\x02\x25\xdc\x80\xf7\xc8\x7b\xe7\xd3\x8f\xe5\xe3\x5c\x01\x13\xdb\xf2\xfc\xc2\x21\xc7\x8c\x8a\xad\xf0\x8c\xcd\xb0\x7f\x15\xe6\x62\x1d\x5a\x89\xaa\x13\xa1\xa2\x49\x50\x9f\x7a\xe9\x41\x86\x63\x20\xbe\x4f\x7e\xf8\xb1\x59\x31\x5c\xe0\x30\x1a\x58\x57\x54\xcb\xaf\xcf\x24\x0a\xf6\x88\xbc\xdc\x02\x8a\xfc\xfa\x92\x99\x60\x1f\x14\xde\x31\x48\xb7\xd6\x21\xcf\xed\xb2\x11\x0c\xeb\x0b\x5b\x3b\xda\x02\xc6\xaf\x5d\x32\x8f\x04\xa3\x5b\x0d\x2d\xef\x29\x24\x9e\xa4\xb4\xcc\xfa\xbc\x01\xb7\x59\x4f\x7e\xdd\x71\x03\x9e\x25\xde\x4a\xc3\xe4\x0c\x07\x43\x7f\x86\x6e\xc1\xb3\xac\x7d\x67\xf5\x16\x1d\x64\x00\xac\xaf\x71\xff\xcc\xb2\x33\x6a\xdb\xfc\x28\xf3\x99\x61\xc8\xeb\x8c\x08\xee\xba\xac\x41\xc6\x35\x7b\x1a\x60\xec\x4f\x21\x95\x30\xe0\x6e\x77\xca\xb2\x84\xb1\xa8\x3b\x22\xee\x40\x29\xde\x97\x57\x0d\xe0\x67\x85\x49\x03\x3f\xb4\x8f\xfe\x34\x31\xa7\x5a\x53\x30\xb5\x2f\xea\x34\x43\xa9\x32\x25\xde\x60\x87\x6d\x69\x90\x92\xc7\xea\x52\xef\xbf\x79\x69\x18\xbe\xe4\xb4\x09\xd3\xc5\xf7\x9a\x78\xc5\x03\xec\xc4\xb0\x91\xc3\xd2\x3e\xb1\x18\x1e\xc3\xc8\x88\x87\xf1\x35\x86\x4a\x60\x1a\xea\xe3\x06\xdd\xad\x70\x1a\x57\x18\xdf\x82\x8a\xbe\xfc\xc7\x4f\x4a\xc3\x02\x53\x1c\x4f\x90\x59\x73\xca\xeb\xa5\x56\xe0\xba\x64\xb4\x46\xfa\x49\xc0\x57\xc1\xcf\x84\xed\x5a\x23\x4b\xcc\x8b\x4b\x0e\x28\x01\x15\x89\x4b\x3d\x19\x68\xd3\xcf\x35\xcd\xb0\x26\x3a\x55\x44\x8a\xcc\x25\xc6\xc2\xd4\x14\x0e\x17\xfc\x51\x2a\xb0\xaf\xb6\x1e\x88\xad\x3e\x93\xd8\x27\xaf\xd8\x2b\x38\x4e\xcb\xf0\x64\xe5\xf1\x05\x2b\x87\x6e\xbd\x2e\xbd\xd4\xf3\xce\xc8\x7b\x4d\xd4\x2e\x4b\x91\x3a\xdd\x51\x79\xad\x46\x74\xdc\xd5\x1a\x5b\x09\xd8\x1d\xe3\xe0\x35\xbf\x75\xf5\x5f\xce\xfb\x45\x7e\x15\x05\x8a\xc9\xd0\x1d\x8a\xed\xd9\x6e\x57\x82\x6b\x56\x7a\xf5\xee\xf0\x72\xa7\xed\xca\x94\xa4\xc5\x51\xbd\xa3\xe0\x02\x87\xfc\x38\x69\x93\x0c\xd0\xe9\x7d\xa4\x60\x5b\x6a\x27\x9e\x4a\x4f\x46\x91\xb3\xa0\x3d\x59\x76\x15\x2d\xcc\x5d\x29\xea\x48\x60\xe2\x1b\xb0\xb6\xf1\x3f\xa1\x0c\x20\xa3\x9c\x79\x44\x8a\x99\x58\x28\xd5\x45\xf0\x42\x93\x06\x7e\x37\xde\x47\xf0\x75\x41\xab\xed\x1e\x65\x9c\x23\x4c\xf3\x38\x26\xa9\xb7\xbb\xc6\xc7\x62\xa5\xa7\x3a\xe5\xf8\xa0\x9c\x5a\xfa\x28\x28\x70\x31\x12\xd8\x1a\x6c\x06\xd6\xc3\x13\x73\x11\xe2\x12\x03\x53\xb7\x6f\xbe\xf0\x0b\x47\x6f\x0c\x97\xca\x39\x12\x28\x18\x29\x5b\xe3\x18\x5b\x60\xb0\x7d\xc0\x7a\xa6\xd8\x3f\x28\x31\x0d\x6f\x21\xae\x0e\x1b\x0c\x8e\x04\xc3\xa9\xfb\x71\x63\x7a\xc4\xf6\xd7\x05\xbc\xb4\x59\xe1\x6e\x94\x7e\xe5\xdd\x52\x6a\x66\xf7\x27\x0d\x0d\x15\xce\x4a\x2b\xde\x4e\x94\x78\xf0\x44\x9b\xa5\x7d\xd0\x92\x51\x24\x73\x5a\x90\x20\x65\x7d\x9e\x36\x7a\xdb\xb1\x45\x8b\x9c\xcf\x02\xe0\x7a\x49\xa6\xbb\x96\x9c\x03\x1a\xab\xe1\x1a\x93\x71\x1d\xfb\xe2\x48\xbf\xa6\x63\x75\xd0\x31\x7f\xa1\x3e\x42\x6e\xf2\x7d\x44\x4e\xcf\xd7\x81\xd6\x9d\x5a\xe6\xdb\x43\xdd\x3d\x0a\x1b\x50\xa2\x00\x5c\xd2\x8b\x38\x4e\xa5\x8d\xa6\x99\xd1\xf0\xdc\xa2\x62\xbd\xb8\x26\x6c\xab\x6b\x41\x78\x7d\x96\xfa\xcf\xfc\xe9\xfb\x5a\x1b\x35\x7f\x29\xf5\x82\x8c\x6b\xcd\xbb\x3e\x15\xaf\x81\x58\x2d\xe3\xb5\x0f\xc0\xe0\x38\x36\xb6\x26\x8f\x61\x84\x99\x14\xa1\x87\xa8\x4e\xa0\x96\x69\x97\x9c\xea\x12\xf3\x69\x70\xc5\x67\xa2\xbc\x9c\x50\x9c\x18\x11\x9d\x69\x18\x48\xbd\xfd\xdc\x52\x2f\xdd\x77\x21\x78\x4d\x00\x64\x3f\x5e\xb0\x45\x2e\x83\x11\x1a\x2b\x43\x6f\xc5\x0a\xe9\x94\x18\xba\xc2\xfd\x27\xd2\x69\x6d\xe2\x57\x08\x0e\xc6\x1b\xf4\xca\x6b\x21\x42\xf1\x93\x5c\x7f\x01\x23\x6d\xcc\x6d\x62\x5b\x03\xb7\x0e\x1f\x4a\x6f\x71\x53\xbd\x9f\xc3\x41\x37\x58\x0e\xf7\x7b\xf8\xbc\x35\xca\xf6\xd1\x1c\xb9\x22\xff\x6a\x34\x1a\xd4\xf8\x1a\xba\xe3\x94\x08\xd4\xc7\xd2\x33\x99\x75\x99\xa3\xdf\x2e\x2f\xd1\x81\x38\x1b\x42\x03\x29\xf5\x6e\x76\xf6\x7b\x54\x2d\x9c\x4c\x86\x38\x67\xd5\x70\xc9\x72\x27\xc8\x2a\xef\x33\x13\xf8\xb3\xea\xdd\x5e\xd5\x73\x6d\x77\x79\xa4\x6a\xad\xdf\x0b\xaf\xe2\xce\xc5\xca\xf9\xc0\x51\x96\xfc\xe4\xe7\x36\xe4\xd6\xe0\xac\xdb\x79\x77\xf2\xfe\x24\x2c\xa0\x23\x43\xb6\x60\x97\x75\xbe\x43\x78\x8b\x9c\xda\x2f\xdb\xae\x78\xb0\x90\xa4\x18\xf1\x59\xfa\xf2\x3f\x4e\x16\x65\x4d\x21\x00\x54\x8e\x30\x8d\x0d\xcb\x87\xbd\x52\x96\xfc\x21\x09\xc8\x84\xae\xd2\x43\x24\x01\x83\x99\xf8\x81\xe3\xef\xfb\xb0\x31\xc8\xcc\xf0\xc3\xef\x5c\xdd\x46\xf1\x5c\x40\xd7\x6a\xc2\xd7\x0a\x4f\xf2\x67\x05\xba\x91\x8f\x52\x5b\xba\x3f\x29\x41\x5b\xe7\xef\x07\x9f\xde\x75\x9d\x09\x38\x8d\x0f\x2d\x99\xef\xad\xaf\xaa\x11\xaa\xef\xb1\x05\x8e\xfe\x91\xe2\xcd\xad\xdf\x2d\x65\xa2\xaa\x24\x5e\x76\xfe\x8e\x14\xe0\x2e\x43\x4f\x40\x49\x2e\x7f\xdd\xd9\xfc\x43\x98\x65\x21\x18\x29\xf8\x95\x4c\x7d\x79\xd0\x7e\xba\x7a\x18\x8b\x0d\x31\x38\x97\x70\xa7\x7c\xf5\x32\x1c\xed\x0d\x66\x58\xe6\xa2\x88\x47\xe2\x00\x74\x78\x0e\x23\xf4\xec\x30\xdf\x60\x2d\xec\xb8\x34\x1c\x57\x69\x74\x58\x59\x96\xf7\xfd\x7e\xf4\xce\xfd\x46\xaa\x74\x92\xa6\xf5\x9e\x75\x25\x5c\xc7\xe2\xdb\xa1\xc4\x2b\xd0\xeb\x06\xe4\x59\xf2\x80\x63\x97\x88\x53\xe3\x5b\x31\x9a\xfc\xd2\x21\xa6\xf4\xe8\x87\x20\xbe\x1f\xc8\x0b\xfa\xb0\x17\x63\xde\xa6\x70\xb4\x41\xe5\xd7\x69\x66\x07\x3f\x58\x3e\x72\xc2\x69\x1e\x2f\xe5\x78\xc2\x73\x35\x65\xa0\xa4\xf2\x67\x69\xcc\x33\x51\xe1\xed\x16\x3b\x52\xe3\x0f\x31\x04\xed\x45\x3c\xd9\xdf\x58\xc5\x55\x23\xb4\x31\x96\x0b\xe9\xfd\x6e\x64\xa3\x6b\x1a\x29\xc7\x3f\xf9\xf7\xc6\xdb\x39\xd3\x73\x47\x8b\x00\xca\x46\x28\x44\x24\x69\x5d\x46\x72\x1e\x53\x90\x53\xc3\x64\x28\xdb\x56\x1f\x56\x72\xb9\x56\xdf\x94\x7a\xf4\xc3\x2e\x3b\x90\xcf\xf6\x68\x73\x5a\x75\x5a\xc3\x4f\x08\xa7\x54\xa9\x57\xc3\x66\xee\x00\xf5\x0e\x4b\xec\x7f\x03\x67\xb8\xea\xa2\x79\x8f\x53\x2d\xbf\x01\xd9\x7f\x8b\x24\x9b\xc7\xf5\x39\x24\xd4\x80\x2c\x2c\x8f\x98\x49\x61\x19\xa6\x58\x2c\x89\x1a\xe7\xea\xce\x33\xac\xe9\xc5\x4a\x70\x01\x03\x1f\x76\x73\x3b\xc9\xaa\x60\xe5\x5c\x4c\x5c\x4c\x22\xe1\xb8\x87\xda\x95\x8c\x10\x90\x3b\x76\xa8\x78\x46\x3d\xd4\x26\x1b\x9c\xf3\x70\x85\xfc\x1e\xab\x69\x7b\x77\x19\xcf\x0a\xfc\xef\x7c\xdc\x40\x1d\x07\xe7\x7b\xc2\xe6\x31\x90\x33\x69\xda\xcf\x1d\xec\x33\x0c\x17\x20\xea\xef\xfa\x5a\xc2\x4c\xfe\xbd\x25\xd8\x52\x6e\xea\x7a\x62\x96\xea\xd1\x48\xcd\x42\x02\xca\x9a\x12\xad\xfd\xfa\xab\x4b\xc0\xa4\x94\xa2\x7d\xd1\x96\x99\xcd\xd8\x4a\xf9\x37\xbb\x49\x74\x52\xf0\x7e\x9a\x16\x36\x88\x88\x75\x5f\x39\x0e\x0e\xa5\xf6\xa5\x45\x57\x42\xb1\x92\x17\xd1\x4b\x4f\xef\xef\x67\x1b\x55\xa2\x37\x2c\xe0\x17\x66\xe8\xc9\x33\x26\x18\x5b\x59\xf9\x69\x59\x29\x9e\x9b\x6f\xeb\x96\x1f\x15\xde\x6f\x76\x5b\xdf\x45\x9e\x87\x90\x34\xdd\x0f\x3f\xb2\x75\x58\x8e\x6b\xde\x95\x28\xa4\x67\x73\x82\x62\x50\x91\x01\x17\xf9\xc8\x2c\x6e\x31\xb1\x3d\x7d\x41\xa1\x0f\xbf\xf9\x2a\x96\x10\x97\x7c\x64\xfb\x48\x6d\x45\x1e\xbb\x9e\xef\xb1\xcf\xad\xfd\xf6\x06\x24\x5f\x84\x3c\xa4\xdf\x59\x90\xe0\xb4\x11\x04\xde\xd7\x33\xc0\x18\x06\x57\x55\x02\xcb\xe8\xc5\xee\x74\xa5\x42\xf2\x3b\x92\xd4\xc4\x2a\x12\x19\xf5\xb0\xb3\x45\x8d\xca\xf2\x16\x70\x83\xd5\x17\xb2\x5a\x45\x21\xf7\x2a\x4a\x59\xc5\x1c\xf8\xd7\x71\x74\xf5\x01\x0c\x56\x8d\x7b\x55\xd4\x2a\x47\x06\x6f\x4f\xab\x24\x58\x99\xa7\x22\x18\x1d\x15\xba\xbc\xa9\x22\x0d\xf7\x23\xfe\xe0\x61\x4f\x52\xa5\xf7\x15\x41\x46\x9d\x2a\x11\x74\xdb\x98\x0c\x12\xaa\x3f\x7d\x02\x2c\x2f\x43\x68\x93\x39\x92\xd7\x64\xc5\x1c\x66\x4c\x1a\x67\x73\xa3\x2f\x53\xf4\x80\xbe\x90\xe3\x52\xee\xf8\xb9\x11\xe2\x09\xd8\x5d\xbb\x1c\xdb\x0e\xac\xaa\xba\x95\xf7\x09\x79\xa2\x04\xed\xf9\x03\x24\x95\xaa\x6a\x64\x8b\x04\x2a\x68\x52\x0e\x7b\x47\x2d\xa2\x51\xdc\x5e\x86\x4d\x7c\x53\xa6\x0c\x93\xdc\x8f\xd4\x0e\x5d\x50\x98\xb8\x40\x2b\x7e\x3d\x21\x01\x1e\xe7\x76\x07\xeb\x71\x6c\x3e\x47\x02\x28\x03\x21\x7d\xd1\x2b\x25\x7f\x09\x7d\xae\x02\x06\x73\xbe\x17\x20\x1b\x42\x77\xae\x81\x0b\x2e\xe3\x73\x9b\x95\xbd\xfe\x1a\xaa\xea\xbb\x07\x2f\xf6\x85\xe4\x4e\x08\x30\x9f\xec\xfa\xef\x97\xf3\x4b\x60\x36\x8f\x99\xc8\x09\xde\x7c\xd5\xe6\xa2\x65\x13\xfe\x09\xbe\xc8\x9a\x7b\x13\x04\x8a\xb3\x2a\xff\x05\xfe\x74\x89\x60\xf2\xff\x6e\x95\xff\x05\x00\x00\xff\xff\xb6\xd9\x3f\x1c\xc8\x0c\x00\x00")

func certKeyPemBytes() ([]byte, error) {
	return bindataRead(
		_certKeyPem,
		"cert/key.pem",
	)
}

func certKeyPem() (*asset, error) {
	bytes, err := certKeyPemBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cert/key.pem", size: 3272, mode: os.FileMode(444), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	"cert/cert.pem": certCertPem,
	"cert/key.pem":  certKeyPem,
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
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"cert": &bintree{nil, map[string]*bintree{
		"cert.pem": &bintree{certCertPem, map[string]*bintree{}},
		"key.pem":  &bintree{certKeyPem, map[string]*bintree{}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
