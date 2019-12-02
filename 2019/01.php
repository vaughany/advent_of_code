#!/usr/bin/php
<?php

// Advent of Code 2019. Day One. http://adventofcode.com/2019/day/1

echo 'Advent of Code 2019, Day One.' . PHP_EOL;

$time_start = microtime(true);

$out1 = $out2 = 0;

$instructions = file('01.txt', FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES) or die('No file.' . PHP_EOL);

// Part One.
foreach ($instructions as $line) {
    $out1 += (int) ($line / 3) - 2;
}

// Part Two.
foreach ($instructions as $line) {
    while (true) {
        $line = (int) ($line / 3) - 2;
        if ($line <= 0) {
            break;
        }
        $out2 += $line;
    }
}

// Part One: 3363929
// Part Two: 5043026
echo "Part One:\t" . $out1 . PHP_EOL;
if ($out2) {
    echo "Part Two:\t" . $out2 . PHP_EOL;
}
echo "Memory usage:\t" . number_format(memory_get_peak_usage() / 1024) . 'Kb' . PHP_EOL;
// echo "Time taken:\t" . number_format(microtime(true) - $time_start, 6) . ' seconds' . PHP_EOL;
echo "Time taken:\t" . number_format((microtime(true) - $time_start) * 1000, 3) . ' ms' . PHP_EOL;
