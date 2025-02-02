package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ValidateRQLIdentity represents the Identity of the object.
var ValidateRQLIdentity = elemental.Identity{
	Name:     "validaterql",
	Category: "validaterql",
	Package:  "larl",
	Private:  false,
}

// ValidateRQLsList represents a list of ValidateRQLs
type ValidateRQLsList []*ValidateRQL

// Identity returns the identity of the objects in the list.
func (o ValidateRQLsList) Identity() elemental.Identity {

	return ValidateRQLIdentity
}

// Copy returns a pointer to a copy the ValidateRQLsList.
func (o ValidateRQLsList) Copy() elemental.Identifiables {

	copy := append(ValidateRQLsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ValidateRQLsList.
func (o ValidateRQLsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ValidateRQLsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ValidateRQL))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ValidateRQLsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ValidateRQLsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ValidateRQLsList converted to SparseValidateRQLsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ValidateRQLsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseValidateRQLsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseValidateRQL)
	}

	return out
}

// Version returns the version of the content.
func (o ValidateRQLsList) Version() int {

	return 1
}

// ValidateRQL represents the model of a validaterql
type ValidateRQL struct {
	// The error message if the query fails validation.
	Error string `json:"error" msgpack:"error" bson:"-" mapstructure:"error,omitempty"`

	// The type of the policy.
	PolicyType string `json:"policyType" msgpack:"policyType" bson:"-" mapstructure:"policyType,omitempty"`

	// The Prisma Cloud ID.
	PrismaID int `json:"prismaID" msgpack:"prismaID" bson:"-" mapstructure:"prismaID,omitempty"`

	// The query to validate.
	Query string `json:"query" msgpack:"query" bson:"-" mapstructure:"query,omitempty"`

	// The search type.
	SearchType string `json:"searchType" msgpack:"searchType" bson:"-" mapstructure:"searchType,omitempty"`

	// The http status code of the response.
	Status int `json:"status" msgpack:"status" bson:"-" mapstructure:"status,omitempty"`

	// The timestamp of the query in nanoseconds.
	Timestamp int `json:"timestamp" msgpack:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewValidateRQL returns a new *ValidateRQL
func NewValidateRQL() *ValidateRQL {

	return &ValidateRQL{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *ValidateRQL) Identity() elemental.Identity {

	return ValidateRQLIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ValidateRQL) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ValidateRQL) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ValidateRQL) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesValidateRQL{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ValidateRQL) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesValidateRQL{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *ValidateRQL) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *ValidateRQL) BleveType() string {

	return "validaterql"
}

// DefaultOrder returns the list of default ordering fields.
func (o *ValidateRQL) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *ValidateRQL) Doc() string {

	return `Used to validate RQL queries.`
}

func (o *ValidateRQL) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ValidateRQL) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseValidateRQL{
			Error:      &o.Error,
			PolicyType: &o.PolicyType,
			PrismaID:   &o.PrismaID,
			Query:      &o.Query,
			SearchType: &o.SearchType,
			Status:     &o.Status,
			Timestamp:  &o.Timestamp,
		}
	}

	sp := &SparseValidateRQL{}
	for _, f := range fields {
		switch f {
		case "error":
			sp.Error = &(o.Error)
		case "policyType":
			sp.PolicyType = &(o.PolicyType)
		case "prismaID":
			sp.PrismaID = &(o.PrismaID)
		case "query":
			sp.Query = &(o.Query)
		case "searchType":
			sp.SearchType = &(o.SearchType)
		case "status":
			sp.Status = &(o.Status)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseValidateRQL to the object.
func (o *ValidateRQL) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseValidateRQL)
	if so.Error != nil {
		o.Error = *so.Error
	}
	if so.PolicyType != nil {
		o.PolicyType = *so.PolicyType
	}
	if so.PrismaID != nil {
		o.PrismaID = *so.PrismaID
	}
	if so.Query != nil {
		o.Query = *so.Query
	}
	if so.SearchType != nil {
		o.SearchType = *so.SearchType
	}
	if so.Status != nil {
		o.Status = *so.Status
	}
	if so.Timestamp != nil {
		o.Timestamp = *so.Timestamp
	}
}

// DeepCopy returns a deep copy if the ValidateRQL.
func (o *ValidateRQL) DeepCopy() *ValidateRQL {

	if o == nil {
		return nil
	}

	out := &ValidateRQL{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ValidateRQL.
func (o *ValidateRQL) DeepCopyInto(out *ValidateRQL) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ValidateRQL: %s", err))
	}

	*out = *target.(*ValidateRQL)
}

// Validate valides the current information stored into the structure.
func (o *ValidateRQL) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("query", o.Query); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*ValidateRQL) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ValidateRQLAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ValidateRQLLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ValidateRQL) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ValidateRQLAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ValidateRQL) ValueForAttribute(name string) interface{} {

	switch name {
	case "error":
		return o.Error
	case "policyType":
		return o.PolicyType
	case "prismaID":
		return o.PrismaID
	case "query":
		return o.Query
	case "searchType":
		return o.SearchType
	case "status":
		return o.Status
	case "timestamp":
		return o.Timestamp
	}

	return nil
}

// ValidateRQLAttributesMap represents the map of attribute for ValidateRQL.
var ValidateRQLAttributesMap = map[string]elemental.AttributeSpecification{
	"Error": {
		AllowedChoices: []string{},
		ConvertedName:  "Error",
		Description:    `The error message if the query fails validation.`,
		Exposed:        true,
		Name:           "error",
		Type:           "string",
	},
	"PolicyType": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyType",
		Description:    `The type of the policy.`,
		Exposed:        true,
		Name:           "policyType",
		Type:           "string",
	},
	"PrismaID": {
		AllowedChoices: []string{},
		ConvertedName:  "PrismaID",
		Description:    `The Prisma Cloud ID.`,
		Exposed:        true,
		Name:           "prismaID",
		Type:           "integer",
	},
	"Query": {
		AllowedChoices: []string{},
		ConvertedName:  "Query",
		Description:    `The query to validate.`,
		Exposed:        true,
		Name:           "query",
		Required:       true,
		Type:           "string",
	},
	"SearchType": {
		AllowedChoices: []string{},
		ConvertedName:  "SearchType",
		Description:    `The search type.`,
		Exposed:        true,
		Name:           "searchType",
		Type:           "string",
	},
	"Status": {
		AllowedChoices: []string{},
		ConvertedName:  "Status",
		Description:    `The http status code of the response.`,
		Exposed:        true,
		Name:           "status",
		Type:           "integer",
	},
	"Timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `The timestamp of the query in nanoseconds.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "integer",
	},
}

