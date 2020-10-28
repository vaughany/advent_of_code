#!/usr/bin/php
<?php

// Advent of Code 2018. Day Eleven. http://adventofcode.com/2018/day/11

echo 'Advent of Code 2018, Day Eleven.' . PHP_EOL;

$time_start = microtime(true);

$out1 = $out2 = $largest = $largestx = $largesty = 0;
$grid = [];

$instructions = (int) file('11.txt', FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES) or die('No file.' . PHP_EOL);

function draw3x3(int $inx, int $iny) : string {
    global $grid;
    $out = PHP_EOL;
    for ($y = $iny; $y < $iny + 3; $y++ ) {
        for ($x = $inx; $x < $inx + 3; $x++ ) {
            if (isset($grid[$x][$y])) {
                // $out .= $grid[$x][$y]['int'] . "\t";
                $out .= $grid[$x][$y]['pwr'] . "\t";
            } else {
                throw new Exception("Grid element doesn't exist");
            }
        }
        $out .= PHP_EOL;
    }
    return $out . PHP_EOL;
}

function sum_xys(int $inx, int $iny, int $s) : int {
    global $grid;
    $out = 0;
    for ($y = $iny; $y < $iny + $s; $y++) {
        for ($x = $inx; $x < $inx + $s; $x++) {
            if (isset($grid[$x][$y][$s])) {
                $out += $grid[$x][$y][$s]['pwr'];
            } else {
                throw new Exception("Grid element doesn't exist");
            }
        }
    }
    return $out;
}

function get_hundred(int $in) : int {
    $in = (string) $in;
    if (strlen($in) < 3) {
        return 0;
    }
    return (int) substr($in, -3, 1);
}

$s = 3;
for ($y = 1; $y <= 300; $y++ ) {
    for ($x = 1; $x <= 300; $x++ ) {
        $grid[$x][$y][$s]['rid'] = $x + 10;
        $grid[$x][$y][$s]['pwr'] = ($grid[$x][$y][$s]['rid'] * $y) + $instructions;
        $grid[$x][$y][$s]['pwr'] = $grid[$x][$y][$s]['pwr'] * $grid[$x][$y][$s]['rid'];
        $grid[$x][$y][$s]['pwr'] = get_hundred($grid[$x][$y][$s]['pwr']) - 5;
    }
}

// Part One.
for ($y = 1; $y <= 298; $y++ ) {
    for ($x = 1; $x <= 298; $x++ ) {
        $grid[$x][$y][$s]['totalpwr'] = sum_xys($x, $y, $s);
        if ($grid[$x][$y][$s]['totalpwr'] > $largest) {
            $largest = $grid[$x][$y][$s]['totalpwr'];
            $largestx = $x;
            $largesty = $y;
        }
    }
}
$out1 = "{$largestx},{$largesty}";

// Part One: "34,72"
$time_part_one = microtime(true);
echo "Part One:\t" . $out1 . "\t(time taken:\t" . number_format(($time_part_one - $time_start) * 1000, 3) . ' ms)' . PHP_EOL;

// Part Two.
for ($s = 1; $s <= 300; $s++) {
    $grid = [];
    echo $s . '.';
    for ($y = 1; $y <= 300; $y++ ) {
        for ($x = 1; $x <= 300; $x++ ) {
            $grid[$x][$y][$s]['rid'] = $x + 10;
            $grid[$x][$y][$s]['pwr'] = ($grid[$x][$y][$s]['rid'] * $y) + $instructions;
            $grid[$x][$y][$s]['pwr'] = $grid[$x][$y][$s]['pwr'] * $grid[$x][$y][$s]['rid'];
            $grid[$x][$y][$s]['pwr'] = get_hundred($grid[$x][$y][$s]['pwr']) - 5;
        }
    }
    for ($y = 1; $y <= 300 - ($s - 1); $y++ ) {
        for ($x = 1; $x <= 300 - ($s - 1); $x++ ) {
            $grid[$x][$y][$s]['totalpwr'] = sum_xys($x, $y, $s);
            if ($grid[$x][$y][$s]['totalpwr'] > $largest) {
                $largest = $grid[$x][$y][$s]['totalpwr'];
                $largestx = $x;
                $largesty = $y;
                $largests = $s;

                echo PHP_EOL . "{$largestx},{$largesty},{$largests}" . PHP_EOL;
            }
        }
    }
}
$out2 = "{$largestx},{$largesty},{$largests}";

$time_end = microtime(true);

// Part Two: "233,187,13"
if ($out2) {
    echo "Part Two:\t" . $out2 . "\t(time taken:\t" . number_format($time_end - $time_part_one, 6) . ' s)' . PHP_EOL;
}
echo "Memory usage:\t" . number_format(memory_get_peak_usage() / 1024) . 'Kb' . PHP_EOL;
echo "Time Taken:\t" . number_format($time_end - $time_start, 6) . ' s' . PHP_EOL;
