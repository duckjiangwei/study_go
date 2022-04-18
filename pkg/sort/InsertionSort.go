package sort

//插入排序
func InsertionSort(s []int) {
	length := len(s)
	if length <= 1 {
		return
	}
	//第一个元素是已排区域，从第二个元素开始
	for i := 1; i < length; i++ {
		value := s[i]
		j := i - 1
		//升序排列，从已排区域的右侧开始比对
		for ; j >= 0; j-- {
			if s[j] > value {
				s[j+1] = s[j] //右移1位
			} else {
				break
			}
		}
		//插入空出的位置
		s[j+1] = value
	}
}
