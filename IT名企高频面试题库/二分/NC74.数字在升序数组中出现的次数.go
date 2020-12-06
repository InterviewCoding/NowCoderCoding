package main

import "fmt"

/**
 * @Author: yirufeng
 * @Date: 2020/12/6 2:37 下午
 * @Desc:
 **/

/**
 *
 * @param data int整型一维数组
 * @param k int整型
 * @return int整型
思路：求出数字在数组中出现的第一个位置的下标以及最后一个位置的下标，如果没有出现则返回-1
✅
*/
func GetNumberOfK(data []int, k int) int {
	// write code here
	//如果有一个没有出现直接返回0即可
	l, r := GetLeftBounder(data, k), GetRightBounder(data, k)

	//注意点：这里如果没有出现，是返回0，不是返回r - l + 1
	if l == -1 || r == -1 {
		return 0
	}
	return r - l + 1
}

//获取目标值左边界的对应位置
func GetLeftBounder(data []int, k int) int {
	l, r := 0, len(data)-1
	for l <= r {
		mid := l + (r-l)>>1
		if data[mid] > k {
			r = mid - 1
		} else if data[mid] < k {
			l = mid + 1
		} else if data[mid] == k {
			r = mid - 1
		}
	}

	//最后循环结束一定有 l = r + 1，此时l指向第一个出现目标值的元素或者目标值没有出现时第一个大于目标值的元素，或者在len(data)位置
	if l < len(data) && data[l] == k {
		return l //说明有出现
	}

	//说明没有出现
	return -1
}

//获取目标值右边界的对应位置
func GetRightBounder(data []int, k int) int {
	l, r := 0, len(data)-1
	for l <= r {
		mid := l + (r-l)>>1
		if data[mid] > k {
			r = mid - 1
		} else if data[mid] < k {
			l = mid + 1
		} else if data[mid] == k {
			l = mid + 1
		}
	}

	if r >= 0 && data[r] == k {
		return r
	}

	//说明没有出现
	return -1
}

func main() {
	fmt.Println(GetNumberOfK([]int{1, 2, 3, 3, 3, 3, 4, 5}, 3))
}
