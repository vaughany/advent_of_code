#!/usr/bin/python3

# Advent of Code 2019. Day Three. http://adventofcode.com/2019/day/3

import time

print ("Advent of Code 2019, Day Three.")

timestart = time.time()

out1 = out2 = 0

with open('03.txt') as instructions_file:
    instructions = instructions_file.read().splitlines()

# Part One.
for line in instructions:
    #

# Part Two.
for line in instructions:
    #

timetaken = time.time() - timestart

# Part One:
# Part Two: 
print("Part One:\t", out1)
if out2:
    print("Part Two:\t", out2)
print("Time taken:\t", round(timetaken * 1000, 3), 'ms')
