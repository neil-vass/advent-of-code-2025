import re
import numpy as np
from scipy.optimize import LinearConstraint
from scipy.optimize import milp

class machine:
    def __init__(self, description):
        tokens = description.split()
        self.lights = tokens[0]
        self.buttons = [[int(n) for n in re.findall(r'\d+', btn)] for btn in tokens[1:-1]]
        self.joltage = [int(n) for n in re.findall(r'\d+', tokens[-1])]
    
    def fewest_presses_for_joltage(self):
        # Aim: minimize total number of button presses
        variablesToMinimize = np.ones(len(self.buttons))
        all_integers = np.ones_like(variablesToMinimize)

        # Constraints: each button press affects some joltages
        joltages_affected_by_btns = np.zeros((len(self.joltage), len(self.buttons)))
        for btnPos, btn in enumerate(self.buttons):
            for jPos in btn:
                joltages_affected_by_btns[jPos][btnPos] = 1

        # Constraint targets: we need to reach these exact joltages
        min_joltage_results = np.array(self.joltage)
        max_joltage_results = min_joltage_results

        constraints = LinearConstraint(joltages_affected_by_btns, min_joltage_results, max_joltage_results)

        result = milp(c=variablesToMinimize, constraints=constraints, integrality=all_integers)
        if not result.success:
            raise ValueError(result.message)
        return sum(result.x)
    
def solve_part_2(lines):
    machines = (machine(ln) for ln in lines)
    return sum(m.fewest_presses_for_joltage() for m in machines)

if __name__ == "__main__":
    c = np.array([1, 1, 1, 1, 1, 1])

    A = np.array([
        [0, 0, 0, 0, 1, 1], 
        [0, 1, 0, 0, 0, 1], 
        [0, 0, 1, 1, 1, 0],
        [1, 1, 0, 1, 0, 0],
    ])
    b_u = np.array([3, 5, 4, 7])
    b_l = b_u


    constraints = LinearConstraint(A, b_l, b_u)
    integrality = np.ones_like(c)


    res = milp(c=c, constraints=constraints, integrality=integrality)
    print(res.x)