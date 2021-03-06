// Copyright 2016 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package testserver_test

import (
	"strings"
	"testing"

	"github.com/cockroachdb/cockroach-go/testserver"
)

func TestRunServer(t *testing.T) {
	db, stop := testserver.NewDBForTest(t)
	defer stop()

	if _, err := db.Exec("SELECT 1"); err != nil {
		t.Fatal(err)
	}
}

func TestPGURLWhitespace(t *testing.T) {
	ts, err := testserver.NewTestServer()
	if err != nil {
		t.Fatal(err)
	}
	if err := ts.Start(); err != nil {
		t.Fatal(err)
	}
	url := ts.PGURL().String()
	if trimmed := strings.TrimSpace(url); url != trimmed {
		t.Errorf("unexpected whitespace in server URL: %q", url)
	}
}
