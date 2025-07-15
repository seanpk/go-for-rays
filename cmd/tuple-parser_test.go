package cmd

import (
	"testing"
)

func TestParseTupleDefault(t *testing.T) {
	tuple, err := parseTuple("(1,2,3)")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if tuple.X() != 1.0 || tuple.Y() != 2.0 || tuple.Z() != 3.0 || !tuple.IsVector() {
		t.Errorf("expected Vector(1,2,3), got %s", tuple.String())
	}
}

func TestParse2DTuple(t *testing.T) {
	tuple, err := parseTuple("(1,2)", parseTupleOptions{dimensions: 2})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if tuple.X() != 1.0 || tuple.Y() != 2.0 || tuple.Z() != 0.0 || !tuple.IsVector() {
		t.Errorf("expected Vector(1,2,0), got %s", tuple.String())
	}
}

func TestParsePoint(t *testing.T) {
	tuple, err := parseTuple("(1,2,3)", parseTupleOptions{kind: "Point"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if tuple.X() != 1.0 || tuple.Y() != 2.0 || tuple.Z() != 3.0 || !tuple.IsPoint() {
		t.Errorf("expected Point(1,2,3), got %s", tuple.String())
	}
}
