#!/usr/bin/php
<?php

// Advent of Code 2020. Day One. http://adventofcode.com/2020/day/1

echo 'Advent of Code 2020, Day One.' . PHP_EOL;

$time_start = microtime(true);

$out1 = $out2 = 0;

$instructions = file('01.txt', FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES) or die('No file.' . PHP_EOL);

// Part One.
foreach ($instructions as $line) {
    //
}

// Part Two.
foreach ($instructions as $line) {
    //
}

// Part One:
// Part Two:
echo "Part One:\t" . $out1 . PHP_EOL;
if ($out2) {
    echo "Part Two:\t" . $out2 . PHP_EOL;
}
echo "Memory usage:\t" . number_format(memory_get_peak_usage() / 1024) . 'Kb' . PHP_EOL;
echo "Time taken:\t" . number_format((microtime(true) - $time_start) * 1000, 3) . ' ms' . PHP_EOL;
