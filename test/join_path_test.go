package test

import (
	"github.com/terraform-modules-krish/terratest"
	"testing"
)

func TestJoinPath(t *testing.T) {
	t.Parallel()

	resourceCollection := createBaseRandomResourceCollection(t)
	terratestOptions := createBaseTerratestOptions(t, "TestJoinPath", "../examples/join-path", resourceCollection)
	defer terratest.Destroy(terratestOptions, resourceCollection)

	apply(t, terratestOptions)

	assertOutputEquals(t, "path", "foo/bar/baz.txt", terratestOptions)
}

