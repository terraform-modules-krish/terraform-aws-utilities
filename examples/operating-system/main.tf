module "os" {
  # When using these modules in your own templates, you will need to use a Git URL with a ref attribute that pins you
  # to a specific version of the modules, such as the following example:
  # source = "git::git@github.com:terraform-modules-krish/package-terraform-utilities.git//modules/operating-system?ref=v1.0.8"
  source = "../../modules/operating-system"
}