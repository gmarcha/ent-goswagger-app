package eventType

import (
	"context"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/eventtype"
	"github.com/google/uuid"
)

type Service struct {
	Type *ent.EventTypeClient
}

func (s *Service) ListType(ctx context.Context) ([]*ent.EventType, error) {

	res, err := s.Type.Query().WithEvents().All(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) CreateType(ctx context.Context, category *ent.EventType) (*ent.EventType, error) {

	builder := s.Type.Create()
	setEventType(builder.Mutation(), category)
	res, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) ReadTypeByID(ctx context.Context, id uuid.UUID) (*ent.EventType, error) {

	res, err := s.Type.Query().Where(eventtype.ID(id)).WithEvents().Only(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) UpdateTypeByID(ctx context.Context, id uuid.UUID, category *ent.EventType) (*ent.EventType, error) {

	builder := s.Type.UpdateOneID(id)
	setEventType(builder.Mutation(), category)
	res, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) DeleteTypeByID(ctx context.Context, id uuid.UUID) error {

	return s.Type.DeleteOneID(id).Exec(ctx)
}

func (s *Service) ListTypeEventsByID(ctx context.Context, id uuid.UUID) ([]*ent.Event, error) {

	res, err := s.ReadTypeByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res.Edges.Events, nil
}

func setEventType(m *ent.EventTypeMutation, category *ent.EventType) {

	m.SetName(category.Name)
	m.SetColor(category.Color)
}
