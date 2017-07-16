package rtmp

import (
	"fmt"
	"testing"
)

func TestAll(*testing.T) {
	url := "rtmpe://netradio-r1-flash.nhk.jp/live/NetRadio_R1_flash@63346 swfUrl=http://www3.nhk.or.jp/netradio/files/swf/rtmpe_ver2015.swf swfVfy=1 live=1"

	r, _ := Init()
	r.SetupURL(url)
	r.Connect()
	defer r.Close()

	b := make([]byte, 1024)
	r.Read(b)

	fmt.Println(b)
}
