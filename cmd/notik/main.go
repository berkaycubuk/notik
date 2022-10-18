/*
 * Copyright (c) Berkay Çubuk <berkay@berkaycubuk.com>, 2022
 */

package main

import (
	"os"
	"fmt"

	"git.berkaycubuk.com/berkay/notik/pkg/cmd/version"
	"git.berkaycubuk.com/berkay/notik/pkg/cmd/help"
)

// Constants
const Version = "v1.0.0"

func main() {
	if len(os.Args) == 1 {
		fmt.Println("hello")
		return
	}

	// Command switcher
	switch os.Args[1] {
	case "new":
	case "open":
	case "git":
	case "version":
		version.VersionCmd()
	case "help":
		help.HelpCmd()
	}

	fmt.Println("done.")
}
