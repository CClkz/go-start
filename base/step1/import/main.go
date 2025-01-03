// 包声明，指明文件属于哪个包，必须在非注释的第一行
// package main表示这是个可执行的程序，可以go run
package main

// 引入包
import (
	"fmt"
	// 此处modelAlt是别名，也可不加，不加下面函数里引用时需用包名model
	// 注意，包名main是被当做程序而不是包，不能引入的
	modelAlt "my-app/base/step1/user/model"
)

func main() {
	/* 这是我的第一个简单的程序 */
	fmt.Println("Hello, World!")

	user := modelAlt.User{ID: 1, Name: "Alice"}
	fmt.Println("User:", user)
}
