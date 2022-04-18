package sort

import "errors"

//二分查找
func BinarySearch(s []int, value int, low int, hight int) (int, error) {
	if low >= hight {
		return 0, errors.New("err")
	}
	mid := (low + hight) / 2
	midV := s[mid]
	if midV == value {
		return mid, nil
	} else if midV > value {
		return BinarySearch(s, value, low, mid)
	} else {
		return BinarySearch(s, value, mid, hight)
	}
}
