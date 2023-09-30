package main

import (
	// "flag"
	"fmt"
	"os"
)

// func createFile(fileName string) error {
//     file, err := os.Create(fileName)
//     if err != nil {
//         return err
//     }
//     defer file.Close()
//     fmt.Println("Created file: %s\n",fileName)
//     return nil

// }

func help() {
	fmt.Println("Only GOD Can Help You")
}

func helloName(name []string) {
	fmt.Println("Hello ",name[0])
}

func main() {
	// name := flag.String("name", "", "Write Your name")
	// flag.Parse()
	// fmt.Printf("Hello, %s!\n", *name)
	// var avail_cmd[] string = ["-help","-name","-create","-delete"]
	var command string = os.Args[1]
	// fmt.Println(command)
	switch command {
	case "-help":
		help()
	case "-name":
		if os.Args[2:] != nil {
			helloName(os.Args[2:])
		} else {
			fmt.Println("Enter Your Name")
		}
	}
	// help := flag.String("help")
	// flag.Parse()
	// fmt.Printf("")
}
