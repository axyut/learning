// Simulation of card movement on a 2D grid
// Description
// Given an application that organizes information by arranging cards with notes on them into lists on a two-dimensional grid. These notes can be moved across multiple lists.

// Examples include billboard-based task management applications such as Trello or Jira.

// Image is bellow.

// The first argument, cards, is given a value indicating the initial position of all cards in the board, each as a two-dimensional array, with the following rules
// [CardID, RowIndex, ColumnIndex]
// Here, the row and column indexes are assumed to be zero-based-indexed.

// The second argument, moves, is a two-dimensional array of values indicating the movement of the cards, each with the following rules
// [CardID, Row index before move, Column index before move, Destination row index, Destination column index]
// Here, the row and column indexes are zero-based-indexed and valid values (you don't have to assume that every index is not valid).

// The third argument, query, is an int type value indicating the ID of the card to be moved.
// Note that there are cases where moving one card will cause other cards to be moved.

// Based on the given cards and moves, implement a function that returns the final position of the card specified in the third argument, query, as [row index, column index].

// Note that the following rules are assumed to exist for moving cards.

// When a card is moved and if there is a card at the destination, the existing card is moved down one row.
// Also, when a card is moved, other cards in the original column move up one.
// Cards can be moved to the same row and column as before.
// There shall be no gaps in the initial placement.
// We can assume that the coordinates before a move is correct.
// Examples
// Example 1
// Input: cards=[[1,1,0],[3,0,0],[6,0,1],[4,0,2],[5,2,0],[7,1,1],[2,1,2]], moves=[[6,0,1,2,0]], query=6
// Output: [2,0]
// Description: This is a case where the target moves. In this case, the card with card ID 6 moves to [2,0], so [2,0] is the correct value.
// Example 2
// Input: cards=[[1,1,0],[3,0,0],[6,0,1],[4,0,2],[5,2,0],[7,1,1],[2,1,2]], moves=[[6,0,1,2,0],[7,0,1,0,2]
// ], query=2
// Output: [2,2]
// Description: This is a case where a card is moved to the same column as the target. In this case, the target is not moved directly, but is moved down one row as the other card (card ID 7) is moved, so [2,2] is the correct value.
// Example 3
// Input: cards=[[1,1,0],[3,0,0],[6,0,1],[4,0,2],[5,2,0],[7,1,1],[2,1,2]], moves=[[6,0,1,3,0],[7,0,1,0,2]
// ], query=5
// Output: [2,0]
// Description: This is a case where the card moves to the same column as the target, but the target does not move because it is placed below it.
// Constraints
// 0 <= len(cards) <= 10^5
// 0 <= len(moves) <= 10^5

package main

import (
	"fmt"
)

func finalCardPosition(cards [][]int, moves [][]int, query int) []int {
	// Map to store current position of each card
	cardPositions := make(map[int][2]int)

	// Initialize card positions based on initial cards setup
	for _, card := range cards {
		cardID, row, col := card[0], card[1], card[2]
		cardPositions[cardID] = [2]int{row, col}
	}

	// Process each move
	for _, move := range moves {
		cardID := move[0]
		fromRow, fromCol := move[1], move[2]
		toRow, toCol := move[3], move[4]

		// Move other cards in the original column up one row
		for otherCardID, pos := range cardPositions {
			otherRow, otherCol := pos[0], pos[1]
			if otherCol == fromCol && otherRow < fromRow {
				cardPositions[otherCardID] = [2]int{otherRow + 1, otherCol}
			}
		}

		// Move any card at the destination down one row
		for otherCardID, pos := range cardPositions {
			otherRow, otherCol := pos[0], pos[1]
			if otherRow == toRow && otherCol == toCol && otherCardID != cardID {
				cardPositions[otherCardID] = [2]int{otherRow + 1, otherCol}
			}
		}

		// Update the moving card's position to its destination
		cardPositions[cardID] = [2]int{toRow, toCol}
	}

	// Return the final position of the queried card
	return []int{cardPositions[query][0], cardPositions[query][1]}
}

func CardMovementExecute() {
	cards := [][]int{{1, 1, 0}, {3, 0, 0}, {6, 0, 1}, {4, 0, 2}, {5, 2, 0}, {7, 1, 1}, {2, 1, 2}}
	moves := [][]int{{6, 0, 1, 2, 0}}
	query := 6
	fmt.Println(finalCardPosition(cards, moves, query)) // Output: [2 0]

	moves = [][]int{{6, 0, 1, 2, 0}, {7, 0, 1, 0, 2}}
	query = 2
	fmt.Println(finalCardPosition(cards, moves, query)) // Output: [2 2]

	moves = [][]int{{6, 0, 1, 3, 0}, {7, 0, 1, 0, 2}}
	query = 5
	fmt.Println(finalCardPosition(cards, moves, query)) // Output: [2 0]
}
