package fps

import (
	"afr/settings"
	"afr/utils"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var reFrameInfo *regexp.Regexp

func getVDecFps() float64 {
	fpss := ""
	if buf, err := ioutil.ReadFile("/proc/msp/vdec00"); err == nil {
		fnd := utils.GetLine(string(buf), "FrameRate(")
		if reFrameInfo == nil {
			reFrameInfo, _ = regexp.Compile(`FrameInfo\((.+?)\)`)
		}
		fpsinfo := reFrcInRate.FindAllStringSubmatch(fnd, -1)
		if len(fpsinfo) > 0 && len(fpsinfo[0]) > 0 && len(fpsinfo[0][1]) > 2 {
			fpss = fpsinfo[0][1]
		}
	} else {
		return -1
	}
	fps, err := strconv.ParseFloat(fpss, 64)
	if err != nil {
		return -1
	}
	fps = fps / 1000
	return fps
}

func getOmxFps() float64 {
	if settings.LoadSettings().DisableOmx() {
		return -1
	}
	fpss := ""
	if buf, err := ioutil.ReadFile("/proc/msp/omxvdec"); err == nil {
		fnd := utils.GetLine(string(buf), "FrameRate(")
		//FrameRate(Src)            :25000(25000)
		arr := strings.Split(fnd, ":")
		if len(arr) > 1 {
			arr = strings.Split(arr[1], "(")
			if len(arr) > 1 {
				fpss = arr[0]
				//if len(fpss) > 0 {
				//	fpss = fpss[:len(fpss)-1]
				//}
			}
		}
	} else {
		return -1
	}
	fps, err := strconv.ParseFloat(fpss, 64)
	if err != nil {
		return -1
	}
	fps = fps / 1000
	return fps
}
