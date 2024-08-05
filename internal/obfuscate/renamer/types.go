package renamer

func isNonrenameableType(in string) bool {
	switch in {
	case "bool":
		return true
	case "int":
		return true
	case "int8":
		return true
	case "int16":
		return true
	case "int32":
		return true
	case "int64":
		return true
	case "uint":
		return true
	case "uint8":
		return true
	case "uint16":
		return true
	case "uint32":
		return true
	case "uint64":
		return true
	case "uintptr":
		return true
	case "float32":
		return true
	case "float64":
		return true
	case "complex64":
		return true
	case "complex128":
		return true
	case "string":
		return true
	case "byte":
		return true
	case "rune":
		return true
	case "error":
		return true
	default:
		return false
	}
}

func isNonrenameableValue(in string) bool {
	switch in {
	case "true":
		return true
	case "false":
		return true
	default:
		return false
	}
}

func isNonrenameableFuncCall(in string) bool {
	if isNonrenameableType(in) {
		return true
	}

	switch in {
	case "append":
		return true
	default:
		return false
	}
}