// ValidateRQLLowerCaseAttributesMap represents the map of attribute for ValidateRQL.
var ValidateRQLLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"error": {
		AllowedChoices: []string{},
		ConvertedName:  "Error",
		Description:    `The error message if the query fails validation.`,
		Exposed:        true,
		Name:           "error",
		Type:           "string",
	},
	"policytype": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyType",
		Description:    `The type of the policy.`,
		Exposed:        true,
		Name:           "policyType",
		Type:           "string",
	},
	"prismaid": {
		AllowedChoices: []string{},
		ConvertedName:  "PrismaID",
		Description:    `The Prisma Cloud ID.`,
		Exposed:        true,
		Name:           "prismaID",
		Type:           "integer",
	},
	"query": {
		AllowedChoices: []string{},
		ConvertedName:  "Query",
		Description:    `The query to validate.`,
		Exposed:        true,
		Name:           "query",
		Required:       true,
		Type:           "string",
	},
	"searchtype": {
		AllowedChoices: []string{},
		ConvertedName:  "SearchType",
		Description:    `The search type.`,
		Exposed:        true,
		Name:           "searchType",
		Type:           "string",
	},
	"status": {
		AllowedChoices: []string{},
		ConvertedName:  "Status",
		Description:    `The http status code of the response.`,
		Exposed:        true,
		Name:           "status",
		Type:           "integer",
	},
	"timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `The timestamp of the query in nanoseconds.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "integer",
	},
}

