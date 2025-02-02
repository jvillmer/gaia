package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PUTrafficActionIncomingValue represents the possible values for attribute "Incoming".
type PUTrafficActionIncomingValue string

const (
	// PUTrafficActionIncomingAllow represents the value Allow.
	PUTrafficActionIncomingAllow PUTrafficActionIncomingValue = "Allow"

	// PUTrafficActionIncomingInherit represents the value Inherit.
	PUTrafficActionIncomingInherit PUTrafficActionIncomingValue = "Inherit"

	// PUTrafficActionIncomingReject represents the value Reject.
	PUTrafficActionIncomingReject PUTrafficActionIncomingValue = "Reject"
)

// PUTrafficActionOutgoingValue represents the possible values for attribute "Outgoing".
type PUTrafficActionOutgoingValue string

const (
	// PUTrafficActionOutgoingAllow represents the value Allow.
	PUTrafficActionOutgoingAllow PUTrafficActionOutgoingValue = "Allow"

	// PUTrafficActionOutgoingInherit represents the value Inherit.
	PUTrafficActionOutgoingInherit PUTrafficActionOutgoingValue = "Inherit"

	// PUTrafficActionOutgoingReject represents the value Reject.
	PUTrafficActionOutgoingReject PUTrafficActionOutgoingValue = "Reject"
)

// PUTrafficActionIdentity represents the Identity of the object.
var PUTrafficActionIdentity = elemental.Identity{
	Name:     "putrafficaction",
	Category: "putrafficactions",
	Package:  "squall",
	Private:  false,
}

// PUTrafficActionsList represents a list of PUTrafficActions
type PUTrafficActionsList []*PUTrafficAction

// Identity returns the identity of the objects in the list.
func (o PUTrafficActionsList) Identity() elemental.Identity {

	return PUTrafficActionIdentity
}

// Copy returns a pointer to a copy the PUTrafficActionsList.
func (o PUTrafficActionsList) Copy() elemental.Identifiables {

	copy := append(PUTrafficActionsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PUTrafficActionsList.
func (o PUTrafficActionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PUTrafficActionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PUTrafficAction))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PUTrafficActionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PUTrafficActionsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PUTrafficActionsList converted to SparsePUTrafficActionsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PUTrafficActionsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePUTrafficActionsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePUTrafficAction)
	}

	return out
}

// Version returns the version of the content.
func (o PUTrafficActionsList) Version() int {

	return 1
}

