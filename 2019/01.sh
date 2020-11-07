#!/bin/bash

# Run with `./01.sh` (assuming you made it executable with `chmod +x`) or `time ./01.sh` to get timings.

shellcheck "$0" || exit 1

echo -e '\e[1mAdvent of Code 2019, Day One.\e[0m'

IN="${0:2:2}.txt"
OUT1=
OUT2=

# Part One.
while IFS= read -r LINE; do
  OUT1=$(( OUT1 += (( LINE / 3 - 2 ))))
done < "$IN"

# Part Two.
while IFS= read -r LINE; do
  while true; do
    LINE=$((LINE / 3 - 2))
    if test $LINE -le 0; then
      break;
    fi
    OUT2=$(( OUT2 += (( LINE ))))
  done
done < "$IN"

# Part One: 3363929
# Part Two: 5043026
echo "Part One: ${OUT1}"
echo "Part Two: ${OUT2}"
