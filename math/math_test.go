package math

import (
  "testing"
)

var data = []int{
 4, 7, -2, -5, 9, 2, 5, 2, 1, -8,
}

var min3 = []int{
 -2, -5, -5, -5, 2, 2, 1, -8,
}

var min2 = []int{
 4, -2, -5, -5, 2, 2, 2, 1, -8,
}

func Test_MinInt_works_with_tuples(t *testing.T) {
  for i, v := range data {
    if i == 0 {
      continue
    }
    v2 := data[i-1]
    expected := min2[i-1]
    testMin([]int{v, v2}, expected, t)
    testMin([]int{v2, v}, expected, t)
  }
}

func Test_MinInt_works_with_triples(t *testing.T) {
  for i, v := range data {
    if i < 2 {
      continue
    }
    v2 := data[i-2]
    v3 := data[i-1]
    expected := min3[i-2]
    testMin([]int{v, v2, v3}, expected, t)
    testMin([]int{v, v3, v2}, expected, t)
    testMin([]int{v2, v, v3}, expected, t)
    testMin([]int{v2, v3, v}, expected, t)
    testMin([]int{v3, v, v2}, expected, t)
    testMin([]int{v3, v2, v}, expected, t)
  }
}

func testMin(testData []int, expected int, t *testing.T) {
  if actual:= MinInt(testData...); actual != expected {
    t.Errorf("MinInt(%v) = %d, expected %d", testData, actual, expected)
  }
}
