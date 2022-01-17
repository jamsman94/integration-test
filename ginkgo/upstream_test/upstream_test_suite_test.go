package upstream_test_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUpstreamTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UpstreamTest Suite")
}
