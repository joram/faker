package faker

import "testing"

func TestGetIPAddress(t *testing.T) {
	type test struct {
		input string
		seed  int64

		expectedOutput string
		expectedError  error
	}
	tests := []test{
		{
			input:          "",
			seed:           0,
			expectedOutput: "",
			expectedError:  ErrInvalidIPAddress,
		},
		{
			input:          "127.0.0",
			seed:           0,
			expectedOutput: "",
			expectedError:  ErrInvalidIPAddress,
		},
		{
			input:          "127.0.0.1",
			seed:           0,
			expectedOutput: "68.66.66.202",
			expectedError:  nil,
		},
		{
			input:          "127.0.0.1",
			seed:           1,
			expectedOutput: "37.243.243.239",
			expectedError:  nil,
		},
	}
	for _, test := range tests {
		ipAddress, err := GetIPAddress(test.input, test.seed)
		if ipAddress != test.expectedOutput {
			t.Errorf("Expected %s but got %s\n", test.expectedOutput, ipAddress)
		}
		if err != test.expectedError {
			t.Errorf("Expected %s but got %s\n", test.expectedError, err)
		}
	}
}
