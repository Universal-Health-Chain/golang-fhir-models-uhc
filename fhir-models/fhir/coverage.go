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

// Coverage is documented here http://hl7.org/fhir/StructureDefinition/Coverage
type Coverage struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                      `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                   `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            FinancialResourceStatusCodes `bson:"status" json:"status"`
	Type              *CodeableConcept             `bson:"type,omitempty" json:"type,omitempty"`
	PolicyHolder      *Reference                   `bson:"policyHolder,omitempty" json:"policyHolder,omitempty"`
	Subscriber        *Reference                   `bson:"subscriber,omitempty" json:"subscriber,omitempty"`
	SubscriberId      *string                      `bson:"subscriberId,omitempty" json:"subscriberId,omitempty"`
	Beneficiary       Reference                    `bson:"beneficiary" json:"beneficiary"`
	Dependent         *string                      `bson:"dependent,omitempty" json:"dependent,omitempty"`
	Relationship      *CodeableConcept             `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Period            *Period                      `bson:"period,omitempty" json:"period,omitempty"`
	Payor             []Reference                  `bson:"payor" json:"payor"`
	Class             []CoverageClass              `bson:"class,omitempty" json:"class,omitempty"`
	Order             *int                         `bson:"order,omitempty" json:"order,omitempty"`
	Network           *string                      `bson:"network,omitempty" json:"network,omitempty"`
	CostToBeneficiary []CoverageCostToBeneficiary  `bson:"costToBeneficiary,omitempty" json:"costToBeneficiary,omitempty"`
	Subrogation       *bool                        `bson:"subrogation,omitempty" json:"subrogation,omitempty"`
	Contract          []Reference                  `bson:"contract,omitempty" json:"contract,omitempty"`
}
type CoverageClass struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept `bson:"type" json:"type"`
	Value             string          `bson:"value" json:"value"`
	Name              *string         `bson:"name,omitempty" json:"name,omitempty"`
}
type CoverageCostToBeneficiary struct {
	Id                *string                              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept                     `bson:"type,omitempty" json:"type,omitempty"`
	Exception         []CoverageCostToBeneficiaryException `bson:"exception,omitempty" json:"exception,omitempty"`
}
type CoverageCostToBeneficiaryException struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept `bson:"type" json:"type"`
	Period            *Period         `bson:"period,omitempty" json:"period,omitempty"`
}
type OtherCoverage Coverage

// MarshalJSON marshals the given Coverage as JSON into a byte slice
func (r Coverage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCoverage
		ResourceType string `json:"resourceType"`
	}{
		OtherCoverage: OtherCoverage(r),
		ResourceType:  "Coverage",
	})
}

// UnmarshalCoverage unmarshals a Coverage.
func UnmarshalCoverage(b []byte) (Coverage, error) {
	var coverage Coverage
	if err := json.Unmarshal(b, &coverage); err != nil {
		return coverage, err
	}
	return coverage, nil
}
