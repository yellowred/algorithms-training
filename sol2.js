#!/usr/bin/env node

console.log(LookAndSay(1, 5));
function LookAndSay(start, n) {
	var ns = start
	for (i = 1; i <= n; i += 1) {
		ns = iterate(ns)
	}
	return ns
}

function iterate(s) {
	var s = s.toString()
	var ns = ''
	var curNum = null
	var lenNum = 0
	for (var i = 0; i < s.length; i += 1) {
		if (curNum != s[i]) {
			if (curNum != null) {
				ns = ns.concat(lenNum)
				ns = ns.concat(curNum);
			}
			curNum = s[i]
			lenNum = 1
		} else {
			lenNum += 1
		}
	}
	if (curNum != null) {
		ns = ns.concat(lenNum)
		ns = ns.concat(curNum);
	}
	return ns
}
