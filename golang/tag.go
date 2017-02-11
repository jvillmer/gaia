package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

// TagIdentity represents the Identity of the object
var TagIdentity = elemental.Identity{
	Name:     "tag",
	Category: "tags",
}

// TagsList represents a list of Tags
type TagsList []*Tag

// Tag represents the model of a tag
type Tag struct {
	// ID is the identifier of the object.
	ID string `json:"ID" cql:"id,omitempty" bson:"id"`

	// Count represents the number of time the tag is used.
	Count int `json:"count" cql:"count,omitempty" bson:"count"`

	// Namespace represents the namespace of the counted tag.
	Namespace string `json:"namespace" cql:"namespace,primarykey,omitempty" bson:"_namespace"`

	// Value represents the value of the tag.
	Value string `json:"value" cql:"value,primarykey,omitempty" bson:"_value"`
}

// NewTag returns a new *Tag
func NewTag() *Tag {

	return &Tag{}
}

// Identity returns the Identity of the object.
func (o *Tag) Identity() elemental.Identity {

	return TagIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Tag) Identifier() string {

	return o.ID
}

func (o *Tag) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Tag) SetIdentifier(ID string) {

	o.ID = ID
}

// Validate valides the current information stored into the structure.
func (o *Tag) Validate() error {

	errors := elemental.Errors{}

	if err := elemental.ValidatePattern("value", o.Value, `^[\w\d\*\$\+\.:,|@<>/-]+=[= \w\d\*\$\+\.:,|@<>/-]+$`); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("value", o.Value); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (Tag) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return TagAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (Tag) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TagAttributesMap
}

// TagAttributesMap represents the map of attribute for Tag.
var TagAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"Count": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Name:           "count",
		ReadOnly:       true,
		Stored:         true,
		Type:           "integer",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Format:         "free",
		Name:           "namespace",
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Value": elemental.AttributeSpecification{
		AllowedChars:   `^[\w\d\*\$\+\.:,|@<>/-]+=[= \w\d\*\$\+\.:,|@<>/-]+$`,
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Format:         "free",
		Name:           "value",
		PrimaryKey:     true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}
