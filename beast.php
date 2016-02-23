<?php

$handle = fopen ("php://stdin","r");
fscanf($handle,"%d",$t);
for($a0 = 0; $a0 < $t; $a0++){
    fscanf($handle,"%d",$n);
    echo beast($n);
}


function beast($n) {
	if ($n < 3)
		return -1;
	$fives = floor($n / 3) * 3;
	$threes = $n - $fives;

	while ($fives % 3 > 0 || $threes % 5 > 0) {
		$fives -= 5 - $threes % 5;
		$threes += 5 - $threes % 5;
	}
	if ($fives < 0)
		return -1;
	return beastOut($fives, $threes);
}


function beastOut($fives, $threes) {
	$d = [];
	for ($i = 0; $i<$fives; $i += 1) {
		$d[] = 5;
	}
	for ($i = 0; $i<$threes; $i += 1) {
		$d[] = 3;
	}
	return implode('', $d);
}
