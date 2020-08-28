package fps

import "log"

func GetFps() (float64, bool) {
	fps := getFileInfoFps()
	if fps > 0 {
		return FixFps(fps), false
	}

	fps = getAVPlayFps()
	if fps > 0 {
		return FixFps(fps), false
	}

	fps = getVDecFps()
	if fps > 0 {
		return FixFps(fps), true
	}

	fps = getOmxFps()
	if fps > 0 {
		return roundFps(FixFps(fps)), true
	}
	return -1, false
}

var fpsstd = []float64{
	23.976,
	24,
	25,
	29.97,
	30,
	50,
	59.94,
	60,
}

func roundFps(fps float64) float64 {
	var bi, le float64
	for i, f := range fpsstd {
		if f == fps {
			return fps
		}
		if f > fps {
			bi = f
			if i > 0 {
				le = fpsstd[i-1]
			}
		}
	}
	bid := bi - fps
	led := fps - le
	set := fps

	if bid < led {
		set = bi
	}
	if led < bid {
		set = le
	}

	log.Printf("Round fps %v -> %v\n", fps, set)

	return set
}

func FixFps(fps float64) float64 {
	if fps <= 0 {
		return fps
	}

	fi := int64(fps)

	//undouble
	if fi > 45 && fi < 48 {
		fi = 23
	}
	if fi >= 48 && fi < 50 {
		fi = 24
	}
	if fi == 100 {
		fi = 50
	}
	if fi > 116 && fi < 120 {
		fi = 59
	}
	if fi >= 120 {
		fi = 60
	}

	//fix float
	if fi == 59 {
		return 59.94
	}
	if fi == 29 {
		return 29.97
	}
	if fi == 23 {
		return 23.976
	}
	return float64(fi)
}
