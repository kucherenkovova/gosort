package qsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type QsortSuite struct {
	suite.Suite
}

func (suite *QsortSuite) TestSortInt() {
	arr := []int{5, 2, 4, -8, 6, -1, 0, 3}
	expected := []int{-8, -1, 0, 2, 3, 4, 5, 6}

	Sort(arr)

	assert.Equal(suite.T(), expected, arr)
}

func (suite *QsortSuite) TestSortInt8() {
	arr := []int8{5, 2, 4, 6, 1, 3}
	expected := []int8{1, 2, 3, 4, 5, 6}

	Sort(arr)

	assert.Equal(suite.T(), expected, arr)
}

func (suite *QsortSuite) TestSortFloat64() {
	arr := []int64{5, 2, 4, 6, 1, 3}
	expected := []int64{1, 2, 3, 4, 5, 6}

	Sort(arr)

	assert.Equal(suite.T(), expected, arr)
}

func (suite *QsortSuite) TestSortEmptySlice() {
	arr := []int{}

	Sort(arr)

	assert.Empty(suite.T(), arr)
}

func (suite *QsortSuite) TestSortNil() {
	var arr []uint = nil
	require.NotPanics(suite.T(), func() {
		Sort(arr)
	})
	assert.Nil(suite.T(), arr)
}

func (suite *QsortSuite) TestSortAlreadySortedSlice() {
	arr := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}

	Sort(arr)

	assert.Equal(suite.T(), expected, arr)
}

func TestQsortSuite(t *testing.T) {
	suite.Run(t, new(QsortSuite))
}
