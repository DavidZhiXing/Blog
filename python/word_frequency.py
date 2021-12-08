"""
    create a word frequency list from a text file
"""
def word_frequency(filename):
    """
        create a word frequency list from a text file
    """
    with open(filename, 'r') as f:
        words = f.read().split()
    d = {}
    for word in words:
        if word in d:
            d[word] += 1
        else:
            d[word] = 1
    return d

def main():
    """
        main function
    """
    print(word_frequency('test.txt'))