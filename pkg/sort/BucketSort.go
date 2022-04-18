package sort

//桶排序
func BucketSort(s []int) {
	length := len(s)
	if length <= 1 {
		return
	}

	//找出最大数
	max := s[0]
	for i := 1; i < length; i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	//生成二维的切片
	buckets := make([][]int, max)

	//将元素入桶
	index := 0
	for i := 0; i < length; i++ {
		// index := 0
		if s[i] < max {
			index = s[i]%max - 1 // 桶序号
		} else {
			index = max - 1
		}
		// index = s[i] * (length - 1) / max // 桶序号
		buckets[index] = append(buckets[index], s[i])
	}
	tmpIndex := 0 //初始位置
	//循环，对每个桶的元素进行排序
	for i := 0; i < max; i++ {
		BubbleSort(buckets[i]) //每个桶的排序-这里使用冒泡
		copy(s[tmpIndex:], buckets[i])
		tmpIndex += len(buckets[i])
	}
}
