***WARNING: THIS REPO IS AN AUTO-GENERATED COPY.*** *This repo has been copied from [Gruntwork’s](https://gruntwork.io/) GitHub repositories so that you can consume it from your company’s own internal Git repositories. This copy is automatically created and updated by the `repo-copier` CLI tool. If you need to make changes to this repo, you should make the changes in a separate fork, and NOT make changes directly in this repo, as otherwise, the `repo-copier` will overwrite your changes! Please see the `repo-copier` [documentation](https://github.com/terraform-modules-krish/repo-copier) for more information on how the code is copied, how cross-references are updated, how the changelog is handled, etc.*

***

_You may find it valuable to view the following resources in the original repo. If these links give you a 404, visit https://app.gruntwork.io to gain access or email support@gruntwork.io if you need assistance._

[Home Page](https://github.com/gruntwork-io/terraform-aws-utilities/) |
[Pull Requests](https://github.com/gruntwork-io/terraform-aws-utilities/pulls) |
[Issues](https://github.com/gruntwork-io/terraform-aws-utilities/issues) |
[Releases and Assets](https://github.com/gruntwork-io/terraform-aws-utilities/releases)

_Alternatively, you can view a copied version of the resources listed above._

[Pull Requests](https://github.com/terraform-modules-krish/terraform-aws-utilities/blob/main/.github/PULL_REQUESTS.md) |
[Issues](https://github.com/terraform-modules-krish/terraform-aws-utilities/blob/main/.github/ISSUES.md) |
[ChangeLog](https://github.com/terraform-modules-krish/terraform-aws-utilities/blob/main/.github/CHANGELOG.md)

***

# Terraform Utility Modules

This repo contains miscellaneous utility and helper modules for use with Terraform. 

The modules are:

* [intermediate-variable](https://github.com/terraform-modules-krish/terraform-aws-utilities/blob/v0.0.2/modules/intermediate-variable): A simple module that returns as output the exact variables 
  you pass to it as inputs. This gives you a way to store intermediate values that contain interpolations.

Click on each module above to see its documentation. Head over to the [examples](https://github.com/terraform-modules-krish/terraform-aws-utilities/blob/v0.0.2/examples) folder for examples.




## What is a module?

At [Gruntwork](http://www.gruntwork.io), we've taken the thousands of hours we spent building infrastructure on AWS and
condensed all that experience and code into pre-built **packages** or **modules**. Each module is a battle-tested,
best-practices definition of a piece of infrastructure, such as a VPC, ECS cluster, or an Auto Scaling Group. Modules
are versioned using [Semantic Versioning](http://semver.org/) to allow Gruntwork clients to keep up to date with the
latest infrastructure best practices in a systematic way.




## How do you use a module?

Most of our modules contain either:

1. [Terraform](https://www.terraform.io/) code
1. Scripts & binaries


### Using a Terraform Module

To use a module in your Terraform templates, create a `module` resource and set its `source` field to the Git URL of
this repo. You should also set the `ref` parameter so you're fixed to a specific version of this repo, as the `master`
branch may have backwards incompatible changes (see [module
sources](https://www.terraform.io/docs/modules/sources.html)).

For example, to use `v1.0.8` of the ecs-cluster module, you would add the following:

```hcl
module "ecs_cluster" {
  source = "git::git@github.com:gruntwork-io/module-ecs.git//modules/ecs-cluster?ref=v1.0.8"

  // set the parameters for the ECS cluster module
}
```

*Note: the double slash (`//`) is intentional and required. It's part of Terraform's Git syntax (see [module
sources](https://www.terraform.io/docs/modules/sources.html)).*

See the module's documentation and `vars.tf` file for all the parameters you can set. Run `terraform get -update` to
pull the latest version of this module from this repo before runnin gthe standard  `terraform plan` and
`terraform apply` commands.


### Using scripts & binaries

You can install the scripts and binaries in the `modules` folder of any repo using the [Gruntwork
Installer](https://github.com/terraform-modules-krish/gruntwork-installer). For example, if the scripts you want to install are
in the `modules/ecs-scripts` folder of the https://github.com/terraform-modules-krish/module-ecs repo, you could install them
as follows:

```bash
gruntwork-install --module-name "ecs-scripts" --repo "https://github.com/terraform-modules-krish/module-ecs" --tag "0.0.1"
```

See the docs for each script & binary for detailed instructions on how to use them.




## Developing a module

### Versioning

We are following the principles of [Semantic Versioning](http://semver.org/). During initial development, the major
version is to 0 (e.g., `0.x.y`), which indicates the code does not yet have a stable API. Once we hit `1.0.0`, we will
follow these rules:

1. Increment the patch version for backwards-compatible bug fixes (e.g., `v1.0.8 -> v1.0.9`).
2. Increment the minor version for new features that are backwards-compatible (e.g., `v1.0.8 -> 1.1.0`).
3. Increment the major version for any backwards-incompatible changes (e.g. `1.0.8 -> 2.0.0`).

The version is defined using Git tags.  Use GitHub to create a release, which will have the effect of adding a git tag.


### Tests

See the [test](https://github.com/terraform-modules-krish/terraform-aws-utilities/blob/v0.0.2/test) folder for details.




## License

Please see [LICENSE.txt](https://github.com/terraform-modules-krish/terraform-aws-utilities/blob/v0.0.2/LICENSE.txt) for details on how the code in this repo is licensed.
