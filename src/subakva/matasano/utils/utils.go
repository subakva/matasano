/*
 * Kitchen sink module with generic functions that are shared across problems.
 */

package utils

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"os"
	"strings"
)

// Reads a file relative to the current working directory.
func ReadRelative(filename string) []byte {
	wd, _ := os.Getwd()
	path := wd + "/" + filename
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return data
}

// Reads a file line-by-line and passes it to the calback function. If the return
// value is not an empty string, stop scanning and return the value.
func ReadAndScan(filename string, cb func(*bufio.Scanner) string) string {
	wd, _ := os.Getwd()
	path := wd + "/" + filename

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result := cb(scanner)
		if result != "" {
			return result
		}
	}
	return ""
}

// Returns true if the int is in the array of ints
func IntInArray(haystack []int, needle int) bool {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}

// Decodes a b64-encoded byte array
func DecodeBase64(encoded []byte) []byte {
	decoded, err := base64.StdEncoding.DecodeString(string(encoded))
	if err != nil {
		panic(err)
	}
	return decoded
}

// Decodes a hex-encoded byte array
func DecodeHex(encoded []byte) []byte {
	trimmed := strings.TrimSpace(string(encoded))
	decoded, err := hex.DecodeString(trimmed)
	if err != nil {
		panic(err)
	}
	return decoded
}

// Converts an array of strings into an array of byte arrays
func StringsToBytes(strings []string) [][]byte {
	bytes := make([][]byte, len(strings))
	for i := 0; i < len(strings); i++ {
		bytes[i] = []byte(strings[i])
	}
	return bytes
}

// Converts an array of byte arrays into an array of strings
func BytesToStrings(bytes [][]byte) []string {
	strings := make([]string, len(bytes))
	for i := 0; i < len(bytes); i++ {
		strings[i] = string(bytes[i])
	}
	return strings
}

// Splits the byte array into an array of byte arrays of the specified size.
func ChunkBytes(chunkMe []byte, chunkSize int) [][]byte {
	numChunks := len(chunkMe) / chunkSize
	if len(chunkMe)%chunkSize != 0 {
		numChunks += 1
	}
	chunks := make([][]byte, numChunks)
	for i := 0; i < numChunks; i++ {
		startIndex := chunkSize * i
		endIndex := chunkSize * (i + 1)
		if endIndex > len(chunkMe) {
			endIndex = len(chunkMe)
		}
		chunks[i] = chunkMe[startIndex:endIndex]
	}
	return chunks
}
