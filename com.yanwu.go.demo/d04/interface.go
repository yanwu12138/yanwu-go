package d04

import (
	"fmt"
)

/**
 * @Author Baofeng Xu
 * @Date 2021/5/13 16:53
 *
 * Description: 接口
 **/

type Device interface {
	// ----- 上线
	online() bool
	// ----- 离线
	offline()
	// ----- 开关
	onOff(onOff bool)
}

type Light struct {
}

func (light Light) online() bool {
	fmt.Println("light online")
	return true
}

func (light Light) offline() {
	fmt.Println("light offline")
}

func (light Light) onOff(onOff bool) {
	if onOff {
		fmt.Println("light on")
	} else {
		fmt.Println("light off")
	}
}

type Edge struct {
}

func (edge Edge) online() bool {
	fmt.Println("edge online")
	return true
}

func (edge Edge) offline() {
	fmt.Println("edge offline")
}

func (edge Edge) onOff(onOff bool) {
	if onOff {
		fmt.Println("edge on")
	} else {
		fmt.Println("edge off")
	}
}

func TestInterface() {
	fmt.Println("====================================== 接口 ======================================")
	var device Device
	device = new(Light)
	device.online()
	device.onOff(true)
	device.onOff(false)
	device.offline()
	fmt.Println("----------")
	device = new(Edge)
	device.online()
	device.onOff(true)
	device.onOff(false)
	device.offline()
}
