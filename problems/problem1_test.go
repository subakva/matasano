package problems

import (
  . "launchpad.net/gocheck"
)

type Problem1Suite struct{}
var _ = Suite(&Problem1Suite{})

func (s *Problem1Suite) TestHexToBase64(c *C) {
  hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
  c.Check(HexToBase64(hex), Equals, "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t")
}
