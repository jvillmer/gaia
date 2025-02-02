package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// X509CertificateSignerValue represents the possible values for attribute "signer".
type X509CertificateSignerValue string

const (
	// X509CertificateSignerNamespace represents the value Namespace.
	X509CertificateSignerNamespace X509CertificateSignerValue = "Namespace"

	// X509CertificateSignerPublic represents the value Public.
	X509CertificateSignerPublic X509CertificateSignerValue = "Public"

	// X509CertificateSignerSystem represents the value System.
	X509CertificateSignerSystem X509CertificateSignerValue = "System"
)

// X509CertificateUsageValue represents the possible values for attribute "usage".
type X509CertificateUsageValue string

const (
	// X509CertificateUsageClient represents the value Client.
	X509CertificateUsageClient X509CertificateUsageValue = "Client"

	// X509CertificateUsageServer represents the value Server.
	X509CertificateUsageServer X509CertificateUsageValue = "Server"

	// X509CertificateUsageServerClient represents the value ServerClient.
	X509CertificateUsageServerClient X509CertificateUsageValue = "ServerClient"
)

// X509CertificateIdentity represents the Identity of the object.
var X509CertificateIdentity = elemental.Identity{
	Name:     "x509certificate",
	Category: "x509certificates",
	Package:  "barret",
	Private:  true,
}

// X509CertificatesList represents a list of X509Certificates
type X509CertificatesList []*X509Certificate

// Identity returns the identity of the objects in the list.
func (o X509CertificatesList) Identity() elemental.Identity {

	return X509CertificateIdentity
}

// Copy returns a pointer to a copy the X509CertificatesList.
func (o X509CertificatesList) Copy() elemental.Identifiables {

	copy := append(X509CertificatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the X509CertificatesList.
func (o X509CertificatesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(X509CertificatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*X509Certificate))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o X509CertificatesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o X509CertificatesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the X509CertificatesList converted to SparseX509CertificatesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o X509CertificatesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseX509CertificatesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseX509Certificate)
	}

	return out
}

// Version returns the version of the content.
func (o X509CertificatesList) Version() int {

	return 1
}

