#include <stdio.h>

void swap(int& a, int& b) {
    int tmp = a;
    a = b;
    b = tmp;
}

void reverse(int* nums, int n) {
    int i = 0;
    int j = n - 1;
    while (i < j) {
        swap(nums[i], nums[j]);
        i++;
        j--;
    }
}

void nextPermutation(int* nums, int n) {
    int i = n - 2;
    while (i >= 0 && nums[i] >= nums[i + 1]) {
        i--;
    }
    if (i >= 0) {
        int j = n - 1;
        while (j >= 0 && nums[j] <= nums[i]) {
            j--;
        }
        swap(nums[i], nums[j]);
    }
    reverse(nums + i + 1, n - i - 1);
}

void test_nextPermutation() {
    int nums[] = {1, 2, 3};
    nextPermutation(nums, 3);
    for (int i = 0; i < 3; i++) {
        printf("%d ", nums[i]);
    }
    printf("\n");
}

int main() {
    test_nextPermutation();
}