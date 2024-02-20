package authorizationservice

import (
	"github.com/ankeshnirala/http-server-go1.22/internals/applications/interfaces/services"
	"github.com/ankeshnirala/http-server-go1.22/internals/applications/models"
)

// Role represents a user role.
const (
	User    models.Role = "user"
	Partner models.Role = "partner"
	Admin   models.Role = "admin"
)

// Action represents an action that can be performed on a resource.
const (
	Create models.Action = "Create"
	Read   models.Action = "Read"
	Delete models.Action = "Delete"
)

// Resource represents a resource/entity.
const (
	Accounts models.Resource = "Accounts"
	Users    models.Resource = "Users"
	Currency models.Resource = "Currency"
)

type authorizationService struct{}

// NewAuthorizationService creates a new instance of AuthorizationService.
func NewAuthorizationService() services.AuthorizationService {
	return &authorizationService{}
}

func (svc *authorizationService) AuthorizeOrThrow(principal models.SecurityPrincipal, resource models.Resource, action models.Action, context map[string]interface{}) (models.Policy, error) {
	// Here you would implement your authorization logic based on the provided principal, resource, and action.
	// For demonstration purposes, we'll just return a dummy policy.
	policy := models.Policy{
		Resource:              resource,
		CanAccess:             func(context map[string]interface{}) (bool, error) { return true, nil },
		SelectionRestrictions: nil,
		FieldAccess:           []string{"*"},
	}
	return policy, nil
}

func (svc *authorizationService) FilterData(fieldAccess []string, entity map[string]interface{}) map[string]interface{} {
	filteredData := make(map[string]interface{})
	for _, field := range fieldAccess {
		filteredData[field] = entity[field]
	}
	return filteredData
}

// authService := NewAuthorizationService()
// 	principal := SecurityPrincipal{ID: "user123", Role: User}
// 	resource := Account
// 	action := Read
// 	context := map[string]interface{}{"entity": map[string]interface{}{"id": "123", "name": "Account123"}}
// 	policy, err := authService.AuthorizeOrThrow(principal, resource, action, context)
// 	if err != nil {
// 		fmt.Println("Authorization error:", err)
// 		return
// 	}
// 	fmt.Println("Authorization granted for resource:", policy.Resource)
