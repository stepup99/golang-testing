package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort1(t *testing.T) {
	// initialization
	// execution
	// validation
	arrElm := GenerateArr(10)
	BubbleSort(arrElm)
	assert.NotNil(t, arrElm)
	assert.Len(t, arrElm, 10, "expected length 10 but got %v", len(arrElm))
	assert.EqualValues(t, 10, len(arrElm))
}
