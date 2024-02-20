package services

import "github.com/ankeshnirala/http-server-go1.22/internals/applications/models"

// AuthorizationService defines methods for authorizing actions on resources and filtering data based on field access restrictions.
type AuthorizationService interface {
	AuthorizeOrThrow(principal models.SecurityPrincipal, resource models.Resource, action models.Action, context map[string]interface{}) (models.Policy, error)
	FilterData(fieldAccess []string, entity map[string]interface{}) map[string]interface{}
}
