#!/usr/bin/python3

# Advent of Code 2019. Day Two. http://adventofcode.com/2019/day/2

import time

print ("Advent of Code 2019, Day Two.")

timestart = time.time()

out1 = out2 = 0
part2_output = 19690720

with open('02.txt') as instructions_file:
    instructions = list(map(int, instructions_file.read().split(',')))

# Part One.
ins, ins[1], ins[2] = instructions[:], 12, 2
for j in range(0, len(ins)):
    if j % 4 == 0:
        if ins[j] == 1:
            ins[ins[j + 3]] = ins[ins[j + 1]] + ins[ins[j + 2]];
        elif ins[j] == 2:
            ins[ins[j + 3]] = ins[ins[j + 1]] * ins[ins[j + 2]];
        elif ins[j] == 99:
            break
out1 = ins[0];

# Part Two.
for noun in range(0, 100):
    for verb in range(0, 100):
        ins, ins[1], ins[2] = instructions[:], noun, verb
        for j in range(0, len(ins), 4):
            if ins[j] == 1:
                ins[ins[j + 3]] = ins[ins[j + 1]] + ins[ins[j + 2]];
            elif ins[j] == 2:
                ins[ins[j + 3]] = ins[ins[j + 1]] * ins[ins[j + 2]];
            elif ins[j] == 99:
                break
        if (ins[0] == part2_output):
            out2 = (100 * noun + verb)
            break

timetaken = time.time() - timestart

# Part One: 3790645
# Part Two: 6577
print("Part One:\t", out1)
if out2:
    print("Part Two:\t", out2)
print("Time taken:\t", round(timetaken * 1000, 3), 'ms')
