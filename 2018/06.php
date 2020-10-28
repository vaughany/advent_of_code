#!/usr/bin/php
<?php

// Advent of Code 2018. Day Six. http://adventofcode.com/2018/day/6

echo 'Advent of Code 2018, Day Six.' . PHP_EOL;

$time_start = microtime(true);

$out1 = $out2 = 0;
$letters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz';
$ins = [];

$instructions = file('06.txt', FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES) or die('No file.' . PHP_EOL);

/**
 * Draws the grid.
 *
 * @param array $grid
 * @return string
 */
function draw(array $grid) : string {
    global $maxx, $maxy;

    $out = PHP_EOL;
    for ($y = 1; $y <= $maxy; $y++ ) {
        for ($x = 1; $x <= $maxx; $x++ ) {
            $out .= $grid[$x][$y];
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
    $maxx = $maxy = 0;

    foreach ($instructions as $line) {
        $tmp = explode(', ', $line);
        $ins[] = ['x' => (int) $tmp[0], 'y' => (int) $tmp[1]];
        if ($tmp[0] > $maxx) {
            $maxx = (int) $tmp[0];
        }
        if ($tmp[1] > $maxy) {
            $maxy = (int) $tmp[1];
        }
    }

    return [$ins, $maxx, $maxy];
}

/**
 * Create a gird of dots. Add letters at coords from instructions.
 *
 * @param array $ins
 * @param int $maxx
 * @param int $maxy
 * @param string $letters
 * @return array
 */
function create_grid(array $ins, int $maxx, int $maxy, string $letters) : array {
    $grid = [];
    for ($y = 1; $y <= $maxy + 1; $y++ ) {
        for ($x = 1; $x <= $maxx + 1; $x++ ) {
            $grid[$x][$y] = '.';
        }
    }

    foreach ($ins as $key => $i) {
        $grid[$i['x']][$i['y']] = substr($letters, $key, 1);
    }

    return $grid;
}

// Part One.
list($ins, $maxx, $maxy) = create_instructions($instructions);
$grid = create_grid($ins, $maxx, $maxy, $letters);

for ($y = 1; $y <= $maxy; $y++) {
    for ($x = 1; $x <= $maxx; $x++) {

        if ($grid[$x][$y] == '.') {

            $distances = [];
            foreach ($ins as $key => $i) {
                $distances[substr($letters, $key, 1)] = distance([$x, $y], [$i['x'], $i['y']]);
            }
            $min = min($distances);
            $mins = array_count_values($distances);
            if ($mins[$min] > 1) {
                // Leave the grid as it is.
            } else {
                $letter = (string) array_keys($distances, $min)[0];
                $grid[$x][$y] = $letter;
            }
        }
        // echo PHP_EOL . "Inspecting {$x},{$y}:" . draw($grid);
    }
}

echo draw($grid);

// Get full counts of all areas.
$counts = [];
for ($y = 1; $y <= $maxy; $y++) {
    for ($x = 1; $x <= $maxx; $x++) {
        if (isset($counts[$grid[$x][$y]])) {
            $counts[$grid[$x][$y]]++;
        } else {
            $counts[$grid[$x][$y]] = 1;
        }
    }
}
unset($counts['.']);

// Get full counts of all areas.
$infinites = [];
for ($y = 1; $y <= $maxy; $y++) {
    if (!isset($infinites[$grid[1][$y]])) {
        $infinites[$grid[1][$y]] = 1;
    }
    if (!isset($infinites[$grid[$maxx][$y]])) {
        $infinites[$grid[$maxx][$y]] = 1;
    }
}
for ($x = 1; $x <= $maxx; $x++) {
    if (!isset($infinites[$grid[$x][1]])) {
        $infinites[$grid[$x][1]] = 1;
    }
    if (!isset($infinites[$grid[$x][$maxy]])) {
        $infinites[$grid[$x][$maxy]] = 1;
    }
}

// Unset the 'infinites' from the count.
foreach ($infinites as $key => $i) {
    unset($counts[$key]);
}

$out1 = max($counts);

// Part One: 5035
$time_part_one = microtime(true);
echo "Part One:\t" . $out1 . "\t(time taken:\t" . number_format($time_part_one - $time_start, 2) . ' s)' . PHP_EOL;

// Part Two.

// Recreate grid. TODO: DRY.
for ($y = 1; $y <= $maxy + 1; $y++ ) {
    for ($x = 1; $x <= $maxx + 1; $x++ ) {
        $grid[$x][$y] = '.';
    }
}
// Put the letters on the correct coordinates on the grid.  TODO: DRY.
foreach ($ins as $key => $i) {
    $grid[$i['x']][$i['y']] = substr($letters, $key, 1);
}
echo draw($grid);







// Part Two:
// echo "Part Two:\t" . $out2 . PHP_EOL;
if ($out2) {
    echo "Part Two:\t" . $out1 . "\t(time taken:\t" . number_format($time_end - $time_part_one, 2) . ' s)' . PHP_EOL;
}
echo "Memory usage:\t" . number_format(memory_get_peak_usage() / 1024) . 'Kb' . PHP_EOL;
echo "Time Taken:\t" . number_format(microtime(true) - $time_start, 3) . ' s' . PHP_EOL;
