# Lyfecycle - Simple application lifecycle via callbacks for Go


## General Usage

```go
// Within your own app config package (or hey maybe right into your only .go file)

import "lyfecycle" // gotta have the package imported to use it

// Declare your lifecycle stages with unique integer values
const (
  LyfeCycleStage1 = iota
  LyfeCycleStage2
  LyfeCycleStage3
  LyfeCycleStage4
  LyfeCycleStage5
)

// during the soonest possible execution moment, register the full list of lifecycle stages
// module init() is about as early as you can get
func init() {
  // sequential stage list from the delcared stage IDs
  stageList := lyfecycle.StageIDsList{
    LyfeCycleStage1,
    LyfeCycleStage2,
    LyfeCycleStage3,
    LyfeCycleStage4,
    LyfeCycleStage5,
  }

  // declare this stage progression with the lyfecyle system
  lyfecycle.DefineStages(stageList)

  // register your callbacks as desired
  lyfecycle.RegisterEvent(LyfeCycleStage1, func() {
    // do something
  })
}

func main() {
  // sometime within your application, run all the stages (stages run in the order declared)
  lyfecycle.PerformAllStages() // all callbacks fire within the stage they belong and in the order they were registered
}
```
