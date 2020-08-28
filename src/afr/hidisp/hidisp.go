package hidisp

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Setfmt(fmt string) string {
	return Hidisp("setfmt", fmt)
}

func SetColors(cs, dc string) bool {
	gcs, gdc := GetColors()
	log.Println("Set colors:", cs, dc)
	if gcs == cs && gdc == dc {
		log.Println("Colors already set")
		return true
	}

	hdmi, err := os.OpenFile("/proc/msp/hdmi0", os.O_RDWR, 0666)
	if err != nil {
		ret := Hidisp("setcolorspaceanddeepcolor", cs, dc)
		return strings.Contains(ret, "ret = 0")
	}
	defer hdmi.Close()

	log.Println("set color space and deep color hdmi:", cs, dc)

	css := fmt.Sprintf("outclrspace %v\n", cs)
	dcs := fmt.Sprintf("deepclr %v\n", dc)
	hdmi.WriteString(css)
	hdmi.WriteString(dcs)
	return true
}

func GetColors() (cs, dc string) {
	cs = "-1"
	dc = "-1"
	buf := Hidisp("getcolorspacemode")
	arr := strings.Split(string(buf), " = ")
	if len(arr) >= 2 {
		cs = arr[1]
		cs = strings.Replace(cs, "=", "", -1)
		cs = strings.TrimSpace(cs)
	}

	buf = Hidisp("getdeepcolormode")
	arr = strings.Split(string(buf), " = ")
	if len(arr) >= 2 {
		dc = arr[1]
		dc = strings.Replace(dc, "=", "", -1)
		dc = strings.TrimSpace(dc)
	}

	log.Println("Get colors:", cs, dc)
	return
}

func Hidisp(arg ...string) string {
	log.Println("hidisp", arg)
	buf, _ := exec.Command("/system/bin/hidisp", arg...).Output()
	if len(buf) > 0 {
		return string(buf)
	}
	return ""
}
