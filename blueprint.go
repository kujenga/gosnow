package gosnow

/*
The structs in this file should be 1:1 with the Snow Crash AST
counterparts (https://github.com/apiaryio/snowcrash/blob/master/src/Blueprint.h).
*/

// TransactionExample Blueprint AST node
type TransactionExample struct {
	Name        string
	Description string
	// example request payloads
	Requests []Payload
	// example response payloads
	Responses []Payload
}

// Action Blueprint AST node represents 'action section'
type Action struct {
	Name        string
	Description string
	// HTTP request method or nil
	Method string
	// action-specific URI parameters or nil
	Parameters []Parameter
	// TODO: in RedSnow, these are flattened out. Ideally these should be too
	Attributes struct {
		// action relation attribute
		Relation string
		// action uri template attribute
		UTITemplate string
	}
	Content []interface{}
	// action transaction examples
	Examples []TransactionExample
}

// Parameter is a URI parameter Blueprint AST node
// represents one 'parameters section' parameter
type Parameter struct {
	Name string
	// an arbitrary type of the parameter or nil
	Type        string
	Description string
	// parameter necessity flag
	// TODO: change this to be an enum indicating required`, optional or undefined
	// Where undefined implies required according to the API Blueprint Specification
	Required bool
	// default value of the parameter or nil
	DefaultValue string `json:"default"`
	// example value of the parameter or nil
	ExampleValue string `json:"example"`
	// an enumeration of possible parameter values
	Values []struct {
		Value string
	}
}

const contentTypeHeaderKey = "Content-Type"

// Payload is a HTTP message payload Blueprint AST node
type Payload struct {
	Name        string
	Description string
	// array of HTTP header fields of the message or nil
	Headers []struct {
		Name  string
		Value string
	}
	// HTTP-message body or nil
	Body string
	// HTTP-message body validation schema or nil
	Schema string
	// Symbol Reference if the payload is a reference
	Reference struct {
		ID string
	}
	Content []interface{}
}

// Resource is a resource within the blueprint
type Resource struct {
	Element     string
	Name        string
	Description string
	URITemplate string
	Model       Payload
	Parameters  []Parameter
	Actions     []Action
}

// Blueprint is a data representation of the ast blueprints
type Blueprint struct {
	Version string `json:"_version"`
	// Metadata collection Blueprint AST node
	Metadata []struct {
		Name  string
		Value string
	}
	// Resource group Blueprint AST node
	// array of resources in the group
	ResourceGroups []struct {
		Name        string
		Description string
		Resources   []Resource
	}
	Name        string
	Description string
	Element     string
}
