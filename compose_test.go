package fp

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombine(t *testing.T) {
	isStringEven := CombineError(strconv.Atoi, even)
	result, err := isStringEven("4")
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestCompose(t *testing.T) {
	addFive := ComposeError(addOne, addOne, addOne, addOne, addOne)

	result, err := addFive(8)
	assert.NoError(t, err)
	assert.Equal(t, 13, result)

}

func even(i int) (bool, error) {
	return i%2 == 0, nil
}
func addOne(i int) (int, error) {
	return i + 1, nil
}
