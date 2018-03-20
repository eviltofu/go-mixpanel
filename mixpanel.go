// Package mixpanel provides primitives for interacting with mixpanel.com.
package mixpanel

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	trackURL  string = "http://api.mixpanel.com/track/"
	engageURL string = "http://api.mixpanel.com/engage/"
)

// MixPanel represents a client interface to the MixPanel HTTP interface
type MixPanel struct {
	Token string
}

// NewMixPanel creates a new MixPanel.
func NewMixPanel(token string) *MixPanel {
	return &MixPanel{Token: token}
}

// NewMixPanelFromEnv creates a new MixPanel using a token from the environment.
func NewMixPanelFromEnv(env string) *MixPanel {
	return NewMixPanel(os.Getenv(env))
}

func (m *MixPanel) event(data map[string]interface{}) error {
	if err := m.handleHTTPCall(data, trackURL); err != nil {
		fmt.Print(err, data)
		return err
	}
	return nil
}

func (m *MixPanel) profile(data map[string]interface{}) error {
	if err := m.handleHTTPCall(data, engageURL); err != nil {
		fmt.Print(err, data)
		return err
	}
	return nil
}

func (m *MixPanel) handleHTTPCall(data map[string]interface{}, url string) error {
	// convert to JSON
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// fmt.Println(string(jsonBytes))
	// convert to base64
	base64String := base64.StdEncoding.EncodeToString(jsonBytes)
	// make http call
	requestURL := url + "?data=" + base64String
	response, err := http.Get(requestURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, bodyErr := ioutil.ReadAll(response.Body)
	if bodyErr != nil {
		return bodyErr
	}
	if string(body) != "1" && string(body) != "1\n" {
		return errors.New("Error response from mixpanel server")
	}
	return nil
}

// TrackEvent tracks the event.
func (m *MixPanel) TrackEvent(
	event string,
	userID *string,
	timeStamp *time.Time,
	ipAddress *string,
	parameters *map[string]interface{}) error {
	var properties = map[string]interface{}{
		"token": m.Token,
	}
	if userID != nil {
		properties["distinct_id"] = *userID
	}
	if timeStamp != nil {
		properties["time"] = timeStamp.Unix()
	}
	if ipAddress != nil {
		properties["ip"] = *ipAddress
	}
	if parameters != nil {
		properties = mergeMapsCopy(*parameters, properties)
	}
	var packet = map[string]interface{}{
		"event":      event,
		"properties": properties,
	}
	return m.event(packet)
}

// TrackEventOnly tracks an event.
func (m *MixPanel) TrackEventOnly(event string) error {
	var now = time.Now()
	return m.TrackEvent(event, nil, &now, nil, nil)
}

// TrackEventWithParameters tracks an event with parameters.
func (m *MixPanel) TrackEventWithParameters(event string, parameters map[string]interface{}) error {
	var now = time.Now()
	return m.TrackEvent(event, nil, &now, nil, &parameters)
}

// TrackEventForUser tracks an event for the user.
func (m *MixPanel) TrackEventForUser(event string, userID string) error {
	var now = time.Now()
	return m.TrackEvent(event, &userID, &now, nil, nil)
}

// TrackEventForUserWithParameters tracks an event for the user with parameters.
func (m *MixPanel) TrackEventForUserWithParameters(event string, userID string, parameters map[string]interface{}) error {
	var now = time.Now()
	return m.TrackEvent(event, &userID, &now, nil, &parameters)
}

// TrackEventForUserFromIP tracks and event from a user from an ip address.
func (m *MixPanel) TrackEventForUserFromIP(event string, userID string, ipAddress string) error {
	var now = time.Now()
	return m.TrackEvent(event, &userID, &now, &ipAddress, nil)
}

// TrackEventForUserFromIPWithParameters tracks and event from a user from an ip address with parameters.
func (m *MixPanel) TrackEventForUserFromIPWithParameters(event string, userID string, ipAddress string, parameters map[string]interface{}) error {
	var now = time.Now()
	return m.TrackEvent(event, &userID, &now, &ipAddress, &parameters)
}

// ProfileSet follows the http documentation.
// Takes a JSON object containing names and values of profile properties.
// If the profile does not exist, it creates it with these properties.
// If it does exist, it sets the properties to these values, overwriting existing values.
func (m *MixPanel) ProfileSet(userID string, attributes map[string]interface{}) error {
	var properties = map[string]interface{}{
		"$token":       m.Token,
		"$distinct_id": userID,
		"$set":         attributes,
	}
	return m.profile(properties)
}

// ProfileSetOnce follows the http documentation.
// Works just like "$set", except it will not overwrite existing property values.
// This is useful for properties like "First login date".
func (m *MixPanel) ProfileSetOnce(userID string, attributes map[string]interface{}) error {
	var properties = map[string]interface{}{
		"$token":       m.Token,
		"$distinct_id": userID,
		"$set_once":    attributes,
	}
	return m.profile(properties)
}

// ProfileAdd follows the http documentation.
// Takes a JSON object containing keys and numerical values.
// When processed, the property values are added to the existing values of the properties on the profile.
// If the property is not present on the profile, the value will be added to 0.
// It is possible to decrement by calling "$add" with negative values.
// This is useful for maintaining the values of properties like "Number of Logins" or "Files Uploaded".
func (m *MixPanel) ProfileAdd(userID string, attributes map[string]int64) error {
	var properties = map[string]interface{}{
		"$token":       m.Token,
		"$distinct_id": userID,
		"$add":         attributes,
	}
	return m.profile(properties)
}

// ProfileAppend follows the http documentation.
// Takes a JSON object containing keys and values, and appends each to a list associated with the corresponding property name.
// $appending to a property that doesn't exist will result in assigning a list with one element to that property.
func (m *MixPanel) ProfileAppend(userID string, attributes map[string]interface{}) error {
	var properties = map[string]interface{}{
		"$token":       m.Token,
		"$distinct_id": userID,
		"$append":      attributes,
	}
	return m.profile(properties)
}

// ProfileUnion follows the http documentation.
// Takes a JSON object containing keys and list values.
// The list values in the request are merged with the existing list on the user profile, ignoring duplicate list values.
func (m *MixPanel) ProfileUnion(userID string, attributes map[string]interface{}) error {
	var properties = map[string]interface{}{
		"$token":       m.Token,
		"$distinct_id": userID,
		"$union":       attributes,
	}
	return m.profile(properties)
}

// ProfileRemove follows the http documentation.
// The value cannot be a list, array, or object that resolves to a list/array.
// Takes a JSON object containing keys and values.
// The value in the request is removed from the existing list on the user profile.
// If it does not exist, no updates are made.
func (m *MixPanel) ProfileRemove(userID string, attributes map[string]interface{}) error {
	var properties = map[string]interface{}{
		"$token":       m.Token,
		"$distinct_id": userID,
		"$remove":      attributes,
	}
	return m.profile(properties)
}

// ProfileUnset follows the http documentation.
// Takes a JSON list of string property names, and permanently removes the properties and their values from a profile.
func (m *MixPanel) ProfileUnset(userID string, keyList []string) error {
	var properties = map[string]interface{}{
		"$token":       m.Token,
		"$distinct_id": userID,
		"$unset":       keyList,
	}
	return m.profile(properties)
}

// ProfileDelete follows the http documentation.
// Permanently delete the profile from Mixpanel, along with all of its properties.
// The value is ignored - the profile is determined by the $distinct_id from the request itself.
func (m *MixPanel) ProfileDelete(userID string) error {
	var properties = map[string]interface{}{
		"$token":       m.Token,
		"$distinct_id": userID,
		"$delete":      "",
	}
	return m.profile(properties)
}

// CurrentTimeString returns the current time as a string format suitable for use in the mixpanel.
func CurrentTimeString() string {
	var currentTime = time.Now().UTC()
	return TimeString(currentTime)
}

// TimeString returns the time as a string in a format suitable for use in the mixpanel.
func TimeString(requiredTime time.Time) string {
	return requiredTime.Format("2006-01-02T15:04:05")
}

// mergeMapsCopy creates a new map, copies the source map and overwrites the new map with the overwrite map
func mergeMapsCopy(source map[string]interface{}, overwrite map[string]interface{}) (copy map[string]interface{}) {
	copy = make(map[string]interface{})
	for k, v := range source {
		copy[k] = v
	}
	for k, v := range overwrite {
		copy[k] = v
	}
	return
}

// Custom functions

// ProfilePropertyIncrement increments the userID's property by 1
func (m *MixPanel) ProfilePropertyIncrement(userID string, property string) error {
	return m.profilePropertyAdjustBy(userID, property, 1)
}

// ProfilePropertyIncrementBy increments the userID's property by value
// value here must be positive and cannot be zero
func (m *MixPanel) ProfilePropertyIncrementBy(userID string, property string, value int64) error {
	if value <= 0 {
		return errors.New("Value must be greater than zero")
	}
	return m.profilePropertyAdjustBy(userID, property, value)
}

// ProfilePropertyDecrement decrements the userID's property by 1
func (m *MixPanel) ProfilePropertyDecrement(userID string, property string) error {
	return m.profilePropertyAdjustBy(userID, property, -1)
}

// ProfilePropertyDecrementBy decrements the userID's property by value
// value here must be positive and cannot be zero
func (m *MixPanel) ProfilePropertyDecrementBy(userID string, property string, value int64) error {
	if value <= 0 {
		return errors.New("Value must be greater than zero")
	}
	return m.profilePropertyAdjustBy(userID, property, -value)
}

// ProfilePropertyAdjustBy changes the userID's property by value
// value here must be negative
func (m *MixPanel) profilePropertyAdjustBy(userID string, property string, value int64) error {
	var attributes = map[string]int64{
		property: value,
	}
	return m.ProfileAdd(userID, attributes)
}

// ProfileAddRevenueTransaction adds a transaction to the mixpanel
// Not tested yet
func (m *MixPanel) ProfileAddRevenueTransaction(userID string, timeStamp time.Time, productCode string, amount float64) error {
	var properties = map[string]interface{}{
		"$token":       m.Token,
		"$distinct_id": userID,
		"$append": map[string]interface{}{
			"$transactions": map[string]interface{}{
				"$time":        TimeString(timeStamp),
				"$amount":      amount,
				"product_code": productCode,
			},
		},
	}
	return m.profile(properties)
}
