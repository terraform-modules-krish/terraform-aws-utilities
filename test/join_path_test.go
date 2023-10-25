package test

import (
	"github.com/terraform-modules-krish/terratest/modules/terraform"
	"testing"
)

func TestJoinPath(t *testing.T) {
	t.Parallel()

	terratestOptions := createBaseTerratestOptions(t, "../examples/join-path")
	defer terraform.Destroy(t, terratestOptions)

	terraform.InitAndApply(t, terratestOptions)

	assertOutputEquals(t, "path", "foo/bar/baz.txt", terratestOptions)
}
