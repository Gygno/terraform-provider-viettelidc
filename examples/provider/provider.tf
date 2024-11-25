# Configuration-base authentication
provider "viettelidc" {
  domain_id = "3b3e6994-4b04-40ea-bedc-5befd874d73a"
  username  = "iac"
  password  = "Vtdc@12345"
  mfa_code  = var.mfa_code
}
