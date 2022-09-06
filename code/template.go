package main

import "fmt"

type Downloader interface {
	Download(uri string)
}

//模版类
type Template struct {
	Implement
	uri string
}

//模版接口
type Implement interface {
	download()
	save()
}

func NewTemplate(impl Implement) *Template {
	return &Template{
		Implement: impl,
	}
}

func (t *Template) Download(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	t.Implement.download()
	t.Implement.save()
	fmt.Print("finish downloading\n")
}

func (t *Template) save() {
	fmt.Print("default save\n")
}

//具体类
type HTTPDownloader struct {
	*Template
}

func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	temlate := NewTemplate(downloader)
	downloader.Template = temlate
	return downloader
}

//具体实现
func (d *HTTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

func (*HTTPDownloader) save() {
	fmt.Printf("http save\n")
}

type FTPDownloader struct {
	*Template
}

func NewFTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	temlate := NewTemplate(downloader)
	downloader.Template = temlate
	return downloader
}

func (d *FTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

func (*FTPDownloader) save() {
	fmt.Printf("http save\n")
}

func main() {
	n := NewHTTPDownloader()
	n.Download("http://www.baidu.com")
}
