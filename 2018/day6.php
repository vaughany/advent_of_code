#!/usr/bin/php
<?php

// Advent of Code 2018. Day Six. http://adventofcode.com/2018/day/6

// https://github.com/jmcastagnetto/Math_Distance/blob/master/src/Math/Distance/Manhattan.php

$timestart = microtime(true);

echo 'Advent of Code 2018, Day Six.' . PHP_EOL;

$out1 = $out2 = 0;
$letters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz';
$ins = [];

$instructions = '311, 74
240, 84
54, 241
99, 336
53, 244
269, 353
175, 75
119, 271
267, 301
251, 248
216, 259
327, 50
120, 248
56, 162
42, 278
309, 269
176, 74
305, 86
93, 359
311, 189
85, 111
255, 106
286, 108
233, 228
105, 211
66, 256
213, 291
67, 53
308, 190
320, 131
254, 179
338, 44
88, 70
296, 113
278, 75
92, 316
274, 92
147, 121
71, 181
113, 268
317, 53
188, 180
42, 267
251, 98
278, 85
268, 266
334, 337
74, 69
102, 227
194, 239';

$instructions = '1, 1
1, 6
8, 3
3, 4
5, 5
8, 9';

function draw(array $in, array $grid) : string {
    global $letters, $maxx, $maxy;
    // $maxx++;
    // $maxy++;
    // for ($y = 0; $y <= $maxy; $y++ ) {
    //     for ($x = 0; $x <= $maxx; $x++ ) {
    //         $grid[$x][$y] = '.';
    //     }
    // }

    // foreach ($in as $key => $i) {
    //     $grid[$i['x']][$i['y']] = substr($letters, $key, 1);
    // }

    $out = '';
    for ($y = 0; $y <= $maxy; $y++ ) {
        for ($x = 0; $x <= $maxx; $x++ ) {
            $out .= $grid[$x][$y];
        }
        $out .= PHP_EOL;
    }
    return $out;
}

function distance(array $x, array $y) : int {
    $sum = 0;
    for ($j = 0, $c = count($x); $j < $c; $j++) {
        $sum += abs($x[$j] - $y[$j]);
    }
    return $sum;
}
// echo distance([0,1], [7,8]); die();

// Part One.
$split_lineend = explode(PHP_EOL, $instructions);
$maxx = $maxy = 0;
$grid = [];

foreach ($split_lineend as $line) {
    $tmp = explode(', ', $line);
    $ins[] = ['x' => (int) $tmp[0], 'y' => (int) $tmp[1]];
    if ($tmp[0] > $maxx) {
        $maxx = (int) $tmp[0];
    }
    if ($tmp[1] > $maxy) {
        $maxy = (int) $tmp[1];
    }
}
// Create grid.
$maxx++;
$maxy++;
for ($y = 0; $y <= $maxy + 1; $y++ ) {
    for ($x = 0; $x <= $maxx + 1; $x++ ) {
        $grid[$x][$y] = '.';
    }
}
// Letters on the grid.
foreach ($ins as $key => $i) {
    $grid[$i['x']][$i['y']] = substr($letters, $key, 1);
}

// var_dump($grid);

foreach ($ins as $key => $i) {
    for ($y = 0; $y <= $maxy + 1; $y++ ) {
        for ($x = 0; $x <= $maxx + 1; $x++ ) {
            // $grid[$x][$y] = '.';
            // echo implode(',', $i) . ' x ' . $i['x'] . ',' . $i['y'] . ': ' . distance([$x, $y], [$i['x'], $i['y']]) . PHP_EOL;
            // echo substr($letters, $key, 1) . ' (' . implode(',', $i) . "): {$x},{$y}" . PHP_EOL;
            // echo substr($letters, $key, 1) . ' (' . implode(',', $i) . ") to {$x},{$y}: " . distance([$i['x'], $i['y']], [$x, $y]) . PHP_EOL;

        }
    }
}



// foreach ($ins as $v1) {
//     foreach ($ins as $v2) {
//         if ($v1 == $v2) {
//             continue;
//         }
//         echo implode(',', $v1) . ' x ' . implode(',', $v2) . ': ' . distance([$v1['x'], $v1['y']], [$v2['x'], $v2['y']]) . PHP_EOL;
//     }
// }



// foreach () {
//
// }















echo draw($ins, $grid);


// Part One:
$timepartone = microtime(true);
echo "Part One:\t" . $out1 . "\t(time taken:\t" . number_format($timepartone - $timestart, 2) . 's)' . PHP_EOL;

// Part Two.



$timeend = microtime(true);

// Part Two:
// echo "Part Two:\t" . $out2 . PHP_EOL;
echo "Part Two:\t" . $out1 . "\t(time taken:\t" . number_format($timeend - $timepartone, 2) . 's)' . PHP_EOL;
echo "Memory usage:\t" . number_format(memory_get_peak_usage() / 1024) . 'Kb' . PHP_EOL;
echo "Time Taken:\t" . number_format($timeend - $timestart, 6) . 's' . PHP_EOL;
