package problems

// 4. Detect single-character XOR
//
// One of the 60-character strings at:
//
//   https://gist.github.com/3132713
//
// has been encrypted by single-character XOR. Find it. (Your code from
// #3 should help.)

// ------------------------------------------------------------
import "os"
import "bufio"

func DetectRepeatingCharacterXOR(filename string) (string) {
  wd, _ := os.Getwd();
  path := wd + "/" + filename

  file, err := os.Open(path)
  if err != nil { panic(err) }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    decoded, _ := RepeatingCharacterXORDecrypt(scanner.Text())
    if decoded != "" {
      return decoded
    }
  }
  return ""
}
