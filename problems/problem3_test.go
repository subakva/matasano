package problems

import (
  . "launchpad.net/gocheck"
)

type Problem3Suite struct{}
var _ = Suite(&Problem3Suite{})

func (s *Problem3Suite) TestFixedXOR(c *C) {
  message := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
  c.Check(DecipherSingleCharacterXOR(message), Equals, "Vanilla Ice")
}
