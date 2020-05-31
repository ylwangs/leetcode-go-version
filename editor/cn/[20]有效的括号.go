//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。 
//
// 有效字符串需满足： 
//
// 
// 左括号必须用相同类型的右括号闭合。 
// 左括号必须以正确的顺序闭合。 
// 
//
// 注意空字符串可被认为是有效字符串。 
//
// 示例 1: 
//
// 输入: "()"
//输出: true
// 
//
// 示例 2: 
//
// 输入: "()[]{}"
//输出: true
// 
//
// 示例 3: 
//
// 输入: "(]"
//输出: false
// 
//
// 示例 4: 
//
// 输入: "([)]"
//输出: false
// 
//
// 示例 5: 
//
// 输入: "{[]}"
//输出: true 
// Related Topics 栈 字符串

package cn


//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	if s == "" {
		return true
	}

	cMap := map[byte]byte{')':'(','}':'{',']':'['}

	// 定义一个栈通过字符匹配判定
	stack := make([]byte, 0, 1)
	for _, v := range s {
		if byte(v) == ' ' {
			continue
		}
		switch byte(v) {
		case '(', '{', '[':
			stack = append(stack, byte(v))
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack) - 1] != cMap[byte(v)] {
				return false
			}
			stack = stack[:len(stack) - 1]
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
