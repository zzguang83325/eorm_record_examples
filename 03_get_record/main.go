package main

import (
	"fmt"

	"github.com/zzguang83325/eorm"
)

// 示例3：GetRecord 和 GetRecords 使用
// 演示如何获取嵌套 Record 和 Record 数组
func main() {
	fmt.Println("========== GetRecord 和 GetRecords 示例 ==========")

	// 1. 创建包含嵌套 Record 的数据
	record := eorm.NewRecord().
		Set("id", 1).
		Set("name", "张三").
		Set("user", eorm.NewRecord().
			Set("id", 1).
			Set("name", "李四").
			Set("age", 30))

	fmt.Printf("1. 创建包含嵌套 Record 的数据: %v\n", record.ToJson())

	// 2. 获取嵌套 Record
	user, err := record.GetRecord("user")
	if err != nil {
		fmt.Printf("2. 获取 user 失败: %v\n", err)
	} else {
		fmt.Printf("2. 获取 user: %v\n", user.ToJson())
		fmt.Printf("   user.name = %s\n", user.GetString("name"))
		fmt.Printf("   user.age = %d\n", user.GetInt("age"))
	}

	// 3. 获取不存在的 Record
	missing, err := record.GetRecord("missing")
	if err != nil {
		fmt.Printf("3. 获取不存在的 Record (预期错误): %v\n", err)
	} else {
		fmt.Printf("3. 获取不存在的 Record: %v\n", missing)
	}

	// 4. 创建包含 Record 数组的数据
	ordersRecord := eorm.NewRecord().
		Set("id", 1).
		Set("orders", []*eorm.Record{
			eorm.NewRecord().Set("order_id", "001").Set("amount", 100),
			eorm.NewRecord().Set("order_id", "002").Set("amount", 200),
			eorm.NewRecord().Set("order_id", "003").Set("amount", 300),
		})

	fmt.Printf("4. 创建包含 Record 数组的数据: %v\n", ordersRecord.ToJson())

	// 5. 获取 Record 数组
	orders, err := ordersRecord.GetRecords("orders")
	if err != nil {
		fmt.Printf("5. 获取 orders 失败: %v\n", err)
	} else {
		fmt.Printf("5. 获取 orders: 共 %d 个\n", len(orders))
		for i, order := range orders {
			fmt.Printf("   orders[%d]: order_id=%s, amount=%d\n", i, order.GetString("order_id"), order.GetInt("amount"))
		}
	}

	// 6. 获取不存在的 Record 数组
	missingOrders, err := ordersRecord.GetRecords("missing")
	if err != nil {
		fmt.Printf("6. 获取不存在的 Record 数组 (预期错误): %v\n", err)
	} else {
		fmt.Printf("6. 获取不存在的 Record 数组: %v\n", missingOrders)
	}

	// 7. 从 JSON 创建并获取嵌套数据
	jsonData := `{
		"id": 1,
		"user": {
			"id": 1,
			"name": "张三",
			"profile": {
				"age": 25,
				"email": "zhangsan@example.com"
			}
		},
		"orders": [
			{"order_id": "001", "amount": 100},
			{"order_id": "002", "amount": 200}
		]
	}`

	jsonRecord := eorm.NewRecord().FromJson(jsonData)
	fmt.Printf("7. 从 JSON 创建: %v\n", jsonRecord.ToJson())

	// 8. 获取嵌套的 Record
	jsonUser, err := jsonRecord.GetRecord("user")
	if err != nil {
		fmt.Printf("8. 获取 user 失败: %v\n", err)
	} else {
		fmt.Printf("8. 获取 user: %v\n", jsonUser.ToJson())
	}

	// 9. 获取嵌套的 Record 数组
	jsonOrders, err := jsonRecord.GetRecords("orders")
	if err != nil {
		fmt.Printf("9. 获取 orders 失败: %v\n", err)
	} else {
		fmt.Printf("9. 获取 orders: 共 %d 个\n", len(jsonOrders))
		for i, order := range jsonOrders {
			fmt.Printf("   orders[%d]: %v\n", i, order.ToJson())
		}
	}

	fmt.Println("\n========== 示例完成 ==========")
}
