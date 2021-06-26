package main

import "fmt"

type NeighborCount struct {
  adult int
  total int
}

type NeighborIndex struct {
  row int
  cell int
}

const (
  empty = 0
  newborn = 1
  adult = 2
  senior = 3
  gridLength = 10 // TODO: have this be customizable
)

type Grid [][]int

func main() {
  givenIterations := 20
  iteration := 1
  firstGeneration := Grid{
    {0,0,0,0,0,0,0,0,0,0},
    {0,0,1,1,0,0,0,0,0,0},
    {0,0,0,0,2,0,0,0,0,0},
    {0,0,0,1,2,0,0,0,0,0},
    {0,0,0,0,0,0,0,0,0,0},
    {0,0,1,0,0,0,0,0,0,0},
    {0,2,1,0,0,0,0,0,0,0},
    {0,2,0,0,0,0,0,0,0,0},
    {0,0,0,0,0,0,0,0,0,0},
    {0,0,0,0,0,0,0,0,0,0},
  }
  generation := populateNewGeneration(firstGeneration)
  for iteration < givenIterations {
    generation = populateNewGeneration(generation)
    iteration++
  }
  fmt.Printf("\n\n 20th Generation is: %v \n\n", generation)

}

func createEmptyGeneration() Grid {
  grid := make(Grid, gridLength)
  for i := range grid {
    grid[i] = make([]int, gridLength)
  }

  return grid
}

func populateNewGeneration(grid Grid) Grid {
  newGeneration := createEmptyGeneration()
  for rowIndex := range grid {
    row := grid[rowIndex]
    for cellIndex := range row {
      cell := row[cellIndex]
      count := getNeighborCount(grid, rowIndex, cellIndex)
      newGeneration[rowIndex][cellIndex] = getNewValue(cell, count)
    }
  }
  return newGeneration
}

func getNeighborCount(grid Grid, row, cell int) NeighborCount {
  indexes := []NeighborIndex{
    {
      row: row + 1,
      cell: cell,
    },
    {
      row: row + 1,
      cell: cell + 1,
    },
    {
      row: row,
      cell: cell + 1,
    },
    {
      row: row - 1,
      cell: cell + 1,
    },
    {
      row: row - 1,
      cell: cell,
    },
    {
      row: row - 1,
      cell: cell -1,
    },
    {
      row: row,
      cell: cell -1,
    },
    {
      row: row + 1,
      cell: cell - 1,
    },
}
  total := 0
  adultCount := 0
  for _, nIndex := range indexes {
    if nIndex.row > 9 || nIndex.row < 0 {
      continue
    }

    if nIndex.cell > 9 || nIndex.cell < 0 {
      continue
    }

    oldCell := grid[nIndex.row][nIndex.cell]
    if oldCell != empty {
      total++
      if (oldCell == adult) {
        adultCount++
      }
    }
  }

  return NeighborCount{
    adult: adultCount,
    total: total,
  }
}

func getNewValue(cell int, neighborCount NeighborCount) int {
  switch cell {
  case empty:
    if neighborCount.adult == 2 {
      return newborn
    } else {
      return empty
    }
  case newborn:
    // TODO: handle comment diffs
    if neighborCount.total <= 1 || neighborCount.total >= 5 {
      return empty;
    } else {
      return adult;
    }
  case adult:
    // TODO: handle comment differences
    if neighborCount.total == 0 || neighborCount.total >= 3 {
      return empty;
    } else {
      return senior
    }
  default:
    return empty
  }
}
