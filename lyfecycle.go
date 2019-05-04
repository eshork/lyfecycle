/*
Package lyfecycle is a simple application lifecycle management package, making it simple for applications to declare
linear lifecycles at initial startup, and for the various modules and plugins embedded within that application
to register callbacks to be invoked as stages of the lifecycle progress.
*/
package lyfecycle

func init() {
	// perform initial setup, get storage objects ready to work
}

// StageID is the type used to individually identify each Lyfecycle stage (tldr: it's an int, numeric order is insignificant, just don't reuse them)
type StageID int

// StageIDsList is what you'd think - an order list of StageID's
type StageIDsList []StageID

// this is the app execution declared StageIDsList -- set once and never again
var definedStageIDsList StageIDsList

// DefineStages declares the Lyfecycle stages for this application.
// It can only be called once per application execution, otherwise it will panic (without updating the Lyfecycle).
func DefineStages(stages StageIDsList) {
	if definedStageIDsList != nil {
		panic("cannot call DefineStages() multiple times; set once and never again")
	}
	// if stages re-uses an integer ID, we should panic

	// if it all passes muster, then this is our new list! (well a copy of it, because we should trust the caller 0%)
	definedStageIDsList = make(StageIDsList, len(stages))
	copy(definedStageIDsList, stages)
}

// GetDefinedStages provides the current defined stages. This will literally be an array of integers, so it's not particularly useful unless
// you know what the integer values specifically mean.
func GetDefinedStages() (stages StageIDsList) {
	// give a copy
	retStages := make(StageIDsList, len(definedStageIDsList))
	copy(retStages, definedStageIDsList)
	return retStages
}

// PerformAllStages executes all Lyfecycle stages, in order.
// Each stage to be ran will run all registered callbacks, in the order in which they were registered.
func PerformAllStages() {
	if definedStageIDsList == nil {
		panic("cannot call PerformAllStages() prior to defining stages via DefineStages()")
	}
	PerformExplicitStages(definedStageIDsList)
}

// PerformExplicitStages executes all given Lyfecycle stages, in the order they were given.
// Each stage to be ran will run all registered callbacks, in the order in which they were registered.
// If a stage is encountered that was not given to DefineStages(), this function will panic.
// This function is literally ran by `PerformAllStages()`, asking it to run all known stages in order.
func PerformExplicitStages(stages StageIDsList) {
	for _, stage := range stages {
		PerformStage(stage)
	}
}

// PerformStage executes the given Lyfecycle stage
func PerformStage(stage StageID) {
	callbackList := stageCallbackListMap[stage]
	for _, callback := range callbackList {
		callback.(func())()
	}
}

type callbackListType []interface{}

type stageCallbackListMapType map[StageID]callbackListType

var stageCallbackListMap stageCallbackListMapType

// RegisterEvent ...
func RegisterEvent(stage StageID, callback interface{}) {
	if stageCallbackListMap == nil {
		stageCallbackListMap = make(stageCallbackListMapType)
	}
	callbackList := stageCallbackListMap[stage]
	stageCallbackListMap[stage] = append(callbackList, callback)
}
