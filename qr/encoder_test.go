package qr

import (
	"fmt"
	"github.com/boombuler/barcode"
	"image/png"
	"os"
	"testing"
)

var qrHelloWorldHUni []bool = []bool{true, true, true, true, true, true, true, false, true, false, true, false, true, false, false, false, true, false, true, true, true, true, true, true, true,
	true, false, false, false, false, false, true, false, true, true, false, false, false, true, true, true, false, false, true, false, false, false, false, false, true,
	true, false, true, true, true, false, true, false, true, false, true, false, true, true, false, true, true, false, true, false, true, true, true, false, true,
	true, false, true, true, true, false, true, false, false, false, false, true, true, false, true, true, false, false, true, false, true, true, true, false, true,
	true, false, true, true, true, false, true, false, false, true, false, false, false, true, true, false, true, false, true, false, true, true, true, false, true,
	true, false, false, false, false, false, true, false, true, false, false, true, false, false, true, true, true, false, true, false, false, false, false, false, true,
	true, true, true, true, true, true, true, false, true, false, true, false, true, false, true, false, true, false, true, true, true, true, true, true, true,
	false, false, false, false, false, false, false, false, true, true, false, false, true, false, false, true, false, false, false, false, false, false, false, false, false,
	false, false, true, true, true, false, true, false, true, true, true, false, true, false, true, true, true, true, true, true, false, false, true, true, true,
	true, true, true, false, false, true, false, false, true, false, false, false, true, true, false, true, false, false, false, true, false, false, true, false, false,
	true, false, false, false, true, false, true, true, true, true, false, false, false, false, true, true, false, true, false, false, true, true, false, true, true,
	true, true, false, true, false, true, false, true, true, false, false, false, true, false, false, false, true, false, true, false, false, false, false, true, true,
	false, false, true, false, false, true, true, true, false, true, false, true, true, true, true, true, false, true, true, true, true, true, true, true, true,
	true, false, true, true, true, false, false, false, true, false, false, true, true, false, false, true, true, false, false, true, false, false, true, false, false,
	true, false, false, false, false, false, true, false, false, true, false, true, false, false, false, false, false, true, true, true, true, true, false, true, true,
	true, false, true, true, true, false, false, false, false, false, true, false, false, false, true, false, true, false, true, true, true, false, false, false, true,
	true, false, true, false, false, true, true, true, false, false, false, true, true, false, true, false, true, true, true, true, true, true, true, false, false,
	false, false, false, false, false, false, false, false, true, false, false, false, false, true, true, false, true, false, false, false, true, false, true, false, false,
	true, true, true, true, true, true, true, false, false, false, false, false, false, true, true, true, true, false, true, false, true, false, true, true, true,
	true, false, false, false, false, false, true, false, false, false, false, true, false, false, false, true, true, false, false, false, true, true, false, true, false,
	true, false, true, true, true, false, true, false, true, false, true, false, false, false, true, true, true, true, true, true, true, true, true, false, false,
	true, false, true, true, true, false, true, false, true, true, false, false, false, true, true, false, false, false, true, false, true, true, false, false, true,
	true, false, true, true, true, false, true, false, true, true, false, true, true, true, true, true, false, false, true, true, false, true, false, false, true,
	true, false, false, false, false, false, true, false, false, true, true, true, false, false, true, true, false, true, false, true, true, false, false, false, true,
	true, true, true, true, true, true, true, false, false, false, false, true, false, false, true, false, true, false, false, true, false, false, true, true, true,
}

func Test_Encode(t *testing.T) {
	res, err := Encode("hello world", H, Unicode)
	if err != nil {
		t.Error(err)
	}
	qrCode, ok := res.(*qrcode)
	if !ok {
		t.Fail()
	}
	if (qrCode.dimension * qrCode.dimension) != len(qrHelloWorldHUni) {
		t.Fail()
	}
	t.Logf("dim %d", qrCode.dimension)
	for i := 0; i < len(qrHelloWorldHUni); i++ {
		x := i % qrCode.dimension
		y := i / qrCode.dimension
		if qrCode.Get(x, y) != qrHelloWorldHUni[i] {
			t.Errorf("Failed at index %d", i)
		}
	}
}

func ExampleEncode() {
	f, _ := os.Create("qrcode.png")
	defer f.Close()

	qrcode, err := Encode("hello world", L, Auto)
	if err != nil {
		fmt.Println(err)
	} else {
		qrcode, err = barcode.Scale(qrcode, 100, 100)
		if err != nil {
			fmt.Println(err)
		} else {
			png.Encode(f, qrcode)
		}
	}
}
