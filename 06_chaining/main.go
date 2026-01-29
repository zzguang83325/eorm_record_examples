package main

import (
	"fmt"

	"github.com/zzguang83325/eorm"
)

// 示例6：Record 链式调用
// 演示 Record 的链式调用特性
func main() {
	fmt.Println("========== Record 链式调用示例 ==========")

	// 1. 基础链式调用
	record := eorm.NewRecord().
		Set("id", 1).
		Set("name", "张三").
		Set("age", 25).
		Set("email", "zhangsan@example.com").
		Set("active", true)

	fmt.Printf("1. 基础链式调用: %v\n", record.ToJson())

	// 2. 链式调用 + 获取
	name := record.
		Set("name", "李四").
		Get("name")

	fmt.Printf("2. 链式调用 + 获取: name = %v\n", name)

	// 3. 链式调用 + 类型安全方法
	age := record.
		Set("age", 30).
		GetInt("age")

	fmt.Printf("3. 链式调用 + 类型安全: age = %d\n", age)

	// 4. 链式调用 + 删除
	record.
		Set("new_field", "new value").
		Delete("new_field")

	fmt.Printf("4. 链式调用 + 删除: %v\n", record.ToJson())

	// 5. 复杂的链式调用
	complexRecord := eorm.NewRecord().
		Set("id", 1).
		Set("name", "张三").
		Set("profile", eorm.NewRecord().
			Set("age", 25).
			Set("email", "zhangsan@example.com").
			Set("address", eorm.NewRecord().
				Set("city", "上海").
				Set("street", "浦东").
				Set("zip_code", "200000"))).
		Set("orders", []*eorm.Record{
			eorm.NewRecord().
				Set("order_id", "001").
				Set("amount", 100).
				Set("items", []interface{}{
					eorm.NewRecord().Set("item_id", "i1").Set("name", "商品1"),
					eorm.NewRecord().Set("item_id", "i2").Set("name", "商品2"),
				}),
			eorm.NewRecord().
				Set("order_id", "002").
				Set("amount", 200).
				Set("items", []interface{}{
					eorm.NewRecord().Set("item_id", "i3").Set("name", "商品3"),
				}),
		})

	fmt.Printf("5. 复杂链式调用: %v\n", complexRecord.ToJson())

	// 6. 链式调用 + JSON 转换
	jsonStr := eorm.NewRecord().
		Set("id", 1).
		Set("name", "张三").
		Set("age", 25).
		ToJson()

	fmt.Printf("6. 链式调用 + JSON: %s\n", jsonStr)

	// 7. 链式调用 + struct 转换
	type User struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Age   int    `json:"age"`
	}

	var user User
	err := eorm.NewRecord().
		Set("id", 1).
		Set("name", "张三").
		Set("age", 25).
		ToStruct(&user)

	if err != nil {
		fmt.Printf("7. 链式调用 + struct 转换失败: %v\n", err)
	} else {
		fmt.Printf("7. 链式调用 + struct 转换: ID=%d, Name=%s, Age=%d\n", user.ID, user.Name, user.Age)
	}

	// 8. 链式调用 + 条件操作
	conditionalRecord := eorm.NewRecord().
		Set("id", 1).
		Set("name", "张三")

	// 根据条件设置字段
	if true {
		conditionalRecord.Set("active", true)
	}

	if false {
		conditionalRecord.Set("deleted", true)
	}

	fmt.Printf("8. 链式调用 + 条件操作: %v\n", conditionalRecord.ToJson())

	fmt.Println("\n========== 示例完成 ==========")
}
