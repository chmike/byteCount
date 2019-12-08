package bytecount

import (
	"math"
)

// writeFloatSI writes value as float %.1f.
// requires val has been rounded and val < 9995
func writeFloatSI(dst []byte, n, val int, unit byte) int {
	switch {
	case val >= 1000:
		dst[n] = '0' + byte(val/1000)
		n++
		val %= 1000
		fallthrough
	case val >= 100:
		dst[n] = '0' + byte(val/100)
		n++
		val %= 100
	}
	dst[n] = '0' + byte(val/10)
	n++
	dst[n] = '.'
	n++
	dst[n] = '0' + byte(val%10)
	n++
	dst[n] = ' '
	n++
	dst[n] = unit
	n++
	dst[n] = 'B'
	return n + 1
}

// SI return the human readable byte count using the international system notation.
func SI(byteCount int) string {
	var buf [9]byte
	var n int
	if byteCount < 0 {
		buf[0] = '-'
		n++
		if byteCount == math.MinInt64 {
			byteCount = math.MaxInt64
		} else {
			byteCount = -byteCount
		}
	}
	switch {
	case byteCount < 1000:
		switch {
		case byteCount >= 100:
			buf[n] = '0' + byte(byteCount/100)
			n++
			byteCount %= 100
			fallthrough
		case byteCount >= 10:
			buf[n] = '0' + byte(byteCount/10)
			n++
			byteCount %= 10
		}
		buf[n] = '0' + byte(byteCount)
		n++
		buf[n] = ' '
		n++
		buf[n] = 'B'
		return string(buf[:n+1])
	case byteCount < 999950:
		return string(buf[:writeFloatSI(buf[:], n, (byteCount+50)/100, 'k')])
	case byteCount < 999950000:
		return string(buf[:writeFloatSI(buf[:], n, (byteCount+50000)/100000, 'M')])
	case byteCount < 999950000000:
		return string(buf[:writeFloatSI(buf[:], n, (byteCount+50000000)/100000000, 'G')])
	case byteCount < 999950000000000:
		return string(buf[:writeFloatSI(buf[:], n, (byteCount+50000000000)/100000000000, 'T')])
	case byteCount < 999950000000000000:
		return string(buf[:writeFloatSI(buf[:], n, (byteCount+50000000000000)/100000000000000, 'P')])
	default:
		return string(buf[:writeFloatSI(buf[:], n, ((byteCount/1000)+50000000000000)/100000000000000, 'E')])
	}
}

// writeFloatBin writes value as float %.1f.
// requires val has been rounded and val < 10235
func writeFloatBin(dst []byte, n, val int, unit byte) int {
	switch {
	case val >= 10000:
		dst[n] = '0' + byte(val/10000)
		n++
		val %= 10000
		fallthrough
	case val >= 1000:
		dst[n] = '0' + byte(val/1000)
		n++
		val %= 1000
		fallthrough
	case val >= 100:
		dst[n] = '0' + byte(val/100)
		n++
		val %= 100
	}
	dst[n] = '0' + byte(val/10)
	n++
	dst[n] = '.'
	n++
	dst[n] = '0' + byte(val%10)
	n++
	dst[n] = ' '
	n++
	dst[n] = unit
	n++
	dst[n] = 'i'
	n++
	dst[n] = 'B'
	return n + 1
}

// Bin return the human readable byte count using the binary prefix notation.
// See https://en.wikipedia.org/wiki/Binary_prefix.
func Bin(byteCount int) string {
	var buf [11]byte
	var n int
	if byteCount < 0 {
		buf[0] = '-'
		n++
		if byteCount == math.MinInt64 {
			byteCount = math.MaxInt64
		} else {
			byteCount = -byteCount
		}
	}
	switch {
	case byteCount < 1024:
		switch {
		case byteCount >= 1000:
			buf[n] = '0' + byte(byteCount/1000)
			n++
			byteCount %= 1000
			fallthrough
		case byteCount >= 100:
			buf[n] = '0' + byte(byteCount/100)
			n++
			byteCount %= 100
			fallthrough
		case byteCount >= 10:
			buf[n] = '0' + byte(byteCount/10)
			n++
			byteCount %= 10
		}
		buf[n] = '0' + byte(byteCount)
		n++
		buf[n] = ' '
		n++
		buf[n] = 'B'
		return string(buf[:n+1])
	case byteCount <= 0xfffcccccccccccc>>40:
		return string(buf[:writeFloatBin(buf[:], n, ((byteCount*10)+512)>>10, 'k')])
	case byteCount <= 0xfffcccccccccccc>>30:
		return string(buf[:writeFloatBin(buf[:], n, (((byteCount*10)>>10)+512)>>10, 'M')])
	case byteCount <= 0xfffcccccccccccc>>20:
		return string(buf[:writeFloatBin(buf[:], n, (((byteCount*10)>>20)+512)>>10, 'G')])
	case byteCount <= 0xfffcccccccccccc>>10:
		return string(buf[:writeFloatBin(buf[:], n, (((byteCount*10)>>30)+512)>>10, 'T')])
	case byteCount <= 0xfffcccccccccccc:
		return string(buf[:writeFloatBin(buf[:], n, ((((byteCount>>10)*10)>>30)+512)>>10, 'P')])
	default:
		return string(buf[:writeFloatBin(buf[:], n, ((((byteCount>>10)*10)>>40)+512)>>10, 'E')])
	}
}
