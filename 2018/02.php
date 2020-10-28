#!/usr/bin/php
<?php

// Advent of Code 2018. Day Two. http://adventofcode.com/2018/day/2

echo 'Advent of Code 2018, Day Two.' . PHP_EOL;

$time_start = microtime(true);

$out1 = $two = $three = 0;
$out2 = '';

$instructions = file('02.txt', FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES) or die('No file.' . PHP_EOL);

$len = 26;

// Part One.
foreach ($instructions as $line) {
    $letters = [];
    for ($j=0; $j <= ($len - 1); $j++) {
        $l = substr($line, $j, 1);
        if (isset($letters[$l])) {
            $letters[$l]++;
        } else {
            $letters[$l] = 1;
        }
    }
    $twodone = $threedone = false;
    foreach ($letters as $key => $value) {
        if ($value == 2 && $twodone == false) {
            $two++;
            $twodone = true;
        } else if ($value == 3 && $threedone == false) {
            $three++;
            $threedone = true;
        } else {
            unset ($letters[$key]);
        }
    }
    $out1 = $two * $three;
}

// Part Two.
foreach ($instructions as $line) {
    for ($j = 0, $count = count($instructions) - 1; $j <= $count; $j++) {
        $diffs = 0;
        $out2 = '';
        for ($k = 0; $k <= $len - 1; $k++ ) {
            if (substr($line, $k, 1) !== substr($instructions[$j], $k, 1)) {
                $diffs++;
            } else {
                $out2 .= substr($line, $k, 1);
            }
        }
        if ($diffs == 1) {
            break 2;
        }
    }
}

// Part One: 5456
// Part Two: megsdlpulxvinkatfoyzxcbvq
echo "Part One:\t" . $out1 . PHP_EOL;
if ($out2) {
    echo "Part Two:\t" . $out2 . PHP_EOL;
}
echo "Memory usage:\t" . number_format(memory_get_peak_usage() / 1024) . 'Kb' . PHP_EOL;
echo "Time taken:\t" . number_format((microtime(true) - $time_start) * 1000, 3) . ' ms' . PHP_EOL;
