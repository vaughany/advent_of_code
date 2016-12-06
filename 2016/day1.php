<?php

// Advent of Code 2016. Day One.

$instructions = explode(', ', 'L4, L3, R1, L4, R2, R2, L1, L2, R1, R1, L3, R5, L2, R5, L4, L3, R2, R2, L5, L1, R4, L1, R3, L3, R5, R2, L5, R2, R1, R1, L5, R1, L3, L2, L5, R4, R4, L2, L1, L1, R1, R1, L185, R4, L1, L1, R5, R1, L1, L3, L2, L1, R2, R2, R2, L1, L1, R4, R5, R53, L1, R1, R78, R3, R4, L1, R5, L1, L4, R3, R3, L3, L3, R191, R4, R1, L4, L1, R3, L1, L2, R3, R2, R4, R5, R5, L3, L5, R2, R3, L1, L1, L3, R1, R4, R1, R3, R4, R4, R4, R5, R2, L5, R1, R2, R5, L3, L4, R1, L5, R1, L4, L3, R5, R5, L3, L4, L4, R2, R2, L5, R3, R1, R2, R5, L5, L3, R4, L5, R5, L3, R1, L1, R4, R4, L3, R2, R5, R1, R2, L1, R4, R1, L3, L3, L5, R2, R5, L1, L4, R3, R3, L3, R2, L5, R1, R3, L3, R2, L1, R4, R3, L4, R5, L2, L2, R5, R1, R2, L4, L4, L5, R3, L4');

$facing = 0;
$coord = ['x' => 0, 'y' => 0];
$coordhistory['0|0'] = [0, 0];

$firstlocationtwice = [];

foreach ($instructions as $i) {
    $turn = substr($i, 0, 1);
    $steps = substr($i, 1);

    $turn == 'L' ? $facing-- : $facing++;
    if ($facing < 0) { $facing = 3; }
    if ($facing > 3) { $facing = 0; }

    for ($s = 1; $s <= $steps; $s++) {
        if ($facing == 0) {
            $coord['y']++;
        } else if ($facing == 1) {
            $coord['x']++;
        } else if ($facing == 2) {
            $coord['y']--;
        } else if ($facing == 3) {
            $coord['x']--;
        }

        if (in_array($coord['x'] . '|' . $coord['y'], $coordhistory)) {
            $firstlocationtwice[] = [$coord['x'] . '|' . $coord['y'], abs($coord['x']) + abs($coord['y'])];
        }
        $coordhistory[] = $coord['x'] . '|' . $coord['y'];
    }
}

echo 'Final coordinate: ' . $coord['x'] . '|' . $coord['y'] . PHP_EOL;
echo 'Total distance away from start: ' . (abs($coord['x']) + abs($coord['y'])) . PHP_EOL;
echo 'First coordinate visited twice: ' . $firstlocationtwice[0][0] . PHP_EOL;
echo 'First coordinate visited twice distance away from start: ' . $firstlocationtwice[0][1] . PHP_EOL;

// Final coordinate: -173|-159
// Total distance away from start: 332
// First coordinate visited twice: -158|8
// First coordinate visited twice distance away from start: 166
