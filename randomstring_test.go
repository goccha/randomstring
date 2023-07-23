package randomstring

import (
	"fmt"
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
	if !(len(value) >= 31 && len(value) <= 41) {
		t.Errorf("expect=(35 <= x <= 45) actual=%d", len(value))
	}
	fmt.Println(value)
}

func TestBuilder(t *testing.T) {
	buf := &strings.Builder{}
	buf.WriteString("*TEST*")
	value := Gen(Fix("Prefix"), Builder(buf), Now(time.RFC3339, time.UTC), AlphaNumeric(10))
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
