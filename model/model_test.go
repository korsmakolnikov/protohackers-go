package model

import (
	"testing"
)

func testUnmarshalRequest(t *testing.T) {
	r, err := unmarshalRequest("{\"method\":\"isPrime\", \"number\":\"5.5\"}")
	if err != nil {
		t.Errorf("Request unmarshal fail with error %q", err)
	}

	if rFloat, ok := r.(Request[FlexFloat]); ok {
		if !rFloat.isValid() {
			t.Errorf("Request isn't valid on %#v\n", rFloat)
		}
		if rFloat.Number != 5.5 {
			t.Errorf("Request hasn't been unmarshaled properly %#v\n", rFloat)
		}
	}

	r, err = unmarshalRequest("{\"method\":\"isPrime\", \"number\":\"5\"}")
	if err != nil {
		t.Errorf("Request unmarshal fail with error %q", err)
	}

	if rInt, ok := r.(Request[FlexInt]); ok {
		if !rInt.isValid() {
			t.Errorf("Request isn't valid on %#v\n", rInt)
		}
		if rInt.Number != 5 {
			t.Errorf("Request hasn't been unmarshaled properly %#v\n", rInt)
		}
	}

}

func TestUnmarshalAndValidateStringRequests(t *testing.T) {
	r, err := unmarshalRequestFloat("{\"method\":\"isPrime\", \"number\":\"5.5\"}")
	if err != nil {
		t.Errorf("Request unmarshal fail with error %q", err)
	}

	if !r.isValid() {
		t.Errorf("Request isn't valid on %#v\n", r)
	}

	if r.Number != 5.5 {
		t.Errorf("Request hasn't been unmarshaled properly %#v\n", r)
	}
}

func TestUnmarshalAndValidateFloatRequests(t *testing.T) {
	r, err := unmarshalRequestFloat("{\"method\":\"isPrime\", \"number\":5.5}")
	if err != nil {
		t.Errorf("Request unmarshal fail with error %q", err)
	}

	if !r.isValid() {
		t.Errorf("Request isn't valid on %#v\n", r)
	}

	if r.Number != 5.5 {
		t.Errorf("Request hasn't been unmarshaled properly %#v\n", r)
	}
}

func TestUnmarshalAndValidateIntegerRequests(t *testing.T) {
	r, err := unmarshalRequestInt("{\"method\":\"isPrime\", \"number\":3}")
	if err != nil {
		t.Errorf("Request unmarshal fail with error %q", err)
	}

	if !r.isValid() {
		t.Errorf("Request isn't valid on %q", r)
	}
	if r.Number != 3 {
		t.Errorf("Request hasn't been unmarshaled properly %#v\n", r)
	}
}
