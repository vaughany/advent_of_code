#!/usr/bin/php
<?php

// Advent of Code 2018. Day Five. http://adventofcode.com/2018/day/5

echo 'Advent of Code 2018, Day Five.' . PHP_EOL;

$time_start = microtime(true);

$out1 = $out2 = 0;
$letters = 'A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z';
$upper = explode(',', $letters);
$lower = explode(',', strtolower($letters));

$instructions = file('05.txt', FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES) or die('No file.' . PHP_EOL);
$instructions = $instructions[0];

// Part One.
$p1ins = $instructions;
$count = $oldlen = $newlen = 0;
while(true) {
    $count++;
    $newlen = strlen($p1ins);
    if ($oldlen == $newlen) {
        break;
    }
    for ($j = 0; $j <= $newlen - 2; $j++) {
        $c1 = substr($p1ins, $j, 1);
        $c2 = substr($p1ins, $j + 1, 1);
        if ((in_array($c1, $upper) && in_array($c2, $lower)) || (in_array($c1, $lower) && in_array($c2, $upper))) {
            if (array_keys($upper, strtoupper($c1)) == array_keys($lower, strtolower($c2))) {
                $p1ins = substr_replace($p1ins ,'', $j, 2);
            }
        }
    }
    $oldlen = $newlen;
}
$out1 = strlen($p1ins);

// Part One: 9370
$time_part_one = microtime(true);
echo "Part One:\t" . $out1 . "\t(time taken:\t" . number_format($time_part_one - $time_start, 2) . ' s)' . PHP_EOL;

// Part Two.
$lengths = [];
$time_prev_letter = microtime(true);
foreach ($lower as $remove) {
    echo ' Processing ' . strtoupper($remove) . ":\t";
    $count = $oldlen = $newlen = 0;
    $p2ins = str_ireplace($remove, '', $instructions);
    while(true) {
        $count++;
        $newlen = strlen($p2ins);
        if ($oldlen == $newlen) {
            break;
        }
        for ($j = 0; $j <= $newlen - 2; $j++) {
            $c1 = substr($p2ins, $j, 1);
            $c2 = substr($p2ins, $j + 1, 1);
            if ((in_array($c1, $upper) && in_array($c2, $lower)) || (in_array($c1, $lower) && in_array($c2, $upper))) {
                if (array_keys($upper, strtoupper($c1)) == array_keys($lower, strtolower($c2))) {
                    $p2ins = substr_replace($p2ins ,'', $j, 2);
                }
            }
        }
        $oldlen = $newlen;
    }
    $lengths[$remove] = strlen($p2ins);
    echo "{$newlen}\t(time taken:\t" . number_format(microtime(true) - $time_prev_letter, 2) . ' s)' . PHP_EOL;
    $time_prev_letter = microtime(true);
}
$out2 = min($lengths);

// Part Two: 6390 (D)
if ($out2) {
    echo "Part Two:\t" . $out2 . "\t(time taken:\t" . number_format(microtime(true) - $time_part_one, 2) . ' s)' . PHP_EOL;
}
echo "Memory usage:\t" . number_format(memory_get_peak_usage() / 1024) . 'Kb' . PHP_EOL;
echo "Time Taken:\t" . number_format(microtime(true) - $time_start, 2) . ' s' . PHP_EOL;
