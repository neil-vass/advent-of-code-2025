
import math

import numpy as np

class Rectangle:
    def __init__(self, min_pos, max_pos):
        self.min_pos = min_pos
        self.max_pos = max_pos

def area(rect):
    return ((rect.max_pos[0] - rect.min_pos[0] + 1) *
            (rect.max_pos[1] - rect.min_pos[1] + 1))

class Polygon:
    def  __init__(self, vertices, min_bound, max_bound, candidates):
        self.vertices = vertices
        self.min_bound = min_bound
        self.max_bound = max_bound
        self.candidates = candidates
        self.mark_outside()
    
    def mark_outside(self):
        self.outside = np.zeros((self.max_bound[0]+3, self.max_bound[1]+3), dtype=bool)

        # Make clockwise.
        signed_area = 0
        for i in range(len(self.vertices) - 1):
            start, end = self.vertices[i], self.vertices[i+1]
            signed_area += (start[0]*end[1] - end[0]*start[1])
	
        if signed_area > 0:
            self.vertices.reverse()

        # Mark the "you're stepping outside" border.
        for i in range(len(self.vertices) - 1):
            edge = (self.vertices[i], self.vertices[i+1])
            if edge[0][0] == edge[1][0]: # is_horizontal
                if edge[0][1] < edge[1][1]:
                    x = edge[0][0] -1
                    y_start = edge[0][1]
                    y_end = edge[1][1] +1
                    self.outside[x, y_start:y_end] = True
                else:
                    x = edge[0][0] +1
                    y_start = edge[1][1]
                    y_end = edge[0][1] +1
                    self.outside[x, y_start:y_end] = True
            else:
                if edge[0][0] < edge[1][0]:
                    y = edge[0][1] +1
                    x_start = edge[0][0]
                    x_end = edge[1][0] +1
                    self.outside[x_start:x_end, y] = True
                else:
                    y = edge[0][1] -1
                    x_start = edge[1][0]
                    x_end = edge[0][0] +1
                    self.outside[x_start:x_end, y] = True

        # Remove edges that got mistakenly marked as "outside" above.
        for i in range(len(self.vertices) - 1):
            edge = (self.vertices[i], self.vertices[i+1])
            if edge[0][0] == edge[1][0]: # is_horizontal
                x = edge[0][0]
                y_start = min(edge[0][1], edge[1][1])
                y_end = max(edge[0][1], edge[1][1]) +1
                self.outside[x, y_start:y_end] = False
            else:
                y = edge[0][1]
                x_start = min(edge[0][0], edge[1][0])
                x_end = max(edge[0][0], edge[1][0]) +1
                self.outside[x_start:x_end, y] = False			


    def solve_part_1(self): 
        winner = self.candidates[0]
        return area(winner)


    def solve_part_2(self):
        for i, c in enumerate(self.candidates):
            if not (self.outside[c.min_pos[0], c.min_pos[1]:c.max_pos[1]+1].any() or
                    self.outside[c.max_pos[0], c.min_pos[1]:c.max_pos[1]+1].any() or
                    self.outside[c.min_pos[0]:c.max_pos[0]+1, c.min_pos[1]].any() or
                    self.outside[c.min_pos[0]:c.max_pos[0]+1, c.max_pos[1]].any()):
                return area(c)

        raise ValueError("No suitable rectangles at all")
        

def parse_polygon(lines):
    tiles = []
    rectanges = []
    min_bound = [math.inf, math.inf]
    max_bound = [-math.inf, -math.inf]

    for ln in lines:
        tile = [int(n)+1 for n in ln.split(',')]
        min_bound = [min(min_bound[0], tile[0]), min(min_bound[1], tile[1])]
        max_bound = [max(max_bound[0], tile[0]), max(max_bound[1], tile[1])]

        for otherTile in tiles:
            min_pos = [min(tile[0], otherTile[0]), min(tile[1], otherTile[1])]
            max_pos = [max(tile[0], otherTile[0]), max(tile[1], otherTile[1])]
            rectanges.append(Rectangle(min_pos, max_pos))
		
        tiles.append(tile)

    # Sort by area, greatest first
    rectanges.sort(key=area, reverse=True)

    # One more tile on the end to close the loop
    tiles.append(tiles[0])
    
    poly = Polygon(tiles, min_bound, max_bound, rectanges)

    return poly
    

def fetch_data(path):
    with open(path, 'r') as f:
        for ln in f:
            yield ln 

if __name__ == "__main__":
    lines = fetch_data("./input.txt")
    poly = parse_polygon(lines)
    print(poly.solve_part_1())
    print(poly.solve_part_2())
