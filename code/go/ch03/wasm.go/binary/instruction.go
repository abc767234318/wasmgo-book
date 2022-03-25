package binary

type Expr = []Instruction

// 指令类型
type Instruction struct {
	Opcode byte
	Args   interface{} // go里的空接口能保存任意类型值，取出的时候需要明确类型
}

func (instr Instruction) GetOpname() string {
	return opnames[instr.Opcode]
}
func (instr Instruction) String() string {
	return opnames[instr.Opcode]
}

// block & loop
type BlockArgs struct {
	BT     BlockType // 决定块返回值类型
	Instrs []Instruction
}

// if指令结构
type IfArgs struct {
	BT      BlockType
	Instrs1 []Instruction
	Instrs2 []Instruction
}

// br_table 指令结构
type BrTableArgs struct {
	Labels  []LabelIdx
	Default LabelIdx
}

// 内存加载/存储系列指令需要指定内存偏移量和对齐提示
type MemArg struct {
	Align  uint32
	Offset uint32
}
