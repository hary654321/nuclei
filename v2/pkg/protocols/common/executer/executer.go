package executer

import (
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/pkg/errors"

	"github.com/hary654321/nuclei/v2/pkg/operators/common/dsl"
	"github.com/hary654321/nuclei/v2/pkg/output"
	"github.com/hary654321/nuclei/v2/pkg/protocols"
	"github.com/hary654321/nuclei/v2/pkg/protocols/common/contextargs"
	"github.com/hary654321/nuclei/v2/pkg/protocols/common/helpers/writer"
	"github.com/projectdiscovery/gologger"
)

// Executer executes a group of requests for a protocol
type Executer struct {
	requests []protocols.Request
	options  *protocols.ExecuterOptions
}

var _ protocols.Executer = &Executer{}

// NewExecuter creates a new request executer for list of requests
func NewExecuter(requests []protocols.Request, options *protocols.ExecuterOptions) *Executer {
	return &Executer{requests: requests, options: options}
}

// Compile compiles the execution generators preparing any requests possible.
func (e *Executer) Compile() error {
	cliOptions := e.options.Options

	for _, request := range e.requests {
		if err := request.Compile(e.options); err != nil {
			var dslCompilationError *dsl.CompilationError
			if errors.As(err, &dslCompilationError) {
				if cliOptions.Verbose {
					rawErrorMessage := dslCompilationError.Error()
					formattedErrorMessage := strings.ToUpper(rawErrorMessage[:1]) + rawErrorMessage[1:] + "."
					gologger.Warning().Msgf(formattedErrorMessage)
					gologger.Info().Msgf("The available custom DSL functions are:")
					fmt.Println(dsl.GetPrintableDslFunctionSignatures(cliOptions.NoColor))
				}
			}
			return err
		}
	}
	return nil
}

// Requests returns the total number of requests the rule will perform
func (e *Executer) Requests() int {
	var count int
	for _, request := range e.requests {
		count += request.Requests()
	}
	return count
}

// Execute executes the protocol group and returns true or false if results were found.
func (e *Executer) Execute(executerOpts protocols.ExecuterOptions, input *contextargs.Context) (bool, error) {
	results := &atomic.Bool{}
	// fmt.Printf("%s", debug.Stack())
	dynamicValues := make(map[string]interface{})
	if input.HasArgs() {
		input.ForEach(func(key string, value interface{}) {
			dynamicValues[key] = value
		})
	}
	previous := make(map[string]interface{})
	for _, req := range e.requests {
		inputItem := input.Clone()
		if e.options.InputHelper != nil && input.MetaInput.Input != "" {
			if inputItem.MetaInput.Input = e.options.InputHelper.Transform(inputItem.MetaInput.Input, req.Type()); inputItem.MetaInput.Input == "" {
				return false, nil
			}
		}

		err := req.ExecuteWithResults(inputItem, dynamicValues, previous, func(event *output.InternalWrappedEvent) {
			ID := req.GetID()
			if ID != "" {
				builder := &strings.Builder{}
				for k, v := range event.InternalEvent {
					builder.WriteString(ID)
					builder.WriteString("_")
					builder.WriteString(k)
					previous[builder.String()] = v
					builder.Reset()
				}
			}
			// If no results were found, and also interactsh is not being used
			// in that case we can skip it, otherwise we've to show failure in
			// case of matcher-status flag.
			if !event.HasOperatorResult() && !event.UsesInteractsh {
				if err := e.options.Output.WriteFailure(event.InternalEvent); err != nil {
					gologger.Warning().Msgf("Could not write failure event to output: %s\n", err)
				}
			} else {
				// slog.Println(slog.DEBUG, "c.options.Output", e.options.Output)
				// slog.Println(slog.DEBUG, "executerOpts.Output", executerOpts.Output)
				if writer.WriteResult(event, executerOpts.Output, e.options.Progress, e.options.IssuesClient) {
					results.CompareAndSwap(false, true)
				} else {
					if err := e.options.Output.WriteFailure(event.InternalEvent); err != nil {
						gologger.Warning().Msgf("Could not write failure event to output: %s\n", err)
					}
				}
			}
		})
		if err != nil {
			if e.options.HostErrorsCache != nil {
				e.options.HostErrorsCache.MarkFailed(input.MetaInput.ID(), err)
			}
			gologger.Warning().Msgf("[%s] Could not execute request for %s: %s\n", e.options.TemplateID, input.MetaInput.PrettyPrint(), err)
		}
		// If a match was found and stop at first match is set, break out of the loop and return
		if results.Load() && (e.options.StopAtFirstMatch || e.options.Options.StopAtFirstMatch) {
			break
		}
	}
	return results.Load(), nil
}

// ExecuteWithResults executes the protocol requests and returns results instead of writing them.
func (e *Executer) ExecuteWithResults(input *contextargs.Context, callback protocols.OutputEventCallback) error {
	dynamicValues := make(map[string]interface{})
	if input.HasArgs() {
		input.ForEach(func(key string, value interface{}) {
			dynamicValues[key] = value
		})
	}
	previous := make(map[string]interface{})
	results := &atomic.Bool{}

	for _, req := range e.requests {
		req := req

		inputItem := input.Clone()
		if e.options.InputHelper != nil && input.MetaInput.Input != "" {
			if inputItem.MetaInput.Input = e.options.InputHelper.Transform(input.MetaInput.Input, req.Type()); input.MetaInput.Input == "" {
				return nil
			}
		}

		err := req.ExecuteWithResults(inputItem, dynamicValues, previous, func(event *output.InternalWrappedEvent) {
			ID := req.GetID()
			if ID != "" {
				builder := &strings.Builder{}
				for k, v := range event.InternalEvent {
					builder.WriteString(ID)
					builder.WriteString("_")
					builder.WriteString(k)
					previous[builder.String()] = v
					builder.Reset()
				}
			}
			if event.OperatorsResult == nil {
				return
			}
			results.CompareAndSwap(false, true)
			callback(event)
		})
		if err != nil {
			if e.options.HostErrorsCache != nil {
				e.options.HostErrorsCache.MarkFailed(input.MetaInput.ID(), err)
			}
			gologger.Warning().Msgf("[%s] Could not execute request for %s: %s\n", e.options.TemplateID, input.MetaInput.PrettyPrint(), err)
		}
		// If a match was found and stop at first match is set, break out of the loop and return
		if results.Load() && (e.options.StopAtFirstMatch || e.options.Options.StopAtFirstMatch) {
			break
		}
	}
	return nil
}
