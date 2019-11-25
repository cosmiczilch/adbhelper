package main

import "log"

type TaskType int
const (
	INSTALL TaskType = iota
	START_APP
	NONE
	INVALID
)

type Option int
const (
	HELP Option = iota
	DELETE_FIRST
	START_AUTOMATICALLY
)

func PrintUsage() {
	log.Println("\n" +
		"Print usage:\n" +
		"  adbhelper --help\n" +
		"Install an apk:\n" +
		"  adbhelper install [options] [package-name]\n" +
		"    valid options:\n" +
		"      -r    reinstall if app already installed\n" +
		"      -s    start app after installing\n" +
		"    package name is something like com.company.appname\n" +
		"      if package name is absent it will be taken \n" +
		"      from the environment variable $PACKAGE_NAME\n" +
		"Start an already installed app:\n" +
		"  adbhelper start [package-name]\n" +
		"    package name is something like com.company.appname\n" +
		"      if package name is absent it will be taken \n" +
		"      from the environment variable $PACKAGE_NAME\n" +
		"    note: if app is already running it will be killed first")
}
