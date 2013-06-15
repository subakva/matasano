package problem6

import (
  "testing"
  "bytes"
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
  // filename := "problem5.b64.txt"
  expectedMessage := "UNKNOWN"
  expectedKey     := "UNKNOWN"
  assertions.CallAndAssertEquals(t, BreakRepeatingKeyXOR, []interface{}{filename}, []interface{}{expectedMessage, expectedKey})
}

func StringsToBytes(strings []string) [][]byte {
  bytes := make([][]byte, len(strings))
  for i := 0; i < len(strings); i++ {
    bytes[i] = []byte(strings[i])
  }
  return bytes
}

func BytesToStrings(bytes [][]byte) []string {
  strings := make([]string, len(bytes))
  for i := 0; i < len(bytes); i++ {
    strings[i] = string(bytes[i])
  }
  return strings
}

func AssertTransposeChunks(t *testing.T, chunks []string, expected []string) {
  byteChunks := StringsToBytes(chunks)
  actual := TransposeChunks(byteChunks)
  AssertStringArraysEqual(t, actual, expected)
}

func AssertChunkBytes(t *testing.T, chunkMe string, chunkSize int, expected []string) {
  actual := ChunkBytes([]byte(chunkMe), chunkSize)
  AssertStringArraysEqual(t, actual, expected)
}

func AssertStringArraysEqual(t *testing.T, actual [][]byte, expected []string) {
  actualStrings := BytesToStrings(actual)
  if len(actual) != len(expected) {
    t.Errorf("Actual length does not match expected: %v != %v", len(actual), len(expected))
    t.Errorf("Actual  : %v\n", actualStrings)
    t.Errorf("Expected: %v\n", expected)
  } else {
    for i := 0; i < len(actual); i++ {
      if ! bytes.Equal(actual[i], []byte(expected[i])) {
        t.Errorf("Actual does not match expected at index: %v", i)
        t.Errorf("Actual  : %v\n", actualStrings)
        t.Errorf("Expected: %v\n", expected)
        t.Errorf("Actual[%v]  : %v\n", i, actualStrings[i])
        t.Errorf("Expected[%v]: %v\n", i, expected[i])
      }
    }
  }
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
