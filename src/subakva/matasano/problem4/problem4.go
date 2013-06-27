package problem4

// 4. Detect single-character XOR
//
// One of the 60-character strings at:
//
//   https://gist.github.com/3132713
//
// has been encrypted by single-character XOR. Find it. (Your code from
// #3 should help.)

import (
	"bufio"
	"strings"
	"subakva/matasano/problem3"
	"subakva/matasano/utils"
)

// Returns the decrypted content of the scanner if it appears to be
// encrypted with a single character.
func ScanXORDecrypt(scanner *bufio.Scanner) string {
	decrypted, _ := problem3.RepeatingCharacterXORDecrypt(scanner.Text())
	if decrypted != "" {
		return strings.TrimSpace(decrypted)
	} else {
		return ""
	}
}

// Returns the decrypted content of the line in the file that
// was encrypted with a single-repeating-character XOR.
func DetectRepeatingCharacterXOR(filename string) string {
	return utils.ReadAndScan(filename, ScanXORDecrypt)
}
