// Copyright 2020 Nimiq community.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nimiqrpc

import (
	"testing"
)

func TestLunaToNIM(t *testing.T) {
	// 0 luna is 0 NIM
	if FormatNIM(0) != "0" {
		t.Fail()
	}
	// 1 luna is 0.00001 NIM
	if FormatNIM(1) != "0.00001" {
		t.Fail()
	}
	// 100000 luna is 1 NIM
	if FormatNIM(100000) != "1" {
		t.Fail()
	}
	// 1234567 is 12.34567 NIM
	if FormatNIM(1234567) != "12.34567" {
		t.Fail()
	}
	// 1200000 is 12 NIM
	if FormatNIM(1200000) != "12" {
		t.Fail()
	}
	// 123456789 is 1234.56789 NIM
	if FormatNIM(123456789) != "1234.56789" {
		t.Fail()
	}
}

func TestNIMToLuna(t *testing.T) {
	// 0 NIM is 0 Luna
	if l, _ := FormatLuna("0"); l != 0 {
		t.Fail()
	}
	// 0.00001 NIM is 1 Luna
	if l, _ := FormatLuna("0.00001"); l != 1 {
		t.Fail()
	}
	// 1 NIM is 100000 Luna
	if l, _ := FormatLuna("1"); l != 100000 {
		t.Fail()
	}
	// 12.34567 NIM is 1234567 Luna
	if l, _ := FormatLuna("12.34567"); l != 1234567 {
		t.Fail()
	}
	// 12 NIM is 1200000 Luna
	if l, _ := FormatLuna("12"); l != 1200000 {
		t.Fail()
	}
	// 1234.56789 NIM is 123456789 Luna
	if l, _ := FormatLuna("1234.56789"); l != 123456789 {
		t.Fail()
	}
}

// The Nimiq Network has been designed for a total supply of 21 Billion NIM.
// The smallest unit of NIM is called Luna and 100’000 (1e5) Luna equal 1 NIM,
// which results in a total supply of 21e14 Luna
func TestMaxLuna(t *testing.T) {
	max, _ := FormatLuna("21000000000")
	if max != 2100000000000000 {
		t.Fail()
	}
}
