
def balanced_parentheses(parentheses: str)->bool:
    """
    :param parentheses: string of parentheses
    :return: True if parentheses are balanced, False otherwise
    """
    if len(parentheses) % 2 != 0:
        return False
    stack = []
    for p in parentheses:
        if p == '(':
            stack.append(p)
        elif p == ')':
            if len(stack) == 0:
                return False
            stack.pop()
    return len(stack) == 0

# Test
print(balanced_parentheses('()'))
print(balanced_parentheses('(()'))
print(balanced_parentheses('()()'))
print(balanced_parentheses('(()()'))
print(balanced_parentheses('(()()()'))
print(balanced_parentheses('(()()()()'))