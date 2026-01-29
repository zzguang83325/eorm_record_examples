package main

import (
	"fmt"

	"github.com/zzguang83325/eorm"
)

// 示例：测试 GetRecordByPath 方法
func main() {
	fmt.Println("========== GetRecordByPath 方法测试 ==========")

	// 1. 基本测试：获取嵌套的 Record
	fmt.Println("\n1. 基本测试：获取嵌套的 Record")
	record1 := eorm.NewRecord().FromJson(`{
		"user": {
			"name": "张三",
			"email": "zhangsan@example.com",
			"age": 25
		}
	}`)

	user, err := record1.GetRecordByPath("user")
	if err != nil {
		fmt.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ user = %s\n", user.ToJson())
	}

	// 2. 多层嵌套测试
	fmt.Println("\n2. 多层嵌套测试")
	record2 := eorm.NewRecord().FromJson(`{
		"data": {
			"profile": {
				"basic": {
					"firstname": "张",
					"lastname": "三",
					"age": 25
				}
			}
		}
	}`)

	basic, err := record2.GetRecordByPath("data.profile.basic")
	if err != nil {
		fmt.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ data.profile.basic = %s\n", basic.ToJson())
	}

	// 3. 获取中间层的 Record
	fmt.Println("\n3. 获取中间层的 Record")
	record3 := eorm.NewRecord().FromJson(`{
		"data": {
			"profile": {
				"basic": {
					"name": "张三"
				}
			}
		}
	}`)

	profile, err := record3.GetRecordByPath("data.profile")
	if err != nil {
		fmt.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ data.profile = %s\n", profile.ToJson())
	}

	// 4. 错误处理：路径不存在
	fmt.Println("\n4. 错误处理：路径不存在")
	record4 := eorm.NewRecord().FromJson(`{
		"user": {
			"name": "张三"
		}
	}`)

	_, err = record4.GetRecordByPath("user.email")
	if err != nil {
		fmt.Printf("   ✅ user.email 不存在，正确返回错误: %v\n", err)
	} else {
		fmt.Printf("   ❌ 应该返回错误\n")
	}

	_, err = record4.GetRecordByPath("user.profile.name")
	if err != nil {
		fmt.Printf("   ✅ user.profile.name 不存在，正确返回错误: %v\n", err)
	} else {
		fmt.Printf("   ❌ 应该返回错误\n")
	}

	// 5. 错误处理：空路径
	fmt.Println("\n5. 错误处理：空路径")
	record5 := eorm.NewRecord().FromJson(`{"name": "张三"}`)

	_, err = record5.GetRecordByPath("")
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

	_, err = record6.GetRecordByPath("user.name.first")
	if err != nil {
		fmt.Printf("   ✅ user.name 不是 Record，正确返回错误: %v\n", err)
	} else {
		fmt.Printf("   ❌ 应该返回错误\n")
	}

	// 7. 错误处理：最终值不是 Record
	fmt.Println("\n7. 错误处理：最终值不是 Record")
	record7 := eorm.NewRecord().FromJson(`{
		"user": {
			"name": "张三",
			"age": 25
		}
	}`)

	_, err = record7.GetRecordByPath("user.name")
	if err != nil {
		fmt.Printf("   ✅ user.name 不是 Record，正确返回错误: %v\n", err)
	} else {
		fmt.Printf("   ❌ 应该返回错误\n")
	}

	// 8. 复杂场景：从多个来源合并后获取嵌套 Record
	fmt.Println("\n8. 复杂场景：从多个来源合并后获取嵌套 Record")
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

	contact, err := record8.GetRecordByPath("contact")
	if err != nil {
		fmt.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ contact = %s\n", contact.ToJson())
	}

	// 9. 获取 Record 后继续操作
	fmt.Println("\n9. 获取 Record 后继续操作")
	record9 := eorm.NewRecord().FromJson(`{
		"user": {
			"name": "张三",
			"age": 25
		}
	}`)

	user9, err := record9.GetRecordByPath("user")
	if err != nil {
		fmt.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ user = %s\n", user9.ToJson())
		name := user9.GetString("name")
		age := user9.GetInt("age")
		fmt.Printf("   ✅ name = %s, age = %d\n", name, age)
		user9.Set("email", "zhangsan@example.com")
		fmt.Printf("   ✅ 添加 email 后: %s\n", user9.ToJson())
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

	dbConfig, err := record10.GetRecordByPath("database")
	if err != nil {
		fmt.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ 数据库配置: %s\n", dbConfig.ToJson())
		host := dbConfig.GetString("host")
		port := dbConfig.GetInt("port")
		username := dbConfig.GetString("username")
		password := dbConfig.GetString("password")
		dbname := dbConfig.GetString("dbname")
		fmt.Printf("      Host: %s\n", host)
		fmt.Printf("      Port: %d\n", port)
		fmt.Printf("      Username: %s\n", username)
		fmt.Printf("      Password: %s\n", password)
		fmt.Printf("      Database: %s\n", dbname)
	}

	serverConfig, err := record10.GetRecordByPath("server")
	if err != nil {
		fmt.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ 服务器配置: %s\n", serverConfig.ToJson())
		port := serverConfig.GetInt("port")
		mode := serverConfig.GetString("mode")
		fmt.Printf("      Port: %d\n", port)
		fmt.Printf("      Mode: %s\n", mode)
	}

	// 11. 嵌套 Record 的嵌套 Record
	fmt.Println("\n11. 嵌套 Record 的嵌套 Record")
	record11 := eorm.NewRecord().FromJson(`{
		"level1": {
			"level2": {
				"level3": {
					"value": "deep"
				}
			}
		}
	}`)

	level3, err := record11.GetRecordByPath("level1.level2.level3")
	if err != nil {
		fmt.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ level1.level2.level3 = %s\n", level3.ToJson())
		value := level3.GetString("value")
		fmt.Printf("   ✅ value = %s\n", value)
	}

	// 12. 混合使用 GetRecordByPath 和 GetStringByPath
	fmt.Println("\n12. 混合使用 GetRecordByPath 和 GetStringByPath")
	record12 := eorm.NewRecord().FromJson(`{
		"user": {
			"profile": {
				"name": "张三",
				"age": 25
			},
			"contact": {
				"email": "zhangsan@example.com",
				"phone": "13800138000"
			}
		}
	}`)

	profile, err = record12.GetRecordByPath("user.profile")
	if err != nil {
		fmt.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ profile = %s\n", profile.ToJson())
	}

	email, err := record12.GetStringByPath("user.contact.email")
	if err != nil {
		fmt.Printf("   ❌ 获取失败: %v\n", err)
	} else {
		fmt.Printf("   ✅ user.contact.email = %s\n", email)
	}

	fmt.Println("\n========== 测试完成 ==========")
}
