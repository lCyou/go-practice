package iteration

import "testing"

func TestRepeat(t *testing.T) {
    repeated := Repeat("a")
    expected := "aaaaa"

    if repeated != expected {
        t.Errorf("expected %q but got %q", expected, repeated)
    }
}

func TestMyRepeat( t *testing.T) {
	repeated := MyRepeat("k", 7)
	ans := "kkkkkkk"

	if repeated != ans {
		t.Errorf("expected: %q ,result: %q", ans, repeated)
	}
}

