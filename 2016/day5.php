<?php

// Advent of Code 2016. Day Five.

$instructions = 'reyedfim';

$c = 1;
$out = '';
$hashes = [];
for ($j = 1; $j <= 10000000; $j++) {
    $plaintext = $instructions . $j;
    $hash = md5($plaintext);
    if (preg_match('/^00000/', $hash)) {
        $out .= substr($hash, 5, 1);
        echo $out . PHP_EOL;
        if ($c == 8) {
            break;
        }
        $c++;
    }
}
echo PHP_EOL . "Password (Part One): '$out'." . PHP_EOL . PHP_EOL;

$password = '--------';
for ($j = 1; $j <= 100000000; $j++) {
    $plaintext = $instructions . $j;
    $hash = md5($plaintext);
    if (preg_match('/^00000/', $hash)) {
        $pos = substr($hash, 5, 1);
        if (is_numeric($pos) && $pos <= 7 && substr($password, $pos, 1) == '-') {
            $char = substr($hash, 6, 1);
            $password = substr_replace($password, $char, $pos, 1);
            echo $password . PHP_EOL;
            if (!strstr($password, '-')) {
                break;
            }
        }
    }
}
echo "Password (Part Two): '$password'." . PHP_EOL;

// Password (Part One): 'f97c354d'.
// Password (Part Two): '863dde27'.
