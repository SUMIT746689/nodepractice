package userrepo

import (
	"context"
	"pos/ent"
	entuser "pos/ent/user"
	"pos/internal/app"
	"pos/pkg"
)

func Save(ctx context.Context, u *app.User) (*ent.User, error) {
	createQuery := pkg.EntClient().User.Create().SetFirstName(u.FirstName).SetLastName(u.LastName).SetUsername(u.Username).
		SetPassword(pkg.Hash(u.Password)).SetRole(entuser.Role(u.Role))

	if u.PhoneNumber != "" {
		createQuery.SetPhoneNumber(u.PhoneNumber)
	}

	if u.Email != "" {
		createQuery.SetEmail(u.Email)
	}

	savedUser, err := createQuery.Save(ctx)
	if err != nil {
		return nil, err
	}

	return savedUser, nil
}
