package bytecount

import (
	"fmt"
	"math"
	"testing"
)

func TestSI(t *testing.T) {
	for i := -1000000; i < 1000000; i++ {
		if SI(i) != SI1(i) {
			t.Fatalf("compare mismatch SI(%d)=%s SI1(%d)=%s", i, SI(i), i, SI1(i))
		}
	}

	tests := []struct {
		id  int
		val int
		res string
	}{
		{id: 0, val: 0, res: "0 B"},
		{id: 1, val: math.MinInt64, res: "-9.2 EB"},
		{id: 2, val: math.MaxInt64, res: "9.2 EB"},
		{id: 3, val: math.MinInt64 + 10000000000, res: "-9.2 EB"},
		{id: 4, val: math.MaxInt64 - 10000000000, res: "9.2 EB"},
		{id: 5, val: 999949, res: "999.9 kB"},
		{id: 6, val: 999950, res: "1.0 MB"},
		{id: 7, val: -999949, res: "-999.9 kB"},
		{id: 8, val: -999950, res: "-1.0 MB"},
		{id: 9, val: 999949000, res: "999.9 MB"},
		{id: 10, val: 999950000, res: "1.0 GB"},
		{id: 11, val: -999949000, res: "-999.9 MB"},
		{id: 12, val: -999950000, res: "-1.0 GB"},
		{id: 13, val: 999949000000, res: "999.9 GB"},
		{id: 14, val: 999950000000, res: "1.0 TB"},
		{id: 15, val: -999949000000, res: "-999.9 GB"},
		{id: 16, val: -999950000000, res: "-1.0 TB"},
		{id: 17, val: 999949000000000, res: "999.9 TB"},
		{id: 18, val: 999950000000000, res: "1.0 PB"},
		{id: 19, val: -999949000000000, res: "-999.9 TB"},
		{id: 20, val: -999950000000000, res: "-1.0 PB"},
		{id: 21, val: 999949000000000000, res: "999.9 PB"},
		{id: 22, val: 999950000000000000, res: "1.0 EB"},
		{id: 23, val: -999949000000000000, res: "-999.9 PB"},
		{id: 24, val: -999950000000000000, res: "-1.0 EB"},
	}
	for _, test := range tests {
		if SI(test.val) != test.res {
			t.Errorf("%d. expected SI(%d)=%s, got %s", test.id, test.val, test.res, SI(test.val))
		}
		if SI(test.val) != SI1(test.val) {
			t.Errorf("%d. compare mismatch SI(%d)=%s SI1(%d)=%s", test.id, test.val, SI(test.val), test.val, SI1(test.val))
		}
	}
}

func SI1(byteCount int) string {
	var s string
	if byteCount < 0 {
		s = "-"
		if byteCount == math.MinInt64 {
			byteCount = math.MaxInt64
		} else {
			byteCount = -byteCount
		}
	}
	switch {
	case byteCount < 1000:
		return fmt.Sprintf("%s%d B", s, byteCount)
	case byteCount < 999950:
		return fmt.Sprintf("%s%.1f kB", s, math.Round(float64(byteCount*10)/1e3)/1e1)
	case byteCount < 999950000:
		return fmt.Sprintf("%s%.1f MB", s, math.Round(float64(byteCount*10)/1e6)/1e1)
	case byteCount < 999950000000:
		return fmt.Sprintf("%s%.1f GB", s, math.Round(float64(byteCount*10)/1e9)/1e1)
	case byteCount < 999950000000000:
		return fmt.Sprintf("%s%.1f TB", s, math.Round(float64(byteCount*10)/1e12)/1e1)
	case byteCount < 999950000000000000:
		return fmt.Sprintf("%s%.1f PB", s, math.Round(float64(byteCount/100)/1e12)/1e1)
	default:
		return fmt.Sprintf("%s%.1f EB", s, math.Round(float64(byteCount/100)/1e15)/1e1)
	}
}

