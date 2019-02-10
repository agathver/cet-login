package main

import (
	"fmt"

	docopt "github.com/docopt/docopt-go"
)

func main() {
	usage := `CET Login: A utility for automatic login on the WiFi Network of 
College of Engineering and Technology, Bhubaneswar.

Usage:
  cetlogin profile create <name>
  cetlogin profile edit <name>
  cetlogin profile use <name>
  cetlogin profile set-default <name>
  cetlogin profile remove <name>
  cetlogin login [--profile=<profile>]
  cetlogin -h | --help
  cetlogin --version

Options:
  -h --help                 Show this screen.
  --version                 Show version.
  -d --device=<device>      Device presented to the server during login.
                            Available choices: "mobile","pc". Default: "mobile"
  -p --profile=<profile>    Use the specified profile for login`

	arguments, _ := docopt.ParseDoc(usage)
	fmt.Println(arguments)
}
