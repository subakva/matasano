package problem6

import (
  "testing"
  "subakva/matasano/assertions"
  utils "subakva/matasano/utils"
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

func TestChunkBytes(t *testing.T) {
  AssertChunkBytes(t, "aabbcc", 2, []string{"aa","bb", "cc"})
  AssertChunkBytes(t, "aabbc",  2, []string{"aa","bb", "c"})
}

func TestTransposeChunks(t *testing.T) {
  AssertTransposeChunks(t, []string{"aa", "bb", "cc"}, []string{"abc", "abc"})
  AssertTransposeChunks(t, []string{"aa", "bb", "c"}, []string{"abc", "ab"})
}

func TestHammingDistance(t *testing.T) {
  input1 := []byte("this is a test")
  input2 := []byte("wokka wokka!!!")
  assertions.CallAndAssertEquals(t, HammingDistance, []interface{}{input1, input2}, []interface{}{37})
}

func TestBreakRepeatingKeyXOR(t *testing.T) {
  filename := "problem6.txt"
  expectedMessage := "I'm back and I'm ringin' the bell"
  expectedKey     := "Terminator X: Bring the noise"
  actualMessage, actualKey := BreakRepeatingKeyXOR(filename)
  assertions.AssertEquals(t, actualKey, expectedKey, "Key did not match!")
  assertions.AssertEquals(t, actualMessage[0:33], expectedMessage, "Messages did not match!")
}

func AssertTransposeChunks(t *testing.T, chunks []string, expected []string) {
  byteChunks := utils.StringsToBytes(chunks)
  actual := TransposeChunks(byteChunks)
  assertions.AssertStringArraysEqual(t, actual, expected)
}

func AssertChunkBytes(t *testing.T, chunkMe string, chunkSize int, expected []string) {
  actual := utils.ChunkBytes([]byte(chunkMe), chunkSize)
  assertions.AssertStringArraysEqual(t, actual, expected)
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
