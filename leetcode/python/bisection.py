from typing import Callable, List, Tuple

"""
Arithemetic analysis is a branch of mathematics that deals with solving linear equations.
"""

def bisection(f: Callable[[float], float], a: float, b: float, eps: float) -> Tuple[float, float]:
    """
    Finds the root of a function using bisection method.
    :param f: The function to find the root of.
    :param a: The left bound of the interval.
    :param b: The right bound of the interval.
    :param eps: The precision of the root.
    :return: The root of the function.
    """
    while abs(b - a) > eps:
        c = (a + b) / 2
        if f(c) == 0:
            return c
        elif f(c) * f(a) < 0:
            b = c
        else:
            a = c
    return (a + b) / 2

def bisection_multiple(f: Callable[[List[float]], float], a: List[float], b: List[float], eps: float) -> List[float]:
    """
    Finds the root of a function using bisection method.
    :param f: The function to find the root of.
    :param a: The left bound of the interval.
    :param b: The right bound of the interval.
    :param eps: The precision of the root.
    :return: The root of the function.
    """
    while abs(b - a) > eps:
        c = [(a[i] + b[i]) / 2 for i in range(len(a))]
        if f(c) == 0:
            return c
        elif f(c) * f(a) < 0:
            b = c
        else:
            a = c
    return [(a[i] + b[i]) / 2 for i in range(len(a))]

if __name__ == "__main__":
    print(bisection(lambda x: x ** 2 - 1, 0, 1, 1e-5))
    print(bisection_multiple(lambda x: x[0] ** 2 + x[1] ** 2 - 1, [0, 0], [1, 1], 1e-5))

    import doctest
    doctest.testmod()