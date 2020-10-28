#!/usr/bin/php
<?php

// Advent of Code 2018. Day Three. http://adventofcode.com/2018/day/3

echo 'Advent of Code 2018, Day Three.' . PHP_EOL;

$time_start = microtime(true);

$out1 = $out2 = 0;
$fabric = [];
$unalloc = '.';
$alloc = 'o';
$overalloc = 'x';

$instructions = file('03.txt', FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES) or die('No file.' . PHP_EOL);

// Prepare the grid.
$w = $h = 1000;
for ($y = 0; $y <= $h; $y++) {
    for ($x = 0; $x <= $w; $x++) {
        $fabric[$y][$x] = $unalloc;
    }
}

// Part One.
foreach ($instructions as $key => $line) {
    // Just some stuff to make the source data nicer.
    $line = explode(' ', $line);
    $line['id'] = $line[0];
    $tmp = explode(',', $line[2]);
    $line['x'] = (int) $tmp[0];
    $line['y'] = (int) substr($tmp[1], 0, strlen($tmp[1]) - 1 );
    $tmp = explode('x', $line[3]);
    $line['w'] = (int) $tmp[0];
    $line['h'] = (int) $tmp[1];
    for ($j = 0; $j <= 3; $j++) {
        unset($line[$j]);
    }
    // Save it back for part 2.
    $instructions[$key] = $line;

    // Fill.
    for ($k = $line['y']; $k < $line['y'] + $line['h']; $k++ ) {
        for ($j = $line['x']; $j < $line['x'] + $line['w']; $j++ ) {
            if ($fabric[$j][$k] == $alloc) {
                $fabric[$j][$k] = $overalloc;
            } else if ($fabric[$j][$k] == $overalloc) {
                // Do nothing.
            } else {
                $fabric[$j][$k] = $alloc;
            }
        }
    }
}
for ($y = 0; $y <= $h; $y++) {
    for ($x = 0; $x <= $w; $x++) {
        if ($fabric[$x][$y] == $overalloc) {
            $out1++;
        }
    }
}

// Part Two.
foreach ($instructions as $key => $line) {
    $overwritten = false;
    for ($k = $line['y']; $k < $line['y'] + $line['h']; $k++ ) {
        for ($j = $line['x']; $j < $line['x'] + $line['w']; $j++ ) {
            if ($fabric[$j][$k] == $overalloc) {
                $overwritten = true;
                break 2;
            }
        }
    }
    if ($overwritten == false) {
        $out2 = substr($line['id'], 1);
        break;
    }
}

// Part One: 107043
// Part Two: 346
echo "Part One:\t" . $out1 . PHP_EOL;
if ($out2) {
    echo "Part Two:\t" . $out2 . PHP_EOL;
}
echo "Memory usage:\t" . number_format(memory_get_peak_usage() / 1024) . 'Kb' . PHP_EOL;
echo "Time taken:\t" . number_format((microtime(true) - $time_start) * 1000, 3) . ' ms' . PHP_EOL;
