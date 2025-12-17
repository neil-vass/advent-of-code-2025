
import math
import numpy as np

class Pos:
    def __init__(self, x, y):
        self.x = x
        self.y = y

class Rectangle:
    def __init__(self, min_pos, max_pos):
        self.min_pos = min_pos
        self.max_pos = max_pos

def area(rect):
    return ((rect.max_pos.x - rect.min_pos.x + 1) *
            (rect.max_pos.y - rect.min_pos.y + 1))

class Polygon:
    def  __init__(self, vertices, min_bound, max_bound, candidates):
        self.vertices = vertices
        self.min_bound = min_bound
        self.max_bound = max_bound
        self.candidates = candidates
        self.mark_outside()
    
    def mark_outside(self):
        self.outside = np.zeros((self.max_bound.x+3, self.max_bound.y+3), dtype=bool)

        # Make clockwise.
        signed_area = 0
        for i in range(len(self.vertices) - 1):
            start, end = self.vertices[i], self.vertices[i+1]
            signed_area += (start.x*end.y - end.x*start.y)
	
        if signed_area > 0:
            self.vertices.reverse()

        # Mark the "you're stepping outside" border.
        for i in range(len(self.vertices) - 1):
            edge = (self.vertices[i], self.vertices[i+1])
            if edge[0].x == edge[1].x: # is_horizontal
                if edge[0].y < edge[1].y:
                    x = edge[0].x -1
                    y_start = edge[0].y
                    y_end = edge[1].y +1
                    self.outside[x, y_start:y_end] = True
                else:
                    x = edge[0].x +1
                    y_start = edge[1].y
                    y_end = edge[0].y +1
                    self.outside[x, y_start:y_end] = True
            else:
                if edge[0].x < edge[1].x:
                    y = edge[0].y +1
                    x_start = edge[0].x
                    x_end = edge[1].x +1
                    self.outside[x_start:x_end, y] = True
                else:
                    y = edge[0].y -1
                    x_start = edge[1].x
                    x_end = edge[0].x +1
                    self.outside[x_start:x_end, y] = True

        # Remove edges that got mistakenly marked as "outside" above.
        for i in range(len(self.vertices) - 1):
            edge = (self.vertices[i], self.vertices[i+1])
            if edge[0].x == edge[1].x: # is_horizontal
                x = edge[0].x
                y_start = min(edge[0].y, edge[1].y)
                y_end = max(edge[0].y, edge[1].y) +1
                self.outside[x, y_start:y_end] = False
            else:
                y = edge[0].y
                x_start = min(edge[0].x, edge[1].x)
                x_end = max(edge[0].x, edge[1].x) +1
                self.outside[x_start:x_end, y] = False			


    def solve_part_1(self): 
        winner = self.candidates[0]
        return area(winner)


    def solve_part_2(self):
        for i, c in enumerate(self.candidates):
            if not (self.outside[c.min_pos.x, c.min_pos.y:c.max_pos.y+1].any() or
                    self.outside[c.max_pos.x, c.min_pos.y:c.max_pos.y+1].any() or
                    self.outside[c.min_pos.x:c.max_pos.x+1, c.min_pos.y].any() or
                    self.outside[c.min_pos.x:c.max_pos.x+1, c.max_pos.y].any()):
                return area(c)

        raise ValueError("No suitable rectangles at all")
        

def parse_polygon(lines):
    tiles = []
    rectanges = []
    min_bound = Pos(math.inf, math.inf)
    max_bound = Pos(-math.inf, -math.inf)

    for ln in lines:
        nums = [int(n)+1 for n in ln.split(',')]
        tile = Pos(x=nums[0], y=nums[1])
        min_bound = Pos(min(min_bound.x, tile.x), min(min_bound.y, tile.y))
        max_bound = Pos(max(max_bound.x, tile.x), max(max_bound.y, tile.y))

        for otherTile in tiles:
            min_pos = Pos(min(tile.x, otherTile.x), min(tile.y, otherTile.y))
            max_pos = Pos(max(tile.x, otherTile.x), max(tile.y, otherTile.y))
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
