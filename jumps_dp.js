// Jump Game
// https://leetcode.com/problems/jump-game/

/**
 * @param {number[]} nums
 * @return {boolean}
 */
const MEMO_GOOD = 1
const MEMO_BAD = 2
const MEMO_UNKNOWN = 3

var canJump = function(nums) {
    let memo = Array(nums.length).fill(MEMO_UNKNOWN)
    memo[memo.length-1] = MEMO_GOOD
    for (let i = memo.length-1; i>=0; i--) {
        let furthestJumpPosition = Math.min(i + nums[i], nums.length-1)
        for (let j = i + 1; j<=furthestJumpPosition; j++) {
            if (memo[j] === MEMO_GOOD) {
                memo[i] = MEMO_GOOD
                break
            }
        }
        if (memo[i] != MEMO_GOOD) {
            memo[i] = MEMO_BAD
        }
    }
    return memo[0] === MEMO_GOOD
};

console.log("RES:", canJump([2,4,2,1,0,2,0]))