<?php

// Advent of code. Day One.

class dayone {
    private $directions, $facing, $coord, $coordhistory, $firstlocationtwice;

    public function __construct() {
        $this->directions = explode(', ', 'L4, L3, R1, L4, R2, R2, L1, L2, R1, R1, L3, R5, L2, R5, L4, L3, R2, R2, L5, L1, R4, L1, R3, L3, R5, R2, L5, R2, R1, R1, L5, R1, L3, L2, L5, R4, R4, L2, L1, L1, R1, R1, L185, R4, L1, L1, R5, R1, L1, L3, L2, L1, R2, R2, R2, L1, L1, R4, R5, R53, L1, R1, R78, R3, R4, L1, R5, L1, L4, R3, R3, L3, L3, R191, R4, R1, L4, L1, R3, L1, L2, R3, R2, R4, R5, R5, L3, L5, R2, R3, L1, L1, L3, R1, R4, R1, R3, R4, R4, R4, R5, R2, L5, R1, R2, R5, L3, L4, R1, L5, R1, L4, L3, R5, R5, L3, L4, L4, R2, R2, L5, R3, R1, R2, R5, L5, L3, R4, L5, R5, L3, R1, L1, R4, R4, L3, R2, R5, R1, R2, L1, R4, R1, L3, L3, L5, R2, R5, L1, L4, R3, R3, L3, R2, L5, R1, R3, L3, R2, L1, R4, R3, L4, R5, L2, L2, R5, R1, R2, L4, L4, L5, R3, L4');

        $this->facing = 'n';
        $this->coord = ['x' => 0, 'y' => 0];
        $this->coordhistory['0|0'] = [0, 0];
    }

    private function turnleft() {
        if ($this->facing == 'n') {
            $this->facing = 'w';

        } else if ($this->facing == 'e') {
            $this->facing = 'n';

        } else if ($this->facing == 's') {
            $this->facing = 'e';

        } else if ($this->facing == 'w') {
            $this->facing = 's';
        }
    }

    private function turnright() {
        if ($this->facing == 'n') {
            $this->facing = 'e';

        } else if ($this->facing == 'e') {
            $this->facing = 's';

        } else if ($this->facing == 's') {
            $this->facing = 'w';

        } else if ($this->facing == 'w') {
            $this->facing = 'n';
        }
    }

    private function move(int $steps) {
        if ($this->facing == 'n') {
            $this->coord['y'] += $steps;

        } else if ($this->facing == 'e') {
            $this->coord['x'] += $steps;

        } else if ($this->facing == 's') {
            $this->coord['y'] -= $steps;

        } else if ($this->facing == 'w') {
            $this->coord['x'] -= $steps;

        }
    }

    public function go() {
        foreach ($this->directions as $direction) {
            $turn = substr($direction, 0, 1);
            $steps = substr($direction, 1);

            if ($turn == 'L') {
                $this->turnleft();
            } else {
                $this->turnright();
            }
            echo 'Turn ' . $turn . ' / ' . $this->facing . PHP_EOL;

            for ($s = 1; $s <= $steps; $s++) {
                $this->move(1);

                foreach ($this->coordhistory as $key => $value) {
                    if ($key != $this->coord['x'] . '|' . $this->coord['y']) {
                        $this->coordhistory[$this->coord['x'] . '|' . $this->coord['y']] = [$this->coord['x'], $this->coord['y']];

                    } else {
                        echo 'Visited this location already! x:' . $this->coord['x'] . '; y:' . $this->coord['y'] . ';' . PHP_EOL;
                        $this->firstlocationtwice[] = [$this->coord['x'], $this->coord['y']];
                    }
                }
            }

            echo 'Walked ' . $steps . PHP_EOL;
            echo 'x:' . $this->coord['x'] . '; y:' . $this->coord['y'] . ';' . PHP_EOL . PHP_EOL;
        }

        echo 'Total distance away from start: ' . (abs($this->coord['x']) + abs($this->coord['y'])) . PHP_EOL . PHP_EOL;
        echo 'First coordinate visited twice: x:' . $this->firstlocationtwice[0][0] . '; y:' . $this->firstlocationtwice[0][1] . ';' . PHP_EOL;
        echo 'First coordinate visited twice distance away from start: ' . (abs($this->firstlocationtwice[0][0]) + abs($this->firstlocationtwice[0][1])) . PHP_EOL;
    }

}

$dayone = new dayone();
$dayone->go();

// Final coordinate:                                x:-173; y:-159;
// Total distance from start (x:0; y:0;):           332
// First twice-visited coordinate:                  x:-158; y:8
// Distance from first twice-visited coordinate:    166
