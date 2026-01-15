/**
 * @author Evan Maddicks
 * @version 1.0.0
 * @date 2026-01-15
 * @fileoverview this program creates an playable checkers game, requiring two people to play.
 */
package main

import (
	"fmt"
	"strings"
)

const size = 8

const (
	White = "♙"
	Black = "♟"
	Empty = "□"
	Dark  = "■"
)

func main() {
	var board [size][size]string
	player := White

	// Set up board
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			if (r+c)%2 == 1 {
				if r < 3 {
					board[r][c] = White
				} else if r > 4 {
					board[r][c] = Black
				} else {
					board[r][c] = Empty
				}
			} else {
				board[r][c] = Dark
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

		fc = strings.ToUpper(fc)
		tc = strings.ToUpper(tc)

		if len(fc) != 1 || len(tc) != 1 {
			fmt.Println("Invalid input")
			continue
		}

		fr--
		tr--
		fcIndex := int(fc[0] - 'A')
		tcIndex := int(tc[0] - 'A')

		if validMove(board, player, fr, fcIndex, tr, tcIndex) {
			// Check for jump
			if abs(tr-fr) == 2 {
				midR := (fr + tr) / 2
				midC := (fcIndex + tcIndex) / 2
				board[midR][midC] = Empty
			}

			board[tr][tcIndex] = player
			board[fr][fcIndex] = Empty

			// Switch player
			if player == White {
				player = Black
			} else {
				player = White
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
	// Bounds check
	if fr < 0 || fr >= size || fc < 0 || fc >= size ||
		tr < 0 || tr >= size || tc < 0 || tc >= size {
		return false
	}

	if b[fr][fc] != p || b[tr][tc] != Empty {
		return false
	}

	rowDiff := tr - fr
	colDiff := tc - fc

	// Normal move
	if abs(rowDiff) == 1 && abs(colDiff) == 1 {
		if p == White && rowDiff == 1 {
			return true
		}
		if p == Black && rowDiff == -1 {
			return true
		}
	}

	// Jump move
	if abs(rowDiff) == 2 && abs(colDiff) == 2 {
		midR := (fr + tr) / 2
		midC := (fc + tc) / 2

		if p == White && rowDiff != 2 {
			return false
		}
		if p == Black && rowDiff != -2 {
			return false
		}

		// Must jump opponent
		if p == White && b[midR][midC] == Black {
			return true
		}
		if p == Black && b[midR][midC] == White {
			return true
		}
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}