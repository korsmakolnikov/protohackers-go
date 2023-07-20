package model

import (
	"encoding/json"
	"strconv"
)

type Number interface {
	FlexFloat | FlexInt
}

type FlexInt int64

func (n *FlexInt) UnmarshalJSON(b []byte) error {
	// Integer
	if b[0] != '"' {
		return json.Unmarshal(b, (*int64)(n))
	}

	// String
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	*n = FlexInt(i)

	return nil
}

type FlexFloat float64

func (n *FlexFloat) UnmarshalJSON(b []byte) error {
	// Integer
	if b[0] != '"' {
		return json.Unmarshal(b, (*float64)(n))
	}

	// String
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	*n = FlexFloat(i)

	return nil
}

/*
 * a request is a json string ended by newline character
 * a valid request must be valid json
 * must contain a field "method" (and it have to be set to "isPrime")
 * must contain a field "number" that contains any number
 */
type Request[T Number] struct {
	Method string `json:"method"`
	Number T      `json:"number"`
}

func (r *Request[T]) isValid() bool {
	if r.Method != "isPrime" {
		return false
	}

	return true
}

func (r *Request[T]) unmarshalRequest(requestString string) (*Request[T], error) {
	err := json.Unmarshal([]byte(requestString), r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

type Response struct {
}
