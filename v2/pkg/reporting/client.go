package reporting

import (
	"github.com/hary654321/nuclei/v2/pkg/output"
)

// Client is a client for nuclei issue tracking module
type Client interface {
	RegisterTracker(tracker Tracker)
	RegisterExporter(exporter Exporter)
	Close()
	Clear()
	CreateIssue(event *output.ResultEvent) error
	GetReportingOptions() *Options
}
