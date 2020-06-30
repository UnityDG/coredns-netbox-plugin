// Copyright 2020 Oz Tiram <oz.tiram@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netbox

import (
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestQuery(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution
	gock.New("https://example.org/api/ipam/ip-addresses/").MatchParams(
		map[string]string{"dns_name": "my_host"}).Reply(
		200).BodyString(
		`{"count":1, "results":[{"address": "10.0.0.2/25", "dns_name": "my_host"}]}`)

	want := "10.0.0.2"
	got := query("https://example.org/api/ipam/ip-addresses", "mytoken", "my_host")
	if got != want {
		t.Fatalf("Expected %s but got %s", want, got)
	}

}

func TestNoSuchHost(t *testing.T) {

	defer gock.Off() // Flush pending mocks after test execution
	gock.New("https://example.org/api/ipam/ip-addresses/").MatchParams(
		map[string]string{"dns_name": "NoSuchHost"}).Reply(
		200).BodyString(`{"count":0,"next":null,"previous":null,"results":[]}`)

	want := ""
	got := query("https://example.org/api/ipam/ip-addresses", "mytoken", "NoSuchHost")
	if got != want {
		t.Fatalf("Expected empty string but got %s", got)
	}

}