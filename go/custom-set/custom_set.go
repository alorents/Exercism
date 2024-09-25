package stringset

import (
	"fmt"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set map[string]string

func New() Set {
	return make(map[string]string)
}

func NewFromSlice(l []string) Set {
	set := New()
	for _, element := range l {
		set.Add(element)
	}
	return set
}

func (s Set) String() string {
	elementString := ""
	keys := make([]string, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	if len(keys) == 1 {
		return fmt.Sprintf("{\"%s\"}", keys[0])
	}
	for i, key := range keys {
		if i == 0 {
			elementString += fmt.Sprintf("\"%s\"", key)
		} else {
			elementString += fmt.Sprintf(", \"%s\"", key)
		}
	}

	return fmt.Sprintf("{%s}", elementString)
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s[elem]
	return ok
}

func (s Set) Add(elem string) {
	s[elem] = elem
}

func Subset(s1, s2 Set) bool {
	if len(s1) > len(s2) {
		return false
	}
	for elem := range s1 {
		if !s2.Has(elem) {
			return false
		}
	}

	return true
}

func Disjoint(s1, s2 Set) bool {
	for elem := range s1 {
		if _, ok := s2[elem]; ok {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	return Subset(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	s3 := New()
	for elem := range s1 {
		if s2.Has(elem) {
			s3.Add(elem)
		}
	}

	return s3
}

func Difference(s1, s2 Set) Set {
	if s2.IsEmpty() {
		return s1
	}
	s3 := New()
	for elem := range s1 {
		if !s2.Has(elem) {
			s3.Add(elem)
		}
	}
	return s3
}

func Union(s1, s2 Set) Set {
	s3 := New()
	for elem := range s1 {
		s3.Add(elem)
	}
	for elem := range s2 {
		s3.Add(elem)
	}
	return s3
}
