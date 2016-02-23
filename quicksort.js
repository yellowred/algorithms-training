#!/usr/bin/env node

function processData(input) {
	//Enter your code here
	var lines = input.split(/\r?\n/)
	var size = lines[0]
	var arr = lines[1].split(' ').map(Number)

	console.log(divideAndConqure(arr).join(' '));
}

function divideAndConqure(arr) {
	var p = arr[0]
	var left = []
	var right = []
	var equal = []
	for (var i = 0; i < arr.length; i += 1) {
		if (arr[i] < p) {
			left.push(arr[i])
		} else if (arr[i] > p) {
			right.push(arr[i])
		} else {
			equal.push(arr[i])
		}
	}
	if (left.length > 1) {
		left = divideAndConqure(left)
		console.log(left.join(' '));
	} else {

	}
	if (right.length > 1) {
		right = divideAndConqure(right)
		console.log(right.join(' '));
	} else {

	}
	return left.concat(equal.concat(right))
}


process.stdin.resume();
process.stdin.setEncoding("ascii");
_input = "";
process.stdin.on("data", function (input) {
	_input += input;
});

process.stdin.on("end", function () {
	processData(_input);
});
