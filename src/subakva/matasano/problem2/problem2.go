package problem2

// 2. Fixed XOR
//
// Write a function that takes two equal-length buffers and produces
// their XOR sum.
//
// The string:
//
//  1c0111001f010100061a024b53535009181c
//
// ... after hex decoding, when xor'd against:
//
//  686974207468652062756c6c277320657965
//
// ... should produce:
//
//  746865206b696420646f6e277420706c6179

import "encoding/hex"

// Calculates the XOR of two hex-encoded strings
func FixedXOR(hex1 string, hex2 string) string {
  if len(hex1) != len(hex2) { panic("FixedXOR: String lengths must match!") }

  bytes1, _ := hex.DecodeString(hex1)
  bytes2, _ := hex.DecodeString(hex2)

  xored := make([]byte, len(bytes1))
  for i := 0; i < len(bytes1); i++ {
    xored[i] = bytes1[i] ^ bytes2[i]
  }
  return hex.EncodeToString(xored)
}
