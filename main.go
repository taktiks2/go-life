package main

import (
	"fmt"
	"github.com/mattn/go-tty"
	"log"
	"os"
	"os/exec"
	"runtime"
)

const (
	WIDTH  = 45
	HEIGHT = 45
)

// golangのenum実装
const (
	DEATH = iota
	LIVE
	CURSOR
)

const (
	SWITCH_CELL = 13
	NEXT_GEN    = 32
	UP          = 119
	DOWN        = 115
	LEFT        = 97
	RIGHT       = 100
)

var board [][]int
var cursorY, cursorX int

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

		// キー入力待ち
		rune, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		switch rune {
		case SWITCH_CELL:
			board[cursorY][cursorX] ^= LIVE
		case NEXT_GEN:
			updateBoard()
		case UP:
			if cursorY == 0 {
				cursorY = HEIGHT - 1
			} else {
				cursorY--
			}
		case DOWN:
			if cursorY == HEIGHT-1 {
				cursorY = 0
			} else {
				cursorY++
			}
		case LEFT:
			if cursorX == 0 {
				cursorX = WIDTH - 1
			} else {
				cursorX--
			}
		case RIGHT:
			if cursorX == WIDTH-1 {
				cursorX = 0
			} else {
				cursorX++
			}
		}
	}
}

func createBoard() (newBoard [][]int) {
	newBoard = make([][]int, HEIGHT)
	for i := range newBoard {
		newBoard[i] = make([]int, WIDTH)
	}
	return
}

func initBoard() {
	board = createBoard()
	cursorY = HEIGHT / 2
	cursorX = WIDTH / 2
}

func updateBoard() {
	tempBoard := createBoard()
	// ボード全体用ループ
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			// 周囲の生きたセルカウント用
			aroundLives := 0
			// 周囲のセル用ループ
			for i := y - 1; i <= y+1; i++ {
				for j := x - 1; j <= x+1; j++ {
					// ボードの端同士をつなぐ
					ii := (i + HEIGHT) % HEIGHT
					jj := (j + WIDTH) % WIDTH
					if ii == y && jj == x {
						continue
					}
					if board[ii][jj] == LIVE {
						aroundLives++
					}
				}
			}
			// 次世代の生死ジャッジ
			if board[y][x] == LIVE {
				if aroundLives < 2 || aroundLives > 3 {
					// 過疎 or 過密
					tempBoard[y][x] = DEATH
				} else {
					// 生存
					tempBoard[y][x] = LIVE
				}
			} else {
				if aroundLives == 3 {
					// 誕生
					tempBoard[y][x] = LIVE
				} else {
					tempBoard[y][x] = DEATH
				}
			}
		}
	}
	board = tempBoard
}

func printBoard(board [][]int) {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			switch {
			case y == cursorY && x == cursorX:
				fmt.Print("<>")
			case board[y][x] == LIVE:
				fmt.Print("[]")
			default:
				fmt.Print(" .")
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
