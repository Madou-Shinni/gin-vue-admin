package tools

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"strings"
)

var (
	ErrorFileSizeNotAllow   = errors.New("文件大小不符合要求")
	ErrorFileDecode         = errors.New("文件解析失败")
	ErrorFileBoundsNotAllow = errors.New("文件尺寸不符合要求")
	ErrorFileSuffixNotAllow = errors.New("文件格式不符合要求")
)

var (
	SuffixImages = []string{"jpg", "jpeg", "png"}
)

// FileLimitAttributes 文件上传属性限制
type FileLimitAttributes struct {
	MaxFileSize int64
	Suffix      []string // 限制后缀
	Width       int      // 限制图片宽度
	Height      int      // 限制图片高度
}

// FileLimit 文件上传限制
func FileLimit(fileHeader *multipart.FileHeader, fileAttributes FileLimitAttributes) error {
	// 限制大小
	if fileHeader.Size > fileAttributes.MaxFileSize {
		return ErrorFileSizeNotAllow
	}

	// 获取后缀
	suffix := getFileSuffix(fileHeader)

	// 限制后缀
	flag := contains(fileAttributes.Suffix, suffix)
	if !flag {
		return ErrorFileSuffixNotAllow
	}

	// 如果是图片，限制图片尺寸
	if contains(SuffixImages, suffix) {
		// 限制图片尺寸
		img, err := processImage(fileHeader)
		if err != nil {
			return ErrorFileDecode
		}
		width := img.Bounds().Dx()
		height := img.Bounds().Dy()
		if width != fileAttributes.Width || height != fileAttributes.Height {
			return ErrorFileBoundsNotAllow
		}
	}

	return nil
}

// GetFileSuffix 获取文件后缀
func getFileSuffix(fileHeader *multipart.FileHeader) string {
	fileHeader.Filename = strings.ToLower(fileHeader.Filename)
	return fileHeader.Filename[strings.LastIndex(fileHeader.Filename, ".")+1:]
}

func processImage(file *multipart.FileHeader) (image.Image, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// Determine the image format
	_, format, err := image.DecodeConfig(src)
	if err != nil {
		return nil, err
	}

	// Reset the reader to the beginning
	src.Seek(0, io.SeekStart)

	// Decode the image based on the format
	var img image.Image
	switch strings.ToLower(format) {
	case "jpeg", "jpg":
		img, err = jpeg.Decode(src)
	case "png":
		img, err = png.Decode(src)
	default:
		return nil, fmt.Errorf("unsupported image format: %s", format)
	}

	if err != nil {
		return nil, err
	}

	return img, nil
}

// contains 函数检查字符串是否在字符串数组中
func contains(arr []string, str string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}
