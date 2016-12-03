<?php

// Advent of Code 2016. Day One.

class dayone {
    private $instructions, $facing, $coord, $coordhistory, $firstlocationtwice;
    const DEBUG = false;
    const DIRECTIONS = ['N', 'E', 'S', 'W'];

    public function __construct() {
        $this->instructions = explode(', ', 'L4, L3, R1, L4, R2, R2, L1, L2, R1, R1, L3, R5, L2, R5, L4, L3, R2, R2, L5, L1, R4, L1, R3, L3, R5, R2, L5, R2, R1, R1, L5, R1, L3, L2, L5, R4, R4, L2, L1, L1, R1, R1, L185, R4, L1, L1, R5, R1, L1, L3, L2, L1, R2, R2, R2, L1, L1, R4, R5, R53, L1, R1, R78, R3, R4, L1, R5, L1, L4, R3, R3, L3, L3, R191, R4, R1, L4, L1, R3, L1, L2, R3, R2, R4, R5, R5, L3, L5, R2, R3, L1, L1, L3, R1, R4, R1, R3, R4, R4, R4, R5, R2, L5, R1, R2, R5, L3, L4, R1, L5, R1, L4, L3, R5, R5, L3, L4, L4, R2, R2, L5, R3, R1, R2, R5, L5, L3, R4, L5, R5, L3, R1, L1, R4, R4, L3, R2, R5, R1, R2, L1, R4, R1, L3, L3, L5, R2, R5, L1, L4, R3, R3, L3, R2, L5, R1, R3, L3, R2, L1, R4, R3, L4, R5, L2, L2, R5, R1, R2, L4, L4, L5, R3, L4');

        $this->facing = 0;
        $this->coord = ['x' => 0, 'y' => 0];
        $this->coordhistory['0|0'] = [0, 0];
    }

    private function direction() {
        return self::DIRECTIONS[$this->facing];
    }

    private function left() {
        $this->facing--;
        if ($this->facing < 0) {
            $this->facing = 3;
        }
    }

    private function right() {
        $this->facing++;
        if ($this->facing > 3) {
            $this->facing = 0;
        }
    }

    private function move() {
        if ($this->facing == 0) {
            $this->coord['y']++;
        } else if ($this->facing == 1) {
            $this->coord['x']++;
        } else if ($this->facing == 2) {
            $this->coord['y']--;
        } else if ($this->facing == 3) {
            $this->coord['x']--;
        }
    }

    public function go() {
        foreach ($this->instructions as $i) {
            $turn = substr($i, 0, 1);
            $steps = substr($i, 1);

            $turn == 'L' ? $this->left() : $this->right();
            if (self::DEBUG) {
                echo 'Turn ' . $turn . ' (' . $this->direction() . ')' . PHP_EOL;
            }

            for ($s = 1; $s <= $steps; $s++) {
                $this->move();

                if (in_array($this->coord['x'] . '|' . $this->coord['y'], $this->coordhistory)) {
                    if (self::DEBUG) {
                        echo '  Visited ' . $this->coord['x'] . '|' . $this->coord['y'] . ' already.' . PHP_EOL;
                    }
                    $this->firstlocationtwice[] = [$this->coord['x'] . '|' . $this->coord['y'], abs($this->coord['x']) + abs($this->coord['y'])];
                }
                $this->coordhistory[] = $this->coord['x'] . '|' . $this->coord['y'];
            }

            if (self::DEBUG) {
                echo 'Walked ' . $steps . PHP_EOL;
                echo $this->coord['x'] . '|' . $this->coord['y'] . PHP_EOL . PHP_EOL;
            }
        }

        echo 'Final coordinate: ' . $this->coord['x'] . '|' . $this->coord['y'] . PHP_EOL;
        echo 'Total distance away from start: ' . (abs($this->coord['x']) + abs($this->coord['y'])) . PHP_EOL;
        echo 'First coordinate visited twice: ' . $this->firstlocationtwice[0][0] . PHP_EOL;
        echo 'First coordinate visited twice distance away from start: ' . $this->firstlocationtwice[0][1] . PHP_EOL;
    }

}

$dayone = new dayone();
$dayone->go();

// Final coordinate:                                x:-173; y:-159;
// Total distance from start (x:0; y:0;):           332
// First twice-visited coordinate:                  x:-158; y:8
// Distance from first twice-visited coordinate:    166
