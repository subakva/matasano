package problem8

// 8. Detecting ECB
//
// At the following URL are a bunch of hex-encoded ciphertexts:
//
//    https://gist.github.com/3132928
//
// One of them is ECB encrypted. Detect it.
//
// Remember that the problem with ECB is that it is stateless and
// deterministic; the same 16 byte plaintext block will always produce
// the same 16 byte ciphertext.

import (
	"bufio"
	"bytes"
	"subakva/matasano/utils"
)

// Returns true if the byte array contains repeated blocks of the given size
func HasRepeatedBlocks(decoded []byte, blockSize int) bool {
	chunks := utils.ChunkBytes(decoded, blockSize)
	for i := 0; i < len(chunks); i++ {
		for j := 0; j < len(chunks); j++ {
			if i == j {
				continue
			}
			if bytes.Equal(chunks[i], chunks[j]) {
				return true
			}
		}
	}
	return false
}

// Returns the text from scanner if the b64-decoded content has repeated blocks
func ScanRepeatedBlocks(scanner *bufio.Scanner) string {
	decoded := utils.DecodeBase64(scanner.Bytes())
	if HasRepeatedBlocks(decoded, 16) {
		return scanner.Text()
	} else {
		return ""
	}
}

// Opens the file and finds the first line that appears to be encrypted with ECB
func DetectECB(filename string) string {
	return utils.ReadAndScan(filename, ScanRepeatedBlocks)
}
