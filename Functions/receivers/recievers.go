package receivers

// method receivers are just like function parameters
// they are just a way to pass the receiver to the method
// so that the method can access the receiver's fields
// and methods &  more like extension methods on other languages

import "fmt"

type SemanticVersion struct {
	major, minor, patch int
}

func NewSemanticVersion(major, minor, patch int) SemanticVersion {
	return SemanticVersion{major, minor, patch}
}

// value typed receivers
func (s SemanticVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", s.major, s.minor, s.patch)
}

func (s SemanticVersion) EqualityCheck(other SemanticVersion) bool {
	return s.major == other.major && s.minor == other.minor && s.patch == other.patch
}

// this is a pointer receiver because we are modifying the value of the receiver
// pointer typed receivers
func (s *SemanticVersion) IncrementMajor() {
	s.major++
}

func (s *SemanticVersion) IncrementMinor() {
	s.minor++
}

func (s *SemanticVersion) IncrementPatch() {
	s.patch++
}

func (s *SemanticVersion) IncrementByMajor(major int) {
	s.major += major
}

func (s *SemanticVersion) IncrementByMinor(minor int) {
	s.minor += minor
}

func (s *SemanticVersion) IncrementByPatch(patch int) {
	s.patch += patch
}
