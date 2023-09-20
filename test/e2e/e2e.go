package e2e

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	// register framework
	_ "github.com/cafe24-dhkim05/operator-for-redis-cluster/test/e2e/framework"
)

// RunE2ETests runs e2e test
func RunE2ETests(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "RedisCluster Suite")
}
