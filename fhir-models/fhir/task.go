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

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// Task is documented here http://hl7.org/fhir/StructureDefinition/Task
type Task struct {
	Id                    *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	InstantiatesCanonical *string           `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       *string           `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	BasedOn               []Reference       `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	GroupIdentifier       *Identifier       `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	PartOf                []Reference       `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                TaskStatus        `bson:"status" json:"status"`
	StatusReason          *CodeableConcept  `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	BusinessStatus        *CodeableConcept  `bson:"businessStatus,omitempty" json:"businessStatus,omitempty"`
	Intent                string            `bson:"intent" json:"intent"`
	Priority              *RequestPriority  `bson:"priority,omitempty" json:"priority,omitempty"`
	Code                  *CodeableConcept  `bson:"code,omitempty" json:"code,omitempty"`
	Description           *string           `bson:"description,omitempty" json:"description,omitempty"`
	Focus                 *Reference        `bson:"focus,omitempty" json:"focus,omitempty"`
	For                   *Reference        `bson:"for,omitempty" json:"for,omitempty"`
	Encounter             *Reference        `bson:"encounter,omitempty" json:"encounter,omitempty"`
	ExecutionPeriod       *Period           `bson:"executionPeriod,omitempty" json:"executionPeriod,omitempty"`
	AuthoredOn            *string           `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	LastModified          *string           `bson:"lastModified,omitempty" json:"lastModified,omitempty"`
	Requester             *Reference        `bson:"requester,omitempty" json:"requester,omitempty"`
	PerformerType         []CodeableConcept `bson:"performerType,omitempty" json:"performerType,omitempty"`
	Owner                 *Reference        `bson:"owner,omitempty" json:"owner,omitempty"`
	Location              *Reference        `bson:"location,omitempty" json:"location,omitempty"`
	ReasonCode            *CodeableConcept  `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference       *Reference        `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Insurance             []Reference       `bson:"insurance,omitempty" json:"insurance,omitempty"`
	Note                  []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
	RelevantHistory       []Reference       `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
	Restriction           *TaskRestriction  `bson:"restriction,omitempty" json:"restriction,omitempty"`
	Input                 []TaskInput       `bson:"input,omitempty" json:"input,omitempty"`
	Output                []TaskOutput      `bson:"output,omitempty" json:"output,omitempty"`
}
type TaskRestriction struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Repetitions       *int        `bson:"repetitions,omitempty" json:"repetitions,omitempty"`
	Period            *Period     `bson:"period,omitempty" json:"period,omitempty"`
	Recipient         []Reference `bson:"recipient,omitempty" json:"recipient,omitempty"`
}
type TaskInput struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept `bson:"type" json:"type"`
}
type TaskOutput struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept `bson:"type" json:"type"`
}
type OtherTask Task

// MarshalJSON marshals the given Task as JSON into a byte slice
func (r Task) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTask
		ResourceType string `json:"resourceType"`
	}{
		OtherTask:    OtherTask(r),
		ResourceType: "Task",
	})
}

// UnmarshalTask unmarshals a Task.
func UnmarshalTask(b []byte) (Task, error) {
	var task Task
	if err := json.Unmarshal(b, &task); err != nil {
		return task, err
	}
	return task, nil
}
