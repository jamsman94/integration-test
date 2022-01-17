package downstream_test_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDownstreamTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DownstreamTest Suite")
}
