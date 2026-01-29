# Record 测试用例

本目录包含 Record 功能的各种测试用例，每个测试用例都在独立的目录中。

## 目录结构

```
records/
├── 01_basic_usage/          # 基础用法
├── 02_json_handling/          # JSON 处理
├── 03_get_record/            # GetRecord 功能
├── 04_merge_data/            # 数据合并
├── 05_type_conversion/        # 类型转换
├── 06_chaining/              # 链式调用
├── 07_error_handling/        # 错误处理
├── 08_method_chaining/        # 方法链式调用
├── 09_get_record_by_path/    # GetRecordByPath 功能
├── 10_get_string_by_path/     # GetStringByPath 功能
├── 11_deep_clone/            # 深拷贝功能
├── 12_get_slice/            # 获取切片功能
├── README.md                 # 本文件
├── go.mod                   # Go 模块文件（共享）
└── go.sum                   # Go 依赖校验文件（共享）
```

## 测试用例列表

### 01. 基础用法 (01_basic_usage/)
演示 Record 的基本操作：创建、设置、获取、删除、清空等

```bash
cd 01_basic_usage
go run main.go
```

**主要功能**：
- 创建空 Record
- 设置字段值
- 获取字段值
- 类型安全获取（GetString、GetInt 等）
- 检查字段是否存在
- 删除字段
- 获取所有字段名
- 清空 Record

---

### 02. JSON 处理 (02_json_handling/)
演示 Record 的 JSON 序列化和反序列化功能

```bash
cd 02_json_handling
go run main.go
```

**主要功能**：
- FromJson：从 JSON 字符串创建 Record
- ToJson：将 Record 转换为 JSON 字符串
- 复杂 JSON 结构处理
- 嵌套对象处理

---

### 03. GetRecord (03_get_record/)
演示从 Record 中获取嵌套 Record 的功能

```bash
cd 03_get_record
go run main.go
```

**主要功能**：
- GetRecord：获取单个嵌套 Record
- GetRecords：获取 Record 数组
- 嵌套 Record 的操作

---

### 04. 数据合并 (04_merge_data/)
演示多个 FromJson、FromMap、FromStruct 连续调用的合并功能

```bash
cd 04_merge_data
go run main.go
```

**主要功能**：
- FromJson + FromJson 合并
- FromMap + FromMap 合并
- FromStruct + FromStruct 合并
- 混合调用合并
- 字段覆盖规则

---

### 05. 类型转换 (05_type_conversion/)
演示各种数据类型的转换功能

```bash
cd 05_type_conversion
go run main.go
```

**主要功能**：
- GetString：转字符串
- GetInt、GetInt64、GetInt32：转整数
- GetFloat、GetFloat32：转浮点数
- GetBool：转布尔值
- GetBytes：转字节数组
- GetTime：转时间类型

---

### 06. 链式调用 (06_chaining/)
演示 Record 的链式调用功能

```bash
cd 06_chaining
go run main.go
```

**主要功能**：
- Set 方法链式调用
- FromJson 链式调用
- FromMap 链式调用
- FromStruct 链式调用
- 混合链式调用

---

### 07. 错误处理 (07_error_handling/)
演示 Record 的错误处理机制

```bash
cd 07_error_handling
go run main.go
```

**主要功能**：
- Get 系列方法的错误处理
- 路径访问的错误处理
- 类型转换错误处理
- 错误信息展示

---

### 08. 方法链式调用 (08_method_chaining/)
演示 FromJson、FromMap、FromStruct 的综合链式调用

```bash
cd 08_method_chaining
go run main.go
```

**主要功能**：
- 多个 FromJson 链式调用
- 多个 FromMap 链式调用
- 多个 FromStruct 链式调用
- 混合方法链式调用
- 复杂场景应用

---

### 09. GetRecordByPath (09_get_record_by_path/)
演示通过点分路径获取嵌套 Record 的功能

```bash
cd 09_get_record_by_path
go run main.go
```

