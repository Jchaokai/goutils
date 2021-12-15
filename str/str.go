package str

// SameNthFront 两个字符串的前几位是否一样
func SameNthFront(one string, another string, n int) bool {
	for i := 0; i < n; i++ {
		if one[i] != another[i] {
			return false
		}
	}
	return true
}

// ChildOfStr 判断一个字符串是否是另一个的子集
//TODO have bugs
func ChildOfStr(child string, str string) bool {
	for i := 0; i < len(child); i++ {
		if child[i] != str[i] {
			return false
		}
	}
	return true
}

// SpaceStr 返回len长度的空格str
func SpaceStr(len int) (res string) {
	for i := 0; i < len; i++ {
		res += " "
	}
	return
}
