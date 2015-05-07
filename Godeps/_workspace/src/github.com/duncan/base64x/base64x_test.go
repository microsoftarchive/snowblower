/*
Copyright 2015 James Duncan Davidson

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package base64x

import "testing"

func TestAutoconvertIsFalseByDefault(t *testing.T) {
	e := Encoding{}
	if e.convert {
		t.Errorf("Expected autoconvert to be false!")
	}
}

func TestConvertEncoding(t *testing.T) {
	s := convertEncoding("ABCDEF+12345/DKDLS")
	if s != "ABCDEF-12345_DKDLS" {
		t.Errorf("Encoding conversation fail. Got %s", s)
	}
}

func TestStdDecodeAcceptsNormallyPaddedData(t *testing.T) {
	s := "YWJjZGU="
	d := make([]byte, len(s))
	i, err := StdEncoding.Decode(d, []byte(s))
	if string(d[0:i]) != "abcde" {
		t.Errorf("Expected abcde, got %s", string(d))
	}
	if i > len(s) {
		t.Errorf("Expected %v bytes decoded, was %v", len(s), i)
	}
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}

func TestStdDecodeAcceptsUnpaddedData(t *testing.T) {
	s := "YWJjZGU"
	d := make([]byte, len(s))
	i, err := StdEncoding.Decode(d, []byte(s))
	if string(d[0:i]) != "abcde" {
		t.Errorf("Expected abcde, got %s", string(d))
	}
	if i > len(s) {
		t.Errorf("Expected %v bytes decoded, was %v", len(s), i)
	}
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}

func TestURLDecodeAcceptsNormallyPaddedData(t *testing.T) {
	s := "YWJjZGU="
	d := make([]byte, len(s))
	i, err := URLEncoding.Decode(d, []byte(s))
	if string(d[0:i]) != "abcde" {
		t.Errorf("Expected abcde, got %s", string(d))
	}
	if i > len(s) {
		t.Errorf("Expected %v bytes decoded, was %v", len(s), i)
	}
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}

func TestURLDecodeAcceptsUnpaddedData(t *testing.T) {
	s := "YWJjZGU"
	d := make([]byte, len(s))
	i, err := URLEncoding.Decode(d, []byte(s))
	if string(d[0:i]) != "abcde" {
		t.Errorf("Expected abcde, got %s", string(d))
	}
	if i > len(s) {
		t.Errorf("Expected %v bytes decoded, was %v", len(s), i)
	}
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}

func TestAutoDecodeAcceptsNormallyPaddedData(t *testing.T) {
	s := "YWJjZGU="
	d := make([]byte, len(s))
	i, err := AutoEncoding.Decode(d, []byte(s))
	if string(d[0:i]) != "abcde" {
		t.Errorf("Expected abcde, got %s", string(d))
	}
	if i > len(s) {
		t.Errorf("Expected %v bytes decoded, was %v", len(s), i)
	}
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}

func TestAutoDecodeAcceptsUnpaddedData(t *testing.T) {
	s := "YWJjZGU"
	d := make([]byte, len(s))
	i, err := AutoEncoding.Decode(d, []byte(s))
	if string(d[0:i]) != "abcde" {
		t.Errorf("Expected abcde, got %s", string(d))
	}
	if i > len(s) {
		t.Errorf("Expected %v bytes decoded, was %v", len(s), i)
	}
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}

func TestStdDecodeAcceptsNormallyPaddedStrings(t *testing.T) {
	s, err := StdEncoding.DecodeString("YWJjZGU=")
	if string(s) != "abcde" {
		t.Errorf("Expected abcde, got %s", string(s))
	}
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}

func TestStdDecodeStringAcceptsUnpaddedStrings(t *testing.T) {
	s, err := StdEncoding.DecodeString("YWJjZGU")
	if string(s) != "abcde" {
		t.Errorf("Expected abcde, got %s", string(s))
	}
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}

func TestURLDecodeAcceptsNormallyPaddedStrings(t *testing.T) {
	s, err := URLEncoding.DecodeString("YWJjZGU=")
	if string(s) != "abcde" {
		t.Errorf("Expected abcde, got %s", string(s))
	}
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}

func TestURLDecodeStringAcceptsUnpaddedStrings(t *testing.T) {
	s, err := URLEncoding.DecodeString("YWJjZGU")
	if string(s) != "abcde" {
		t.Errorf("Expected abcde, got %s", string(s))
	}
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}

func TestAutoDecodeAcceptsNormallyPaddedStrings(t *testing.T) {
	s, err := AutoEncoding.DecodeString("YWJjZGU=")
	if string(s) != "abcde" {
		t.Errorf("Expected abcde, got %s", string(s))
	}
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}

func TestAutoDecodeStringAcceptsUnpaddedStrings(t *testing.T) {
	s, err := AutoEncoding.DecodeString("YWJjZGU")
	if string(s) != "abcde" {
		t.Errorf("Expected abcde, got %s", string(s))
	}
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}
