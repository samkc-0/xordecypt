package main

import (
	"encoding/hex"
	"testing"
)

const m = "THM{p1alntExtAtt4ckcAnr3alLyhUrty0urxOr}"
const k = "Ytewn"
const c = "0d3c280c1e681509191a1c0c11361a2d40061c0d181a17440f35381c1f3b2b001c471b2b0c2a0513"

func TestXorDecrypt(t *testing.T) {
	cipher, err := hex.DecodeString(c)
	if err != nil {
		t.Fatalf("bad hex: %v", err)
	}
	key := []byte(k)
	want := []byte(m)
	got := xorDecrypt(cipher, key)
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("xorDecrypt[%d] = %x, want %x", i, got[i], want[i])
		}
	}
}

func TestRecoverPartialKey(t *testing.T) {
	cipher, err := hex.DecodeString(c)
	if err != nil {
		t.Fatalf("bad hex: %v", err)
	}
	plain := []byte(m)
	want := []byte(k)
	got := recoverPartialKey(cipher, plain)
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("recoverPartialKey[%d] = %x, want %x", i, got[i], want[i])
		}
	}
}

func TestHexToBytesAndBack(t *testing.T) {
	raw := []byte(m)
	hexed := hex.EncodeToString(raw)
	back, err := hex.DecodeString(hexed)
	if err != nil {
		t.Fatalf("decode fail: %v", err)
	}
	if string(raw) != string(back) {
		t.Errorf("round trip failed: got %s, want %s", string(back), string(raw))
	}
}

func TestIsPrintable(t *testing.T) {
	if !isPrintable([]byte("hello123")) {
		t.Error("expected printable input to be printable but it wasn't")
	}
	if isPrintable([]byte{0x01, 0x02, 0x02}) {
		t.Error("expected control bytes to be non-printable, but they are?")
	}
}
