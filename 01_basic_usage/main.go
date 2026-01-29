package main

import (
	"fmt"

	"github.com/zzguang83325/eorm"
)

// 示例1：Record 基础用法
// 演示 Record 的基本操作：创建、设置、获取、删除、清空等
func main() {
	fmt.Println("========== Record 基础用法示例 ==========")

	// 1. 创建空 Record
	record := eorm.NewRecord()
	fmt.Printf("1. 创建空 Record: %v\n", record.ToJson())

	// 2. 设置多个字段
	record.Set("name", "张三")
	record.Set("age", 25)
	record.Set("active", true)
	record.Set("email", "zhangsan@example.com")
	fmt.Printf("2. 设置多个字段: %v\n", record.ToJson())

	// 3. 获取值（推荐使用类型安全的方法）
	name := record.GetString("name")
	age := record.GetInt("age")
	fmt.Printf("3. 获取值: name=%s, age=%d\n", name, age)

	// 4. 更多类型安全获取
	activeBool := record.GetBool("active")
	fmt.Printf("4. 更多类型安全获取: active=%t\n", activeBool)

	// 5. 获取不存在的字段（类型安全方法返回零值）
	missing := record.GetString("missing")
	fmt.Printf("5. 获取不存在的字段: %s\n", missing)

	// 6. 检查字段是否存在
	hasName := record.Has("name")
	hasMissing := record.Has("missing")
	fmt.Printf("6. 字段 'name' 存在: %t\n", hasName)
	fmt.Printf("   字段 'missing' 存在: %t\n", hasMissing)

	// 7. 删除字段
	record.Delete("email")
	fmt.Printf("7. 删除字段 'email': %v\n", record.ToJson())

	// 8. 获取所有字段名
	columns := record.Columns()
	fmt.Printf("8. 所有字段名: %v\n", columns)

	// 9. 清空 Record
	record.Clear()
	fmt.Printf("9. 清空 Record: %v\n", record.ToJson())

	fmt.Println("\n========== 示例完成 ==========")
}
