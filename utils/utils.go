package utils

import b64 "encoding/base64"
import hex "encoding/hex"
import "strings"
// import "fmt"

// Returns true if the int is in the array of ints
func IntInArray(haystack []int, needle int) bool {
  for i := 0; i < len(haystack); i++ {
    if haystack[i] == needle { return true }
  }
  return false
}

// Decodes a b64-encoded byte array
func DecodeBase64(encoded []byte) []byte {
  decoded, err := b64.StdEncoding.DecodeString(string(encoded))
  if err != nil { panic(err) }
  return decoded
}

// Decodes a hex-encoded byte array
func DecodeHex(encoded []byte) []byte {
  trimmed := strings.TrimSpace(string(encoded))
  decoded, err := hex.DecodeString(trimmed)
  if err != nil { panic(err) }
  return decoded
}

// Converts an array of strings into an array of byte arrays
func StringsToBytes(strings []string) [][]byte {
  bytes := make([][]byte, len(strings))
  for i := 0; i < len(strings); i++ {
    bytes[i] = []byte(strings[i])
  }
  return bytes
}

// Converts an array of byte arrays into an array of strings
func BytesToStrings(bytes [][]byte) []string {
  strings := make([]string, len(bytes))
  for i := 0; i < len(bytes); i++ {
    strings[i] = string(bytes[i])
  }
  return strings
}

// func FirstLine(s string) string {
//   parts := strings.SplitN(s, "\n", 1)
//   fmt.Println(parts)
//   return parts[0]
// }

// func RatioCeil(total int, part int) int {
//   return int(math.Ceil(float64(total) / float64(part)))
// }

// func IntMin(first int, second int) int {
//   if first <= second {
//     return first
//   } else {
//     return second
//   }
// }

