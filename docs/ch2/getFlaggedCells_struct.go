package ch2

type Cell struct{}

func (c *Cell) IsFlagged() bool {
    return true
}

func getFlaggedCells_struct(gameBoard []Cell) []Cell {
    var flaggedCells []Cell
    for _, cell := range gameBoard {
        if cell.IsFlagged() {
            flaggedCells = append(flaggedCells, cell)
        }
    }
    return flaggedCells
}
