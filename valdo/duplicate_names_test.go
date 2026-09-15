package valdo_test

import (
	"strings"
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/jsony"
	"github.com/orsinium-labs/valdo/valdo"
)

func TestUnmarshalRejectsDuplicateObjects(t *testing.T) {
	is := is.New(t)
	type request struct {
		Profile struct {
			Admin bool `json:"admin"`
		} `json:"profile"`
	}
	v := valdo.Object(valdo.Property("profile", valdo.Object(
		valdo.Property("admin", valdo.BoolConst(false)).Optional(),
	)))
	input := []byte(`{"profile":{"admin":true},"profile":{}}`)
	got, err := valdo.Unmarshal[request](v, input)
	t.Logf("error=%v; admin=%v", err, got.Profile.Admin)
	is.True(err != nil) // Repeated objects must not bypass validation of admin.
	is.Equal(got, request{})
}

func TestValidateRejectsDuplicateNames(t *testing.T) {
	cases := []struct{ name, input string }{
		{"root", `{"a":1,"a":2}`},
		{"equal_values", `{"a":1,"a":1}`},
		{"first_null", `{"a":null,"a":2}`},
		{"empty_name", `{"":1,"":2}`},
		{"nested", `{"outer":{"a":1,"a":2}}`},
		{"array", `[{"a":1,"a":2}]`},
		{"parent_after_child", `{"a":1,"child":{"a":2},"a":3}`},
		{"escaped_ascii", `{"a":1,"\u0061":2}`},
		{"escaped_quote", `{"a\"b":1,"a\u0022b":2}`},
		{"escaped_unicode", `{"é":1,"\u00e9":2}`},
		{"surrogate_pair", `{"😀":1,"\ud83d\ude00":2}`},
		{"unpaired_surrogate", `{"\ud800":1,"\ufffd":2}`},
		{"invalid_utf8", "{\"\xff\":1,\"\ufffd\":2}"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			is := is.New(t)
			v := &recordingValidator{}
			err := valdo.Validate(v, []byte(tc.input))
			is.True(err != nil)
			is.Equal(v.calls, 0) // Reject before invoking even a custom validator.
		})
	}
}

func TestValidateAcceptsNamesInSeparateObjects(t *testing.T) {
	for _, input := range []string{
		`{"a":1,"A":2}`,
		`{"a":1,"child":{"a":2}}`,
		`[{"a":1},{"a":2}]`,
		`{"left":{"a":1},"right":{"a":2}}`,
		`{"é":1,"e\u0301":2}`,
		`{"a":"\"a\":1,\"a\":2","b":"{[\\\"}]"}`,
		` [true,false,null,42,"a",{},[]] `,
		`{"escaped\nname" : 1,"different":2}`,
		strings.Repeat(`{"a":`, 1000) + `0` + strings.Repeat(`}`, 1000),
	} {
		is.New(t).NoErr(valdo.Validate(valdo.Any(), []byte(input)))
	}
}

func TestValidatePreservesJSONErrorsAndNumberTypes(t *testing.T) {
	is := is.New(t)
	for _, input := range []string{"", " ", `{`, `{"a":}`, `{"a":1,}`, `{"a":1} {}`, `{"a":"\q"}`, `1e999`, strings.Repeat("[", 10001) + "0" + strings.Repeat("]", 10001)} {
		v := &recordingValidator{}
		is.True(valdo.Validate(v, []byte(input)) != nil)
		is.Equal(v.calls, 0)
	}
	v := &recordingValidator{}
	is.NoErr(valdo.Validate(v, []byte(`42`)))
	is.Equal(v.value, any(float64(42))) // Existing custom validators still receive float64.
}

type recordingValidator struct {
	calls int
	value any
}

func (v *recordingValidator) Validate(value any) valdo.Error {
	v.calls++
	v.value = value
	return nil
}

func (*recordingValidator) Schema() jsony.Object { return nil }
