/*
Copyright 2023 Flant JSC

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

package cr

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []string{
		"registry.deckhouse.io/deckhouse/fe",
		"registry.deckhouse.io:5123/deckhouse/fe",
		"192.168.1.1/deckhouse/fe",
		"192.168.1.1:8080/deckhouse/fe",
		"2001:db8:3333:4444:5555:6666:7777:8888/deckhouse/fe",
		"[2001:db8::1]:8080/deckhouse/fe",
		"192.168.1.1:5123/deckhouse/fe",
	}

	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			u, err := parse(tt)
			if err != nil {
				t.Errorf("got error: %s", err)
			}
			if u.String() != "//"+tt {
				t.Errorf("got: %s, wanted: %s", u, tt)
			}
		})
	}
}
