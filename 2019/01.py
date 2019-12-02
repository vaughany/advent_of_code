#!/usr/bin/python3

# Advent of Code 2019. Day One. http://adventofcode.com/2019/day/1

import sys, time

print ("Advent of Code 2019, Day One.")

timestart = time.time()

out1 = out2 = 0

with open('01.txt') as instructions_file:
    # instructions = instructions_file.readlines()
    instructions = instructions_file.read().splitlines()

# Part One.
for line in instructions:
    # print(line.rstrip())
    # out1 = out1 + int( int(line.rstrip()) / 3 - 2 )
    out1 += int( int(line) / 3 - 2 )

# Part Two.
for line in instructions:
    while True:
        line = int( int(line) / 3 - 2 )
        if line <= 0:
            break
        out2 += line

timetaken = time.time() - timestart

# Part One: 3363929
# Part Two: 5043026
print("Part One:\t", out1)
if out2:
    print("Part Two:\t", out2)
print("Time taken:\t", round(timetaken * 1000, 3), 'ms')
