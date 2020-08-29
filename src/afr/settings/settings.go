package settings

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Settings map[string]string

var defSettings4K = map[string]string{
	"3840X2160":      "2, 2",
        "3840X2160@50, 3, 0",
	"3840X2160@60, 3, 0",
	"1080":          "2, 2",
        "720":           "1080, 2, 2",
	"480":           "1080, 2, 2",
        "576":           "1080, 2, 2",

	"MainMenu":      "1080@60, 2, 2",
	"Default":       "3840X2160@60, 3, 0",
	"Disabled":      "0",
	"PauseMainMenu": "2",
	"OMXDisable":    "0",
}

var defSettings1080 = map[string]string{
	"3840X2160":   "1080, 1, 2",
        "1080":        "1, 2",
        "720":         "1080, 1, 2",
        "480":         "1080, 1, 2",
        "576":         "1080, 1, 2",

	"MainMenu":      "1080@60, 1, 2",
	"Default":       "1080@60, 1, 2",
	"Disabled":      "0",
	"PauseMainMenu": "2",
	"OMXDisable":    "0",
}

var defSettings map[string]string

func Set4K(val bool) {
	if val {
		defSettings = defSettings4K
	} else {
		defSettings = defSettings1080
	}
}

func LoadSettings() Settings {
	buf, err := ioutil.ReadFile(SettingsPath)
	if buf == nil {
		fmt.Println("Error load settings:", err)
		SaveSettings(defSettings)
		return defSettings
	}

	var set Settings = defSettings
	lines := strings.Split(string(buf), "\n")

	for _, l := range lines {
		if strings.HasPrefix(l, "#") {
			continue
		}
		md := strings.Split(l, ":")
		if len(md) == 2 {
			if strings.ToLower(md[0]) == "default" {
				set["Default"] = md[1]
			} else {
				set[md[0]] = md[1]
			}
		}
	}
	return set
}

func SaveSettings(sets Settings) {
	lines := make([]string, 0)
	for k, v := range sets {
		lines = append(lines, k+":"+v)
	}
	sort.Slice(lines, func(i, j int) bool {
		return lines[i] < lines[j]
	})
	str := strings.Join(lines, "\n") + "\n"
	ioutil.WriteFile(SettingsPath, []byte(str), 0666)
}

func (s Settings) IsDisabled() bool {
	if v, ok := s["Disabled"]; ok {
		return v != "0"
	}
	return false
}

func (s Settings) DisableOmx() bool {
	if v, ok := s["OMXDisable"]; ok {
		return v != "0"
	}
	return true
}

func (s Settings) PauseMainMenu() int {
	if v, ok := s["PauseMainMenu"]; ok {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 1
		}
		return int(i)
	}
	return 1
}
