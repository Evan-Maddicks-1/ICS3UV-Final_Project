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
	White     = "♙"
	Black     = "♟"
	WhiteKing = "♔"
	BlackKing = "♚"
	Empty     = "□"
	Dark      = "■"
)

func main() {
	var board [size][size]string
	player := White

	// Set up board
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			if (r+c)%2 == 1 {
				switch {
				case r < 3:
					board[r][c] = White
				case r > 4:
					board[r][c] = Black
				default:
					board[r][c] = Empty
				}
			} else {
				board[r][c] = Dark
			}
		}
	}

	for {
		printBoard(board)

		// Check win
		if checkWin(board, White) {
			fmt.Println("White wins!")
			break
		}
		if checkWin(board, Black) {
			fmt.Println("Black wins!")
			break
		}

		fmt.Println("Player:", player)
		fmt.Print("Move (example: B 3 C 4 or QUIT): ")

		var fc string
		fmt.Scan(&fc)
		fc = strings.ToUpper(fc)

		if fc == "QUIT" {
			fmt.Println("Game quit.")
			break
		}

		var fr int
		var tc string
		var tr int
		fmt.Scan(&fr, &tc, &tr)

		tc = strings.ToUpper(tc)

		fr--
		tr--
		fcIndex := int(fc[0] - 'A')
		tcIndex := int(tc[0] - 'A')

		if validMove(board, player, fr, fcIndex, tr, tcIndex) {
			// Jump removal
			if abs(tr-fr) == 2 {
				board[(fr+tr)/2][(fcIndex+tcIndex)/2] = Empty
			}

			// Move piece
			board[tr][tcIndex] = board[fr][fcIndex]
			board[fr][fcIndex] = Empty

			// Kinging
			if board[tr][tcIndex] == White && tr == size-1 {
				board[tr][tcIndex] = WhiteKing
			}
			if board[tr][tcIndex] == Black && tr == 0 {
				board[tr][tcIndex] = BlackKing
			}

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

// Print the board
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

// Check if a move is valid
func validMove(b [size][size]string, p string, fr, fc, tr, tc int) bool {
	if fr < 0 || fr >= size || fc < 0 || fc >= size ||
		tr < 0 || tr >= size || tc < 0 || tc >= size {
		return false
	}

	piece := b[fr][fc]
	if !isPlayersPiece(piece, p) || b[tr][tc] != Empty {
		return false
	}

	rowDiff := tr - fr
	colDiff := tc - fc

	// Normal move
	if abs(rowDiff) == 1 && abs(colDiff) == 1 {
		return validDirection(piece, rowDiff)
	}

	// Jump move
	if abs(rowDiff) == 2 && abs(colDiff) == 2 {
		midR := (fr + tr) / 2
		midC := (fc + tc) / 2
		return validDirection(piece, rowDiff) && isOpponentPiece(piece, b[midR][midC])
	}

	return false
}

// Check if piece moves in correct direction
func validDirection(piece string, rowDiff int) bool {
	switch piece {
	case White:
		return rowDiff == 1 || rowDiff == 2
	case Black:
		return rowDiff == -1 || rowDiff == -2
	case WhiteKing, BlackKing:
		return abs(rowDiff) == 1 || abs(rowDiff) == 2
	}
	return false
}

// Check if piece belongs to current player
func isPlayersPiece(piece, player string) bool {
	return (player == White && (piece == White || piece == WhiteKing)) ||
		(player == Black && (piece == Black || piece == BlackKing))
}

// Check if piece belongs to opponent
func isOpponentPiece(piece, target string) bool {
	return (piece == White || piece == WhiteKing) && (target == Black || target == BlackKing) ||
		(piece == Black || piece == BlackKing) && (target == White || target == WhiteKing)
}

// Check if a player has no pieces left
func checkWin(b [size][size]string, player string) bool {
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			if isPlayersPiece(b[r][c], player) {
				return false
			}
		}
	}
	return true
}

// Absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
