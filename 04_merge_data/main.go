package main

import (
	"fmt"

	"github.com/zzguang83325/eorm"
)

// 示例：多个 FromJson、FromMap、FromStruct 连续调用合并数据
func main() {
	fmt.Println("========== 多个 FromJson、FromMap、FromStruct 合并数据示例 ==========")

	// 1. 多个 FromJson 合并
	fmt.Println("\n1. 多个 FromJson 合并")
	result1 := eorm.NewRecord().
		FromJson(`{"id": 1, "name": "张三"}`).
		FromJson(`{"age": 25, "email": "zhangsan@example.com"}`).
		FromJson(`{"active": true, "created_at": "2024-01-01"}`).
		FromJson(`{"active": false, "updated_at": "2024-01-01"}`)
	fmt.Printf("   结果: %s\n", result1.ToJson())

	// 2. 多个 FromMap 合并
	fmt.Println("\n2. 多个 FromMap 合并")
	result2 := eorm.NewRecord().
		FromMap(map[string]interface{}{
			"id":   1,
			"name": "张三",
		}).
		FromMap(map[string]interface{}{
			"age":   25,
			"email": "zhangsan@example.com",
		}).
		FromMap(map[string]interface{}{
			"active":     true,
			"created_at": "2024-01-01",
		})
	fmt.Printf("   结果: %s\n", result2.ToJson())

	// 3. 多个 FromStruct 合并
	fmt.Println("\n3. 多个 FromStruct 合并")
	type BasicInfo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	type ContactInfo struct {
		Age   int    `json:"age"`
		Email string `json:"email"`
	}
	type StatusInfo struct {
		Active    bool   `json:"active"`
		CreatedAt string `json:"created_at"`
	}

	result3 := eorm.NewRecord().
		FromStruct(BasicInfo{ID: 1, Name: "张三3"}).
		FromStruct(ContactInfo{Age: 25, Email: "zhangsan3@example.com"}).
		FromStruct(StatusInfo{Active: true, CreatedAt: "2024-01-01"})
	fmt.Printf("   结果: %s\n", result3.ToJson())

	// 4. 混合调用：FromJson + FromMap + FromStruct
	fmt.Println("\n4. 混合调用：FromJson + FromMap + FromStruct")
	result4 := eorm.NewRecord().
		FromJson(`{"id": 1, "name": "张三"}`).
		FromMap(map[string]interface{}{
			"age":   25,
			"email": "zhangsan@example.com",
		}).
		FromStruct(StatusInfo{Active: true, CreatedAt: "2024-01-01"})
	fmt.Printf("   结果: %s\n", result4.ToJson())

	// 5. 混合调用：FromStruct + FromJson + FromMap
	fmt.Println("\n5. 混合调用：FromStruct + FromJson + FromMap")
	result5 := eorm.NewRecord().
		FromStruct(BasicInfo{ID: 1, Name: "张三"}).
		FromJson(`{"age": 25, "email": "zhangsan@example.com"}`).
		FromMap(map[string]interface{}{
			"active":     true,
			"created_at": "2024-01-01",
		})
	fmt.Printf("   结果: %s\n", result5.ToJson())

	// 6. 混合调用：FromMap + FromStruct + FromJson
	fmt.Println("\n6. 混合调用：FromMap + FromStruct + FromJson")
	result6 := eorm.NewRecord().
		FromMap(map[string]interface{}{
			"id":   1,
			"name": "张三",
		}).
		FromStruct(ContactInfo{Age: 25, Email: "zhangsan@example.com"}).
		FromJson(`{"active": true, "created_at": "2024-01-01"}`)
	fmt.Printf("   结果: %s\n", result6.ToJson())

	// 7. 混合调用 + Set
	fmt.Println("\n7. 混合调用 + Set")
	result7 := eorm.NewRecord().
		FromJson(`{"id": 1, "name": "张三"}`).
		FromMap(map[string]interface{}{
			"age": 25,
		}).
		FromStruct(ContactInfo{Email: "zhangsan@example.com"}).
		Set("active", false).
		Set("created_at", "2024-01-01")
	fmt.Printf("   结果: %s\n", result7.ToJson())

	// 8. 演示字段覆盖（后面的覆盖前面的）
	fmt.Println("\n8. 演示字段覆盖（后面的覆盖前面的）")
	result8 := eorm.NewRecord().
		FromJson(`{"id": 1, "name": "张三", "age": 20}`).
		FromMap(map[string]interface{}{
			"age":   25, // 覆盖前面的 age: 20
			"email": "zhangsan@example.com",
		}).
		FromStruct(BasicInfo{ID: 2, Name: "李四"}).
		FromJson(`{"age": 30, "active": true}`)
	fmt.Printf("   结果: %s\n", result8.ToJson())
	fmt.Printf("   说明: id 被覆盖为 2, name 被覆盖为 李四, age 最终为 30\n")

	// 9. 复杂场景：从多个来源合并用户数据
	fmt.Println("\n9. 复杂场景：从多个来源合并用户数据")
	type Profile struct {
		Name     string `json:"name"`
		Age      int    `json:"age"`
		Gender   string `json:"gender"`
		Birthday string `json:"birthday"`
	}
	type Account struct {
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Username string `json:"username"`
	}
	type Settings struct {
		Theme        string `json:"theme"`
		Language     string `json:"language"`
		Notification bool   `json:"notification"`
	}

	result9 := eorm.NewRecord().
		FromJson(`{"id": 1001, "created_at": "2024-01-01"}`).
		FromStruct(Profile{
			Name:     "张三",
			Age:      25,
			Gender:   "男",
			Birthday: "1999-01-01",
		}).
		FromMap(map[string]interface{}{
			"email":    "zhangsan@example.com",
			"phone":    "13800138000",
			"username": "zhangsan",
		}).
		FromStruct(Settings{
			Theme:        "dark",
			Language:     "zh-CN",
			Notification: true,
		}).
		Set("updated_at", "2024-01-28").
		Set("status", "active")
	fmt.Printf("   结果: %s\n", result9.ToJson())

	// 10. 嵌套 Record 合并
	fmt.Println("\n10. 嵌套 Record 合并")
	profileRecord := eorm.NewRecord().
		FromJson(`{"name": "张三", "age": 25}`)
	addressRecord := eorm.NewRecord().
		FromMap(map[string]interface{}{
			"city":    "北京",
			"country": "中国",
		})

	result10 := eorm.NewRecord().
		FromJson(`{"id": 1, "email": "zhangsan@example.com"}`).
		Set("profile", profileRecord).
		Set("address", addressRecord).
		FromJson(`{"active": true}`)
	fmt.Printf("   结果: %s\n", result10.ToJson())

	fmt.Println("\n========== 示例完成 ==========")
	fmt.Println("\n说明：")
	fmt.Println("1. FromJson、FromMap、FromStruct 都会合并数据到现有 Record")
	fmt.Println("2. 相同字段会被后面的值覆盖")
	fmt.Println("3. 可以混合使用这三个方法")
	fmt.Println("4. Set 方法也可以在链式调用中使用")
}
