#!/usr/bin/php
<?php

// Advent of Code 2018. Day Seven. http://adventofcode.com/2018/day/7

echo 'Advent of Code 2018, Day Seven.' . PHP_EOL;

$time_start = microtime(true);

$out1 = $out2 = 0;
$off = ' ';
$on = '#';

$instructions = file('07.txt', FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES) or die('No file.' . PHP_EOL);

function draw(array $in) : string {
    global $minx, $miny, $maxx, $maxy, $off, $on;
    $out = PHP_EOL;
    for ($y = $miny; $y <= $maxy; $y++ ) {
        for ($x = $minx; $x <= $maxx; $x++ ) {
            $out .= (isset($in[$x][$y])) ? $on : $off;
        }
        $out .= PHP_EOL;
    }
    $out .= PHP_EOL;
    return $out;
}

// Part One.
$ins = [];
$split_lineend = explode(PHP_EOL, $instructions);
foreach ($split_lineend as $line) {
    $x = $y = $vx = $vy = 0;
    $x = (int) substr($line, 10, 6);
    $y = (int) substr($line, 18, 6);
    $vx = (int) substr($line, 36, 2);
    $vy = (int) substr($line, 40, 2);
    $ins[] = ['x' => $x, 'y' => $y, 'vx' => $vx, 'vy' => $vy];
}

for ($j = 0; $j <= 11000; $j++) {
    // Get the min and max X and Y values.
    // Gotcha: the eventual words are not located at 0,0 as the sample data was...
    $minx = $miny = $maxx = $maxy = 140;
    foreach ($ins as $i) {
        if ($i['x'] < $minx) {
            $minx = $i['x'];
        } else if ($i['x'] > $maxx) {
            $maxx = $i['x'];
        }
        if ($i['y'] < $miny) {
            $miny = $i['y'];
        } else if ($i['y'] > $maxy) {
            $maxy = $i['y'];
        }
    }

    // Draw the instructions.
    $grid = [];
    foreach ($ins as $i) {
        $grid[$i['x']][$i['y']] = true;
    }
    // $grid[0][0] = '0';

    // The number here was reached by watching the output and iteratively limiting the grids which were drawn.
    if ($maxy < 149) {
        echo draw($grid);
        // Part Two.
        $out2 = $j;
    }

    // Edit the grid.
    foreach ($ins as $key => $i) {
        $ins[$key]['x'] += $i['vx'];
        $ins[$key]['y'] += $i['vy'];
    }
}
$out1 = 'GJNKBZEE';

// Part One: GJNKBZEE
$time_part_one = microtime(true);
echo "Part One:\t" . $out1 . "\t(time taken:\t" . number_format(($time_part_one - $time_start) * 1000, 3) . ' ms)' . PHP_EOL;

// Part Two: 10727
if ($out2) {
    echo "Part Two:\t" . $out2 . "\t\t(time taken:\t" . number_format((microtime(true) - $time_part_one) * 1000, 3) . ' ms)' . PHP_EOL;
}
echo "Memory usage:\t" . number_format(memory_get_peak_usage() / 1024) . 'Kb' . PHP_EOL;
echo "Time Taken:\t" . number_format((microtime(true) - $time_start) * 1000, 3) . ' ms' . PHP_EOL;
