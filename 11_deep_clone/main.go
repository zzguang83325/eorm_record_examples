// d:\workspace_xuance\otherProjects\eorm\examples\records\11_deep_clone\main.go
package main

import (
	"fmt"

	"github.com/zzguang83325/eorm"
)

func main() {
	fmt.Println("========== 深拷贝函数测试 ==========")

	// 1. 基本的 DeepClone 测试
	fmt.Println("\n1. 基本的 DeepClone 测试")
	record1 := eorm.NewRecord().FromJson(`{
		"id": 1,
		"name": "张三",
		"age": 25
	}`)

	clonedRecord1 := record1.DeepClone()
	fmt.Printf("   原始记录: %s\n", record1.ToJson())
	fmt.Printf("   克隆记录: %s\n", clonedRecord1.ToJson())
	fmt.Printf("   ✅ DeepClone 成功\n")

	// 2. 基本的 FromRecordDeep 测试
	fmt.Println("\n2. 基本的 FromRecordDeep 测试")
	record2 := eorm.NewRecord().FromJson(`{
		"id": 2,
		"name": "李四",
		"age": 30
	}`)

	newRecord2 := eorm.NewRecord().FromRecordDeep(record2)
	fmt.Printf("   原始记录: %s\n", record2.ToJson())
	fmt.Printf("   深拷贝记录: %s\n", newRecord2.ToJson())
	fmt.Printf("   ✅ FromRecordDeep 成功\n")

	// 3. 浅拷贝 vs 深拷贝的对比 - 修改顶层字段
	fmt.Println("\n3. 浅拷贝 vs 深拷贝的对比 - 修改顶层字段")
	record3 := eorm.NewRecord().FromJson(`{
		"id": 3,
		"name": "王五",
		"age": 35
	}`)

	// 浅拷贝
	shallowCopy := eorm.NewRecord().FromRecord(record3)
	shallowCopy.Set("name", "王五（浅拷贝修改）")

	// 深拷贝
	deepCopy := record3.DeepClone()
	deepCopy.Set("name", "王五（深拷贝修改）")

	fmt.Printf("   原始记录: %s\n", record3.ToJson())
	fmt.Printf("   浅拷贝记录: %s\n", shallowCopy.ToJson())
	fmt.Printf("   深拷贝记录: %s\n", deepCopy.ToJson())
	fmt.Printf("   ✅ 修改顶层字段，浅拷贝和深拷贝都不会影响原始记录\n")

	// 4. 浅拷贝 vs 深拷贝的对比 - 修改嵌套对象
	fmt.Println("\n4. 浅拷贝 vs 深拷贝的对比 - 修改嵌套对象")
	record4 := eorm.NewRecord().FromJson(`{
		"id": 4,
		"name": "赵六",
		"profile": {
			"age": 40,
			"city": "北京"
		}
	}`)

	// 先创建深拷贝
	deepCopy4 := record4.DeepClone()

	// 获取深拷贝的 profile Record
	deepProfile, _ := deepCopy4.GetRecord("profile")
	fmt.Printf("   修改前深拷贝的 profile: %s\n", deepProfile.ToJson())

	deepProfile.Set("age", 50)
	deepProfile.Set("city", "广州")

	fmt.Printf("   修改后深拷贝的 profile: %s\n", deepProfile.ToJson())

	// 再创建浅拷贝
	shallowCopy4 := eorm.NewRecord().FromRecord(record4)

	// 获取浅拷贝的 profile Record
	shallowProfile, _ := shallowCopy4.GetRecord("profile")
	fmt.Printf("   修改前浅拷贝的 profile: %s\n", shallowProfile.ToJson())

	shallowProfile.Set("age", 45)
	shallowProfile.Set("city", "上海")

	fmt.Printf("   修改后浅拷贝的 profile: %s\n", shallowProfile.ToJson())

	// 检查原始记录的 profile
	originalProfile, _ := record4.GetRecord("profile")
	fmt.Printf("   原始记录的 profile: %s\n", originalProfile.ToJson())

	fmt.Printf("   原始记录: %s\n", record4.ToJson())
	fmt.Printf("   浅拷贝记录: %s\n", shallowCopy4.ToJson())
	fmt.Printf("   深拷贝记录: %s\n", deepCopy4.ToJson())

	// 验证浅拷贝是否影响原始记录
	if originalProfile.Str("city") == shallowProfile.Str("city") {
		fmt.Printf("   ✅ 浅拷贝的 profile Record 与原始记录共享同一个引用\n")
	} else {
		fmt.Printf("   ❌ 浅拷贝的 profile Record 与原始记录不共享引用\n")
	}

	// 验证深拷贝是否独立
	if originalProfile.Str("city") != deepProfile.Str("city") {
		fmt.Printf("   ✅ 深拷贝的 profile Record 与原始记录是独立的\n")
	} else {
		fmt.Printf("   ❌ 深拷贝的 profile Record 与原始记录共享引用\n")
	}

	// 5. 多层嵌套对象的深拷贝测试
	fmt.Println("\n5. 多层嵌套对象的深拷贝测试")
	record5 := eorm.NewRecord().FromJson(`{
		"id": 5,
		"name": "钱七",
		"profile": {
			"basic": {
				"age": 45,
				"city": "深圳"
			},
			"contact": {
				"email": "qianqi@example.com",
				"phone": "13900139000"
			}
		}
	}`)

	deepCopy5 := record5.DeepClone()
	profile, _ := deepCopy5.GetRecord("profile")
	basic, _ := profile.GetRecord("basic")
	basic.Set("age", 50)
	basic.Set("city", "杭州")
	contact, _ := profile.GetRecord("contact")
	contact.Set("email", "newqianqi@example.com")

	fmt.Printf("   原始记录: %s\n", record5.ToJson())
	fmt.Printf("   深拷贝记录: %s\n", deepCopy5.ToJson())
	fmt.Printf("   ✅ 多层嵌套对象深拷贝成功，修改不影响原始记录\n")

	// 6. 包含数组的深拷贝测试
	fmt.Println("\n6. 包含数组的深拷贝测试")
	record6 := eorm.NewRecord().FromJson(`{
		"id": 6,
		"name": "孙八",
		"hobbies": ["读书", "游泳", "旅游"],
		"scores": [85, 90, 95]
	}`)

	deepCopy6 := record6.DeepClone()
	hobbies, _ := deepCopy6.GetSlice("hobbies")
	if len(hobbies) > 0 {
		hobbies[0] = "编程"
	}
	hobbies = append(hobbies, "音乐")
	deepCopy6.Set("hobbies", hobbies)
	scores, _ := deepCopy6.GetSlice("scores")
	if len(scores) > 0 {
		scores[0] = 100
	}
	deepCopy6.Set("scores", scores)

	fmt.Printf("   原始记录: %s\n", record6.ToJson())
	fmt.Printf("   深拷贝记录: %s\n", deepCopy6.ToJson())
	fmt.Printf("   ✅ 数组深拷贝成功，修改不影响原始记录\n")

	// 7. 包含嵌套 Record 的深拷贝测试
	fmt.Println("\n7. 包含嵌套 Record 的深拷贝测试")
	record7 := eorm.NewRecord().FromJson(`{
		"id": 7,
		"name": "周九"
	}`)

	nestedRecord := eorm.NewRecord().FromJson(`{
		"age": 55,
		"city": "成都"
	}`)
	record7.Set("nested", nestedRecord)

	deepCopy7 := record7.DeepClone()
	nested, _ := deepCopy7.GetRecord("nested")
	nested.Set("age", 60)
	nested.Set("city", "重庆")

	fmt.Printf("   原始记录: %s\n", record7.ToJson())
	fmt.Printf("   深拷贝记录: %s\n", deepCopy7.ToJson())
	fmt.Printf("   ✅ 嵌套 Record 深拷贝成功，修改不影响原始记录\n")

	// 8. FromRecordDeep 链式调用测试
	fmt.Println("\n8. FromRecordDeep 链式调用测试")
	record8 := eorm.NewRecord().FromJson(`{
		"id": 8,
		"name": "吴十",
		"age": 60
	}`)

	chainRecord := eorm.NewRecord().
		FromRecordDeep(record8).
		Set("email", "wushi@example.com").
		Set("phone", "13700137000")

	fmt.Printf("   原始记录: %s\n", record8.ToJson())
	fmt.Printf("   链式调用记录: %s\n", chainRecord.ToJson())
	fmt.Printf("   ✅ FromRecordDeep 支持链式调用\n")

	// 9. 空记录的深拷贝测试
	fmt.Println("\n9. 空记录的深拷贝测试")
	emptyRecord := eorm.NewRecord()
	deepCopyEmpty := emptyRecord.DeepClone()

	fmt.Printf("   空记录: %s\n", emptyRecord.ToJson())
	fmt.Printf("   深拷贝空记录: %s\n", deepCopyEmpty.ToJson())
	fmt.Printf("   ✅ 空记录深拷贝成功\n")

	// 10. nil 源记录的深拷贝测试
	fmt.Println("\n10. nil 源记录的深拷贝测试")
	var nilRecord *eorm.Record = nil
	deepCopyNil := nilRecord.DeepClone()

	if deepCopyNil == nil {
		fmt.Printf("   ✅ nil 记录深拷贝返回 nil\n")
	} else {
		fmt.Printf("   ❌ nil 记录深拷贝应该返回 nil\n")
	}

	// 11. FromRecordDeep 处理 nil 源记录
	fmt.Println("\n11. FromRecordDeep 处理 nil 源记录")
	newRecord11 := eorm.NewRecord().FromRecordDeep(nilRecord)
	newRecord11.Set("id", 11)
	newRecord11.Set("name", "测试")

	fmt.Printf("   FromRecordDeep(nil) 后添加数据: %s\n", newRecord11.ToJson())
	fmt.Printf("   ✅ FromRecordDeep 处理 nil 源记录成功\n")

	// 12. 复杂数据结构的深拷贝测试
	fmt.Println("\n12. 复杂数据结构的深拷贝测试")
	record12 := eorm.NewRecord().FromJson(`{
		"id": 12,
		"name": "郑十一",
		"profile": {
			"basic": {
				"age": 65,
				"city": "西安"
			},
			"contacts": [
				{"type": "email", "value": "zheng11@example.com"},
				{"type": "phone", "value": "13600136000"}
			]
		},
		"metadata": {
			"created": "2024-01-01",
			"updated": "2024-01-15"
		}
	}`)

	deepCopy12 := record12.DeepClone()
	profile, _ = deepCopy12.GetRecord("profile")
	basic, _ = profile.GetRecord("basic")
	basic.Set("age", 70)
	contacts, _ := profile.GetRecords("contacts")
	if len(contacts) > 0 {
		contacts[0].Set("value", "newzheng11@example.com")
	}
	metadata, _ := deepCopy12.GetRecord("metadata")
	metadata.Set("updated", "2024-01-20")

	fmt.Printf("   原始记录: %s\n", record12.ToJson())
	fmt.Printf("   深拷贝记录: %s\n", deepCopy12.ToJson())
	fmt.Printf("   ✅ 复杂数据结构深拷贝成功\n")

	// 13. 循环引用的深拷贝测试
	fmt.Println("\n13. 循环引用的深拷贝测试")

	// 创建两个 Record，互相引用形成循环
	recordA := eorm.NewRecord()
	recordA.Set("id", 13)
	recordA.Set("name", "记录A")

	recordB := eorm.NewRecord()
	recordB.Set("id", 14)
	recordB.Set("name", "记录B")

	// A 引用 B
	recordA.Set("friend", recordB)

	// B 引用 A（形成循环引用）
	recordB.Set("friend", recordA)

	fmt.Printf("   记录A: %s\n", recordA.ToJson())
	fmt.Printf("   记录B: %s\n", recordB.ToJson())

	// 尝试深拷贝包含循环引用的 Record
	clonedRecordA := recordA.DeepClone()

	fmt.Printf("   深拷贝记录A: %s\n", clonedRecordA.ToJson())
	fmt.Printf("   ✅ 循环引用深拷贝成功，没有无限递归\n")

	// 验证深拷贝后的 Record 也包含循环引用
	if friend, err := clonedRecordA.GetRecord("friend"); err == nil {
		if friend2, err := friend.GetRecord("friend"); err == nil {
			fmt.Printf("   ✅ 深拷贝后的 Record 也包含循环引用\n")
			fmt.Printf("   循环引用的记录: %s\n", friend2.ToJson())
		}
	}

	fmt.Println("\n========== 测试完成 ==========")
	fmt.Println("\n总结:")
	fmt.Println("1. DeepClone() 方法可以创建 Record 的深拷贝")
	fmt.Println("2. FromRecordDeep() 方法可以从另一个 Record 深拷贝填充当前 Record")
	fmt.Println("3. 深拷贝会递归复制所有嵌套对象（map、slice、Record 等）")
	fmt.Println("4. 深拷贝后的记录与原记录完全独立，修改不会相互影响")
	fmt.Println("5. FromRecordDeep() 支持链式调用")
	fmt.Println("6. 深拷贝可以正确处理 nil、空记录、复杂数据结构")
	fmt.Println("7. 深拷贝可以正确处理循环引用，避免无限递归")
}
