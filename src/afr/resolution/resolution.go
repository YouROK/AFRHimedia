package resolution

func GetResolution() string {
	width, height := getFileInfoRes()
	if width == -1 || height == -1 {
		width, height = getOmxRes()
	}
	return GetResType(width, height)
}

func GetResType(width, height int) string {
	if width == -1 && height == -1 {
		return "MainMenu"
	} else if height <= 480 {
		return "480"
	} else if width <= 768 && height <= 576 {
		return "576"
	} else if width <= 1280 || height <= 720 {
		return "720"
	} else if width <= 1440 {
		return "1080i"
	} else if width <= 1920 || height <= 1080 {
		return "1080"
	} else if width <= 3840 && height <= 2160 {
		return "3840x2160"
	} else if width <= 4096 || height <= 2160 {
		return "4096x2160"
	} else {
		return "1080"
	}
}
