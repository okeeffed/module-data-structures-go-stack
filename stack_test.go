package stack

import "testing"

func TestStackPushFunc(t *testing.T) {
	for _, tt := range pushTestCases {
		res, err := tt.inputQ.push(tt.input)

		if err != nil {
			t.Log(err)
			t.Errorf("FAIL: Returned an error")
			return
		}

		if res != tt.expected {
			t.Errorf("FAIL: Return value expected %d but got %d", tt.expected, res)
			return
		}

		for i, d := range tt.inputQ.data {
			if d != tt.expectedQ.data[i] {
				t.Errorf("FAIL: Queue data expected %d but got %d", tt.expectedQ.data[i], d)
				return
			}
		}

		t.Logf("PASS: Expected %d and got %d", tt.expected, res)
		t.Logf("PASS: Expected %+v and got %+v", tt.expectedQ.data, tt.inputQ.data)
	}
}

func TestStackPopFunc(t *testing.T) {
	for _, tt := range popTestCases {
		res, err := tt.inputQ.pop()

		if err != nil {
			t.Log(err)
			t.Errorf("FAIL: Returned an error")
			return
		}

		if res != tt.expected {
			t.Errorf("FAIL: Return value expected %d but got %d", tt.expected, res)
			return
		}

		for i, d := range tt.inputQ.data {
			if d != tt.expectedQ.data[i] {
				t.Errorf("FAIL: Queue data expected %d but got %d", tt.expectedQ.data[i], d)
				return
			}
		}

		t.Logf("PASS: Expected %d and got %d", tt.expected, res)
		t.Logf("PASS: Expected %+v and got %+v", tt.expectedQ.data, tt.inputQ.data)
	}
}

func TestStackPopHandlesError(t *testing.T) {
	for _, tt := range popErrorTestCases {
		res, err := tt.inputQ.pop()

		if err == nil {
			t.Errorf("FAIL: Did not correctly handle errors")
			return
		}

		if res != tt.expected {
			t.Errorf("FAIL: Return value expected %d but got %d", tt.expected, res)
			return
		}

		if len(tt.inputQ.data) != 0 {
			t.Errorf("FAIL: Queue length expected to be 0 if an error is thrown")
			return
		}

		t.Logf("PASS: Dequeue error correctly not nil and queue length is 0")
	}
}

func TestStackPeekFunc(t *testing.T) {
	for _, tt := range peekTestCases {
		res, err := tt.inputQ.peek()

		if err != nil {
			t.Log(err)
			t.Errorf("FAIL: Returned an error")
			return
		}

		if res != tt.expected {
			t.Errorf("FAIL: Return value expected %d but got %d", tt.expected, res)
			return
		}

		for i, d := range tt.inputQ.data {
			if d != tt.expectedQ.data[i] {
				t.Errorf("FAIL: Queue data expected %d but got %d", tt.expectedQ.data[i], d)
				return
			}
		}

		t.Logf("PASS: Expected %d and got %d", tt.expected, res)
		t.Logf("PASS: Expected %+v and got %+v", tt.expectedQ.data, tt.inputQ.data)

	}
}
