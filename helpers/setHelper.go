package helpers

import (
	"github.com/juliangruber/go-intersect/v2"
	"syncute-go/messages/resources"
)

func Difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func DifferenceResources(a, b []resources.Resource) []resources.Resource {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x.RelativePath] = struct{}{}
	}
	var diff []resources.Resource
	for _, x := range a {
		if _, found := mb[x.RelativePath]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func Intersect(a, b []string) []string {
	res := intersect.SortedGeneric(a, b)
	return res
}

func IntersectResources(a, b []resources.Resource) []resources.Resource {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x.RelativePath] = struct{}{}
	}
	var diff []resources.Resource
	for _, x := range a {
		if _, found := mb[x.RelativePath]; found {
			diff = append(diff, x)
		}
	}
	return diff
}
