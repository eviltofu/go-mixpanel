# go-mixpanel

Go bindings to talk to Mixpanel.

## Documentation

### Usage

Import the package in order to use it.

```golang
import mixpanel
```

### Creating the mixpanel struct

You need to provide a valid mixpanel token in order to use this. You can provide the token or take it from an environment variable.

```golang
var mixpanel = mixpanel.NewMixPanel("ValidToken")
var mixpanelFromEnv = mixpanel.NewMixPanelFromEnv("Environment name")
```

### Tracking

The method TrackEvent is used for tracking events. Convenience methods are supplied for commonly used combinations. The convenience methods always use the current time for the time stamp of the event.

```golang
TrackEvent(
    event string,
    userID *string,
    timeStamp *time.Time,
    ipAddress *string,
    parameters *map[string]interface{}) error
```

This method returns nil unless an error occurs. Use the following code format:

```golang
import mixpanel
var mixpanel = mixpanel.NewMixPanel("ValidToken")
if err := mixpanel.TrackEventOnly("My Event"); err != nil {
    // report error etc
}
```

#### Tracking convenience methods

When you only wish to track an event.

```golang
TrackEventOnly(event string) error
```

When you want to track an event and add parameters.

```golang
TrackEventWithParameters(event string, parameters map[string]interface{}) error
```

When you want to track an event and associate it with a user.

```golang
TrackEventForUser(event string, userID string) error {
```

When you want to track and event and associate it with a user and some parameters.

```golang
TrackEventForUserWithParameters(event string, userID string, parameters map[string]interface{}) error
```

When you want to track an event and associate it with a user and an ip address.

```golang
TrackEventForUserFromIP(event string, userID string, ipAddress string) error
```

When you want to track an event and associate it with a user, ip address, and some parameters.

```golang
TrackEventForUserFromIPWithParameters(event string, userID string, ipAddress string, parameters map[string]interface{}) error
```

### Profile

When you want to create a user.

```golang
ProfileSet(userID string, attributes map[string]interface{}) error
```

When you want to set some parameters once only. See the [HTTP specifications](https://mixpanel.com/help/reference/http).

```golang
ProfileSetOnce(userID string, attributes map[string]interface{}) error
```

When you want to modify some property counters associated with a user. See the HTTP specificatons.

```golang
ProfileAdd(userID string, attributes map[string]int64) error
```

When you want to append some elements to properties associated with a user. See the [HTTP specifications](https://mixpanel.com/help/reference/http).

```golang
ProfileAppend(userID string, attributes map[string]interface{}) error
```

When you want to append some elements to properties associated with a user and do not want duplicates. See the [HTTP specifications](https://mixpanel.com/help/reference/http).

```golang
ProfileUnion(userID string, attributes map[string]interface{}) error
```

When you want to remove an item from a property. See the [HTTP specifications](https://mixpanel.com/help/reference/http).

```golang
ProfileRemove(userID string, attributes map[string]interface{}) error
```

When you want to remove a property. See the [HTTP specifications](https://mixpanel.com/help/reference/http).

```golang
ProfileUnset(userID string, keyList []string) error
```

When you want to delete a user from mixpanel. You might not see an immediate effect after executing this method. See the [HTTP specifications](https://mixpanel.com/help/reference/http).

```golang
ProfileDelete(userID string) error
```

#### Convenience methods

Increases a property of a user by 1.

```golang
ProfilePropertyIncrement(userID string, property string) error
```

Increases a property of a user by value. Value has to be positive (>0).

```golang
ProfilePropertyIncrementBy(userID string, property string, value int64) error
```

Decreases a property of a user by 1.

```golang
ProfilePropertyDecrement(userID string, property string) error
```

Decreases a property of a user by value. Value has to be positive (>0).

```golang
ProfilePropertyDecrementBy(userID string, property string, value int64) error
```

Adding a transcation with a product code, time stamp, and amount to a user.

```golang
ProfileAddRevenueTransaction(userID string, timeStamp time.Time, productCode string, amount float64) error
```

### Utility functions

Returns the current time in the format mixpanel uses.

```golang
CurrentTimeString() string
```

Returns the given time.Time struct as a string in the format mixpanel uses.

```golang
TimeString(requiredTime time.Time) string
```