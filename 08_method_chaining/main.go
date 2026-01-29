package main

import (
	"fmt"

	"github.com/zzguang83325/eorm"
)

// 示例10：FromJson、FromMap、FromStruct、FromRecord、Clone 综合链式调用
// 演示这些方法与其他方法的链式调用
// 注意：只有返回 *Record 的方法才能用于链式调用
// Get、GetString、GetInt 等方法返回值，不能继续链式调用
// Delete、Clear、Has 等方法没有返回值或返回值，不能用于链式调用
func main() {
	fmt.Println("========== FromJson、FromMap、FromStruct、FromRecord、Clone 综合链式调用示例 ==========")

	// 1. FromJson + Set + ToJson 链式调用
	fmt.Println("\n1. FromJson + Set + ToJson 链式调用")
	jsonStr1 := eorm.NewRecord().
		FromJson(`{"id": 1, "name": "张三"}`).
		Set("age", 25).
		Set("email", "zhangsan@example.com").
		ToJson()
	fmt.Printf("   结果: %s\n", jsonStr1)

	// 2. FromJson + 多次 Set + ToJson 链式调用
	fmt.Println("\n2. FromJson + 多次 Set + ToJson 链式调用")
	jsonStr2 := eorm.NewRecord().
		FromJson(`{"id": 1, "name": "张三"}`).
		Set("age", 25).
		Set("email", "zhangsan@example.com").
		Set("active", true).
		Set("created_at", "2024-01-01").
		ToJson()
	fmt.Printf("   结果: %s\n", jsonStr2)

	// 3. FromMap + Set + ToJson 链式调用
	fmt.Println("\n3. FromMap + Set + ToJson 链式调用")
	mapData := map[string]interface{}{
		"id":     1,
		"name":   "张三",
		"age":    25,
		"email":  "zhangsan@example.com",
		"active": true,
	}
	jsonStr3 := eorm.NewRecord().
		FromMap(mapData).
		Set("updated_at", "2024-01-01").
		Set("created_by", "admin").
		ToJson()
	fmt.Printf("   结果: %s\n", jsonStr3)

	// 4. FromStruct + Set + ToJson 链式调用
	fmt.Println("\n4. FromStruct + Set + ToJson 链式调用")
	type User struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Email  string `json:"email"`
		Active bool   `json:"active"`
	}

	user := User{
		ID:     1,
		Name:   "张三",
		Age:    25,
		Email:  "zhangsan@example.com",
		Active: true,
	}
	jsonStr4 := eorm.NewRecord().
		FromStruct(user).
		Set("updated_at", "2024-01-01").
		Set("created_by", "admin").
		ToJson()
	fmt.Printf("   结果: %s\n", jsonStr4)

	// 5. FromJson + Set + FromStruct + ToJson 链式调用
	fmt.Println("\n5. FromJson + Set + FromStruct + ToJson 链式调用")
	jsonStr5 := eorm.NewRecord().
		FromJson(`{"id": 1, "name": "张三"}`).
		Set("age", 25).
		FromStruct(User{
			ID:    2,
			Name:  "李四",
			Age:   30,
			Email: "lisi@example.com",
		}).
		ToJson()
	fmt.Printf("   结果: %s\n", jsonStr5)

	// 6. FromMap + Set + FromJson + ToJson 链式调用
	fmt.Println("\n6. FromMap + Set + FromJson + ToJson 链式调用")
	jsonStr6 := eorm.NewRecord().
		FromMap(map[string]interface{}{
			"id":   1,
			"name": "张三",
		}).
		Set("age", 25).
		FromJson(`{"id": 2, "name": "李四", "age": 30}`).
		ToJson()
	fmt.Printf("   结果: %s\n", jsonStr6)

	// 7. FromStruct + Set + FromMap + ToJson 链式调用
	fmt.Println("\n7. FromStruct + Set + FromMap + ToJson 链式调用")
	jsonStr7 := eorm.NewRecord().
		FromStruct(User{
			ID:   1,
			Name: "张三",
			Age:  25,
		}).
		Set("email", "zhangsan@example.com").
		FromMap(map[string]interface{}{
			"active":     true,
			"created_at": "2024-01-01",
		}).
		ToJson()
	fmt.Printf("   结果: %s\n", jsonStr7)

	// 8. FromJson + Set + Delete + ToJson 链式调用（Delete 不能链式调用，需要分开）
	fmt.Println("\n8. FromJson + Set + Delete + ToJson 链式调用")
	record8 := eorm.NewRecord().
		FromJson(`{"id": 1, "name": "张三", "age": 25, "email": "old@example.com"}`).
		Set("email", "new@example.com")
	record8.Delete("age")
	jsonStr8 := record8.ToJson()
	fmt.Printf("   结果: %s\n", jsonStr8)

	// 9. FromJson + Set + Clear + ToJson 链式调用（Clear 不能链式调用，需要分开）
	fmt.Println("\n9. FromJson + Set + Clear + Set + ToJson 链式调用")
	record9 := eorm.NewRecord().
		FromJson(`{"id": 1, "name": "张三", "age": 25}`).
		Set("email", "zhangsan@example.com")
	record9.Clear()
	jsonStr9 := record9.
		Set("new_field", "new_value").
		ToJson()
	fmt.Printf("   结果: %s\n", jsonStr9)

	// 10. FromJson + Set + Clone + ToJson 链式调用（Clone 不能链式调用，需要分开）
	fmt.Println("\n10. FromJson + Set + Clone + ToJson 链式调用")
	record10 := eorm.NewRecord().
		FromJson(`{"id": 1, "name": "张三", "age": 25}`).
		Set("email", "zhangsan@example.com")
	clonedRecord := record10.Clone()
	clonedRecord.Set("cloned", true)
	fmt.Printf("   原始: %s\n", record10.ToJson())
	fmt.Printf("   克隆: %s\n", clonedRecord.ToJson())

	// 11. FromJson + Set + ToMap + ToJson 链式调用（ToMap 不能链式调用，需要分开）
	fmt.Println("\n11. FromJson + Set + ToMap + ToJson 链式调用")
	record11 := eorm.NewRecord().
		FromJson(`{"id": 1, "name": "张三", "age": 25}`).
		Set("email", "zhangsan@example.com")
	map11 := record11.ToMap()
	fmt.Printf("   Map 数据: %v\n", map11)
	jsonStr11 := record11.ToJson()
	fmt.Printf("   结果: %s\n", jsonStr11)

	// 12. FromJson + Set + Keys + ToJson 链式调用（Keys 不能链式调用，需要分开）
	fmt.Println("\n12. FromJson + Set + Keys + ToJson 链式调用")
	record12 := eorm.NewRecord().
		FromJson(`{"id": 1, "name": "张三", "age": 25}`).
		Set("email", "zhangsan@example.com")
	keys12 := record12.Keys()
	fmt.Printf("   字段列表: %v\n", keys12)
	jsonStr12 := record12.ToJson()
	fmt.Printf("   结果: %s\n", jsonStr12)

	// 13. 综合示例：FromJson + Set + Get + Set + ToJson（Get 返回值，不能链式调用）
	fmt.Println("\n13. 综合示例：FromJson + Set + Get + Set + ToJson")
	record13 := eorm.NewRecord().
		FromJson(`{"id": 1, "name": "张三", "age": 25}`).
		Set("email", "zhangsan@example.com")

	name13 := record13.GetString("name")
	fmt.Printf("   获取 name: %s\n", name13)

	record13.Set("active", true)
	jsonStr13 := record13.ToJson()
	fmt.Printf("   结果: %s\n", jsonStr13)

	// 14. 综合示例：FromJson + Set + Has + Set + ToJson（Has 返回 bool，不能链式调用）
	fmt.Println("\n14. 综合示例：FromJson + Set + Has + Set + ToJson")
	record14 := eorm.NewRecord().
		FromJson(`{"id": 1, "name": "张三", "age": 25}`).
		Set("email", "zhangsan@example.com")

	hasName := record14.Has("name")
	fmt.Printf("   检查 name 是否存在: %v\n", hasName)

	record14.Set("active", true)
	jsonStr14 := record14.ToJson()
	fmt.Printf("   结果: %s\n", jsonStr14)

	// 15. 复杂综合示例
	fmt.Println("\n15. 复杂综合示例：FromJson + Set + FromMap + Set + FromStruct + ToJson")
	jsonStr15 := eorm.NewRecord().
		FromJson(`{"id": 1, "name": "张三"}`).
		Set("age", 25).
		FromMap(map[string]interface{}{
			"email":  "zhangsan@example.com",
			"active": true,
		}).
		Set("created_at", "2024-01-01").
		FromStruct(User{
			ID:   1,
			Name: "张三",
			Age:  25,
		}).
		ToJson()
	fmt.Printf("   结果: %s\n", jsonStr15)

	// 16. FromRecord 链式调用示例
	fmt.Println("\n16. FromRecord 链式调用示例")
	sourceRecord := eorm.NewRecord().
		FromJson(`{"id": 1, "name": "张三", "age": 25}`)

	jsonStr16 := eorm.NewRecord().
		FromRecord(sourceRecord).
		Set("email", "zhangsan@example.com").
		Set("active", true).
		ToJson()
	fmt.Printf("   结果: %s\n", jsonStr16)

	// 17. FromRecord + FromJson 综合链式调用
	fmt.Println("\n17. FromRecord + FromJson 综合链式调用")
	sourceRecord2 := eorm.NewRecord().
		FromJson(`{"id": 1, "name": "张三"}`)

	jsonStr17 := eorm.NewRecord().
		FromRecord(sourceRecord2).
		Set("age", 25).
		FromJson(`{"email": "zhangsan@example.com", "active": true}`).
		ToJson()
	fmt.Printf("   结果: %s\n", jsonStr17)

	// 18. Clone + 链式调用示例
	fmt.Println("\n18. Clone + 链式调用示例")
	originalRecord := eorm.NewRecord().
		FromJson(`{"id": 1, "name": "张三", "age": 25}`)
	fmt.Printf("   原始: %s\n", originalRecord.ToJson())
	clonedRecord2 := originalRecord.Clone()
	fmt.Println("   完整克隆 Record :", clonedRecord2.ToJson())
	jsonStr18 := clonedRecord2.
		Set("cloned", true).
		Set("email", "zhangsan@example.com").
		ToJson()

	fmt.Printf("   克隆并修改: %s\n", jsonStr18)

	// 19. FromRecord + Clone 综合示例
	fmt.Println("\n19. FromRecord + Clone 综合示例")
	sourceRecord3 := eorm.NewRecord().
		FromJson(`{"id": 1, "name": "张三", "age": 25, "profile": {"city": "北京"}}`)

	// 先使用 FromRecord 复制一份，再克隆一份，各自修改
	recordA := eorm.NewRecord().
		FromRecord(sourceRecord3).
		Set("type", "A").
		ToJson()

	recordB := sourceRecord3.Clone().
		Set("type", "B").
		ToJson()

	// 修改复制的记录不会影响原始记录
	recordAObj, _ := eorm.NewRecord().FromJson(recordA).GetRecord("profile")
	recordAObj.Set("city", "上海")

	fmt.Printf("   原始: %s\n", sourceRecord3.ToJson())
	fmt.Printf("   FromRecord复制并修改: %s\n", recordA)
	fmt.Printf("   克隆并修改: %s\n", recordB)
	fmt.Printf("   修改复制后的原始记录仍不变: %s\n", sourceRecord3.ToJson())

	// 20. 浅拷贝丢失数据示例
	fmt.Println("\n20. 浅拷贝丢失数据示例")
	// 创建一个包含嵌套对象的 Record
	originalRecord4 := eorm.NewRecord()
	originalRecord4.Set("name", "原始记录")
	originalRecord4.Set("profile", map[string]interface{}{
		"age":   25,
		"city":  "北京",
		"hobby": []string{"读书", "游泳"},
		"detail": map[string]interface{}{
			"level": "VIP",
			"score": 100,
		},
	})

	fmt.Printf("   原始记录: %s\n", originalRecord4.ToJson())

	// 使用 Clone 进行浅拷贝
	clonedRecord4 := originalRecord4.Clone()
	fmt.Printf("   克隆后的: %s\n", clonedRecord4.ToJson())

	// 直接修改克隆记录中的 map 值（浅拷贝会共享引用）
	profile, _ := clonedRecord4.GetRecord("profile")
	profile.Set("age", 30)
	profile.Set("city", "上海")
	detail, _ := profile.GetRecord("detail")
	detail.Set("level", "SVIP")

	fmt.Printf("   修改克隆记录后的克隆记录: %s\n", clonedRecord4.ToJson())
	fmt.Printf("   修改克隆记录后的原始记录: %s\n", originalRecord4.ToJson())
	fmt.Printf("   说明：由于是浅拷贝，修改克隆记录的嵌套对象会影响原始记录\n")

	// 21. 循环引用处理示例
	fmt.Println("\n21. 循环引用处理示例")
	// 创建一个循环引用的 Record
	circularRecord := eorm.NewRecord()
	circularRecord.Set("name", "循环引用测试")

	// 创建另一个 Record 并引用第一个 Record
	anotherRecord := eorm.NewRecord()
	anotherRecord.Set("title", "另一个记录")
	anotherRecord.Set("ref", circularRecord)

	// 将第一个 Record 也设置为引用第二个 Record，形成循环
	circularRecord.Set("back_ref", anotherRecord)

	// 克隆包含循环引用的 Record
	clonedCircular := circularRecord.Clone()
	fmt.Printf("   原始记录（包含循环引用）: %s\n", circularRecord.ToJson())
	fmt.Printf("   克隆后的记录（处理了循环引用）: %s\n", clonedCircular.ToJson())
	fmt.Printf("   说明：循环引用被检测并处理，避免了无限递归\n")

	fmt.Println("\n========== 示例完成 ==========")
	fmt.Println("\n注意：")
	fmt.Println("1. FromJson、FromMap、FromStruct、FromRecord 返回 *Record，可以链式调用")
	fmt.Println("2. Set 返回 *Record，可以链式调用")
	fmt.Println("3. Get、GetString、GetInt 等返回值，不能继续链式调用")
	fmt.Println("4. Delete、Clear、Has 等没有返回值或返回值，不能用于链式调用")
	fmt.Println("5. ToJson 返回 string，不能继续链式调用")
	fmt.Println("6. Clone 返回 *Record，可以链式调用，但需要先调用再链式")
	fmt.Println("7. 正确的链式调用：NewRecord().FromJson().Set().Set().ToJson()")
	fmt.Println("8. FromRecord 用于从另一个 Record 复制数据，是浅拷贝")
	fmt.Println("9. Clone 用于创建 Record 的浅拷贝")
	fmt.Println("10. FromRecord 和 Clone 都是浅拷贝，嵌套对象不会被复制")
	fmt.Println("11. 浅拷贝时，修改克隆记录的嵌套对象会影响原始记录")
	fmt.Println("12. 循环引用会被检测并处理，避免无限递归")
	fmt.Println("13. 循环引用的对象会被标记为 {\"$ref\": \"circular_reference\"}")
}
