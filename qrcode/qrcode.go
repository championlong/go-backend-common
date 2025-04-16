package qrcode

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
)

// CreateQrCodeBs64WithLogo  黑色条纹带logo的二维码图片生成 content-二维码内容 size-像素单位 logoPath-logo文件路径
func CreateQrCodeBs64WithLogo(content, logoPath, outPath string, size int) (data string, err error) {
	code, err := qrcode.New(content, qrcode.High)
	if err != nil {
		return
	}
	code.DisableBorder = true
	//设置文件大小并创建画板
	qrcodeImg := code.Image(size)
	outImg := image.NewRGBA(qrcodeImg.Bounds())

	//读取logo文件
	logoFile, err := os.Open(logoPath)
	if err != nil {
		return
	}
	logoImg, _, _ := image.Decode(logoFile)
	logoImg = resize.Resize(uint(size/5), 0, logoImg, resize.Lanczos3)

	// 添加边框
	// 图片到边框距离
	pic2FramePadding := logoImg.Bounds().Dx() / 10

	// 新建一个边框图层
	transparentImg := image.NewRGBA(image.Rect(0, 0, logoImg.Bounds().Dx()+pic2FramePadding, logoImg.Bounds().Dy()+pic2FramePadding))
	// 图层颜色设为白色
	draw.Draw(transparentImg, transparentImg.Bounds(), image.White, image.Point{}, draw.Over)
	// 将缩略图放到透明图层上
	draw.Draw(transparentImg,
		image.Rect(pic2FramePadding/2, pic2FramePadding/2, transparentImg.Bounds().Dx(), transparentImg.Bounds().Dy()),
		logoImg,
		image.Point{},
		draw.Over)

	//logo和二维码拼接
	draw.Draw(outImg, outImg.Bounds(), qrcodeImg, image.Pt(0, 0), draw.Over)
	offset := image.Pt((outImg.Bounds().Max.X-transparentImg.Bounds().Max.X)/2, (outImg.Bounds().Max.Y-transparentImg.Bounds().Max.Y)/2)
	draw.Draw(outImg, outImg.Bounds().Add(offset), transparentImg, image.Pt(0, 0), draw.Over)

	buf := new(bytes.Buffer)
	_ = png.Encode(buf, outImg)

	// 写入文件
	f, err := os.Create(outPath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if err = png.Encode(f, outImg); err != nil {
		return "", err
	}

	res := base64.StdEncoding.EncodeToString(buf.Bytes())
	return res, nil
}

func CreateQrCodeWithBackground(content, logoPath, bgPath, outPath string, qrSize, bgSize int) (data string, err error) {
	code, err := qrcode.New(content, qrcode.High)
	if err != nil {
		return
	}
	code.DisableBorder = true

	// 创建二维码图像
	qrcodeImg := code.Image(qrSize)
	outImg := image.NewRGBA(image.Rect(0, 0, bgSize, bgSize)) // 创建背景尺寸画板

	// 读取背景图片
	bgFile, err := os.Open(bgPath)
	if err != nil {
		return
	}
	defer bgFile.Close()
	bgImg, _, err := image.Decode(bgFile)
	if err != nil {
		return
	}
	bgImg = resize.Resize(uint(bgSize), uint(bgSize), bgImg, resize.Lanczos3)

	// 将背景图绘制到画板上
	draw.Draw(outImg, outImg.Bounds(), bgImg, image.Point{}, draw.Over)

	// 读取 logo 文件
	logoFile, err := os.Open(logoPath)
	if err != nil {
		return
	}
	defer logoFile.Close()
	logoImg, _, err := image.Decode(logoFile)
	if err != nil {
		return
	}
	logoImg = resize.Resize(uint(qrSize/5), 0, logoImg, resize.Lanczos3)

	// 添加边框到 logo
	pic2FramePadding := logoImg.Bounds().Dx() / 10
	transparentImg := image.NewRGBA(image.Rect(0, 0, logoImg.Bounds().Dx()+pic2FramePadding, logoImg.Bounds().Dy()+pic2FramePadding))
	draw.Draw(transparentImg, transparentImg.Bounds(), image.White, image.Point{}, draw.Over)
	draw.Draw(transparentImg,
		image.Rect(pic2FramePadding/2, pic2FramePadding/2, transparentImg.Bounds().Dx(), transparentImg.Bounds().Dy()),
		logoImg,
		image.Point{},
		draw.Over)

	// 将二维码和 logo 合成
	qrWithLogoImg := image.NewRGBA(qrcodeImg.Bounds())
	draw.Draw(qrWithLogoImg, qrWithLogoImg.Bounds(), qrcodeImg, image.Pt(0, 0), draw.Over)
	offsetLogo := image.Pt((qrWithLogoImg.Bounds().Max.X-transparentImg.Bounds().Max.X)/2, (qrWithLogoImg.Bounds().Max.Y-transparentImg.Bounds().Max.Y)/2)
	draw.Draw(qrWithLogoImg, qrWithLogoImg.Bounds().Add(offsetLogo), transparentImg, image.Pt(0, 0), draw.Over)

	// 将二维码绘制到背景图中央
	offsetQR := image.Pt((outImg.Bounds().Dx()-qrWithLogoImg.Bounds().Dx())/2, (outImg.Bounds().Dy()-qrWithLogoImg.Bounds().Dy())/2)
	draw.Draw(outImg, outImg.Bounds().Add(offsetQR), qrWithLogoImg, image.Point{}, draw.Over)

	// 编码为 PNG 格式并保存到文件
	buf := new(bytes.Buffer)
	err = png.Encode(buf, outImg)
	if err != nil {
		return
	}
	f, err := os.Create(outPath)
	if err != nil {
		return
	}
	defer f.Close()
	err = png.Encode(f, outImg)
	if err != nil {
		return
	}

	// 转为 Base64 数据返回
	data = base64.StdEncoding.EncodeToString(buf.Bytes())
	return data, nil
}

// CreateWhiteQrCodeBs64WithLogo 白色纹路的二维码
func CreateWhiteQrCodeBs64WithLogo(content, logoPath, outPath string, size int) (data string, err error) {
	// 创建二维码
	code, err := qrcode.New(content, qrcode.High)
	if err != nil {
		return "", err
	}
	code.DisableBorder = true

	// 设置二维码大小
	qrcodeImg := code.Image(size)
	// 创建全黑色背景的输出图像
	outImg := image.NewRGBA(qrcodeImg.Bounds())
	draw.Draw(outImg, outImg.Bounds(), &image.Uniform{color.Black}, image.Point{}, draw.Src)

	// 反转二维码颜色
	for y := 0; y < qrcodeImg.Bounds().Dy(); y++ {
		for x := 0; x < qrcodeImg.Bounds().Dx(); x++ {
			if qrcodeImg.At(x, y) == color.Black {
				outImg.Set(x, y, color.White) // 白色线条
			}
		}
	}

	// 读取logo文件
	logoFile, err := os.Open(logoPath)
	if err != nil {
		return "", err
	}
	defer logoFile.Close()
	logoImg, _, err := image.Decode(logoFile)
	if err != nil {
		return "", err
	}
	logoImg = resize.Resize(uint(size/5), 0, logoImg, resize.Lanczos3)

	// 添加logo到二维码中心
	offset := image.Pt((outImg.Bounds().Dx()-logoImg.Bounds().Dx())/2, (outImg.Bounds().Dy()-logoImg.Bounds().Dy())/2)
	draw.Draw(outImg, outImg.Bounds().Add(offset), logoImg, image.Point{}, draw.Over)

	// 编码为PNG
	buf := new(bytes.Buffer)
	if err = png.Encode(buf, outImg); err != nil {
		return "", err
	}

	// 写入文件
	f, err := os.Create(outPath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if err = png.Encode(f, outImg); err != nil {
		return "", err
	}

	// 编码为Base64
	res := base64.StdEncoding.EncodeToString(buf.Bytes())
	return res, nil
}
