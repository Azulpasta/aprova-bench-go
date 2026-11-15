package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix struct {
	data [][]int
}

func New(s string) (Matrix, error) {
	lines := strings.Split(s, "\n")
	var data [][]int
	var expectedCols int

	for i, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			return Matrix{}, fmt.Errorf("empty row at index %d", i)
		}

		parts := strings.Fields(line)
		if i == 0 {
			expectedCols = len(parts)
		}

		if len(parts) != expectedCols {
			return Matrix{}, fmt.Errorf("uneven row: expected %d columns, got %d", expectedCols, len(parts))
		}

		row := make([]int, len(parts))
		for j, part := range parts {
			num, err := strconv.ParseInt(part, 10, 64)
			if err != nil {
				return Matrix{}, err
			}
			row[j] = int(num)
		}
		data = append(data, row)
	}

	return Matrix{data: data}, nil
}

func (m Matrix) Rows() [][]int {
	if len(m.data) == 0 {
		return [][]int{}
	}

	result := make([][]int, len(m.data))
	for i, row := range m.data {
		result[i] = make([]int, len(row))
		copy(result[i], row)
	}
	return result
}

func (m Matrix) Cols() [][]int {
	if len(m.data) == 0 {
		return [][]int{}
	}

	numRows := len(m.data)
	numCols := len(m.data[0])

	result := make([][]int, numCols)
	for c := 0; c < numCols; c++ {
		result[c] = make([]int, numRows)
		for r := 0; r < numRows; r++ {
			result[c][r] = m.data[r][c]
		}
	}
	return result
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m.data) {
		return false
	}
	if col < 0 || col >= len(m.data[row]) {
		return false
	}
	m.data[row][col] = val
	return true
}
