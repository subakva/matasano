package problem2

import (
	"subakva/matasano/assertions"
	"testing"
)

func TestFixedXOR(t *testing.T) {
	hex1 := "1c0111001f010100061a024b53535009181c"
	hex2 := "686974207468652062756c6c277320657965"
	expect := "746865206b696420646f6e277420706c6179"
	assertions.AssertEquals(t, FixedXOR(hex1, hex2), expect, "FixedXOR did not match expected.")
}
