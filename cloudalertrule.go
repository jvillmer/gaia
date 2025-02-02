package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudAlertRuleIdentity represents the Identity of the object.
var CloudAlertRuleIdentity = elemental.Identity{
	Name:     "cloudalertrule",
	Category: "cloudalertsrule",
	Package:  "vargid",
	Private:  false,
}

// CloudAlertRulesList represents a list of CloudAlertRules
type CloudAlertRulesList []*CloudAlertRule

// Identity returns the identity of the objects in the list.
func (o CloudAlertRulesList) Identity() elemental.Identity {

	return CloudAlertRuleIdentity
}

// Copy returns a pointer to a copy the CloudAlertRulesList.
func (o CloudAlertRulesList) Copy() elemental.Identifiables {

	copy := append(CloudAlertRulesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CloudAlertRulesList.
func (o CloudAlertRulesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CloudAlertRulesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CloudAlertRule))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CloudAlertRulesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CloudAlertRulesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the CloudAlertRulesList converted to SparseCloudAlertRulesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CloudAlertRulesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCloudAlertRulesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCloudAlertRule)
	}

	return out
}

// Version returns the version of the content.
func (o CloudAlertRulesList) Version() int {

	return 1
}

// CloudAlertRule represents the model of a cloudalertrule
type CloudAlertRule struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Defines whether the Alert rule is enabled.
	Enabled bool `json:"enabled" msgpack:"enabled" bson:"enabled" mapstructure:"enabled,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Prisma Cloud Alert Rule id.
	PrismaCloudAlertRuleID string `json:"prismaCloudAlertRuleID" msgpack:"prismaCloudAlertRuleID" bson:"prismacloudalertruleid" mapstructure:"prismaCloudAlertRuleID,omitempty"`

	// List of Policy IDs associated to an Alert rule.
	PrismaCloudPolicyIDs []string `json:"prismaCloudPolicyIDs" msgpack:"prismaCloudPolicyIDs" bson:"prismacloudpolicyids" mapstructure:"prismaCloudPolicyIDs,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// List of regions where the Alert rule is enforced.
	Regions []string `json:"regions" msgpack:"regions" bson:"regions" mapstructure:"regions,omitempty"`

	// List of accounts associated to an Alert rule.
	TargetAccountIDs []string `json:"targetAccountIDs" msgpack:"targetAccountIDs" bson:"targetaccountids" mapstructure:"targetAccountIDs,omitempty"`

	// Prisma ID of the tenant in which the Alert Rule is created.
	TenantPrismaID string `json:"tenantPrismaID" msgpack:"tenantPrismaID" bson:"tenantprismaid" mapstructure:"tenantPrismaID,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudAlertRule returns a new *CloudAlertRule
func NewCloudAlertRule() *CloudAlertRule {

	return &CloudAlertRule{
		ModelVersion:         1,
		Annotations:          map[string][]string{},
		AssociatedTags:       []string{},
		NormalizedTags:       []string{},
		Regions:              []string{},
		MigrationsLog:        map[string]string{},
		PrismaCloudPolicyIDs: []string{},
		TargetAccountIDs:     []string{},
	}
}

// Identity returns the Identity of the object.
func (o *CloudAlertRule) Identity() elemental.Identity {

	return CloudAlertRuleIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CloudAlertRule) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CloudAlertRule) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudAlertRule) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudAlertRule{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.Description = o.Description
	s.Enabled = o.Enabled
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.PrismaCloudAlertRuleID = o.PrismaCloudAlertRuleID
	s.PrismaCloudPolicyIDs = o.PrismaCloudPolicyIDs
	s.Protected = o.Protected
	s.Regions = o.Regions
	s.TargetAccountIDs = o.TargetAccountIDs
	s.TenantPrismaID = o.TenantPrismaID
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudAlertRule) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudAlertRule{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.Description = s.Description
	o.Enabled = s.Enabled
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.PrismaCloudAlertRuleID = s.PrismaCloudAlertRuleID
	o.PrismaCloudPolicyIDs = s.PrismaCloudPolicyIDs
	o.Protected = s.Protected
	o.Regions = s.Regions
	o.TargetAccountIDs = s.TargetAccountIDs
	o.TenantPrismaID = s.TenantPrismaID
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CloudAlertRule) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudAlertRule) BleveType() string {

	return "cloudalertrule"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CloudAlertRule) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *CloudAlertRule) Doc() string {

	return `Represents an Alert rule along with policies and accounts associated
