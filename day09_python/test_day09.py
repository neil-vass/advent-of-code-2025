from src.day09 import *

square = ["0,0", "0,1", "1,1", "1,0"]

L_shape = ["0,0", "0,2", "3,2", "3,5", "5,5", "5,0"]

example = [
	"7,1",
	"11,1",
	"11,7",
	"9,7",
	"9,5",
	"2,5",
	"2,3",
	"7,3",
]


def test_solve_part_2():
    assert parse_polygon(square).solve_part_2() == 4
    assert parse_polygon(L_shape).solve_part_2() == 18
    assert parse_polygon(example).solve_part_2() == 24
