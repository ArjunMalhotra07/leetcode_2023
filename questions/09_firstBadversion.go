package questions

func IsBadVersion(version int) bool {
	return false
}

func FirstBadVersion(n int) int {
	s := 0
	e := n - 1
	for s <= e {
		mid := s + ((e - s) / 2)
		if IsBadVersion(mid) == true {
			e = mid - 1
			if IsBadVersion(mid-1) == false {
				return mid
			}
		} else {
			s = mid + 1
			if IsBadVersion(mid+1) == true {
				return mid + 1
			}
		}
	}
	return 1
}
