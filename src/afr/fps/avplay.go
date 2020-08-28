package fps

import (
	"afr/utils"
	"io/ioutil"
	"regexp"
	"strconv"
)

var reFrcInRate *regexp.Regexp

func getAVPlayFps() float64 {
	fpss := ""
	if buf, err := ioutil.ReadFile("/proc/msp/avplay00"); err == nil {
		fnd := utils.GetLine(string(buf), "FrcInRate")
		if reFrcInRate == nil {
			reFrcInRate, _ = regexp.Compile(`FrcInRate.+?:(.+?)\s+|.+`)
		}
		fpsinfo := reFrcInRate.FindAllStringSubmatch(fnd, -1)
		if len(fpsinfo) > 0 && len(fpsinfo[0]) > 0 && len(fpsinfo[0][0]) > 0 {
			fpss = fpsinfo[0][1]
		}
	} else {
		return -1
	}
	fps, err := strconv.ParseFloat(fpss, 64)
	if err != nil {
		return -1
	}
	return fps
}
