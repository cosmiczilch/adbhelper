package main

import (
	"log"
	"os"
	"os/exec"
)

func GetPackageName() string {
	p := os.Getenv("ADBHELPER_PACKAGE_NAME")
	if p == "" {
		log.Fatal("Package name not set. Set the environment variable $ADBHELPER_PACKAGE_NAME to com.yourcompany.yourapp")
	}
	return p
}

func InstallApk(apk string) {
	runCommand("adb", "install", apk)
}

func UninstallApk(packageName string) {
	runCommand("adb", "uninstall", packageName)
}

func StartApp(packageName string) {
	runCommand("adb", "shell", "am", "start", "-n", packageName + "/com.unity3d.player.UnityPlayerActivity")
}

func StopApp(packageName string) {
	// adb shell am force-stop
	runCommand("adb", "shell", "am", "force-stop", packageName)
}

func runCommand(name string, arg ...string) {
	command := exec.Command(name, arg...)
	command.Env = os.Environ()
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		log.Fatal("Error executing command: " + err.Error())
	}
}