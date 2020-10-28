#!/usr/bin/php
<?php

// Advent of Code 2019. Day Two. http://adventofcode.com/2019/day/2

echo 'Advent of Code 2019, Day Two.' . PHP_EOL;

$time_start = microtime(true);

$out1 = $out2 = 0;

$part2_output = 19690720;

$instructions = file_get_contents('02.txt', FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES) or die('No file.' . PHP_EOL);
$instructions = explode(',', $instructions);

// Part One.
$ins = $instructions;
$ins[1] = 12;
$ins[2] = 2;
for ($j = 0; $j <= count($ins) + 1; $j += 4) {
    if ($ins[$j] == 1) {
        $ins[$ins[$j + 3]] = $ins[$ins[$j + 1]] + $ins[$ins[$j + 2]];
    } else if ($ins[$j] == 2) {
        $ins[$ins[$j + 3]] = $ins[$ins[$j + 1]] * $ins[$ins[$j + 2]];
    } else if ($ins[$j] == 99) {
        break;
    }
}
$out1 = $ins[0];

// Part Two.
$loops = 0;
for ($noun = 0; $noun <= 99; $noun++) {
    for ($verb = 0; $verb <= 99; $verb++) {
        $ins = $instructions;
        $ins[1] = $noun;
        $ins[2] = $verb;
        $loops++;
        for ($j = 0; $j <= count($ins) + 1; $j += 4) {
            if ($ins[$j] == 1) {
                $ins[$ins[$j + 3]] = $ins[$ins[$j + 1]] + $ins[$ins[$j + 2]];
            } else if ($ins[$j] == 2) {
                $ins[$ins[$j + 3]] = $ins[$ins[$j + 1]] * $ins[$ins[$j + 2]];
            } else if ($ins[$j] == 99) {
                break;
            }
        }

        if ($ins[0] == $part2_output) {
            $out2 = (100 * $noun + $verb) . " ({$loops} loops)";
            break;
        }
    }
}


// Part One: 3790645
// Part Two: 6577
echo "Part One:\t" . $out1 . PHP_EOL;
if ($out2) {
    echo "Part Two:\t" . $out2 . PHP_EOL;
}
echo "Memory usage:\t" . number_format(memory_get_peak_usage() / 1024) . 'Kb' . PHP_EOL;
// echo "Time taken:\t" . number_format(microtime(true) - $time_start, 6) . ' seconds' . PHP_EOL;
echo "Time taken:\t" . number_format((microtime(true) - $time_start) * 1000, 3) . ' ms' . PHP_EOL;
