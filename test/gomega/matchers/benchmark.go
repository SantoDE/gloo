package matchers

import (
	"time"

	"github.com/solo-io/gloo/test/gomega/transforms"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

// HavePercentileLessThan returns a matcher requiring the given slice of durations to be less than the given upperBound at the given percentile
func HavePercentileLessThan(percentile int, upperBound time.Duration) types.GomegaMatcher {
	return gomega.WithTransform(transforms.WithPercentile(percentile), gomega.BeNumerically("<", upperBound))
}

// HaveMedianLessThan returns a matcher requiring the given slice of durations have a median value less than the given upperBound
func HaveMedianLessThan(upperBound time.Duration) types.GomegaMatcher {
	return gomega.WithTransform(transforms.WithMedian(), gomega.BeNumerically("<", upperBound))
}
