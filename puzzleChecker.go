package eightPuzzleSolver

import (
	"errors"
)

func isValidBoard(board *[3][3]int) bool {
	checker := [9]int{0,0,0,0,0,0,0,0,0};
	flattenBoard := flattenBoard(board)

	for _,val := range flattenBoard {
		if val >=0 && val <= 8 {
			checker[val] = 1
		}
	}

	for _,val := range checker {
		if val != 1 {
			return false
		}
	}
	return true
}

func flattenBoard(board *[3][3]int) []int {
	flattenBoard := make([]int,0)
	for i:=0;i<3;i++ {
		for j:=0;j<3;j++ {
			flattenBoard = append(flattenBoard,(*board)[i][j])
		}
	}
	return flattenBoard;
}

func getInversionCount(board *[3][3]int) int {
	count := 0;
	flattenBoard := flattenBoard(board)
	for i:=0;i<8;i++ {
		for j:=i+1;j<9;j++ {
			if (flattenBoard[i] > flattenBoard[j]) && (flattenBoard[i] != 0) && (flattenBoard[j] != 0) {
				count++;
			}
		}
	}
	return count;
}

func isSolvable(board *[3][3]int) error {
	if isValidBoard(board) == false {
		return errors.New("invalid board")
	}
	if getInversionCount(board) % 2 == 1 {
		return errors.New("this board is not solvable")
	}
	return nil
} 