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
