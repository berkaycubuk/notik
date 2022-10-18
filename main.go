/*
 * Copyright (c) Berkay Çubuk <berkay@berkaycubuk.com>, 2022
 */

package main

import (
	"os"
	"fmt"
)

// Constants
const Version = "v1.0.0"

func main() {
	// Help text
	if len(os.Args) == 1 {
		fmt.Println("hello")
		return
	}

	// Command switcher
	switch os.Args[1] {
	case "new":
	case "open":
	case "git":
	case "help":
	}

	fmt.Println("done.")
}
