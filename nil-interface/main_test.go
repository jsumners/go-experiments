package main

import (
	"reflect"
	"testing"
)

type foo interface {
	value() string
}

type bar struct {
	name string
}

func (b bar) value() string {
	return b.name
}

func reflection(item foo) bool {
	if reflect.ValueOf(item).IsNil() == false {
		return false
	}
	return true
}

func typeAssertion(item foo) bool {
	switch item.(type) {
	case nil:
		return true
	case bar:
		return false
	}
	return false
}

func Benchmark_reflection(b *testing.B) {
	for _ = range b.N {
		reflection(&bar{})
	}
}

func Benchmark_assertion(b *testing.B) {
	for _ = range b.N {
		typeAssertion(&bar{})
	}
}
