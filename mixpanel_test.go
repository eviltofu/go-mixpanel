package mixpanel

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"
)

var ValidTestToken = os.Getenv("FOO")

func TestCreation(t *testing.T) {
	var mixpanel = NewMixPanel(ValidTestToken)
	if mixpanel.Token != ValidTestToken {
		t.Error("Token failure")
	}
}

func TestTrackEventOnly(t *testing.T) {
	var mixpanel = NewMixPanel(ValidTestToken)
	if err := mixpanel.TrackEventOnly("Test TrackEventOnly"); err != nil {
		t.Error(err)
	}
}

func TestTrackEventWithParameters(t *testing.T) {
	var mixpanel = NewMixPanel(ValidTestToken)
	if err := mixpanel.TrackEventWithParameters("Test TrackEventWithParameters", parameters()); err != nil {
		t.Error(err)
	}
}

func TestTrackEventForUser(t *testing.T) {
	var mixpanel = NewMixPanel(ValidTestToken)
	if err := mixpanel.TrackEventForUser("Test TrackEventForUser", "User 0001"); err != nil {
		t.Error(err)
	}
}

func TestTrackEventForUserWithParameters(t *testing.T) {
	var mixpanel = NewMixPanel(ValidTestToken)
	if err := mixpanel.TrackEventForUserWithParameters("Test TrackEventForUserWithParameters", "User 0001", parameters()); err != nil {
		t.Error(err)
	}
}

func TestTrackEventForUserFromIP(t *testing.T) {
	var mixpanel = NewMixPanel(ValidTestToken)
	if err := mixpanel.TrackEventForUserFromIP("Test TrackEventForUserFromIP", "User 0001", "64.2.4.1"); err != nil {
		t.Error(err)
	}
}

func TestTrackEventForUserFromIPWithParameters(t *testing.T) {
	var mixpanel = NewMixPanel(ValidTestToken)
	if err := mixpanel.TrackEventForUserFromIPWithParameters("Test TrackEventForUserFromIPWithParameters", "User 0001", "123.234.5.2", parameters()); err != nil {
		t.Error(err)
	}
}

func parameters() map[string]interface{} {
	return map[string]interface{}{
		"key1": 100,
		"key2": "200",
		"key3": 300.1,
		"key4": map[string]interface{}{
			"innerKey1": map[string]interface{}{"name": "Ancient One", "age": 12343},
		},
		"key5": time.Now().Unix(),
	}
}

// TestProfileSet results in an account being created.
func TestProfileSet(t *testing.T) {
	var distinctID = uniqueID("TestProfileForSet")
	var properties = map[string]interface{}{
		"$first_name": "AccountSet",
		"$last_name":  "ProfileSet",
		"$name":       distinctID,
		"$created":    CurrentTimeString(),
		"$email":      "set.email@someplace.com",
		"$phone":      "6500000000",
	}
	var mixpanel = NewMixPanel(ValidTestToken)
	if err := mixpanel.ProfileSet(distinctID, properties); err != nil {
		t.Error(err)
	}
}

// TestProfileSetOnce fun should be 1.
func TestProfileSetOnce(t *testing.T) {
	var distinctID = uniqueID("TestProfileForSetOnce")
	var properties = map[string]interface{}{
		"$first_name": "AccountSetOnce",
		"$last_name":  "ProfileSetOnce",
		"$name":       distinctID,
		"$created":    CurrentTimeString(),
		"$email":      "setonce.email@someplace.com",
		"$phone":      "6500000001",
		"fun":         1,
	}
	var mixpanel = NewMixPanel(ValidTestToken)
	if err := mixpanel.ProfileSet(distinctID, properties); err != nil {
		t.Error(err)
		return
	}
	var moreProperties = map[string]interface{}{
		"fun": 2,
	}
	if err := mixpanel.ProfileSetOnce(distinctID, moreProperties); err != nil {
		t.Error(err)
	}
}

