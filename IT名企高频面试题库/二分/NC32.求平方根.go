package main

/**
 * @Author: yirufeng
 * @Date: 2020/12/6 5:21 下午
 * @Desc:
 **/

/**
 * ✅代码
 * @param x int整型
 * @return int整型
 */
func sqrt(x int) int {
	// write code here
	//一个数的平方根肯定不会超过它自己，不过直觉还告诉我们，一个数的平方根最多不会超过它的一半
	//l, r := 1, x
	//➡️针对上面代码的优化：
	l, r := 1, x>>1+1 //使用x>>1+1是为了普及一下x=1的情况
	for l <= r {
		mid := l + (r-l)>>1
		if mid*mid > x {
			r = mid - 1
		} else if mid*mid < x {
			l = mid + 1
		} else if mid*mid == x {
			return mid
		}
	}

	//1, 3, 4 我们要找2的平方根
	//最后循环结束的时候有l = r + 1，并且一定是执行了r = mid - 1 或者 l = mid + 1最后结束了循环。
	//在执行这一步之前一定有l = r - 1，根据此时满足的情况分别执行对应的一步造成l > r，从而我们r指向的数一定是平方根的下取整(或者)，但是l指向的数一定是平方根的上取整数
	return r
}

func main() {

}
