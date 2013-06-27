package problem1

import "testing"
import "subakva/matasano/assertions"

func TestHexToBase64(t *testing.T) {
	hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	b64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	assertions.AssertEquals(t, HexToBase64(hex), b64, "HexToBase64 did not match expected.")
}

func TestBase64ToHex(t *testing.T) {
	b64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	assertions.AssertEquals(t, Base64ToHex(b64), hex, "Base64ToHex did not match expected.")
}
