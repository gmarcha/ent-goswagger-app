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
}

func (s *Service) ListUser(ctx context.Context, tutor *bool) ([]*ent.User, error) {

	builder := s.User.Query()
	if tutor != nil && *tutor {
		builder.WithRoles(func(q *ent.RoleQuery) {
			q.Where(role.Name("tutor"))
		})
	}
	res, err := builder.All(ctx)
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

func (s *Service) ReadUserByLogin(ctx context.Context, login string) (*ent.User, error) {

	res, err := s.User.Query().Where(user.Login(login)).WithRoles().Only(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) UpdateUserByLogin(ctx context.Context, login string, userInfo *ent.User) (*ent.User, error) {

	builder := s.User.UpdateOne(userInfo)
	setUser(builder.Mutation(), userInfo)
	res, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) DeleteUserByLogin(ctx context.Context, login string) error {

	_, err := s.User.Delete().Where(user.Login(login)).Exec(ctx)
	return err
}

func (s *Service) ListUserEventsByLogin(ctx context.Context, login string) ([]*ent.Event, error) {

	res, err := s.ReadUserByLogin(ctx, login)
	if err != nil {
		return nil, err
	}
	return res.Edges.Events, err
}

func (s *Service) SubscribeUserByLogin(ctx context.Context, login string, id uuid.UUID) ([]*ent.Event, error) {

	user, err := s.ReadUserByLogin(ctx, login)
	if err != nil {
		return nil, err
	}
	res, err := user.Update().AddEventIDs(id).Save(ctx)
	if err != nil {
		return nil, err
	}
	return res.Edges.Events, err
}

func (s *Service) UnsubscribeUserByLogin(ctx context.Context, login string, id uuid.UUID) error {

	user, err := s.ReadUserByLogin(ctx, login)
	if err != nil {
		return err
	}
	_, err = user.Update().RemoveEventIDs(id).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ReadUserByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {

	res, err := s.User.Query().Where(user.ID(id)).WithRoles().Only(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) UpdateUserByID(ctx context.Context, id uuid.UUID, user *ent.User) (*ent.User, error) {

	builder := s.User.UpdateOneID(id)
	setUser(builder.Mutation(), user)
	res, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) DeleteUserByID(ctx context.Context, id uuid.UUID) error {

	return s.User.DeleteOneID(id).Exec(ctx)
}

func (s *Service) ListUserEventsByID(ctx context.Context, id uuid.UUID) ([]*ent.Event, error) {

	res, err := s.ReadUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res.Edges.Events, err
}

func (s *Service) SubscribeUserByID(ctx context.Context, userId uuid.UUID, eventId uuid.UUID) ([]*ent.Event, error) {

	user, err := s.ReadUserByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	res, err := user.Update().AddEventIDs(eventId).Save(ctx)
	if err != nil {
		return nil, err
	}
	return res.Edges.Events, err
}

func (s *Service) UnsubscribeUserByID(ctx context.Context, userId uuid.UUID, eventId uuid.UUID) error {

	user, err := s.ReadUserByID(ctx, userId)
	if err != nil {
		return err
	}
	_, err = user.Update().RemoveEventIDs(eventId).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SetUserOnLogin(ctx context.Context, user *ent.User) (*ent.User, error) {

	res, err := s.ReadUserByLogin(ctx, user.Login)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}
		res, err = s.CreateUser(ctx, user)
		if err != nil {
			return nil, err
		}
	} else {
		res, err = s.UpdateUserByID(ctx, res.ID, user)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func setUser(m *ent.UserMutation, u *ent.User) {

	m.SetFirstName(u.FirstName)
	m.SetLastName(u.LastName)
	m.SetImagePath(u.ImagePath)
}
