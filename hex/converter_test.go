package hex_test

import (
	. "github.com/chentom88/ecrypto/hex"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Converter", func() {
	Describe("Hex To String", func() {
		DescribeTable("converts valid hex", func(input, expectedOut string) {
			output, err := HexStringToString(input)
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(Equal(expectedOut))
		},
			Entry("basic lower case", "48656c6c6f426f62", "HelloBob"),
			Entry("basic upper case", "48656C6C6F426F62", "HelloBob"),
			Entry("basic mix case", "48656c6C6f426F62", "HelloBob"),
			Entry("advanced", "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
				"I'm killing your brain like a poisonous mushroom"))

		DescribeTable("checks for valid hex", func(input string) {
			_, err := HexStringToString(input)
			Expect(err).To(HaveOccurred())
		},
			Entry("odd length", "1a3"),
			Entry("empty string", ""),
			Entry("invalid character", "1z2a"))
	})

	Describe("String to Hex", func() {
		DescribeTable("converts string to valid hex", func(input, expectedOut string) {
			output, err := StringToHexString(input)
			Expect(err).ToNot(HaveOccurred())
			Expect(output).To(Equal(expectedOut))
		},
			Entry("basic", "HelloBob", "48656c6c6f426f62"),
			Entry("advanced", "I'm killing your brain like a poisonous mushroom", "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
	})

	It("string to hex string checks for valid input", func() {
		_, err := StringToHexString("")
		Expect(err).To(HaveOccurred())
	})
})
