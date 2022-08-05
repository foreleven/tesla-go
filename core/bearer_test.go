package core_test

import (
	"strings"
	"testing"
)

func TestBearer(t *testing.T) {
	const test = "Bearer abc"
	parts := strings.SplitN(test, " ", 2)

	if parts[0] == "Bearer" {
		t.Logf("Bearer: %s", parts[1])
	}
}
