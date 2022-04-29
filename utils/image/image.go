package image

import (
	"bytes"
	"golang.org/x/image/draw"
	"image"
	"image/jpeg"
)

func ResizeJPEG(srcData []byte, width, height int64) ([]byte, error) {
	src, _, err := image.Decode(bytes.NewReader(srcData))
	if err != nil {
		return nil, err
	}

	dst := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	draw.CatmullRom.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, dst, &jpeg.Options{Quality: 100}); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func GetSize(srcData []byte) (int64, int64, error) {
	src, _, err := image.Decode(bytes.NewReader(srcData))
	if err != nil {
		return 0, 0, err
	}

	bounds := src.Bounds()
	return int64(bounds.Dx()), int64(bounds.Dy()), nil
}

func CalcToResizeWidth[T ~int64 | ~int](srcWidth, srcHeight, dstWidth T) T {
	return T((float64(dstWidth) / float64(srcWidth)) * float64(srcHeight))
}

func CalcToResizeHeight[T ~int64 | ~int](srcWidth, srcHeight, dstHeight T) T {
	return T((float64(dstHeight) / float64(srcHeight)) * float64(srcWidth))
}
