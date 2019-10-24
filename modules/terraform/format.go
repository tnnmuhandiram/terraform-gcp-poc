package terraform

import (
	"fmt"
	"reflect"
	"strings"
)

func FormatArgs(options *Options, args ...string) []string {
	var terraformArgs []string
	terraformArgs = append(terraformArgs, args...)
	terraformArgs = append(terraformArgs, FormatTerraformVarsAsArgs(options.Vars)...)
	terraformArgs = append(terraformArgs, FormatTerraformArgs("-var-file", options.VarFiles)...)
	terraformArgs = append(terraformArgs, FormatTerraformArgs("-target", options.Targets)...)
	return terraformArgs
}

func FormatTerraformVarsAsArgs(vars map[string]interface{}) []string {
	return formatTerraformArgs(vars, "-var", true)
}

func FormatTerraformArgs(argName string, args []string) []string {
	argsList := []string{}
	for _, argValue := range args {
		argsList = append(argsList, argName, argValue)
	}
	return argsList
}

func FormatTerraformBackendConfigAsArgs(vars map[string]interface{}) []string {
	return formatTerraformArgs(vars, "-backend-config", false)
}

func formatTerraformArgs(vars map[string]interface{}, prefix string, useSpaceAsSeparator bool) []string {
	var args []string

	for key, value := range vars {
		hclString := toHclString(value, false)
		argValue := fmt.Sprintf("%s=%s", key, hclString)
		if useSpaceAsSeparator {
			args = append(args, prefix, argValue)
		} else {
			args = append(args, fmt.Sprintf("%s=%s", prefix, argValue))
		}
	}

	return args
}

func toHclString(value interface{}, isNested bool) string {

	if slice, isSlice := tryToConvertToGenericSlice(value); isSlice {
		return sliceToHclString(slice)
	} else if m, isMap := tryToConvertToGenericMap(value); isMap {
		return mapToHclString(m)
	} else {
		return primitiveToHclString(value, isNested)
	}
}

func tryToConvertToGenericSlice(value interface{}) ([]interface{}, bool) {
	reflectValue := reflect.ValueOf(value)
	if reflectValue.Kind() != reflect.Slice {
		return []interface{}{}, false
	}

	genericSlice := make([]interface{}, reflectValue.Len())

	for i := 0; i < reflectValue.Len(); i++ {
		genericSlice[i] = reflectValue.Index(i).Interface()
	}

	return genericSlice, true
}

func tryToConvertToGenericMap(value interface{}) (map[string]interface{}, bool) {
	reflectValue := reflect.ValueOf(value)
	if reflectValue.Kind() != reflect.Map {
		return map[string]interface{}{}, false
	}

	reflectType := reflect.TypeOf(value)
	if reflectType.Key().Kind() != reflect.String {
		return map[string]interface{}{}, false
	}

	genericMap := make(map[string]interface{}, reflectValue.Len())

	mapKeys := reflectValue.MapKeys()
	for _, key := range mapKeys {
		genericMap[key.String()] = reflectValue.MapIndex(key).Interface()
	}

	return genericMap, true
}

func sliceToHclString(slice []interface{}) string {
	hclValues := []string{}

	for _, value := range slice {
		hclValue := toHclString(value, true)
		hclValues = append(hclValues, hclValue)
	}

	return fmt.Sprintf("[%s]", strings.Join(hclValues, ", "))
}

func mapToHclString(m map[string]interface{}) string {
	keyValuePairs := []string{}

	for key, value := range m {
		keyValuePair := fmt.Sprintf(`"%s" = %s`, key, toHclString(value, true))
		keyValuePairs = append(keyValuePairs, keyValuePair)
	}

	return fmt.Sprintf("{%s}", strings.Join(keyValuePairs, ", "))
}

func primitiveToHclString(value interface{}, isNested bool) string {
	switch v := value.(type) {

	case bool:
		if v {
			return "1"
		}
		return "0"

	case string:
		if isNested {
			return fmt.Sprintf("\"%v\"", v)
		}

		return fmt.Sprintf("%v", v)

	default:
		return fmt.Sprintf("%v", v)
	}
}
