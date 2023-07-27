package pkg

import (
	"context"
	"pos/ent"
	"pos/ent/user"
	"pos/internal/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func UserHasPermission(userID int, permission domain.USER_PERMISSION) bool {
	// TODO: use context
	usr, err := EntClient().User.Query().Where(user.ID(userID)).WithPermissions().WithRole(func(query *ent.RoleQuery) {
		query.WithPermissions()
	}).First(context.Background())
	if err != nil {
		return false
	}

	for _, p := range usr.Edges.Permissions {
		if p.Value == string(permission) {
			return true
		}
	}

	for _, p := range usr.Edges.Role.Edges.Permissions {
		if p.Value == string(permission) {
			return true
		}
	}

	return false
}

func UserHasAllPermissions(userID int, permissions []domain.USER_PERMISSION) bool {
	// TODO: use context
	usr, err := EntClient().User.Query().Where(user.ID(userID)).WithPermissions().WithRole(func(query *ent.RoleQuery) {
		query.WithPermissions()
	}).First(context.Background())
	if err != nil {
		return false
	}

	// flat all permissions
	var allPermissions []string

	for _, p := range usr.Edges.Permissions {
		allPermissions = append(allPermissions, p.Value)
	}

	for _, p := range usr.Edges.Role.Edges.Permissions {
		allPermissions = append(allPermissions, p.Value)
	}
	println("allPermissions", allPermissions)
	println("rePermissions", permissions)

	for _, requiredPermission := range permissions {
		var isVerify = false
		for _, permission := range allPermissions {
			if permission == string(requiredPermission) {
				isVerify = true
			}
		}
		if !isVerify {
			return false
		}
	}
	return true
}

func GetAuthedUser(c *fiber.Ctx) (int, int) {
	jwtToken := c.Locals("user").(*jwt.Token)
	claims := jwtToken.Claims.(jwt.MapClaims)
	id := int(claims["id"].(float64))
	roleID := int(claims["role_id"].(float64))
	return id, roleID
}
