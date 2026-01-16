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

const boardSize = 8

const (
	White     = "♙"
	Black     = "♟"
	whiteKing = "♔"
	blackKing = "♚"
	Empty     = "□"
	Dark      = "■"
)

func main() {
	var board [boardSize][boardSize]string
	currentPlayer := White

	setupBoard(&board)

	for {
		printBoard(board)

		if hasWon(board, White) {
			fmt.Println("White wins!")
			return
		}
		if hasWon(board, Black) {
			fmt.Println("Black wins!")
			return
		}

		fmt.Println("Player:", currentPlayer)
		fmt.Print("Move (example: B 3 C 4 or QUIT): ")

		var fromCol string
		fmt.Scan(&fromCol)
		fromCol = strings.ToUpper(fromCol)

		if fromCol == "QUIT" {
			fmt.Println("Game quit.")
			return
		}

		var fromRow int
		var toCol string
		var toRow int
		fmt.Scan(&fromRow, &toCol, &toRow)

		fromRow--
		toRow--
		fromC := int(fromCol[0] - 'A')
		toC := int(strings.ToUpper(toCol)[0] - 'A')

		if isValidMove(board, currentPlayer, fromRow, fromC, toRow, toC) {
			makeMove(&board, fromRow, fromC, toRow, toC)
			currentPlayer = switchPlayer(currentPlayer)
		} else {
			fmt.Println("Invalid move")
		}
	}
}

func setupBoard(board *[boardSize][boardSize]string) {
	for r := 0; r < boardSize; r++ {
		for c := 0; c < boardSize; c++ {
			if (r+c)%2 == 0 {
				board[r][c] = Dark
			} else if r < 3 {
				board[r][c] = White
			} else if r > 4 {
				board[r][c] = Black
			} else {
				board[r][c] = Empty
			}
		}
	}
}

func printBoard(board [boardSize][boardSize]string) {
	fmt.Println("\n  A B C D E F G H")
	for r := 0; r < boardSize; r++ {
		fmt.Print(r+1, " ")
		for c := 0; c < boardSize; c++ {
			fmt.Print(board[r][c], " ")
		}
		fmt.Println()
	}
}

func isValidMove(board [boardSize][boardSize]string, player string, fr, fc, tr, tc int) bool {
	if !inBounds(fr, fc) || !inBounds(tr, tc) {
		return false
	}

	piece := board[fr][fc]
	if !isPlayersPiece(piece, player) || board[tr][tc] != Empty {
		return false
	}

	rowDiff := tr - fr
	colDiff := tc - fc

	if abs(rowDiff) == 1 && abs(colDiff) == 1 {
		return correctDirection(piece, rowDiff)
	}

	if abs(rowDiff) == 2 && abs(colDiff) == 2 {
		midR := (fr + tr) / 2
		midC := (fc + tc) / 2
		return correctDirection(piece, rowDiff) &&
			isOpponent(piece, board[midR][midC])
	}

	return false
}

func makeMove(board *[boardSize][boardSize]string, fr, fc, tr, tc int) {
	if abs(tr-fr) == 2 {
		board[(fr+tr)/2][(fc+tc)/2] = Empty
	}

	board[tr][tc] = board[fr][fc]
	board[fr][fc] = Empty

	if board[tr][tc] == White && tr == boardSize-1 {
		board[tr][tc] = whiteKing
	}
	if board[tr][tc] == Black && tr == 0 {
		board[tr][tc] = blackKing
	}
}

func correctDirection(piece string, rowDiff int) bool {
	if piece == White {
		return rowDiff > 0
	}
	if piece == Black {
		return rowDiff < 0
	}
	return abs(rowDiff) <= 2
}

func isPlayersPiece(piece, player string) bool {
	if player == White {
		return piece == White || piece == whiteKing
	}
	return piece == Black || piece == blackKing
}

func isOpponent(piece, target string) bool {
	return isPlayersPiece(piece, White) && isPlayersPiece(target, Black) ||
		isPlayersPiece(piece, Black) && isPlayersPiece(target, White)
}

func hasWon(board [boardSize][boardSize]string, player string) bool {
	for r := 0; r < boardSize; r++ {
		for c := 0; c < boardSize; c++ {
			if isPlayersPiece(board[r][c], player) {
				return false
			}
		}
	}
	return true
}

func switchPlayer(player string) string {
	if player == White {
		return Black
	}
	return White
}

func inBounds(r, c int) bool {
	return r >= 0 && r < boardSize && c >= 0 && c < boardSize
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
