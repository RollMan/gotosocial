// Code generated by astool. DO NOT EDIT.

package vocab

import "net/url"

// Specifies an identifier for a Branch, that is used in the Repository to
// uniquely refer to it. For example, in Git, "refs/heads/master" would be the
// ref of the master branch.
//
//   {
//     "@context": [
//       "https://www.w3.org/ns/activitystreams",
//       "https://forgefed.peers.community/ns"
//     ],
//     "context": "https://example.dev/luke/myrepo",
//     "id": "https://example.dev/luke/myrepo/branches/master",
//     "name": "master",
//     "ref": "refs/heads/master",
//     "type": "Branch"
//   }
type ForgeFedRefProperty interface {
	// Clear ensures no value of this property is set. Calling
	// IsXMLSchemaString afterwards will return false.
	Clear()
	// Get returns the value of this property. When IsXMLSchemaString returns
	// false, Get will return any arbitrary value.
	Get() string
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return any arbitrary value.
	GetIRI() *url.URL
	// HasAny returns true if the value or IRI is set.
	HasAny() bool
	// IsIRI returns true if this property is an IRI.
	IsIRI() bool
	// IsXMLSchemaString returns true if this property is set and not an IRI.
	IsXMLSchemaString() bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API detail only for folks looking to replace the
	// go-fed implementation. Applications should not use this method.
	KindIndex() int
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o ForgeFedRefProperty) bool
	// Name returns the name of this property: "ref".
	Name() string
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// Set sets the value of this property. Calling IsXMLSchemaString
	// afterwards will return true.
	Set(v string)
	// SetIRI sets the value of this property. Calling IsIRI afterwards will
	// return true.
	SetIRI(v *url.URL)
}