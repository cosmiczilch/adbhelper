package main

import (
	"errors"
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Version = "1.0"
	app.Usage = "A simple helper for common adb tasks like [re]installing an apk and [re]starting an app"

	app.Commands = []cli.Command{
		{
			Name:    "packagename",
			Aliases: []string{"p"},
			Usage:   "print the current package name",

			Action:  func(c *cli.Context) error {

				fmt.Println("Package name:", GetPackageName())
				fmt.Println("To change the package name set the environment variable $ADBHELPER_PACKAGE_NAME to com.yourcompany.yourapp")

				return nil
			},
		},
		{
			Name:    "install",
			Aliases: []string{"i"},
			Flags: []cli.Flag {
				cli.BoolFlag{
					Name: "delete, d",
					Usage: "Delete the app first if already installed",
				},
				cli.BoolFlag{
					Name: "start, s",
					Usage: "Start the app automatically after installing (Will be restarted if already running)",
				},
			},
			Usage:   "install the apk of an app",

			Action:  func(c *cli.Context) error {

				if c.Bool("delete") {
					p := GetPackageName()
					fmt.Println("Uninstalling " + p)
					UninstallApk(p)
				}

				if len(c.Args()) != 1 {
					return errors.New("please provide valid path to apk")
				}
				apk := c.Args()[0]
				InstallApk(apk)

				if c.Bool("start") {
					p := GetPackageName()
					fmt.Println("Starting " + p)
					StartApp(p)
				}

				return nil
			},
		},
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "Start an installed app. If the app is already running, it will be restarted",

			Action:  func(c *cli.Context) error {

				p := GetPackageName()

				fmt.Println("Stopping " + p)
				StopApp(p)

				fmt.Println("Starting " + p)
				StartApp(p)

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
