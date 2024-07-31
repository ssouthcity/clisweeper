package minesweeper_test

import (
	"testing"

	"github.com/ssouthcity/clisweeper/pkg/minesweeper"
)

func TestDefaultCellValues(t *testing.T) {
	cell := minesweeper.Cell{}

	if cell.State() != minesweeper.Hidden {
		t.Errorf("Expected newly created cell to be hidden, got '%v'", cell.State())
	}
}

func TestRevealCell(t *testing.T) {
	cell := minesweeper.Cell{}
	cell.Uncover()

	if cell.State() != minesweeper.Revealed {
		t.Errorf("Expected cell to be revealed after uncovering, got '%v'", cell.State())
	}
}

func TestRevealCellMultipleTimes(t *testing.T) {
	cell := minesweeper.Cell{}

	if err := cell.Uncover(); err != nil {
		t.Errorf("Expected uncovering to succeed, got '%v'", err)
	}

	if err := cell.Uncover(); err != minesweeper.ErrCellAlreadyRevealed {
		t.Errorf("Expected subsequent reveals of cell to return ErrCellAlreadyRevealed, got '%v'", err)
	}
}

func TestToggleFlaggingCell(t *testing.T) {
	cell := minesweeper.Cell{}

	if err := cell.ToggleFlag(); err != nil {
		t.Errorf("Expected flagging of new cell to succeed, got '%v'", err)
	}

	if cell.State() != minesweeper.Flagged {
		t.Errorf("Expected cell to be flagged after calling ToggleFlag() once, got '%v'", cell.State())
	}

	if err := cell.ToggleFlag(); err != nil {
		t.Errorf("Expected toggling from flagged.State() to succeed, got '%v'", err)
	}

	if cell.State() == minesweeper.Flagged {
		t.Errorf("Expected cell to be unflagged after calling ToggleFlag() twice, got '%v'", cell.State())
	}
}

func TestUncoveringFlaggedCell(t *testing.T) {
	cell := minesweeper.Cell{}
	cell.ToggleFlag()

	err := cell.Uncover()
	if err != minesweeper.ErrUncoveringFlaggedCell {
		t.Errorf("Expected uncovering of flagged cell to return ErrUncoveringFlaggedCell, got '%v'", err)
	}
}

func TestFlaggingRevealedCell(t *testing.T) {
	cell := minesweeper.Cell{}
	cell.Uncover()

	err := cell.ToggleFlag()
	if err != minesweeper.ErrCannotFlagRevealedCell {
		t.Errorf("Expected uncovering of cell to return ErrCannotFlagRevealedCell, got '%v'", err)
	}
}
