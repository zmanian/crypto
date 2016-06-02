package sha3

import (
	"gopkg.in/dedis/crypto.v0/test"
	"testing"
)

func TestAES(t *testing.T) {
	test.CipherTest(t, NewCipher224)
}
