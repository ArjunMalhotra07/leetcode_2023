package main

func isBadVersion(version int) bool

func firstBadVersion(n int) int {
	s := 0
	e := n - 1
	for s <= e {
		mid := s + ((e - s) / 2)
		if isBadVersion(mid) == true {
			e = mid - 1
			if isBadVersion(mid-1) == false {
				return mid
			}
		} else {
			s = mid + 1
			if isBadVersion(mid+1) == true {
				return mid + 1
			}
		}
	}
	return 1
}
