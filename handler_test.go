package lab2

import (
	"bytes"
	"fmt"
	"strings"

	. "gopkg.in/check.v1"
)

func (s *MySuite) TestInputFile(c *C) {
	var buff bytes.Buffer

	handler := &ComputeHandler{
		Input:  strings.NewReader("12 2 3 4 * 10 5 / + * +"),
		Output: &buff,
	}
	handler.Compute()
	c.Assert(buff.String(), Equals, "40")
}

func (s *MySuite) TestInputErr(c *C) {
	var buff bytes.Buffer

	handler := &ComputeHandler{
		Input:  strings.NewReader("c*322"),
		Output: &buff,
	}
	err := handler.Compute()
	c.Assert(fmt.Sprint(err), Equals, "Expression contains invalid characters")
}