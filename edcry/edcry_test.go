package edcry

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMd5(t *testing.T) {
	t.Log(Md5("hello, world", ""))
	t.Log(Md5("hello, world", "awesome"))
}

func TestSha256(t *testing.T) {
	t.Log(Sha256("hello", ""))
	t.Log(Sha256("hello, world", "awesome"))
}

func TestGenAesKey(t *testing.T) {
	t.Log(GenAesKey16())
	t.Log(GenAesKey32())
}

// plaintext: 明文; ciphertext: 密文
func TestAesEDcrypt(t *testing.T) {
	var (
		plaintext = "窗前明月光,，我是郭德纲"
		key       = GenAesKey32()
	)

	ciphertext := AesEncrypt(key, plaintext)
	plaintext2 := AesDecrypt(key, ciphertext)

	require.Equal(t, plaintext, plaintext2)
}
