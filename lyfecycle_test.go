/*
Package lyfecycle_test is a test package based on an external representation of normal usage.
It does not have access to Lyfecycle internals, nor does it need them, and nor _should_ it need them.
If this test package passes, then users should expect to be satisfied with the Lyfecycle implementation.
*/
package lyfecycle_test

import (
	"github.com/eshork/lyfecycle" // gotta have the package imported to use it
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// From here, for a bit, we're defining an app as we normally would
// BEGIH NORMAL'ISH APP BEHAVIOUR

// Declare your lifecycle stages with unique integer values
const (
	LyfeCycleStage1 = iota
	LyfeCycleStage2
	LyfeCycleStage3
	LyfeCycleStage4
	LyfeCycleStage5
)

var calledLyfeCycleStage1 bool
var calledLyfeCycleStage2 bool
var calledLyfeCycleStage3 bool
var calledLyfeCycleStage4 bool
var calledLyfeCycleStage5 bool

var test_stageList lyfecycle.StageIDsList // in a normal app, you probably wouldn't keep a global reference to this var - this one is only here to support testing

// during the soonest possible execution moment, register the full list of lifecycle stages
// module init() is about as early as you can get
func init() {
	// sequential stage list from the delcared stage IDs
	test_stageList = lyfecycle.StageIDsList{ // shoving this into a global variable so we can validate it later...
		LyfeCycleStage1,
		LyfeCycleStage2,
		LyfeCycleStage3,
		LyfeCycleStage4,
		LyfeCycleStage5,
	}
	// declare this stage progression with the lyfecyle system
	lyfecycle.DefineStages(test_stageList)

	// register your callbacks as desired
	lyfecycle.RegisterEvent(LyfeCycleStage1, func() {
		calledLyfeCycleStage1 = true
	})
	lyfecycle.RegisterEvent(LyfeCycleStage2, func() {
		calledLyfeCycleStage2 = true
	})
	lyfecycle.RegisterEvent(LyfeCycleStage3, func() {
		calledLyfeCycleStage3 = true
	})
	lyfecycle.RegisterEvent(LyfeCycleStage4, func() {
		calledLyfeCycleStage4 = true
	})
	lyfecycle.RegisterEvent(LyfeCycleStage5, func() {
		calledLyfeCycleStage5 = true
	})
	lyfecycle.PerformAllStages()
}

// END NORMAL'ISH APP BEHAVIOUR
// This is where we start testing

var _ = Describe("Lyfecycle", func() {

	BeforeEach(func() {
		calledLyfeCycleStage1 = false
		calledLyfeCycleStage2 = false
		calledLyfeCycleStage3 = false
		calledLyfeCycleStage4 = false
		calledLyfeCycleStage5 = false
	})

	Context("GetDefinedStages()", func() {
		It("reflects the stages we created in our app init", func() {
			Expect(lyfecycle.GetDefinedStages()).Should(BeEquivalentTo(test_stageList))
		})
	})
	Context("PerformAllStages()", func() {
		It("runs each stage", func() {
			Expect(calledLyfeCycleStage1).Should(BeFalse())
			Expect(calledLyfeCycleStage2).Should(BeFalse())
			Expect(calledLyfeCycleStage3).Should(BeFalse())
			Expect(calledLyfeCycleStage4).Should(BeFalse())
			Expect(calledLyfeCycleStage5).Should(BeFalse())
			lyfecycle.PerformAllStages()
			Expect(calledLyfeCycleStage1).Should(BeTrue())
			Expect(calledLyfeCycleStage2).Should(BeTrue())
			Expect(calledLyfeCycleStage3).Should(BeTrue())
			Expect(calledLyfeCycleStage4).Should(BeTrue())
			Expect(calledLyfeCycleStage5).Should(BeTrue())
		})
	})
})
