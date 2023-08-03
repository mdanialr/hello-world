package app_test

import (
	"testing"

	"github.com/mdanialr/hello-world/app"
)

func TestName(t *testing.T) {
	res := app.Add(1, 2)
	if res != 3 {
		t.Errorf("expect 3 got %d instead", res)
	}
}