**主要功能**：
- 多层嵌套路径访问
- 中间层 Record 获取
- 错误处理（路径不存在、空路径、中间节点不是 Record）
- 获取 Record 后继续操作
- 实际应用场景（配置文件读取）

---

### 10. GetStringByPath (10_get_string_by_path/)
演示通过点分路径获取嵌套字符串值的功能

```bash
cd 10_get_string_by_path
go run main.go
```

**主要功能**：
- 多层嵌套路径访问
- 类型自动转换（数字、布尔值等转字符串）
- Record 对象转 JSON 字符串
- 错误处理
- 实际应用场景

---

### 11. 深拷贝功能 (11_deep_clone/)
演示 Record 的深拷贝和浅拷贝功能

```bash
cd 11_deep_clone
go run main.go
```

**主要功能**：
- DeepClone：创建 Record 的深拷贝
- FromRecordDeep：从另一个 Record 深拷贝填充当前 Record
- Clone：创建 Record 的浅拷贝
- FromRecord：从另一个 Record 浅拷贝填充当前 Record
- 深拷贝 vs 浅拷贝对比
- 修改顶层字段（浅拷贝和深拷贝都不会影响原始记录）
- 修改嵌套对象（浅拷贝会影响原始记录，深拷贝不会）
- 多层嵌套对象的深拷贝
- 包含数组的深拷贝
- 包含嵌套 Record 的深拷贝
- FromRecordDeep 链式调用
- 空记录和 nil 源记录的处理
- 复杂数据结构的深拷贝
- 循环引用的深拷贝（避免无限递归）

---

### 12. 获取切片功能 (12_get_slice/)
演示 Record 的切片获取方法

```bash
cd 12_get_slice
go run main.go
```

**主要功能**：
- GetSlice：获取切片值（返回 []interface{}）
- GetStringSlice：获取字符串切片（返回 []string）
- GetIntSlice：获取整数切片（返回 []int）
- GetSliceByPath：通过点分路径获取嵌套切片
- 从 JSON 创建包含数组的 Record
- 遍历切片元素
- 字符串自动分割功能（支持逗号、分号、竖线、空格分隔符）
- 单元素切片处理
- 空切片处理
- 混合类型切片处理
- 错误处理（字段不存在、路径不存在）

## 说明

1. **独立运行**：每个测试用例都可以独立运行，互不干扰
2. **共享依赖**：所有测试用例共享根目录的 go.mod 和 go.sum
3. **完整示例**：每个测试用例都包含详细的功能演示和注释
4. **学习路径**：建议按照编号顺序逐步了解 Record 的功能
5. **实际应用**：测试用例包含实际应用场景，如配置文件读取
6. **类型安全**：推荐使用类型安全的方法（GetString、GetInt 等）而不是通用的 Get 方法

## 最佳实践

### 获取值的方法选择

| 方法 | 返回类型 | 推荐场景 |
|------|---------|---------|
| `GetString(key)` | `string` | 获取字符串值 |
| `GetInt(key)` | `int` | 获取整数值 |
| `GetInt64(key)` | `int64` | 获取 64 位整数值 |
| `GetFloat(key)` | `float64` | 获取浮点数值 |
| `GetBool(key)` | `bool` | 获取布尔值 |
| `GetRecord(key)` | `*Record` | 获取嵌套的 Record 对象 |
| `GetRecords(key)` | `[]*Record` | 获取 Record 数组 |
| `Get(key)` | `interface{}` | 通用方法，不推荐（除非特殊需求） |

### 拷贝方法选择

| 方法 | 类型 | 推荐场景 |
|------|------|---------|
| `Clone()` | 浅拷贝 | 只读、共享引用、性能敏感 |
| `FromRecord(src)` | 浅拷贝 | 只读、共享引用、性能敏感 |
| `DeepClone()` | 深拷贝 | 需要修改、需要独立、安全隔离 |
| `FromRecordDeep(src)` | 深拷贝 | 需要修改、需要独立、安全隔离 |

## 注意事项

- 运行测试前请确保已安装 Go 环境
- 每个测试用例都是独立的，可以单独运行
- 如需运行所有测试，请使用上面提供的脚本