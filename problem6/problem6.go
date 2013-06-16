package problem6

// 6. Break repeating-key XOR
//
// The buffer at the following location:
//
//  https://gist.github.com/3132752
//
// is base64-encoded repeating-key XOR. Break it.
//
// Here's how:
//
// a. Let KEYSIZE be the guessed length of the key; try values from 2 to
// (say) 40.
//
// b. Write a function to compute the edit distance/Hamming distance
// between two strings. The Hamming distance is just the number of
// differing bits. The distance between:
//
//   this is a test
//
// and:
//
//   wokka wokka!!!
//
// is 37.
//
// c. For each KEYSIZE, take the FIRST KEYSIZE worth of bytes, and the
// SECOND KEYSIZE worth of bytes, and find the edit distance between
// them. Normalize this result by dividing by KEYSIZE.
//
// d. The KEYSIZE with the smallest normalized edit distance is probably
// the key. You could proceed perhaps with the smallest 2-3 KEYSIZE
// values. Or take 4 KEYSIZE blocks instead of 2 and average the
// distances.
//
// e. Now that you probably know the KEYSIZE: break the ciphertext into
// blocks of KEYSIZE length.
//
// f. Now transpose the blocks: make a block that is the first byte of
// every block, and a block that is the second byte of every block, and
// so on.
//
// g. Solve each block as if it was single-character XOR. You already
// have code to do this.
//
// e. For each block, the single-byte XOR key that produces the best
// looking histogram is the repeating-key XOR key byte for that
// block. Put them together and you have the key.
//

import hex "encoding/hex"
import "fmt"
import "math"
import "subakva/matasano/problem3"
import "subakva/matasano/utils"

func BitCount(n uint8) (num int) {
  for i := uint8(0); i < 8; i++ {
    leftShift   := (1 << i)               // shift left to compare the bit in position i
    bitwiseAnd  := (n & uint8(leftShift)) // bitwise-and with the number to compare the bit in position i
    rightShift  := bitwiseAnd >> i        // shift the anded bit back to position 1, so that it is either 0 or 1
    num += int(rightShift)                // convert 0 or 1 back to an int and add it to the sum
  }
  return
}

func Factorial(n int) int {
  if n <= 1 {
    return 1
  } else {
    return n * Factorial(n - 1)
  }
}

func CountCombinations(n int, k int) int {
  return Factorial(n) / (Factorial(k) * Factorial(n - k))
}

func Float64Average(values []float64) float64 {
  sum := float64(0)
  for _, v := range values {
    sum += v
  }
  return sum / float64(len(values))
}

func HammingDistance(first []byte, second []byte) (distance int) {
  if len(first) != len(second) {
    panic("Cannot calculate Hamming distance unless the lengths match.")
  }
  for i := 0; i < len(first); i++ {
    b1 := first[i]
    b2 := second[i]
    distance += BitCount(b1 ^ b2) // count the number of bits in the XOR result
  }
  return
}

func GuessKeySize(bytes []byte, numChunks int) (likelyKeySize int) {
  minDistance      := float64(math.MaxFloat64)
  numCombinations  := CountCombinations(numChunks, 2)

  maxKeySize := 40
  maxKeySizeForNumChunks := len(bytes) / numChunks
  if maxKeySizeForNumChunks < maxKeySize {
    maxKeySize = maxKeySizeForNumChunks
  }

  for keySize := 2; keySize <= maxKeySize; keySize++ {
    chunks := utils.ChunkBytes(bytes, keySize)
    distances := make([]float64, numCombinations)
    di := 0
    for i := 0; i < numChunks; i++ {
      for j := i + 1; j < numChunks; j++ {
        distance      := HammingDistance(chunks[i], chunks[j])
        normDistance  := float64(distance) / float64(keySize)
        distances[di] = normDistance
        di++
      }
    }
    averageDistance := Float64Average(distances)
    // fmt.Printf("keySize, averageDistance  = %v, %v\n", keySize, averageDistance)
    // fmt.Printf("keySize, distances        = %v, %v\n", keySize, distances)
    if averageDistance < minDistance {
      minDistance   = averageDistance
      likelyKeySize = keySize
    }
  }
  return
}

func TransposeChunks(chunks [][]byte) (transposed [][]byte) {
  numChunks         := len(chunks)
  maxBytesPerChunk  := len(chunks[0])
  minBytesPerChunk  := len(chunks[len(chunks) - 1])

  transposed = make([][]byte, maxBytesPerChunk)
  for byteIndex := 0; byteIndex < maxBytesPerChunk; byteIndex++ {
    blockLength := numChunks
    if byteIndex >= minBytesPerChunk {
      blockLength = numChunks - 1
    }
    block := make([]byte, blockLength)
    for j := 0; j < numChunks; j++ {
      if byteIndex < len(chunks[j]) {
        block[j] = chunks[j][byteIndex]
      }
    }
    transposed[byteIndex] = block
  }
  // for i := 0; i < len(transposed); i++ {
  //   fmt.Printf("len(transposed[%v]) : %v\n", i, len(transposed[i]))
  //   fmt.Printf("transposed[%v] : %v\n", i, transposed[i])
  // }
  return
}

func ComposeParts(parts []string) (message string) {
  numParts := len(parts)
  // for i := 0; i < numParts; i++ {
  //   fmt.Printf("parts[%v] : %v\n", i, parts[i])
  // }
  for i := 0; i < len(parts[0]); i++ {
    for j := 0; j < numParts; j++ {
      if i < len(parts[j]) {
        message += string(parts[j][i])
      }
    }
  }
  return
}

func BreakRepeatingKeyXOR(filename string) (message string, key string) {
  encoded := utils.ReadRelative(filename)
  decoded := utils.DecodeBase64(encoded)

  likelyKeySize := GuessKeySize(decoded, 4)
  fmt.Printf(" => Likely Key Size: %v\n", likelyKeySize)
  chunks     := utils.ChunkBytes(decoded, likelyKeySize)
  transposed := TransposeChunks(chunks)

  messageParts := make([]string, likelyKeySize)
  for i := 0; i < likelyKeySize; i++ {
    hexEncoded := hex.EncodeToString(transposed[i])

    decrypted, keyChar := problem3.RepeatingCharacterXORDecrypt(hexEncoded)
    if keyChar == "" {
      fmt.Printf(" => Unable to find key at index: %v\n", i)
    }
    messageParts[i] = decrypted
    key += keyChar
  }

  message = ComposeParts(messageParts)

  return message, key
}

