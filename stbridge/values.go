package stbridge

import (
	"fmt"
	"strconv"
	"strings"
)

func float64Attr(node *Node, name AttributeName) (float64, bool, error) {
	raw, ok := node.Attr(name)
	if !ok {
		return 0, false, nil
	}

	value, err := strconv.ParseFloat(strings.TrimSpace(raw), 64)
	if err != nil {
		return 0, true, invalidAttrValueError(name, raw, "float64", err)
	}
	return value, true, nil
}

func invalidAttrValueError(name AttributeName, raw string, target string, err error) error {
	return fmt.Errorf("stbridge: parse attribute %q value %q as %s: %w", name, raw, target, err)
}
