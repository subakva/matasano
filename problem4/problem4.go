package problem4

// 4. Detect single-character XOR
//
// One of the 60-character strings at:
//
//   https://gist.github.com/3132713
//
// has been encrypted by single-character XOR. Find it. (Your code from
// #3 should help.)

import "bufio"
import "strings"
import "subakva/matasano/problem3"
import "subakva/matasano/utils"

func ScanXORDecrypt(scanner *bufio.Scanner) string {
  decoded, _ := problem3.RepeatingCharacterXORDecrypt(scanner.Text())
  if decoded != "" {
    return strings.TrimSpace(decoded)
  } else {
    return ""
  }
}

func DetectRepeatingCharacterXOR(filename string) (string) {
  return utils.ReadAndScan(filename, ScanXORDecrypt)
}
