<?php
//https://www.hackerrank.com/challenges/matrix-rotation-algo
print "\n******\n";
$handle = fopen ("php://stdin","r");
$line = fgets($handle);
$line = explode(' ', $line);
$M = $line[0];
$N = $line[1];
$R = $line[2];
$A = [];
for($a0 = 0; $a0 < $M; $a0++){
    $line = fgets($handle);
    $A[] = explode(' ', $line);
}
$AN = [];
//print sum1($N, $M, 10)."";
//$X = plainCoord(3, 3, $M, $N);
//print $X."\n";
//var_dump(plainTo2D(2, 2, 3));
//die;

printMatrix($A, $M, $N);
$margin = 0;
while ($M - $margin * 2 >= 2 && $N - $margin * 2 >= 2) {
	$MN = $M - $margin * 2;
	$NN = $N - $margin * 2;
	for ($x = 0; $x <= perimeter($MN, $NN) - 1; $x += 1) {
		list($i, $j) = plainTo2D($x, $MN, $NN);
		list($in, $jn) = plainTo2D($x + $R, $MN, $NN);
		$AN[$in + $margin][$jn + $margin] = intval($A[$i + $margin][$j + $margin]);
		print 'x='.$x.', i='.$i.', j='.$j.
			', M='.($M - $margin*2).', N='.($N - $margin*2).
			', in='.$in.', jn='.$jn."\n";
	}
	$margin += 1;
}
printMatrix($AN, $M, $N);

print "\n******\n";exit(1);

function printMatrix($matrix, $M, $N) {
	for ($i = 0; $i < $M; $i += 1) {
		for ($j = 0; $j < $N; $j += 1) {
			print sprintf('%3d', isset($matrix[$i][$j]) ? $matrix[$i][$j] : 0);
		}
		print PHP_EOL;
	}
}

function plainTo2D($plainX, $M, $N) {
	$perimeter = perimeter($M, $N);

	//normalize
	$plainX = $plainX > $perimeter ? $plainX % $perimeter : $plainX;

	//calculate
	if ($plainX <= $perimeter / 2) {
		$i = min($M - 1, $plainX);
		$j = $plainX - $i;
	} else {
		$i = $M - 1 - min($plainX - $perimeter / 2, $M - 1);
		$j = $perimeter - $plainX - $i;
	}
	return [$i, $j];
}

function perimeter($M, $N) {
	return 2 * ($M + $N - 2);
}
/*
1 + 3  = 4 1
1 + 5  = 4 2
1 + 9  = 1 4
1 + 11 = 1 2

$M /
*/
