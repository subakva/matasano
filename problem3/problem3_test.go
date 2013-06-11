package problem3

import (
  "fmt"
  "testing"
  "subakva/matasano/assertions"
)

func TestProbablyEnglish(t *testing.T) {
  if ProbablyEnglish("01234567890") {
    t.Errorf("Numbers should not be english")
  }
  if !ProbablyEnglish("englishletterswithoutspaces") {
    t.Errorf("English letters without spaces should be English")
  }
  if !ProbablyEnglish("English should be English") {
    t.Errorf("English should be English")
  }
  if !ProbablyEnglish("Cooking MC's like a pound of bacon") {
    t.Errorf("Vanilla ICE should rap in English")
  }
}

func TestRepeatingCharacterXORDecrypt(t *testing.T) {
  message := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
  expectedMessage := "Cooking MC's like a pound of bacon"
  expectedKey := "X"
  actualMessage, actualKey := RepeatingCharacterXORDecrypt(message)

  errorMessage := fmt.Sprintf("RepeatingCharacterXORDecrypt(%v)", message)
  assertions.AssertEquals(t, actualMessage, expectedMessage, errorMessage)
  assertions.AssertEquals(t, actualKey, expectedKey, errorMessage)
}

