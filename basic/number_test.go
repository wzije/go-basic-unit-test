package tdd

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Basic test
func TestOddOrEven(t *testing.T) {
	assert.Equal(t, "20 is even number", OddOrEven(20))
	assert.Equal(t, "25 is odd number", OddOrEven(25))
	assert.Equal(t, "0 is even number", OddOrEven(0))
	assert.Equal(t, "-1 is odd number", OddOrEven(-1))
}

// Test using testtable
func TestOddOrEvenTestTable(t *testing.T) {

	attributes := []struct {
		num         int
		expectation string
	}{
		{num: 20, expectation: "20 is even number"},
		{num: 25, expectation: "25 is odd number"},
		{num: 0, expectation: "0 is even number"},
		{num: -1, expectation: "-1 is odd number"},
	}

	for _, d := range attributes {
		assert.Equal(t, d.expectation, OddOrEven(d.num))
	}
}

// Test using subtest
func testOddOrEvenSubtest(t *testing.T) {
	t.Run("test positive number", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			if i%2 == 0 {
				assert.Equal(t, fmt.Sprintf("%d is even number", i), OddOrEven(i))
			} else {
				assert.Equal(t, fmt.Sprintf("%d is odd number", i), OddOrEven(i))
			}
		}
	})

	t.Run("test negative number", func(t *testing.T) {
		for i := -100; i < 0; i-- {
			if i%2 == 0 {
				assert.Equal(t, fmt.Sprintf("%d is even number", i), OddOrEven(i))
			} else {
				assert.Equal(t, fmt.Sprintf("%d is odd number", i), OddOrEven(i))
			}
		}
	})
}

// Skip Test
func TestAddNeedsToBeSkip(t *testing.T) {
	t.Skip("this will be skipped")
}

// Skip long test
func TestCallToHoldSort(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip because calling db is way to long.")
	}
	<-time.After(3 * time.Second)
}
