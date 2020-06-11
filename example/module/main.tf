resource "null_resource" "nested" {
  triggers = {
    nested_foo = "foo"
    nested_bar = "barCHANGED"
    nested_baz = "baz"
  }
}
