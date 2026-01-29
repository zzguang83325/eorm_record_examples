package main

import (
	"fmt"

	"github.com/zzguang83325/eorm"
)

// 示例12：获取切片
// 演示 Record 的切片获取方法：GetSlice、GetStringSlice、GetIntSlice、GetSliceByPath
func main() {
	fmt.Println("========== Record 获取切片示例 ==========")

	// 1. 从 JSON 创建包含数组的 Record
	record := eorm.NewRecord().FromJson(`{
		"name": "张三",
		"hobbies": ["读书", "游泳", "旅游"],
		"scores": [85, 90, 95],
		"tags": ["developer", "golang", "database"],
		"contact": {
			"phones": ["13800138000", "13900139000"],
			"emails": ["zhangsan@example.com", "zhangsan@work.com"]
		}
	}`)
	fmt.Printf("1. 从 JSON 创建 Record: %v\n\n", record.ToJson())

	// 2. GetSlice - 获取切片值
	fmt.Println("2. GetSlice - 获取切片值")
	hobbies, err := record.GetSlice("hobbies")
	if err != nil {
		fmt.Printf("   获取 hobbies 失败: %v\n", err)
	} else {
		fmt.Printf("   hobbies: %v\n", hobbies)
	}

	scores, err := record.GetSlice("scores")
	if err != nil {
		fmt.Printf("   获取 scores 失败: %v\n", err)
	} else {
		fmt.Printf("   scores: %v\n", scores)
	}
	fmt.Println()

	// 3. GetStringSlice - 获取字符串切片
	fmt.Println("3. GetStringSlice - 获取字符串切片")
	hobbyStrings, err := record.GetStringSlice("hobbies")
	if err != nil {
		fmt.Printf("   获取 hobbyStrings 失败: %v\n", err)
	} else {
		fmt.Printf("   hobbyStrings: %v\n", hobbyStrings)
		for i, hobby := range hobbyStrings {
			fmt.Printf("   [%d] %s\n", i, hobby)
		}
	}

	tags, err := record.GetStringSlice("tags")
	if err != nil {
		fmt.Printf("   获取 tags 失败: %v\n", err)
	} else {
		fmt.Printf("   tags: %v\n", tags)
		for i, tag := range tags {
			fmt.Printf("   [%d] %s\n", i, tag)
		}
	}
	fmt.Println()

	// 4. GetIntSlice - 获取整数切片
	fmt.Println("4. GetIntSlice - 获取整数切片")
	scoreInts, err := record.GetIntSlice("scores")
	if err != nil {
		fmt.Printf("   获取 scoreInts 失败: %v\n", err)
	} else {
		fmt.Printf("   scoreInts: %v\n", scoreInts)
		for i, score := range scoreInts {
			fmt.Printf("   [%d] %d\n", i, score)
		}
		total := 0
		for _, score := range scoreInts {
			total += score
		}
		average := float64(total) / float64(len(scoreInts))
		fmt.Printf("   总分: %d, 平均分: %.2f\n", total, average)
	}
	fmt.Println()

	// 5. GetSliceByPath - 通过路径获取切片
	fmt.Println("5. GetSliceByPath - 通过路径获取切片")
	phones, err := record.GetSliceByPath("contact.phones")
	if err != nil {
		fmt.Printf("   获取 phones 失败: %v\n", err)
	} else {
		fmt.Printf("   phones: %v\n", phones)
		for i, phone := range phones {
			fmt.Printf("   [%d] %v\n", i, phone)
		}
	}

	emails, err := record.GetSliceByPath("contact.emails")
	if err != nil {
		fmt.Printf("   获取 emails 失败: %v\n", err)
	} else {
		fmt.Printf("   emails: %v\n", emails)
		for i, email := range emails {
			fmt.Printf("   [%d] %v\n", i, email)
		}
	}
	fmt.Println()

	// 6. 错误处理 - 获取不存在的字段
	fmt.Println("6. 错误处理 - 获取不存在的字段")
	_, err = record.GetStringSlice("not_exist")
	if err != nil {
		fmt.Printf("   获取不存在的字段（预期错误）: %v\n", err)
	}

	_, err = record.GetIntSlice("not_exist")
	if err != nil {
		fmt.Printf("   获取不存在的字段（预期错误）: %v\n", err)
	}

	_, err = record.GetSliceByPath("not.exist")
	if err != nil {
		fmt.Printf("   获取不存在的路径（预期错误）: %v\n", err)
	}
	fmt.Println()

	// 7. 字符串分割功能
	fmt.Println("7. 字符串分割功能")
	record2 := eorm.NewRecord()
	record2.Set("comma_separated", "apple,banana,orange")
	record2.Set("semicolon_separated", "red;green;blue")
	record2.Set("pipe_separated", "cat|dog|bird")
	record2.Set("space_separated", "hello world go")

	commaSlice, _ := record2.GetStringSlice("comma_separated")
	fmt.Printf("   逗号分隔: %v\n", commaSlice)

	semicolonSlice, _ := record2.GetStringSlice("semicolon_separated")
	fmt.Printf("   分号分隔: %v\n", semicolonSlice)

	pipeSlice, _ := record2.GetStringSlice("pipe_separated")
	fmt.Printf("   竖线分隔: %v\n", pipeSlice)

	spaceSlice, _ := record2.GetStringSlice("space_separated")
	fmt.Printf("   空格分隔: %v\n", spaceSlice)
	fmt.Println()

	// 8. 单元素切片
	fmt.Println("8. 单元素切片")
	record3 := eorm.NewRecord()
	record3.Set("single_value", "hello")
	singleSlice, _ := record3.GetStringSlice("single_value")
	fmt.Printf("   单元素切片: %v\n", singleSlice)
	fmt.Println()

	// 9. 空切片处理
	fmt.Println("9. 空切片处理")
	record4 := eorm.NewRecord()
	record4.Set("empty_array", []interface{}{})
	emptySlice, _ := record4.GetSlice("empty_array")
	fmt.Printf("   空切片: %v (长度: %d)\n", emptySlice, len(emptySlice))
	fmt.Println()

	// 10. 混合类型切片
	fmt.Println("10. 混合类型切片")
	record5 := eorm.NewRecord()
	record5.Set("mixed", []interface{}{"string", 123, true, 45.67})
	mixedSlice, _ := record5.GetSlice("mixed")
	fmt.Printf("   混合类型切片: %v\n", mixedSlice)
	for i, item := range mixedSlice {
		fmt.Printf("   [%d] %v (类型: %T)\n", i, item, item)
	}

	fmt.Println("========== 测试完成 ==========")
}
