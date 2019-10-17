package main


import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"reflect"
)
type SA struct {
	Names []string
	Nick string
}
type SB struct {
	Names []string
}

func main()  {
	CopyStructToMap()
}

func TestCopySlice()  {
	src := SA{}
	//src.Names = []string{}
	dst := SB{}
	IterateStruct(&src)
	copier.Copy(&dst, &src)
	data, _ := json.Marshal(dst)
	fmt.Println(string(data))
}

func CopyStructToMap()  {
	type Student struct {
		Name string
		Age int
	}
	s := Student{Name:"ha", Age:5}
	d := map[string]interface{}{}
	copier.Copy(&d, &s)
	fmt.Println(d)
}

//func Fill(v SA)  {
//	objT := reflect.TypeOf(v)
//	objV := reflect.ValueOf(v)
//}

func IterateStruct(v interface{})  {
	objT := reflect.TypeOf(v)
	objV := reflect.ValueOf(v)
	objKind := objT.Kind()

	fmt.Println(objKind)
	fmt.Println(objV.CanAddr())
	if objKind == reflect.Struct {
		filedCount := objV.NumField()
		for x := 0; x < filedCount; x++ {
			subObjV := objV.Field(x)
			subObjKind := subObjV.Kind()
			fmt.Println(subObjKind)
			fmt.Println(subObjV.CanAddr())
			continue
			if subObjKind == reflect.Slice {
				if subObjV.Len() == 0 {
					fmt.Println("edwin 40 ", subObjV.CanSet())
				}
				return
			} else if subObjKind == reflect.Struct {
				IterateStruct(subObjV)
			}
		}
	}
}


//
//func main() {
//	type t struct {
//		N int
//	}
//	var n = t{42}
//	// N at start
//	fmt.Println(n.N)
//	// pointer to struct - addressable
//	ps := reflect.ValueOf(&n)
//	// struct
//	s := ps.Elem()
//	if s.Kind() == reflect.Struct {
//		// exported field
//		f := s.FieldByName("N")
//		if f.IsValid() {
//			// A Value can be changed only if it is
//			// addressable and was not obtained by
//			// the use of unexported struct fields.
//			if f.CanSet() {
//				// change value of N
//				if f.Kind() == reflect.Int {
//					x := int64(7)
//					if !f.OverflowInt(x) {
//						f.SetInt(x)
//					}
//				}
//			}
//		}
//	}
//	// N at end
//	fmt.Println(n.N)
//}
