package main

import (
	"cli_proj/url_shortner"
	"fmt"
	"math/rand"
	"os"
	"os/exec"

	"github.com/enescakir/emoji"
)

func ticktacktoe() {
	board := [][]string{{"O", "", "X"}, {"O", "X", "O"}, {"X", "O", "X"}}
	fmt.Println("Tic-Tac-Toe")
	fmt.Println("-------------")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j > 0 {
				fmt.Print(" | ")
			}
			if board[i][j] == "" {
				fmt.Print(" ")
			} else {
				fmt.Print(board[i][j])
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("-------------")
		}
	}
	fmt.Println("-------------")
}

func help() {
	fmt.Println("Only GOD Can Help You", emoji.IndexPointingUp)
}

func helloName(name []string) {
	fmt.Println("Hello ", name[0]+"!", emoji.WavingHand)
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

func removeFile(fileName []string) error {
	err := os.Remove(fileName[0])
	if err != nil {
		fmt.Println("Error while deleting file:", err)
		os.Exit(1)
	}
	fmt.Printf("Deleted file: %s\n", fileName)
	return nil
}

func guessGame() {
	var number int
	randNum := rand.Intn(10)
	fmt.Println("Enter any number between 1 to 10:")
	fmt.Scan(&number)
	if number == randNum {
		fmt.Println("You got it !!")
	} else {
		fmt.Println("Try again You are in Loop !!", emoji.SkullAndCrossbones)
		guessGame()
	}
}

func tickTackToe() {
	ticktacktoe()
}

func chess() {
	fmt.Println("Coming Soon...")
}

func playGame(gameNum int) error {
	if gameNum == 1 {
		guessGame()
	} else if gameNum == 2 {
		tickTackToe()
	} else if gameNum == 3 {
		chess()
	}
	return nil
}

func ping(packets string, target []string) error {
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
		fmt.Println("-name [yourName]          Prints Hello [yourName]")
		fmt.Println("-ping [packets] [target]  Pinging Server")
		fmt.Println("-create [fileName]        This Command will create file", emoji.FloppyDisk)
		fmt.Println("-remove [fileName]        Command for Deleting File", emoji.Wastebasket)
		fmt.Println("-play                     for playing games", emoji.VideoGame)
	}
	if length >= 2 {
		var command string = os.Args[1]
		// fmt.Println(command)
		switch command {
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
				removeFile(os.Args[2:])
			}
		case "-play":
			var number int
			if length == 2 {
				fmt.Println(emoji.BowAndArrow, "Checkout below list of Games", emoji.Joystick)
				fmt.Println("1.Guess the Number", emoji.ThinkingFace)
				fmt.Println("2.Tick-tack-toe", emoji.GameDie)
				fmt.Println("3.Chess", emoji.ChessPawn)
				fmt.Print("Enter the number of the game: ")
				fmt.Scan(&number)
				playGame(number)
			}
		case "-shorten":
			if length >= 3 {
				url_shortner.WillWOrk(os.Args[2:])
			}
		case "-ping":
			if length >= 4 {
				packets := os.Args[2]
				ping(packets, os.Args[3:])
			}
		}
	}
}
