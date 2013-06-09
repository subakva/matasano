package problem3

// 3. Single-character XOR Cipher
//
// The hex encoded string:
//
//       1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
//
// ... has been XOR'd against a single character. Find the key, decrypt
// the message.
//
// Write code to do this for you. How? Devise some method for "scoring" a
// piece of English plaintext. (Character frequency is a good metric.)
// Evaluate each output and choose the one with the best score.
//
// Tune your algorithm until this works.

// import "fmt"
import hex "encoding/hex"
import "subakva/matasano/problem2"
import "regexp"

var vowels = regexp.MustCompile(`(?i)[aeiou]`)
var spaces = regexp.MustCompile(`(?i)[ ]`)

// returns true if the string is likely english
func ProbablyEnglish(decodedString string) (bool) {
  numVowels     := len(vowels.FindAllStringIndex(decodedString, len(decodedString)))
  numSpaces     := len(spaces.FindAllStringIndex(decodedString, len(decodedString)))
  numCharacters := len(decodedString)
  vowelRatio    := float32(numVowels) / float32(numCharacters)
  spaceRatio    := float32(numSpaces) / float32(numCharacters)

  // if spaceRatio > 0.1 && vowelRatio > 0.2 {
  //   fmt.Printf("Space: %v Vowels: %v Decoded: %v\n", spaceRatio, vowelRatio, decodedString)
  // }

  return vowelRatio > 0.2 && spaceRatio > 0.1
}

func RepeatingCharacterXORDecrypt(message string) (string, string) {
  // fmt.Printf("Match : %v\n", len(vowels.FindAllStringIndex("aaaaa", 0)))
  for c := 32; c <= 126; c++ {
    key := string(c)
    comp := ""
    for i := 0; i < len(message) / 2; i++ {
      comp += key
    }
    hexComp       := hex.EncodeToString([]byte(comp))
    xorResult     := problem2.FixedXOR(message, hexComp)
    xorDecoded, _ := hex.DecodeString(xorResult)
    decodedString := string(xorDecoded)

    if ProbablyEnglish(decodedString) {
      return decodedString, key
    }
  }
  return "", ""
}
