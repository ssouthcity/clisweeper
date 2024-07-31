package minesweeper

import "errors"

var (
	ErrCellAlreadyRevealed    = errors.New("Cell has already been revealed")
	ErrUncoveringFlaggedCell  = errors.New("A flagged cell cannot be uncovered")
	ErrCannotFlagRevealedCell = errors.New("A cell that has been revealed cannot be flagged afterwards")
)

type CellState int

const (
	Hidden CellState = iota
	Revealed
	Flagged
)

type Cell struct {
	state CellState
}

func (c *Cell) State() CellState {
	return c.state
}

func (c *Cell) Uncover() error {
	if c.state == Revealed {
		return ErrCellAlreadyRevealed
	}

	if c.state == Flagged {
		return ErrUncoveringFlaggedCell
	}

	c.state = Revealed

	return nil
}

func (c *Cell) ToggleFlag() error {
	if c.state == Revealed {
		return ErrCannotFlagRevealedCell
	}

	if c.state == Flagged {
		c.state = Hidden
		return nil
	}

	if c.state == Hidden {
		c.state = Flagged
		return nil
	}

	return nil
}
