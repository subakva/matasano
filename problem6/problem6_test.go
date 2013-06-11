package problem6

import (
  "testing"
  "subakva/matasano/assertions"
)

func TestBitCount(t *testing.T) {
  AssertBitCount(t, 0,   0)
  AssertBitCount(t, 1,   1)
  AssertBitCount(t, 2,   1)
  AssertBitCount(t, 4,   1)
  AssertBitCount(t, 8,   1)
  AssertBitCount(t, 16,  1)
  AssertBitCount(t, 32,  1)
  AssertBitCount(t, 64,  1)
  AssertBitCount(t, 128, 1)
  AssertBitCount(t, 255, 8)
}

func TestFactorial(t *testing.T) {
  AssertFactorial(t, 0, 1)
  AssertFactorial(t, 1, 1)
  AssertFactorial(t, 2, 2)
  AssertFactorial(t, 3, 6)
  AssertFactorial(t, 4, 24)
}

func TestCountCombinations(t *testing.T) {
  AssertCountCombinations(t, 1, 2, 0)
  AssertCountCombinations(t, 2, 2, 1)
  AssertCountCombinations(t, 4, 2, 6)
  AssertCountCombinations(t, 5, 2, 10)
}

func TestBreakRepeatingKeyXOR(t *testing.T) {
  filename := "problem6.txt"
  expectedMessage := "UNKNOWN"
  expectedKey     := "UNKNOWN"
  assertions.CallAndAssertEquals(t, BreakRepeatingKeyXOR, []interface{}{filename}, []interface{}{expectedMessage, expectedKey})
}

func AssertCountCombinations(t *testing.T, first int, second int, expected int) {
  assertions.CallAndAssertEquals(t, CountCombinations, []interface{}{first, second}, []interface{}{expected})
}

func AssertBitCount(t *testing.T, input int, expected int) {
  assertions.CallAndAssertEquals(t, BitCount, []interface{}{uint8(input)}, []interface{}{expected})
}

func AssertFactorial(t *testing.T, input int, expected int) {
  assertions.CallAndAssertEquals(t, Factorial, []interface{}{input}, []interface{}{expected})
}
