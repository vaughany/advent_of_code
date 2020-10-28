#!/usr/bin/php
<?php

// Advent of Code 2019. Day Four. http://adventofcode.com/2019/day/4

echo 'Advent of Code 2019, Day Four.' . PHP_EOL;

$time_start = microtime(true);

$out1 = $out2 = 0;

const DEBUG = false;

$file = (DEBUG) ? '04-sample.txt' : '04.txt';
$instructions = file($file, FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES) or die('No file.' . PHP_EOL);

function six_digits(string $in) : bool {
    if (strlen($in) == 6) {
        return true;
    }
    return false;
}

function adjacent(string $in) : bool {
    $matches = ['00', '11', '22', '33', '44', '55', '66', '77', '88', '99'];
    foreach ($matches as $m) {
        $position = strpos($in, $m);
        if ($position !== false) {
            return true;
        }
    }
    return false;
}

function increasing(string $in) : bool {
    // We don't need to check for six digits, as we already did that.
    $ints = [];
    for ($j = 0; $j <= 5; $j++) {
        $ints[$j] = (int) $in[$j];
    }
    if ($ints[0] <= $ints[1] && $ints[1] <= $ints[2] && $ints[2] <= $ints[3] && $ints[3] <= $ints[4] && $ints[4] <= $ints[5]) {
        return true;
    }
    return false;
}

function pairs(string $in) : bool {
    $letters = [];
    for ($i = 0; $i <= 5; $i++) {
        if (isset($letters[$in[$i]])) {
            $letters[$in[$i]]++;
        } else {
            $letters[$in[$i]] = 1;
        }
    }
    if (in_array(2, $letters)) {
        return true;
    }
    return false;
}

if (DEBUG) {
    $lower = (int) $instructions[0];
    $upper = (int) $instructions[0] + 1;
} else {
    list($lower, $upper) = array_map('intval', explode('-', $instructions[0]));
}

for ($j = $lower; $j <= $upper; $j++) {
    // Part One.
    if (six_digits($j) && adjacent($j) && increasing($j)) {
        $out1++;
        // Part Two.
        if (pairs($j)) {
            $out2++;
        }
    }
}

// Part One: 1330
// Part Two: 876
echo "Part One:\t" . $out1 . PHP_EOL;
if ($out2) {
    echo "Part Two:\t" . $out2 . PHP_EOL;
}
echo "Memory usage:\t" . number_format(memory_get_peak_usage() / 1024) . 'Kb' . PHP_EOL;
// echo "Time taken:\t" . number_format(microtime(true) - $time_start, 6) . ' seconds' . PHP_EOL;
echo "Time taken:\t" . number_format((microtime(true) - $time_start) * 1000, 3) . ' ms' . PHP_EOL;
