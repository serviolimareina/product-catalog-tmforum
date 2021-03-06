/*
 * Product Catalog Management
 *
 * ## TMF API Reference: TMF620 - Product Catalog Management  ### Release : 19.0 - June 2019  Product Catalog API is one of Catalog Management API Family. Product Catalog API goal is to provide a catalog of products.   ### Operations Product Catalog API performs the following operations on the resources : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity - Manage notification of events
 *
 * API version: 4.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Describes a non-composite relationship between product offering prices. For example one price might be an discount alteration for another price.
type ProductOfferingPriceRelationship struct {

	// Unique identifier of the associated product offering price
	Id string `json:"id,omitempty"`

	// hyperlink reference of the associated product offering price
	Href string `json:"href,omitempty"`

	// Name of the associated product offering price
	Name string `json:"name,omitempty"`

	// type of the relationship, for example override, discount, etc.
	RelationshipType string `json:"relationshipType,omitempty"`

	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class entity name
	Type_ string `json:"@type,omitempty"`
}
