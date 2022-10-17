package main

import (
	"project/common/cmd"
	simplebetCmd "project/project/cmd"
)

func main() {
	// initialize additional flag options
	cmd.Initialize()
	simplebetCmd.Execute()
}
