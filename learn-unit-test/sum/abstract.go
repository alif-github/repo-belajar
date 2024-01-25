package sum

//var (
//	MyVariable     = normalStruct{}
//	MyVariableMock = mockStruct{}
//)

var ClientConfig Client

type Client interface {
	GetSum(int, int) int
}
