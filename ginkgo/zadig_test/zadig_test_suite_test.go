package zadig_test_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestZadigTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ZadigTest Suite")
}
