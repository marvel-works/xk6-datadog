/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// LogsURLParserType Type of logs URL parser.
type LogsURLParserType string

// List of LogsURLParserType
const (
	LOGSURLPARSERTYPE_URL_PARSER LogsURLParserType = "url-parser"
)

func (v *LogsURLParserType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogsURLParserType(value)
	for _, existing := range []LogsURLParserType{"url-parser"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogsURLParserType", value)
}

// Ptr returns reference to LogsURLParserType value
func (v LogsURLParserType) Ptr() *LogsURLParserType {
	return &v
}

type NullableLogsURLParserType struct {
	value *LogsURLParserType
	isSet bool
}

func (v NullableLogsURLParserType) Get() *LogsURLParserType {
	return v.value
}

func (v *NullableLogsURLParserType) Set(val *LogsURLParserType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsURLParserType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsURLParserType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsURLParserType(val *LogsURLParserType) *NullableLogsURLParserType {
	return &NullableLogsURLParserType{value: val, isSet: true}
}

func (v NullableLogsURLParserType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsURLParserType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
