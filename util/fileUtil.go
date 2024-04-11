package util

import (
	"bufio"
	"os"
	"strings"
)

type fileUtil struct {
}

var fileUtilInstance fileUtil

func GetInstanceByFileUtil() *fileUtil {
	return &fileUtilInstance
}

// IsExists 判断所给路径文件/文件夹是否存在
func (*fileUtil) IsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// GetNewFile 获取文件夹输入句柄，不存在会新建，存在会覆盖
// defer f.Close()
// w.WriteString
// w.Flush()
func (*fileUtil) GetNewFile(filePath string) (*os.File, *bufio.Writer, error) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, nil, err
	}
	wr := bufio.NewWriter(file)
	return file, wr, nil
}

// IsImageMimeType 是图片 返回true
// image/gif	gif	GIF 图像格式
// image/jpeg	jpg, jpeg	JPG(JPEG) 图像格式
// image/jp2	jpg2	JPG2 图像格式
// image/png	png	PNG 图像格式
// image/tiff	tif, tiff	TIF(TIFF) 图像格式
// image/bmp	bmp	BMP 图像格式（位图格式）
// image/svg+xml	svg, svgz	SVG 图像格式
// image/webp	webp	WebP 图像格式
// image/x-icon	ico	ico 图像格式，通常用于浏览器 Favicon 图标
func (*fileUtil) IsImageMimeType(mimeType string) bool {
	return strings.HasPrefix(mimeType, "image")
}

// IsVideoMimeType 是视频文件返回 true
// video/mp4	mp4	mp4 视频格式
// video/mpeg	mpg, mpe, mpeg	mpeg 视频格式
// video/quicktime	qt, mov	QuickTime 视频格式
// video/x-m4v	m4v	m4v 视频格式
// video/x-ms-wmv	wmv	wmv 视频格式（Windows 操作系统上的一种视频格式）
// video/x-msvideo	avi	avi 视频格式
// video/webm	webm	webm 视频格式
// video/x-flv	flv	一种基于 flash 技术的视频格式
func (*fileUtil) IsVideoMimeType(mimeType string) bool {
	return strings.HasPrefix(mimeType, "video")
}

// IsWordMimeType 文档、文本、XML 类型
// application/msword	doc	微软 Office Word 格式（Microsoft Word 97 - 2004 document）
// application/vnd.openxmlformats-officedocument.wordprocessingml.document	docx	微软 Office Word 文档格式
// application/vnd.ms-excel	xls	微软 Office Excel 格式（Microsoft Excel 97 - 2004 Workbook
// application/vnd.openxmlformats-officedocument.spreadsheetml.sheet	xlsx	微软 Office Excel 文档格式
// application/vnd.ms-powerpoint	ppt	微软 Office PowerPoint 格式（Microsoft PowerPoint 97 - 2003 演示文稿）
// application/vnd.openxmlformats-officedocument.presentationml.presentation	pptx	微软 Office PowerPoint 文稿格式
// application/x-gzip	gz, gzip	GZ 压缩文件格式
// application/zip	zip, 7zip	ZIP 压缩文件格式
// application/rar	rar	RAR 压缩文件格式
// application/x-tar	tar, tgz	TAR 压缩文件格式
// application/pdf	pdf	PDF 是 Portable Document Format 的简称，即便携式文档格式
// application/rtf	rtf	RTF 是指 Rich Text Format，即通常所说的富文本格式
// application/kswps	wps	金山 Office 文字排版文件格式
// application/kset	et	金山 Office 表格文件格式
// application/ksdps	dps	金山 Office 演示文稿格式
// application/x-photoshop	psd	Photoshop 源文件格式
// application/x-coreldraw	cdr	Coreldraw 源文件格式
// application/x-shockwave-flash	swf	Adobe Flash 源文件格式
// text/plain	txt	普通文本格式
// text/xml	xml	XML 文件格式
func (*fileUtil) IsWordMimeType(mimeType string) bool {

	if "text/plain" == mimeType {
		return true
	}
	if "text/xml" == mimeType {
		return true
	}

	if !strings.HasPrefix(mimeType, "application") {
		return false
	}

	if "application/x-httpd-php" == mimeType {
		return false
	}
	if "application/java-archive" == mimeType {
		return false
	}
	if "application/vnd.android.package-archive" == mimeType {
		return false
	}
	if "application/octet-stream" == mimeType {
		return false
	}
	if "application/x-x509-user-cert" == mimeType {
		return false
	}
	if "application/x-javascript" == mimeType {
		return false
	}
	if "application/xhtml+xml" == mimeType {
		return false
	}

	return true
}

// IsAudioMimeType 音频格式
// audio/mpeg	mp3	mpeg 音频格式
// audio/midi	mid, midi	mid 音频格式
// audio/x-wav	wav	wav 音频格式
// audio/x-mpegurl	m3u	m3u 音频格式
// audio/x-m4a	m4a	m4a 音频格式
// audio/ogg	ogg	ogg 音频格式
// audio/x-realaudio	ra	Real Audio 音频格式
func (*fileUtil) IsAudioMimeType(mimeType string) bool {
	return strings.HasPrefix(mimeType, "audio")
}

//application/x-javascript	js	Javascript 文件类型
//text/javascript	js	表示 Javascript 脚本文件
//text/css	css	表示 CSS 样式表
//text/html	htm, html, shtml	HTML 文件格式
//application/xhtml+xml	xht, xhtml	XHTML 文件格式
//text/x-vcard	vcf	VCF 文件格式
//application/x-httpd-php	php, php3, php4, phtml	PHP 文件格式
//application/java-archive	jar	Java 归档文件格式
//application/vnd.android.package-archive	apk	Android 平台包文件格式
//application/octet-stream	exe	Windows 系统可执行文件格式
//application/x-x509-user-cert	crt, pem	PEM 文件格式
