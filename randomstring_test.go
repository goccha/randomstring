package randomstring

import "testing"

func TestGen(t *testing.T) {
	value := Gen(Grow(30), Fix("A"), Now("200601021504"), Numbers(1), Lowers(5), Uppers(3), Format("%05d", 1))
	if len(value) != 27 {
		t.Errorf("expect=%d actual=%d", 27, len(value))
	}
}

func TestGen2(t *testing.T) {
	value := Gen(Grow(70), AlphaNumeric(10), All(10), Base64(10), Base64Url(10), LowersAlphaNumeric(10), UppersAlphaNumeric(10), SymbolAll(10))
	if len(value) != 70 {
		t.Errorf("expect=%d actual=%d", 70, len(value))
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
