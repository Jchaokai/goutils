package sort_test

import (
	"fmt"
	"goutils/sort"
	"testing"
)

func TestSortObjSlice(t *testing.T) {
	var s []interface{}
	for i := 0; i < 5; i++ {
		obj1 := struct {
			intV   int64
			strV   string
			floatV float64
		}{intV: int64(i + 100), strV: string(97 + i), floatV: 9.9 + float64(i)}
		s = append(s, obj1)
	}
	fmt.Println(s)
	sort.SortObjSlice(s, "strV", true)
	fmt.Println(s)
}

func ss() {

}
