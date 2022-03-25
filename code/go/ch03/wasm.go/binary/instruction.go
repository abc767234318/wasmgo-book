package binary

type Expr = []Instruction

// ָ������
type Instruction struct {
	Opcode byte
	Args   interface{} // go��Ŀսӿ��ܱ�����������ֵ��ȡ����ʱ����Ҫ��ȷ����
}

func (instr Instruction) GetOpname() string {
	return opnames[instr.Opcode]
}
func (instr Instruction) String() string {
	return opnames[instr.Opcode]
}

// block & loop
type BlockArgs struct {
	BT     BlockType // �����鷵��ֵ����
	Instrs []Instruction
}

// ifָ��ṹ
type IfArgs struct {
	BT      BlockType
	Instrs1 []Instruction
	Instrs2 []Instruction
}

// br_table ָ��ṹ
type BrTableArgs struct {
	Labels  []LabelIdx
	Default LabelIdx
}

// �ڴ����/�洢ϵ��ָ����Ҫָ���ڴ�ƫ�����Ͷ�����ʾ
type MemArg struct {
	Align  uint32
	Offset uint32
}
