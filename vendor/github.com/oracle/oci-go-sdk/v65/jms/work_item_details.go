// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WorkItemDetails The minimum details of a work item.
type WorkItemDetails interface {

	// The work item type.
	GetWorkItemType() WorkItemTypeEnum
}

type workitemdetails struct {
	JsonData     []byte
	WorkItemType WorkItemTypeEnum `mandatory:"false" json:"workItemType,omitempty"`
	Kind         string           `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *workitemdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerworkitemdetails workitemdetails
	s := struct {
		Model Unmarshalerworkitemdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.WorkItemType = s.Model.WorkItemType
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *workitemdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "BASIC":
		mm := BasicWorkItemDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APPLICATION":
		mm := ApplicationWorkItemDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for WorkItemDetails: %s.", m.Kind)
		return *m, nil
	}
}

// GetWorkItemType returns WorkItemType
func (m workitemdetails) GetWorkItemType() WorkItemTypeEnum {
	return m.WorkItemType
}

func (m workitemdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m workitemdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingWorkItemTypeEnum(string(m.WorkItemType)); !ok && m.WorkItemType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkItemType: %s. Supported values are: %s.", m.WorkItemType, strings.Join(GetWorkItemTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WorkItemDetailsKindEnum Enum with underlying type: string
type WorkItemDetailsKindEnum string

// Set of constants representing the allowable values for WorkItemDetailsKindEnum
const (
	WorkItemDetailsKindBasic       WorkItemDetailsKindEnum = "BASIC"
	WorkItemDetailsKindApplication WorkItemDetailsKindEnum = "APPLICATION"
)

var mappingWorkItemDetailsKindEnum = map[string]WorkItemDetailsKindEnum{
	"BASIC":       WorkItemDetailsKindBasic,
	"APPLICATION": WorkItemDetailsKindApplication,
}

var mappingWorkItemDetailsKindEnumLowerCase = map[string]WorkItemDetailsKindEnum{
	"basic":       WorkItemDetailsKindBasic,
	"application": WorkItemDetailsKindApplication,
}

// GetWorkItemDetailsKindEnumValues Enumerates the set of values for WorkItemDetailsKindEnum
func GetWorkItemDetailsKindEnumValues() []WorkItemDetailsKindEnum {
	values := make([]WorkItemDetailsKindEnum, 0)
	for _, v := range mappingWorkItemDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkItemDetailsKindEnumStringValues Enumerates the set of values in String for WorkItemDetailsKindEnum
func GetWorkItemDetailsKindEnumStringValues() []string {
	return []string{
		"BASIC",
		"APPLICATION",
	}
}

// GetMappingWorkItemDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkItemDetailsKindEnum(val string) (WorkItemDetailsKindEnum, bool) {
	enum, ok := mappingWorkItemDetailsKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
