//
// Partial Sort
//

package partialsort

import (
	"math"
	"reflect"
	"sort"
	"unsafe"
)

func Partition(data sort.Interface, begin int, end int, mid int) int {
	for {
		// stop on element greater than mid
		for begin < mid && data.Less(begin, mid) == true {
			begin++
		}
		// stop on element less than mid
		for mid < end && data.Less(end, mid) == false {
			end--
		}
		if begin == end {
			return mid
		}
		data.Swap(begin, end)
		if begin == mid {
			mid = end
		} else {
			mid = begin
		}
	}
}

func Nth_elementBegin(data sort.Interface, begin int, end int, n int) {
	for begin < end {
		mid := begin + (end-begin)/2
		mid = Partition(data, begin, end, mid)
		if n < mid {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}
}

func Nth_element(data sort.Interface, n int) {
	if my_len := data.Len(); my_len > 1 {
		Nth_elementBegin(data, 0, my_len-1, n)
	}
}

func QuickSortPlainBegin(data sort.Interface, begin int, end int) {
	if begin < end {
		mid := begin + (end-begin)/2
		mid = Partition(data, begin, end, mid)
		QuickSortPlainBegin(data, begin, mid)
		QuickSortPlainBegin(data, mid+1, end)
	}
}

func QuickSortPlain(data sort.Interface) {
	if my_len := data.Len(); my_len > 1 {
		QuickSortPlainBegin(data, 0, my_len-1)
	}
}

func QuickSortBegin(data sort.Interface, begin int, end int) {
	for begin < end {
		mid := begin + (end-begin)/2
		mid = Partition(data, begin, end, mid)
		if mid-begin > end-mid {
			QuickSortBegin(data, mid+1, end)
			end = mid - 1
		} else {
			QuickSortBegin(data, begin, mid)
			begin = mid + 1
		}
	}
}

func QuickSort(data sort.Interface) {
	if my_len := data.Len(); my_len > 1 {
		QuickSortBegin(data, 0, my_len-1)
	}
}

func PartialSortBegin(data sort.Interface, begin int, end int, n int) {
	for begin < end {
		mid := begin + (end-begin)/2
		mid = Partition(data, begin, end, mid)
		if n <= mid {
			end = mid - 1
		} else if mid-begin > end-mid {
			QuickSortBegin(data, mid+1, end)
			end = mid - 1
		} else {
			QuickSortBegin(data, begin, mid-1)
			begin = mid + 1
		}
	}
}

func PartialSort(data sort.Interface, n int) {
	if my_len := data.Len(); my_len > 1 {
		PartialSortBegin(data, 0, my_len-1, n)
	}
}

func Fsum(in []float64) (res float64) {
	var partials []float64
	for _, a := range in {
		for i, b := range partials {
			if math.Abs(a) < math.Abs(b) {
				a, b = b, a
			}
			hi := a + b
			lo := b - (hi - a)
			if lo != 0 {
				partials[i] = lo
			} else {
				partials[i], partials[len(partials)-1] = partials[len(partials)-1], partials[i]
				partials = partials[:len(partials)-1]
			}
			a = hi
		}
		partials = append(partials, a)
	}
	for _, v := range partials {
		res += v
	}
	return
}

func b2s(in []byte) string {
	bs := (*reflect.SliceHeader)(unsafe.Pointer(&in))
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{Data: bs.Data, Len: bs.Len}))
}

func s2b(in string) []byte {
	ss := (*reflect.StringHeader)(unsafe.Pointer(&in))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{Data: ss.Data, Len: ss.Len, Cap: ss.Len}))
}
