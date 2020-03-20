package splits

import (
	"reflect"
	"testing"
)

type testCase struct {
	str  string
	sep  string
	want []string
}

func TestSplit(t *testing.T) {
	ret := Splits("wlw|test|kkk", "|")
	want := []string{"wlw", "test", "kkk"}
	if !reflect.DeepEqual(ret, want) {
		t.Errorf("want: %v but ret: %v", want, ret)
	}
}

// go test -v || go test -run TestGroupSplit/test1
func TestGroupSplit(t *testing.T) {
	testGroup := map[string]testCase{
		"test1": testCase{str: "wlw,aab,ccc", sep: ",", want: []string{"wlw", "aab", "ccc"}},
		"test2": testCase{str: "www.baidu.com", sep: ".", want: []string{"www", "baidu", "com"}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			ret := Splits(tc.str, tc.sep)
			if !reflect.DeepEqual(ret, tc.want) {
				t.Errorf("want: %v but ret: %v", tc.want, ret)
			}
		})

	}
}

// go test -bench=Split 运行时间
// go test -bench=Split -benchmem 内存申请
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Splits("wlw|test|kkk", "|")
	}
}

// 性能比较测试
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

// go test -bench=Fib1 -benchmem
func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