// PUTrafficAction represents the model of a putrafficaction
// +k8s:openapi-gen=true
type PUTrafficAction struct {
	// The processing unit action for incoming traffic for the namespace.
	Incoming PUTrafficActionIncomingValue `json:"Incoming" msgpack:"Incoming" bson:"-" mapstructure:"Incoming,omitempty"`

	// The processing unit action for outgoing traffic for the namespace.
	Outgoing PUTrafficActionOutgoingValue `json:"Outgoing" msgpack:"Outgoing" bson:"-" mapstructure:"Outgoing,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPUTrafficAction returns a new *PUTrafficAction
func NewPUTrafficAction() *PUTrafficAction {

	return &PUTrafficAction{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *PUTrafficAction) Identity() elemental.Identity {

	return PUTrafficActionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PUTrafficAction) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PUTrafficAction) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PUTrafficAction) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPUTrafficAction{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PUTrafficAction) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPUTrafficAction{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *PUTrafficAction) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *PUTrafficAction) BleveType() string {

	return "putrafficaction"
}

// DefaultOrder returns the list of default ordering fields.
func (o *PUTrafficAction) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PUTrafficAction) Doc() string {

	return `Returns the processing unit traffic actions for the specified namespace.`
}

func (o *PUTrafficAction) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PUTrafficAction) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePUTrafficAction{
			Incoming: &o.Incoming,
			Outgoing: &o.Outgoing,
		}
	}

	sp := &SparsePUTrafficAction{}
	for _, f := range fields {
		switch f {
		case "Incoming":
			sp.Incoming = &(o.Incoming)
		case "Outgoing":
			sp.Outgoing = &(o.Outgoing)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePUTrafficAction to the object.
func (o *PUTrafficAction) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePUTrafficAction)
	if so.Incoming != nil {
		o.Incoming = *so.Incoming
	}
	if so.Outgoing != nil {
		o.Outgoing = *so.Outgoing
	}
}

// DeepCopy returns a deep copy if the PUTrafficAction.
func (o *PUTrafficAction) DeepCopy() *PUTrafficAction {

	if o == nil {
		return nil
	}

	out := &PUTrafficAction{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PUTrafficAction.
func (o *PUTrafficAction) DeepCopyInto(out *PUTrafficAction) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PUTrafficAction: %s", err))
	}

	*out = *target.(*PUTrafficAction)
}

// Validate valides the current information stored into the structure.
func (o *PUTrafficAction) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("Incoming", string(o.Incoming), []string{"Allow", "Reject", "Inherit"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("Outgoing", string(o.Outgoing), []string{"Allow", "Reject", "Inherit"}, false); err != nil {
		errors = errors.Append(err)
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
func (*PUTrafficAction) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PUTrafficActionAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PUTrafficActionLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PUTrafficAction) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PUTrafficActionAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PUTrafficAction) ValueForAttribute(name string) interface{} {

	switch name {
	case "Incoming":
		return o.Incoming
	case "Outgoing":
		return o.Outgoing
	}

	return nil
}

// PUTrafficActionAttributesMap represents the map of attribute for PUTrafficAction.
var PUTrafficActionAttributesMap = map[string]elemental.AttributeSpecification{
	"Incoming": {
		AllowedChoices: []string{"Allow", "Reject", "Inherit"},
		ConvertedName:  "Incoming",
		Description:    `The processing unit action for incoming traffic for the namespace.`,
		Exposed:        true,
		Name:           "Incoming",
		Type:           "enum",
	},
	"Outgoing": {
		AllowedChoices: []string{"Allow", "Reject", "Inherit"},
		ConvertedName:  "Outgoing",
		Description:    `The processing unit action for outgoing traffic for the namespace.`,
		Exposed:        true,
		Name:           "Outgoing",
		Type:           "enum",
	},
}

// PUTrafficActionLowerCaseAttributesMap represents the map of attribute for PUTrafficAction.
var PUTrafficActionLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"incoming": {
		AllowedChoices: []string{"Allow", "Reject", "Inherit"},
		ConvertedName:  "Incoming",
		Description:    `The processing unit action for incoming traffic for the namespace.`,
		Exposed:        true,
		Name:           "Incoming",
		Type:           "enum",
	},
	"outgoing": {
		AllowedChoices: []string{"Allow", "Reject", "Inherit"},
		ConvertedName:  "Outgoing",
		Description:    `The processing unit action for outgoing traffic for the namespace.`,
		Exposed:        true,
		Name:           "Outgoing",
		Type:           "enum",
	},
}

// SparsePUTrafficActionsList represents a list of SparsePUTrafficActions
type SparsePUTrafficActionsList []*SparsePUTrafficAction

// Identity returns the identity of the objects in the list.
func (o SparsePUTrafficActionsList) Identity() elemental.Identity {

	return PUTrafficActionIdentity
}

// Copy returns a pointer to a copy the SparsePUTrafficActionsList.
func (o SparsePUTrafficActionsList) Copy() elemental.Identifiables {

	copy := append(SparsePUTrafficActionsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePUTrafficActionsList.
func (o SparsePUTrafficActionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePUTrafficActionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePUTrafficAction))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePUTrafficActionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePUTrafficActionsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePUTrafficActionsList converted to PUTrafficActionsList.
func (o SparsePUTrafficActionsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePUTrafficActionsList) Version() int {

	return 1
}

// SparsePUTrafficAction represents the sparse version of a putrafficaction.
// +k8s:openapi-gen=true
type SparsePUTrafficAction struct {
	// The processing unit action for incoming traffic for the namespace.
	Incoming *PUTrafficActionIncomingValue `json:"Incoming,omitempty" msgpack:"Incoming,omitempty" bson:"-" mapstructure:"Incoming,omitempty"`

	// The processing unit action for outgoing traffic for the namespace.
	Outgoing *PUTrafficActionOutgoingValue `json:"Outgoing,omitempty" msgpack:"Outgoing,omitempty" bson:"-" mapstructure:"Outgoing,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePUTrafficAction returns a new  SparsePUTrafficAction.
func NewSparsePUTrafficAction() *SparsePUTrafficAction {
	return &SparsePUTrafficAction{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePUTrafficAction) Identity() elemental.Identity {

	return PUTrafficActionIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePUTrafficAction) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePUTrafficAction) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePUTrafficAction) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePUTrafficAction{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePUTrafficAction) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePUTrafficAction{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparsePUTrafficAction) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePUTrafficAction) ToPlain() elemental.PlainIdentifiable {

	out := NewPUTrafficAction()
	if o.Incoming != nil {
		out.Incoming = *o.Incoming
	}
	if o.Outgoing != nil {
		out.Outgoing = *o.Outgoing
	}

	return out
}

// DeepCopy returns a deep copy if the SparsePUTrafficAction.
func (o *SparsePUTrafficAction) DeepCopy() *SparsePUTrafficAction {

	if o == nil {
		return nil
	}

	out := &SparsePUTrafficAction{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePUTrafficAction.
func (o *SparsePUTrafficAction) DeepCopyInto(out *SparsePUTrafficAction) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePUTrafficAction: %s", err))
	}

	*out = *target.(*SparsePUTrafficAction)
}

type mongoAttributesPUTrafficAction struct {
}
type mongoAttributesSparsePUTrafficAction struct {
}
