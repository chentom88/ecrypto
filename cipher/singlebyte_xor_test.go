package cipher_test

import (
	. "github.com/chentom88/ecrypto/cipher"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
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

	It("single byte xor encodes string to hex", func() {
		input := "Cooking MC's like a pound of bacon"
		key := string([]byte{0x58})
		expected := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

		actual, err := EncodeSingleByteXORString(input, key)
		Expect(err).ToNot(HaveOccurred())
		Expect(actual).To(Equal(expected))
	})

	DescribeTable("cracks single byte xor", func(input, expected string, key byte) {
		results, err := CrackSingleByteXOR(input)

		Expect(err).ToNot(HaveOccurred())
		Expect(results[0].Decrypted).To(Equal(expected))
		Expect(results[0].Key).To(Equal(key))
	},
		Entry("from challenge", "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736", "Cooking MC's like a pound of bacon", byte(0x58)),
		Entry("mine", myEncoded, myLongString, byte(0x58)),
		Entry("mine with different key", myEncoded2, myString2, byte(0x5a)))
})

const myLongString = `What are the most common n-letter sequences (called "n-grams") for various values of n? You can see the 50 most common for each value of n from 1 to 9 in the table below. The counts and percentages are not shown, but don't worry -- you'll get lots of counts in the next section.`

const myEncoded = `0f30392c78392a3d782c303d7835372b2c783b3735353736783675343d2c2c3d2a782b3d292d3d363b3d2b78703b3934343d3c787a36753f2a39352b7a71783e372a782e392a31372d2b782e39342d3d2b78373e7836677801372d783b3936782b3d3d782c303d786d687835372b2c783b3735353736783e372a783d393b30782e39342d3d78373e7836783e2a37357869782c377861783136782c303d782c393a343d783a3d34372f76780c303d783b372d362c2b7839363c78283d2a3b3d362c393f3d2b78392a3d7836372c782b30372f3674783a2d2c783c37367f2c782f372a2a217875757821372d7f3434783f3d2c7834372c2b78373e783b372d362c2b783136782c303d78363d202c782b3d3b2c31373676`

const myString2 = `In 1894 Griffith's family moved to Denver. She taught first as a substitute teacher and then as a full-time teacher. Griffith became the Deputy State Superintendent of Schools in 1904. She served for six years, leaving twice to go back to the classroom to work with students before returning to her post.`

const myEncoded2 = `13347a6b62636e7a1d28333c3c332e327d297a3c3b373336237a37352c3f3e7a2e357a1e3f342c3f28747a09323f7a2e3b2f3d322e7a3c3328292e7a3b297a3b7a292f38292e332e2f2e3f7a2e3f3b39323f287a3b343e7a2e323f347a3b297a3b7a3c2f3636772e33373f7a2e3f3b39323f28747a1d28333c3c332e327a383f393b373f7a2e323f7a1e3f2a2f2e237a092e3b2e3f7a092f2a3f2833342e3f343e3f342e7a353c7a093932353536297a33347a6b636a6e747a09323f7a293f282c3f3e7a3c35287a2933227a233f3b2829767a363f3b2c33343d7a2e2d33393f7a2e357a3d357a383b39317a2e357a2e323f7a39363b2929283535377a2e357a2d3528317a2d332e327a292e2f3e3f342e297a383f3c35283f7a283f2e2f283433343d7a2e357a323f287a2a35292e74`
