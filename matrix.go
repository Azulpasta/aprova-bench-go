package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Matrix represents a two-dimensional grid of integers.
type Matrix [][]int

func New(s string) (Matrix, error) {
	lines := strings.Split(s, "\n")
	m := make(Matrix, len(lines))
	width := -1
	for i, line := range lines {
		fields := strings.Fields(line)
		if width == -1 {
			width = len(fields)
		}
		if len(fields) != width {
			return nil, fmt.Errorf("row %d has %d columns, want %d", i, len(fields), width)
		}
		row := make([]int, len(fields))
		for j, f := range fields {
			v, err := strconv.Atoi(f)
			if err != nil {
				return nil, fmt.Errorf("row %d, col %d: %w", i, j, err)
			}
			row[j] = v
		}
		m[i] = row
	}
	return m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	if len(m) == 0 {
		return nil
	}
	cols := make([][]int, len(m[0]))
	for j := range cols {
		col := make([]int, len(m))
		for i := range m {
			col[i] = m[i][j]
		}
		cols[j] = col
	}
	return cols
}

func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m))
	for i, r := range m {
		rows[i] = append([]int(nil), r...)
	}
	return rows
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m) {
		return false
	}
	if col < 0 || col >= len(m[row]) {
		return false
	}
	m[row][col] = val
	return true
}
