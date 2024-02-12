package operation

const (
	OpPush = iota
	OpPlus
	OpMinus
	OpMultiply
	OpDivision
	OpEqual
	OpLess
	OpGreater
	OpDump
	OpCount
)

type Operation struct {
	Code  int
	Value int64
}

func Push(value int64) Operation {
	return Operation{
		Code:  OpPush,
		Value: value,
	}
}

func Plus() Operation {
	return Operation{
		Code: OpPlus,
	}
}

func Minus() Operation {
	return Operation{
		Code: OpMinus,
	}
}

func Multiply() Operation {
	return Operation{
		Code: OpMultiply,
	}
}

func Division() Operation {
	return Operation{
		Code: OpDivision,
	}
}

func Equal() Operation {
	return Operation{
		Code: OpEqual,
	}
}

func Less() Operation {
	return Operation{
		Code: OpLess,
	}
}

func Greater() Operation {
	return Operation{
		Code: OpGreater,
	}
}

func Dump() Operation {
	return Operation{
		Code: OpDump,
	}
}
