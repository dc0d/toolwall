package twerrors_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMark_first(t *testing.T) {
	var (
		assert        = assert.New(t)
		expectedError string
	)

	{
		expectedError = "[marker_call_fixture_test.go:8 twerrors_test.firstFn]\nROOT CAUSE ERROR"
	}

	actualError := firstFn()

	assert.Contains(actualError.Error(), expectedError)
}

func TestMark_another(t *testing.T) {
	var (
		assert        = assert.New(t)
		expectedError string
	)

	{
		expectedError = "[marker_call_fixture_test.go:12 twerrors_test.anotherFn]\nROOT CAUSE ERROR"
	}

	actualError := anotherFn()

	assert.Contains(actualError.Error(), expectedError)
}

func TestMark_third(t *testing.T) {
	var (
		assert        = assert.New(t)
		expectedError string
	)

	{
		expectedError = "[marker_call_fixture_test.go:20 twerrors_test.thirdFn]\n[marker_call_fixture_test.go:16 twerrors_test.secondFn]\n[marker_call_fixture_test.go:8 twerrors_test.firstFn]\nROOT CAUSE ERROR"
	}

	actualError := thirdFn()

	assert.Contains(actualError.Error(), expectedError)
}

func TestMark_anonymous_func(t *testing.T) {
	var (
		assert        = assert.New(t)
		expectedError string
	)

	{
		expectedError = "[marker_call_fixture_test.go:25 twerrors_test.callAnonymousFunc]\n[marker_call_fixture_test.go:24 twerrors_test.callAnonymousFunc.func1]\nROOT CAUSE ERROR"
	}

	actualError := callAnonymousFunc()

	assert.Contains(actualError.Error(), expectedError)
}

func TestMark_first_unwrap(t *testing.T) {
	var (
		assert        = assert.New(t)
		expectedError error
	)

	{
		expectedError = rootCause
	}

	actualError := errors.Unwrap(firstFn())

	assert.Equal(expectedError, actualError)
}

func TestMark_third_unwrap(t *testing.T) {
	var (
		assert        = assert.New(t)
		expectedError error
	)

	{
		expectedError = rootCause
	}

	actualError := errors.Unwrap(thirdFn())

	assert.Equal(expectedError, actualError)
}

func TestMark_nil_unwrap(t *testing.T) {
	originalRootCause := rootCause
	defer func() { rootCause = originalRootCause }()
	rootCause = nil

	var (
		assert        = assert.New(t)
		expectedError error
	)

	{
		expectedError = rootCause
	}

	actualError := errors.Unwrap(thirdFn())

	assert.Equal(expectedError, actualError)
}

func TestMark_json(t *testing.T) {
	var (
		assert       = assert.New(t)
		expectedJSON string
	)

	{
		expectedJSON = "{\"Calls\":[\"[marker_call_fixture_test.go:20 twerrors_test.thirdFn]\",\"[marker_call_fixture_test.go:16 twerrors_test.secondFn]\",\"[marker_call_fixture_test.go:8 twerrors_test.firstFn]\"],\"Cause\":\"ROOT CAUSE ERROR\"}"
	}

	err := thirdFn()

	js, err := json.Marshal(err)
	if err != nil {
		t.Fatal(err)
	}
	jsonStr := string(js)

	assert.Equal(expectedJSON, jsonStr)
}

func TestMark_json_nil(t *testing.T) {
	originalRootCause := rootCause
	defer func() { rootCause = originalRootCause }()
	rootCause = nil

	var (
		assert       = assert.New(t)
		expectedJSON string
	)

	{
		expectedJSON = "{\"Calls\":[\"[marker_call_fixture_test.go:20 twerrors_test.thirdFn]\",\"[marker_call_fixture_test.go:16 twerrors_test.secondFn]\",\"[marker_call_fixture_test.go:8 twerrors_test.firstFn]\"],\"Cause\":null}"
	}

	err := thirdFn()

	js, err := json.Marshal(err)
	if err != nil {
		t.Fatal(err)
	}
	jsonStr := string(js)

	assert.Equal(expectedJSON, jsonStr)
}

func TestMark_json_sentinel(t *testing.T) {
	originalRootCause := rootCause
	defer func() { rootCause = originalRootCause }()
	rootCause = sentinelErr("ROOT CAUSE ERROR")

	var (
		assert       = assert.New(t)
		expectedJSON string
	)

	{
		expectedJSON = "{\"Calls\":[\"[marker_call_fixture_test.go:20 twerrors_test.thirdFn]\",\"[marker_call_fixture_test.go:16 twerrors_test.secondFn]\",\"[marker_call_fixture_test.go:8 twerrors_test.firstFn]\"],\"Cause\":\"ROOT CAUSE ERROR\"}"
	}

	err := thirdFn()

	js, err := json.Marshal(err)
	if err != nil {
		t.Fatal(err)
	}
	jsonStr := string(js)

	assert.Equal(expectedJSON, jsonStr)
}

var (
	rootCause = errors.New("ROOT CAUSE ERROR")
)

type sentinelErr string

func (s sentinelErr) Error() string { return string(s) }

// ?: recursive
