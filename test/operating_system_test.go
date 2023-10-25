package test

import (
	"github.com/terraform-modules-krish/terratest/modules/terraform"
	"runtime"
	"strings"
	"testing"
)

func TestOperatingSystem(t *testing.T) {
	t.Parallel()

	terratestOptions := createBaseTerratestOptions(t, "../examples/operating-system")
	defer terraform.Destroy(t, terratestOptions)

	terraform.InitAndApply(t, terratestOptions)

	assertOutputEquals(t, "os_name", strings.Title(runtime.GOOS), terratestOptions)
	assertOutputEquals(t, "path_separator", "/", terratestOptions)
}
