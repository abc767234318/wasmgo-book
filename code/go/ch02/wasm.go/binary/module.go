package binary

// 魔数和版本号，前八字节
const (
	MagicNumber = 0x6D736100 // `\0asm`
	Version     = 0x00000001 // 1
)

const (
	SecCustomID = iota
	SecTypeID
	SecImportID
	SecFuncID
	SecTableID
	SecMemID
	SecGlobalID
	SecExportID
	SecStartID
	SecElemID
	SecCodeID
	SecDataID
)

// 导入段tag常量
const (
	ImportTagFunc   = 0
	ImportTagTable  = 1
	ImportTagMem    = 2
	ImportTagGlobal = 3
)

// 导出段tag 常量
const (
	ExportTagFunc   = 0
	ExportTagTable  = 1
	ExportTagMem    = 2
	ExportTagGlobal = 3
)

// 索引空间
type (
	TypeIdx   = uint32
	FuncIdx   = uint32
	TableIdx  = uint32
	MemIdx    = uint32
	GlobalIdx = uint32
	LocalIdx  = uint32 // 函数的内部变量索引空间(包含函数参数和局部变量)
	LabelIdx  = uint32 // 函数跳转标签索引空间
)

type Module struct {
	Magic      uint32
	Version    uint32
	CustomSecs []CustomSec
	TypeSec    []FuncType
	ImportSec  []Import
	FuncSec    []TypeIdx
	TableSec   []TableType
	MemSec     []MemType
	GlobalSec  []Global
	ExportSec  []Export
	StartSec   *FuncIdx
	ElemSec    []Elem
	CodeSec    []Code
	DataSec    []Data
}

//type TypeSec   = []FuncType
//type ImportSec = []Import
//type FuncSec   = []TypeIdx
//type TableSec  = []TableType
//type MemSec    = []MemType
//type GlobalSec = []Global
//type ExportSec = []Export
//type StartSec  = FuncIdx
//type ElemSec   = []Elem
//type CodeSec   = []Code
//type DataSec   = []Data

// 自定义段结构
type CustomSec struct {
	Name  string
	Bytes []byte // TODO
}

// 导入项的结构定义
type Import struct {
	Module string
	Name   string
	Desc   ImportDesc
}

// go中不支持C语言的联合体，所以把四种类型都加进来了，当然解析的时候只有一个成员有意义
type ImportDesc struct {
	Tag      byte
	FuncType TypeIdx    // tag=0
	Table    TableType  // tag=1
	Mem      MemType    // tag=2
	Global   GlobalType // tag=3
}

// 全局段类型，全局项需要指定全局变量类型、初始值
type Global struct {
	Type GlobalType
	Init Expr
}

// 导出段类型，函数、表、内存、全局变量
type Export struct {
	Name string
	Desc ExportDesc
}
type ExportDesc struct {
	Tag byte
	Idx uint32
}

// 元素段结构
type Elem struct {
	Table  TableIdx
	Offset Expr
	Init   []FuncIdx
}

// 代码段结构
type Code struct {
	Locals []Locals // 局部变量
	Expr   Expr // 字节码
}
type Locals struct {
	N    uint32
	Type ValType
}

// 数据段结构
type Data struct {
	Mem    MemIdx
	Offset Expr
	Init   []byte
}

func (code Code) GetLocalCount() uint64 {
	n := uint64(0)
	for _, locals := range code.Locals {
		n += uint64(locals.N)
	}
	return n
}
