package valdo_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/valdo/valdo"
)

func FuzzValidateDuplicateNames(f *testing.F) {
	for _, input := range []string{
		`null`, `42`, `1e999`, `[]`, `{"a":1} {}`,
		`{"a":1,"a":2}`, `{"a":null,"\u0061":2}`,
		`{"a":{"b":1,"c":2},"a":{}}`,
		`[{"a":1},{"a":2}]`, `{"a":1,"child":{"a":2}}`,
		`{"a":"\"a\":1,\"a\":2","b":"{[\\\"}]"}`,
		`{"😀":0,"\ud83d\ude00":1}`,
		`{"\ud800":0,"\ufffd":1}`,
		"{\"\xff\":1,\"\ufffd\":2}",
	} {
		f.Add([]byte(input))
	}
	f.Fuzz(func(t *testing.T, input []byte) {
		is := is.New(t)
		err := valdo.Validate(valdo.Any(), input)
		var decoded any
		if json.Unmarshal(input, &decoded) != nil {
			is.True(err != nil)
			return
		}
		// Independent oracle: decoded names from the standard token decoder,
		// with a separate set for each object. Do not use member counts here.
		d := json.NewDecoder(bytes.NewReader(input))
		duplicate, refErr := referenceDuplicateNames(d)
		is.NoErr(refErr)
		is.Equal(err != nil, duplicate)
	})
}

func referenceDuplicateNames(d *json.Decoder) (bool, error) {
	token, err := d.Token()
	if err != nil {
		return false, err
	}
	duplicate := false
	switch token {
	case json.Delim('{'):
		seen := map[string]bool{}
		for d.More() {
			key, err := d.Token()
			if err != nil {
				return false, err
			}
			name := key.(string)
			duplicate = duplicate || seen[name]
			seen[name] = true
			sub, err := referenceDuplicateNames(d)
			if err != nil {
				return false, err
			}
			duplicate = duplicate || sub
		}
		_, err = d.Token()
	case json.Delim('['):
		for d.More() {
			sub, err := referenceDuplicateNames(d)
			if err != nil {
				return false, err
			}
			duplicate = duplicate || sub
		}
		_, err = d.Token()
	}
	return duplicate, err
}
