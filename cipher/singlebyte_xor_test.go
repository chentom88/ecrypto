package cipher_test

import (
	. "github.com/chentom88/ecrypto/cipher"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SinglebyteXor", func() {
	It("encodes a hex string with byte", func() {
		inputHex := "1a2b3c4d5e6f"
		key := byte(0x01)
		expected := []byte{0x1b, 0x2a, 0x3d, 0x4c, 0x5f, 0x6e}

		actual, err := EncodeSingleByteXOR(inputHex, key)
		Expect(err).ToNot(HaveOccurred())
		Expect(actual).To(Equal(expected))
	})

	It("decodes by encoding ciphertext with same key", func() {
		inputHex := "1b2a3d4c5f6e"
		key := byte(0x01)
		expected := []byte{0x1a, 0x2b, 0x3c, 0x4d, 0x5e, 0x6f}

		actual, err := EncodeSingleByteXOR(inputHex, key)
		Expect(err).ToNot(HaveOccurred())
		Expect(actual).To(Equal(expected))
	})
})