with the Alert rule.`
}

func (o *CloudAlertRule) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *CloudAlertRule) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *CloudAlertRule) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *CloudAlertRule) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *CloudAlertRule) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *CloudAlertRule) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *CloudAlertRule) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *CloudAlertRule) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *CloudAlertRule) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *CloudAlertRule) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *CloudAlertRule) SetDescription(description string) {

	o.Description = description
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *CloudAlertRule) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *CloudAlertRule) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *CloudAlertRule) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *CloudAlertRule) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *CloudAlertRule) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *CloudAlertRule) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *CloudAlertRule) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *CloudAlertRule) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *CloudAlertRule) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *CloudAlertRule) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *CloudAlertRule) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *CloudAlertRule) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *CloudAlertRule) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *CloudAlertRule) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *CloudAlertRule) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *CloudAlertRule) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *CloudAlertRule) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *CloudAlertRule) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CloudAlertRule) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCloudAlertRule{
			ID:                     &o.ID,
			Annotations:            &o.Annotations,
			AssociatedTags:         &o.AssociatedTags,
			CreateIdempotencyKey:   &o.CreateIdempotencyKey,
			CreateTime:             &o.CreateTime,
			Description:            &o.Description,
			Enabled:                &o.Enabled,
			MigrationsLog:          &o.MigrationsLog,
			Name:                   &o.Name,
			Namespace:              &o.Namespace,
			NormalizedTags:         &o.NormalizedTags,
			PrismaCloudAlertRuleID: &o.PrismaCloudAlertRuleID,
			PrismaCloudPolicyIDs:   &o.PrismaCloudPolicyIDs,
			Protected:              &o.Protected,
			Regions:                &o.Regions,
			TargetAccountIDs:       &o.TargetAccountIDs,
			TenantPrismaID:         &o.TenantPrismaID,
			UpdateIdempotencyKey:   &o.UpdateIdempotencyKey,
			UpdateTime:             &o.UpdateTime,
			ZHash:                  &o.ZHash,
			Zone:                   &o.Zone,
		}
	}

	sp := &SparseCloudAlertRule{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "enabled":
			sp.Enabled = &(o.Enabled)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "prismaCloudAlertRuleID":
			sp.PrismaCloudAlertRuleID = &(o.PrismaCloudAlertRuleID)
		case "prismaCloudPolicyIDs":
			sp.PrismaCloudPolicyIDs = &(o.PrismaCloudPolicyIDs)
		case "protected":
			sp.Protected = &(o.Protected)
		case "regions":
			sp.Regions = &(o.Regions)
		case "targetAccountIDs":
			sp.TargetAccountIDs = &(o.TargetAccountIDs)
		case "tenantPrismaID":
			sp.TenantPrismaID = &(o.TenantPrismaID)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCloudAlertRule to the object.
func (o *CloudAlertRule) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCloudAlertRule)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Enabled != nil {
		o.Enabled = *so.Enabled
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.PrismaCloudAlertRuleID != nil {
		o.PrismaCloudAlertRuleID = *so.PrismaCloudAlertRuleID
	}
	if so.PrismaCloudPolicyIDs != nil {
		o.PrismaCloudPolicyIDs = *so.PrismaCloudPolicyIDs
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Regions != nil {
		o.Regions = *so.Regions
	}
	if so.TargetAccountIDs != nil {
		o.TargetAccountIDs = *so.TargetAccountIDs
	}
	if so.TenantPrismaID != nil {
		o.TenantPrismaID = *so.TenantPrismaID
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the CloudAlertRule.
func (o *CloudAlertRule) DeepCopy() *CloudAlertRule {

	if o == nil {
		return nil
	}

	out := &CloudAlertRule{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudAlertRule.
func (o *CloudAlertRule) DeepCopyInto(out *CloudAlertRule) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudAlertRule: %s", err))
	}

	*out = *target.(*CloudAlertRule)
}

// Validate valides the current information stored into the structure.
func (o *CloudAlertRule) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
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
func (*CloudAlertRule) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CloudAlertRuleAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CloudAlertRuleLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CloudAlertRule) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CloudAlertRuleAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CloudAlertRule) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "enabled":
		return o.Enabled
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "prismaCloudAlertRuleID":
		return o.PrismaCloudAlertRuleID
	case "prismaCloudPolicyIDs":
		return o.PrismaCloudPolicyIDs
	case "protected":
		return o.Protected
	case "regions":
		return o.Regions
	case "targetAccountIDs":
		return o.TargetAccountIDs
	case "tenantPrismaID":
		return o.TenantPrismaID
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// CloudAlertRuleAttributesMap represents the map of attribute for CloudAlertRule.
var CloudAlertRuleAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Annotations": {
		AllowedChoices: []string{},
		BSONFieldName:  "annotations",
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"AssociatedTags": {
		AllowedChoices: []string{},
		BSONFieldName:  "associatedtags",
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"CreateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createidempotencykey",
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createtime",
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Enabled": {
		AllowedChoices: []string{},
		BSONFieldName:  "enabled",
		ConvertedName:  "Enabled",
		Description:    `Defines whether the Alert rule is enabled.`,
		Exposed:        true,
		Name:           "enabled",
		Stored:         true,
		SubType:        "boolean",
		Type:           "boolean",
	},
	"MigrationsLog": {
		AllowedChoices: []string{},
		BSONFieldName:  "migrationslog",
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "normalizedtags",
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"PrismaCloudAlertRuleID": {
		AllowedChoices: []string{},
		BSONFieldName:  "prismacloudalertruleid",
		ConvertedName:  "PrismaCloudAlertRuleID",
		Description:    `Prisma Cloud Alert Rule id.`,
		Exposed:        true,
		Name:           "prismaCloudAlertRuleID",
		Stored:         true,
		SubType:        "string",
		Type:           "string",
	},
	"PrismaCloudPolicyIDs": {
		AllowedChoices: []string{},
		BSONFieldName:  "prismacloudpolicyids",
		ConvertedName:  "PrismaCloudPolicyIDs",
		Description:    `List of Policy IDs associated to an Alert rule.`,
		Exposed:        true,
		Name:           "prismaCloudPolicyIDs",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Protected": {
		AllowedChoices: []string{},
		BSONFieldName:  "protected",
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Regions": {
		AllowedChoices: []string{},
		BSONFieldName:  "regions",
		ConvertedName:  "Regions",
		Description:    `List of regions where the Alert rule is enforced.`,
		Exposed:        true,
		Name:           "regions",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"TargetAccountIDs": {
		AllowedChoices: []string{},
		BSONFieldName:  "targetaccountids",
		ConvertedName:  "TargetAccountIDs",
		Description:    `List of accounts associated to an Alert rule.`,
		Exposed:        true,
		Name:           "targetAccountIDs",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"TenantPrismaID": {
		AllowedChoices: []string{},
		BSONFieldName:  "tenantprismaid",
		ConvertedName:  "TenantPrismaID",
		Description:    `Prisma ID of the tenant in which the Alert Rule is created.`,
		Exposed:        true,
		Name:           "tenantPrismaID",
		Stored:         true,
		SubType:        "string",
		Type:           "string",
	},
	"UpdateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updateidempotencykey",
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"UpdateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updatetime",
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"ZHash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zhash",
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zone",
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// CloudAlertRuleLowerCaseAttributesMap represents the map of attribute for CloudAlertRule.
var CloudAlertRuleLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"annotations": {
		AllowedChoices: []string{},
		BSONFieldName:  "annotations",
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"associatedtags": {
		AllowedChoices: []string{},
		BSONFieldName:  "associatedtags",
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"createidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createidempotencykey",
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createtime",
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"enabled": {
		AllowedChoices: []string{},
		BSONFieldName:  "enabled",
		ConvertedName:  "Enabled",
		Description:    `Defines whether the Alert rule is enabled.`,
		Exposed:        true,
		Name:           "enabled",
		Stored:         true,
		SubType:        "boolean",
		Type:           "boolean",
	},
	"migrationslog": {
		AllowedChoices: []string{},
		BSONFieldName:  "migrationslog",
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"normalizedtags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "normalizedtags",
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"prismacloudalertruleid": {
		AllowedChoices: []string{},
		BSONFieldName:  "prismacloudalertruleid",
		ConvertedName:  "PrismaCloudAlertRuleID",
		Description:    `Prisma Cloud Alert Rule id.`,
		Exposed:        true,
		Name:           "prismaCloudAlertRuleID",
		Stored:         true,
		SubType:        "string",
		Type:           "string",
	},
	"prismacloudpolicyids": {
		AllowedChoices: []string{},
		BSONFieldName:  "prismacloudpolicyids",
		ConvertedName:  "PrismaCloudPolicyIDs",
		Description:    `List of Policy IDs associated to an Alert rule.`,
		Exposed:        true,
		Name:           "prismaCloudPolicyIDs",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"protected": {
		AllowedChoices: []string{},
		BSONFieldName:  "protected",
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"regions": {
		AllowedChoices: []string{},
		BSONFieldName:  "regions",
		ConvertedName:  "Regions",
		Description:    `List of regions where the Alert rule is enforced.`,
		Exposed:        true,
		Name:           "regions",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"targetaccountids": {
		AllowedChoices: []string{},
		BSONFieldName:  "targetaccountids",
		ConvertedName:  "TargetAccountIDs",
		Description:    `List of accounts associated to an Alert rule.`,
		Exposed:        true,
		Name:           "targetAccountIDs",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"tenantprismaid": {
		AllowedChoices: []string{},
		BSONFieldName:  "tenantprismaid",
		ConvertedName:  "TenantPrismaID",
		Description:    `Prisma ID of the tenant in which the Alert Rule is created.`,
		Exposed:        true,
		Name:           "tenantPrismaID",
		Stored:         true,
		SubType:        "string",
		Type:           "string",
	},
	"updateidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updateidempotencykey",
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"updatetime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updatetime",
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"zhash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zhash",
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zone",
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SparseCloudAlertRulesList represents a list of SparseCloudAlertRules
type SparseCloudAlertRulesList []*SparseCloudAlertRule

// Identity returns the identity of the objects in the list.
func (o SparseCloudAlertRulesList) Identity() elemental.Identity {

	return CloudAlertRuleIdentity
}

// Copy returns a pointer to a copy the SparseCloudAlertRulesList.
func (o SparseCloudAlertRulesList) Copy() elemental.Identifiables {

	copy := append(SparseCloudAlertRulesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCloudAlertRulesList.
func (o SparseCloudAlertRulesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCloudAlertRulesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCloudAlertRule))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCloudAlertRulesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCloudAlertRulesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseCloudAlertRulesList converted to CloudAlertRulesList.
func (o SparseCloudAlertRulesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCloudAlertRulesList) Version() int {

	return 1
}

// SparseCloudAlertRule represents the sparse version of a cloudalertrule.
type SparseCloudAlertRule struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Defines whether the Alert rule is enabled.
	Enabled *bool `json:"enabled,omitempty" msgpack:"enabled,omitempty" bson:"enabled,omitempty" mapstructure:"enabled,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Prisma Cloud Alert Rule id.
	PrismaCloudAlertRuleID *string `json:"prismaCloudAlertRuleID,omitempty" msgpack:"prismaCloudAlertRuleID,omitempty" bson:"prismacloudalertruleid,omitempty" mapstructure:"prismaCloudAlertRuleID,omitempty"`

	// List of Policy IDs associated to an Alert rule.
	PrismaCloudPolicyIDs *[]string `json:"prismaCloudPolicyIDs,omitempty" msgpack:"prismaCloudPolicyIDs,omitempty" bson:"prismacloudpolicyids,omitempty" mapstructure:"prismaCloudPolicyIDs,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// List of regions where the Alert rule is enforced.
	Regions *[]string `json:"regions,omitempty" msgpack:"regions,omitempty" bson:"regions,omitempty" mapstructure:"regions,omitempty"`

	// List of accounts associated to an Alert rule.
	TargetAccountIDs *[]string `json:"targetAccountIDs,omitempty" msgpack:"targetAccountIDs,omitempty" bson:"targetaccountids,omitempty" mapstructure:"targetAccountIDs,omitempty"`

	// Prisma ID of the tenant in which the Alert Rule is created.
	TenantPrismaID *string `json:"tenantPrismaID,omitempty" msgpack:"tenantPrismaID,omitempty" bson:"tenantprismaid,omitempty" mapstructure:"tenantPrismaID,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCloudAlertRule returns a new  SparseCloudAlertRule.
