<?php

// Advent of Code 2017. Day Three. http://adventofcode.com/2017/day/3

$out1 = $out2 = 0;
$xmin = $ymin = -263;
$xmax = $ymax = abs($xmin);
$grid = [];
$marker = 0;
$dir = 'r';

$instructions = 277678;

// Part One.
for ($y = $ymax; $y >= $ymin; $y--) {
    for ($x = $xmin; $x <= $xmax; $x++) {
        $grid[$x][$y] = $marker;
    }
}

$x = $y = 0;
$grid[$x][$y] = 1;
for ($j = 2; $j <= $instructions; $j++) {

    if ($dir == 'r') {
        if ($grid[$x + 1][$y] == $marker) {
            $grid[++$x][$y] = $j;
            if ($grid[$x][$y + 1] == $marker) {
                $dir = 'u';
                continue;
            }
        }
    }
    if ($dir == 'u') {
        if ($grid[$x][$y + 1] == $marker) {
            $grid[$x][++$y] = $j;
            if ($grid[$x - 1][$y] == $marker) {
                $dir = 'l';
                continue;
            }
        }
    }
    if ($dir == 'l') {
        if ($grid[$x - 1][$y] == $marker) {
            $grid[--$x][$y] = $j;
            if ($grid[$x][$y - 1] == $marker) {
                $dir = 'd';
                continue;
            }
        }
    }
    if ($dir == 'd') {
        if ($grid[$x][$y - 1] == $marker) {
            $grid[$x][--$y] = $j;
            if ($grid[$x + 1][$y] == $marker) {
                $dir = 'r';
                continue;
            }
        }
    }

    if ($j == $instructions) {
        $out1 = abs($x) + abs($y);
        break;
    }
}


// Part Two.
function add_neighbours($xx = 0, $yy = 0) {
    global $grid;
    $sum = 0;

    for ($y = $yy - 1; $y <= $yy + 1; $y++) {
        for ($x = $xx - 1; $x <= $xx + 1; $x++) {
            $sum += $grid[$x][$y];
        }
    }
    return $sum;
}

$grid = [];
for ($y = $ymax; $y >= $ymin; $y--) {
    for ($x = $xmin; $x <= $xmax; $x++) {
        $grid[$x][$y] = $marker;
    }
}

$x = $y = 0;
$grid[$x][$y] = 1;
for ($j = 2; $j <= $instructions; $j++) {

    if ($dir == 'r') {
        if ($grid[$x + 1][$y] == $marker) {
            $grid[++$x][$y] = add_neighbours($x, $y);
            if ($grid[$x][$y + 1] == $marker) {
                $dir = 'u';
                continue;
            }
        }
    }
    if ($dir == 'u') {
        if ($grid[$x][$y + 1] == $marker) {
            $grid[$x][++$y] = add_neighbours($x, $y);
            if ($grid[$x - 1][$y] == $marker) {
                $dir = 'l';
                continue;
            }
        }
    }
    if ($dir == 'l') {
        if ($grid[$x - 1][$y] == $marker) {
            $grid[--$x][$y] = add_neighbours($x, $y);
            if ($grid[$x][$y - 1] == $marker) {
                $dir = 'd';
                continue;
            }
        }
    }
    if ($dir == 'd') {
        if ($grid[$x][$y - 1] == $marker) {
            $grid[$x][--$y] = add_neighbours($x, $y);
            if ($grid[$x + 1][$y] == $marker) {
                $dir = 'r';
                continue;
            }
        }
    }

    if ($grid[$x][$y] > $instructions) {
        $out2 = $grid[$x][$y];
        break;
    }
}

// CLI output. For testing. Small grids only.
// for ($y = $ymax; $y >= $ymin; $y--) {
//     for ($x = $xmin; $x <= $xmax; $x++) {
//         echo $grid[$x][$y] . "\t";
//     }
//     echo "\n";
// }

// Part One: 475
// Part Two: 279138
echo 'Part One: ' . $out1 . PHP_EOL;
echo 'Part Two: ' . $out2 . PHP_EOL;
echo 'Memory usage: ' . number_format(memory_get_peak_usage() / 1024) . 'Kb' . PHP_EOL;
