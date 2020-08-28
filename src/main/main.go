package main

import (
	"afr"
	"afr/settings"
	"afr/utils"
	"github.com/sevlyar/go-daemon"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const logfile = "/sdcard/afr.log"
const Version = "1.09"

func main() {

	if len(os.Args) == 2 && os.Args[1] == "install" {
		if !install() {
			os.Exit(1)
		}
		return
	}

	if len(os.Args) == 2 && os.Args[1] == "uninstall" {
		uninstall()
		return
	}

	settings.Set4K(utils.Is4K())

	if len(os.Args) == 2 {
		wh := strings.ToLower(os.Args[1])
		wha := strings.Split(wh, "x")
		if len(wha) == 2 {
			w, _ := strconv.Atoi(wha[0])
			h, _ := strconv.Atoi(wha[1])
			afr.Test(w, h)
			return
		}
	}
	if len(os.Args) > 2 {
		ws := os.Args[1]
		hs := os.Args[2]
		w, _ := strconv.Atoi(ws)
		h, _ := strconv.Atoi(hs)
		afr.Test(w, h)
		return
	}

	if len(os.Args) == 2 {
		afr.TestAll()
		return
	}

	cntxt := &daemon.Context{}
	child, err := cntxt.Reborn()

	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if child != nil {
		os.Exit(0)
	}
	defer cntxt.Release()

	for true {
		_, err := os.Lstat("/sdcard/")
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	os.Remove(logfile)
	ff, err := os.OpenFile(logfile, os.O_CREATE|os.O_RDWR, 0666)
	if err == nil {
		log.SetOutput(ff)
		defer ff.Close()
	}

	log.Println("Start afr " + Version)

	fr := afr.NewAFR()
	fr.Start()
}