func NewSparseCloudAlertRule() *SparseCloudAlertRule {
	return &SparseCloudAlertRule{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCloudAlertRule) Identity() elemental.Identity {

	return CloudAlertRuleIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCloudAlertRule) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCloudAlertRule) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudAlertRule) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCloudAlertRule{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.Enabled != nil {
		s.Enabled = o.Enabled
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.PrismaCloudAlertRuleID != nil {
		s.PrismaCloudAlertRuleID = o.PrismaCloudAlertRuleID
	}
	if o.PrismaCloudPolicyIDs != nil {
		s.PrismaCloudPolicyIDs = o.PrismaCloudPolicyIDs
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.Regions != nil {
		s.Regions = o.Regions
	}
	if o.TargetAccountIDs != nil {
		s.TargetAccountIDs = o.TargetAccountIDs
	}
	if o.TenantPrismaID != nil {
		s.TenantPrismaID = o.TenantPrismaID
	}
	if o.UpdateIdempotencyKey != nil {
		s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		s.UpdateTime = o.UpdateTime
	}
	if o.ZHash != nil {
		s.ZHash = o.ZHash
	}
	if o.Zone != nil {
		s.Zone = o.Zone
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCloudAlertRule) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCloudAlertRule{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Enabled != nil {
		o.Enabled = s.Enabled
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.PrismaCloudAlertRuleID != nil {
		o.PrismaCloudAlertRuleID = s.PrismaCloudAlertRuleID
	}
	if s.PrismaCloudPolicyIDs != nil {
		o.PrismaCloudPolicyIDs = s.PrismaCloudPolicyIDs
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.Regions != nil {
		o.Regions = s.Regions
	}
	if s.TargetAccountIDs != nil {
		o.TargetAccountIDs = s.TargetAccountIDs
	}
	if s.TenantPrismaID != nil {
		o.TenantPrismaID = s.TenantPrismaID
	}
	if s.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	}
	if s.UpdateTime != nil {
		o.UpdateTime = s.UpdateTime
	}
	if s.ZHash != nil {
		o.ZHash = s.ZHash
	}
	if s.Zone != nil {
		o.Zone = s.Zone
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCloudAlertRule) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCloudAlertRule) ToPlain() elemental.PlainIdentifiable {

	out := NewCloudAlertRule()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Enabled != nil {
		out.Enabled = *o.Enabled
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.PrismaCloudAlertRuleID != nil {
		out.PrismaCloudAlertRuleID = *o.PrismaCloudAlertRuleID
	}
	if o.PrismaCloudPolicyIDs != nil {
		out.PrismaCloudPolicyIDs = *o.PrismaCloudPolicyIDs
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Regions != nil {
		out.Regions = *o.Regions
	}
	if o.TargetAccountIDs != nil {
		out.TargetAccountIDs = *o.TargetAccountIDs
	}
	if o.TenantPrismaID != nil {
		out.TenantPrismaID = *o.TenantPrismaID
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseCloudAlertRule) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseCloudAlertRule) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseCloudAlertRule) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseCloudAlertRule) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseCloudAlertRule) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetDescription(description string) {

	o.Description = &description
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseCloudAlertRule) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseCloudAlertRule) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseCloudAlertRule) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseCloudAlertRule) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseCloudAlertRule) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseCloudAlertRule) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseCloudAlertRule) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseCloudAlertRule) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseCloudAlertRule) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseCloudAlertRule) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseCloudAlertRule.
func (o *SparseCloudAlertRule) DeepCopy() *SparseCloudAlertRule {

	if o == nil {
		return nil
	}

	out := &SparseCloudAlertRule{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCloudAlertRule.
func (o *SparseCloudAlertRule) DeepCopyInto(out *SparseCloudAlertRule) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCloudAlertRule: %s", err))
	}

	*out = *target.(*SparseCloudAlertRule)
}

