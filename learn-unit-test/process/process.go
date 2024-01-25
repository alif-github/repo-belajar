package process

import (
	"fmt"
	sum_sum "learn-unit-test/sum"
)

var ProcService = AProcessService{}

type AProcessService struct {
	Repository sum_sum.ClientConfig
}

func (input AProcessService) Process(num1, num2 int) (int, error) {
	result := input.Repository.GetSum(num1, num2)
	return result, nil
}

func NotTracked() {
	fmt.Println("Im Not Tracked")
}
