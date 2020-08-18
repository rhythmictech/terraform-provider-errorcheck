# terraform-provider-errorcheck Provider
docs for `terraform-provider-errorcheck`

## Example Usage

code:
```terraform
locals {
  compare     = "success"
  testSuccess = "success"
  testFail    = "fail"
}

resource "errorcheck_is_valid" "shouldMatch" {
  name = "check_something"
  test = {
    assert = local.compare == local.testSuccess
    error_message = "Your assertion is not valid"
  }
}

resource "errorcheck_is_valid" "Not_valid_if_not_match" {
  name = "Should not match"
  test = {
    assert = local.compare == local.testFail
    error_message = "Your assertion is not valid"
  }
}

```

output:

```bash
terraform validate .

Error: Your assertion is not valid

  on main.tf line 11, in resource "errorcheck_is_valid" "Not_valid_if_not_match":
  11: resource "errorcheck_is_valid" "Not_valid_if_not_match" {
```
