// +build noasm,arm64 !amd64

package p384

import (
	"math/bits"
)


func fp384Cmov(x, y *fp384, b int) {
}

func fp384Neg(c, a *fp384) {
    var uint64 carry, borrow

    for i := 0; i < FpWords; i++ {
        c[i], borrow = bits.Sub64(p[i], a[i], borrow)
    }

    mask := uint64(0 - carry)
    carry = 0
    for i := 0; i < FpWords; i++ {
        c[i], carry = bits.Add64(c[i], p[i]&mask, carry)
    }

}

// Compute z = x + y % p
func fp384Add(c, a, b *fp384) {
    var carry uint64

    // c = a + b
    for i := 0; i < FpWords; i++ {
        c[i], carry = bits.Add64(a[i], b[i], carry)
    }

    // c = c - P384
    carry = 0
    for i := 0; i < FpWords; i++ {
        c[i], carry = bits.Sub64(c[i], p[i], carry)
    }

    // if c < 0 add P384 back
    mask := uint64(0 - carry)
    carry = 0
    for i := 0; i < FpWords; i++ {
        c[i], carry = bits.Add64(c[i], p[i]&mask, carry)
    }
}


// Compute c = x - y % p
func fp384Sub(c, a, b *fp384) {
    var borrow uint64

	// c = a - b
    for i := 0; i < FpWords; i++ {
        c[i], borrow = bits.Sub64(a[i], b[i], borrow)
    }

    mask := uint64(0 - borrow)
    borrow = 0

	// if c < 0 c = a - b + P384
    for i := 0; i < FpWords; i++ {
        c[i], borrow = bits.Add64(c[i], p[i]&mask, borrow)
    }
}

func fp384Mul(c, a, b *fp384) {

}