// TestProfileAdd results in add = 5, minus = -5.
func TestProfileAdd(t *testing.T) {
	var distinctID = uniqueID("TestProfileForAdd")
	var properties = map[string]interface{}{
		"$name":    distinctID,
		"$created": CurrentTimeString(),
		"$email":   "add.email@someplace.com",
		"$phone":   "6500000002",
	}
	var mixpanel = NewMixPanel(ValidTestToken)
	if err := mixpanel.ProfileSet(distinctID, properties); err != nil {
		t.Error(err)
		return
	}
	var addProperties = map[string]int64{
		"add": 5,
	}
	if err := mixpanel.ProfileAdd(distinctID, addProperties); err != nil {
		t.Error(err)
		return
	}
	var minusProperties = map[string]int64{
		"minus": -5,
	}
	if err := mixpanel.ProfileAdd(distinctID, minusProperties); err != nil {
		t.Error(err)
	}
}

// TestProfileAppend results in hobbies = cats, dogs, fish, cats, dogs, rabbits
// movies = movie 1, movie 2
func TestProfileAppend(t *testing.T) {
	var distinctID = uniqueID("TestProfileForAppend")
	var properties = map[string]interface{}{
		"$first_name": "AccountAppend",
		"$last_name":  "ProfileAppend",
		"$name":       distinctID,
		"$created":    CurrentTimeString(),
		"$email":      "append.email@someplace.com",
		"$phone":      "6500000003",
		"hobbies":     []string{"cats", "dogs", "fish"},
	}
	var mixpanel = NewMixPanel(ValidTestToken)
	if err := mixpanel.ProfileSet(distinctID, properties); err != nil {
		t.Error(err)
		return
	}
	var moreProperties = map[string]interface{}{
		"hobbies": []string{"cats", "dogs", "rabbits"},
		"movies":  []string{"movie 1", "movie 2"},
	}
	if err := mixpanel.ProfileAppend(distinctID, moreProperties); err != nil {
		t.Error(err)
	}
}

// TestProfileUnion results in hobbies = cats, dogs, fish, rabbits
// movies = movie 1, movie 2
func TestProfileUnion(t *testing.T) {
	var distinctID = uniqueID("TestProfileForUnion")
	var properties = map[string]interface{}{
		"$first_name": "AccountUnion",
		"$last_name":  "ProfileUnion",
		"$name":       distinctID,
		"$created":    CurrentTimeString(),
		"$email":      "union.email@someplace.com",
		"$phone":      "6500000004",
		"hobbies":     []string{"cats", "dogs", "fish"},
	}
	var mixpanel = NewMixPanel(ValidTestToken)
	if err := mixpanel.ProfileSet(distinctID, properties); err != nil {
		t.Error(err)
		return
	}
	var moreProperties = map[string]interface{}{
		"hobbies": []string{"cats", "dogs", "rabbits"},
		"movies":  []string{"movie 1", "movie 2"},
	}
	if err := mixpanel.ProfileUnion(distinctID, moreProperties); err != nil {
		t.Error(err)
	}
}

// TestProfileRemove results in hobbies = dogs, fish
func TestProfileRemove(t *testing.T) {
	var distinctID = uniqueID("TestProfileForRemove")
	var properties = map[string]interface{}{
		"$first_name": "AccountRemove",
		"$last_name":  "ProfileRemove",
		"$name":       distinctID,
		"$created":    CurrentTimeString(),
		"$email":      "remove.email@someplace.com",
		"$phone":      "6500000005",
		"hobbies":     []string{"cats", "dogs", "fish"},
	}
	var mixpanel = NewMixPanel(ValidTestToken)
	if err := mixpanel.ProfileSet(distinctID, properties); err != nil {
		t.Error(err)
		return
	}
	var moreProperties = map[string]interface{}{
		"hobbies": "cats",
	}
	if err := mixpanel.ProfileRemove(distinctID, moreProperties); err != nil {
		t.Error(err)
		return
	}
}

