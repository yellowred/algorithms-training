#!/usr/bin/env node

console.log(compute("PMLPMMMLPMLPMML"));
function compute(instructions) {
	var s = instructions
	var pos = 0
	var has_block = false
	var arr = Array.apply(null, Array(10)).map(Number.prototype.valueOf,0);
	for (i = 0; i <= s.length; i += 1) {
		if (s[i] == 'P') {
			pos = 0
			has_block = true
		} else if (s[i] == 'M') {
			if (pos < 9) {
				pos += 1
			}
		} else if (s[i] == 'L' && has_block) {
			if (arr[pos] < 15)
			{
				arr[pos] += 1
				has_block = false
			}
		}
	}
	return arr.join('')
}
