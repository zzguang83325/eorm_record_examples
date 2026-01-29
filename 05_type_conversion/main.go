package main

import (
	"fmt"

	"github.com/zzguang83325/eorm"
)

// 示例5：Record 类型转换
// 演示 Record 与 struct 之间的转换
func main() {
	fmt.Println("========== Record 类型转换示例 ==========")

	// 1. 定义一个 struct
	type User struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Email  string `json:"email"`
		Active bool   `json:"active"`
	}

	// 2. 从 struct 创建 Record
	user := User{
		ID:     1,
		Name:   "张三",
		Age:    25,
		Email:  "zhangsan@example.com",
		Active: true,
	}

	record := eorm.NewRecord().FromStruct(user)
	fmt.Printf("1. 从 struct 创建 Record: %v\n", record.ToJson())

	// 3. 从 Record 转换为 struct
	var result User
	err := record.ToStruct(&result)
	if err != nil {
		fmt.Printf("2. 转换为 struct 失败: %v\n", err)
	} else {
		fmt.Printf("2. 转换为 struct: ID=%d, Name=%s, Age=%d, Email=%s, Active=%v\n",
			result.ID, result.Name, result.Age, result.Email, result.Active)
	}

	// 4. 修改 Record 并转换回 struct
	record.Set("age", 30).Set("active", false)
	err = record.ToStruct(&result)
	if err != nil {
		fmt.Printf("3. 修改后转换失败: %v\n", err)
	} else {
		fmt.Printf("3. 修改后转换: ID=%d, Name=%s, Age=%d, Email=%s, Active=%v\n",
			result.ID, result.Name, result.Age, result.Email, result.Active)
	}

	// 5. 嵌套 struct 转换
	type Address struct {
		City    string `json:"city"`
		Street  string `json:"street"`
		ZipCode string `json:"zip_code"`
	}

	type Profile struct {
		Name    string  `json:"name"`
		Age     int     `json:"age"`
		Address Address `json:"address"`
	}

	type NestedUser struct {
		ID      int     `json:"id"`
		Name    string  `json:"name"`
		Profile Profile `json:"profile"`
	}

	nestedUser := NestedUser{
		ID:   1,
		Name: "张三",
		Profile: Profile{
			Name: "张三",
			Age:  25,
			Address: Address{
				City:    "上海",
				Street:  "浦东",
				ZipCode: "200000",
			},
		},
	}

	nestedRecord := eorm.NewRecord().FromStruct(nestedUser)
	fmt.Printf("4. 嵌套 struct 转换: %v\n", nestedRecord.ToJson())

	// 6. 转换回嵌套 struct
	var nestedResult NestedUser
	err = nestedRecord.ToStruct(&nestedResult)
	if err != nil {
		fmt.Printf("5. 嵌套转换失败: %v\n", err)
	} else {
		fmt.Printf("5. 嵌套转换结果: ID=%d, Name=%s, Profile.Name=%s, Profile.Address.City=%s\n",
			nestedResult.ID, nestedResult.Name, nestedResult.Profile.Name, nestedResult.Profile.Address.City)
	}

	// 7. 数组 struct 转换
	type Item struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Price float64 `json:"price"`
	}

	items := []Item{
		{ID: 1, Name: "商品1", Price: 100},
		{ID: 2, Name: "商品2", Price: 200},
		{ID: 3, Name: "商品3", Price: 300},
	}

	itemsRecord := eorm.NewRecord().Set("items", items)
	fmt.Printf("6. 数组 struct 转换: %v\n", itemsRecord.ToJson())

	fmt.Println("\n========== 示例完成 ==========")
}
