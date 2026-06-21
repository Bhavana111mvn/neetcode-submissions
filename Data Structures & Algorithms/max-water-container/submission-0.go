func maxArea(heights []int) int {
    n := len(heights)
	left, right := 0, n-1
	product :=0
    for left < right {
		if heights[left] < heights[right] {
			if (heights[left])*(right-left) > product {
				product =  (heights[left])*(right-left)
			}
			left++
		} else {
			if (heights[right])*(right-left)> product {
				product =  (heights[right])*(right-left)
			}
			right--
		}
	}
	return product
}
