package terraform

import (
	"fmt"
	"strings"
)

func Output(options *Options, key string) string {
	out, err := OutputE(options, key)
	// require.NoError(t, err)
	print(err)
	return out
}

func OutputE(options *Options, key string) (string, error) {
	output, err := RunTerraformCommandAndGetStdoutE(options, "output", "-no-color", key)

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(output), nil
}

func parseListOutputTerraform12OrNewer(outputList []interface{}, key string) ([]string, error) {
	list := []string{}

	for _, item := range outputList {
		list = append(list, fmt.Sprintf("%v", item))
	}

	return list, nil
}

func parseListOutputTerraform11OrOlder(outputMap map[string]interface{}, key string) ([]string, error) {
	value, containsValue := outputMap["value"]
	if !containsValue {
		return nil, OutputKeyNotFound(key)
	}

	list := []string{}
	switch t := value.(type) {
	case []interface{}:
		for _, item := range t {
			list = append(list, fmt.Sprintf("%v", item))
		}
	default:
		return nil, OutputValueNotList{Value: value}
	}

	return list, nil
}
