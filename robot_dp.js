// https://leetcode.com/problems/unique-paths/

/**
 * @param {number} m
 * @param {number} n
 * @return {number}
 */
var uniquePaths = function(m, n) {
    let memo = {cells:Array(m).fill(0).map(()=> Array(n).fill(0)), queue:[]}
    memo.cells[m-1][n-1] = 1
    console.log(memo)
    memo.queue.push([m-1,n-1])
    path(memo)
    console.log(memo)
    return memo.cells[0][0]
};

/**
 * @param {object} total
 */
var path = (memo) => {
    let counter = 0
    while (memo.queue.length>0) {
        let cell = memo.queue.shift()
        console.log("C", cell, memo.cells)

        let x = cell[0]
        let y = cell[1]

        let xd = x-1
        let yd = y-1

        if (xd>=0 && memo.cells[xd][y] === 0) {
            memo.cells[xd][y] = memo.cells[x][y] + (memo.cells[xd][y+1] ? memo.cells[xd][y+1] : 0)
            console.log("M", xd, y, memo.cells[xd][y], memo.cells[x][y])
            memo.queue.push([xd,y])
        }

        if (yd>=0 && memo.cells[x][yd] === 0) {
            memo.cells[x][yd] = memo.cells[x][y] + (memo.cells[x+1] ? memo.cells[x+1][yd] : 0)
            console.log("M", x, yd, memo.cells[x][yd], memo.cells[x][y])
            memo.queue.push([x,yd])
        }
        counter++
    }
    console.log("counter=", counter)
}

console.log(uniquePaths(7,3))
