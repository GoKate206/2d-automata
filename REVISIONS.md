### Refactors I would do
- Allow for dynamic grids ( I used a variable for grid length )
- I would use a mock database ( or real database for production ) to store / query / update the value of a cell for a grid
  - This would particular help because I wouldn't have to have a reference slice to store the row / cell indices to check a cell's neighbor.
  - I could have also used the comments for why a cell's age changes
- I would have liked to build a frontend for this project that could show the changes between generations
