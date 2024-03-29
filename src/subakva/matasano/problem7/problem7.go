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

import (
	"bytes"
	"crypto/aes"
	"subakva/matasano/utils"
)

// Decrypts the given byte array with the given key and returns the result.
func AESECBDecrypt(message []byte, key []byte) []byte {
	cipher, cerr := aes.NewCipher([]byte(key))
	if cerr != nil {
		panic(cerr)
	}

	encrypted := bytes.NewBuffer(message)
	decrypted := bytes.NewBuffer(make([]byte, 0))

	bufferIn := make([]byte, len(key))
	bufferOut := make([]byte, len(key))
	for {
		i, _ := encrypted.Read(bufferIn)
		if i == 0 {
			break
		}
		cipher.Decrypt(bufferOut, bufferIn)
		decrypted.Write(bufferOut)
	}

	return decrypted.Bytes()
}

// Opens the a base64-encoded file and decrypts it with the given key
func AESECBDecryptFile(filename string, key string) string {
	encoded := utils.ReadRelative(filename)
	decoded := utils.DecodeBase64(encoded)

	return string(AESECBDecrypt(decoded, []byte(key)))
}
