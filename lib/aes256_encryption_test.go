package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryption(t *testing.T) {
	key := GenerateHex256keyAsStr()
	message := []byte("Hello World!")
	cypher := Aes256Encrypt(key, message)
	messageDecrypted := Aes256Decrypt(key, cypher)
	assert.Equal(t, message, messageDecrypted)
}
