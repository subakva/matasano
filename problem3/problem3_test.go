package problem3

import (
  "testing"
)

func TestRepeatingCharacterXORDecrypt(t *testing.T) {
  message := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
  expectedMessage := "Cooking MC's like a pound of bacon"
  expectedKey := "X"
  actualMessage, actualKey := RepeatingCharacterXORDecrypt(message)

  if actualMessage != expectedMessage {
    t.Errorf("RepeatingCharacterXORDecrypt(%v)", message)
    t.Errorf("Expected: %v", expectedMessage)
    t.Errorf("Actual  : %v", actualMessage)
  }

  if actualKey != expectedKey {
    t.Errorf("RepeatingCharacterXORDecrypt(%v)", message)
    t.Errorf("Expected: %v", expectedKey)
    t.Errorf("Actual  : %v", actualKey)
  }

}

