package user

import (
	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/role"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/user"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type Service struct {
	User *ent.UserClient
	Role *ent.RoleClient
}

func (s *Service) ListUser(ctx context.Context, tutor *bool) ([]*ent.User, error) {

	var res []*ent.User
	var err error

	if tutor != nil && *tutor {
		res, err = s.Role.Query().Where(role.Name("tutor")).QueryUsers().All(ctx)
	} else {
		res, err = s.User.Query().All(ctx)
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) CreateUser(ctx context.Context, user *ent.User) (*ent.User, error) {
	builder := s.User.Create().SetLogin(user.Login)
	setUser(builder.Mutation(), user)
	res, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) ReadUser(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	res, err := s.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) UpdateUser(ctx context.Context, id uuid.UUID, user *ent.User) (*ent.User, error) {
	builder := s.User.UpdateOneID(id)
	setUser(builder.Mutation(), user)
	res, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.User.DeleteOneID(id).Exec(ctx)
}

func (s *Service) ReadMe(ctx context.Context, login string) (*ent.User, error) {
	res, err := s.User.Query().Where(user.Login(login)).WithRoles().Only(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func setUser(m *ent.UserMutation, u *ent.User) {
	m.SetFirstName(u.FirstName)
	m.SetLastName(u.LastName)
	m.SetImagePath(u.ImagePath)
}
