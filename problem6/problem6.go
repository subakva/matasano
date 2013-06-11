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

import b64 "encoding/base64"
import "io/ioutil"
import "os"
import "fmt"
import "math"
import hex "encoding/hex"
import "subakva/matasano/problem3"

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

func HammingDistance(first string, second string) (distance int) {
  if len(first) != len(second) {
    panic("Cannot calculate Hamming distance unless the lengths match.")
  }
  for i := 0; i < len(first); i++ {
    b1 := []byte(first)[i]
    b2 := []byte(second)[i]
    distance += BitCount(b1 ^ b2) // count the number of bits in the XOR result
  }
  return
}

func GuessKeySize(decoded []byte, numChunks int) (likelyKeySize int) {
  minDistance      := float64(math.MaxFloat64)
  numCombinations  := CountCombinations(numChunks, 2)

  for keySize := 2; keySize <= 40; keySize++ {
    chunks := ChunkString(string(decoded), keySize)
    // chunks := make([]string, 4)
    // for i := 0; i < numChunks; i++ {
    //   startIndex := keySize * i
    //   endIndex   := keySize * (i + 1)
    //   chunks[i] = string(decoded[startIndex:endIndex])
    // }

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
    if averageDistance < minDistance {
      minDistance   = averageDistance
      likelyKeySize = keySize
    }
  }
  return
}

func ChunkString(chunkMe string, chunkSize int) []string {
  numChunks := len(string(chunkMe)) / chunkSize
  chunks    := make([]string, numChunks)
  for i := 0; i < numChunks; i++ {
    startIndex := chunkSize * i
    endIndex   := chunkSize * (i + 1)
    chunks[i] = string(chunkMe[startIndex:endIndex])
  }
  return chunks
}

func BreakRepeatingKeyXOR(filename string) (message string, key string) {
  wd, _ := os.Getwd();
  path := wd + "/" + filename
  encoded, _ := ioutil.ReadFile(path)
  decoded, _ := b64.StdEncoding.DecodeString(string(encoded))

  likelyKeySize := GuessKeySize(decoded, 4)
  fmt.Printf("likelyKeySize : %v\n", likelyKeySize)
  fmt.Printf("len(string(decoded)) : %v\n", len(string(decoded)))


  // decoded       = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
  // decoded       = []byte("THEQUICKBROWNFOXJUMPEDOVERTHELAZYDOG")
  // likelyKeySize = 2
  chunks            := ChunkString(string(decoded), likelyKeySize)
  transposedBlocks  := make([][]byte, likelyKeySize)
  for i := 0; i < likelyKeySize; i++ {
    maxLength := len(decoded) / likelyKeySize
    block := make([]byte, maxLength)
    for j := 0; j < maxLength; j++ {
      block[j] = chunks[j][i]
    }
    transposedBlocks[i] = block
    // fmt.Printf("transposedBlocks[%v] : %v\n", i, string(block))
  }

  for i := 0; i < likelyKeySize; i++ {
    // fmt.Printf("transposedBlocks[%v] : %v\n", i, string(transposedBlocks[i]))
    hexEncoded := hex.EncodeToString(transposedBlocks[i])
    // fmt.Printf("hexEncoded : %v\n", hexEncoded)
    // decoded, _ := hex.DecodeString(hexEncoded)
    // fmt.Printf("hexDecoded : %v\n", string(decoded))
    decrypted, keyChar := problem3.RepeatingCharacterXORDecrypt(hexEncoded)
    message += decrypted
    key     += keyChar
  }

  return message, key
}

