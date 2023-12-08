package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFailSum(t *testing.T) {
	result := Sum(100, 100)

	require.Equal(t, 100, result, "result should be 200")

	fmt.Println("test failure")
}

func TestFailSum2(t *testing.T) {
	result := Sum(100, 100)

	assert.Equal(t, 190, result, "result should be 200")

	fmt.Println("test failsum2")
}

func TestSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		t.Fatal("result should be 20")
	}

	fmt.Println("Testing")

	result2 := Sum(12, 8)

	if result2 != 22 {
		t.Error("result should be 22")
	}
}

/*func TestSum2(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		t.Fatal("result should be 20")
	}

	fmt.Println("Testing")

	result2 := Sum(12, 8)

	if result2 != 22 {
		t.Error("result should be 22")
	}
}
*/
