package model

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

type FlexInt int64

type FlexIntError struct{}

func (r *FlexIntError) Error() string {
	return "FailedParsingError"
}

func (n *FlexInt) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		return unmarshalNumber(b, n)
	} else {
		return unmarshalString(b, n)
	}
}

func (n *FlexInt) isPrime() bool {
	in := int64(*n)
	bn := big.NewInt(in)
	return bn.ProbablyPrime(0)
}

func unmarshalNumber(b []byte, n *FlexInt) error {
	err := json.Unmarshal(b, (*int64)(n))
	if err != nil {
		var f float64
		err = json.Unmarshal(b, &f)
		if err != nil {
			return &FlexIntError{}
		}

		*n = FlexInt(math.Round(f))
	}

	return err

}

func unmarshalString(b []byte, n *FlexInt) error {

	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return &FlexIntError{}
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("%+v\n", err)
		x, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Printf("%+v\n", err)
			return &FlexIntError{}
		} else {
			i = int(math.Round(x))
		}
	}

	*n = FlexInt(i)
	return nil

}

/*
 * a request is a json string ended by newline character
 * a valid request must be valid json
 * must contain a field "method" (and it have to be set to "isPrime")
 * must contain a field "number" that contains any number
 */
type Request struct {
	Method        string  `json:"method"`
	Number        FlexInt `json:"number"`
	FailedParsing bool
}

func (r *Request) isValid() bool {
	if r.Method != "isPrime" {
		return false
	}

	if r.FailedParsing {
		return false
	}

	return true
}

func (r *Request) isPrime() bool {
	return r.Number.isPrime()
}

func unmarshalRequest(requestString string) (r Request, err error) {
	err = json.Unmarshal([]byte(requestString), &r)
	if err != nil {
		r.FailedParsing = true
	}

	return
}

type Response struct {
}
