package main

import "fmt"

/**
 * @Author: yirufeng
 * @Date: 2020/12/6 2:22 下午
 * @Desc:
 **/

/**
 * ✅
 * 二分查找
 * @param n int整型 数组长度
 * @param v int整型 查找值
 * @param a int整型一维数组 有序数组
 * @return int整型
 * 思路：获取左边界，如果左边界对应位置的元素是我们的v，返回对应的索引，如果不是v我们直接返回-1
 */
func upper_bound_(n int, v int, a []int) int {
	// write code here
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)>>1
		if a[mid] > v {
			r = mid - 1
		} else if a[mid] < v {
			l = mid + 1
		} else if a[mid] == v {
			//因为是要找左侧的元素
			r = mid - 1
		}
	}

	//循环结束：此时有l = r + 1
	//如果此时l没有越界并且指向位置的值大于等于目标值，就返回l，否则返回数组长度+1
	//这里其实就包含了两种情况，一种是目标值不在另外一种是目标值在里面，并且这两种情况下一定有大于等于查找值的位置(也就是说查找值的位置不会越界)
	if l == n {  //说明肯定不存在大于等于目标值的数
		return n + 1
	}

	//注意：因为题目中说明了输出位置从1开始
	return l + 1
}

func main() {
	fmt.Println(upper_bound_(5, 4, []int{1, 2, 4, 4, 5}))
}
