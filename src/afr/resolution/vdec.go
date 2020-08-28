package resolution

import (
	"afr/utils"
	"io/ioutil"
	"strconv"
	"strings"
)

func getOmxRes() (width, height int) {
	if buf, err := ioutil.ReadFile("/proc/msp/omxvdec"); err == nil {
		fnd := utils.GetLine(string(buf), "Resolution")
		//Resolution                :1920x1080
		arr := strings.Split(fnd, ":")
		if len(arr) > 1 {
			arr = strings.Split(arr[1], "x")
			if len(arr) > 1 {
				swidth := strings.TrimSpace(arr[0])
				sheight := strings.TrimSpace(arr[1])
				width, _ = strconv.Atoi(swidth)
				height, _ = strconv.Atoi(sheight)
				return
			}
		}
	}
	return -1, -1
}
