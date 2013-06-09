package problem2

import (
  "testing"
)

func TestFixedXOR(t *testing.T) {
  hex1   := "1c0111001f010100061a024b53535009181c"
  hex2   := "686974207468652062756c6c277320657965"
  expect := "746865206b696420646f6e277420706c6179"
  actual := FixedXOR(hex1, hex2)
  if actual != expect {
    t.Errorf("FixedXOR(%v, %v)", hex1, hex2)
    t.Errorf("Expected: %v", expect)
    t.Errorf("Actual  : %v", actual)
  }

}

