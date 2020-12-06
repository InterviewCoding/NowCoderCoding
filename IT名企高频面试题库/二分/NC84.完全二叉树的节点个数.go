package main

/**
 * @Author: yirufeng
 * @Date: 2020/12/6 5:15 下午
 * @Desc:
 **/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * ✅ 【推荐】方法一：每次判断是否为满二叉树
 * @param head TreeNode类
 * @return int整型
 */
func nodeNum(head *TreeNode) int {
	// write code here
	if head == nil {
		return 0
	}

	hl, hr := 0, 0
	l, r := head.Left, head.Right
	for l != nil {
		hl++
		l = l.Left
	}

	for r != nil {
		hr++
		r = r.Right
	}

	//说明是一颗满二叉树
	if hl == hr {
		return 1<<(hl+1) - 1
	}

	return 1 + nodeNum(head.Left) + nodeNum(head.Right)
}

// ✅ 方法二：每次判断左右子树哪一个是满二叉树
func nodeNum(head *TreeNode) int {
	if head == nil {
		return 0
	}

	//获得左右节点的高度
	hl, hr := getHeight(head.Left), getHeight(head.Right)

	//如果高度相同，说明左子树是一颗满二叉树
	if hl == hr {
		return 1<<hl + nodeNum(head.Right)
	}

	return 1<<hr + nodeNum(head.Left)
}

func getHeight(head *TreeNode) int {
	if head == nil {
		return 0
	}

	return 1 + max(getHeight(head.Left), getHeight(head.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
