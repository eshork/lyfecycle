package lyfecycle_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLyfecycle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lyfecycle")
}
