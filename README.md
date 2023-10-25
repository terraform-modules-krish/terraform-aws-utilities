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

[![Maintained by Gruntwork.io](https://img.shields.io/badge/maintained%20by-gruntwork.io-%235849a6.svg)](https://gruntwork.io/?ref=repo_package_terraform_utilities)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/gruntwork-io/terraform-aws-utilities.svg?label=latest)](https://github.com/gruntwork-io/terraform-aws-utilities/releases/latest)
![Terraform Version](https://img.shields.io/badge/tf-%3E%3D1.0.0-blue.svg)

# Terraform Utility Modules

This repo contains miscellaneous utility and helper modules for use with Terraform.

## What is in this repo

This repo provides a Gruntwork IaC Package and has the following folder structure:

* [modules](https://github.com/terraform-modules-krish/terraform-aws-utilities/blob/v0.6.0/modules): This folder contains the main implementation code for this repository, broken down into multiple
  standalone modules.
* [examples](https://github.com/terraform-modules-krish/terraform-aws-utilities/blob/v0.6.0/examples): This folder contains examples of how to use the modules.
* [test](https://github.com/terraform-modules-krish/terraform-aws-utilities/blob/v0.6.0/test): Automated tests for the modules and examples.

The following modules are available:

* [join-path](https://github.com/terraform-modules-krish/terraform-aws-utilities/blob/v0.6.0/modules/join-path): This module can be used to join a list of given path parts into a single path that is
  platform/operating system aware. **(This module requires Python)**
* [operating-system](https://github.com/terraform-modules-krish/terraform-aws-utilities/blob/v0.6.0/modules/operating-system): This module can be used to figure out what operating system is being
  used to run Terraform. **(This module requires Python)**
* [require-executable](https://github.com/terraform-modules-krish/terraform-aws-utilities/blob/v0.6.0/modules/require-executable): This is a module that can be used to ensure particular executables
  is available in the `PATH`. **(This module requires Python)**
* [run-pex-as-data-source](https://github.com/terraform-modules-krish/terraform-aws-utilities/blob/v0.6.0/modules/run-pex-as-data-source): This module prepares a portable environment for running PEX
  files and runs them as an external data source. PEX files are python executables that contain all the requirements
  necessary to run the script. **(This module requires Python)**
* [run-pex-as-resource](https://github.com/terraform-modules-krish/terraform-aws-utilities/blob/v0.6.0/modules/run-pex-as-resource): This module prepares a portable environment for running PEX files
  and runs them as an local-exec provisioner on a null_resource. PEX files are python executables that contain all the
  requirements necessary to run the script. **(This module requires Python)**

The following modules were deprecated and removed:

* `intermediate-variable`: This module has been superseded by [terraform local
  values](https://www.terraform.io/docs/configuration/locals.html). To upgrade, switch usage of `intermediate-variable`
  with `locals`.
* `enabled-aws-regions`: This module has been superseded by [terraform aws_regions data
  source](https://www.terraform.io/docs/providers/aws/d/regions.html). To upgrade, switch the module block with:

    data "aws_regions" "enabled_regions" {}

  Then, you can get the list of enabled regions using `data.aws_regions.enabled_regions.names`.


Click on each module above to see its documentation. Head over to the [examples](https://github.com/terraform-modules-krish/terraform-aws-utilities/blob/v0.6.0/examples) folder for example usage.




## What is a module?

A Module is a canonical, reusable, best-practices definition for how to run a single piece of infrastructure, such as a
database or server cluster. Each Module is written using a combination of Terraform and scripts (mostly bash) and
include automated tests, documentation, and examples. It is maintained both by the open source community and companies
that provide commercial support.

Instead of figuring out the details of how to run a piece of infrastructure from scratch, you can reuse existing code
that has been proven in production. And instead of maintaining all that infrastructure code yourself, you can leverage
the work of the Module community to pick up infrastructure improvements through a version number bump.



## Who maintains this Module?

This Module is maintained by [Gruntwork](http://www.gruntwork.io/). If you're looking for help or commercial
support, send an email to [modules@gruntwork.io](mailto:modules@gruntwork.io?Subject=Terraform%20Utilities%20Module).
Gruntwork can help with:

* Setup, customization, and support for this Module.
* Modules for other types of infrastructure, such as VPCs, Docker clusters, databases, and continuous integration.
* Modules that meet compliance requirements, such as HIPAA.
* Consulting & Training on AWS, Terraform, and DevOps.




## How is this Module versioned?

This Module follows the principles of [Semantic Versioning](http://semver.org/). You can find each new release,
along with the changelog, in the [Releases Page](../../releases).

During initial development, the major version will be 0 (e.g., `0.x.y`), which indicates the code does not yet have a
stable API. Once we hit `1.0.0`, we will make every effort to maintain a backwards compatible API and use the MAJOR,
MINOR, and PATCH versions on each release to indicate any incompatibilities.





## License

Please see [LICENSE.txt](https://github.com/terraform-modules-krish/terraform-aws-utilities/blob/v0.6.0/LICENSE.txt) and [NOTICE](https://github.com/terraform-modules-krish/terraform-aws-utilities/blob/v0.6.0/NOTICE) for details on how the code in this repo is licensed.
