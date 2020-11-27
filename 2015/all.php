#!/usr/bin/php
<?php

// Advent of Code 2015. http://adventofcode.com/2015/

echo 'Advent of Code 2015' . PHP_EOL . PHP_EOL;

for ($i = 1; $i <= 25; $i++) {
    $filename = "day" . $i . ".php";
    if (file_exists($filename)) {
        echo shell_exec("./" . $filename) . PHP_EOL;
    }
}
