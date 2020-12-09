package main

/**
 * @Author: yirufeng
 * @Date: 2020/12/9 9:59 上午
 * @Desc: ✅
 **/

/**
 *
 * @param A int整型一维数组
 * @param target int整型
 * @return int整型
 */
func search(A []int, target int) int {
	// write code here
	if len(A) == 0 {
		return -1
	}

	l, r := 0, len(A)-1
	for l <= r {
		mid := l + (r-l)>>1
		if target == A[mid] {
			return mid
		}

		if A[mid] > A[r] { //说明此时[l, mid]是单调递增的
			//if target > nums[mid] {
			//	l = mid + 1
			//} else if target < nums[mid] {
			//	if target >= nums[l] {
			//		r = mid - 1
			//	} else if target < nums[l] {
			//		l = mid + 1
			//	}
			//}
			if target < A[mid] && target >= A[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if A[mid] < A[r] {

			if target > A[mid] && target <= A[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}

		} else if A[mid] == A[r] { //关键点：不同于33题的
			//因为数组中有重复元素，我们不能确定l, mid, r是指向同一个元素的
			//因此让r不断缩减
			r--
		}
	}

	//最后一定不会执行到这里
	return -1
}
