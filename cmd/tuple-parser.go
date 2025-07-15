package cmd

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/seanpk/go-for-rays/internal/geometry"
)

// parseTuple parses a string representation of a tuple (point or vector) in the format "(x,y[,z])".
// It supports both 2D and 3D tuples, and can return either a Point or a Vector based on the options provided.
// The default behavior is to parse a 3D Vector.
// If the input is invalid, it returns a NaN tuple (IsNaN()==true) and an error.
func parseTuple(input string, options ...parseTupleOptions) (geometry.HomogeneousTuple, error) {
	opts := resolveTupleOptions(options...)

	components, err := extractComponentsFromTupleText(input, opts.dimensions)
	if err != nil {
		return geometry.NaNTuple(), err
	}

	x, err := strconv.ParseFloat(components[0], 64)
	if err != nil {
		if opts.dimensions == 2 {
			return geometry.NewVector(math.NaN(), math.NaN(), 0), err
		}
		return geometry.NewVector(math.NaN(), math.NaN(), math.NaN()), err
	}

	y, err := strconv.ParseFloat(components[1], 64)
	if err != nil {
		if opts.dimensions == 2 {
			return geometry.NewVector(x, math.NaN(), 0), err
		}
		return geometry.NewVector(x, math.NaN(), math.NaN()), err
	}

	var z float64
	if opts.dimensions == 3 {
		z, err = strconv.ParseFloat(components[2], 64)
		if err != nil {
			return geometry.NewVector(x, y, math.NaN()), err
		}
	} else {
		z = 0
	}

	switch opts.kind {
	case "Point":
		return geometry.NewPoint(x, y, z), nil
	default: // "Vector"
		return geometry.NewVector(x, y, z), nil
	}
}

type parseTupleOptions struct {
	dimensions int    // valid values: 2 or 3 (default: 3)
	kind       string // valid values: "Point" or "Vector" (default: "Vector")
}

func resolveTupleOptions(options ...parseTupleOptions) parseTupleOptions {
	if len(options) == 0 {
		return parseTupleOptions{
			dimensions: 3,
			kind:       "Vector",
		}
	}
	opts := options[0]
	if opts.dimensions != 2 && opts.dimensions != 3 {
		opts.dimensions = 3
	}
	if opts.kind != "Point" && opts.kind != "Vector" {
		opts.kind = "Vector"
	}
	return opts
}

var tupleMatcher = regexp.MustCompile(`^\s*\(?\s*([-+]?\d*\.?\d+)\s*,\s*([-+]?\d*\.?\d+)(?:\s*,\s*([-+]?\d*\.?\d+))?\s*\)?\s*$`)

func extractComponentsFromTupleText(input string, expectedDimensions int) ([]string, error) {
	parts := tupleMatcher.FindStringSubmatch(input)
	if parts == nil {
		return nil, fmt.Errorf("invalid input: expected format '(x,y[,z])'")
	}
	if (expectedDimensions == 2 && parts[3] != "") || (expectedDimensions == 3 && parts[3] == "") {
		return nil, fmt.Errorf("invalid input: expected %d dimensions", expectedDimensions)
	}
	return parts[1:], nil
}
