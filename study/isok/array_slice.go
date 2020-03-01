package main

import "fmt"

func main() {
	// 一维数组
	var arr1 [3]int
	arr1 = [3]int{1, 2, 3}

	var arr2 = [4]int{1, 2, 3, 4}

	arr3 := [3]int{1, 2, 3}

	var arr4 = [...]int{1, 3}

	arr5 := [...]int{0: 100, 4: 100}

	fmt.Println(arr1, arr2, arr3, arr4, arr5)

	var arr2Arr2 [2][2]int
	arr2Arr2 = [2][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
	}
	fmt.Println(arr2Arr2)

	// 1. 求数组[1,3,5,7,8]所有元素的和
	arr6 := [...]int{1, 3, 5, 7, 8}
	var sum = 0
	for _, v := range arr6 {
		sum += v
	}
	fmt.Println(sum)

	// 2. 找到数组中和为指定值的两个元素下标，比如从数组[1,3,5,7,8]中找到和为8的两个元素下标分别为(0,3)和(1,2)
	for k1, v1 := range arr6 {
		for k2, v2 := range arr6 {
			if v1+v2 == 8 {
				fmt.Println(k1, k2)
			}
		}
	}

	for i := 0; i < len(arr6); i++ {
		for j := i; j < len(arr6); j++ {
			if arr6[i]+arr6[j] == 8 {
				fmt.Println(i, j)
			}
		}
	}

	var arr10 = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 切片是数组的引用，因此切片是引用类型
	sl1 := arr10[2:7]
	fmt.Println(sl1, len(sl1), cap(sl1)) // [2 3 4 5 6] 5 8
	arr10[7] = 110110
	sl1_1 := sl1[1:6]
	fmt.Println(sl1_1, len(sl1_1), cap(sl1_1)) // 3 4 5 6 110110] 5 7
	fmt.Println(&arr10[0], &sl1[0], &sl1_1[0])

	sl2 := make([]int, 10, 100)
	fmt.Println(sl2, len(sl2), cap(sl2)) // [0 0 0 0 0 0 0 0 0 0] 10 100

	// arrNil == nil，nil代表没有开辟内存空间
	var arrNil []int
	fmt.Println(arrNil == nil, len(arrNil), cap(arrNil))

	var sl3 = make([]int, 0, 1)
	fmt.Println(sl3 == nil, len(sl3), cap(sl3))

	// append & copy
	var sl4 = make([]int, 0, 3)
	// 使用append，如果底层数组容量不够，会申请新的地址空间
	sl4 = append(sl4, 1)
	fmt.Println(sl4, &sl4[0])
	sl4 = append(sl4, 2, 3)
	fmt.Println(sl4, &sl4[0])
	sl5 := []int{44, 55}
	sl4 = append(sl4, sl5...)
	fmt.Println(sl4, &sl4[0])

	var s1 = []int{1, 2, 3}
	var s2 []int = s1
	var s3 = make([]int, 3, 3)
	copy(s3, s1)
	fmt.Println(s1, s2, s3)
	fmt.Println(&s1[0], &s2[0], &s3[0])

	// 删除切片中的某个数字
	arry1 := []int{1, 2, 3, 4, 5}
	var sli1 = arry1[:]
	sli1 = append(sli1[:2], sli1[3:]...)
	fmt.Println(sli1, len(sli1), cap(sli1), arry1)

	// len和cap的关系
	var arry2 = make([]int, 3, 10)
	for i := 0; i < 10; i++ {
		arry2 = append(arry2, i)
	}
	fmt.Println(arry2)

}
