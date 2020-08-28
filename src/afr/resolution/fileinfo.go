package resolution

import (
	"afr/utils"
	"io/ioutil"
	"strconv"
	"strings"
)

func getFileInfoRes() (width, height int) {
	if buf, err := ioutil.ReadFile("/proc/hisi/hiplayer00/fileinfo"); err == nil {
		fnd := utils.GetLine(string(buf), "w * h:")
		tmp := strings.Split(fnd, ":")
		if len(tmp) > 1 {
			tmp := strings.Split(strings.TrimSpace(tmp[1]), "*")
			if len(tmp) == 2 {
				swidth := strings.TrimSpace(tmp[0])
				sheight := strings.TrimSpace(tmp[1])
				width, _ = strconv.Atoi(swidth)
				height, _ = strconv.Atoi(sheight)
				return
			}
		}
	}
	return -1, -1
}
