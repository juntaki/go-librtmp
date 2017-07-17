package rtmp

import "testing"

func TestAll(t *testing.T) {
	//url := "rtmpe://f-radiko.smartstream.ne.jp/TBS/_definst_/simul-stream.stream swfUrl=http://radiko.jp/apps/js/flash/myplayer-release.swf swfVfy=1 conn=S: conn=S: conn=S: conn=S:oalF9-E_E0g98jxPFpajJA live=1 stop=100"
	url := "rtmpe://netradio-r1-flash.nhk.jp/live/NetRadio_R1_flash@63346 swfUrl=http://www3.nhk.or.jp/netradio/files/swf/rtmpe_ver2015.swf swfVfy=1 live=1"

	r, err := Alloc()
	if err != nil {
		t.Fatal(err)
	}

	r.Init()

	err = r.SetupURL(url)
	if err != nil {
		t.Fatal(err)
	}

	err = r.Connect()
	if err != nil {
		t.Fatal(err)
	}

	b := make([]byte, 64*1024)

	size, err := r.Read(b)
	if err != nil {
		t.Fatal(err)
	}
	if size == 0 {
		t.Fatal(size)
	}

	err = r.Close()
	if err != nil {
		t.Fatal(err)
	}

	err = r.Free()
	if err != nil {
		t.Fatal(err)
	}
}
