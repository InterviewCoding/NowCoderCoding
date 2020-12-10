package main

/**
 * @Author: yirufeng
 * @Date: 2020/12/9 11:57 上午
 * @Desc: 验证给定的字符串是否符合括号的规则
 **/

/**
 *
 * @param s string字符串
 * @return bool布尔型
 */
func isValid(s string) bool {
	// write code here
	//如果字符串长度为奇数，直接返回false
	if len(s)&1 != 0 {
		return false
	}

	stack := []byte{}
	for i := 0; i < len(s); i++ {
		//如果是左括号就压入
		if s[i] == '[' || s[i] == '{' {
			stack = append(stack, stack[i])
		} else if s[i] == ']' { //如果是右括号就判断当前栈顶与其是否匹配，若不匹配直接返回false，

			//注意点：确保栈不会溢出
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else if s[i] == '}' {
			//注意点：确保栈不会溢出
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	//最后判断一下栈是否为空
	return len(stack) == 0
}

func main() {

}
