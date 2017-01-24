package base64_test

import (
	"github.com/chentom88/ecrypto/base64"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Coder", func() {
	Describe("String to Base64", func() {
		DescribeTable("encodes string to base64", func(input, expected string) {
			output, err := base64.Encode(input)
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(Equal(expected))
		},
			Entry("basic1", "hellobob", "aGVsbG9ib2I="),
			Entry("basic2", "ellobob", "ZWxsb2JvYg=="),
			Entry("basic3", "llobob", "bGxvYm9i"),
			Entry("basic5", "B", "Qg=="))

		It("checks for empty string", func() {
			_, err := base64.Encode("")
			Expect(err).To(HaveOccurred())
		})

		DescribeTable("decodes base64 to string", func(input, expected string) {
			output, err := base64.Decode(input)
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(Equal(expected))
		},
			Entry("basic1", "aGVsbG9ib2I=", "hellobob"),
			Entry("basic2", "ZWxsb2JvYg==", "ellobob"),
			Entry("basic3", "bGxvYm9i", "llobob"),
			Entry("basic4", "Qg==", "B"))

		DescribeTable("decode validates input", func(input string) {
			_, err := base64.Decode(input)
			Expect(err).To(HaveOccurred())
		},
			Entry("invalid length", "aGV"),
			Entry("empty string", ""),
			Entry("invalid characters", "aGV?"))

	})

	Describe("Hex string to Base64", func() {
		DescribeTable("checks for valid hex string", func(input string) {
			_, err := base64.EncodeHex(input)

			Expect(err).To(HaveOccurred())
		},
			Entry("Invalid char", "742AT"),
			Entry("Invalid length", "742A0"))

		It("checks for empty string", func() {
			_, err := base64.EncodeHex("")
			Expect(err).To(HaveOccurred())
		})

		DescribeTable("encodes hex string to base64", func(input, expected string) {
			output, err := base64.EncodeHex(input)
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(Equal(expected))
		},
			Entry("basic lower case", "1a2b3c4e", "Gis8Tg=="),
			Entry("basic upper case", "1A2B3C4E", "Gis8Tg=="),
			Entry("basic mixed case", "1a2B3C4e", "Gis8Tg=="),
			Entry("advanced", "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
				"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"))

		DescribeTable("decodes from base64 to hex string", func(input, expected string) {
			output, err := base64.DecodeToHex(input)
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(Equal(expected))
		},
			Entry("basic", "Gis8Tg==", "1a2b3c4e"),
			Entry("advanced", "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
				"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))

		DescribeTable("decode validates input", func(input string) {
			_, err := base64.DecodeToHex(input)
			Expect(err).To(HaveOccurred())
		},
			Entry("invalid length", "aGV"),
			Entry("empty string", ""),
			Entry("invalid characters", "aGV?"))
	})
})
