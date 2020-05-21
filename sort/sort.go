//copyright by jck 2020
package sort

import (
	"goutils/str"
	"reflect"
	"sort"
)

type SortType struct {
	intV int64
	strV string
	floatV float64
	complexV complex128
}


/*
	比较对象的 string,int,float  字段值
	slice 保存对象的slice
	field 根据对象的哪个字段排序
	asc   true 升序 false 降序
 */
func SortObjSlice(slice []struct{},field string,asc bool) {
	if len(slice) <=1 { return }
	structField, b := reflect.TypeOf(slice[0]).FieldByName(field)
	if !b {
		panic("要排序的字段不存在")
	}
	Type := structField.Type.String()
	sort.Slice(slice, func(i, j int) bool {
		var si,sj SortType
		switch  {
		case str.SameNthFront( Type,"int",3):
			si.intV = reflect.ValueOf(slice[i]).FieldByName(field).Int()
			sj.intV = reflect.ValueOf(slice[j]).FieldByName(field).Int()
		case Type == "string" :
			si.strV = reflect.ValueOf(slice[i]).FieldByName(field).String()
			sj.strV = reflect.ValueOf(slice[j]).FieldByName(field).String()
		case str.SameNthFront(Type,"float",5):
			si.floatV = reflect.ValueOf(slice[i]).FieldByName(field).Float()
			sj.floatV = reflect.ValueOf(slice[j]).FieldByName(field).Float()
		case str.SameNthFront(Type, "complex", 7):
			si.complexV = reflect.ValueOf(slice[i]).FieldByName(field).Complex()
			sj.complexV = reflect.ValueOf(slice[j]).FieldByName(field).Complex()
		}

		if asc {
			return true
		}else{
			return false
		}
	})
}
