package iptools

import (
	"net"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type IsPrivateSuite struct{}

var _ = Suite(&IsPrivateSuite{})

func (s *IsPrivateSuite) TestV4(c *C) {
	c.Assert(IsPrivate(net.ParseIP("1.1.1.1")), Equals, false)

	c.Assert(IsPrivate(net.ParseIP("10.0.0.1")), Equals, true)

	c.Assert(IsPrivate(net.ParseIP("169.254.0.1")), Equals, true)
}

func (s *IsPrivateSuite) TestV6(c *C) {
	c.Assert(IsPrivate(net.ParseIP("2607:f8b0:4000:802::200e")), Equals, false)

	c.Assert(IsPrivate(net.ParseIP("fd12:3456:789a:1::1")), Equals, true)

	c.Assert(IsPrivate(net.ParseIP("fe80::5054:8fff:fe9f:af62")), Equals, true)
}
