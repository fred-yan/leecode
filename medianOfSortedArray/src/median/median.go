package median

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		nums1, nums2 = nums2, nums2
		m, n = n, m
	}

	iMin := 0
	iMax := m
	halfLen := (m + n + 1) / 2
	for iMin < iMax {
		i := (iMin + iMax) / 2
		// when halfLen is maintain,  i increase and then j decrease
		j := halfLen - i

		if i < iMax && nums2[j-1] > nums1[i] {
			// when nums2[j-1] > nums1[i], j need to decrease and i need to increase
			iMin = iMin + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			// when nums1[i-1] > nums2[j], i need to decrease and j need to increase
			iMax = iMax - 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = Max(nums1[i-1], nums2[j-1])
			}

			if (m+n)%2 == 1 {
				return maxLeft
			}

			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = Min(nums1[i], nums2[j])
			}

			return (maxLeft + minRight) / 2
		}
	}
	return 0
}
