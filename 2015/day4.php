#!/usr/bin/php
<?php

// Advent of Code 2015. Day Four.

$instructions = 'yzbqklnj';

for ($j = 1; $j <= 10000000; $j++) {
    $plaintext = $instructions . $j;
    $hash = md5($plaintext);
    if (preg_match('/^00000/', $hash)) {
        echo 'First hash with "00000" at the start: ' . $hash . ' from input: "' . $plaintext . '" is ' . $j . '.' . PHP_EOL;
        break;
    }
}

for ($j = 1; $j <= 10000000; $j++) {
    $plaintext = $instructions . $j;
    $hash = md5($plaintext);
    if (preg_match('/^000000/', $hash)) {
        echo 'First hash with "000000" at the start: ' . $hash . ' from input: "' . $plaintext . '" is ' . $j . '.' . PHP_EOL;
        break;
    }
}

// First hash with "00000" at the start: 000002c655df7738246e88f6c1c43eb7 from input: "yzbqklnj282749".
// First hash with "000000" at the start: 0000004b347bf4b398b3f62ace7cd301 from input: "yzbqklnj9962624".
