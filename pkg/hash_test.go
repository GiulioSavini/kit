package kit

import "testing"

func TestSHA256Hex(t *testing.T) {
	t.Parallel()

	const emptyDigest = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	const abcDigest = "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"

	t.Run("empty string", func(t *testing.T) {
		t.Parallel()
		if got := SHA256Hex(""); got != emptyDigest {
			t.Errorf("SHA256Hex(empty string) = %q, want %q", got, emptyDigest)
		}
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		if got := SHA256Hex("abc"); got != abcDigest {
			t.Errorf("SHA256Hex(abc) = %q, want %q", got, abcDigest)
		}
	})

	t.Run("bytes", func(t *testing.T) {
		t.Parallel()
		if got := SHA256Hex([]byte("abc")); got != abcDigest {
			t.Errorf("SHA256Hex(abc bytes) = %q, want %q", got, abcDigest)
		}
	})

	t.Run("binary bytes", func(t *testing.T) {
		t.Parallel()
		const want = "6e340b9cffb37a989ca544e6bb780a2c78901d3fb33738768511a30617afa01d"
		if got := SHA256Hex([]byte{0}); got != want {
			t.Errorf("SHA256Hex(zero byte) = %q, want %q", got, want)
		}
	})
}
