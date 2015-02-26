package justext

import (
		"bytes"
		"compress/gzip"
		"io"
		"reflect"
		"unsafe"
)

var _DetailedTemplate = ""+
"\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe4\x59\x59\x73\xe2\x3a"+
"\xf6\x7f\x4e\x3e\x85\xff\xfc\x6f\x4d\x27\x45\x9a\x1d\x02\x09\x64"+
"\xc6\x18\x13\x76\x02\x06\x02\x99\x9a\xea\x12\xb6\x6c\x14\xaf\xb1"+
"\x65\x8c\xc9\xf0\xdd\x47\x32\x66\x0b\x24\x4d\xcf\x43\xdf\x5b\x35"+
"\xbc\xd8\x3e\x3a\xfb\xf9\x49\x3a\x12\xc5\x19\xd6\xb5\x87\xcb\xa2"+
"\x83\x7d\x0d\x3e\x5c\x5e\xfc\x03\xe9\x96\x69\x63\xc6\xb5\xb5\xab"+
"\x19\xc6\xd6\x5d\x3c\x2e\x9b\x06\x76\x62\x8a\x69\x2a\x1a\x04\x16"+
"\x72\x62\xa2\xa9\xc7\x45\xc7\xf9\xbb\x0c\x74\xa4\xf9\xa5\xae\x05"+
"\x8d\xa8\x00\x0c\xe7\x2e\x93\x48\xdc\xa4\x13\x89\xbf\x39\xee\xd4"+
"\x81\xb8\xa4\x01\x8c\x8c\x1b\xd1\xb7\x91\xa6\x21\xf1\x3b\x5c\xe0"+
"\xed\xc7\x8d\x62\x43\xa8\x06\xa4\x39\x82\xd8\x00\x3a\x74\xe0\x4d"+
"\x20\x40\x89\xd7\xf7\x97\x17\x97\x17\x53\x53\xf2\x2f\x2f\xde\x2f"+
"\x2f\x2e\xa8\x0b\xdf\xd7\xe6\xee\x98\x6f\xd4\x20\x43\x0d\x7e\xbb"+
"\x61\x1c\xf2\xf8\xee\x40\x1b\xc9\xf7\x1b\x3e\x0f\x22\x65\x86\xef"+
"\x18\xe2\xc9\x3d\x43\x88\x3a\xb0\x15\xa2\x56\x83\x32\x21\x02\x17"+
"\x9b\xf7\x3b\xa2\xbd\x66\xdd\x50\x3d\x24\xe1\xd9\x1d\x93\x4f\x24"+
"\xac\x05\xf9\x5e\x5d\x5e\x5e\x38\x16\x30\x62\x0e\x36\x2d\xcf\xb4"+
"\x25\x86\x3a\x83\x89\x83\xdf\x25\x28\x9a\x36\x71\xd7\x34\xee\x18"+
"\xd7\x90\xa0\xad\x21\x03\x86\x22\x16\x4d\x96\x74\xc3\x58\xb1\x19"+
"\x04\x12\x32\x14\xfa\x3a\x05\x81\x34\x43\x7e\xa1\x95\x6c\x76\x6d"+
"\x85\x92\x02\x9d\x40\x43\x0a\x51\xf7\xea\x3a\x18\xc9\x7e\x38\x12"+
"\x3a\x4a\x1c\xb8\x63\x52\x5b\xfe\x90\x3a\x35\x31\x36\xf5\xfd\x01"+
"\x0b\x48\xd4\x62\x18\x6d\xfa\x88\x1e\x06\x9c\xde\xc6\x77\xe4\xec"+
"\xc6\xcb\x29\x10\x55\xc5\x36\x49\x70\xdf\x45\x53\x33\xed\x3b\xe6"+
"\xff\x45\x59\xdc\x4a\xed\x71\x1f\xa6\x3d\x13\xa6\x3d\x20\x3a\x68"+
"\x09\xef\x98\x64\x2c\x0d\xf5\xad\xe4\x5e\x26\x4e\xd8\x90\x25\x29"+
"\xe4\x94\xd0\x3c\x66\x01\x1b\x28\x36\xb0\x66\x3f\x24\x88\x01\xd2"+
"\x9c\x0f\x39\x4c\x65\x76\x39\x94\x35\x13\x10\xfb\x41\x84\x21\x49"+
"\x42\x8e\xa5\x01\x82\x19\xc3\xdc\x56\xe7\xb4\x5a\x0c\xa6\x1a\xdc"+
"\xba\x45\x4a\x0d\x6d\xea\x92\x06\x2c\x87\x04\xb0\x79\x3b\x43\x05"+
"\xde\x06\x77\x58\x8a\xe3\x12\x85\xa5\xd8\x0d\xac\xcd\x92\x74\x59"+
"\x0b\xc6\x31\x35\x24\x91\x6c\xc8\xf2\x57\x36\x83\xda\xad\x0d\xdf"+
"\x7c\xc6\xb2\x29\xd4\x89\x08\x0f\x4c\x15\xe4\xc2\x97\xa6\x68\xd9"+
"\x7e\xa6\x43\x06\xe0\x4c\x77\x19\x6c\x9f\xe9\x31\xb6\xbf\x40\xcb"+
"\xf9\x4e\x7f\xa9\x66\xe7\x77\x0c\xfa\xd0\x24\xeb\x4b\xb8\xee\xac"+
"\x21\x45\x2b\x78\xff\x71\x25\xf9\x9e\x24\x40\xdf\x11\x83\x09\xba"+
"\x9e\x56\x17\x5b\xd8\x4d\x35\x53\x54\xf7\x16\x96\x74\x72\xcd\x30"+
"\x0b\xe7\x4a\x2a\xb3\xfe\xde\x73\xc9\x86\x16\xa4\x36\x0d\x33\x7c"+
"\xfd\x30\x8e\x74\xa0\x10\x44\xd2\xf5\x59\x02\x18\xdc\x05\xdf\x71"+
"\xcb\x50\xee\xa7\xc0\x81\xb9\xcc\x0d\x1a\x95\xbb\x7d\x2f\xd1\x7c"+
"\x54\x4c\x96\xfc\x3a\xc2\x70\xc6\x0f\x15\xf2\x56\xce\xd3\xef\x3e"+
"\xc7\x4e\xe8\x53\xb3\x6a\x52\x83\x52\x1b\x09\x8d\xef\x8d\xfa\x99"+
"\xd4\x5b\x6a\x34\xee\xa5\x2b\x6d\xd6\xe7\x17\x3e\xdb\xe3\xde\xf8"+
"\x1e\xab\xf3\xc1\x93\x43\x7c\xaf\x8c\xc2\x77\x87\xef\xb5\x5f\x0b"+
"\x62\xa3\x8d\x9b\x53\x0c\x52\x42\x27\x5a\x49\xe4\x07\xc6\xb8\x76"+
"\x9b\x03\xcd\x54\x35\x31\x9d\x98\xcf\x03\x55\x74\xb3\xf5\x45\x4b"+
"\x5d\xf0\x6d\xae\x96\xb5\xa6\x3d\x76\x24\x69\x28\xa3\x2d\x13\xc3"+
"\x7a\xaf\xde\x86\xd9\xa7\xbe\xe8\xf3\x6d\x1f\xa4\x9f\xd2\xec\x5b"+
"\x53\x10\xdc\xe9\xe8\x09\x18\x8b\xaa\x66\xc1\x16\x18\xa3\x7e\x79"+
"\xf8\x72\xcb\xab\x13\x7b\xde\x19\xbd\xb1\xd5\x41\x6e\x22\x8d\x27"+
"\xed\xc9\x68\x24\x24\xeb\x8f\x2f\xc9\x49\x33\xca\xc7\x1f\xcb\xe2"+
"\x04\xd5\x52\xb5\xea\xf8\x35\x5b\xad\x8d\xbb\x99\x6a\x66\x91\x6b"+
"\x45\x93\xaf\xb7\xac\x50\x43\x9e\x33\x28\x3c\xb5\x66\xfd\xf9\x4b"+
"\x27\x6a\x44\x0b\x2e\x1a\x62\x81\xef\xcb\x4e\x74\xdc\x99\xb0\x9c"+
"\x97\xcc\x0f\xa7\xdd\x64\xe7\xcd\x86\x6a\xf2\x79\x6e\xcd\xa3\x65"+
"\xf1\x55\x4d\x34\x52\x95\xd6\x44\x28\x24\x6a\xe5\x6a\x05\x60\x2e"+
"\xaf\xf8\x8b\xac\x06\xe2\x51\x3b\x5d\x4b\xb3\x03\xbe\xda\x7d\x92"+
"\x2b\xad\xaa\xfb\xda\xcd\x65\xdc\xe6\xa0\xba\x30\x3b\xe9\x27\xcb"+
"\xc8\x38\x6f\x78\xc9\xdd\xbe\xe9\xa3\x7c\x42\x97\x04\x77\xe0\x4d"+
"\xbb\x59\x3d\xe5\x8e\x86\x80\xeb\x75\xc1\x70\x2c\x88\x42\x83\x2b"+
"\x14\x16\x8e\x5d\x90\x0a\xf5\x42\x2a\xee\x2f\x9b\x8b\x81\xee\xb9"+
"\x43\x0f\x3e\xaa\xb5\x78\x5a\x2d\x64\x66\x15\x29\xe5\x2f\xec\x11"+
"\x6c\x69\x62\xa7\xc7\x06\xa5\xe9\x0f\xb3\xbc\xad\x36\x14\x45\x29"+
"\x95\xae\xf7\x10\x2a\x6a\xa6\x03\xa5\xf3\x31\xfa\x57\x06\xe4\x24"+
"\x04\x24\xf7\xe6\x83\x32\xa5\x4e\xf9\x10\x90\x76\xea\xb9\x22\xa6"+
"\xf9\x3a\xeb\x3f\x2e\x3d\x5e\x68\xe8\xbc\xa0\xd6\x39\x5e\x18\xb2"+
"\x1c\x6f\xb1\x1e\x37\x44\x93\x19\xaf\x4f\xca\xdc\xc0\x01\x8b\x2e"+
"\x28\x74\x32\x6c\x34\x8e\x53\x8d\xc7\xbc\x5b\x4e\xfa\xb2\x66\xbd"+
"\xf8\x46\x3a\x2d\xc7\x93\xc3\x44\xed\x51\x9d\x73\x76\xaa\x30\xcd"+
"\xb3\xe3\x72\xbf\x10\x2d\xeb\x70\xc2\xa2\x42\x2d\x9a\xcb\x03\x9c"+
"\xeb\xb3\x48\xf7\xf3\xcd\xc6\x30\xdb\xcf\x9a\x23\xb9\x65\x56\xfb"+
"\x39\x76\x32\x07\xe5\x7c\xc3\x7a\x75\xf9\xa1\x1d\x9d\x22\x24\xa4"+
"\x9b\x29\xce\x8c\x3e\xce\xbd\x6a\x45\x89\xf2\x26\xa8\x44\xf9\xc4"+
"\xb2\xae\xbe\x2e\x15\xd9\x06\x99\xda\xeb\x74\xb1\xe4\xf8\xe9\x22"+
"\xa5\xd4\xb3\xea\x20\xd1\x2b\xa8\xcf\x0d\xdb\xee\x8d\x9e\x0c\x61"+
"\xf6\xfa\xe2\xc5\xdd\x4e\x4e\xe8\xcc\xf3\x8e\xe1\xd7\x6e\xcb\x09"+
"\xa9\x2b\xd8\x04\x65\xec\x52\xe8\x69\xfa\x14\x5a\x4b\x47\xf4\xa2"+
"\x65\xde\x29\xf4\x58\xac\x76\x9f\x47\x39\xb8\x18\xc1\xf2\x4b\x56"+
"\xab\x22\xfc\xc8\xdb\xb9\x96\xd9\x83\x75\xa1\x3e\x00\xa3\x51\xc3"+
"\x18\x71\xee\x4b\xed\x16\x4b\x2d\x6d\xda\xec\xea\x6d\x37\x3a\x98"+
"\xf5\x75\x1b\x8f\x5b\x5a\xd9\xd1\x0c\x33\xda\x28\x2c\x5e\x26\x96"+
"\xe8\x77\xd9\xbe\xef\x67\xd8\x9a\xb1\x48\x35\xe7\xb3\xee\x5b\xbd"+
"\x62\xa9\xac\x17\xf7\xf8\x0a\x28\x24\x97\x5c\x62\x32\x6e\xb6\x9a"+
"\xaa\x69\xf0\x60\xd4\x66\x5f\x8c\x4e\x54\xc9\x65\x35\x38\xcb\x3e"+
"\xd5\x13\x4e\x35\xa9\x3b\x8f\xd6\x78\x1e\x97\x85\x6e\xf6\x6d\x7a"+
"\x2b\x4d\x92\xfd\xf4\xab\xe5\x56\xd5\x4c\x6b\x68\xd8\x05\x2b\x21"+
"\xb6\xf2\x05\x2e\xeb\x95\x73\x95\x44\x66\x6e\xd4\xe1\x30\xd9\xd2"+
"\x59\xcf\x91\xd5\xb6\xe4\x55\xab\xfd\x84\x37\x1f\xa6\x33\x12\x1a"+
"\xa6\x9b\xf1\xd9\xf3\xb8\xde\x02\xd1\x38\x07\x86\xcb\x0c\x8e\xa3"+
"\x46\x5c\xa9\xbc\xc8\x5a\x5a\x2f\xe4\xc4\x71\xdb\xfb\x0c\xec\xc5"+
"\x78\xd8\x96\x16\x1d\xd1\x46\x16\x26\xfd\xa9\xec\x1a\x22\xed\xba"+
"\x18\x67\x66\x7a\x3f\x80\x34\x07\x86\x08\xa5\xab\xeb\xcd\xe2\xbe"+
"\xa1\xfc\x30\x2d\xca\xe6\x30\x25\x46\x32\x45\x57\x87\x06\x8e\x29"+
"\x10\xf3\x1a\xa4\xaf\x65\xbf\x2e\x5d\x45\x3e\xf2\x46\xae\xef\x4f"+
"\x2b\x89\x05\x6e\xc4\xc2\x59\x44\x54\x46\x82\x79\x14\x09\xd9\x0f"+
"\x5c\xf9\x41\x1a\x41\xf5\x2b\xab\xc7\xdc\x5b\xbb\xc7\x43\xc7\x96"+
"\x69\x27\xb3\x31\x3c\x43\x12\x3c\xdf\xf0\x31\xf7\xd6\xf0\xf1\xd0"+
"\xb1\xe1\x48\xb8\xfe\x6c\x2b\x70\x20\xf4\xdb\x2b\xb0\x9f\x87\xdf"+
"\x5a\x80\x3f\x25\xf9\x9b\x68\x0f\x0a\x80\x4d\x85\x1c\xc9\x7e\x4c"+
"\x4d\xa4\x41\x9b\x70\x62\xf2\xee\x92\x03\x81\xb1\x2e\xc6\xc5\x7a"+
"\xbc\xbc\x1b\xfe\xca\xc1\x23\xe6\xc0\xbf\x0b\x24\x33\x57\x47\x43"+
"\x31\x51\x03\x8e\xd3\x21\x87\xb6\x18\x22\x67\x9f\x45\x57\xbe\x8a"+
"\x6c\x77\xa5\xc8\x35\xf3\x7f\x25\xb2\xf3\xac\x7d\x38\x76\x62\x27"+
"\x4c\x03\x0b\xdb\x2d\x1a\x1b\xd9\x85\x68\x26\xf6\xc2\xb9\x0a\x5c"+
"\x58\x31\x50\x73\xe0\xb9\xda\x42\x27\x02\x7d\x41\x39\x8f\xf5\x9d"+
"\xc2\xf1\x01\xd7\xae\x83\x0f\xfb\xc9\x4f\x40\xec\x94\xfd\x01\x50"+
"\xa8\xf1\xab\x88\xb5\xad\xa7\x6c\xda\xcc\xd5\x1c\xd8\x0c\x2a\x91"+
"\x93\x10\x2a\xee\xb4\xc4\x34\x68\x28\x78\x46\x88\xd1\xe8\xd6\x08"+
"\xfd\x51\xee\x2d\x1b\xb1\xb5\x13\xf9\x27\xfa\xd7\xfd\x8e\x8f\x56"+
"\x63\x3b\x76\xaa\x0a\xa4\xdf\x3d\xc8\x3f\xb3\xf7\xdb\x09\x7e\x39"+
"\x99\xe8\x6f\xb5\x7e\x5d\xed\x01\x74\x1f\x65\x67\x81\xfd\xa3\xc0"+
"\x21\xde\x3f\x8e\x9e\x31\xc1\x7f\xc5\x83\x93\x02\x87\xd3\xfc\xe7"+
"\x1e\x1c\x4d\xb8\x63\x3c\xfd\xcf\x21\xe5\x73\x94\xfc\x39\x35\x3a"+
"\xda\x0e\x7f\x3b\x4e\x4f\xa3\xe4\xe8\x28\xba\xab\xc6\x0d\x63\x49"+
"\x48\x3a\x86\xce\xf6\x22\xe1\x73\x8f\x03\xc1\xfb\x4f\xe4\x3e\xed"+
"\x52\x98\xb5\xc0\x2f\x41\x62\x57\xe5\x8f\x60\xd8\xb5\xfe\x1c\x3d"+
"\x3d\x53\x3b\xf4\xfc\xbc\x29\x02\x5d\xab\x7f\x4d\x98\x9c\xe1\x4f"+
"\xf7\x16\x7f\x89\x1c\xae\x01\xf6\x1b\x52\x28\x49\xff\x7d\x0a\x45"+
"\x59\x8c\x6c\x3a\xe6\xb0\x53\x2e\xd2\xdb\x53\xd2\x30\x17\x25\x34"+
"\x67\x90\x54\x3a\xb1\xc1\x33\x41\x00\xa5\xbd\x5d\x93\x31\x0d\x51"+
"\x43\xa2\xba\xe1\x3e\xd9\x5b\x44\x1e\x8a\x71\xa2\xf4\xe1\xf2\xfd"+
"\xdd\x06\x86\x02\x99\x3f\x82\xf0\x6f\x98\x3f\x2c\xe6\xae\xc4\xc4"+
"\x9e\xb6\x4b\xd2\x6a\xb5\x67\xdf\x92\xde\xdf\xd7\x9c\xab\xd5\xd6"+
"\xf4\x89\xfb\xbd\x77\x92\xe4\x58\x6d\x7d\x05\xb4\x5a\x85\x77\x41"+
"\xef\xef\xd0\x90\x56\x2b\x32\x1a\xe3\xa8\x24\x55\x11\xa4\xa3\x14"+
"\xd9\x9c\x6d\x83\x2b\xbe\xc8\xc3\x05\xd9\xe1\x8b\xc1\x95\xcf\x03"+
"\x6d\x01\x8a\xd8\x0e\x9e\xe4\x45\x7a\xa8\x22\x03\x68\x6b\xd3\x77"+
"\x4c\x31\x4e\x28\x94\xba\xd3\x19\x90\x02\xa9\x78\x28\x76\x20\xce"+
"\x99\x46\x70\x4f\x2b\xdb\x10\x9e\xd4\x22\x9f\xa7\x27\x0c\xee\x50"+
"\x78\x1b\xf1\x4f\x84\x5b\xc1\xc6\xc0\x5c\x21\x83\x11\x67\x24\x7b"+
"\x22\x86\xb6\x73\x7d\xa0\x8b\xec\x1d\x4c\x6c\x40\x3c\xfd\xa9\xb2"+
"\x8e\x4b\xce\x9e\x36\x63\xca\x7b\xba\x18\x0f\x11\xfd\x74\xad\xfb"+
"\x10\x5e\x8b\x90\xa0\xc4\x11\x46\x8e\xa0\xef\xe7\xca\x29\x3f\x23"+
"\x41\xc3\x41\xd8\x3f\xd6\x54\x59\x0f\xfc\x82\x8b\xf4\x02\xfe\x83"+
"\x4b\xcf\x84\x74\x9e\x33\x3b\x35\xf4\x2e\xff\x94\x2e\x21\xbc\xe3"+
"\x3f\x4f\x9f\xb0\xd1\x72\x3a\xc2\x8d\xb2\x33\xa3\xa4\x17\xcc\xf4"+
"\x8f\x86\x52\x24\x15\xa1\xe2\x15\x53\x7f\x02\x78\x76\x42\x8c\x3c"+
"\x43\x6c\xaf\x27\x21\x81\x7b\xd1\xda\xcc\xa6\xf3\xe7\x8e\x69\xe8"+
"\xa6\xeb\x40\x73\x0e\xed\x52\xe4\x93\x0d\x0b\xcf\x90\x73\xc3\x7c"+
"\xdb\x9f\xb7\xdf\xae\xef\x77\xb2\x2e\x2e\x45\x3e\x59\xa7\x3f\x11"+
"\x25\xa1\x0d\x6c\xa4\x0b\x16\x10\xe1\x1a\xa1\xcc\xbf\x99\x36\xb0"+
"\xd5\x4d\xba\x82\xb9\x63\xd1\x95\x25\x70\x99\x2c\x67\xeb\x55\xac"+
"\x18\x0f\xfe\xa6\xfa\x4f\x00\x00\x00\xff\xff\xdb\x07\xd3\x5c\xad"+
"\x1a\x00\x00"

// DetailedTemplate returns the binary data for a given file.
func DetailedTemplate() []byte {
		// This bit of black magic ensures we do not get
		// unneccesary memcpy's and can read directly from
		// the .rodata section.
		var empty [0]byte
		sx := (*reflect.StringHeader)(unsafe.Pointer(&_DetailedTemplate))
		b := empty[:]
		bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
		bx.Data = sx.Data
		bx.Len = len(_DetailedTemplate)
		bx.Cap = bx.Len
		
		gz, err := gzip.NewReader(bytes.NewBuffer(b))

		if err != nil {
			panic("Decompression failed: " + err.Error())
		}

		var buf bytes.Buffer
		io.Copy(&buf, gz)
		gz.Close()

		return buf.Bytes()
}