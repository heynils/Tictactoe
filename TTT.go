package main

import "fmt"

var board [3][3]string      //2D array that holds the state of the game
var currentPlayerMark = "X" //'X' or 'O' char representing which player's turn it is

// Set/Reset the board back to all empty values
func initializeBoard() {
	//Loop through rows
	for i := 0; i < 3; i++ {
		//Loop through columns
		for j := 0; j < 3; j++ {
			board[i][j] = "-"
		}
	}
}

// Present the board with its current status
func printBoard() {
	println("-------------")
	for i := 0; i < 3; i++ {
		print("| ")
		for j := 0; j < 3; j++ {
			print(board[i][j] + " | ")
		}
		println()
		println("-------------")
	}
}

// Returns true if board is full, otherwise false
func isBoardFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "-" {
				return false
			}
		}
	}
	return true
}

// Checks if any player has won in any way
func checkForWin() bool {
	return (checkRowsForWin() || checkColumnsForWin() || checkDiagonalsForWin())
}

// Checks if any player has won with rows
func checkRowsForWin() bool {
	for i := 0; i < 3; i++ {
		if checkRowCol(board[i][0], board[i][1], board[i][2]) == true {
			return true
		}
	}
	return false
}

// Checks if any player has won with columns
func checkColumnsForWin() bool {
	for i := 0; i < 3; i++ {
		if checkRowCol(board[0][i], board[1][i], board[2][i]) == true {
			return true
		}
	}
	return false
}

// Checks if any player has won with diagonals
func checkDiagonalsForWin() bool {
	return (checkRowCol(board[0][0], board[1][1], board[2][2]) ||
		checkRowCol(board[2][0], board[1][1], board[0][2]))
}

// Compares three values (input) and returns true if they are all the same (excluding '-')
func checkRowCol(c1, c2, c3 string) bool {
	return (c1 != "-" && c1 == c2 && c2 == c3)
}

// Changes the current player mark
func changePlayer() {
	if currentPlayerMark == "X" {
		currentPlayerMark = "O"
	} else {
		currentPlayerMark = "X"
	}
}

// Places a mark at the cell specified by row and col with the mark of the current player
func placeMark(row, col int) bool {
	// Make sure that row and col are in bounds of the board
	if row >= 0 && row < 3 {
		if col >= 0 && col < 3 {
			if board[row][col] == "-" {
				board[row][col] = currentPlayerMark
				return true
			}
		}
	}
	return false
}

func main() {

	initializeBoard()
	println("Tic-Tac-Toe!")

	for {
		println("Current board layout:")
		printBoard()
		for {
			println("Player " + currentPlayerMark + ", enter an empty row and column to place your mark!")
			var row int
			fmt.Scan(&row)
			var col int
			fmt.Scan(&col)
			row--
			col--
			if placeMark(row, col) {
				break
			}
		}
		changePlayer()
		if checkForWin() || isBoardFull() {
			break
		}

	}
	if isBoardFull() && !checkForWin() {
		printBoard()
		println("The game was a tie!")
	} else {
		println("Current board layout:")
		printBoard()
		changePlayer()
		println(currentPlayerMark + " wins!")
	}
}
