func merge(nums1 []int, m int, nums2 []int, n int) {
	write := m+n-1
	left, right := m-1, n-1
	for left>=0 && right>=0 {
		if nums1[left] > nums2[right] {
			nums1[write] = nums1[left]
			left--
		} else {
			nums1[write] = nums2[right]
			right--
		}
		write--
	}
	for left >= 0 {
		nums1[write] = nums1[left]
		write--
		left--
	}
	for right >=0 {
		nums1[write] = nums2[right]
		write--
		right--
	}
}
