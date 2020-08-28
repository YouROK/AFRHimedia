package fps

import (
	"afr/utils"
	"io/ioutil"
	"strconv"
	"strings"
)

func getFileInfoFps() float64 {
	fpss := ""
	if buf, err := ioutil.ReadFile("/proc/hisi/hiplayer00/fileinfo"); err == nil {
		line := utils.GetLine(string(buf), "fps:")
		fpsa := strings.Split(line, ":")
		if len(fpsa) > 1 {
			fpss = strings.TrimSpace(fpsa[1])
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
