package sdl

import (
	"testing"
)

func TestDisplays(t *testing.T) {
	err := Init(INIT_VIDEO)
	if err != nil {
		t.Errorf("Init(INIT_VIDEO) error: %s", err)
	}

	displays := GetDisplays()
	t.Log(displays)
}
