/*
Copyright (c) 2016 Palantir Technologies

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"os"
	"os/exec"
	"testing"
)

const expected = `{
    "diff": {
        "": {
            "destroy": false
        },
        "nestedModule": {
            "destroy": false
        },
        "nestedModule.null_resource.nested": {
            "destroy": true,
            "destroy_tainted": false,
            "id": "",
            "triggers.%": "3",
            "triggers.nested_bar": "barCHANGED",
            "triggers.nested_baz": "baz",
            "triggers.nested_foo": "foo"
        },
        "null_resource.foo": {
            "destroy": true,
            "destroy_tainted": false,
            "id": "",
            "triggers.%": "3",
            "triggers.bar": "bar",
            "triggers.baz": "",
            "triggers.blarg": "blarg",
            "triggers.foo": "fooCHANGED"
        }
    },
    "state": {
        "nestedModule.null_resource.nested": {
            "id": "1574642143301947708",
            "triggers.%": "3",
            "triggers.nested_bar": "bar",
            "triggers.nested_baz": "baz",
            "triggers.nested_foo": "foo"
        },
        "null_resource.foo": {
            "id": "1678235666534884518",
            "triggers.%": "3",
            "triggers.bar": "bar",
            "triggers.baz": "baz",
            "triggers.foo": "foo"
        }
    }
}`

func Test(t *testing.T) {
	os.Chdir("example/")
	mustRun(t, "terraform", "init")
	mustRun(t, "terraform", "plan", "-out=terraform.tfplan")

	j, err := tfjson_string("terraform.tfplan")
	if err != nil {
		t.Fatal(err)
	}

	if j != expected {
		t.Errorf("Expected: %s\nActual: %s", expected, j)
	}
}

func mustRun(t *testing.T, name string, arg ...string) {
	if _, err := exec.Command(name, arg...).Output(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			t.Fatal(string(exitError.Stderr))
		} else {
			t.Fatal(err)
		}
	}
}
