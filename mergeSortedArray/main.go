package main

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	// make a copy of the first array
// 	// we need the copy because we are
// 	// going to overrite the first nums1
// 	// instead of trying to shift the elements over (which is costly)
// 	nums1Copy := make([]int, m)
// 	for i := 0; i < m; i++ {
// 		nums1Copy[i] = nums1[i]
// 	}

// 	// keep 3 pointers (p is in the for loop)
// 	// we need 3 because we now have 3 slices (a copy of one)
// 	p1 := 0
// 	p2 := 0

// 	// lets loop through, stopping when the total length of
// 	// both are reached.
// 	for p := 0; p < m+n; p++ {
// 		// if p2 is greater then n (the size of nums2) then we've
// 		// exhausted nums2 but still have nums1 values (because)
// 		// we have m+n length.
// 		// we then compare the current values of the copy and
// 		// the second array but also have to check that p1 is no over
// 		// its own array length (p1) otherwise you'd get and overflow
// 		if p2 >= n || (p1 < m && nums1Copy[p1] < nums2[p2]) {
// 			// if nums1copy at p1 is less the nums2 at p2 then
// 			// we overrite nums1 at p with the copy value at p1
// 			// we then move p1 up
// 			nums1[p] = nums1Copy[p1]
// 			p1++
// 		} else {
// 			// otherwise we overite nums1 at p with nums2 at p2
// 			// and move p2 up
// 			nums1[p] = nums2[p2]
// 			p2++
// 		}
// 	}
// }

func merge(nums1 []int, m int, nums2 []int, n int) {
	// copy and three pointers
	// from here on out .. nums1 is just our result set
	// think of it as recycling an array thats already allocated
	n1Copy := make([]int, m)

	for i := 0; i < m; i++ {
		n1Copy[i] = nums1[i]
	}

	p1 := 0
	p2 := 0

	for i := 0; i < m+n; i++ {
		if p2 >= n || (p1 < m && n1Copy[p1] < nums2[p2]) {
			nums1[i] = n1Copy[p1]
			p1++
		} else {
			nums1[i] = nums2[p2]
			p2++
		}
	}
}
