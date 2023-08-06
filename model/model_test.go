package model

import (
	"fmt"
	"testing"
)

func TestIsPrimeNumber(t *testing.T) {
	r, err := unmarshalRequest("{\"method\":\"isPrime\", \"number\":\"11\"}")
	if err != nil {
		t.Errorf("Request unmarshal fail with error %q", err)
	}

	if !r.isValid() {
		t.Errorf("Request isn't valid on %#v\n", r)
	}

	if !r.isPrime() {
		t.Errorf("Request hasn't been unmarshaled properly %#v\n", r)
	}
}

func TestIsNotPrimeNumber(t *testing.T) {
	r, err := unmarshalRequest("{\"method\":\"isPrime\", \"number\":\"5.7\"}")
	if err != nil {
		t.Errorf("Request unmarshal fail with error %q", err)
	}

	if !r.isValid() {
		t.Errorf("Request isn't valid on %#v\n", r)
	}

	if r.isPrime() {
		t.Errorf("Request hasn't been unmarshaled properly %#v\n", r)
	}
}

func TestUnmarshalRequestString(t *testing.T) {
	r, err := unmarshalRequest("{\"method\":\"isPrime\", \"number\":\"5.7\"}")
	if err != nil {
		t.Errorf("Request unmarshal fail with error %q", err)
	}

	if !r.isValid() {
		t.Errorf("Request isn't valid on %#v\n", r)
	}

	if r.Number != 6 {
		t.Errorf("Request hasn't been unmarshaled properly %#v\n", r)
	}
}

func TestUnmarshalRequestFloat(t *testing.T) {
	r, err := unmarshalRequest("{\"method\":\"isPrime\", \"number\":5.5}")
	if err != nil {
		t.Errorf("Request unmarshal fail with error %q", err)
	}

	if !r.isValid() {
		t.Errorf("Request isn't valid on %#v\n", r)
	}

	if r.Number != 6 {
		t.Errorf("Request hasn't been unmarshaled properly %#v\n", r)
	}
}

func TestUnmarshalRequestInt(t *testing.T) {
	r, err := unmarshalRequest("{\"method\":\"isPrime\", \"number\":5}")
	if err != nil {
		t.Errorf("Request unmarshal fail with error %q", err)
	}

	if !r.isValid() {
		t.Errorf("Request isn't valid on %#v\n", r)
	}

	if r.Number != 5 {
		t.Errorf("Request hasn't been unmarshaled properly %#v\n", r)
	}
}

func TestUnmarshalRequestInvalidMethod(t *testing.T) {
	r, err := unmarshalRequest("{\"method\":\"invalid\", \"number\":\"5\"}")
	if err != nil {
		t.Errorf("Request unmarshal fail with error %q", err)
	}

	if r.isValid() {
		t.Errorf("Request isn't be valid %#v\n", r)
	}
}

func TestUnmarshalRequestUnparsable(t *testing.T) {
	r, err := unmarshalRequest("{\"method\":\"isPrime\", \"number\":\"unparsable\"}")

	fmt.Printf("%+v\n", err)
	if err == nil {
		t.Errorf("Request unmarshal should fail")
	}

	if r.isValid() {
		t.Errorf("Request shouldn't be valid %#v\n", r)
	}
}
