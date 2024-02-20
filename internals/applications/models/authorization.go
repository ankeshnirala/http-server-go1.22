package models

// Role represents a user role.
type Role string

// Action represents an action that can be performed on a resource.
type Action string

// Resource represents a resource/entity.
type Resource string

// SecurityPrincipal represents a security principal with an ID and a role.
type SecurityPrincipal struct {
	ID   string
	Role Role
}

// Policy represents a policy for a specific resource.
type Policy struct {
	Resource              Resource
	CanAccess             func(context map[string]interface{}) (bool, error)
	SelectionRestrictions map[string]interface{}
	FieldAccess           []string
}
