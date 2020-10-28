#!/usr/bin/php
<?php

// Advent of Code 2018. Day Four. http://adventofcode.com/2018/day/4

echo 'Advent of Code 2018, Day Four.' . PHP_EOL;

$time_start = microtime(true);

$out1 = $out2 = $maxasleepminutes = $maxhighestminute = $maxguardid = 0;
$actions = $guards = [];
$toggle = 1;

$instructions = file('04.txt', FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES) or die('No file.' . PHP_EOL);

asort($instructions);

// Part One.
foreach ($instructions as $key => $line) {
    if (strpos($line, 'begins shift')) {
        $tmp = explode(' ', $line);
        $guardid = substr($tmp[3], 1);

    } else {
        $tmp = explode(' ', $line);
        $minute = substr($tmp[1], 3, 2);
        $action = ($tmp[2] == 'wakes') ? 'awake' : 'asleep';
        $actions[] = "{$guardid} {$action} {$minute}";
    }
}
foreach ($actions as $key => $action) {
    $toggle = 1 - $toggle;
    if ($toggle == 1) {
        continue;
    }

    $tmp = explode(' ', $action);
    $guardid = (int) $tmp[0];
    $asleep = (int) $tmp[2];

    $tmp = explode(' ', $actions[$key + 1]);
    $awake = (int) $tmp[2];

    for ($j = $asleep; $j < $awake; $j++) {
        if (!isset($guards[$guardid]['minutes'][$j])) {
            $guards[$guardid]['minutes'][$j] = 1;
        } else {
            $guards[$guardid]['minutes'][$j]++;
        }
        if (!isset($guards[$guardid]['asleepminutes'])) {
            $guards[$guardid]['asleepminutes'] = 1;
        } else {
            $guards[$guardid]['asleepminutes']++;
        }
    }
}
foreach ($guards as $guardid => $guard) {
    if ($guard['asleepminutes'] > $maxasleepminutes) {
        $maxasleepminutes = $guard['asleepminutes'];
        $maxhighestminute = array_search(max($guard['minutes']), $guard['minutes']);
        $maxguardid = $guardid;
    }
}
$out1 = $maxguardid * $maxhighestminute;

// Part Two.
$mgm = $minnumber = $maxminutes = $maxguardid = $maxminnumber = 0;
foreach ($guards as $guardid => $guard) {
    $mgm = max($guard['minutes']);
    $minnumber = array_keys($guard['minutes'], $mgm);
    if ($mgm > $maxminutes && count($minnumber) == 1) {
        $maxminutes = $mgm;
        $maxguardid = $guardid;
        $maxminnumber = $minnumber[0];
    }
}
$out2 = $maxguardid * $maxminnumber;

// Part One: 101194
// Part Two: 102095
echo "Part One:\t" . $out1 . PHP_EOL;
if ($out2) {
    echo "Part Two:\t" . $out2 . PHP_EOL;
}
echo "Memory usage:\t" . number_format(memory_get_peak_usage() / 1024) . 'Kb' . PHP_EOL;
echo "Time Taken:\t" . number_format((microtime(true) - $time_start) * 1000, 3) . ' ms' . PHP_EOL;
