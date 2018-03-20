# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]
* Added MixPanel creation methods.
* Added TrackEvent call and the commonly used varients. 
* Added ProfileSet, ProfileSetOnce, ProfileAdd, ProfileAppend, ProfileUnion, ProfileAppend, ProfileRemove, ProfileUnsetDelete api calls which correspond to the mixpanel http specification.
* Added TimeString and CurrentTimeString functions to change time.Time into a string format used by mixpanel.
* Added ProfilePropertyIncrement, ProfilePropertyIncrementBy, ProfilePropertyDecrement, ProfilePropertyDecrementBy convenience functions.
* Added ProfileAddRevenueTransaction convenience function for adding revenue transactions.