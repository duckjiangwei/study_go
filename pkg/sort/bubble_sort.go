package sort

//冒泡排序-升序
func BubbleSort(data []int) []int {
	length := len(data)
	if length <= 1 {
		return data
	}
	for i := 0; i < length; i++ {
		breakTag := false
		for j := 0; j < length-1; j++ {
			if data[j] > data[j+1] {
				//交换位置
				tmp := data[j+1]
				data[j+1] = data[j]
				data[j] = tmp
				breakTag = true
			}
		}
		if !breakTag {
			break
		}
	}
	return data
}
