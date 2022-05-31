package algorithm

import (
	"fmt"
	"os"
	"reflect"

	utils "strconv"
)

// SearchInSlice func
func SearchInSlice(slice, searchValue interface{}, fieldName string) int {
	var idx = -1

	ref := reflect.ValueOf(slice)
	if ref.Kind() == reflect.Ptr {
		ref = reflect.ValueOf(slice).Elem()
	}
	if ref.Kind() != reflect.Slice {
		return idx
	}

	n := ref.Len()
	first, last := 0, n-1
	for first <= last {
		mid := (first + last) / 2

		var valueInData int
		var err error
		isExist := false
		for i := 0; i < ref.Index(mid).NumField(); i++ {
			key := ref.Index(mid).Type().Field(i).Name
			if key == fieldName {
				valueI := ref.Index(mid).Field(i).Interface()
				valueInData, err  = utils.Atoi(fmt.Sprint(valueI))
				if err != nil {
					fmt.Println(err)
					os.Exit(99)
				}
				isExist = true
				break
			}
		}
		if !isExist {
			return -1
		}

		inSearch, err := utils.Atoi(fmt.Sprint(searchValue))
		if err != nil {
			fmt.Println(err)
			os.Exit(98)
		}

		if inSearch > valueInData {
			first = mid + 1
		} else if inSearch < valueInData {
			last = mid - 1
		} else {
			idx = mid
			break
		}
	}

	return idx
}
