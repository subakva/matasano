package problem1

import "testing"

func TestHexToBase64(t *testing.T) {
  hex     := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
  expect  := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
  actual  := HexToBase64(hex)
  if actual != expect {
    t.Errorf("HexToBase64(%v)", hex)
    t.Errorf("Expected: %v", expect)
    t.Errorf("Actual  : %v", actual)
  }
}
