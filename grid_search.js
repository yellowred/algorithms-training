#!/usr/bin/env node

process.stdin.resume();
process.stdin.setEncoding('ascii');

var input_stdin = "";
var input_stdin_array = "";
var input_currentline = 0;

process.stdin.on('data', function (data) {
    input_stdin += data;
});

process.stdin.on('end', function () {
    input_stdin_array = input_stdin.split("\n");
    main();
});

function readLine() {
    return input_stdin_array[input_currentline++];
}

/////////////// ignore above this line ////////////////////

function main() {
    var t = parseInt(readLine());
    for(var a0 = 0; a0 < t; a0++){
        var R_temp = readLine().split(' ');
        var R = parseInt(R_temp[0]);
        var C = parseInt(R_temp[1]);
        var G = [];
        for(var G_i = 0; G_i < R; G_i++){
           G[G_i] = readLine();
        }
        var r_temp = readLine().split(' ');
        var r = parseInt(r_temp[0]);
        var c = parseInt(r_temp[1]);
        var P = [];
        for(var P_i = 0; P_i < r; P_i++){
           P[P_i] = readLine();
        }
		console.log(proccessMatrix(G, P, R, r, C, c) ? 'YES' : 'NO')
    }
}

function proccessMatrix(G, P, R, r, C, c) {
	var match = false;
	for (var i = 0; i <= R - r; i += 1) {
		for (var j = 0; j <= C - c; j += 1) {
			match = true

matchLoop:	for (var l = 0; l < r; l += 1) {
				for (var m = 0; m < c; m += 1) {
					if (G[i+l].charAt(j+m) != P[l].charAt(m))
					{
						match = false
						break matchLoop
					}
				}
			}

			if (match) {
				return true
			}
		}
	}
	return false
}
