package main

import (
	"os"

	"github.com/hary654321/nuclei/v2/pkg/testutils"
	errorutil "github.com/projectdiscovery/utils/errors"
)

var templatesDirTestCases = map[string]testutils.TestCase{
	"dns/cname-fingerprint.yaml": &templateDirWithTargetTest{},
}

type templateDirWithTargetTest struct{}

// Execute executes a test case and returns an error if occurred
func (h *templateDirWithTargetTest) Execute(filePath string) error {
	tempdir, err := os.MkdirTemp("", "nuclei-update-dir-*")
	if err != nil {
		return errorutil.NewWithErr(err).Msgf("failed to create temp dir")
	}
	defer os.RemoveAll(tempdir)

	results, err := testutils.RunNucleiTemplateAndGetResults(filePath, "8x8exch02.8x8.com", debug, "-ud", tempdir)
	if err != nil {
		return err
	}

	return expectResultsCount(results, 1)
}
