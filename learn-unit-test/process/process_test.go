package process

//
//import (
//	"testing"
//)
//
//func TestProcessTabularMethod_Positive(t *testing.T) {
//	cases := []struct {
//		name        string
//		num1        int
//		num2        int
//		res         int
//		expectError bool
//	}{
//		{
//			name:        "Case 1 Small Number",
//			num1:        1,
//			num2:        2,
//			res:         3,
//			expectError: false,
//		},
//		{
//			name:        "Case 2 Big Number",
//			num1:        5,
//			num2:        10,
//			res:         15,
//			expectError: false,
//		},
//	}
//
//	for _, tc := range cases {
//		t.Run(tc.name, func(t *testing.T) {
//			res, err := ProcService.Process(tc.num1, tc.num2)
//			if err != nil && !tc.expectError {
//				t.Fatal(err)
//			}
//
//			if res != tc.res {
//				t.Fatalf("Expected result %d got %d", tc.res, res)
//			}
//		})
//	}
//}
