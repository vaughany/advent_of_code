#!/usr/bin/php
<?php

// Advent of Code 2019. Day Five. http://adventofcode.com/2019/day/5

echo 'Advent of Code 2019, Day Five.' . PHP_EOL;

$time_start = microtime(true);

$out1 = $out2 = [];

const DEBUG = false;
const VERBOSE = true;

$file = (DEBUG) ? '05-sample2.txt' : '05.txt';
$instructions = file_get_contents($file, FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES) or die('No file.' . PHP_EOL);
$instructions = array_map('intval', explode(',', $instructions));

function get_param (int $j, bool $type) : array
{
    global $ins;

    if ($type) {
        return [$ins[$j], $ins[$j]];
    } else {
        return [$ins[$j], $ins[$ins[$j]]];
    }
}

function get_params (int $j, int $num = 2) : array
{
    global $ins;

    $param1type = $param2type = $param3type = 0;
    if (strlen( $ins[$j]) > 1) {
        $param1type = (strlen($ins[$j]) >= 3) ? substr($ins[$j], strlen($ins[$j]) - 3, 1) : 0;
        $param2type = (strlen($ins[$j]) >= 4) ? substr($ins[$j], strlen($ins[$j]) - 4, 1) : 0;
        if ($num == 3) {
            $param3type = (strlen($ins[$j]) >= 5) ? substr($ins[$j], strlen($ins[$j]) - 5, 1) : 0;
        }
    }
    $param1number = $j + 1;
    $param2number = $j + 2;
    if ($num == 3) {
        $param3number = $j + 3;
    }

    list ($param1, $value1) = get_param ($param1number, $param1type);
    list ($param2, $value2) = get_param ($param2number, $param2type);
    if ($num == 3) {
        list ($param3, $value3) = get_param ($param3number, $param3type);
    }

    echo (VERBOSE) ? "    param1 type: {$param1type}. param1 number: {$param1number}. param: {$param1}. value: {$value1}" . PHP_EOL : '';
    echo (VERBOSE) ? "    param2 type: {$param2type}. param1 number: {$param2number}. param: {$param2}. value: {$value2}" . PHP_EOL : '';
    if ($num == 3) {
        echo (VERBOSE) ? "    param3 type: {$param3type}. param3 number: {$param3number}. param: {$param3}. value: {$value3}" . PHP_EOL : '';
        return [$value1, $value2, $value3];
    } else {
        return [$value1, $value2];
    }
}

function set_param (int $j, int $result) : void
{
    global $ins;

    $param3type = 0;
    if (strlen( $ins[$j]) > 1) {
        $param3type = (strlen($ins[$j]) >= 5) ? substr($ins[$j], strlen($ins[$j]) - 5, 1) : 0;
    }
    $param3number = $j + 3;

    list ($param3, $value3) = get_param ($param3number, $param3type);

    if ($param3type) {
        $ins[$param3number] = $result;
    } else {
        $ins[$ins[$param3number]] = $result;
    }

    list ($param3, $value3new) = get_param ($param3number, $param3type);

    echo (VERBOSE) ? "    param3 type: {$param3type}. param1 number: {$param3number}. param: {$param3}. value: {$value3}. new value: {$value3new} " . PHP_EOL : '';
}

function actually_set_param (int $j, int $result) : void
{
    global $ins;

    $paramtype = 0;
    if (strlen( $ins[$j]) > 1) {
        $paramtype = (strlen($ins[$j]) >= 5) ? substr($ins[$j], strlen($ins[$j]) - 5, 1) : 0;
    }
    $paramnumber = $j + 3;

    list ($param, $value) = get_param ($paramnumber, $paramtype);

    if ($paramtype) {
        $ins[$paramnumber] = $result;
    } else {
        $ins[$ins[$paramnumber]] = $result;
    }

    list ($param, $valuenew) = get_param ($paramnumber, $paramtype);

    echo (VERBOSE) ? "    param type: {$paramtype}. param number: {$paramnumber}. param: {$param}. value: {$value}. new value: {$valuenew} " . PHP_EOL : '';
}


