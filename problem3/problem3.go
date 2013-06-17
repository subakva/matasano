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

import (
  "fmt"
  "regexp"
  "os"
  "encoding/hex"
  "subakva/matasano/problem2"
)

var spaces              = regexp.MustCompile(`(?i)[ ]`)
var alphabet            = regexp.MustCompile(`(?i)[a-z]`)
var vowels              = regexp.MustCompile(`(?i)[aeiou]`)
var consonants          = regexp.MustCompile(`(?i)[bcdfghjklmnpqrstvwxyz]`)
var commonLetters       = regexp.MustCompile(`(?i)[tnshrdlc]`)
var commonCharacters    = regexp.MustCompile(`(?i)[a-z \.\,\'\;\:\-\?\!\n]`)
var debug               = os.Getenv("DEBUG") != ""

// Switches on debugging statements
func EnableDebug() {
  debug = true
}

// Switches off debugging statements
func DisableDebug() {
  debug = false
}

// Returns the ratio of matching to non-matching characters in the string.
func MatchFrequency(testString string, expression *regexp.Regexp) float32 {
  numCharacters := len(testString)
  numMatches := len(expression.FindAllStringIndex(testString, numCharacters))
  return float32(numMatches) / float32(numCharacters)
}

// Returns true if the string contains characters that appear to come from English sentences.
func ProbablyEnglish(decodedString string) (bool) {
  vowelRatio        := MatchFrequency(decodedString, vowels)
  spaceRatio        := MatchFrequency(decodedString, spaces)
  alphabetRatio     := MatchFrequency(decodedString, alphabet)
  commonRatio       := MatchFrequency(decodedString, commonCharacters)

  if debug && commonRatio > 0.95 {
    fmt.Printf("Given: %v\n", decodedString)
    fmt.Printf("  Alphabet        : %v\n", alphabetRatio)
    fmt.Printf("  Common          : %v\n", commonRatio)
    fmt.Printf("  Vowels          : %v\n", vowelRatio)
    fmt.Printf("  Spaces          : %v\n", spaceRatio)
    fmt.Println("----------------------------------------")
  }

  // These values make this particular set of problems work, but they're just magic numbers.
  // Without a better statistical model, these would probably require some tweaking for any
  // other set of problems.
  return commonRatio > 0.95 && alphabetRatio > 0.6 && spaceRatio > 0 && vowelRatio > 0.17
}

func DumpDebug(key string, hexMessage string, decodedString string) {
  fmt.Printf("Key: %v\n", key)
  fmt.Printf("  Encrypted: %v\n", hexMessage)
  fmt.Printf("  Decrypted: %v\n", decodedString)
  fmt.Printf("  English? : %v\n", ProbablyEnglish(decodedString))
  fmt.Printf("  vowelRatio        :%v\n", MatchFrequency(decodedString, vowels))
  fmt.Printf("  spaceRatio        :%v\n", MatchFrequency(decodedString, spaces))
  fmt.Printf("  alphabetRatio     :%v\n", MatchFrequency(decodedString, alphabet))
  fmt.Printf("  commonRatio       :%v\n", MatchFrequency(decodedString, commonCharacters))
  fmt.Printf("  consonantRatio    :%v\n", MatchFrequency(decodedString, consonants))
  fmt.Printf("  commonLetterRatio :%v\n", MatchFrequency(decodedString, commonLetters))
}

// Attempts to decrypt the XOR encrypted string by finding single-character keys that generate
// English-like results.
func RepeatingCharacterXORDecrypt(hexMessage string) (string, string) {
  for c := 32; c <= 126; c++ {
    key := string(c)
    comp := ""
    for i := 0; i < len(hexMessage) / 2; i++ {
      comp += key
    }
    hexComp       := hex.EncodeToString([]byte(comp))
    xorResult     := problem2.FixedXOR(hexMessage, hexComp)
    xorDecoded, _ := hex.DecodeString(xorResult)
    decodedString := string(xorDecoded)

    if ProbablyEnglish(decodedString) {
      if debug {
        fmt.Printf("Selected: %v\n", key)
        DumpDebug(key, hexMessage, decodedString)
      }
      return decodedString, key
    }
  }
  return "", ""
}
