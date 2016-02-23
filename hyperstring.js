#!/usr/bin/env node

function processData(input) {
	//Enter your code here
	var lines = input.split(/\r?\n/)
	var arr = lines[0].split(' ').map(Number)
	var N = arr[0]
	var M = arr[1]
	var i = 1
	var ss = []
	while(i <= N) {
		ss.push(lines[i])
		i += 1
	}
	var res = []
	var hs = '';
	i = 0
	var solutions = 1
	while(i >= 0) {
		if (res[i] == undefined) {
			res[i] = 0
		}
		hs = res.slice(0, i+1).map(function(a){return ss[a]}).join('')
		if (hs.length <= M) {
			solutions += 1
			i += 1
		} else {
			do {
				delete res[i]
				i -= 1
				res[i] += 1
			} while (res[i] >= N && i >= 0)
		}
	}

	//#2 929267708
	//#3 167388485
	console.log(solutions % 1000000007)
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
