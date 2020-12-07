package main

/**
 * @Author: yirufeng
 * @Date: 2020/12/7 11:31 上午
 * @Desc:
 **/

/**
 * ✅
 * 思路：每次从左上角或者右下角开始搜索
 * @param mat int整型二维数组
 * @param n int整型
 * @param m int整型
 * @param x int整型
 * @return int整型一维数组
 */
func findElement(mat [][]int, n int, m int, x int) []int {
	// write code here

	//此时右上角对应的坐标
	curX, curY := 0, m-1
	for curX < n && curY > -1  {
		if mat[curX][curY] > x { //说明不在该列
			curY--
		} else if mat[curX][curY] < x { //说明不在该行
			curX++
		} else if mat[curX][curY] == x {
			return []int{curX, curY}
		}
	}

	//说明找不到
	return []int{-1, -1}
}
