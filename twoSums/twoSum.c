#include <stdlib.h>
#include <stdio.h>

/*
* Note: The returned array must be malloced, assume caller calls free().
*/

int* twoSum(int* nums, int numsSize, int target) {
	int* p;
	p = nums;
	int i, j;

	int *res = NULL;
	for (i = 0; i < numsSize; i++) {
		for(j = i+1; j < numsSize; j++) {
			if ((*(p + i) + *(p + j)) == target) {
				goto lable;
			}
		}
	}

	lable:res = (int *)malloc(2*sizeof(int));
	*res = i;
	*(res+1) = j;
	return res;
}

int main() {
	int size = 10;
	int *tests = NULL;
	tests = (int *)malloc(size*sizeof(int));
	int i = 0;
	for (i = 0; i < size; i++) {
		*(tests+i) = i;
	}

	int *p;
	p = twoSum(tests, size, 3);
	for (i = 0; i < 2; i++) {
		printf("%d ", *(p+i));
	}
	return 0;
}