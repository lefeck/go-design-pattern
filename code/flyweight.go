package main

import "fmt"

/*
享元模式是一种结构型设计模式， 它摒弃了在每个对象中保存所有数据的方式， 通过共享多个对象所共有的相同状态， 让你能在有限的内存容量中载入更多对象。
*/

//享元工厂
type ImageFlyweightFactory struct {
	maps map[string]*ImageFlyweight
}

func NewImageFlyweightFactory() *ImageFlyweightFactory {
	return &ImageFlyweightFactory{
		make(map[string]*ImageFlyweight),
	}
}

func (i *ImageFlyweightFactory) Get(filename string) *ImageFlyweight {
	images := i.maps[filename]
	if images == nil {
		images = NewImageFlyweight(filename)
		i.maps[filename] = images
	}
	return images
}

//具体享元对象
type ImageFlyweight struct {
	data string
}

func NewImageFlyweight(filename string) *ImageFlyweight {
	data := fmt.Sprintf("image data %s", filename)
	return &ImageFlyweight{
		data: data,
	}
}

func (i *ImageFlyweight) Data() string {
	return i.data
}

type ImageViewer struct {
	*ImageFlyweight
}

func NewImageViewer(filename string) *ImageViewer {
	image := NewImageFlyweightFactory().Get(filename)
	return &ImageViewer{
		image,
	}
}

func (i *ImageViewer) Display() {
	fmt.Printf("Display: %s\n", i.Data())
}

func main() {
	img := NewImageViewer("image1.png")
	img.Display()
}
