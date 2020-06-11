tfjson
======

Utility to read in a Terraform plan file and dump it out in JSON. Standalone version of [Terraform PR #3170](https://github.com/hashicorp/terraform/pull/3170).

**Note:** This fork has been modified from upstream:

* It produces flat JSON, with ``modulename.type.resourcename`` keys instead of nested module mappings.
* It has two top-level keys, ``diff`` and ``state``, so that you can produce an actual diff with old and new values.

## Installation

```
$ git clone https://github.com/jantman/tfjson.git
$ cd tfjson
$ unset GOPATH
$ go build
```

Either run ``./tfjson``, or copy it somewhere useful (e.g. ``cp tfjson ~/bin/``).

## Usage

Given the terraform configuration in [example/](example/), running `terraform plan -out=terraform.tfplan` produces a Terraform plan file.

The JSON representation produced by `tfjson` looks like:

```json
{
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
}
```

## Compatibility Notice

This library is compatible up to Terraform v0.11.

## tagdiff.py

This is a Python script (I needed it quick, and my Python is much better than my Go) that shows diff output much like ``terraform plan``, but shows a logical key-based diff for tags on ASGs and other resources that display tags indexed numerically instead of by tag key. It operates on the output of ``tfjson``.

## License

This project is made available under the [MIT License](http://opensource.org/licenses/MIT).
