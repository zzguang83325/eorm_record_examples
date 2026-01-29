package main

import (
	"fmt"

	"github.com/zzguang83325/eorm"
)

// 示例：测试 GetStringByPath 方法
func main() {
	fmt.Println("========== GetStringByPath 方法测试 ==========")

	// 1. 基本测试：获取嵌套的字符串值
	fmt.Println("\n1. 基本测试：获取嵌套的字符串值")
	record1 := eorm.NewRecord().FromJson(`{
		"user": {
			"name": "张三",
			"email": "zhangsan@example.com"
		}
	}`)

	name, err := record1.GetStringByPath("user.name")
	if err != nil {
		fmt.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ user.name = %s\n", name)
	}

	email, err := record1.GetStringByPath("user.email")
	if err != nil {
		fmt.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ user.email = %s\n", email)
	}

	// 2. 多层嵌套测试
	fmt.Println("\n2. 多层嵌套测试")
	record2 := eorm.NewRecord().FromJson(`{
		"data": {
			"profile": {
				"basic": {
					"firstname": "张",
					"lastname": "三"
				}
			}
		}
	}`)

	firstname, err := record2.GetStringByPath("data.profile.basic")
	if err != nil {
		fmt.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ data.profile.basic = %s\n", firstname)
	}

	lastname, err := record2.GetStringByPath("data.profile.basic.lastname")
	if err != nil {
		fmt.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ data.profile.basic.lastname = %s\n", lastname)
	}

	// 3. 类型转换测试：将数字转换为字符串
	fmt.Println("\n3. 类型转换测试：将数字转换为字符串")
	record3 := eorm.NewRecord().FromJson(`{
		"user": {
			"id": 12345,
			"age": 25
		}
	}`)

	id, err := record3.GetStringByPath("user.id")
	if err != nil {
		fmt.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ user.id (数字转字符串) = %s\n", id)
	}

	age, err := record3.GetStringByPath("user.age")
	if err != nil {
		fmt.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ user.age (数字转字符串) = %s\n", age)
	}

	// 4. 错误处理：路径不存在
	fmt.Println("\n4. 错误处理：路径不存在")
	record4 := eorm.NewRecord().FromJson(`{
		"user": {
			"name": "张三"
		}
	}`)

	_, err = record4.GetStringByPath("user.email")
	if err != nil {
		fmt.Printf("   ✅ user.email 不存在，正确返回错误: %v\n", err)
	} else {
		fmt.Printf("   ❌ 应该返回错误\n")
	}

	_, err = record4.GetStringByPath("user.profile.name")
	if err != nil {
		fmt.Printf("   ✅ user.profile.name 不存在，正确返回错误: %v\n", err)
	} else {
		fmt.Printf("   ❌ 应该返回错误\n")
	}

	// 5. 错误处理：空路径
	fmt.Println("\n5. 错误处理：空路径")
	record5 := eorm.NewRecord().FromJson(`{"name": "张三"}`)

	_, err = record5.GetStringByPath("")
	if err != nil {
		fmt.Printf("   ✅ 空路径，正确返回错误: %v\n", err)
	} else {
		fmt.Printf("   ❌ 应该返回错误\n")
	}

	// 6. 错误处理：中间路径不是 Record
	fmt.Println("\n6. 错误处理：中间路径不是 Record")
	record6 := eorm.NewRecord().FromJson(`{
		"user": {
			"name": "张三"
		}
	}`)

	_, err = record6.GetStringByPath("user.name.first")
	if err != nil {
		fmt.Printf("   ✅ user.name 不是 Record，正确返回错误: %v\n", err)
	} else {
		fmt.Printf("   ❌ 应该返回错误\n")
	}

	// 7. 布尔值转字符串
	fmt.Println("\n7. 布尔值转字符串")
	record7 := eorm.NewRecord().FromJson(`{
		"user": {
			"active": true,
			"verified": false
		}
	}`)

	active, err := record7.GetStringByPath("user.active")
	if err != nil {
		fmt.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ user.active (布尔转字符串) = %s\n", active)
	}

	verified, err := record7.GetStringByPath("user.verified")
	if err != nil {
		fmt.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ user.verified (布尔转字符串) = %s\n", verified)
	}

	// 8. 复杂场景：从多个来源合并后获取嵌套值
	fmt.Println("\n8. 复杂场景：从多个来源合并后获取嵌套值")
	type Profile struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	type Contact struct {
		Email string `json:"email"`
		Phone string `json:"phone"`
	}

	record8 := eorm.NewRecord().
		FromJson(`{"id": 1}`).
		FromStruct(Profile{Name: "张三", Age: 25}).
		FromMap(map[string]interface{}{
			"contact": map[string]interface{}{
				"email": "zhangsan@example.com",
				"phone": "13800138000",
			},
		})

	email, err = record8.GetStringByPath("contact.email")
	if err != nil {
		fmt.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ contact.email = %s\n", email)
	}

	phoneStr, err := record8.GetStringByPath("contact.phone")
	if err != nil {
		fmt.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ contact.phone = %s\n", phoneStr)
	}

	// 9. 数组中的对象（注意：数组索引需要特殊处理）
	fmt.Println("\n9. 数组中的对象")
	record9 := eorm.NewRecord().FromJson(`{
		"users": [
			{"name": "张三", "age": 25},
			{"name": "李四", "age": 30}
		]
	}`)

	// 注意：当前实现不支持数组索引，需要先获取 Record 再处理
	users, err := record9.GetRecord("users")
	if err != nil {
		fmt.Printf("   ❌ 获取 users 失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ users = %s\n", users.ToJson())
		fmt.Printf("   💡 提示：数组索引需要先获取 Record 再处理\n")
	}

	// 10. 实际应用场景：配置文件读取
	fmt.Println("\n10. 实际应用场景：配置文件读取")
	record10 := eorm.NewRecord().FromJson(`{
		"database": {
			"host": "localhost",
			"port": 3306,
			"username": "root",
			"password": "123456",
			"dbname": "mydb"
		},
		"server": {
			"port": 8080,
			"mode": "production"
		}
	}`)

	dbHost, _ := record10.GetStringByPath("database.host")
	dbPort, _ := record10.GetStringByPath("database.port")
	dbUser, _ := record10.GetStringByPath("database.username")
	dbPass, _ := record10.GetStringByPath("database.password")
	dbName, _ := record10.GetStringByPath("database.dbname")

	fmt.Printf("   ✅ 数据库配置:\n")
	fmt.Printf("      Host: %s\n", dbHost)
	fmt.Printf("      Port: %s\n", dbPort)
	fmt.Printf("      Username: %s\n", dbUser)
	fmt.Printf("      Password: %s\n", dbPass)
	fmt.Printf("      Database: %s\n", dbName)

	serverPort, _ := record10.GetStringByPath("server.port")
	serverMode, _ := record10.GetStringByPath("server.mode")

	fmt.Printf("   ✅ 服务器配置:\n")
	fmt.Printf("      Port: %s\n", serverPort)
	fmt.Printf("      Mode: %s\n", serverMode)

	fmt.Println("\n========== 测试完成 ==========")
}
