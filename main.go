/**
 * @author Evan Maddicks
 * @version 1.0.0
 * @date 2026-01-15
 * @fileoverview this program creates an playable checkers game, requiring two people to play.
 */
package main

import "fmt"

const size = 8

func main() {
	var board [size][size]string
	player := "♙"

	// Set up board
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			if (r+c)%2 == 1 {
				if r < 3 {
					board[r][c] = "♙"
				} else if r > 4 {
					board[r][c] = "♟"
				} else {
					board[r][c] = "□"
				}
			} else {
				board[r][c] = "■"
			}
		}
	}

	for {
		printBoard(board)
		fmt.Println("Player:", player)

		var fc, tc string
		var fr, tr int

		fmt.Print("Move (example: B 3 C 4): ")
		fmt.Scan(&fc, &fr, &tc, &tr)

		fr--
		tr--
		fcIndex := int(fc[0] - 'A')
		tcIndex := int(tc[0] - 'A')

		if validMove(board, player, fr, fcIndex, tr, tcIndex) {
			board[tr][tcIndex] = player
			board[fr][fcIndex] = "□"
			if player == "♙" {
				player = "♟"
			} else {
				player = "♙"
			}
		} else {
			fmt.Println("Invalid move")
		}
	}
}

func printBoard(b [size][size]string) {
	fmt.Println("\n  A B C D E F G H")
	for r := 0; r < size; r++ {
		fmt.Print(r+1, " ")
		for c := 0; c < size; c++ {
			fmt.Print(b[r][c], " ")
		}
		fmt.Println()
	}
}

func validMove(b [size][size]string, p string, fr, fc, tr, tc int) bool {
	if fr < 0 || fr >= size || fc < 0 || fc >= size ||
		tr < 0 || tr >= size || tc < 0 || tc >= size {
		return false
	}

	if b[fr][fc] != p || b[tr][tc] != "□" {
		return false
	}

	if tc-fc != 1 && tc-fc != -1 {
		return false
	}

	if p == "♙" && tr-fr != 1 {
		return false
	}
	if p == "♟" && tr-fr != -1 {
		return false
	}

	return true
}
