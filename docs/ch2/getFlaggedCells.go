package ch2

const (
    STATUS_VALUE = 0
    FLAGGED      = 4
)

func getFlaggedCells(gameBoard [][]int) [][]int {
    flaggedCells := make([][]int, 0)
    for _, cell := range gameBoard {
        if cell[STATUS_VALUE] == FLAGGED {
            flaggedCells = append(flaggedCells, cell)
        }
    }
    return flaggedCells
}
