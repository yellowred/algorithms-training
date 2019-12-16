// 5286. Shortest Path in a Grid + Obstacles Elimination
// hard

package main

import (
	"log"
	"math"
)

func main() {
	//log.Println(shortestPath([][]int{[]int{0, 0, 0}, []int{1, 1, 0}, []int{0, 0, 0}, []int{0, 1, 1}, []int{0, 0, 0}}, 1))
	//log.Println(shortestPath([][]int{[]int{0, 0, 0}, []int{1, 1, 0}, []int{0, 0, 0}}, 1))
	log.Println(shortestPath([][]int{[]int{0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1}, []int{0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 0, 1}, []int{1, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0}, []int{1, 0, 1, 1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1, 1}, []int{0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1, 1}, []int{0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 1, 0, 1, 1}, []int{1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 0}, []int{0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 1, 0, 0, 0}, []int{0, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0}, []int{0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 1, 0}}, 27))
	// log.Println(shortestPath(
	// 	[][]int{
	// 		[]int{0, 1, 0, 0},
	// 		[]int{0, 0, 0, 1},
	// 		[]int{0, 1, 1, 0},
	// 	}, 1))
}

// Lee Algorithm: http://users.eecs.northwestern.edu/~haizhou/357/lec6.pdf
func shortestPath(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[len(grid)-1])
	rowNum := []int{-1, 0, 0, 1}
	colNum := []int{0, -1, 1, 0}

	if grid[0][0] > 0 || grid[m-1][n-1] > 0 {
		return -1
	}

	visited := make([][][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			visited[i][j] = make([]int, k+1)
			for l := 0; l <= k; l++ {
				visited[i][j][l] = math.MaxInt32
			}
		}
	}
	visited[0][0][k] = 0
	// visited[m-1][n-1][k] = -1

	queue := [][]int{} //0 - x, 1 - y, 2 - k
	queue = append(queue, []int{0, 0, k})
	// log.Println(queue)
	for len(queue) > 0 {
		queueNode := queue[0]
		queue = queue[1:len(queue)]
		// log.Println("ITER", queueNode, visited[queueNode[0]][queueNode[1]][queueNode[2]])
		for i := 0; i < 4; i++ {
			x := queueNode[0]
			y := queueNode[1]
			row := queueNode[0] + rowNum[i]
			col := queueNode[1] + colNum[i]

			if row >= 0 && col >= 0 && row < m && col < n {
				nCost := queueNode[2] - grid[row][col]

				if nCost >= 0 && visited[row][col][nCost] > visited[x][y][queueNode[2]]+1 {
					visited[row][col][nCost] = visited[x][y][queueNode[2]] + 1
					queue = append(queue, []int{row, col, nCost})
				}
			}
		}
	}
	ans := -1
	for i := 0; i <= k; i++ {
		c := visited[m-1][n-1][i]
		if c != math.MaxInt32 {
			if ans == -1 || c < ans {
				ans = c
			}
		}
	}
	return ans
}
