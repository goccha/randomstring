package randomstring

import (
	"strings"
	"testing"
	"time"
)

func TestGen(t *testing.T) {
	value := Gen(Grow(30), Fix("A"), Now("200601021504"), Numbers(1), Lowers(5), Uppers(3), Format("%05d", 1))
	if len(value) != 27 {
		t.Errorf("expect=%d actual=%d", 27, len(value))
	}
}

func TestGen2(t *testing.T) {
	value := Gen(Grow(90), Now(time.RFC3339, time.UTC), AlphaNumeric(10), All(10), Base64(10), Base64Url(10), LowersAlphaNumeric(10), UppersAlphaNumeric(10), SymbolAll(10))
	if len(value) != 90 {
		t.Errorf("expect=%d actual=%d", 90, len(value))
	}
}

func TestGen3(t *testing.T) {
	value := Gen(Grow(45), Now(time.RFC3339, time.UTC), Numbers(10, 20))
	if !(len(value) >= 30 && len(value) <= 40) {
		t.Errorf("expect=(30 <= x <= 40) actual=%d", len(value))
	}
}

func TestGen4(t *testing.T) {
	value := Gen(Grow(1), Fix("A", "B", "C"), Numbers(10))
	if len(value) != 11 {
		t.Errorf("expect=%d actual=%d", 11, len(value))
	}
	switch value[0] {
	case 'A', 'B', 'C':
	default:
		t.Errorf("expect=%s actual=%s", "A or B or C", string(value[0]))
	}
}

func TestGen5(t *testing.T) {
	values := make([]string, 0, 5)
	for i := 0; i < 5; i++ {
		value := Gen(Grow(12), Fix("3"),
			CharSet("23456789abcdefghijklmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ-/#@", 10),
			Fix("5"))
		if contains(values, value) {
			t.Errorf("duplicate value=%s", value)
		}
		values = append(values, value)
	}
}

func TestGen6(t *testing.T) {
	values := make([]string, 0, 5)
	builder := New()
	for i := 0; i < 5; i++ {
		value := builder.Gen(Grow(12), Fix("3"),
			CharSet("23456789abcdefghijklmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ-/#@", 10),
			Fix("5"))
		if contains(values, value) {
			t.Errorf("duplicate value=%s", value)
		}
		values = append(values, value)
	}
}

func contains(s []string, v string) bool {
	for i := range s {
		if v == s[i] {
			return true
		}
	}
	return false
}

func TestMerge(t *testing.T) {
	buf := &strings.Builder{}
	buf.WriteString("*TEST*")
	value := Gen(Fix("Prefix"), Merge(buf), Now(time.RFC3339, time.UTC), AlphaNumeric(10))
	if value[6:12] != "*TEST*" {
		t.Errorf("expect=%s actual=%s", "*TEST*", value[6:12])
	}
}

func BenchmarkGen(b *testing.B) {
	value := Gen(Fix("A"), Numbers(1), Lowers(5), Uppers(3), Fix("Z"))
	if len(value) != 11 {
		b.Errorf("expect=%d actual=%d", 11, len(value))
	}
}

func BenchmarkGrowGen(b *testing.B) {
	value := Gen(Grow(11), Fix("A"), Numbers(1), Lowers(5), Uppers(3), Fix("Z"))
	if len(value) != 11 {
		b.Errorf("expect=%d actual=%d", 11, len(value))
	}
}
