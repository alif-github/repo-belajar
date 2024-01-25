package sum

import "github.com/stretchr/testify/mock"

type MockStruct struct {
	Client
	Mock mock.Mock
}

func (input *MockStruct) GetSum(add1 int, add2 int) int {
	args := input.Mock.Called(add1, add2)
	if args.Get(0) == nil {
		return 0
	}

	return args.Get(0).(int)
}
