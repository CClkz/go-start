#### 创建一个测试
比如目录下有sum.go，准备测试Sum函数
```
package main

// Sum calculates the total from an array of numbers.
func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}
```
1. 同目录下创建sum_test.go文件
```
package main

import "testing"

func TestSum(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if want != got {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
```
2. 执行测试
```
go test
```
命令会自动去找_test.go后缀的文件，

执行里面函数名TestXxx且参数为（t *testing.T）的测试函数

（这里不会对Test后面的Xxx和代码的函数名作校验，首字母大写即可）

