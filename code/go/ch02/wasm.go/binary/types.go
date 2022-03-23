package binary

import (
	"fmt"
	"strings"
)

// wasm 值类型
const (
	ValTypeI32 ValType = 0x7F // i32
	ValTypeI64 ValType = 0x7E // i64
	ValTypeF32 ValType = 0x7D // f32
	ValTypeF64 ValType = 0x7C // f64

	FtTag   = 0x60
	FuncRef = 0x70

	MutConst byte = 0
	MutVar   byte = 1
)

type ValType = byte
type BlockType = int32

// 内存类型只需描述内存的页数限制，定义成限制的别名即可
type MemType = Limits

// 定义函数类型结构体
type FuncType struct {
	Tag         byte
	ParamTypes  []ValType
	ResultTypes []ValType
}

// 表类型需要描述表的元素类型和元素数量的限制
type TableType struct {
	ElemType byte
	Limits   Limits
}

// 全局变量类型描述全局变量的类型和可变性
type GlobalType struct {
	ValType ValType
	Mut     byte
}

// 限制类型用于描述表的元素数量或内存页数的上下限
type Limits struct {
	Tag byte
	Min uint32
	Max uint32
}

func ValTypeToStr(vt ValType) string {
	switch vt {
	case ValTypeI32:
		return "i32"
	case ValTypeI64:
		return "i64"
	case ValTypeF32:
		return "f32"
	case ValTypeF64:
		return "f64"
	default:
		panic(fmt.Errorf("invalid valtype: %d", vt))
	}
}

func (ft FuncType) Equal(ft2 FuncType) bool {
	//return reflect.DeepEqual(ft, ft2)
	if len(ft.ParamTypes) != len(ft2.ParamTypes) {
		return false
	}
	if len(ft.ResultTypes) != len(ft2.ResultTypes) {
		return false
	}
	for i, vt := range ft.ParamTypes {
		if vt != ft2.ParamTypes[i] {
			return false
		}
	}
	for i, vt := range ft.ResultTypes {
		if vt != ft2.ResultTypes[i] {
			return false
		}
	}
	return true
}

// (i32,i32)->(i32)
func (ft FuncType) GetSignature() string {
	sb := strings.Builder{}
	sb.WriteString("(")
	for i, vt := range ft.ParamTypes {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(ValTypeToStr(vt))
	}
	sb.WriteString(")->(")
	for i, vt := range ft.ResultTypes {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(ValTypeToStr(vt))
	}
	sb.WriteString(")")
	return sb.String()
}

func (ft FuncType) String() string {
	return ft.GetSignature()
}
func (gt GlobalType) String() string {
	return fmt.Sprintf("{type: %s, mut: %d}",
		ValTypeToStr(gt.ValType), gt.Mut)
}
func (limits Limits) String() string {
	return fmt.Sprintf("{min: %d, max: %d}",
		limits.Min, limits.Max)
}
