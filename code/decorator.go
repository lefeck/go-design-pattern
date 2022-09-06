package main

import (
	"fmt"
	"time"
)

/* 媒体播放接口， Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。我们以下定义
videoPlay和BarrageVideoPlay的struct， 都有共同方法的属性，
*/
type MediaPlayer interface {
	GetMediaName() string
	GetMediaSeconds() int
	GetMediaContext() string
	play()
}

//视频播放类
type VideoPlay struct {
	MediaPlayer
	videoContent string
	videoName    string
}

func (v *VideoPlay) GetMediaName() string {
	return v.videoName
}

func (v *VideoPlay) GetMediaSeconds() int {
	return len(v.videoContent)
}

func (v *VideoPlay) GetMediaContext() string {
	return v.videoContent
}

//模拟媒体播放
func (m *VideoPlay) play() {
	for i := 0; i < m.GetMediaSeconds(); i++ {
		time.Sleep(time.Millisecond)
		fmt.Println("当前播放第", i, "秒"+m.GetMediaContext())
	}
}

//弹幕视频播放
type BarrageVideoPlay struct {
	MediaPlayer
}

type BarrageReVideoPlay struct {
	BarrageVideoPlay
}

func (b *BarrageReVideoPlay) GetMediaName() string {
	return b.MediaPlayer.GetMediaName() + "开启弹幕"
}

func (b *BarrageReVideoPlay) GetMediaSeconds() int {
	return b.MediaPlayer.GetMediaSeconds()
}

func (b *BarrageReVideoPlay) GetMediaContext() string {
	return "弹幕中---" + b.MediaPlayer.GetMediaContext()
}

func (m *BarrageReVideoPlay) play() {
	for i := 0; i < m.GetMediaSeconds(); i++ {
		time.Sleep(time.Millisecond)
		fmt.Println("当前播放第", i, "秒"+m.GetMediaContext())
	}
}

func main() {
	med := &VideoPlay{
		videoContent: "大闹天宫",
		videoName:    "西游记",
	}
	b := &BarrageReVideoPlay{BarrageVideoPlay{med}}
	med.play()
	b.play()
}
