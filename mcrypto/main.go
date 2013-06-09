package main

import (
  "fmt"
  "subakva/matasano/problem1"
  "subakva/matasano/problem2"
  "subakva/matasano/problems"
  hex "encoding/hex"
)

func runProblem1() {
  hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
  decodedString, _ := hex.DecodeString(hexString)
  fmt.Print("\nProblem 1\n")
  fmt.Printf("Hex     : %v\n", hexString)
  fmt.Printf("Decoded : %v\n", string(decodedString))
  fmt.Printf("Base64  : %v\n", problem1.HexToBase64(hexString))
}

func runProblem2() {
  fmt.Print("\nProblem 2\n")
  hex1 := "1c0111001f010100061a024b53535009181c"
  hex2 := "686974207468652062756c6c277320657965"
  fmt.Printf("Input1  : %v\n", hex1)
  fmt.Printf("Input2  : %v\n", hex2)
  fmt.Printf("FixedXOR: %v\n", problem2.FixedXOR(hex1, hex2))
}

func runProblem3() {
  fmt.Print("\nProblem 3\n")
  message := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
  fmt.Printf("Message : %v\n", message)
  decoded, key := problems.RepeatingCharacterXORDecrypt(message)
  fmt.Printf("Key     : %v\n", key)
  fmt.Printf("Decoded : %v\n", decoded)
}

func runProblem4() {
  fmt.Print("\nProblem 4\n")
  filename := "src/subakva/matasano/problems/problem4.txt"
  fmt.Printf("Searching: %v\n", filename)
  detected := problems.DetectRepeatingCharacterXOR(filename)
  fmt.Printf("Detected : %v\n", detected)
}

func runProblem5() {
  fmt.Print("\nProblem 5\n")
  expected := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
  message := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
  key := "ICE"
  encrypted := problems.RepeatingKeyXOREncrypt(message, key)
  fmt.Printf("Expecting: %v\n", expected)
  fmt.Printf("Encrypted: %v\n", encrypted)
  // encryptAndPrint("CAD", "CAD")
  // encryptAndPrint("ABRACADABRA", "CAD")
  // encryptAndPrint("ABRACADABRA", "ABRACADABRA")
  // encryptAndPrint("AAA", "A")
  // encryptAndPrint("BAA", "A")
  // encryptAndPrint("aaa", "A")
  // encryptAndPrint("baa", "A")
  // encryptAndPrint("aaa", "AAA")
  // encryptAndPrint("baa", "AAA")
}

func encryptAndPrint(message string, key string) {
  encrypted := problems.RepeatingKeyXOREncrypt(message, key)
  fmt.Printf("Message  : %v\n", message)
  fmt.Printf("Key      : %v\n", key)
  fmt.Printf("Encrypted: %v\n", encrypted)
}

func main() {
  runProblem1()
  runProblem2()
  runProblem3()
  runProblem4()
  runProblem5()
}
