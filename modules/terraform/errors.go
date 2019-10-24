package terraform

import (
	"fmt"
)

type TgInvalidBinary string

func (err TgInvalidBinary) Error() string {
	return fmt.Sprintf("terragrunt must be set as TerraformBinary to use this function. [ TerraformBinary : %s ]", err)
}

type OutputKeyNotFound string

func (err OutputKeyNotFound) Error() string {
	return fmt.Sprintf("output doesn't contain a value for the key %q", err)
}

type OutputValueNotMap struct {
	Value interface{}
}

func (err OutputValueNotMap) Error() string {
	return fmt.Sprintf("Output value %q is not a map", err.Value)
}

type OutputValueNotList struct {
	Value interface{}
}

func (err OutputValueNotList) Error() string {
	return fmt.Sprintf("Output value %q is not a list", err.Value)
}

type EmptyOutput string

func (outputName EmptyOutput) Error() string {
	return fmt.Sprintf("Required output %s was empty", string(outputName))
}

type UnexpectedOutputType struct {
	Key          string
	ExpectedType string
	ActualType   string
}

func (err UnexpectedOutputType) Error() string {
	return fmt.Sprintf("Expected output '%s' to be of type '%s' but got '%s'", err.Key, err.ExpectedType, err.ActualType)
}
