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
