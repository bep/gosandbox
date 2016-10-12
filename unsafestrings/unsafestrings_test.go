package unsafestrings

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkSafeBytesToString(b *testing.B) {
	var (
		bt = []byte("The quick brown fox jumps over the lazy dog.")
		s  string
	)

	for i := 0; i < b.N; i++ {
		s = SafeBytesToString(bt)
	}

	s = s[:]
}

func BenchmarkUnsafeBytesToString(b *testing.B) {
	var (
		bt = []byte("The quick brown fox jumps over the lazy dog.")
		s  string
	)

	for i := 0; i < b.N; i++ {
		s = UnsafeBytesToString(bt)
	}

	s = s[:]
}

var testString = "The quick brown fox jumps over the lazy dog."

func TestSafeBytesToString(t *testing.T) {
	var (
		b = []byte(testString)
		s = SafeBytesToString(b)
	)

	if s != testString {
		t.Errorf("Expected '%s' was '%s'", testString, s)
	}

	b[0] = byte('S')

	if s == string(b) {
		t.Errorf("Expected '%s' was '%s'", b, s)
	}
}

func TestUnsafeBytesToString(t *testing.T) {
	var (
		b = []byte(testString)
		s = UnsafeBytesToString(b)
	)

	if s != testString {
		t.Errorf("Expected '%s' was '%s'", testString, s)
	}

	b[0] = byte('S')

	if s != string(b) {
		t.Errorf("Expected '%s' was '%s'", b, s)
	}
}

type appendSliceWriter []byte

func (w *appendSliceWriter) Write(p []byte) (int, error) {
	*w = append(*w, p...)
	return len(p), nil
}

func (w *appendSliceWriter) WriteString(s string) (int, error) {
	*w = append(*w, s...)
	return len(s), nil
}

func BenchmarkUnsafeStringsReplacer(b *testing.B) {
	var (
		by = []byte("The quick brown fox jumps over the lazy dog.")
		re = strings.NewReplacer("quick", "slow", "brown", "blue", "lazy", "energetic")
	)

	buf := make(appendSliceWriter, 0, len(by))

	for i := 0; i < b.N; i++ {
		re.WriteString(&buf, UnsafeBytesToString(by))
		if UnsafeBytesToString(buf) !=
			"The slow blue fox jumps over the energetic dog." {
			b.Fatalf("Failed replacement")
		}
		buf = buf[:0] // reuse
	}
}

func BenchmarkSafeStringsReplacer(b *testing.B) {
	var (
		by = []byte("The quick brown fox jumps over the lazy dog.")
		re = strings.NewReplacer("quick", "slow", "brown", "blue", "lazy", "energetic")
	)

	buf := make(appendSliceWriter, 0, len(by))

	for i := 0; i < b.N; i++ {

		re.WriteString(&buf, SafeBytesToString(by))
		if UnsafeBytesToString(buf) !=
			"The slow blue fox jumps over the energetic dog." {
			b.Fatalf("Failed replacement")
		}
		buf = buf[:0]
	}
}

func BenchmarkMultipleBytesReplace(b *testing.B) {
	by := []byte("The quick brown fox jumps over the lazy dog.")

	for i := 0; i < b.N; i++ {
		var replaced []byte

		replaced = bytes.Replace(by, []byte("quick"), []byte("slow"), -1)
		replaced = bytes.Replace(replaced, []byte("brown"), []byte("blue"), -1)
		replaced = bytes.Replace(replaced, []byte("lazy"), []byte("energetic"), -1)

		if UnsafeBytesToString(replaced) != "The slow blue fox jumps over the energetic dog." {
			b.Fatalf("Failed replacement")
		}
	}
}

func BenchmarkMultiplesStringsReplace(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog."

	for i := 0; i < b.N; i++ {
		var replaced string

		replaced = strings.Replace(s, "quick", "slow", -1)
		replaced = strings.Replace(replaced, "brown", "blue", -1)
		replaced = strings.Replace(replaced, "lazy", "energetic", -1)

		if replaced != "The slow blue fox jumps over the energetic dog." {
			b.Fatalf("Failed replacement")
		}
	}
}

func BenchmarkAppendString(b *testing.B) {
	var (
		buf  = make([]byte, 0, 100)
		buf2 []byte
		s    = "bepsays"
	)

	for i := 0; i < b.N; i++ {
		buf2 = append(buf, s...)
	}
	buf2 = buf2[:]
}

func BenchmarkAppendByteString(b *testing.B) {
	var (
		buf  = make([]byte, 0, 100)
		buf2 []byte
		s    = "bepsays"
	)

	for i := 0; i < b.N; i++ {
		buf2 = append(buf, []byte(s)...)
	}
	buf2 = buf2[:]
}
