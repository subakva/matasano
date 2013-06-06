package problems

import (
  . "launchpad.net/gocheck"
)

type Problem3Suite struct{}
var _ = Suite(&Problem3Suite{})

func (s *Problem3Suite) TestFixedXOR(c *C) {
  message := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
  decoded, key := DecipherSingleCharacterXOR(message)
  c.Check(decoded, Equals, "Cooking MC's like a pound of bacon")
  c.Check(key, Equals, "X")
}
