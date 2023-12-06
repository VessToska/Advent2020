package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

func tcast(in_slice []string) []int {
	out_slice := make([]int, len(in_slice))
	for temp_idx, temp_string := range in_slice {
		int_data, int_error := strconv.Atoi(temp_string)
		tcheck(int_error)
		out_slice[temp_idx] = int_data
	}
	return out_slice
}

func tcheck(in_error error) {
	if in_error != nil {
		panic(in_error)
	}
}

func tcopy[T any](in_var []T) []T {
	return slices.Clone(in_var)
}

func tcountdigit(in_int int) int {
	if in_int > 0 {
		return int(math.Log10(float64(in_int)))+1
	} else if in_int == 0 {
		return 1
	}
	return int(math.Log10(float64(-in_int)))+2
}

func tdigit(in_string string) bool {
	switch in_string {
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9": return true 
		default: return false
	}
}

func tdrop[T any](in_list []T, in_index int) []T {
	return append(in_list[:in_index], in_list[in_index+1:]...)
}

func tminidx(in_list []int) int {
	int_min := math.MaxInt
	int_idx := -1
	for temp_idx, temp_val := range in_list {
		if temp_val < int_min {
			int_min = temp_val
			int_idx = temp_idx
		}
	}
	return int_idx
}

func tconcatint(in_x int, in_y int) int {
	int_digits := tcountdigit(in_y)
	int_new := in_x * tpow(10, int_digits) + in_y
	return int_new
}

func tconcatgroup(in_ints []int) int {
	final_int := in_ints[0]
	for _, temp_int := range in_ints[1:] {
		final_int = tconcatint(final_int, temp_int)
	}
	return final_int
}

func tload(in_path string) string {
	file_data, file_error := os.ReadFile(in_path)
	tcheck(file_error)
	return string(file_data)
}

func tmake(in_fill any, in_len ...int) any {
	var slice_final []any
	if len(in_len) == 1 {
		for temp_itr := 1; temp_itr <= in_len[0]; temp_itr++ {
			slice_final = append(slice_final, in_fill)
		}
	} else {
		for temp_itr := 1; temp_itr <= in_len[0]; temp_itr++ {
			slice_next := tmake(in_fill, in_len[1:]...)
			slice_final = append(slice_final, slice_next)
		}
	}
	return slice_final
}

func tpow(in_num int, in_exp int) int {
	if in_exp == 0 { return 1 }
	if in_exp == 1 { return in_num }
	var_y := tpow(in_num, in_exp / 2)
	if in_exp % 2 == 0 { return var_y * var_y }
	return in_num * var_y * var_y
}

func tprint(in_list ...any) {
	if len(in_list) == 0 {
		fmt.Println()
	} else {
		for _, temp_var := range in_list {
			fmt.Printf("%#v\n", temp_var)
		}
	}
}

func tuse(in_list ...any) {
	for _, temp_var := range in_list {
		_ = temp_var
	}
}