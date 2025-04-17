package sdl

import "testing"

func TestInitQuit(t *testing.T) {
	// INIT_CAMERA skipped
	subs := []InitFlags{INIT_AUDIO, INIT_EVENTS, INIT_GAMEPAD, INIT_HAPTIC,
		INIT_JOYSTICK, INIT_SENSOR, INIT_VIDEO}
	var allSubs InitFlags
	for _, s := range subs {
		allSubs |= s
	}
	if err := Init(allSubs); err != nil {
		t.Errorf("Init(EVERYTHING) error: %s", err)
	}
	Quit()
	if WasInit(allSubs) != 0 {
		t.Error("Quit(): subsystems still initialized")
	}
	for _, s := range subs {
		if err := Init(s); err != nil {
			t.Errorf("Init(%d) error: %s", s, err)
		}
		if WasInit(s) != s {
			t.Errorf("Init(%d) subsystem not initialized", s)
		}
		QuitSubSystem(s)
		if WasInit(s) == s {
			t.Errorf("QuitSubSystem(%d): subsystem(s) still initialized", s)
		}
		if err := InitSubSystem(s); err != nil {
			t.Errorf("InitSubSystem(%d) error: %s", s, err)
		}
		if WasInit(s) != s {
			t.Errorf("InitSubSystem(%d): subsystem not initialized", s)
		}
		QuitSubSystem(s)
		if WasInit(s) == s {
			t.Errorf("QuitSubSystem(%d): subsystem(s) still initialized", s)
		}
	}
	Quit()
}
