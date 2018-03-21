# go-mixpanel

Go bindings to talk to Mixpanel

## Sample code

```go
var mixpanel = mixpanel.NewMixPanelFromEnv(“FOO”)
var distinctID = “SomeID"
if success, err := mixpanel.ProfilePropertyIncrement(distinctID, “field 1"); !success {
    if err != nil {
        t.Error(err)
    } else {
        t.Error(“Increment failure”)
    }
}
if success, err := mixpanel.ProfilePropertyIncrementBy(distinctID, “field 2”, 50); !success {
    if err != nil {
        t.Error(err)
    } else {
        t.Error(“IncrementBy failure”)
    }
}
if success, err := mixpanel.ProfilePropertyDecrement(distinctID, “field 3"); !success {
    if err != nil {
        t.Error(err)
    } else {
        t.Error(“Decrement failure”)
    }
}
if success, err := mixpanel.ProfilePropertyDecrementBy(distinctID, “field 4”, 100); !success {
    if err != nil {
        t.Error(err)
    } else {
        t.Error(“DecrementBy failure”)
    }
}
```
