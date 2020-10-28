#!/usr/bin/php
<?php

// Advent of Code 2018. Day One. http://adventofcode.com/2018/day/1

echo 'Advent of Code 2018, Day One.' . PHP_EOL;

$time_start = microtime(true);

$out1 = $out2 = 0;
$frequencies = [];

$instructions = file('01.txt', FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES) or die('No file.' . PHP_EOL);

// Part One.
foreach ($instructions as $line) {
    if (substr($line, 0, 1) == '+') {
        $out1 += substr($line, 1);
    } else {
        $out1 -= substr($line, 1);
    }
}

// Part Two.
while (true) {
    foreach ($instructions as $line) {
        if (substr($line, 0, 1) == '+') {
            $out2 += substr($line, 1);
        } else {
            $out2 -= substr($line, 1);
        }
        if (isset($frequencies[$out2])) {
            break 2;
        }
        $frequencies[$out2] = 1;
    }
}

// Part One: 493
// Part Two: 413
echo "Part One:\t" . $out1 . PHP_EOL;
if ($out2) {
    echo "Part Two:\t" . $out2 . PHP_EOL;
}
echo "Memory usage:\t" . number_format(memory_get_peak_usage() / 1024) . 'Kb' . PHP_EOL;
echo "Time taken:\t" . number_format((microtime(true) - $time_start) * 1000, 3) . ' ms' . PHP_EOL;
