package merge

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type MergeSortSuite struct {
	suite.Suite
}

func (suite *MergeSortSuite) TestSortInt() {
	arr := []int{5, 2, 4, -8, 6, -1, 0, 3}
	expected := []int{-8, -1, 0, 2, 3, 4, 5, 6}

	Sort(arr)

	assert.Equal(suite.T(), expected, arr)
}

func (suite *MergeSortSuite) TestSortInt8() {
	arr := []int8{5, 2, 4, 6, 1, 3}
	expected := []int8{1, 2, 3, 4, 5, 6}

	Sort(arr)

	assert.Equal(suite.T(), expected, arr)
}

func (suite *MergeSortSuite) TestSortFloat64() {
	arr := []int64{5, 2, 4, 6, 1, 3}
	expected := []int64{1, 2, 3, 4, 5, 6}

	Sort(arr)

	assert.Equal(suite.T(), expected, arr)
}

func (suite *MergeSortSuite) TestSortEmptySlice() {
	arr := []int{}

	Sort(arr)

	assert.Empty(suite.T(), arr)
}

func (suite *MergeSortSuite) TestSortNil() {
	var arr []uint = nil
	require.NotPanics(suite.T(), func() {
		Sort(arr)
	})
	assert.Nil(suite.T(), arr)
}

func (suite *MergeSortSuite) TestSortAlreadySortedSlice() {
	arr := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}

	Sort(arr)

	assert.Equal(suite.T(), expected, arr)
}

func TestMergeSortSuite(t *testing.T) {
	suite.Run(t, new(MergeSortSuite))
}
