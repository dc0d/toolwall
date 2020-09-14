package twerrors_test

import (
	"github.com/dc0d/toolwall/twerrors"
)

func firstFn() error {
	return twerrors.Mark(rootCause)
}

func anotherFn() error {
	return twerrors.Mark(rootCause)
}

func secondFn() error {
	return twerrors.Mark(firstFn())
}

func thirdFn() error {
	return twerrors.Mark(secondFn())
}

func callAnonymousFunc() error {
	fn := func() error { return twerrors.Mark(rootCause) }
	return twerrors.Mark(fn())
}
