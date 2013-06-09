package problem5

// 5. Repeating-key XOR Cipher
//
// Write the code to encrypt the string:
//
//   Burning 'em, if you ain't quick and nimble
//   I go crazy when I hear a cymbal
//
// Under the key "ICE", using repeating-key XOR. It should come out to:
//
//   0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f
//
// Encrypt a bunch of stuff using your repeating-key XOR function. Get a
// feel for it.

import hex "encoding/hex"
import "subakva/matasano/problem2"

func RepeatingKeyXOREncrypt(message string, key string) (string) {
  mask := ""
  for len(mask) < len(message) {
    mask += key
  }
  mask = mask[0:len(message)]
  hexMessage := hex.EncodeToString([]byte(message))
  hexMask    := hex.EncodeToString([]byte(mask))
  return problem2.FixedXOR(hexMessage, hexMask)
}
