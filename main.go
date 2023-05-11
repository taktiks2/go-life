package main

import (
	"fmt"
	"github.com/mattn/go-tty"
	"log"
	"os"
	"os/exec"
	"runtime"
)

const width = 45
const height = 45

// golangのenum実装
const (
	Death = iota
	Live
	Cursor
)

const (
	Up    = "w"
	Down  = "s"
	Left  = "a"
	Right = "d"
)

var board [][]int

func put() {
	board[15][15] = 1
	board[15][16] = 2
	board[15][17] = 1
	board[15][18] = 1
}

func initBoard() {
	board = make([][]int, height)
	for i := range board {
		board[i] = make([]int, width)
	}
	board[22][22] = Cursor
}

func printBoard(board [][]int) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			switch board[i][j] {
			case Death:
				fmt.Print(" .")
			case Live:
				fmt.Print("[]")
			case Cursor:
				fmt.Print("<>")
			}
		}
		fmt.Println()
	}
}

func clear() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func init() {
	initBoard()
}

func main() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()
	for {
		clear()
		printBoard(board)
		rune, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		switch string(rune) {
		case Up:
			fmt.Println(rune)
		case Down:
			fmt.Println(rune)
		case Left:
			fmt.Println(rune)
		case Right:
			fmt.Println(rune)
		}
	}
}
