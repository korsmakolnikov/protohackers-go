package model

import (
	"testing"
)

func TestUnmarshalAndValidateStringRequests(t *testing.T) {

	var request Request[FlexFloat]
	r, err := request.unmarshalRequest("{\"method\":\"isPrime\", \"number\":\"5.5\"}")
	if err != nil {
		t.Errorf("Request unmarshal fail with error %q", err)
	}

	if !request.isValid() {
		t.Errorf("Request isn't valid on %#v\n", request)
	}

	if r.Number != 5.5 {
		t.Errorf("Request hasn't been unmarshaled properly %#v\n", request)
	}
}

func TestUnmarshalAndValidateFloatRequests(t *testing.T) {

	var request Request[FlexFloat]
	r, err := request.unmarshalRequest("{\"method\":\"isPrime\", \"number\":5.5}")
	if err != nil {
		t.Errorf("Request unmarshal fail with error %q", err)
	}

	if !request.isValid() {
		t.Errorf("Request isn't valid on %#v\n", request)
	}

	if r.Number != 5.5 {
		t.Errorf("Request hasn't been unmarshaled properly %#v\n", request)
	}
}

func TestUnmarshalAndValidateIntegerRequests(t *testing.T) {

	var request Request[FlexInt]
	r, err := request.unmarshalRequest("{\"method\":\"isPrime\", \"number\":3}")
	if err != nil {
		t.Errorf("Request unmarshal fail with error %q", err)
	}

	if !request.isValid() {
		t.Errorf("Request isn't valid on %q", request)
	}
	if r.Number != 3 {
		t.Errorf("Request hasn't been unmarshaled properly %#v\n", request)
	}
}
