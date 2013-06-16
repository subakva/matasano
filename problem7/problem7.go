package problem7

// 7. AES in ECB Mode
//
// The Base64-encoded content at the following location:
//
//     https://gist.github.com/3132853
//
// Has been encrypted via AES-128 in ECB mode under the key
//
//     "YELLOW SUBMARINE".
//
// (I like "YELLOW SUBMARINE" because it's exactly 16 bytes long).
//
// Decrypt it.
//
// Easiest way:
//
// Use OpenSSL::Cipher and give it AES-128-ECB as the cipher.

import "io/ioutil"
import "os"
import "subakva/matasano/utils"
import "bytes"
import aes "crypto/aes"

func AESECBDecrypt(message []byte, key []byte) []byte {
  cipher, cerr := aes.NewCipher([]byte(key))
  if cerr != nil { panic(cerr) }

  encrypted := bytes.NewBuffer(message)
  decrypted := bytes.NewBuffer(make([]byte, 0))

  bufferIn  := make([]byte, len(key))
  bufferOut := make([]byte, len(key))
  for {
    i, _ := encrypted.Read(bufferIn)
    if i == 0 { break }
    cipher.Decrypt(bufferOut, bufferIn)
    decrypted.Write(bufferOut)
  }

  return decrypted.Bytes()
}

func AESECBDecryptFile(filename string, key string) string {
  wd, _       := os.Getwd();
  path        := wd + "/" + filename
  encoded, _  := ioutil.ReadFile(path)
  decoded     := utils.DecodeBase64(encoded)

  return string(AESECBDecrypt(decoded, []byte(key)))
}
