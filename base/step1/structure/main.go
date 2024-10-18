// 包声明，指明文件属于哪个包，必须在非注释的第一行
// package main表示这是个可执行的程序，可以go run
package main

// 引入包
import (
	"fmt"
	// 引入路径是目录
	// 目录和包名一致时，简洁，推荐一致
	// 目录和包名不一致，还是引入目录，程序里使用的是包名，这里编辑器自动加上了别名（===包名）
	usermodels "my-app/base/step1/user/model"
	// vars包取名main了，不能被引入
	// vars "my-app/base/vars"
)

func main() {
	/* 这是我的第一个简单的程序 */
	fmt.Println("Hello, World!")

	// n1 := vars.GetMale()
	// fmt.Println(n1)

	user := usermodels.User{ID: 1, Name: "Alice"}
	fmt.Println("User:", user)
}
