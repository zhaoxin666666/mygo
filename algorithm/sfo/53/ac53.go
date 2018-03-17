package problem53

func getFirstK(numbers []int, k, start, end int) int {
	if start > end {
		return -1
	}

	middleIndex := (start + end) / 2
	middleData := numbers[middleIndex]
	if middleData == k {
		if (0 < middleIndex && numbers[middleIndex-1] != k) || 0 == middleIndex {
			return middleIndex
		} else {
			end = middleIndex - 1
		}
	} else if middleData > k {
		end = middleIndex - 1
	} else {
		start = middleIndex + 1
	}
	return getFirstK(numbers, k, start, end)
}

func getLastK(numbers []int, k, start, end int) int {
	if start > end {
		return -1
	}

	middleIndex := (start + end) / 2
	middleData := numbers[middleIndex]
	if middleData == k {
		if (len(numbers)-1 > middleIndex && numbers[middleIndex+1] != k) || len(numbers)-1 == middleIndex {
			return middleIndex
		} else {
			start = middleIndex + 1
		}
	} else if middleData < k {
		start = middleIndex + 1
	} else {
		end = middleIndex - 1
	}
	return getLastK(numbers, k, start, end)
}

func getNumberOfK(numbers []int, k int) int {
	if nil == numbers {
		return 0
	}
	first := getFirstK(numbers, k, 0, len(numbers)-1)
	last := getLastK(numbers, k, 0, len(numbers)-1)
	if first > -1 && last > -1 {
		return last - first + 1
	}
	return 0
}