// X509Certificate represents the model of a x509certificate
type X509Certificate struct {
	// Contains the Certificate Signing Request (CSR) as a PEM-encoded string.
	CSR string `json:"CSR" msgpack:"CSR" bson:"-" mapstructure:"CSR,omitempty"`

	// ID contains the identifier of the certificate.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Certificate contains the certificate data in PEM format.
	Certificate string `json:"certificate" msgpack:"certificate" bson:"-" mapstructure:"certificate,omitempty"`

	// ExpirationDate contains the requested expiration date.
	ExpirationDate time.Time `json:"expirationDate" msgpack:"expirationDate" bson:"-" mapstructure:"expirationDate,omitempty"`

	// Extensions to add to the certificate. It must contains ASN1-encoded bytes,
	// themselves encoded in base64.
	Extensions []byte `json:"extensions" msgpack:"extensions" bson:"-" mapstructure:"extensions,omitempty"`

	// Selects what CA should sign the certificate.
	Signer X509CertificateSignerValue `json:"signer" msgpack:"signer" bson:"-" mapstructure:"signer,omitempty"`

	// Selects the ID of the CA that should sign the certificate when signer is set to
	// Namespace.
	SignerID string `json:"signerID" msgpack:"signerID" bson:"-" mapstructure:"signerID,omitempty"`

	// Additional subject information to use to override the ones in the CSR.
	SubjectOverride *PKIXName `json:"subjectOverride,omitempty" msgpack:"subjectOverride,omitempty" bson:"-" mapstructure:"subjectOverride,omitempty"`

	// If set to true, the certificate is considered short lived and it will not be
	// possible to revoke it.
	Unrevocable bool `json:"unrevocable" msgpack:"unrevocable" bson:"-" mapstructure:"unrevocable,omitempty"`

	// Usage defines the requested key usage.
	Usage X509CertificateUsageValue `json:"usage" msgpack:"usage" bson:"-" mapstructure:"usage,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewX509Certificate returns a new *X509Certificate
func NewX509Certificate() *X509Certificate {

	return &X509Certificate{
		ModelVersion:    1,
		Extensions:      []byte{},
		Signer:          X509CertificateSignerPublic,
		SubjectOverride: NewPKIXName(),
		Usage:           X509CertificateUsageClient,
	}
}

// Identity returns the Identity of the object.
func (o *X509Certificate) Identity() elemental.Identity {

	return X509CertificateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *X509Certificate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *X509Certificate) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *X509Certificate) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesX509Certificate{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *X509Certificate) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesX509Certificate{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *X509Certificate) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *X509Certificate) BleveType() string {

	return "x509certificate"
}

// DefaultOrder returns the list of default ordering fields.
func (o *X509Certificate) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *X509Certificate) Doc() string {

	return `This API allows you to retrieve an client certificate for API authentication.`
}

func (o *X509Certificate) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *X509Certificate) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseX509Certificate{
			CSR:             &o.CSR,
			ID:              &o.ID,
			Certificate:     &o.Certificate,
			ExpirationDate:  &o.ExpirationDate,
			Extensions:      &o.Extensions,
			Signer:          &o.Signer,
			SignerID:        &o.SignerID,
			SubjectOverride: o.SubjectOverride,
			Unrevocable:     &o.Unrevocable,
			Usage:           &o.Usage,
		}
	}

	sp := &SparseX509Certificate{}
	for _, f := range fields {
		switch f {
		case "CSR":
			sp.CSR = &(o.CSR)
		case "ID":
			sp.ID = &(o.ID)
		case "certificate":
			sp.Certificate = &(o.Certificate)
		case "expirationDate":
			sp.ExpirationDate = &(o.ExpirationDate)
		case "extensions":
			sp.Extensions = &(o.Extensions)
		case "signer":
			sp.Signer = &(o.Signer)
		case "signerID":
			sp.SignerID = &(o.SignerID)
		case "subjectOverride":
			sp.SubjectOverride = o.SubjectOverride
		case "unrevocable":
			sp.Unrevocable = &(o.Unrevocable)
		case "usage":
			sp.Usage = &(o.Usage)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseX509Certificate to the object.
func (o *X509Certificate) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseX509Certificate)
	if so.CSR != nil {
		o.CSR = *so.CSR
	}
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Certificate != nil {
		o.Certificate = *so.Certificate
	}
	if so.ExpirationDate != nil {
		o.ExpirationDate = *so.ExpirationDate
	}
	if so.Extensions != nil {
		o.Extensions = *so.Extensions
	}
	if so.Signer != nil {
		o.Signer = *so.Signer
	}
	if so.SignerID != nil {
		o.SignerID = *so.SignerID
	}
	if so.SubjectOverride != nil {
		o.SubjectOverride = so.SubjectOverride
	}
	if so.Unrevocable != nil {
		o.Unrevocable = *so.Unrevocable
	}
	if so.Usage != nil {
		o.Usage = *so.Usage
	}
}

// DeepCopy returns a deep copy if the X509Certificate.
func (o *X509Certificate) DeepCopy() *X509Certificate {

	if o == nil {
		return nil
	}

	out := &X509Certificate{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *X509Certificate.
func (o *X509Certificate) DeepCopyInto(out *X509Certificate) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy X509Certificate: %s", err))
	}

	*out = *target.(*X509Certificate)
}

// Validate valides the current information stored into the structure.
func (o *X509Certificate) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("CSR", o.CSR); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("signer", string(o.Signer), []string{"Public", "Namespace", "System"}, false); err != nil {
		errors = errors.Append(err)
	}

	if o.SubjectOverride != nil {
		elemental.ResetDefaultForZeroValues(o.SubjectOverride)
		if err := o.SubjectOverride.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateStringInList("usage", string(o.Usage), []string{"Client", "Server", "ServerClient"}, false); err != nil {
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
func (*X509Certificate) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := X509CertificateAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return X509CertificateLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*X509Certificate) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return X509CertificateAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *X509Certificate) ValueForAttribute(name string) interface{} {

	switch name {
	case "CSR":
		return o.CSR
	case "ID":
		return o.ID
	case "certificate":
		return o.Certificate
	case "expirationDate":
		return o.ExpirationDate
	case "extensions":
		return o.Extensions
	case "signer":
		return o.Signer
	case "signerID":
		return o.SignerID
	case "subjectOverride":
		return o.SubjectOverride
	case "unrevocable":
		return o.Unrevocable
	case "usage":
		return o.Usage
	}

	return nil
}

// X509CertificateAttributesMap represents the map of attribute for X509Certificate.
var X509CertificateAttributesMap = map[string]elemental.AttributeSpecification{
	"CSR": {
		AllowedChoices: []string{},
		ConvertedName:  "CSR",
		CreationOnly:   true,
		Description:    `Contains the Certificate Signing Request (CSR) as a PEM-encoded string.`,
		Exposed:        true,
		Name:           "CSR",
		Required:       true,
		Type:           "string",
	},
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID contains the identifier of the certificate.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		ReadOnly:       true,
		Type:           "string",
	},
	"Certificate": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Certificate",
		Description:    `Certificate contains the certificate data in PEM format.`,
		Exposed:        true,
		Name:           "certificate",
		ReadOnly:       true,
		Type:           "string",
	},
	"ExpirationDate": {
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationDate",
		CreationOnly:   true,
		Description:    `ExpirationDate contains the requested expiration date.`,
		Exposed:        true,
		Name:           "expirationDate",
		SubType:        "string",
		Type:           "time",
	},
	"Extensions": {
		AllowedChoices: []string{},
		ConvertedName:  "Extensions",
		CreationOnly:   true,
		Description: `Extensions to add to the certificate. It must contains ASN1-encoded bytes,
themselves encoded in base64.`,
		Exposed: true,
		Name:    "extensions",
		SubType: "[]byte",
		Type:    "external",
	},
	"Signer": {
		AllowedChoices: []string{"Public", "Namespace", "System"},
		ConvertedName:  "Signer",
		DefaultValue:   X509CertificateSignerPublic,
		Description:    `Selects what CA should sign the certificate.`,
		Exposed:        true,
		Name:           "signer",
		Type:           "enum",
	},
	"SignerID": {
		AllowedChoices: []string{},
		ConvertedName:  "SignerID",
		Description: `Selects the ID of the CA that should sign the certificate when signer is set to
Namespace.`,
		Exposed: true,
		Name:    "signerID",
		Type:    "string",
	},
	"SubjectOverride": {
		AllowedChoices: []string{},
		ConvertedName:  "SubjectOverride",
		CreationOnly:   true,
		Description:    `Additional subject information to use to override the ones in the CSR.`,
		Exposed:        true,
		Name:           "subjectOverride",
		SubType:        "pkixname",
		Type:           "ref",
	},
	"Unrevocable": {
		AllowedChoices: []string{},
		ConvertedName:  "Unrevocable",
		Description: `If set to true, the certificate is considered short lived and it will not be
possible to revoke it.`,
		Exposed: true,
		Name:    "unrevocable",
		Type:    "boolean",
	},
	"Usage": {
		AllowedChoices: []string{"Client", "Server", "ServerClient"},
		ConvertedName:  "Usage",
		DefaultValue:   X509CertificateUsageClient,
		Description:    `Usage defines the requested key usage.`,
		Exposed:        true,
		Name:           "usage",
		Type:           "enum",
	},
}

// X509CertificateLowerCaseAttributesMap represents the map of attribute for X509Certificate.
var X509CertificateLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"csr": {
		AllowedChoices: []string{},
		ConvertedName:  "CSR",
		CreationOnly:   true,
		Description:    `Contains the Certificate Signing Request (CSR) as a PEM-encoded string.`,
		Exposed:        true,
		Name:           "CSR",
		Required:       true,
		Type:           "string",
	},
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID contains the identifier of the certificate.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		ReadOnly:       true,
		Type:           "string",
	},
	"certificate": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Certificate",
		Description:    `Certificate contains the certificate data in PEM format.`,
		Exposed:        true,
		Name:           "certificate",
		ReadOnly:       true,
		Type:           "string",
	},
	"expirationdate": {
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationDate",
		CreationOnly:   true,
		Description:    `ExpirationDate contains the requested expiration date.`,
		Exposed:        true,
		Name:           "expirationDate",
		SubType:        "string",
		Type:           "time",
	},
	"extensions": {
		AllowedChoices: []string{},
		ConvertedName:  "Extensions",
		CreationOnly:   true,
		Description: `Extensions to add to the certificate. It must contains ASN1-encoded bytes,
themselves encoded in base64.`,
		Exposed: true,
		Name:    "extensions",
		SubType: "[]byte",
		Type:    "external",
	},
	"signer": {
		AllowedChoices: []string{"Public", "Namespace", "System"},
		ConvertedName:  "Signer",
		DefaultValue:   X509CertificateSignerPublic,
		Description:    `Selects what CA should sign the certificate.`,
		Exposed:        true,
		Name:           "signer",
		Type:           "enum",
	},
	"signerid": {
		AllowedChoices: []string{},
		ConvertedName:  "SignerID",
		Description: `Selects the ID of the CA that should sign the certificate when signer is set to
Namespace.`,
		Exposed: true,
		Name:    "signerID",
		Type:    "string",
	},
	"subjectoverride": {
		AllowedChoices: []string{},
		ConvertedName:  "SubjectOverride",
		CreationOnly:   true,
		Description:    `Additional subject information to use to override the ones in the CSR.`,
		Exposed:        true,
		Name:           "subjectOverride",
		SubType:        "pkixname",
		Type:           "ref",
	},
	"unrevocable": {
		AllowedChoices: []string{},
		ConvertedName:  "Unrevocable",
		Description: `If set to true, the certificate is considered short lived and it will not be
possible to revoke it.`,
		Exposed: true,
		Name:    "unrevocable",
		Type:    "boolean",
	},
	"usage": {
		AllowedChoices: []string{"Client", "Server", "ServerClient"},
		ConvertedName:  "Usage",
		DefaultValue:   X509CertificateUsageClient,
		Description:    `Usage defines the requested key usage.`,
		Exposed:        true,
		Name:           "usage",
		Type:           "enum",
	},
}

// SparseX509CertificatesList represents a list of SparseX509Certificates
type SparseX509CertificatesList []*SparseX509Certificate

// Identity returns the identity of the objects in the list.
func (o SparseX509CertificatesList) Identity() elemental.Identity {

	return X509CertificateIdentity
}

// Copy returns a pointer to a copy the SparseX509CertificatesList.
func (o SparseX509CertificatesList) Copy() elemental.Identifiables {

	copy := append(SparseX509CertificatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseX509CertificatesList.
func (o SparseX509CertificatesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseX509CertificatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseX509Certificate))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseX509CertificatesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseX509CertificatesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseX509CertificatesList converted to X509CertificatesList.
func (o SparseX509CertificatesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseX509CertificatesList) Version() int {

	return 1
}

// SparseX509Certificate represents the sparse version of a x509certificate.
type SparseX509Certificate struct {
	// Contains the Certificate Signing Request (CSR) as a PEM-encoded string.
	CSR *string `json:"CSR,omitempty" msgpack:"CSR,omitempty" bson:"-" mapstructure:"CSR,omitempty"`

	// ID contains the identifier of the certificate.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Certificate contains the certificate data in PEM format.
	Certificate *string `json:"certificate,omitempty" msgpack:"certificate,omitempty" bson:"-" mapstructure:"certificate,omitempty"`

	// ExpirationDate contains the requested expiration date.
	ExpirationDate *time.Time `json:"expirationDate,omitempty" msgpack:"expirationDate,omitempty" bson:"-" mapstructure:"expirationDate,omitempty"`

	// Extensions to add to the certificate. It must contains ASN1-encoded bytes,
	// themselves encoded in base64.
	Extensions *[]byte `json:"extensions,omitempty" msgpack:"extensions,omitempty" bson:"-" mapstructure:"extensions,omitempty"`

	// Selects what CA should sign the certificate.
	Signer *X509CertificateSignerValue `json:"signer,omitempty" msgpack:"signer,omitempty" bson:"-" mapstructure:"signer,omitempty"`

	// Selects the ID of the CA that should sign the certificate when signer is set to
	// Namespace.
	SignerID *string `json:"signerID,omitempty" msgpack:"signerID,omitempty" bson:"-" mapstructure:"signerID,omitempty"`

	// Additional subject information to use to override the ones in the CSR.
	SubjectOverride *PKIXName `json:"subjectOverride,omitempty" msgpack:"subjectOverride,omitempty" bson:"-" mapstructure:"subjectOverride,omitempty"`

	// If set to true, the certificate is considered short lived and it will not be
	// possible to revoke it.
	Unrevocable *bool `json:"unrevocable,omitempty" msgpack:"unrevocable,omitempty" bson:"-" mapstructure:"unrevocable,omitempty"`

	// Usage defines the requested key usage.
	Usage *X509CertificateUsageValue `json:"usage,omitempty" msgpack:"usage,omitempty" bson:"-" mapstructure:"usage,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseX509Certificate returns a new  SparseX509Certificate.
func NewSparseX509Certificate() *SparseX509Certificate {
	return &SparseX509Certificate{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseX509Certificate) Identity() elemental.Identity {

	return X509CertificateIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseX509Certificate) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseX509Certificate) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseX509Certificate) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseX509Certificate{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseX509Certificate) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseX509Certificate{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseX509Certificate) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseX509Certificate) ToPlain() elemental.PlainIdentifiable {

	out := NewX509Certificate()
	if o.CSR != nil {
		out.CSR = *o.CSR
	}
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Certificate != nil {
		out.Certificate = *o.Certificate
	}
	if o.ExpirationDate != nil {
		out.ExpirationDate = *o.ExpirationDate
	}
	if o.Extensions != nil {
		out.Extensions = *o.Extensions
	}
	if o.Signer != nil {
		out.Signer = *o.Signer
	}
	if o.SignerID != nil {
		out.SignerID = *o.SignerID
	}
	if o.SubjectOverride != nil {
		out.SubjectOverride = o.SubjectOverride
	}
	if o.Unrevocable != nil {
		out.Unrevocable = *o.Unrevocable
	}
	if o.Usage != nil {
		out.Usage = *o.Usage
	}

	return out
}

// DeepCopy returns a deep copy if the SparseX509Certificate.
func (o *SparseX509Certificate) DeepCopy() *SparseX509Certificate {

	if o == nil {
		return nil
	}

	out := &SparseX509Certificate{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseX509Certificate.
func (o *SparseX509Certificate) DeepCopyInto(out *SparseX509Certificate) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseX509Certificate: %s", err))
	}

	*out = *target.(*SparseX509Certificate)
}

type mongoAttributesX509Certificate struct {
}
type mongoAttributesSparseX509Certificate struct {
}
