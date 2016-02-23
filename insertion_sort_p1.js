#!/usr/bin/env node

function processData(input) {
    //Enter your code here
	var lines = input.split(/\r?\n/)
	var size = lines[0]
	var arr = lines[1].split(' ').map(Number)
	for (var j = 2; j<= size; j += 1) {
		var i = j - 2;
		var e = arr[j - 1]
		while (e != null) {
			if (i >= 0 && arr[i] > e)
			{
				arr[i+1] = arr[i]
			} else {
				arr[i+1] = e
				e = null
			}
			i -= 1
		}
		console.log(arr.join(' '));
	}

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
