package main

// "strconv"
// "time"
// "math"

func main() {
	// // fmt.Println(countAndSay(5))
	// a := 3
	// tm := fmt.Sprintf("%b", a)
	// fmt.Println(time.Now().UnixNano())
	// nums1 := [1,2,2,1]
	// nums2

	// fmt.Println(intersect(,[2,2]))

}

func intersect(nums1 []int, nums2 []int) (nums []int) {
	var hash map[int]int = make(map[int]int)
	for _, v := range nums1 {
		hash[v] += 1
	}
	for _, v := range nums2 {
		sult, err := hash[v]
		if err && sult > 0 {
			hash[v] -= 1
			nums = append(nums, v)
		}
	}
	return
}
