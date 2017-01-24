package base64_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBase64(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Base64 Suite")
}
