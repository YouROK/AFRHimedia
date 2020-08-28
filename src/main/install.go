package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

const initdScript = `#!/system/bin/sh
/system/bin/afr
exit 0
`

func install() bool {
	fmt.Println("Install afr")

	if os.Getuid() != 0 {
		fmt.Println("Must be run as root")
		return false
	}

	exec.Command("mount", "-o", "rw,remount", "/system").Run()
	defer exec.Command("mount", "-o", "ro,remount", "/system").Run()

	fmt.Println("Copy afr to bin...")
	_, err := copyFile(os.Args[0], "/system/bin/afr")
	if err != nil {
		fmt.Println("Error copy:", err)
		return false
	}

	fmt.Println("Change file owner")
	err = os.Chown("/system/bin/afr", 0, 2000)
	if err != nil {
		fmt.Println("Error set owner:", err)
	}

	fmt.Println("Change file permission")
	err = os.Chmod("/system/bin/afr", os.ModeSetgid|os.ModeSetuid|0755)
	if err != nil {
		fmt.Println("Error set permission:", err)
		return false
	}

	if _, err := os.Lstat("/system/etc/init.d"); os.IsNotExist(err) {
		fmt.Println("init.d not found")
		return false
	}

	fmt.Println("Create init.d autostart script...")
	ff, err := os.Create("/system/etc/init.d/03afr")
	if err != nil {
		fmt.Println("Error create autostart script:", err)
		return false
	}
	_, err = ff.WriteString(initdScript)
	if err != nil {
		fmt.Println("Error write autostart script:", err)
		return false
	}
	fmt.Println("Change script permission")
	err = os.Chmod("/system/etc/init.d/03afr", 0755)
	if err != nil {
		fmt.Println("Error set permission:", err)
	}
	fmt.Println("Change script owner")
	err = os.Chown("/system/etc/init.d/03afr", 0, 2000)
	if err != nil {
		fmt.Println("Error set owner:", err)
	}
	return true
}

func uninstall() {
	fmt.Println("Uninstall afr")

	exec.Command("mount", "-o", "rw,remount", "/system").Run()
	defer exec.Command("mount", "-o", "ro,remount", "/system").Run()

	if err := os.Remove("/system/bin/afr"); err != nil && !os.IsNotExist(err) {
		fmt.Println("Error remove afr:", err)
	}
	if err := os.Remove("/system/etc/init.d/03afr"); err != nil && !os.IsNotExist(err) {
		fmt.Println("Error remove script:", err)
	}
}

func copyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
