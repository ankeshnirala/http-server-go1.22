package policies

// type PolicySetGenerator func(dataSource DataSource, securityPrincipal SecurityPrincipal) (map[interface{}]Policy, error)

// func userPolicySetGenerator(dataSource DataSource, securityPrincipal SecurityPrincipal) (map[interface{}]Policy, error) {
// 	userEntity, err := dataSource.GetRepository(User).FindBy(map[string]interface{}{"ciamId": securityPrincipal.ID})
// 	if err != nil {
// 		return nil, err
// 	}

// 	policySet := make(map[interface{}]Policy)
// 	policySet[User] = userPolicies(userEntity)
// 	policySet[Currency] = currencyPolicies(userEntity)
// 	policySet[Account] = accountPolicies(userEntity)
// 	return policySet, nil
// }

// func accountPolicies(userEntity *models.User) map[Action]Policy {
// 	// Check if the user exists
// 	canAccess := userEntity != nil

// 	accountSelectionRestrictions := FindOptionsWhere{
// 		OwnerID: userEntity.AccountOwnershipID,
// 		SubType: "POSTED",
// 	}

// 	accountFieldAccess := []AccountFields{}

// 	return map[Action]Policy{
// 		Read: {
// 			canAccess:             func() bool { return canAccess },
// 			selectionRestrictions: accountSelectionRestrictions,
// 			fieldAccess:           accountFieldAccess,
// 		},
// 		Create: {
// 			canAccess:             func() bool { return canAccess },
// 			selectionRestrictions: accountSelectionRestrictions,
// 			fieldAccess:           accountFieldAccess,
// 		},
// 		Delete: {
// 			canAccess: func(context Context) bool {
// 				if !canAccess {
// 					return false
// 				}
// 				entity := context.Entity.(*Account)
// 				return entity.OwnerID == userEntity.AccountOwnershipID
// 			},
// 			selectionRestrictions: accountSelectionRestrictions,
// 			fieldAccess:           accountFieldAccess,
// 		},
// 	}
// }