type mongoAttributesCloudAlertRule struct {
	ID                     bson.ObjectId       `bson:"_id,omitempty"`
	Annotations            map[string][]string `bson:"annotations"`
	AssociatedTags         []string            `bson:"associatedtags"`
	CreateIdempotencyKey   string              `bson:"createidempotencykey"`
	CreateTime             time.Time           `bson:"createtime"`
	Description            string              `bson:"description"`
	Enabled                bool                `bson:"enabled"`
	MigrationsLog          map[string]string   `bson:"migrationslog,omitempty"`
	Name                   string              `bson:"name"`
	Namespace              string              `bson:"namespace"`
	NormalizedTags         []string            `bson:"normalizedtags"`
	PrismaCloudAlertRuleID string              `bson:"prismacloudalertruleid"`
	PrismaCloudPolicyIDs   []string            `bson:"prismacloudpolicyids"`
	Protected              bool                `bson:"protected"`
	Regions                []string            `bson:"regions"`
	TargetAccountIDs       []string            `bson:"targetaccountids"`
	TenantPrismaID         string              `bson:"tenantprismaid"`
	UpdateIdempotencyKey   string              `bson:"updateidempotencykey"`
	UpdateTime             time.Time           `bson:"updatetime"`
	ZHash                  int                 `bson:"zhash"`
	Zone                   int                 `bson:"zone"`
}
type mongoAttributesSparseCloudAlertRule struct {
	ID                     bson.ObjectId        `bson:"_id,omitempty"`
	Annotations            *map[string][]string `bson:"annotations,omitempty"`
	AssociatedTags         *[]string            `bson:"associatedtags,omitempty"`
	CreateIdempotencyKey   *string              `bson:"createidempotencykey,omitempty"`
	CreateTime             *time.Time           `bson:"createtime,omitempty"`
	Description            *string              `bson:"description,omitempty"`
	Enabled                *bool                `bson:"enabled,omitempty"`
	MigrationsLog          *map[string]string   `bson:"migrationslog,omitempty"`
	Name                   *string              `bson:"name,omitempty"`
	Namespace              *string              `bson:"namespace,omitempty"`
	NormalizedTags         *[]string            `bson:"normalizedtags,omitempty"`
	PrismaCloudAlertRuleID *string              `bson:"prismacloudalertruleid,omitempty"`
	PrismaCloudPolicyIDs   *[]string            `bson:"prismacloudpolicyids,omitempty"`
	Protected              *bool                `bson:"protected,omitempty"`
	Regions                *[]string            `bson:"regions,omitempty"`
	TargetAccountIDs       *[]string            `bson:"targetaccountids,omitempty"`
	TenantPrismaID         *string              `bson:"tenantprismaid,omitempty"`
	UpdateIdempotencyKey   *string              `bson:"updateidempotencykey,omitempty"`
	UpdateTime             *time.Time           `bson:"updatetime,omitempty"`
	ZHash                  *int                 `bson:"zhash,omitempty"`
	Zone                   *int                 `bson:"zone,omitempty"`
}
