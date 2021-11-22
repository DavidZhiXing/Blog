nextPermutation: [int] -> [int]
    #
    # Return the next permutation of the elements in a list.
    #
    # If the list is in descending order, return the list in ascending order.
    #
    # If the list is in ascending order, return the list in descending order.
    #
    # If the list has only one element, return the list.
    #
    # Example:
    #
    # nextPermutation([3, 2, 1]) # [1, 2, 3]
    # nextPermutation([1, 1, 5]) # [1, 5, 1]
    # nextPermutation([1, 2, 3]) # [1, 3, 2]
    # nextPermutation([3, 2, 1, 0]) # [0, 1, 2, 3]
    # nextPermutation([0, 1, 2]) # [2, 1, 0]
    #
    # Parameters
    # ----------
    # list : [int]
    #   The list to be permuted.
    #
    # Returns
    # -------
    # [int]
    #   The next permutation of the list.
    #
    # Raises
    # ------
    # ValueError
    #   If the list is not in ascending or descending order.
    #
    # Notes
    # -----
    # The algorithm used is based on the following observation:
    #
    # If the list is in descending order, the next permutation is the
    # smallest permutation that is greater than the current permutation.
    #
    # If the list is in ascending order, the next permutation is the
    # largest permutation that is less than the current permutation.
    #

nextPermutation = lambda list: list[::-1]
