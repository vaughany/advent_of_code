#!/usr/bin/php
<?php

// Advent of Code 2018. Day Twelve. http://adventofcode.com/2018/day/12

echo 'Advent of Code 2018, Day Twelve.' . PHP_EOL;

$timestart = microtime(true);

$out1 = $out2 = 0;

$pots = '#.#.#...#..##..###.##.#...#.##.#....#..#.#....##.#.##...###.#...#######.....##.###.####.#....#.#..##';
$pots = '#..#.#..##......###...###';

$instructions = file('12.txt', FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES) or die('No file.' . PHP_EOL);
//$instructions = '.#..# => #';

$pots = ".....{$pots}.....";

// Part One.
$ins = [];
foreach ($instructions as $key => $line) {
    $tmp = explode(' ', $line);
    $ins[$key]['note'] = $tmp[0];
    $ins[$key]['action'] = ($tmp[2] == '#') ? 'add' : 'remove';
}
// var_dump($ins);

echo $pots . PHP_EOL;
for ($k = 1; $k <= 20; $k++) {
    for ($j = 0; $j <= strlen($pots); $j++) {
        $slice = substr($pots, $j, 5);
        if (strlen($slice) != 5) {
            continue;
        }
        foreach ($ins as $i) {
            // echo 'Checking ' . $i['note'] . ' against ' . $slice . ' : ';
            if ($i['note'] == $slice) {
                // echo 'match! ';
                // change it
                if ($i['action'] == 'add') {
                    // echo 'addition' . PHP_EOL;
                    $pots[$j+2] = '#';
                } else {
                    // echo 'subtraction' . PHP_EOL;
                    $pots[$j+2] = '.';
                }
            } else {
                // echo 'no match' . PHP_EOL;
            }
        }
        // echo PHP_EOL;
    }
    echo $pots . PHP_EOL;
}

$out1 = 0;

// Part One: "34,72"
$timepartone = microtime(true);
echo "Part One:\t" . $out1 . "\t(time taken:\t" . number_format($timepartone - $timestart, 6) . 's)' . PHP_EOL;

// Part Two.
$out2 = 0;

$timeend = microtime(true);

// Part Two: "233,187,13"
echo "Part Two:\t" . $out2 . "\t(time taken:\t" . number_format($timeend - $timepartone, 6) . 's)' . PHP_EOL;
echo "Memory usage:\t" . number_format(memory_get_peak_usage() / 1024) . 'Kb' . PHP_EOL;
echo "Time Taken:\t" . number_format($timeend - $timestart, 6) . 's' . PHP_EOL;
