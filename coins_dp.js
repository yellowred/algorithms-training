/**
 * @param {number[]} coins
 * @param {number} amount
 * @return {number}
 */
var coinChange = function(coins, amount) {
    coins.sort((a,b) => a>b)
    return coinSet(coins, amount)
};

var coinSet = (coins, amount) => {
    let newCoins = [...coins]
    let biggest = newCoins.pop()
    let total = 0
    let count = 0
    while (total + biggest <= amount) {
        total += biggest
        count++
    }
    if (amount == total) {
      console.log("S1", biggest, count)
        return count
    }
    if (count === 0 && newCoins.length === 0) {
        return -1
    }
    let remainder = -1
    while (remainder < 0) {
        remainder = coinSet(newCoins, amount - total)    
        if (remainder < 0 && count>0) {
            total -= biggest
            count--
        } else if (remainder<0 && count==0) {
            return -1
        }
    }
    console.log("S", biggest, count)
    return count + remainder
}

// 83*13=1079
// 186
// 408*4=1632
// 419*8=3352

console.log(coinChange([186,419,83,408], 6249))