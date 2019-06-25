# terraform-provider-errorcheck
## A custom terraform provider that can be used to do complex validation during planning and validation


### Installation:

If you want to use the script directly from this repo:
`curl -sSL https://raw.githubusercontent.com/rhythmictech/terraform-provider-errorcheck/master/update-provider.sh | bash`

Otherwise you can download `update-provider.sh` and run it locally

### example:
code:
```terraform
locals {
  compare     = "success"
  testSuccess = "success"
  testFail    = "fail"
}

resource "errorcheck_is_valid" "shouldMatch" {
  test = local.compare == local.testSuccess
}

resource "errorcheck_is_valid" "Not_valid_if_not_match" {
  test = local.compare == local.testFail
}
```
output:
```bash
terraform validate .

Error: Not Valid

  on main.tf line 11, in resource "errorcheck_is_valid" "Not_valid_if_not_match":
  11: resource "errorcheck_is_valid" "Not_valid_if_not_match" {
```
