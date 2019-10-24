package terraform

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Output(t *testing.T, options *Options, key string) string {
	out, err := OutputE(t, options, key)
	require.NoError(t, err)
	return out
}

func OutputE(t *testing.T, options *Options, key string) (string, error) {
	output, err := RunTerraformCommandAndGetStdoutE(t, options, "output", "-no-color", key)

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(output), nil
}

func OutputRequired(t *testing.T, options *Options, key string) string {
	out, err := OutputRequiredE(t, options, key)
	require.NoError(t, err)
	return out
}

func OutputRequiredE(t *testing.T, options *Options, key string) (string, error) {
	out, err := OutputE(t, options, key)

	if err != nil {
		return "", err
	}
	if out == "" {
		return "", EmptyOutput(key)
	}

	return out, nil
}

func OutputList(t *testing.T, options *Options, key string) []string {
	out, err := OutputListE(t, options, key)
	require.NoError(t, err)
	return out
}

func OutputListE(t *testing.T, options *Options, key string) ([]string, error) {
	out, err := RunTerraformCommandAndGetStdoutE(t, options, "output", "-no-color", "-json", key)
	if err != nil {
		return nil, err
	}

	var output interface{}
	if err := json.Unmarshal([]byte(out), &output); err != nil {
		return nil, err
	}

	if outputMap, isMap := output.(map[string]interface{}); isMap {
		return parseListOutputTerraform11OrOlder(outputMap, key)
	} else if outputList, isList := output.([]interface{}); isList {
		return parseListOutputTerraform12OrNewer(outputList, key)
	}

	return nil, UnexpectedOutputType{Key: key, ExpectedType: "map or list", ActualType: reflect.TypeOf(output).String()}
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

func OutputMap(t *testing.T, options *Options, key string) map[string]string {
	out, err := OutputMapE(t, options, key)
	require.NoError(t, err)
	return out
}

func OutputMapE(t *testing.T, options *Options, key string) (map[string]string, error) {
	out, err := RunTerraformCommandAndGetStdoutE(t, options, "output", "-no-color", "-json", key)
	if err != nil {
		return nil, err
	}

	outputMap := map[string]interface{}{}
	if err := json.Unmarshal([]byte(out), &outputMap); err != nil {
		return nil, err
	}

	value, containsValue := outputMap["value"]
	_, containsSensitive := outputMap["sensitive"]
	_, containsType := outputMap["type"]
	if containsValue && containsSensitive && containsType {
		valueMap, ok := value.(map[string]interface{})
		if !ok {
			return nil, OutputValueNotMap{Value: value}
		}

		outputMap = valueMap
	}

	resultMap := make(map[string]string)
	for k, v := range outputMap {
		resultMap[k] = fmt.Sprintf("%v", v)
	}
	return resultMap, nil
}

func OutputForKeys(t *testing.T, options *Options, keys []string) map[string]interface{} {
	out, err := OutputForKeysE(t, options, keys)
	require.NoError(t, err)
	return out
}

func OutputForKeysE(t *testing.T, options *Options, keys []string) (map[string]interface{}, error) {
	out, err := RunTerraformCommandAndGetStdoutE(t, options, "output", "-no-color", "-json")
	if err != nil {
		return nil, err
	}

	outputMap := map[string]map[string]interface{}{}
	if err := json.Unmarshal([]byte(out), &outputMap); err != nil {
		return nil, err
	}

	if keys == nil {
		outputKeys := make([]string, 0, len(outputMap))
		for k := range outputMap {
			outputKeys = append(outputKeys, k)
		}
		keys = outputKeys
	}

	resultMap := make(map[string]interface{})
	for _, key := range keys {
		value, containsValue := outputMap[key]["value"]
		if !containsValue {
			return nil, OutputKeyNotFound(string(key))
		}
		resultMap[key] = value
	}
	return resultMap, nil
}

func OutputAll(t *testing.T, options *Options) map[string]interface{} {
	out, err := OutputAllE(t, options)
	require.NoError(t, err)
	return out
}

func OutputAllE(t *testing.T, options *Options) (map[string]interface{}, error) {
	return OutputForKeysE(t, options, nil)
}
