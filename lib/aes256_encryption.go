package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func GenerateHex256keyAsStr() (str string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	return fmt.Sprintf("%X", b)
}

func Aes256Encrypt(keyAsStr string, src []byte) []byte {
	key, _ := hex.DecodeString(keyAsStr)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(src))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], src)

	return ciphertext
}

// Aes256Decrypt from base64 to decrypted string
func Aes256Decrypt(keyAsStr string, cipherbytes []byte) []byte {
	key, _ := hex.DecodeString(keyAsStr)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the cipherbytes.
	if len(cipherbytes) < aes.BlockSize {
		panic("cipherbytes too short")
	}
	iv := cipherbytes[:aes.BlockSize]
	cipherbytes = cipherbytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherbytes, cipherbytes)

	return cipherbytes
}
