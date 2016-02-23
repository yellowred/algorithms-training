#!/usr/bin/env node

function processData(input) {
    //Enter your code here
//	var t = new BigNumber(12)
//	console.log(t.pow2().add(16).toString());
	var aa = new BigNumber(12)
	console.log(aa.pow2().toString());
/*
	var arr = input.split(' ').map(Number)
	var A = arr[0]
	var B = arr[1]
	var N = arr[2]
	function fibo(t1, t2, n)
	{
		if (n == 1) {
			return t1
		} else if (n == 2) {
			return t2
		} else {
			return fibo(t1, t2, n - 1).pow2().add(fibo(t1, t2, n - 2))
		}
	}

	console.log(fibo(new BigNumber(A), new BigNumber(B), N).toString())
*/
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

BigNumber = function(a){
    this.arr = [] //Array.apply(null, Array(100)).map(Number.prototype.valueOf, 0)
	var m = 0
	var i = 0
	while(a > 0) {
		m = a % 10
		this.arr[i] = m
		a = Math.floor(a / 10)
		i += 1
	}
};
BigNumber.prototype.pow2 = function() {
	var carry = 0
	var m = 0
	var i = 0
	var e = 0
	var tempArr = []
	console.log('pow', this.arr)
	for (var j = 0; j < this.arr.length; j += 1) {
		e = this.arr[j]
		i = 0
		m = 0
		carry = 0
		while (i < this.arr.length || carry > 0) {
			if (typeof this.arr[i] !== 'undefined') {
				m = this.arr[i] * e
			} else {
				m = 0
			}
			m += carry
			tempArr[i] = m % 10
			carry = Math.floor(m / 10)
			console.log(i, tempArr, m, carry);
			i += 1
		}
		e = tempArr.reverse().join('')
	}
	console.log('pow', this.arr)
	return this
}
BigNumber.prototype.add = function(arrB) {
	var s1 = s2 = 0
	var carry = 0
	var m = 0
	var i = 0

	while (i < this.arr.length || i < arrB.length || carry > 0) {
		s1 = 0
		s2 = 0
		if (typeof this.arr[i] !== 'undefined') {
			s1 = this.arr[i]
		}
		if (typeof arrB[i] !== 'undefined') {
			s2 = arrB[i]
		}
		m = s1 + s2 + carry
		this.arr[i] = m % 10
		carry = Math.floor(m / 10)
		console.log(i, this.arr, m, carry);
		i += 1
	}
	return this
}
BigNumber.prototype.toString = function() {
	return this.arr.reverse().join('')
}
