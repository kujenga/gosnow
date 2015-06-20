package gosnow

type TransactionExample struct {
	Name        string
	Description string
	Requests    []Payload
	Responses   []Payload
}

type Action struct {
	Name        string
	Description string
	Method      string
	Parameters  []Parameter
	Attributes  struct {
		Relation    string
		UTITemplate string
	}
	Content  []interface{}
	Examples []TransactionExample
}

type Parameter struct {
	Name        string
	Type        string
	Description string
	Required    bool
	Default     string
	Example     string
	Values      []struct {
		Value string
	}
}

const contentTypeHeaderKey = "Content-Type"

type Payload struct {
	Name        string
	Description string
	Headers     []struct {
		Name  string
		Value string
	}
	Body    string
	Schema  string
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

// ASTBlueprint is a data representation of the ast blueprints
type ASTBlueprint struct {
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

// SourcemapBlueprint is a blueprint for sourcemaps
// TODO: fill this out
type SourcemapBlueprint interface{}
