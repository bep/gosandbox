package unsafestrings

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkSafeBytesToString(b *testing.B) {
	testBytes := []byte("The quick brown fox jumps over the lazy dog.")

	b.ResetTimer()
	var s string
	for i := 0; i < b.N; i++ {
		s = SafeBytesToString(testBytes)
	}
	s = s[:]
}

func BenchmarkUnsafeBytesToString(b *testing.B) {
	testBytes := []byte("The quick brown fox jumps over the lazy dog.")

	b.ResetTimer()
	var s string
	for i := 0; i < b.N; i++ {
		s = UnsafeBytesToString(testBytes)
	}
	s = s[:]
}

var testString = "The quick brown fox jumps over the lazy dog."

func TestSafeBytesToString(t *testing.T) {
	testBytes := []byte(testString)
	s := SafeBytesToString(testBytes)

	if s != testString {
		t.Errorf("Expected '%s' was '%s'", testString, s)
	}

	testBytes[0] = byte('S')

	if s == string(testBytes) {
		t.Errorf("Expected '%s' was '%s'", testBytes, s)
	}

}

func TestUnsafeBytesToString(t *testing.T) {
	testBytes := []byte(testString)
	s := UnsafeBytesToString(testBytes)

	if s != testString {
		t.Errorf("Expected '%s' was '%s'", testString, s)
	}

	testBytes[0] = byte('S')

	if s != string(testBytes) {
		t.Errorf("Expected '%s' was '%s'", testBytes, s)
		t.Errorf("Expected '%s' was '%s'", testBytes, s)
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
	testBytes := []byte("The quick brown fox jumps over the lazy dog.")
	replacer :=
		strings.NewReplacer("quick", "slow", "brown", "blue", "lazy", "energetic")

	buf := make(appendSliceWriter, 0, len(testBytes))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		replacer.WriteString(&buf, UnsafeBytesToString(testBytes))
		if UnsafeBytesToString(buf) !=
			"The slow blue fox jumps over the energetic dog." {
			b.Fatalf("Failed replacement")
		}
		buf = buf[:0] // reuse
	}
}

func BenchmarkSafeStringsReplacer(b *testing.B) {
	testBytes := []byte("The quick brown fox jumps over the lazy dog.")
	replacer :=
		strings.NewReplacer("quick", "slow", "brown", "blue", "lazy", "energetic")

	buf := make(appendSliceWriter, 0, len(testBytes))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		replacer.WriteString(&buf, SafeBytesToString(testBytes))
		if UnsafeBytesToString(buf) !=
			"The slow blue fox jumps over the energetic dog." {
			b.Fatalf("Failed replacement")
		}
		buf = buf[:0]
	}
}

func BenchmarkMultipleBytesReplace(b *testing.B) {
	testBytes := []byte("The quick brown fox jumps over the lazy dog.")

	for i := 0; i < b.N; i++ {
		var replaced []byte

		replaced = bytes.Replace(testBytes, []byte("quick"), []byte("slow"), -1)
		replaced = bytes.Replace(replaced, []byte("brown"), []byte("blue"), -1)
		replaced = bytes.Replace(replaced, []byte("lazy"), []byte("energetic"), -1)

		if UnsafeBytesToString(replaced) != "The slow blue fox jumps over the energetic dog." {
			b.Fatalf("Failed replacement")
		}
	}
}

func BenchmarkMultiplesStringsReplace(b *testing.B) {
	testString := "The quick brown fox jumps over the lazy dog."

	for i := 0; i < b.N; i++ {
		var replaced string

		replaced = strings.Replace(testString, "quick", "slow", -1)
		replaced = strings.Replace(replaced, "brown", "blue", -1)
		replaced = strings.Replace(replaced, "lazy", "energetic", -1)

		if replaced != "The slow blue fox jumps over the energetic dog." {
			b.Fatalf("Failed replacement")
		}
	}
}

func BenchmarkAppendString(b *testing.B) {
	buf := make([]byte, 0, 100)
	s := "bepsays"
	b.ResetTimer()
	var buf2 []byte
	for i := 0; i < b.N; i++ {
		buf2 = append(buf, s...)
	}
	buf2 = buf2[:]
}

func BenchmarkAppendByteString(b *testing.B) {
	buf := make([]byte, 0, 100)
	s := "bepsays"
	b.ResetTimer()
	var buf2 []byte
	for i := 0; i < b.N; i++ {
		buf2 = append(buf, []byte(s)...)
	}
	buf2 = buf2[:]
}
