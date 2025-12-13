from src.day10 import *

def test_build_machine():
    m = machine("[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}")
    assert m.lights == "[.##.]"
    assert m.buttons == [[3], [1,3], [2], [2,3], [0,2], [0,1]]
    assert m.joltage == [3,5,4,7]

def test_fewest_presses_for_joltage():
    m = machine("[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}")
    assert m.fewest_presses_for_joltage() == 10

def test_solve_part_2():
    lines = [
        "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
        "[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
        "[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
    ]
    assert solve_part_2(lines) == 10 + 12 + 11
    