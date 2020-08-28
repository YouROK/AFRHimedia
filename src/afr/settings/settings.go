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
	"4096X2160":       "444, 12",
	"4096X2160@59.94": "4096X2160@60, 420, 10",
	"4096X2160@30":    "4096X2160@60, 444, 12",
	"4096X2160@25":    "4096X2160@50, 444, 12",

	"3840X2160":       "444, 12",
	"3840X2160@59.94": "3840X2160@60, 420, 10",
	"3840X2160@30":    "3840X2160@60, 444, 12",
	"3840X2160@25":    "3840X2160@50, 444, 12",

	"1080":        "444, 12",
	"1080@60":     "422, 8",
	"1080@30":     "1080@60, 444, 12",
	"1080@29.970": "1080@59.94, 444, 12",
	"1080@25":     "1080@50, 444, 12",
	"1080@24":     "444, 12",

	"1080i":        "1080, 444, 8",
	"1080i@60":     "444, 8",
	"1080i@59.94":  "444, 8",
	"1080i@29.970": "1080i@59.94, 444, 8",
	"1080i@50":     "444, 8",

	"720":        "1080, 444, 8",
	"720@60":     "444, 8",
	"720@59.94":  "444, 8",
	"720@29.970": "720@59.94, 444, 8",
	"720@50":     "444, 8",
	"720@30":     "720@60, 444, 8",
	"720@25":     "720@50, 444, 8",

	"576":    "444, 12",
	"576@25": "576@50, 444, 12",

	"480":    "1080, 444, 12",
	"480@60": "444, 12",

	"MainMenu":      "1080@60, 444, 12",
	"Default":       "3840X2160@23.976, 444, 12",
	"Disabled":      "0",
	"PauseMainMenu": "2",
	"OMXDisable":    "0",
}

var defSettings1080 = map[string]string{
	"4096X2160":    "1080, 444, 8",
	"4096X2160@30": "1080@60, 444, 8",
	"4096X2160@25": "1080@50, 444, 8",

	"3840X2160":    "1080, 444, 8",
	"3840X2160@30": "1080@60, 444, 8",
	"3840X2160@25": "1080@50, 444, 8",

	"1080":        "444, 8",
	"1080@30":     "1080@60, 444, 8",
	"1080@29.970": "1080@59.94, 444, 8",
	"1080@25":     "1080@50, 444, 8",

	"1080i":        "1080, 444, 8",
	"1080i@60":     "444, 8",
	"1080i@59.94":  "444, 8",
	"1080i@29.970": "1080i@59.94, 444, 8",
	"1080i@50":     "444, 8",

	"720":        "1080, 444, 8",
	"720@60":     "444, 8",
	"720@59.94":  "444, 8",
	"720@29.970": "720@59.94, 444, 8",
	"720@50":     "444, 8",
	"720@30":     "720@60, 444, 8",
	"720@25":     "720@50, 444, 8",

	"576":    "444, 8",
	"576@25": "576@50, 444, 8",

	"480":    "1080, 444, 8",
	"480@60": "444, 8",

	"MainMenu":      "1080@60, 444, 8",
	"Default":       "1080, 444, 8",
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
