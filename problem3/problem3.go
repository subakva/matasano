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

import "fmt"
import hex "encoding/hex"
import "subakva/matasano/problem2"
import "regexp"
import "os"

var spaces              = regexp.MustCompile(`(?i)[ ]`)
var alphabet            = regexp.MustCompile(`(?i)[a-z]`)
var vowels              = regexp.MustCompile(`(?i)[aeiou]`)
var consonants          = regexp.MustCompile(`(?i)[bcdfghjklmnpqrstvwxyz]`)
var commonLetters       = regexp.MustCompile(`(?i)[tnshrdlc]`)
var uncommonCharacters  = regexp.MustCompile(`[^a-zA-Z \.\,\'\;\:\-\?\!\n]`)
var debug               = os.Getenv("DEBUG") != ""

func EnableDebug() {
  debug = true
}

func DisableDebug() {
  debug = false
}

func MatchFrequency(testString string, expression *regexp.Regexp) float32 {
  numCharacters := len(testString)
  numMatches := len(expression.FindAllStringIndex(testString, numCharacters))
  return float32(numMatches) / float32(numCharacters)
}

// returns true if the string is likely english
func ProbablyEnglish(decodedString string) (bool) {
  vowelRatio        := MatchFrequency(decodedString, vowels)
  spaceRatio        := MatchFrequency(decodedString, spaces)
  alphabetRatio     := MatchFrequency(decodedString, alphabet)
  uncommonRatio     := MatchFrequency(decodedString, uncommonCharacters)
  consonantRatio    := MatchFrequency(decodedString, consonants)
  commonLetterRatio := MatchFrequency(decodedString, commonLetters)

  // if debug && alphabetRatio > 0.75 && vowelRatio > 0.25 && consonantRatio > 0.45 {
  // if debug && alphabetRatio > 0.75 {
  if debug && uncommonRatio < 0.05 {
    fmt.Printf("Given: %v\n", decodedString)
    fmt.Printf("  Alphabet        : %v\n", alphabetRatio)
    fmt.Printf("  Uncommon        : %v\n", uncommonRatio)
    fmt.Printf("  Consonants      : %v\n", consonantRatio)
    fmt.Printf("  Vowels          : %v\n", vowelRatio)
    fmt.Printf("  Common Letters  : %v\n", commonLetterRatio)
    fmt.Printf("  Spaces          : %v\n", spaceRatio)
    fmt.Println("----------------------------------------")
  }
  // if spaceRatio > 0.1 && vowelRatio > 0.2 {
  //   fmt.Printf("Space: %v Vowels: %v Decoded: %v\n", spaceRatio, vowelRatio, decodedString)
  // }

  // return vowelRatio > 0.2 && spaceRatio > 0.1
  // return alphabetRatio > 0.75 && vowelRatio >= 0.25 && consonantRatio > 0.45
  // return alphabetRatio >= 0.75 && vowelRatio >= 0.25 && consonantRatio >= 0.45 && uncommonRatio < 0.1
  return uncommonRatio < 0.05 && alphabetRatio > 0.6 && spaceRatio > 0 && vowelRatio > 0.17
  // return uncommonRatio < 0.05 && alphabetRatio > 0.65 //&& spaceRatio > 0 && vowelRatio > 0.2
}

func DumpDebug(key string, hexMessage string, decodedString string) {
  fmt.Printf("Key: %v\n", key)
  fmt.Printf("  Encrypted: %v\n", hexMessage)
  fmt.Printf("  Decrypted: %v\n", decodedString)
  fmt.Printf("  English? : %v\n", ProbablyEnglish(decodedString))
  fmt.Printf("  vowelRatio        :%v\n", MatchFrequency(decodedString, vowels))
  fmt.Printf("  spaceRatio        :%v\n", MatchFrequency(decodedString, spaces))
  fmt.Printf("  alphabetRatio     :%v\n", MatchFrequency(decodedString, alphabet))
  fmt.Printf("  uncommonRatio     :%v\n", MatchFrequency(decodedString, uncommonCharacters))
  fmt.Printf("  consonantRatio    :%v\n", MatchFrequency(decodedString, consonants))
  fmt.Printf("  commonLetterRatio :%v\n", MatchFrequency(decodedString, commonLetters))
}

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
