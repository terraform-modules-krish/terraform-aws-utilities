package test

import (
	"testing"

	"github.com/terraform-modules-krish/terratest/modules/collections"
	"github.com/terraform-modules-krish/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestGetEnabledAWSRegions(t *testing.T) {
	t.Parallel()

	terratestOptions := createBaseTerratestOptions(t, "../examples/enabled-aws-regions")
	defer terraform.Destroy(t, terratestOptions)

	terraform.InitAndApply(t, terratestOptions)

	regions := terraform.OutputList(t, terratestOptions, "enabled_regions")

	// Test a few regions that are known to be enabled on our test accounts
	assert.True(t, collections.ListContains(regions, "us-east-1"))
	assert.True(t, collections.ListContains(regions, "us-west-1"))
	assert.True(t, collections.ListContains(regions, "eu-west-1"))

	// ... and verify an opted out region is not included
	assert.False(t, collections.ListContains(regions, "ap-east-1"))      // Hong Kong
	assert.False(t, collections.ListContains(regions, "ap-northeast-3")) // Osaka
}
