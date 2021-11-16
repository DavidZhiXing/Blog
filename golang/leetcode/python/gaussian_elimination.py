"""
Gaussian elimination method for sovling a system linear equations.
Given a matrix of coefficients, the method will solve the system of linear equations.
https://en.wikipedia.org/wiki/Gaussian_elimination

"""

import numpy as np

def gaussian_elimination(matrix):
    """
    Solve a system of linear equations using Gaussian elimination method.
    :param matrix: Matrix of coefficients.
    :return: Solution of the system.
    """
    n = len(matrix)
    for i in range(n):
        # Find the maximum element in the column
        max_row = i
        for j in range(i+1, n):
            if abs(matrix[j][i]) > abs(matrix[max_row][i]):
                max_row = j
        # Swap the row
        matrix[i], matrix[max_row] = matrix[max_row], matrix[i]
        # Eliminate the element in the column
        for j in range(i+1, n):
            c = matrix[j][i] / matrix[i][i]
            for k in range(i, n+1):
                matrix[j][k] -= c * matrix[i][k]
    # Back substitution
    x = [0] * n
    for i in range(n-1, -1, -1):
        x[i] = matrix[i][n] / matrix[i][i]
        for j in range(i-1, -1, -1):
            matrix[j][n] -= matrix[j][i] * x[i]
    return x

def retroactive_resolution(coffficients: np.matrix, vector: np.ndarray) -> np.ndarray:
    """
    Solve a system of linear equations using retroactive resolution method.
    :param coffficients: Matrix of coefficients.
    :param vector: Vector of constants.
    :return: Solution of the system.
    """
    n = len(coffficients)
    x = [0] * n
    for i in range(n-1, -1, -1):
        x[i] = vector[i]
        for j in range(i+1, n):
            x[i] -= coffficients[i, j] * x[j]
        x[i] /= coffficients[i, i]
    return x

if __name__ == "__main__":
    # Test
    matrix = np.matrix([[1, 2, 3], [4, 5, 6], [7, 8, 9]])
    print(gaussian_elimination(matrix))