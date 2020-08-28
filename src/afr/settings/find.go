package settings

import (
	"afr/hidisp"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func GetMode(mode string, fps float64) (ft, cs, dc string, err error) {
	if fps > 0 {
		ft, cs, dc, err = findMode(mode, fps)
		if err == nil {
			return
		}
		//get default mode
		ft, cs, dc, err = findMode("default", fps)
		if err == nil {
			return
		}
	}
	//get mainmenu
	ft, cs, dc, err = findMode("mainmenu", fps)
	return
}

func findMode(mode string, fps float64) (ft, cs, dc string, err error) {
	log.Println("Find in settings:", mode, fps)

	set := LoadSettings()

	ffps := int(fps)

	fndStr := fmt.Sprintf("%v@%v", mode, ffps)
	//find mode with fps
	for smode, scolor := range set {
		if strings.HasPrefix(strings.ToLower(smode), fndStr) {
			tmode, tfps, tocolor, todeep, err := getColors(smode, scolor)
			if err != nil {
				return "", "", "", err
			}
			if tfps != "" {
				tmp, err := strconv.ParseFloat(tfps, 64)
				if err != nil {
					return "", "", "", err
				}
				ffps = int(tmp)
			}
			f, err := findFmt(tmode, ffps)
			if err != nil {
				return "", "", "", err
			}
			return f, tocolor, todeep, nil
		}
	}

	//find mode without fps
	for smode, scolor := range set {
		if !strings.Contains(smode, "@") {
			if strings.ToLower(smode) == mode {
				tmode, tfps, tocolor, todeep, err := getColors(smode, scolor)
				if err != nil {
					return "", "", "", err
				}
				if tfps != "" {
					tmp, err := strconv.ParseFloat(tfps, 64)
					if err != nil {
						return "", "", "", err
					}
					ffps = int(tmp)
				}
				f, err := findFmt(tmode, ffps)
				if err != nil {
					return "", "", "", err
				}
				return f, tocolor, todeep, nil
			}
		}
	}
	return "", "", "", fmt.Errorf("Mode not found")
}

func findFmt(res string, fps int) (string, error) {
	for n, f := range hidisp.Formats {
		if f.Resolution == res && f.Fps == fps {
			return n, nil
		}
	}
	return "", fmt.Errorf("Format not found: %v@%v", res, fps)
}

func getColors(modeF, colorsF string) (mode, fps, cs, dc string, err error) {
	modeFrom := strings.Split(modeF, "@")
	modeAndColorTo := strings.Split(colorsF, ",")
	if len(modeAndColorTo) == 3 {
		modeTo := strings.Split(modeAndColorTo[0], "@")
		mode = strings.TrimSpace(modeTo[0])
		if len(modeTo) > 1 {
			fps = strings.TrimSpace(modeTo[1])
		} else if len(modeFrom) > 1 {
			fps = modeFrom[1]
		}
		cs = strings.TrimSpace(modeAndColorTo[1])
		dc = strings.TrimSpace(modeAndColorTo[2])
		return
	}
	if len(modeAndColorTo) == 2 {
		if len(modeFrom) > 0 {
			mode = modeFrom[0]
		}
		if len(modeFrom) > 1 {
			fps = modeFrom[1]
		}
		cs = strings.TrimSpace(modeAndColorTo[0])
		dc = strings.TrimSpace(modeAndColorTo[1])
		return
	}
	return "", "", "", "", fmt.Errorf("Wrong mode: %v:%v", modeF, colorsF)
}
