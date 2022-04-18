package sort

//选择排序
func SelectionSort(s []int) {
	length := len(s)
	if length <= 1 {
		return
	}
	for i := 0; i < length; i++ {
		// 查找最小值
		minIndex := i
		for j := i + 1; j < length; j++ {
			if s[j] < s[minIndex] {
				minIndex = j
			}
		}
		// 交换
		s[i], s[minIndex] = s[minIndex], s[i]
	}
}
