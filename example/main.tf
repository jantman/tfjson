resource "null_resource" "foo" {
  triggers = {
    foo = "fooCHANGED"
    bar = "bar"
    blarg = "blarg"
  }
}

module "nestedModule" {
  source = "module/"
}
