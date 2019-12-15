// 5286. Shortest Path in a Grid + Obstacles Elimination
// hard

package main

import (
	"log"
)

func main() {
	//log.Println(shortestPath([][]int{[]int{0, 0, 0}, []int{1, 1, 0}, []int{0, 0, 0}, []int{0, 1, 1}, []int{0, 0, 0}}, 1))
	//log.Println(shortestPath([][]int{[]int{0, 0, 0}, []int{1, 1, 0}, []int{0, 0, 0}}, 1))
	log.Println(shortestPath([][]int{[]int{0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1}, []int{0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 0, 1}, []int{1, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0}, []int{1, 0, 1, 1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1, 1}, []int{0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1, 1}, []int{0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 1, 0, 1, 1}, []int{1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 0}, []int{0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 1, 0, 0, 0}, []int{0, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0}, []int{0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 1, 0}}, 27))
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

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	visited[0][0] = true

	queue := [][]int{} //0 - x, 1 - y, 2 - dst, 3 - k
	queue = append(queue, []int{0, 0, 0, 0})

	for len(queue) > 0 {
		queueNode := queue[0]

		// If we have reached the destination cell,
		// we are done
		if queueNode[0] == m-1 && queueNode[1] == n-1 {
			return queueNode[2]
		}

		// Otherwise dequeue the front cell in the queue
		// and enqueue its adjacent cells
		queue = queue[1:len(queue)]

		for i := 0; i < 4; i++ {
			row := queueNode[0] + rowNum[i]
			col := queueNode[1] + colNum[i]

			// if adjacent cell is valid, has path and
			// not visited yet, enqueue it.
			if row >= 0 && col >= 0 && row < m && col < n && !visited[row][col] {
				if grid[row][col] == 0 {
					// mark cell as visited and enqueue it
					visited[row][col] = true
					queue = append(queue, []int{row, col, queueNode[2] + 1, queueNode[3]})
				} else if queueNode[3] < k {
					// mark cell as visited and enqueue it
					visited[row][col] = true
					queue = append(queue, []int{row, col, queueNode[2] + 1, queueNode[3] + 1})
				}
			}
		}
	}
	return -1
}
