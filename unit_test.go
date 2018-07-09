package bitcoinlib

import (
	"testing"
)

func TestUvarintToBytes(t *testing.T) {

	testCase(240, 1, t)
	testCase(241, 2, t)
	testCase(2287, 2, t)
	testCase(2288, 3, t)
	testCase(67823, 3, t)

}

func testCase(testValue uint64, expectedLen int, t *testing.T) {
	_, length := UvarintToBytes(uint64(testValue))
	if length != expectedLen {
		t.Error("error test", testValue)
	}
}
