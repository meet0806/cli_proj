package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

func help() {
	fmt.Println("Only GOD Can Help You")
}

func helloName(name []string) {
	fmt.Printf("Hello %s!\n", name[0])
}

func createFile(fileName []string) error {
	file, err := os.Create(fileName[0])
	if err != nil {
		fmt.Println("Error creating the file:", err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Printf("Created file: %s\n", fileName)
	return nil
}

func deleteFile(fileName []string) error {
	err := os.Remove(fileName[0])
	if err != nil {
		fmt.Println("Error while deleting file:", err)
		os.Exit(1)
	}
	fmt.Printf("Deleted file: %s\n", fileName)
	return nil
}

func tickTackToe() {
	fmt.Println("Coming Soon...")
}

func chess(){
	fmt.Println("Coming Soon...")
}

func playGame(gameNum int) error {
	if gameNum == 1 {
		var number int
		randNum := rand.Intn(10)
		fmt.Println("Enter any number between 1 to 10:")
		fmt.Scan(&number)
		if number == randNum {
			fmt.Println("You got it right")
		} else {
			fmt.Println("Try again!!")
		}
	} else if gameNum == 2 {
		tickTackToe()
	} else if gameNum == 3 {
		chess()
	}
	return nil
}

func ping(packets string, target[] string) error {
	cmd := exec.Command("ping", "-c", packets, target[0]) // -c 4 means send 4 packets
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		fmt.Println("Error running ping command:", err)
		os.Exit(1)
	}
	return nil
}

func main() {
	length := len(os.Args)
	if length == 1 {
		fmt.Println("Get Bored Try Below Commands")
		fmt.Println("-help")
		fmt.Println("-name [yourName]          Prints Hello [yourName]")
		fmt.Println("-ping [packets] [target]  Pinging Server")
		fmt.Println("-create [fileName]        This Command will create file")
		fmt.Println("-delete [fileName]        Command for Deleting File")
		fmt.Println("-play                     for playing games")
	}
	if length >= 2 {
		var command string = os.Args[1]
		// fmt.Println(command)
		switch command {
		case "-help":
			help()
		case "-name":
			if length >= 3 {
				helloName(os.Args[2:])
			} else {
				fmt.Println("Enter Your Name After Command")
			}
		case "-create":
			if length >= 3 {
				createFile(os.Args[2:])
			}
		case "-remove":
			if length >= 3 {
				deleteFile(os.Args[2:])
			}
		case "-play":
			var number int
			if length == 2 {
				fmt.Println("List of Games")
				fmt.Println("1.Guess the Number")
				fmt.Println("2.Tick-tack-toe")
				fmt.Println("3.Chess")
				fmt.Print("Enter the number of the game: ")
				fmt.Scan(&number)
				playGame(number)
			}
		case "-ping":
			if length >= 4 {
				packets := os.Args[2]
				ping(packets, os.Args[3:])
			}
		}
	}
}
