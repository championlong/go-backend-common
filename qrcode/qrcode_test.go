package qrcode

import (
	"fmt"
	"testing"
)

func TestQrCode(t *testing.T) {
	_, err := CreateQrCodeBs64WithLogo("https://www.baidu.com",
		"logo.png",
		"qrcode.png", 1200)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestQrCodeBackground(t *testing.T) {
	_, err := CreateQrCodeWithBackground("https://www.baidu.com",
		"logo.png", "background.png",
		"qrcode.png", 128, 283)
	if err != nil {
		fmt.Println(err)
		return
	}
}
