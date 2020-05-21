package slice

func Same(s1, s2 []interface{}) bool {
	if len(s1) != len(s2) {
		return false
	}else{
		for i := 0; i < len(s1); i++{
		    if s1[i] != s2 [i] {
		    	return false
			}
		}
		return  true
	}
}
