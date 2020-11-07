package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestEvaluatePostfixSimple1(c *C) {
	res, _ := EvaluatePostfix("5 4 - 3 *")
	c.Assert(res, Equals, "3")
}

func (s *MySuite) TestEvaluatePostfixSimple2(c *C) {
	res, _ := EvaluatePostfix("8 6 - 5 ^")
	c.Assert(res, Equals, "32")
}

func (s *MySuite) TestEvaluatePostfixSimple3(c *C) {
	res, _ := EvaluatePostfix("1 4 + 8 -")
	c.Assert(res, Equals, "-3")
}

func (s *MySuite) TestEvaluatePostfixSimple4(c *C) {
	res, _ := EvaluatePostfix("5 4 - 3 *")
	c.Assert(res, Equals, "3")
}

func (s *MySuite) TestEvaluatePostfixSimple5(c *C) {
	res, _ := EvaluatePostfix("3 4 +")
	c.Assert(res, Equals, "7")
}

func (s *MySuite) TestEvaluatePostfixAdvanced1(c *C) {
	res, _ := EvaluatePostfix("12 2 3 4 * 10 5 / + * +")
	c.Assert(res, Equals, "40")
}

func (s *MySuite) TestEvaluatePostfixAdvanced2(c *C) {
	res, _ := EvaluatePostfix("33 22 + 27 - 3 16 10 - 2 ^ * +")
	c.Assert(res, Equals, "136")
}

func (s *MySuite) TestEvaluatePostfixAdvanced3(c *C) {
	res, _ := EvaluatePostfix("3 3 ^ 27 - 4 5 + 3 / + 17 +")
	c.Assert(res, Equals, "20")
}

func (s *MySuite) TestEvaluatePostfixAdvanced4(c *C) {
	res, _ := EvaluatePostfix("4 2 ^ 2 3 ^ - 4 / 5 3 * +")
	c.Assert(res, Equals, "17")
}

func (s *MySuite) TestEvaluatePostfixAdvanced5(c *C) {
	res, _ := EvaluatePostfix("20 10 / 2 ^ 13 2 * - 15 3 / +")
	c.Assert(res, Equals, "-17")
}

func (s *MySuite) TestEvaluatePostfixEmptyString(c *C) {
	res, err := EvaluatePostfix("")
	c.Assert(res, Equals, "")
	c.Assert(fmt.Sprint(err), Equals, "Expression is empty string")
}

func (s *MySuite) TestEvaluatePostfixInvalidExpression(c *C) {
	res, err := EvaluatePostfix("c*322")
	c.Assert(res, Equals, "")
	c.Assert(fmt.Sprint(err), Equals, "Expression contains invalid characters")
}

func (s *MySuite) TestEvaluatePostfixSpacesOnly(c *C) {
	res, err := EvaluatePostfix("  ")
	c.Assert(res, Equals, "")
	c.Assert(fmt.Sprint(err), Equals, "Expression is empty string")
}
