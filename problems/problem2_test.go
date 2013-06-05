package problems

import (
  . "launchpad.net/gocheck"
)

type Problem2Suite struct{}
var _ = Suite(&Problem2Suite{})

func (s *Problem2Suite) TestFixedXOR(c *C) {
  hex1 := "1c0111001f010100061a024b53535009181c"
  hex2 := "686974207468652062756c6c277320657965"
  c.Check(FixedXOR(hex1, hex2), Equals, "746865206b696420646f6e277420706c6179")
}
