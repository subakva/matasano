package problem7

import (
  "testing"
  "subakva/matasano/assertions"
  // utils "subakva/matasano/utils"
)

func TestDecryptAESECB(t *testing.T) {
  filename := "problem7.b64.txt"
  key := "YELLOW SUBMARINE"
  expectedMessage := "I'm back and I'm ringin' the bell"
  actualMessage := AESECBDecryptFile(filename, key)
  assertions.AssertEquals(t, actualMessage[0:33], expectedMessage, "Messages did not match!")
}
