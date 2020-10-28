#!/usr/bin/php
<?php

// Advent of Code 2019. Day Three. http://adventofcode.com/2019/day/3

echo 'Advent of Code 2019, Day Three.' . PHP_EOL;

$time_start = microtime(true);

$out1 = $out2 = 0;

const DEBUG = false;

$file = (DEBUG) ? '03-sample.txt' : '03.txt';
$instructions = file($file, FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES) or die('No file.' . PHP_EOL);

/**
 * Draws the grid.
 *
 * @return string
 */
function draw() : string {
    global $grid, $minx, $miny, $maxx, $maxy;
    $out = PHP_EOL;
    for ($y = $miny; $y <= $maxy; $y++ ) {
        for ($x = $minx; $x <= $maxx; $x++ ) {
            if (isset($grid[$x][$y])) {
                $out .= $grid[$x][$y];
            } else {
                $out .= '.';
            }
        }
        $out .= PHP_EOL;
    }
    return $out .= PHP_EOL;
}

/**
 * Claculates the Manhattan / taxi distance. Can work in n dimensions.
 *
 * @param array $vector1
 * @param array $vector2
 * @return int
 */
function distance(array $vector1, array $vector2) : int {
    $sum = 0;
    for ($j = 0, $c = count($vector1); $j < $c; $j++) {
        $sum += abs($vector1[$j] - $vector2[$j]);
    }
    return $sum;
}

/**
 * Create the 'instructions'; return them and max-x and max-y.
 *
 * @param array $instructions
 * @return array
 */
function create_instructions(array $instructions) : array {
    $ins = [];
    foreach ($instructions as $line) {
        $ins[] = explode(',', $line);
    }
    return $ins;
}

// Part One.
$minx = $miny = $maxx = $maxy = 0;
$grid[0][0] = 'O';
$grid_points[0][0] = 'O';
$crosses = [];
$ins = create_instructions($instructions);

foreach ($ins as $ins_no => $wire) {
    $x = $y = 0;
    $wire_pieces = count($wire);
    $wire_piece = 0;

    foreach ($wire as $w) {
        $wire_piece++;
        $direction = substr($w, 0, 1);
        $distance = (int) substr($w, 1);

        if ($direction == 'L') {
            $grid_points[$x - $distance][$y] = $ins_no;
        } else if ($direction == 'R') {
            $grid_points[$x + $distance][$y] = $ins_no;
        } else if ($direction == 'U') {
            $grid_points[$x][$y + $distance] = $ins_no;
        } else if ($direction == 'D') {
            $grid_points[$x][$y - $distance] = $ins_no;
        }

        for ($i = 1; $i <= $distance; $i++) {

            if ($direction == 'L') {
                --$x; $char = '<';
            } else if ($direction == 'R') {
                ++$x; $char = '>';
            } else if ($direction == 'U') {
                --$y; $char = 'A';
            } else if ($direction == 'D') {
                ++$y; $char = 'V';
            }

            // if (isset($grid[$x][$y])) {
            if (isset($grid[$x][$y]) && $grid[$x][$y] != $ins_no) {
                $grid[$x][$y] = 'x';
                $d = distance([0, 0], [$x, $y]);
                $crosses[$d] = ['x' => $x, 'y' => $y, 'd' => $d];
                $crosses2[$x][$y] = $d;
            } else {
                // $grid[$x][$y] = ($i == $distance && $wire_piece != $wire_pieces) ? '+' : $char;
                // $grid[$x][$y] = ($i == $distance && $wire_piece != $wire_pieces) ? '+' : $ins_no;
                $grid[$x][$y] = $ins_no;
                // $grid[$x][$y] = o;
            }

            $minx = ($x < $minx) ? $x : $minx;
            $maxx = ($x > $maxx) ? $x : $maxx;
            $miny = ($y < $miny) ? $y : $miny;
            $maxy = ($y > $maxy) ? $y : $maxy;
        }
    }
}

if (DEBUG) {
    echo PHP_EOL . draw();
}
ksort($crosses);
$out1 = array_key_first($crosses);

// Part One: 1674
echo "Part One:\t" . $out1 . PHP_EOL;

// var_dump($grid);
// var_dump($crosses2);
// var_dump(array_intersect($grid, $crosses2));

// Part Two.
foreach ($instructions as $line) {
    //
}


// Part Two:
if ($out2) {
    echo "Part Two:\t" . $out2 . PHP_EOL;
}
echo "Memory usage:\t" . number_format(memory_get_peak_usage() / 1024) . 'Kb' . PHP_EOL;
// echo "Time taken:\t" . number_format(microtime(true) - $time_start, 6) . ' seconds' . PHP_EOL;
echo "Time taken:\t" . number_format((microtime(true) - $time_start) * 1000, 3) . ' ms' . PHP_EOL;