// Part One.
$ins = $instructions;
$input = 1;
$j = 0;
while(true) {
    echo (VERBOSE) ? "Loop: {$j}" . ". Instruction: {$ins[$j]}" . PHP_EOL : '';
    $opcode = substr($ins[$j], strlen($ins[$j]) - 2);

    if ($ins[$j] == 1 || ($opcode == '01' && strlen($ins[$j]) >= 3)) {
        echo (VERBOSE) ? "  Addition triggered." . PHP_EOL : '';
        list ($value1, $value2) = get_params ($j);
        set_param ($j, $value1 + $value2);
        $j += 4;

    } else if ($ins[$j] == 2 || ($opcode == '02' && strlen($ins[$j]) >= 3)) {
        echo (VERBOSE) ? "  Multiplication triggered." . PHP_EOL : '';
        list ($value1, $value2) = get_params ($j);
        set_param ($j, $value1 * $value2);
        $j += 4;

    } else if ($ins[$j] == 3) {
        echo (VERBOSE) ? "  Input triggered." . PHP_EOL : '';
        echo (VERBOSE) ? "    param1 number: " . ($j + 1) . ". param: {$ins[$j + 1]}. value (before): {$ins[$ins[$j + 1]]}" . PHP_EOL : '';
        $ins[$ins[$j + 1]] = $input;
        echo (VERBOSE) ? "    param1 number: " . ($j + 1) . ". param: {$ins[$j + 1]}. value (after): {$ins[$ins[$j + 1]]}" . PHP_EOL : '';
        $j += 2;

    } else if ($ins[$j] == 4) {
        echo (VERBOSE) ? PHP_EOL . "  Output triggered." . PHP_EOL : '';
        echo (VERBOSE) ? "    param1 number: " . ($j + 1) . ". param: {$ins[$j + 1]}. value: {$ins[$ins[$j + 1]]}" . PHP_EOL . PHP_EOL : '';
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
echo "================================================================================================================================================================" . PHP_EOL;










// $instructions = array_map('intval', explode(',', '3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9')); // zero, position
// $instructions = array_map('intval', explode(',', '3,3,1105,-1,9,1101,0,0,12,4,12,99,1')); // zero, immediate

// $instructions = array_map('intval', explode(',', '3,9,8,9,10,9,4,9,99,-1,8')); // == 8, position
// $instructions = array_map('intval', explode(',', '3,9,7,9,10,9,4,9,99,-1,8')); // < 8, position
// $instructions = array_map('intval', explode(',', '3,3,1108,-1,8,3,4,3,99')); // == 8, immediate
// $instructions = array_map('intval', explode(',', '3,3,1107,-1,8,3,4,3,99')); // < 8, immediate

// $instructions = array_map('intval', explode(',', '3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99')); // < 8, immediate

// Part Two.
$ins = $instructions;
$input = 5;
$j = 0;
while(true) {
    echo (VERBOSE) ? "Loop: {$j}" . ". Instruction: {$ins[$j]}" . PHP_EOL : '';
    $opcode = substr($ins[$j], strlen($ins[$j]) - 2);

    if ($ins[$j] == 1) {
        echo (VERBOSE) ? "  Addition triggered." . PHP_EOL : '';
        set_param ($j, $ins[$j + 1] + $ins[$j + 2]);
        $j += 4;

    } else if ($opcode == '01' && strlen($ins[$j]) >= 3) {
        echo (VERBOSE) ? "  Addition triggered." . PHP_EOL : '';
        list ($value1, $value2) = get_params ($j);
        set_param ($j, $value1 + $value2);
        $j += 4;

    } else if ($ins[$j] == 2) {
        echo (VERBOSE) ? "  Multiplication triggered." . PHP_EOL : '';
        set_param ($j, $ins[$j + 1] * $ins[$j + 2]);
        $j += 4;

    } else if ($opcode == '02' && strlen($ins[$j]) >= 3) {
        echo (VERBOSE) ? "  Multiplication triggered." . PHP_EOL : '';
        list ($value1, $value2) = get_params ($j);
        set_param ($j, $value1 * $value2);
        $j += 4;

    } else if ($ins[$j] == 3) {
        echo (VERBOSE) ? "  Input triggered." . PHP_EOL : '';
        echo (VERBOSE) ? "    param1 number: " . ($j + 1) . ". param: {$ins[$j + 1]}. value (before): {$ins[$ins[$j + 1]]}" . PHP_EOL : '';
        $ins[$ins[$j + 1]] = $input;
        echo (VERBOSE) ? "    param1 number: " . ($j + 1) . ". param: {$ins[$j + 1]}. value (after): {$ins[$ins[$j + 1]]}" . PHP_EOL : '';
        $j += 2;

    } else if ($ins[$j] == 4) {
        echo (VERBOSE) ? PHP_EOL . "  Output triggered." . PHP_EOL : '';
        echo (VERBOSE) ? "    param1 number: " . ($j + 1) . ". param: {$ins[$j + 1]}. value: {$ins[$ins[$j + 1]]}" . PHP_EOL . PHP_EOL : '';
        $out2[] = $ins[$ins[$j + 1]];
        // if ($ins[$ins[$j + 1]] != 0) {
        //     var_dump($ins[$ins[$j + 1]]);
        //     die();
        // }
        $j += 2;

    } else if ($opcode == '04' && strlen($ins[$j]) >= 3) {
        echo (VERBOSE) ? PHP_EOL . "  Output triggered." . PHP_EOL : '';
        echo (VERBOSE) ? "    param1 number: " . ($j + 1) . ". param: {$ins[$j + 1]}. value: {$ins[$j + 1]}" . PHP_EOL . PHP_EOL : '';
        $out2[] = $ins[$j + 1];
        $j += 2;

    } else if ($ins[$j] == 5 || ($opcode == '05' && strlen($ins[$j]) >= 3)) {
        echo (VERBOSE) ? PHP_EOL . "  Jump-If-True triggered." . PHP_EOL : '';
        list ($value1, $value2) = get_params ($j);
        if ($value1) {
            $j = $value2;
        } else {
            $j += 3;
        }

    } else if ($ins[$j] == 6 || ($opcode == '06' && strlen($ins[$j]) >= 3)) {
        echo (VERBOSE) ? PHP_EOL . "  Jump-If-False triggered." . PHP_EOL : '';
        list ($value1, $value2) = get_params ($j);
        if (!$value1) {
            $j = $value2;
        } else {
            $j += 3;
        }

    } else if ($ins[$j] == 7 || ($opcode == '07' && strlen($ins[$j]) >= 3)) {
        echo (VERBOSE) ? PHP_EOL . "  Less-Than triggered." . PHP_EOL : '';
        list ($value1, $value2) = get_params ($j);
        if ($value1 < $value2) {
            actually_set_param ($j, 1);
        } else {
            actually_set_param ($j, 0);
        }
        $j += 4;

    } else if ($ins[$j] == 8 || ($opcode == '08' && strlen($ins[$j]) >= 3)) {
        echo (VERBOSE) ? PHP_EOL . "  Equals triggered." . PHP_EOL : '';
        list ($value1, $value2) = get_params ($j);
        if ($value1 == $value2) {
            actually_set_param ($j, 1);
        } else {
            actually_set_param ($j, 0);
        }
        $j += 4;

    } else if ($ins[$j] == 99) {
        break;

    // } else if ($ins[$j] > 25) {
    //     die();

    } else {
        $j++;
    }
}



// Part Two: NOT 4877834, not 447, not 21602
var_dump($out2);
foreach ($out2 as $o) {
    if ($o > 0) {
        $out2 = $o;
        break;
    }
}
if ($out2) {
    echo "Part Two:\t" . $out2 . PHP_EOL;
}
echo "Memory usage:\t" . number_format(memory_get_peak_usage() / 1024) . 'Kb' . PHP_EOL;
// echo "Time taken:\t" . number_format(microtime(true) - $time_start, 6) . ' seconds' . PHP_EOL;
echo "Time taken:\t" . number_format((microtime(true) - $time_start) * 1000, 3) . ' ms' . PHP_EOL;