// TestProfileUnset results in friends = me, you
func TestProfileUnset(t *testing.T) {
	var distinctID = uniqueID("TestProfileForUnset")
	var properties = map[string]interface{}{
		"$first_name": "AccountUnset",
		"$last_name":  "ProfileUnset",
		"$name":       distinctID,
		"$created":    CurrentTimeString(),
		"$email":      "Unset.email@someplace.com",
		"$phone":      "6500000006",
		"hobbies":     []string{"cats", "dogs", "fish"},
		"movies":      []string{"aliens", "predator", "tron", "tremors"},
		"friends":     []string{"me", "you"},
	}
	var mixpanel = NewMixPanel(ValidTestToken)
	if err := mixpanel.ProfileSet(distinctID, properties); err != nil {
		t.Error(err)
		return
	}
	var moreProperties = []string{"movies", "hobbies"}
	if err := mixpanel.ProfileUnset(distinctID, moreProperties); err != nil {
		t.Error(err)
	}
}

// TestProfileDelete results in this account not showing up
func TestProfileDelete(t *testing.T) {
	var distinctID = uniqueID("TestProfileForDelete")
	var properties = map[string]interface{}{
		"$first_name": "AccountDelete",
		"$last_name":  "ProfileDelete",
		"$name":       distinctID,
		"$created":    CurrentTimeString(),
		"$email":      "Delete.email@someplace.com",
		"$phone":      "6500000007",
		"hobbies":     []string{"cats", "dogs", "fish"},
	}
	var mixpanel = NewMixPanel(ValidTestToken)
	if err := mixpanel.ProfileSet(distinctID, properties); err != nil {
		t.Error(err)
		return
	}
	time.Sleep(5 * time.Second)
	if err := mixpanel.ProfileDelete(distinctID); err != nil {
		t.Error(err)
	}
}

func uniqueID(label string) string {
	var result = "user " + strconv.FormatInt(time.Now().Unix(), 16) + "->" + label
	fmt.Println(result)
	return result
}

func TestProfilePropertyIncrement(t *testing.T) {
	var distinctID = uniqueID("TestProfileForProperty")
	var properties = map[string]interface{}{
		"$first_name": "AccountDelete",
		"$last_name":  "ProfileDelete",
		"$name":       distinctID,
		"$created":    CurrentTimeString(),
		"$email":      "Delete.email@someplace.com",
		"$phone":      "6500000008",
		"a":           0,
		"b":           0,
		"c":           0,
		"d":           0,
	}
	var mixpanel = NewMixPanel(ValidTestToken)
	if err := mixpanel.ProfileSet(distinctID, properties); err != nil {
		t.Error(err)
		return
	}
	if err := mixpanel.ProfilePropertyIncrement(distinctID, "a"); err != nil {
		t.Error(err)
		return
	}
	if err := mixpanel.ProfilePropertyIncrementBy(distinctID, "b", 20); err != nil {
		t.Error(err)
		return
	}
	if err := mixpanel.ProfilePropertyDecrement(distinctID, "c"); err != nil {
		t.Error(err)
		return
	}
	if err := mixpanel.ProfilePropertyDecrementBy(distinctID, "d", 50); err != nil {
		t.Error(err)
	}
}

// TestProfileProfileAddRevenueTransaction Total Revenue = 28.25
func TestProfileProfileAddRevenueTransaction(t *testing.T) {
	var distinctID = uniqueID("TestProfileProfileAddRevenueTransaction")
	var properties = map[string]interface{}{
		"$first_name": "AddRevenueTransaction",
		"$last_name":  "AddRevenueTransaction",
		"$name":       distinctID,
		"$created":    CurrentTimeString(),
		"$email":      "AddRevenueTransaction.email@someplace.com",
		"$phone":      "6500000009",
	}
	var mixpanel = NewMixPanel(ValidTestToken)
	if err := mixpanel.ProfileSet(distinctID, properties); err != nil {
		t.Error(err)
		return
	}
	if err := mixpanel.ProfileAddRevenueTransaction(distinctID, time.Now(), "Apple 0001", 12.50); err != nil {
		t.Error(err)
		return
	}
	if err := mixpanel.ProfileAddRevenueTransaction(distinctID, time.Now(), "Google 0002", 13.25); err != nil {
		t.Error(err)
		return
	}
	if err := mixpanel.ProfileAddRevenueTransaction(distinctID, time.Now(), "IBM 0003", 2.50); err != nil {
		t.Error(err)
		return
	}
}
