// Package leetcode provides details for the Leetcode applet.
package leetcode

import (
	_ "embed"

	"tidbyt.dev/community/apps/manifest"
)

//go:embed leetcode.star
var source []byte

// New creates a new instance of the Leetcode applet.
func New() manifest.Manifest {
	return manifest.Manifest{
		ID:          "leetcode",
		Name:        "Leetcode",
		Author:      "greek1414",
		Summary:     "Programming challenges",
		Desc:        "Displays the daily Leetcode challenge.",
		FileName:    "leetcode.star",
		PackageName: "leetcode",
		Source:  source,
	}
}
