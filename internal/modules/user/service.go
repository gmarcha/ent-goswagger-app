package user

import (
	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type Service struct {
	User *ent.UserClient
}

func (s *Service) listUser(ctx context.Context) ([]*ent.User, error) {
	res, err := s.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) createUser(ctx context.Context, user *ent.User) (*ent.User, error) {
	builder := s.User.Create()
	setUser(builder.Mutation(), user)
	res, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) readUser(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	res, err := s.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) updateUser(ctx context.Context, id uuid.UUID, user *ent.User) (*ent.User, error) {
	builder := s.User.UpdateOneID(id)
	setUser(builder.Mutation(), user)
	res, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) deleteUser(ctx context.Context, id uuid.UUID) error {
	return s.User.DeleteOneID(id).Exec(ctx)
}

func setUser(m *ent.UserMutation, u *ent.User) {
	m.SetLogin(u.Login)
	m.SetFirstName(u.FirstName)
	m.SetLastName(u.LastName)
	m.SetImagePath(u.ImagePath)
	m.SetCalendarScope(u.CalendarScope)
	m.SetAdminScope(u.AdminScope)
}
