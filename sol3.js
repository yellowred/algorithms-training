#!/usr/bin/env node
console.log(convert(7));

function convert(n) {
	var dec = n
	var base = '0atlsin'
	var ret = '';
	while (dec > 0) {
		ret = base.charAt(dec % 7) + ret;
		dec = Math.floor(dec / 7);
	}
	return ret;
}

function base7(n) {
	var s = n.toString().split('').reverse();
	var arr = []
	var carry = 0
	var m = 0
	console.log(s);
	for (i = 0; i < s.length; i += 1) {
		arr[i] = s[i] % 7 + carry
		carry = Math.floor(s[i] / 7)
	}
}

function base(dec, base){
	var len=base.length;
	var ret='';
	while(dec>0){
		ret = base.charAt(dec%len) + ret;
		dec = Math.floor(dec/len);
	}
	return ret;
}
