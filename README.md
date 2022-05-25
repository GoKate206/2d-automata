### Requirements
- Given a 2D Array to reflect a 10x10 grid of cells
- Cells may be alive or empty. What stage of life is reflected by a number
- Cells "age" according to rules similar to Conway's Game of Life
- Assuming the first grid is the first generation, what would generation 20 look like?

#### How to run this program
- Build the project and dependencies by running<br>
`go mod init src/github.com/GoKate206`<br>
( Make sure there is no leading slash at the end of `GoKate206`)
- Navigate to `src/github.com/GoKate206`
  - To run tests `go test *.go -v`
  - To run the program `go run main.go`
- Technology Used
  - [Testify](https://pkg.go.dev/github.com/stretchr/testify): Provide tools for testing
