package main

import (
	"fmt"

	"github.com/zzguang83325/eorm"
)

// 示例7：Record 错误处理
// 演示如何处理 Record 操作中的错误
func main() {
	fmt.Println("========== Record 错误处理示例 ==========")

	// 1. 获取不存在的字段（类型安全方法返回零值）
	record := eorm.NewRecord().Set("id", 1).Set("name", "张三")
	missing := record.GetString("missing")
	fmt.Printf("1. 获取不存在的字段: %s\n", missing)

	// 2. GetRecord 错误处理
	_, err := record.GetRecord("missing")
	if err != nil {
		fmt.Printf("2. GetRecord 错误 (预期): %v\n", err)
	}

	// 3. GetRecord 类型不匹配
	_, err = record.GetRecord("id")
	if err != nil {
		fmt.Printf("3. GetRecord 类型不匹配 (预期): %v\n", err)
	}

	// 4. GetRecords 错误处理
	_, err = record.GetRecords("missing")
	if err != nil {
		fmt.Printf("4. GetRecords 错误 (预期): %v\n", err)
	}

	// 5. GetRecords 类型不匹配
	_, err = record.GetRecords("name")
	if err != nil {
		fmt.Printf("5. GetRecords 类型不匹配 (预期): %v\n", err)
	}

	// 6. GetRecordByPath 错误处理
	pathRecord := eorm.NewRecord().FromJson(`{
		"a": {
			"b": "value"
		}
	}`)

	_, err = pathRecord.GetRecordByPath("a.c.d")
	if err != nil {
		fmt.Printf("6. GetRecordByPath 错误 (预期): %v\n", err)
	}

	// 7. GetRecordByPath 空路径
	_, err = pathRecord.GetRecordByPath("")
	if err != nil {
		fmt.Printf("7. GetRecordByPath 空路径 (预期): %v\n", err)
	}

	// 8. ToStruct 错误处理
	type User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	var user User
	err = record.ToStruct(&user)
	if err != nil {
		fmt.Printf("8. ToStruct 错误: %v\n", err)
	} else {
		fmt.Printf("8. ToStruct 成功: ID=%d, Name=%s\n", user.ID, user.Name)
	}

	// 9. FromJson 错误处理
	invalidJson := `{not valid json}`
	invalidRecord := eorm.NewRecord().FromJson(invalidJson)
	fmt.Printf("9. FromJson 无效 JSON (静默处理): %v\n", invalidRecord.ToJson())

	// 10. 完整的错误处理流程
	fmt.Println("\n10. 完整的错误处理流程:")
	completeRecord := eorm.NewRecord().FromJson(`{
		"user": {
			"id": 1,
			"name": "张三"
		},
		"orders": [
			{"order_id": "001", "amount": 100}
		]
	}`)

	// 尝试获取 user
	if user, err := completeRecord.GetRecord("user"); err != nil {
		fmt.Printf("   获取 user 失败: %v\n", err)
	} else {
		fmt.Printf("   获取 user 成功: %v\n", user.ToJson())
	}

	// 尝试获取 orders
	if orders, err := completeRecord.GetRecords("orders"); err != nil {
		fmt.Printf("   获取 orders 失败: %v\n", err)
	} else {
		fmt.Printf("   获取 orders 成功: 共 %d 个\n", len(orders))
	}

	// 尝试获取不存在的字段
	if _, err := completeRecord.GetRecord("missing"); err != nil {
		fmt.Printf("   获取 missing 失败 (预期): %v\n", err)
	}

	fmt.Println("\n========== 示例完成 ==========")
}
