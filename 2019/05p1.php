#!/usr/bin/php
<?php

// Advent of Code 2019. Day Five. http://adventofcode.com/2019/day/5

echo 'Advent of Code 2019, Day Five.' . PHP_EOL;

$time_start = microtime(true);

$out1 = $out2 = [];

const DEBUG = false;

$file = (DEBUG) ? '05-sample2.txt' : '05.txt';
$instructions = file_get_contents($file, FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES) or die('No file.' . PHP_EOL);
$instructions = array_map('intval', explode(',', $instructions));

// Part One.
$ins = $instructions;
$input = 1;
$j = 0;
while(true) {
    echo "Loop: {$j}" . ". Instruction: {$ins[$j]}" . PHP_EOL;
    $opcode = substr($ins[$j], strlen($ins[$j]) - 2);

    if ($ins[$j] == 1 || ($opcode == '01' && strlen($ins[$j]) >= 3)) {

        $param1type = $param2type = $param3type = 0;
        if (strlen( $ins[$j]) > 1) {
            $param1type = (strlen($ins[$j]) >= 3) ? substr($ins[$j], strlen($ins[$j]) - 3, 1) : 0;
            $param2type = (strlen($ins[$j]) >= 4) ? substr($ins[$j], strlen($ins[$j]) - 4, 1) : 0;
            $param3type = (strlen($ins[$j]) >= 5) ? substr($ins[$j], strlen($ins[$j]) - 5, 1) : 0;
        }
        $param1number = $j + 1;
        $param2number = $j + 2;
        $param3number = $j + 3;

        $param1 = $ins[$param1number];
        if ($param1type) {
            $value1 = $ins[$param1number];
        } else {
            $value1 = $ins[$ins[$param1number]];
        }

        $param2 = $ins[$param2number];
        if ($param2type) {
            $value2 = $ins[$param2number];
        } else {
            $value2 = $ins[$ins[$param2number]];
        }

        $result = $value1 + $value2;

        $param3 = $ins[$param3number];
        if ($param3type) {
            $value3 = $ins[$param3number];
            $ins[$param3number] = $result;
            $value3new = $ins[$param3number];
        } else {
            $value3 = $ins[$ins[$param3number]];
            $ins[$ins[$param3number]] = $result;
            $value3new = $ins[$ins[$param3number]];
        }

        echo "  Addition triggered." . PHP_EOL;
        echo "    param1 type: {$param1type}. param1 number: {$param1number}. param: {$param1}. value: {$value1}" . PHP_EOL;
        echo "    param2 type: {$param2type}. param1 number: {$param2number}. param: {$param2}. value: {$value2}" . PHP_EOL;
        echo "    param3 type: {$param3type}. param1 number: {$param3number}. param: {$param3}. value: {$value3}. new value: {$value3new} " . PHP_EOL;

        $j += 4;


    } else if ($ins[$j] == 2 || ($opcode == '02' && strlen($ins[$j]) >= 3)) {

        $param1type = $param2type = $param3type = 0;
        if (strlen( $ins[$j]) > 1) {
            $param1type = (strlen($ins[$j]) >= 3) ? substr($ins[$j], strlen($ins[$j]) - 3, 1) : 0;
            $param2type = (strlen($ins[$j]) >= 4) ? substr($ins[$j], strlen($ins[$j]) - 4, 1) : 0;
            $param3type = (strlen($ins[$j]) >= 5) ? substr($ins[$j], strlen($ins[$j]) - 5, 1) : 0;
        }
        $param1number = $j + 1;
        $param2number = $j + 2;
        $param3number = $j + 3;

        $param1 = $ins[$param1number];
        if ($param1type) {
            $value1 = $ins[$param1number];
        } else {
            $value1 = $ins[$ins[$param1number]];
        }

        $param2 = $ins[$param2number];
        if ($param2type) {
            $value2 = $ins[$param2number];
        } else {
            $value2 = $ins[$ins[$param2number]];
        }

        $param3 = $ins[$param3number];
        if ($param3type) {
            $value3 = $ins[$param3number];
        } else {
            $value3 = $ins[$ins[$param3number]];
        }

        $result = $value1 * $value2;
        if ($param3type) {
            $ins[$param3number] = $result;
        } else {
            $ins[$ins[$param3number]] = $result;
        }

        if ($param3type) {
            $value3new = $ins[$param3number];
        } else {
            $value3new = $ins[$ins[$param3number]];
        }

        echo "  Multiplication triggered." . PHP_EOL;
        echo "    param1 type: {$param1type}. param1 number: {$param1number}. param: {$param1}. value: {$value1}" . PHP_EOL;
        echo "    param2 type: {$param2type}. param1 number: {$param2number}. param: {$param2}. value: {$value2}" . PHP_EOL;
        echo "    param3 type: {$param3type}. param1 number: {$param3number}. param: {$param3}. value: {$value3}. new value: {$value3new} " . PHP_EOL;

        $j += 4;


    } else if ($ins[$j] == 3) {
        echo "  Input triggered." . PHP_EOL;
        echo "    param1 number: " . ($j + 1) . ". param: {$ins[$j + 1]}. value (before): {$ins[$ins[$j + 1]]}" . PHP_EOL;
        $ins[$ins[$j + 1]] = $input;
        echo "    param1 number: " . ($j + 1) . ". param: {$ins[$j + 1]}. value (after): {$ins[$ins[$j + 1]]}" . PHP_EOL;
        $j += 2;


    } else if ($ins[$j] == 4) {
        echo PHP_EOL . "  Output triggered." . PHP_EOL;
        echo "    param1 number: " . ($j + 1) . ". param: {$ins[$j + 1]}. value: {$ins[$ins[$j + 1]]}" . PHP_EOL . PHP_EOL;

        $out1[] = $ins[$ins[$j + 1]];
        $j += 2;


    } else if ($ins[$j] == 99) {
        break;

    } else {
        $j++;
    }
}

// Part One: 13285749
foreach ($out1 as $o) {
    if ($o > 0) {
        $out1 = $o;
        break;
    }
}
echo "Part One:\t" . $out1 . PHP_EOL;
// Part Two:
if ($out2) {
    echo "Part Two:\t" . $out2 . PHP_EOL;
}
echo "Memory usage:\t" . number_format(memory_get_peak_usage() / 1024) . 'Kb' . PHP_EOL;
// echo "Time taken:\t" . number_format(microtime(true) - $time_start, 6) . ' seconds' . PHP_EOL;
echo "Time taken:\t" . number_format((microtime(true) - $time_start) * 1000, 3) . ' ms' . PHP_EOL;
