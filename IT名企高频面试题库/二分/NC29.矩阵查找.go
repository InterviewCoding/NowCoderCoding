package main

/**
 * @Author: yirufeng
 * @Date: 2020/12/7 11:32 上午
 * @Desc:
 **/

/**
 * ✅
 * 方法一：暴力
 * 方法二：每行或每列进行二分
 * 方法三：从右上角或者左下角不断缩减行或列
 * 方法四：每次找到矩阵的中心点，然后划分为4个小矩阵，每次排除一个小矩阵
 * 方法五【当前方法】：因为多加了一个条件，所以整个矩阵可以按行遍历之后组成一个有序数组，从而直接使用标准的二分去查找即可
 * 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗：2.7 MB, 在所有 Go 提交中击败了75.17%的用户
 * @param matrix int整型二维数组
 * @param target int整型
 * @return bool布尔型
 */
func searchMatrix(matrix [][]int, target int) bool {
	// write code here

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row, col := len(matrix), len(matrix[0])
	l, r := 0, row*col-1
	for l <= r {
		mid := l + (r-l)>>1
		if matrix[mid/col][mid%col] > target { //说明
			r = mid - 1
		} else if matrix[mid/col][mid%col] < target {
			l = mid + 1
		} else if matrix[mid/col][mid%col] == target {
			return true
		}
	}

	//说明没找到
	return false
}
