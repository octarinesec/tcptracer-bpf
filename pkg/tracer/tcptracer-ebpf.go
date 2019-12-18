// Code generated by go-bindata.
// sources:
// ../dist/tcptracer-ebpf.o
// DO NOT EDIT!

package tracer

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

var _tcptracerEbpfO = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5c\x7d\x6c\x5b\xd7\x75\xbf\x8f\x14\xad\x0f\xc7\x95\x5d\x8f\x09\xcd\x2c\x9b\x9c\x8f\x55\x13\x02\x47\x92\x65\x9b\x56\xbe\xe4\xa4\x49\x14\xcd\x9d\x15\x57\x6c\x34\x6f\x01\xcd\x52\xb4\x49\x4b\xb6\x69\x91\xb6\xf5\x24\x63\xd3\xd2\xa6\x33\x88\xfc\x21\x37\x19\x26\x64\x41\x27\x4a\x4a\xa2\x76\xc1\xe0\xad\xed\xec\xa1\x19\x64\xb8\xc5\x2a\x6c\x05\x66\x64\xd9\x20\x04\x1e\x20\x0c\x41\xa1\x66\xc5\xa0\x05\x86\xa7\x06\x46\x38\xf0\x9e\xdf\xe3\x7b\xef\xbc\xfb\x1e\x29\xdb\x6d\xff\xf1\x03\x92\xcb\x73\xde\x3d\xf7\x9c\x7b\xee\xb9\xe7\xfe\xee\x7d\x57\xfe\x93\x67\xf6\x3e\xeb\xd3\x34\x61\x3c\x9a\xf8\x3f\x61\x52\xe6\xd3\xff\x82\xf9\xbb\x0b\xff\x7f\x50\x68\x62\xfe\x6e\xe2\xbd\x22\x84\xf8\x9c\x10\x62\xac\x61\xa5\x58\xa2\xf5\x78\x46\xf2\xc7\xc2\xab\x92\x9e\x9f\xa6\x7a\xb5\x3e\x21\x56\x8a\xc5\xe2\xfc\x0c\x68\xbf\x10\xab\xc5\x62\x31\xc4\x94\x5e\xa8\x31\xdb\xf5\x95\x68\xf0\xff\x1c\xa5\x7e\x4f\x2f\xd3\xdb\x2d\xf5\x5c\x40\x3b\x63\xe1\x2e\x87\xde\x6e\x85\x9e\x57\x64\x9f\x85\x08\x8a\xcf\xc9\x37\xfa\x3a\xe2\x57\x23\xe7\x17\x42\x1c\xae\x15\xa2\x49\x08\x71\x1a\x65\xb4\x36\xa0\x71\xf9\x2e\x0f\xbd\xf3\xb5\x44\x07\x6b\x3f\xa3\x7e\x9d\x00\xad\xdd\x20\xba\x80\x7e\xf9\xd0\xaf\x96\x56\xf2\x6f\x16\xf5\xfc\x2f\x48\x3b\xa2\xfe\xeb\xb2\x3d\x7d\xb8\x4e\xf2\x5f\xfa\xfc\x35\xf2\x63\x3d\xe4\xde\x58\x2e\x52\xb9\x84\x72\x11\xe5\x15\x94\x0b\x28\x2f\xa1\xbc\x88\xf2\x3c\xca\x39\x94\x53\x28\x27\x51\x4e\xa0\x3c\x8b\x72\x1c\xe5\x48\xd9\x6e\x9f\xb4\x3b\x45\xfc\x70\x86\xec\x3f\xb1\x11\xe3\x44\xf5\xf4\xe1\x10\xfa\xd7\x4c\xf5\x5a\xc6\xc1\x6f\x02\xbf\x09\x7c\xd2\xa3\x67\x9a\x89\xdf\x40\xfa\xf5\x93\xad\x44\x1f\x20\xbb\xf4\x5c\x84\xfc\x37\x42\xd1\x3a\x96\x87\xfd\x3d\x64\xbf\x3e\xd2\x4d\xef\x87\x7b\xd1\x3e\xfa\x9b\x3f\x6f\xf3\x43\x7c\x38\x27\xdf\x27\x5a\xae\x82\x1e\x01\x4d\x7e\x4c\x0f\x9f\x91\xf4\x60\xcb\x47\xa0\xff\x18\xf4\x4f\x51\x3f\x83\xfa\x8b\x78\x3f\x8e\xf7\xcb\x78\x3f\x20\xe9\xc3\x3e\x8a\x9f\xb8\x7e\x50\xd2\xfb\xbe\x49\xe3\x16\xd7\x53\x54\x66\x87\xa8\x9e\x9f\xea\xed\x7b\x0d\xe3\xda\x89\xf1\x6b\xa1\xf1\x4b\x6a\x71\xe9\xef\xa0\x76\x06\x71\xf4\x4d\x8a\x6f\x4d\x93\xfc\x80\xf8\x5b\x8a\x0b\xc4\xe3\x60\x78\xbf\x94\x6b\xfc\x06\xd1\xf3\x05\x2a\x6b\x35\x21\xf6\x17\x8b\x45\x63\xfc\xe6\x13\x66\xdc\x96\xa6\x66\xba\x40\x72\x83\x61\xea\x47\x40\x7c\x0b\xfd\xec\x2f\x52\x79\x83\xde\xb7\xf4\xa1\xbc\x0e\x3b\x53\x78\xbf\x8a\x38\xa0\x71\x37\xda\x6f\xdc\xec\xb4\xa3\x1f\x76\xd4\x58\xf4\xeb\x27\x9a\x2a\xca\xdd\x50\xca\x35\x57\x94\xeb\xb3\xf4\xdb\x94\x6b\xad\x28\x77\x5d\x29\x47\x71\xd8\x78\xb7\xb3\x7e\x0a\xf5\xeb\x14\xfe\xd5\x4f\x50\xdc\xea\x33\xe4\x2f\x95\xbe\x55\x45\xff\xe2\xb3\xf0\x7f\x0f\xe5\x89\xf8\x0c\x8d\x43\xa2\xb3\x83\xe2\x6f\xda\x18\x8f\x08\xd1\x85\xeb\x18\xc7\x4e\xaa\xff\x16\x8d\x4b\xe2\xc0\x63\xb2\x0c\xd6\x0c\xda\xe2\xee\x34\xca\xa0\x9f\xe2\xf2\x45\x9f\x10\xc5\xa2\x10\x41\xdf\x21\xa2\x35\xd0\x1a\xc5\xb5\x91\xbf\xc6\xc2\x29\x96\x9f\x47\x6c\xf9\xc0\xda\xaf\x0c\xfa\xb5\x11\xfd\xda\xc8\xf2\x69\x2b\x5b\x3f\x52\x8a\xfc\x1a\x10\xfb\x65\x19\xd4\xbe\x20\xf3\x63\x50\x7b\x5c\xfa\xa7\x34\x0f\x02\xf2\x3d\xcd\x7b\xc3\x9e\x44\xb8\x0f\xf6\x90\xff\x1a\x9f\x75\x8f\x43\xef\xf1\xea\x77\x1d\xaf\x3e\xc5\x78\x1d\x16\xf0\x2b\xca\x52\x3f\x3e\x29\x16\x8b\x46\x3f\x5e\xda\xd8\x56\xb6\x53\x93\xf6\x5d\xc2\x7c\x6b\x65\xf6\xd3\x38\xeb\x05\xca\x9f\xaa\x78\xb3\xce\x07\x6e\x7f\xbc\x80\x38\x09\x5f\x45\xfb\x0f\xda\xda\x1f\x0c\x53\x9c\x34\xae\x73\xb6\x7b\xdd\x33\x4f\x18\xf1\xf5\x11\xda\xdd\xc2\xec\x46\x1e\x28\x34\xb9\xda\xbd\xea\x69\x37\xe2\x35\xbc\x84\xf6\x37\x30\xbb\x0f\x51\xbb\x5f\x71\xb6\x7b\xc8\xd3\xee\x43\xb0\xfb\xa7\x68\x97\x16\xe0\x78\xa1\x0f\xfa\x16\x1d\xf8\xa2\x9a\xb8\xac\x84\x6b\x7e\x6d\x78\x2a\xc4\xf1\x54\x2f\x9b\xaf\xdd\x0e\xbd\xbd\x5e\xb8\x06\x71\x12\x5c\x47\xa1\x5c\x8d\x5c\x69\x9e\xea\x87\x84\xa3\xbe\x17\x6e\x33\xf1\xd3\x27\xd4\x8f\x32\x7e\x5a\x21\x7a\x9a\xf7\xa3\xa9\x68\xaf\x17\x25\xfc\xa4\x7d\x4c\xf8\xe9\x04\xf0\xd3\xd6\x65\x61\x97\x03\x5e\x0a\x63\xbd\x0d\x03\x2f\x61\x3e\x8e\x85\x81\x1f\xc2\xc0\x0f\x61\xe0\x8d\x30\xf0\x52\x18\x78\x29\x0c\xbc\x14\x06\x5e\x0a\x03\x2f\xb1\x7c\x58\xce\x97\x7e\xe0\xa7\xce\x7e\xac\x9f\x07\x81\x8b\x08\x3f\xe9\x59\xe0\xa6\xce\x8c\x6d\x7d\xd5\xb3\xb4\x3e\xea\x39\xe0\xa4\x1e\xe8\xe9\x04\xde\xca\xb6\xe2\x7d\x04\xef\x61\x57\x27\xf0\x55\x16\xf9\x2c\xd7\x8d\xf7\xe8\x47\x27\xf0\x55\x16\xb8\xa9\x73\xce\xd6\xef\x78\x0e\x78\xa9\xe7\x27\xa0\x81\x97\x7a\xc8\x6f\xe9\x1c\xf0\x52\xcf\x07\xa0\x81\x97\x7a\xfe\x03\xf5\x81\x97\x7a\x16\xf0\x1e\x78\xa9\x67\x11\xef\x81\x83\x6a\x80\x97\x4e\xd2\x3a\xb4\xaf\x0f\xe3\xd4\x03\xdc\x96\x23\x1c\x15\x3f\x09\x7c\x15\x00\x6e\x7a\x06\xf5\x0e\x18\xe3\x15\x41\x49\xeb\x65\xd4\xf7\x6d\x51\x0a\xdd\xf9\x77\x10\x7f\xb5\x42\x5c\x2c\x16\x8b\x8d\x3b\x11\x97\x96\xfc\xd1\x6a\x59\xa7\x1c\xeb\x41\x01\xed\xde\x43\xf1\xa4\x17\x5a\x41\xa3\x1d\x4b\x7c\x37\xb1\xf9\xdb\xaf\x5c\xcf\x5e\xa7\x7e\xd4\x61\xbd\x40\x19\xad\x7b\x8d\xc5\x29\xf4\x20\xce\xc7\xc2\x2b\x28\x97\x91\xcf\x3e\x45\xb9\x0a\xfe\x12\xf2\x19\xd6\x8f\x13\x34\xae\x1c\xe7\x94\xec\xda\xe8\xd2\xff\x3a\xdb\x3a\xe8\x2e\x1f\x62\xf2\x4d\x6b\x90\xb7\xca\xad\xac\x51\x6f\x1d\x93\x5f\x56\xca\x57\xc6\x85\x9f\xde\x24\x2e\x5c\x5d\x23\x2e\x5c\xaa\x0a\x67\x2c\xb9\xe2\x0c\x25\xee\xfd\x36\xe2\xe2\x0d\xca\x23\xfa\xdb\x88\x8f\x23\x94\x4f\xe6\xe3\x24\xbf\xef\x1c\x95\xfa\x3b\x88\x9b\x51\xe4\x93\x39\xec\x1f\xf3\x94\x67\xd2\xb3\x88\xa3\x9e\x11\xe0\xca\x55\xe0\xca\x51\x86\x3b\xcf\x10\x8e\xd4\x9e\x20\x3b\xff\x14\x7a\xca\xf9\xf7\x51\xe2\xef\x21\xda\xc4\x8f\xe4\x9f\xf9\xad\x9c\xbf\x43\xd2\x63\x3d\x75\xc8\x47\x1b\x91\xf7\x42\x2c\xcf\x77\xd9\xe6\xb7\x5e\xe8\xb5\xcd\x73\xab\xbf\x22\x0a\xbc\x19\xad\xa7\x05\xcc\xf0\xcb\x61\x0d\xf3\x0e\x65\x48\xc2\x01\x13\xa7\xdd\xb7\x95\x16\xa2\x68\xed\x16\x5b\x3f\x2b\xc9\xbd\xb4\x95\x06\xf0\xf4\x3a\x7a\x9f\x38\x42\xe3\x62\xcc\xef\x44\xfe\x41\xd8\x4f\xfd\xbd\xaa\x91\x7d\x83\xe1\x66\xf0\x43\x8c\xdf\x02\xfe\x46\xcc\xeb\x87\x1d\xeb\x36\xcf\x37\xad\xca\x7c\x83\xfe\x5b\xe4\xfa\xab\x90\x5b\x0b\xbe\xa9\x93\x38\x23\xc2\xc6\x8d\xf2\xbc\x7e\x98\xf2\xba\x9e\xca\x38\xec\x58\xf4\xc2\x03\x38\xef\x08\xd6\xdf\x47\xf2\x88\xe7\xa8\x76\xaf\x5c\xe7\xe7\x4f\xa8\xc7\xa5\xb4\x2f\xa8\x95\xf5\xee\x96\xed\xf0\x38\x1a\xc4\xfa\x56\xed\xb9\x4e\x10\x78\x51\xc7\xfb\xa0\xb6\x41\xea\xbf\xfd\xed\xe6\xa9\x5d\xe6\x67\x5d\xa3\x75\x7b\xfe\x6b\x44\xab\xf2\xc4\x79\x45\x5e\x4a\x17\xc8\x9e\xa8\xf6\x71\xd1\x67\xc3\xe9\xe7\x11\x4f\x2b\xb0\x7b\x09\xe5\x8d\xa2\xb5\x3f\x09\xac\x33\xfa\xb9\x50\xd5\xfa\x6d\x79\xea\x5c\x53\x45\xb9\x15\xa5\x5c\x73\x45\xb9\x25\x55\x1e\x3e\xd7\x5a\x51\xee\x86\x52\x0e\xf9\x3b\xe4\xac\xdf\x65\xcd\xdf\x5f\x33\xc7\x8f\xe4\x8c\xfc\xdd\xe5\x9a\xbf\x97\x95\xfb\x7a\xf8\xbf\x67\x12\xf9\x75\x05\xf9\xf5\x4d\xe4\x5f\x8c\x47\x0b\xe1\xb4\x34\xf6\x73\x83\xe1\x69\xaa\xff\x16\x8d\x4b\xe2\xc0\x2c\xe5\xe3\x9a\x7f\x90\xa5\x73\x5f\xff\xf7\x92\x6f\xee\xeb\xbf\x47\x74\x39\xff\xfe\x1d\xc9\x39\xce\x3d\x69\xfe\x04\x44\xbb\xe0\xf1\x3c\xe9\x19\xcf\x84\x1f\x2f\x00\xc7\x8f\x8d\x02\x4f\x8f\x02\x9f\x8d\x9e\x77\xe4\x6b\x23\x6e\x24\x9e\xf0\xdb\xfd\xfb\x8a\xb4\xa1\x94\x57\xc8\xbf\xfa\x4c\x44\xd9\xcf\xb1\x4e\xf3\xbc\x58\x93\x71\x4b\xeb\x43\xbc\xf0\x3d\xc4\x71\x1f\xe8\x49\xd0\x29\xd0\x6f\x82\x1e\x82\x9f\xa7\x58\xfe\xc5\xfa\x58\x98\x66\xfc\x61\xc8\xcf\x42\x3e\x47\xe5\xe8\x08\xca\x33\xc8\xdb\xc6\x7e\x02\x38\xbf\x60\xe0\x54\x5a\xb7\xe7\x67\xe1\x87\x1a\x8a\x33\x63\x7d\x0a\xf9\xec\xfe\xbd\x10\xa0\x73\xfa\x92\x3f\xb6\x54\x39\x1e\x7e\xe9\xbb\x53\xa4\xf7\x96\xf3\xc7\xd1\x62\x83\x72\x3d\x9e\xbc\x39\x5c\x7a\xae\x57\x69\x8f\x0a\x97\x76\xa9\xf0\x9d\x87\x3c\xc7\xa5\x93\x6b\x90\xaf\x88\x4b\x3d\xf4\x56\x85\x4b\xab\xc8\x6b\x4a\x5c\x5a\x45\x5e\x53\xe2\x52\x8f\xbc\xb6\x54\x55\x5e\x5b\x23\x2e\x2d\x74\xd9\xf6\x8f\xfa\x5b\x88\x0f\xec\xcf\xf6\x05\x05\xda\x45\xbc\x74\xd2\xfc\xd0\x67\x81\x47\x8d\x7d\xe2\x34\xe2\xa7\x05\xfb\x4c\x9c\x0b\x0d\x86\xdf\xa7\x79\x37\x05\x3c\xda\x40\xfb\xce\x60\x60\x0b\xb5\xdf\x41\xed\x07\x6b\xee\x61\x79\x2f\xc8\xf2\xde\x66\x97\xbc\xb7\x4e\xe6\x3d\xeb\xfc\x3a\xef\x39\xbf\x1a\x1c\xeb\xbd\x5b\x7d\xca\x8f\x35\x2c\x3f\x62\xbf\x3d\x8a\xfd\xf8\xe8\xa4\x23\x3f\x4e\xfe\x52\xf3\xe3\x9b\xb6\xfc\xa8\x63\xbe\xeb\xd3\xc8\x53\x2d\xf6\xf3\x0b\x23\x9f\xe9\xd3\x38\x2f\x69\xb1\x9f\x77\xa4\x0b\x57\x58\x9e\x3c\x0b\xfe\xfb\x8c\x9f\x87\xfe\x0f\xa0\xff\x55\xe4\xcd\x09\x94\xaf\x43\xdf\x1c\xcb\x9f\x53\x37\x9d\x3f\xfb\x2d\xf9\x33\x20\x7e\x48\xf2\xc0\xaf\x1c\xaf\x9a\xe7\x66\xad\x65\x5a\xce\x93\x3a\xc4\x6f\x7a\x8d\x38\xb6\x7c\xae\xf5\x30\xc9\x03\xb7\x46\xb5\x16\x19\x47\x63\x47\x08\xff\xcf\xe3\xbb\xcd\x05\xe8\x19\xcb\x93\x9f\x13\x79\xf2\x93\xb1\xcf\x9c\xc7\xf7\x22\xd5\xbc\x5c\x50\xee\x87\x8d\xf9\xbc\xe0\x3a\x9f\x3f\x50\xcc\xe7\xb1\x3c\xc5\xd3\x60\x7e\xca\x11\xe7\x91\x35\xe1\xda\x46\xc2\xeb\x29\xf8\xef\xed\x66\xb6\x4f\x80\x9e\xf0\xad\xea\xf9\xb2\x37\x7e\xf6\xf0\xdb\x94\x72\xfd\x83\x3d\xcc\xee\xa8\xb6\xcc\xf0\xf4\x44\x05\x3c\x1d\xb1\xe3\x69\xe3\x3b\x99\x87\x3d\x13\x5e\xdf\xc9\x3c\xe4\x94\x78\xda\x38\x0f\xf1\x90\x53\xe2\x69\xe3\x3c\xc4\x43\x4e\x89\xa7\x8d\xf3\x90\x2d\xce\xfa\x11\xeb\xba\xf3\x9a\x39\x9e\xf6\x38\x8d\xac\x11\x4f\xc3\xff\x38\xd7\x34\xf1\xf4\x2c\xc3\xd3\x73\x0c\x4f\x7f\x87\xe1\xe9\x77\x81\xa7\x69\xc2\x3a\xf1\x34\x6d\x88\xcd\x75\x85\xe2\xcf\x58\x57\xa2\xda\x65\xb2\xd7\x12\xbf\xa5\xb8\x0a\x88\xef\xdf\x96\x38\x8c\x6a\xb3\x0c\x87\x19\xe7\x9d\x13\x37\x87\xc3\x8c\x73\x36\x66\x8f\x0a\x87\x45\xbc\xce\xe9\x14\xf2\x1c\x87\x4d\xac\x41\xbe\xea\xf3\x41\x85\xde\x35\x9d\x0f\x7a\xe8\xf5\x3c\x1f\xf4\x90\xf3\x3c\x1f\x54\xcc\x87\xa5\xaa\xe6\xc3\x1a\x71\xd8\x2c\xe2\xc2\xb8\x1f\x31\x63\x7c\x0f\x98\xc3\xfa\x8d\x38\x69\xc1\x7a\x5f\x00\xfe\xc2\x77\x8f\xf4\x14\xe2\xa6\x81\xd6\xfb\x34\xbe\x23\x0f\x1e\xb8\x4c\xf3\xe5\x1d\xe0\xaf\xd1\x1f\x51\x7e\xdc\x45\xf6\xec\x7b\x8e\xca\x60\xed\x1f\x92\x9d\x38\x1f\xdf\xf7\x05\x83\xdf\x2f\xcb\x17\x71\x6a\x16\x14\xf4\x21\xef\xc5\x00\xe8\x00\x1d\xfc\x47\x6b\x02\xb2\xe4\xf3\x2f\x24\xa7\x9f\xe5\x9c\xee\x69\x3a\xd7\x8a\xfa\xc6\xa8\x7e\xc5\xf3\xb9\x53\x65\xff\x06\x2c\xe7\x55\xfc\x1c\x7e\x2c\xdf\x55\x5e\x9f\xfc\x72\xbe\x74\xc3\x4f\x84\x07\x8c\xf6\x13\xe1\xbd\x6c\x1d\xc3\x77\x1d\x03\x5f\x4d\x53\x3e\x49\xe0\x3b\x4f\x7c\xfa\x3d\xd0\x03\xc8\x4b\xe4\xdf\xab\x38\xaf\x1c\xc4\xf7\x9e\xf4\xf4\x65\xc6\x3f\x02\xf9\x1f\x41\x7e\x08\xf3\x38\x83\x32\xe7\xc0\x8d\xbd\x8a\x73\xd0\x74\xe1\x08\xda\xc7\xfe\x77\x86\xec\x88\xcf\x1e\x44\xde\x23\x40\x34\x4f\xc7\xb2\xe5\xef\x2c\xa7\x51\x06\x03\x9f\xa7\xf7\x0f\x09\x36\x6e\x8d\xc4\x0f\x72\xfe\x5d\xc4\xb7\xe0\xb4\x88\x07\x4e\x33\xf6\x21\x7c\xbf\x5b\xc9\xee\x04\xbe\x77\x25\x3a\x5f\x67\x7e\x9b\x64\xb8\xf3\x0d\xe0\x4e\xc3\x7f\x6f\x3a\xf0\xc6\x84\x27\xde\xbf\xf6\x19\xaf\x3f\xe5\x59\xff\x7f\x3e\xfb\xd5\xc6\xdb\x14\x8b\xb7\x59\x16\x6f\x73\x2e\xf1\xf6\x1d\x97\x78\x7b\xf7\x96\xe2\xed\x56\xc7\x3d\x20\xfe\xe9\x57\xea\x3f\x7d\xda\xc8\x97\xd8\x77\x74\xe2\x9e\x9b\xf1\x9d\x76\x1a\xf9\x72\x06\xfb\x23\x7c\x7f\x35\xf6\x49\xce\xf9\x3c\xee\x32\x9f\x5f\x66\xf3\xf9\xeb\xf0\xeb\x59\x94\xaf\xfe\xd2\xfc\xdb\x2f\x7d\xf9\x6f\xd4\xbe\x63\xff\x83\xfb\x1b\x69\xf5\x7d\x4c\xbe\x0f\x0a\xd6\xc7\x04\x9f\x0f\xaa\xfb\x0a\xce\x7d\x10\xad\x03\xe6\x3e\xe8\x2b\xb4\x0f\x32\xc6\x11\xed\x0f\xe6\x71\xdf\x24\x7f\x0d\x78\xa5\xf2\xfa\x7b\x6d\x0d\xeb\x76\x2d\xbe\xc7\x57\xba\xf7\x35\x96\x07\x3e\xcc\xe3\xfe\x46\x15\xfb\x30\x35\xee\x30\xd6\xf3\x65\xd7\xf5\x5c\x85\xdf\xc7\xf2\x13\xf0\x43\x6f\xd5\xf9\x4a\xb3\xdc\x33\x0d\x0a\xca\xcf\xc6\x3e\xc9\xe0\x47\xb5\xbb\xe8\x9e\x46\x15\xb8\xb4\x57\xe1\x97\x12\xad\xc9\xfc\x8c\x7b\x19\xf0\xe7\x7d\x6d\x9a\x4d\x4f\x7a\x96\xc6\x2f\x3d\x43\xe3\xf9\x22\xe4\x54\x76\xaf\xbb\x85\xfe\x56\xfe\x4e\x83\x76\xc3\xb7\xda\x6e\xcf\x4d\xef\x33\xa5\x1f\xeb\xf0\x3d\xcd\x81\xf3\x7b\x81\xf3\x7f\x56\xf4\xd9\xf2\x1e\xee\xc5\xe2\x7c\x8d\xe3\x1c\x7e\x1e\x9b\xc0\x7e\xdd\xdc\x9f\x7a\xdf\x6b\x89\x03\x0f\x26\x5a\x90\x8f\x70\x1f\x26\x11\xfe\x09\xf4\x76\xb3\x3c\x3a\x8b\x79\x50\x79\x1f\x7b\xd1\x63\x1f\xab\xc2\xc3\xef\x59\xef\x87\x31\x3c\x6c\xde\x3f\x03\x2e\x2d\x5c\x63\xeb\xfb\x65\x47\xde\x9c\x53\xe4\xcd\x38\xce\xe1\x83\xda\x8f\x81\x07\xde\x53\xfa\x35\xa8\xfd\x10\x7a\x48\x9f\x79\x7e\x68\xe8\xbf\xcc\xf8\xff\x48\xfa\x2d\x79\x79\xc2\x23\x2f\x1b\xdf\x8d\x02\xe2\xcf\x6e\x73\xfc\xbc\x2b\xf7\x89\xb7\x2f\x7e\x6e\xee\xbe\x94\x19\x57\x57\x11\x4f\xc6\xfe\x93\xee\x0f\xea\x73\x88\xab\xf2\xf7\x71\x23\xae\x2a\xef\x0b\xd5\xe7\x5c\xd5\xef\x0b\xaf\xac\x51\x9e\xef\x87\x2f\x7a\xc8\x1b\x71\x5d\x5b\xef\xdc\x07\x5f\xf2\xd8\xef\x99\xf1\xbd\xe8\x12\xdf\x1f\x56\x15\xdf\xe6\x39\x2d\x8d\x4b\xd4\x87\xef\xff\xe8\x4f\xe5\x7d\x55\x3d\xe2\x93\xc6\x5b\x9f\xa5\xf1\x8e\xd6\xd0\xc6\xc0\x6d\x7f\x40\xd9\xc0\xb2\xdf\x0a\x3c\x52\xee\x1f\xc5\x21\x70\xd6\x3b\x88\xb7\xd1\xfe\x72\x9c\xf9\x64\x5c\x00\xcf\xe7\x07\x18\x4e\xb3\x9f\x43\x9b\xb8\x69\x04\x34\xce\x91\x5b\xce\x00\x6f\x91\xff\x4c\xfc\xf5\x21\xe8\x97\x59\x5c\x56\xc6\x5d\x29\x65\xfe\x18\x41\xfe\x78\x12\xf9\xe3\x8c\x4b\xfe\x78\x14\xe3\x38\xce\xf2\x84\x71\x5f\xfb\x65\xc6\xdf\xe1\xc8\x1f\xbd\x5e\xb8\x0e\x7f\x37\xe2\xc4\xcd\x74\x4f\x64\x1f\xee\xdb\x04\x35\xba\x47\xb2\xef\x69\xd0\xfe\x66\xe8\x5f\x64\xfa\x1f\x02\xff\x43\xc6\xbf\xff\xa6\xf2\x9a\x89\x37\x7f\x6e\xc3\x9b\x26\xae\x24\xbc\x69\xdc\x97\x4e\x34\x00\xe7\x95\xef\xc7\x2e\x3b\xd6\xe9\x6b\x9e\xeb\x74\x83\xa4\xe3\xe5\x75\x9a\x80\x24\xc7\xa7\xd6\x7b\x2f\xcb\xb7\xe5\x5e\xaf\xf3\x1e\x6f\x75\x38\xf8\x41\xa5\x7d\x6e\xfb\x49\x8e\xbf\xcd\x79\x45\xf3\xdc\xb9\x2f\xba\x54\x9e\x47\x35\x96\x75\xdd\x9c\x57\xf6\xf3\x42\xe3\x7c\x90\xdf\x8f\x35\xef\x93\xe3\x1e\x28\x70\xbb\x71\xde\xad\xca\xab\x6e\xfb\x94\x8b\x37\xb1\x0f\xe4\xfe\xaf\xb1\xfc\x27\x6a\xc4\x9d\x07\xcf\x1d\xbf\xa8\x1f\x0d\xf3\x46\x4e\x9b\x3b\x7e\x29\x3f\x77\xfc\xa2\x7e\x34\xe4\x2d\x89\x67\xee\xf8\xa5\xfc\x94\xfc\xd2\x7c\xc7\x2f\x8e\xc7\x98\x47\x77\xf2\xae\xfd\xd1\xee\xac\x47\xca\xc7\x88\x97\xb3\xf8\x7d\xe7\xa1\xe7\xb9\xe3\xb9\xa6\x78\x22\x91\xcc\xe4\x9a\xf6\x3c\x1b\x7b\xfe\xf7\x9f\xe9\x7b\xfc\xa1\x81\x87\x9b\x86\x32\xc7\x87\x73\x4d\xa5\x5f\x03\xf8\x25\x9e\xeb\xdd\x2b\x84\xf8\x0c\x7f\x5b\xb9\x11\xf2\xda\xe8\x7e\x51\x77\x66\xbd\x76\xaf\x35\x7f\x0b\x21\x46\x50\x96\x78\x0f\xe3\xf7\x62\x9d\xc9\xeb\x52\xd4\xc3\x9f\x20\x89\x94\xa5\x1e\xae\x81\x88\x4b\x0d\x26\xef\x8c\x61\xfc\x06\x93\x37\x01\x56\xc6\x52\x0f\x10\x58\x2c\x58\x74\x5c\xc4\xef\xf1\x0a\x41\xf0\x25\xf9\xfe\xb7\x1c\xfc\xdd\x72\x4e\x05\x44\xf3\x3a\x3b\xff\xbb\x7e\xe2\xd7\xd5\xda\xf9\x5d\xe0\x5f\x62\xfc\x2d\xe0\x77\xf9\xed\xfc\xb7\x7d\xc4\x4f\xb1\xfa\x67\xc1\x6f\xaa\xb3\xf3\xe3\xe0\x77\x33\x7b\x7e\x17\x76\x8e\xb0\xfa\xbf\x80\xde\xb3\x8c\xff\xcf\xe0\x4f\x32\x7b\x32\x68\xff\x22\xab\xbf\x07\xfc\x05\xc6\x8f\x48\xf9\x7a\xd1\xcd\xec\xff\x96\x8f\xf8\xe3\x8c\x3f\x0e\x7e\x88\xb5\xf3\x07\xe0\x77\x33\x7e\x7f\x0d\xf1\xfb\x19\x3f\x8d\xfa\x73\xcc\xfe\xbf\x96\xf4\x7a\xb1\xcc\xf8\x9b\xc0\x1f\x09\xd8\xf9\x23\x3e\xe2\xf7\xb3\xf8\x78\x01\xfc\x39\x56\xbf\x19\xed\x2c\x32\x7b\x34\xf0\xeb\xea\xed\xfc\x57\xc1\x9f\xb3\xb3\xc5\x27\x68\x9f\x27\xa7\xbf\x00\x3f\xc5\xf8\x6d\xd2\x0f\x1b\xe0\x0f\xf3\xa9\x07\x3f\xd4\x60\xe7\x7f\xe8\x27\x7e\x8a\xf1\x5b\xc1\x5f\x64\xed\xd7\x83\x9f\x61\xed\x6f\x45\xfb\x11\xd6\xce\x0c\xea\x2f\xb0\xf1\xfd\x85\x8f\xf8\xab\x3c\xfe\xd1\xce\x02\xf3\xcf\xcf\xd0\xce\xdc\x7a\x3b\xff\x32\xf8\x93\x8c\xff\x9a\x8b\xfd\xaf\x18\xf6\x33\xfe\x56\xf0\xf9\xfc\xff\x6f\xd8\xb9\xcc\xf8\x9a\xb4\xb3\xd1\x31\x4f\xdf\xf7\x13\xbf\x97\xf1\x9f\xf7\x11\x7f\x82\xc5\xc9\xff\xa2\x3e\x9f\xbf\x7f\x09\x7e\x8a\xf1\x77\xa3\x1d\xee\x9f\x5e\xd8\x73\x85\xf1\x37\x81\xbf\xc4\xf8\x13\x68\x27\xc5\xf6\xfd\xed\xe0\x2f\x33\xfe\x0f\xa4\x3d\x9b\xca\x79\xd2\x78\xfe\x4b\xf2\x37\x3b\xfc\x29\xe4\x39\x89\xdf\xc9\x94\xfc\x80\x0b\xbf\xd6\x85\x5f\xef\xc2\x5f\xef\xc2\xdf\xe0\xc2\x6f\x74\xe1\x6f\x72\xe1\x6f\x76\xf0\xd6\xcb\xf3\x94\xdf\x76\xf0\xfb\xe4\xb9\x70\x93\x83\x1f\x97\xe7\xd5\xf7\x38\xf8\x3f\x96\xed\x84\x1c\xfc\x27\x64\x3b\xbf\xe9\xe0\x3f\x23\xd7\xa9\x7b\x1d\xfc\xac\xac\xef\xb4\xff\x23\xc9\x77\xfa\xe1\xfb\x92\xef\xf4\xdb\xbf\xcb\xf6\x9d\xe3\x75\x54\xf2\x9d\xe3\xf2\x92\x6c\xc7\xe9\x9f\x5a\x59\xdf\xe9\xe7\x1f\x48\xbe\x73\xdc\xbf\x2c\xf9\xce\xf1\xfd\x92\xf4\x5b\xd0\xc1\xff\xae\xf4\xdb\xdd\x0e\x7e\xb3\xb4\xe7\x3e\x07\xff\xb4\x6c\x67\x8b\x83\xff\x9f\xb2\x9d\xb0\x83\xdf\x8d\xb2\x14\xce\x4f\x97\xd6\x2d\x46\x67\x18\x7d\xde\x42\xef\x2f\xf5\x65\x9d\x49\x97\xfa\x36\x57\x67\x7f\x6f\x6d\xff\x8b\xac\xfd\x2f\xb2\xf6\x4b\xf4\x14\x6b\x7f\xc2\x6f\xa7\x45\x9d\x5d\x5f\xc4\x42\xf7\x31\x7d\xa5\xfa\x57\x18\xdd\xac\xd9\xe9\x55\xbf\xbd\xbd\xfe\x80\x49\xf7\x96\xf0\x46\xc0\xfe\x7e\x99\xe9\x0b\xd5\xdb\xe9\xc9\x06\x93\x7e\x01\x6d\x58\xf5\x75\x31\xfd\x73\x8c\x1e\xaf\xb7\xeb\x0f\x35\xd8\xf5\x77\x37\xd8\xf5\x8d\xac\xb7\xd7\x5f\xbc\xcb\xae\xbf\x9f\xe9\x5f\xe0\xfa\x7c\x26\x5d\xfa\xef\xa2\x8f\x8d\x2f\xf3\x7f\x86\xf9\xe7\x3c\xb3\xb7\x9b\xf5\xdf\x3a\x1e\xa5\x39\x3d\x69\xa1\x9f\xc5\x9e\xde\x4a\x1f\x64\x74\xb7\x66\x6f\xdf\xc8\x14\x0d\x88\xa5\x56\x0b\xbd\x1b\x31\x64\xd0\x9d\x96\xfe\x97\xe8\x47\x85\x10\x43\x16\xfa\xb1\x52\xff\x2d\xf4\xe3\x42\xe0\xaf\xe3\x89\x7e\xa2\x34\x3e\x16\xfa\x49\x21\xc4\x7b\x16\x5a\x62\xf7\x6d\xb9\xe4\x48\x4e\x0c\x0e\x27\x73\x99\xe1\xe3\x5f\x4d\xc6\x62\xe9\x63\xc9\x5c\x2c\x91\x1d\x8c\xd1\x76\x42\x6c\x1b\x4e\x0e\x95\x5f\x3f\xc2\xdf\x1e\x8d\x67\xb2\x8f\x1c\x1a\x48\x1f\xcb\xe6\xe2\x43\x43\xb1\xe1\xa4\xad\xad\x5c\x22\x13\x3b\xb5\x33\x96\x38\x7e\xec\x58\x32\x91\x13\x83\x6a\xb6\x5d\x83\xea\xa5\xf2\x0d\xd7\xd3\xa1\xd6\xd3\xe1\xa5\xa7\xc3\x55\x8f\xf9\x46\x76\x31\x97\xc8\xe4\x86\xe3\x89\xe4\x70\x2c\x9b\x8b\xe7\x4e\x66\x79\xc7\x33\xe9\x81\xac\xd8\xb6\x77\x5b\x36\x37\x2c\x62\xa7\x92\xc3\xd9\xf4\xf1\x63\x56\x0b\x0f\x0d\xc4\x50\xb5\x6c\x9d\x85\x65\xb7\x8c\xbf\x70\x70\xad\xfd\xcb\x26\x73\xd2\xa4\x24\xef\x81\xf9\xc2\x5a\x3d\x31\x74\x3c\xeb\xa8\x4a\xcc\xd8\x50\x3a\x91\x3c\x86\xb7\xdb\x92\xa9\xd8\xa1\xe1\xf8\xd1\xa4\x28\x75\x29\x17\xff\xaa\xd8\x96\xd5\x8f\x96\xca\xbd\x4f\x3d\xb5\x33\xb6\x9b\x8a\x76\x59\xb6\xa1\xdc\x19\x6b\x93\x65\x07\xca\xed\x28\xdb\xca\x74\x04\xd5\x23\xa8\x16\x41\x35\xf0\xa9\xdc\x19\xdb\x45\xaf\x77\xe1\xed\x2e\xbc\xdd\x55\x1e\x8a\x58\xf2\x54\xf2\x58\x2e\x96\xce\x9c\xda\x49\x3c\x8c\x55\xf6\x78\x62\xd0\xc2\xcd\x9d\xcc\x0c\x25\x33\xe9\x01\x62\xed\x7d\xea\xa9\x1d\xb1\x9d\xd4\x34\xa8\x76\x90\xed\xa0\xdb\x40\x53\xd9\x86\x72\x47\x6c\x07\xaa\xed\x40\xb5\x1d\xb0\x6c\x07\xaa\xed\x50\x58\xd6\xa1\xb4\xac\xc3\x69\x59\x07\x35\xda\x41\x3a\x40\x6d\xef\x80\x87\x41\xb7\xe3\x75\x5b\x07\x54\xca\x72\x57\x6c\x3b\xaa\x6f\x87\x67\xb7\x63\x20\xc0\xa7\x32\x12\x6b\x47\xb5\x76\xb4\xda\x8e\xea\xa0\xdb\xf0\xbe\x0d\xf4\xf6\x36\xd4\x47\xd9\xde\x86\xfa\xa0\xdb\x40\xb7\xb5\x89\x6d\xc3\xc7\x07\xe2\xb9\x78\x29\x4a\xda\xb6\xb5\x41\xbc\x95\xfc\xd3\x6e\x64\xb7\x9b\x7f\xee\xf7\x09\x25\x2a\x5d\xdd\x4b\xe5\x41\xb6\xaf\xe1\xb0\xd6\x58\x17\xd8\xb6\xba\x7c\x6e\xc1\x1f\x7e\xd4\xb6\x41\xf3\x96\x1f\x61\x7c\x06\xff\x45\xbd\x26\x94\x68\xf8\xe2\xf3\x54\x1a\xe7\x2f\x0f\xe0\x3b\xa5\x21\x6f\xf0\x57\x5c\xec\x37\x70\x3d\xdf\xaf\x72\xfd\x3f\x17\x6a\xfd\x0b\xd0\xdf\x6b\xd1\x1f\x50\xe8\x7f\xcb\x45\xff\x04\x1a\xad\xd4\xff\x82\x8b\xfe\x50\x8f\x5d\xcf\x03\xc0\xa2\x5c\xff\x59\x17\xfd\x75\x58\xa7\xf9\x7e\x8b\xeb\xff\xba\x8b\xfe\x66\xe8\x3f\x68\xd1\x5f\xaf\xd0\xff\xaf\x2e\xe3\x9f\xa1\xeb\xdc\xe5\xf3\x2c\x37\xfd\xff\xe2\x32\xfe\x67\xa1\x7f\xc2\xa2\x7f\xbd\x42\xff\xa7\x2e\xfa\x5b\x0d\xe8\xbc\xc1\x5b\xff\x75\x17\xfd\xa1\xdf\xa3\x72\xdc\xa2\x7f\x83\x42\xff\xef\xb8\xf8\xbf\xb5\x85\xca\x4c\x85\xfe\xdf\xef\xe2\xff\x49\x85\xfe\x46\x85\xfe\x19\x97\xfe\x4f\xe1\xef\x55\x17\x84\xb7\xfe\xbf\x72\xeb\x3f\xf2\x47\x93\x45\xff\x26\x85\xfe\xbc\x8b\xfe\x14\xee\xf1\xf0\xf3\x07\xae\xff\x1b\x2e\xfa\x23\xd0\xdf\x6a\xd1\xbf\x59\xa1\xbf\xdb\x4f\xfa\x79\x0e\x5c\xc6\xbf\xc3\xc4\xef\x3b\xf0\xfc\xf5\x37\x3e\xb5\xbc\xb1\x91\xa9\x24\xff\x47\x2e\xfa\x43\xbb\xab\x93\xbf\xec\xa2\xbf\xb5\xb3\x3a\xf9\x61\x17\xfd\xdd\x8f\x56\x27\xbf\xe8\xa2\xff\xe0\x63\xd5\xc9\x77\x09\xb5\xfc\xc8\xe3\xd5\xc9\xef\xd6\xd4\xf2\x13\x4f\x54\x27\xff\x80\x8b\xfc\xdc\x93\xd5\xc9\xef\xa9\x21\xf9\x76\xc6\xbf\x84\x05\xac\x85\xf1\x35\x56\xde\xe5\xe2\xbf\x8f\xbb\xd4\xfa\xf8\xfa\x1b\x75\xb1\x7f\xc5\x45\x9e\xd3\x21\xe8\x67\xc7\x5b\x62\x15\xf2\x95\xe6\xff\x6f\xf8\xd4\xf3\x2f\x85\xf9\x67\xcd\xbf\x5b\x15\xf3\xaf\xc5\xe7\xd4\x5d\x7a\x2e\xe1\xdf\x69\x32\xce\x43\x4b\x36\xee\xb1\xc8\x1b\x27\x45\xff\x1f\x00\x00\xff\xff\x9f\x8a\xcc\x8d\x58\x5a\x00\x00")

func tcptracerEbpfOBytes() ([]byte, error) {
	return bindataRead(
		_tcptracerEbpfO,
		"tcptracer-ebpf.o",
	)
}

func tcptracerEbpfO() (*asset, error) {
	bytes, err := tcptracerEbpfOBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tcptracer-ebpf.o", size: 23128, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"tcptracer-ebpf.o": tcptracerEbpfO,
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
	"tcptracer-ebpf.o": &bintree{tcptracerEbpfO, map[string]*bintree{}},
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

