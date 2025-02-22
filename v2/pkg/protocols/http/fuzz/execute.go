package fuzz

import (
	"regexp"
	"strings"

	"github.com/hary654321/nuclei/v2/pkg/protocols"
	"github.com/hary654321/nuclei/v2/pkg/protocols/common/generators"
	"github.com/pkg/errors"
	"github.com/projectdiscovery/retryablehttp-go"
	urlutil "github.com/projectdiscovery/utils/url"
)

// ExecuteRuleInput is the input for rule Execute function
type ExecuteRuleInput struct {
	// URL is the URL for the request
	URL *urlutil.URL
	// Callback is the callback for generated rule requests
	Callback func(GeneratedRequest) bool
	// InteractURLs contains interact urls for execute call
	InteractURLs []string
	// Values contains dynamic values for the rule
	Values map[string]interface{}
	// BaseRequest is the base http request for fuzzing rule
	BaseRequest *retryablehttp.Request
}

// GeneratedRequest is a single generated request for rule
type GeneratedRequest struct {
	// Request is the http request for rule
	Request *retryablehttp.Request
	// InteractURLs is the list of interactsh urls
	InteractURLs []string
	// DynamicValues contains dynamic values map
	DynamicValues map[string]interface{}
}

// Execute executes a fuzzing rule accepting a callback on which
// generated requests are returned.
//
// Input is not thread safe and should not be shared between concurrent
// goroutines.
func (rule *Rule) Execute(input *ExecuteRuleInput) error {
	if !rule.isExecutable(input.URL) {
		return nil
	}
	baseValues := input.Values
	if rule.generator == nil {
		evaluatedValues, interactURLs := rule.options.Variables.EvaluateWithInteractsh(baseValues, rule.options.Interactsh)
		input.Values = generators.MergeMaps(evaluatedValues, baseValues)
		input.InteractURLs = interactURLs
		err := rule.executeRuleValues(input)
		return err
	}
	iterator := rule.generator.NewIterator()
	for {
		values, next := iterator.Value()
		if !next {
			return nil
		}
		evaluatedValues, interactURLs := rule.options.Variables.EvaluateWithInteractsh(generators.MergeMaps(values, baseValues), rule.options.Interactsh)
		input.InteractURLs = interactURLs
		input.Values = generators.MergeMaps(values, evaluatedValues, baseValues)

		if err := rule.executeRuleValues(input); err != nil {
			return err
		}
	}
}

// isExecutable returns true if the rule can be executed based on provided input
func (rule *Rule) isExecutable(parsed *urlutil.URL) bool {
	if len(parsed.Query()) > 0 && rule.partType == queryPartType {
		return true
	}
	return false
}

// executeRuleValues executes a rule with a set of values
func (rule *Rule) executeRuleValues(input *ExecuteRuleInput) error {
	for _, payload := range rule.Fuzz {
		if err := rule.executePartRule(input, payload); err != nil {
			return err
		}
	}
	return nil
}

// Compile compiles a fuzzing rule and initializes it for operation
func (rule *Rule) Compile(generator *generators.PayloadGenerator, options *protocols.ExecuterOptions) error {
	// If a payload generator is specified from base request, use it
	// for payload values.
	if generator != nil {
		rule.generator = generator
	}
	rule.options = options

	// Resolve the default enums
	if rule.Mode != "" {
		if valueType, ok := stringToModeType[rule.Mode]; !ok {
			return errors.Errorf("invalid mode value specified: %s", rule.Mode)
		} else {
			rule.modeType = valueType
		}
	} else {
		rule.modeType = multipleModeType
	}
	if rule.Part != "" {
		if valueType, ok := stringToPartType[rule.Part]; !ok {
			return errors.Errorf("invalid part value specified: %s", rule.Part)
		} else {
			rule.partType = valueType
		}
	} else {
		rule.partType = queryPartType
	}

	if rule.Type != "" {
		if valueType, ok := stringToRuleType[rule.Type]; !ok {
			return errors.Errorf("invalid type value specified: %s", rule.Type)
		} else {
			rule.ruleType = valueType
		}
	} else {
		rule.ruleType = replaceRuleType
	}

	// Initialize other required regexes and maps
	if len(rule.Keys) > 0 {
		rule.keysMap = make(map[string]struct{})
	}
	for _, key := range rule.Keys {
		rule.keysMap[strings.ToLower(key)] = struct{}{}
	}
	for _, value := range rule.ValuesRegex {
		compiled, err := regexp.Compile(value)
		if err != nil {
			return errors.Wrap(err, "could not compile value regex")
		}
		rule.valuesRegex = append(rule.valuesRegex, compiled)
	}
	for _, value := range rule.KeysRegex {
		compiled, err := regexp.Compile(value)
		if err != nil {
			return errors.Wrap(err, "could not compile key regex")
		}
		rule.keysRegex = append(rule.keysRegex, compiled)
	}
	return nil
}
