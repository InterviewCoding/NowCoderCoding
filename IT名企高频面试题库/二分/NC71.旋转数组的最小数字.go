package main

/**
 * @Author: yirufeng
 * @Date: 2020/12/7 8:57 上午
 * @Desc: https://www.nowcoder.com/practice/9f3231a991af4f55b95579b44b7a01ba?tpId=117&&tqId=34993&rp=1&ru=/ta/job-code-high&qru=/ta/job-code-high/question-ranking
 **/

/**
 * ✅
 * @param rotateArray int整型一维数组
 * @return int整型
 * 题目中并没有说数字是否是唯一的，因此可能出现重复
 * 解题思路：因为题目中没有明显的target，所以我们假设target为我们当前有效搜索区间上的最右侧元素
 * 			步骤1：每次判断当前有效搜索区间最左侧元素是否小于最右侧元素，若是，说明当前区间的最小值就是最左侧元素，
 *																	并且由于是当前的搜索区间所以直接返回最左侧元素
 * 			步骤2：如果当前有效区间的中间元素为mid，如果当前区间上的mid位置元素大于target，则说明旋转点在右侧，left = mid + 1
 * 			      如果当前有效区间的中间元素为mid，如果当前区间上的mid位置元素小于target，则说明右侧是区间递增，
																					旋转点在左侧(包含当前点)，right = mid
 * 			      如果当前有效区间的中间元素为mid，如果当前区间上的mid位置元素等于target，我们并不知道旋转点在左侧还是右侧，，
																						因此让last = last - 1收缩区间
 * 时间复杂度：O(logN) - O(N)  如果数组中的数字全部都一样，那么就是O(N)
   我们考虑数组中的最后一个元素 xx：在最小值右侧的元素，它们的值一定都小于等于 xx；而在最小值左侧的元素，它们的值一定都大于等于 xx。因此，我们可以根据这一条性质，通过二分查找的方法找出最小值。
*/
func minNumberInRotateArray(rotateArray []int) int {
	// write code here

	//数组大小为0，请返回0
	if len(rotateArray) == 0 {
		return 0
	}

	l, r := 0, len(rotateArray)-1
	for l <= r {
		if rotateArray[l] < rotateArray[r] {
			return rotateArray[l]
		}

		mid := l + (r - l) >> 1
		if rotateArray[mid] > rotateArray[r] {
			l = mid + 1
		} else if rotateArray[mid] < rotateArray[r] {
			r = mid
		} else if rotateArray[mid] == rotateArray[r] {
			r = r - 1
		}
	}


	//最后结束时一定首先有rotateArray[mid] == rotateArray[r]
	//然后r = r - 1，从而l还是指向那个rotateArray[mid] == rotateArray[r]时mid指向的元素
	return rotateArray[l]
}

func main() {

}
