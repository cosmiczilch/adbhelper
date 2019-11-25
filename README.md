Required packages:
go get gopkg.in/urfave/cli.v1

################################################################################

NAME:
   adbhelper - A simple helper for common adb tasks like [re]installing an apk and [re]starting an app

USAGE:
   adbhelper [global options] command [command options] [arguments...]

VERSION:
   1.0

COMMANDS:
   packagename, p  print the current package name
   install, i      install the apk of an app
   start, s        Start an installed app. If the app is already running, it will be restarted
   help, h         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
