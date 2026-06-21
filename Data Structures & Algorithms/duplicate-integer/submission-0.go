func hasDuplicate(nums []int) bool {
    hashmap := make(map[int]int)
    for _, num := range nums {
        hashmap[num]++
        if hashmap[num]>1 {
            return true
        }

    }
    return false
}
