// Copyright 2019 The Samply Development Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// DeviceUseStatementStatus is documented here http://hl7.org/fhir/ValueSet/device-statement-status
type DeviceUseStatementStatus int

const (
	DeviceUseStatementStatusActive DeviceUseStatementStatus = iota
	DeviceUseStatementStatusCompleted
	DeviceUseStatementStatusEnteredInError
	DeviceUseStatementStatusIntended
	DeviceUseStatementStatusStopped
	DeviceUseStatementStatusOnHold
)

func (code DeviceUseStatementStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DeviceUseStatementStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = DeviceUseStatementStatusActive
	case "completed":
		*code = DeviceUseStatementStatusCompleted
	case "entered-in-error":
		*code = DeviceUseStatementStatusEnteredInError
	case "intended":
		*code = DeviceUseStatementStatusIntended
	case "stopped":
		*code = DeviceUseStatementStatusStopped
	case "on-hold":
		*code = DeviceUseStatementStatusOnHold
	default:
		return fmt.Errorf("unknown DeviceUseStatementStatus code `%s`", s)
	}
	return nil
}
func (code DeviceUseStatementStatus) String() string {
	return code.Code()
}
func (code DeviceUseStatementStatus) Code() string {
	switch code {
	case DeviceUseStatementStatusActive:
		return "active"
	case DeviceUseStatementStatusCompleted:
		return "completed"
	case DeviceUseStatementStatusEnteredInError:
		return "entered-in-error"
	case DeviceUseStatementStatusIntended:
		return "intended"
	case DeviceUseStatementStatusStopped:
		return "stopped"
	case DeviceUseStatementStatusOnHold:
		return "on-hold"
	}
	return "<unknown>"
}
func (code DeviceUseStatementStatus) Display() string {
	switch code {
	case DeviceUseStatementStatusActive:
		return "Active"
	case DeviceUseStatementStatusCompleted:
		return "Completed"
	case DeviceUseStatementStatusEnteredInError:
		return "Entered in Error"
	case DeviceUseStatementStatusIntended:
		return "Intended"
	case DeviceUseStatementStatusStopped:
		return "Stopped"
	case DeviceUseStatementStatusOnHold:
		return "On Hold"
	}
	return "<unknown>"
}
func (code DeviceUseStatementStatus) Definition() string {
	switch code {
	case DeviceUseStatementStatusActive:
		return "The device is still being used."
	case DeviceUseStatementStatusCompleted:
		return "The device is no longer being used."
	case DeviceUseStatementStatusEnteredInError:
		return "The statement was recorded incorrectly."
	case DeviceUseStatementStatusIntended:
		return "The device may be used at some time in the future."
	case DeviceUseStatementStatusStopped:
		return "Actions implied by the statement have been permanently halted, before all of them occurred."
	case DeviceUseStatementStatusOnHold:
		return "Actions implied by the statement have been temporarily halted, but are expected to continue later. May also be called \"suspended\"."
	}
	return "<unknown>"
}