// SparseValidateRQLsList represents a list of SparseValidateRQLs
type SparseValidateRQLsList []*SparseValidateRQL

// Identity returns the identity of the objects in the list.
func (o SparseValidateRQLsList) Identity() elemental.Identity {

	return ValidateRQLIdentity
}

// Copy returns a pointer to a copy the SparseValidateRQLsList.
func (o SparseValidateRQLsList) Copy() elemental.Identifiables {

	copy := append(SparseValidateRQLsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseValidateRQLsList.
func (o SparseValidateRQLsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseValidateRQLsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseValidateRQL))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseValidateRQLsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseValidateRQLsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseValidateRQLsList converted to ValidateRQLsList.
func (o SparseValidateRQLsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseValidateRQLsList) Version() int {

	return 1
}

// SparseValidateRQL represents the sparse version of a validaterql.
type SparseValidateRQL struct {
	// The error message if the query fails validation.
	Error *string `json:"error,omitempty" msgpack:"error,omitempty" bson:"-" mapstructure:"error,omitempty"`

	// The type of the policy.
	PolicyType *string `json:"policyType,omitempty" msgpack:"policyType,omitempty" bson:"-" mapstructure:"policyType,omitempty"`

	// The Prisma Cloud ID.
	PrismaID *int `json:"prismaID,omitempty" msgpack:"prismaID,omitempty" bson:"-" mapstructure:"prismaID,omitempty"`

	// The query to validate.
	Query *string `json:"query,omitempty" msgpack:"query,omitempty" bson:"-" mapstructure:"query,omitempty"`

	// The search type.
	SearchType *string `json:"searchType,omitempty" msgpack:"searchType,omitempty" bson:"-" mapstructure:"searchType,omitempty"`

	// The http status code of the response.
	Status *int `json:"status,omitempty" msgpack:"status,omitempty" bson:"-" mapstructure:"status,omitempty"`

	// The timestamp of the query in nanoseconds.
	Timestamp *int `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseValidateRQL returns a new  SparseValidateRQL.
func NewSparseValidateRQL() *SparseValidateRQL {
	return &SparseValidateRQL{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseValidateRQL) Identity() elemental.Identity {

	return ValidateRQLIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseValidateRQL) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseValidateRQL) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseValidateRQL) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseValidateRQL{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseValidateRQL) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseValidateRQL{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseValidateRQL) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseValidateRQL) ToPlain() elemental.PlainIdentifiable {

	out := NewValidateRQL()
	if o.Error != nil {
		out.Error = *o.Error
	}
	if o.PolicyType != nil {
		out.PolicyType = *o.PolicyType
	}
	if o.PrismaID != nil {
		out.PrismaID = *o.PrismaID
	}
	if o.Query != nil {
		out.Query = *o.Query
	}
	if o.SearchType != nil {
		out.SearchType = *o.SearchType
	}
	if o.Status != nil {
		out.Status = *o.Status
	}
	if o.Timestamp != nil {
		out.Timestamp = *o.Timestamp
	}

	return out
}

// DeepCopy returns a deep copy if the SparseValidateRQL.
func (o *SparseValidateRQL) DeepCopy() *SparseValidateRQL {

	if o == nil {
		return nil
	}

	out := &SparseValidateRQL{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseValidateRQL.
func (o *SparseValidateRQL) DeepCopyInto(out *SparseValidateRQL) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseValidateRQL: %s", err))
	}

	*out = *target.(*SparseValidateRQL)
}

type mongoAttributesValidateRQL struct {
}
type mongoAttributesSparseValidateRQL struct {
}
