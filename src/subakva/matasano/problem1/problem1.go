package problem1

// 1. Convert hex to base64 and back.
//
// The string:
//
//   49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
//
// should produce:
//
//   SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t

import (
	"encoding/base64"
	"encoding/hex"
)

// Converts a hex string into base64 string
func HexToBase64(hexString string) string {
	decodedString, _ := hex.DecodeString(hexString)
	return base64.StdEncoding.EncodeToString([]byte(decodedString))
}

// Converts a base64 string into a hex string
func Base64ToHex(b64String string) string {
	decodedString, _ := base64.StdEncoding.DecodeString(b64String)
	return hex.EncodeToString([]byte(decodedString))
}