func TestBin(t *testing.T) {
	for i := -1000000; i < 1000000; i++ {
		if SI(i) != SI1(i) {
			t.Fatalf("compare mismatch Bin(%d)=%s Bin1(%d)=%s", i, Bin(i), i, Bin1(i))
		}
	}

	tests := []struct {
		id  int
		val int
		res string
	}{
		{id: 0, val: 0, res: "0 B"},
		{id: 1, val: math.MinInt64, res: "-8.0 EiB"},
		{id: 2, val: math.MaxInt64, res: "8.0 EiB"},
		{id: 3, val: math.MinInt64 + 10000000000, res: "-8.0 EiB"},
		{id: 4, val: math.MaxInt64 - 10000000000, res: "8.0 EiB"},
		{id: 5, val: 0xFFFCCCCCCCCCCCC >> 40, res: "1023.9 kiB"},
		{id: 6, val: 0xFFFCCCCCCCCCCCC>>40 + 1, res: "1.0 MiB"},
		{id: 7, val: -(0xFFFCCCCCCCCCCCC >> 40), res: "-1023.9 kiB"},
		{id: 8, val: -(0xFFFCCCCCCCCCCCC>>40 + 1), res: "-1.0 MiB"},
		{id: 9, val: 0xFFFCCCCCCCCCCCC >> 30, res: "1023.9 MiB"},
		{id: 10, val: 0xFFFCCCCCCCCCCCC>>30 + 1, res: "1.0 GiB"},
		{id: 11, val: -(0xFFFCCCCCCCCCCCC >> 30), res: "-1023.9 MiB"},
		{id: 12, val: -(0xFFFCCCCCCCCCCCC>>30 + 1), res: "-1.0 GiB"},
		{id: 13, val: 0xFFFCCCCCCCCCCCC >> 20, res: "1023.9 GiB"},
		{id: 14, val: 0xFFFCCCCCCCCCCCC>>20 + 1, res: "1.0 TiB"},
		{id: 15, val: -(0xFFFCCCCCCCCCCCC >> 20), res: "-1023.9 GiB"},
		{id: 16, val: -(0xFFFCCCCCCCCCCCC>>20 + 1), res: "-1.0 TiB"},
		{id: 17, val: 0xFFFCCCCCCCCCCCC >> 10, res: "1023.9 TiB"},
		{id: 18, val: 0xFFFCCCCCCCCCCCC>>10 + 1, res: "1.0 PiB"},
		{id: 19, val: -(0xFFFCCCCCCCCCCCC >> 10), res: "-1023.9 TiB"},
		{id: 20, val: -(0xFFFCCCCCCCCCCCC>>10 + 1), res: "-1.0 PiB"},
		{id: 21, val: 0xFFFCCCCCCCCCCCC, res: "1023.9 PiB"},
		{id: 22, val: 0xFFFCCCCCCCCCCCC + 1, res: "1.0 EiB"},
		{id: 23, val: -0xFFFCCCCCCCCCCCC, res: "-1023.9 PiB"},
		{id: 24, val: -0xFFFCCCCCCCCCCCC - 1, res: "-1.0 EiB"},
	}
	for _, test := range tests {
		if Bin(test.val) != test.res {
			t.Errorf("%d. expected Bin(%d)=%s, got %s", test.id, test.val, test.res, Bin(test.val))
		}
		if Bin1(test.val) != test.res {
			t.Errorf("%d. expected Bin1(%d)=%s, got %s", test.id, test.val, test.res, Bin1(test.val))
		}
	}
}

func Bin1(byteCount int) string {
	var s string
	if byteCount < 0 {
		s = "-"
		if byteCount == math.MinInt64 {
			byteCount = math.MaxInt64
		} else {
			byteCount = -byteCount
		}
	}
	switch {
	case byteCount < 1024:
		return fmt.Sprintf("%s%d B", s, byteCount)
	case byteCount <= 0xfffcccccccccccc>>40:
		return fmt.Sprintf("%s%.1f kiB", s, math.Round(float64(byteCount*10)/float64(1<<10))/1e1)
	case byteCount <= 0xfffcccccccccccc>>30:
		return fmt.Sprintf("%s%.1f MiB", s, math.Round(float64(byteCount*10)/float64(1<<20))/1e1)
	case byteCount <= 0xfffcccccccccccc>>20:
		return fmt.Sprintf("%s%.1f GiB", s, math.Round(float64(byteCount*10)/float64(1<<30))/1e1)
	case byteCount <= 0xfffcccccccccccc>>10:
		return fmt.Sprintf("%s%.1f TiB", s, math.Round(float64(byteCount*10)/float64(1<<40))/1e1)
	case byteCount <= 0xfffcccccccccccc:
		return fmt.Sprintf("%s%.1f PiB", s, math.Round(float64((byteCount>>10)*10)/float64(1<<40))/1e1)
	default:
		return fmt.Sprintf("%s%.1f EiB", s, math.Round(float64((byteCount>>10)*10)/float64(1<<50))/1e1)
	}
}

var s string

func BenchmarkSI(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for magn := 1000; magn < 1000000000000001; magn *= 1000 {
			for j := magn*1000 - magn; j < magn*1000; j += magn / 20 {
				s = SI(j)
			}
		}
	}
}

func BenchmarkSI1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for magn := 1000; magn < 1000000000000001; magn *= 1000 {
			for j := magn*1000 - magn; j < magn*1000; j += magn / 20 {
				s = SI1(j)
			}
		}
	}
}

func BenchmarkBin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for magn := 1024; magn < 1<<50+1; magn *= 1024 {
			for j := magn*1024 - magn; j < magn*1024; j += magn / 16 {
				s = Bin(j)
			}
		}
	}
}

func BenchmarkBin1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for magn := 1024; magn < 1<<50+1; magn *= 1024 {
			for j := magn*1024 - magn; j < magn*1024; j += magn / 16 {
				s = Bin1(j)
			}
		}
	}
}
