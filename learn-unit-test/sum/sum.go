package sum

type NormalStruct struct {
	Client
}

func (input *NormalStruct) GetSum(add1 int, add2 int) int {
	return add1 + add2
}
