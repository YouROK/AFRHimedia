package afr

import (
	"afr/fps"
	"afr/hidisp"
	"afr/resolution"
	"afr/settings"
	"afr/utils"
	"fmt"
	"log"
	"time"
)

const defPause = 3

type Afr struct {
}

func NewAFR() *Afr {
	a := new(Afr)
	return a
}

func (a *Afr) Start() {
	lastFps := 0.0
	lastRes := ""
mainLoop:
	for {
		time.Sleep(time.Second)

		if settings.LoadSettings().IsDisabled() {
			continue
		}

		f, omx := fps.GetFps()
		res := resolution.GetResolution()

		if lastFps == f && res == lastRes {
			continue
		}
		log.Println()
		log.Println("Receive res/fps/omx:", res, f, omx)

		if f < 0 || res == "MainMenu" { //set main menu
			log.Println("Pause befor main menu")
			pause := settings.LoadSettings().PauseMainMenu()
			pauseTimer := time.NewTimer(time.Duration(pause) * time.Second)
			time.Sleep(time.Second * defPause)
			for f < 0 {
				select {
				case <-pauseTimer.C:
					{
						if setMainMenu() {
							lastFps = f
							lastRes = res
							continue mainLoop
						}
					}
				default:
					{
						time.Sleep(time.Second)
						f, _ = fps.GetFps()
						res = resolution.GetResolution()
						if f != -1 && res != "MainMenu" {
							log.Println("Main menu stop:", f, res)
							pauseTimer.Stop()
							break
						}
					}
				}
			}
		}

		res = resolution.GetResolution()
		if res == "MainMenu" {
			continue
		}

		f, pause := fps.GetFps()
		if pause {
			log.Println("Pause before change")
			lastf := f
			i := 0
			for {
				if f != lastf {
					log.Println("Fps changed", f, ", wait", defPause, "sec")
					lastf = f
					i = 0
				}
				if i >= defPause {
					break
				}
				time.Sleep(time.Second)
				f, _ = fps.GetFps()
				i++
			}
		}

		if f == -1 {
			continue
		}

		log.Printf("Set mode: %v@%v ***\n", res, f)
		fmt, cs, dc, err := settings.GetMode(res, f)
		if err != nil {
			log.Println(err)
			continue
		}
		setMode(fmt, cs, dc)
		lastFps = f
		lastRes = res
	}
}

func setMainMenu() bool {
	log.Println("Set main menu ***")
	fmt, cs, dc, err := settings.GetMode("mainmenu", 0)
	if err != nil {
		log.Println("Error set afr main menu:", err)
		return false
	}
	setMode(fmt, cs, dc)
	return true
}

func setMode(fmt, cs, dc string) {
	f := hidisp.Formats[fmt]
	log.Printf("Find fmt:%v R:%v@%v CS:%v DC:%v", fmt, f.Resolution, f.Fps, cs, dc)
	setFmt(fmt)
	setColors(cs, dc)
}

func setFmt(fmt string) {
	ret := hidisp.Setfmt(fmt)
	if ret != "" {
		log.Println("Result set fmt:", ret)
	}
}

func setColors(cs, dc string) {
	cs = getColorSpace(cs)
	dc = getDeepColor(dc)
	hidisp.SetColors(cs, dc)
}

func Test(width, height int) {
	fpsa := []float64{
		60,
		59.94,
		50,
		30,
		29.97,
		25,
		24,
		23.976,
	}
	fmt.Printf("Test resolution: %dX%d is 4k:%v\n", width, height, utils.Is4K())
	mode := resolution.GetResType(width, height)
	for _, fps := range fpsa {
		ft, cs, dc, err := settings.GetMode(mode, fps)
		r := hidisp.Formats[ft].Resolution
		fmt.Printf("%vX%v@%v\t -> fmt:%v res:%v color space:%v deep color:%v err:%v\n", width, height, fps, ft, r, cs, dc, err)
	}
	fmt.Println()
}

func TestAll() {
	Test(3840, 2160)
	Test(1920, 1080)
	Test(1440, 1080)
	Test(1280, 720)
	Test(768, 576)
	Test(640, 480)
}
