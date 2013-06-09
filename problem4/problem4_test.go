package problem4

import (
  "testing"
)

func TestDetectRepeatingCharacterXOR(t *testing.T) {
  filename := "problem4.txt"
  expected := "Now that the party is jumping"
  actual   := DetectRepeatingCharacterXOR(filename)

  if actual != expected {
    t.Errorf("DetectRepeatingCharacterXOR(%v)", filename)
    t.Errorf("Expected: %v", expected)
    t.Errorf("Actual  : %v", actual)
    t.Errorf("Expected: %v", len(expected))
    t.Errorf("Actual  : %v", len(actual))
  }
}

