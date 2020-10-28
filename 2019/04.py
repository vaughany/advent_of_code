#!/usr/bin/python3

# Advent of Code 2019. Day Four. http://adventofcode.com/2019/day/4

import time, collections

print ("Advent of Code 2019, Day Four.")

timestart = time.time()

out1 = out2 = 0

with open('04.txt') as instructions_file:
    instructions = instructions_file.read().splitlines()

instructions = instructions[0].split('-')

for password in range(int(instructions[0]), int(instructions[1]) + 1):
    password = str(password)
    sorted_pasword = sorted(password)
    if len(sorted_pasword) != len(set(sorted_pasword)) and sorted_pasword == list(password):
        out1 += 1
        if 2 in collections.Counter(sorted_pasword).values():
            out2 += 1

timetaken = time.time() - timestart

# Part One: 1330
# Part Two: 876
print("Part One:\t", out1)
if out2:
    print("Part Two:\t", out2)
print("Time taken:\t", round(timetaken * 1000, 3), 'ms')
