package main

import (
  "fmt"
  "subakva/matasano/problems"
  hex "encoding/hex"
)

func main() {
  hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
  decodedString, _ := hex.DecodeString(hexString)
  fmt.Print("\nProblem 1\n")
  fmt.Printf("Hex     : %v\n", hexString)
  fmt.Printf("Decoded : %v\n", string(decodedString))
  fmt.Printf("Base64  : %v\n", problems.HexToBase64(hexString))

  fmt.Print("\nProblem 2\n")
  hex1 := "1c0111001f010100061a024b53535009181c"
  hex2 := "686974207468652062756c6c277320657965"
  fmt.Printf("Input1  : %v\n", hex1)
  fmt.Printf("Input2  : %v\n", hex2)
  fmt.Printf("FixedXOR: %v\n", problems.FixedXOR(hex1, hex2))

  fmt.Print("\nProblem 3\n")
  message := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
  fmt.Printf("Message : %v\n", message)
  decoded, key := problems.DecipherSingleCharacterXOR(message)
  fmt.Printf("Key     : %v\n", key)
  fmt.Printf("Decoded : %v\n", decoded)
}
