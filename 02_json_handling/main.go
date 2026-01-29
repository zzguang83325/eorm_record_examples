package main

import (
	"fmt"

	"github.com/zzguang83325/eorm"
)

// 示例2：Record JSON 处理
// 演示 FromJson 和 ToJson 的使用
func main() {
	fmt.Println("========== Record JSON 处理示例 ==========")

	// 1. 从 JSON 字符串创建 Record
	jsonStr := `{
		"id": 1,
		"name": "张三",
		"age": 25,
		"email": "zhangsan@example.com",
		"active": true
	}`

	record := eorm.NewRecord().FromJson(jsonStr)
	fmt.Printf("1. 从 JSON 创建 Record: %v\n", record.ToJson())

	// 2. 转换为 JSON 字符串
	jsonOutput := record.ToJson()
	fmt.Printf("2. 转换为 JSON: %s\n", jsonOutput)

	// 3. 处理嵌套 JSON
	nestedJson := `{
		"user": {
			"id": 1,
			"profile": {
				"name": "张三",
				"age": 25,
				"address": {
					"city": "上海",
					"street": "浦东"
				}
			}
		}
	}`

	nestedRecord := eorm.NewRecord().FromJson(nestedJson)
	fmt.Printf("3. 嵌套 JSON: %v\n", nestedRecord.ToJson())

	// 4. 处理数组 JSON
	arrayJson := `{
		"id": 1,
		"items": [
			{"id": 1, "name": "商品1", "price": 100},
			{"id": 2, "name": "商品2", "price": 200},
			{"id": 3, "name": "商品3", "price": 300}
		]
	}`

	arrayRecord := eorm.NewRecord().FromJson(arrayJson)
	fmt.Printf("4. 数组 JSON: %v\n", arrayRecord.ToJson())

	// 5. 处理复杂 JSON
	complexJson := `{
		"id": 1,
		"name": "测试",
		"user": {
			"id": 1,
			"profile": {
				"age": 25,
				"address": {
					"city": "上海",
					"street": "浦东"
				}
			}
		},
		"orders": [
			{
				"order_id": "001",
				"amount": 100,
				"items": [
					{"item_id": "i1", "name": "商品1"},
					{"item_id": "i2", "name": "商品2"}
				]
			},
			{
				"order_id": "002",
				"amount": 200,
				"items": [
					{"item_id": "i3", "name": "商品3"}
				]
			}
		]
	}`

	complexRecord := eorm.NewRecord().FromJson(complexJson)
	fmt.Printf("5. 复杂 JSON: %v\n", complexRecord.ToJson())

	// 6. 获取嵌套 Record
	user, err := complexRecord.GetRecord("user")
	if err != nil {
		fmt.Printf("6. 获取 user 失败: %v\n", err)
	} else {
		fmt.Printf("6. 获取 user: %v\n", user.ToJson())
	}

	// 7. 获取 Record 数组
	orders, err := complexRecord.GetRecords("orders")
	if err != nil {
		fmt.Printf("7. 获取 orders 失败: %v\n", err)
	} else {
		fmt.Printf("7. 获取 orders: 共 %d 个\n", len(orders))
		for i, order := range orders {
			fmt.Printf("   orders[%d]: %v\n", i, order.ToJson())
		}
	}

	fmt.Println("\n========== 示例完成 ==========")
}
