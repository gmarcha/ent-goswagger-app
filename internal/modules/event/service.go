package event

import (
	"context"
	"time"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	entEvent "github.com/gmarcha/ent-goswagger-app/internal/ent/event"
	"github.com/google/uuid"
)

type Service struct {
	Event *ent.EventClient
}

func (s *Service) ListEvent(ctx context.Context, day, month *string) ([]*ent.Event, error) {

	builder := s.Event.Query().WithCategory()
	if day != nil {
		filter, err := time.Parse(time.RFC3339, *day)
		if err != nil {
			return nil, err
		}
		builder.Where(entEvent.Or(
			entEvent.And(
				entEvent.StartAtGT(time.Date(filter.Year(), filter.Month(), filter.Day(), 0, 0, 0, 0, filter.Location())),
				entEvent.StartAtLT(time.Date(filter.Year(), filter.Month(), filter.Day()+1, 0, 0, 0, 0, filter.Location())),
			),
			entEvent.And(
				entEvent.EndAtGT(time.Date(filter.Year(), filter.Month(), filter.Day(), 0, 0, 0, 0, filter.Location())),
				entEvent.EndAtLT(time.Date(filter.Year(), filter.Month(), filter.Day()+1, 0, 0, 0, 0, filter.Location())),
			),
		))
	} else if month != nil {
		filter, err := time.Parse(time.RFC3339, *month)
		if err != nil {
			return nil, err
		}
		builder.Where(entEvent.Or(
			entEvent.And(
				entEvent.StartAtGT(time.Date(filter.Year(), filter.Month(), 0, 0, 0, 0, 0, filter.Location())),
				entEvent.StartAtLT(time.Date(filter.Year(), filter.Month()+1, 0, 0, 0, 0, 0, filter.Location())),
			),
			entEvent.And(
				entEvent.EndAtGT(time.Date(filter.Year(), filter.Month(), 0, 0, 0, 0, 0, filter.Location())),
				entEvent.EndAtLT(time.Date(filter.Year(), filter.Month()+1, 0, 0, 0, 0, 0, filter.Location())),
			),
		))
	}
	res, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) CreateEvent(ctx context.Context, event *ent.Event) (*ent.Event, error) {

	builder := s.Event.Create()
	setEvent(builder.Mutation(), event)
	res, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) ReadEventByID(ctx context.Context, id uuid.UUID) (*ent.Event, error) {

	res, err := s.Event.Query().Where(entEvent.ID(id)).WithCategory().WithUsers().Only(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) UpdateEventByID(ctx context.Context, id uuid.UUID, event *ent.Event) (*ent.Event, error) {

	builder := s.Event.UpdateOneID(id)
	setEvent(builder.Mutation(), event)
	res, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) DeleteEventByID(ctx context.Context, id uuid.UUID) error {

	return s.Event.DeleteOneID(id).Exec(ctx)
}

func (s *Service) ListEventUsersByID(ctx context.Context, id uuid.UUID) ([]*ent.User, error) {

	res, err := s.ReadEventByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res.Edges.Users, nil
}

func (s *Service) ReadEventTypeByID(ctx context.Context, id uuid.UUID) (*ent.EventType, error) {

	event, err := s.ReadEventByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return event.Edges.Category, nil
}

func (s *Service) UpdateEventTypeByID(ctx context.Context, eventID uuid.UUID, typeID uuid.UUID) (*ent.Event, error) {

	event, err := s.ReadEventByID(ctx, eventID)
	if err != nil {
		return nil, err
	}
	res, err := event.Update().SetCategoryID(typeID).Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, err
}

func setEvent(m *ent.EventMutation, event *ent.Event) {

	m.SetName(event.Name)
	m.SetDescription(event.Description)
	if event.TutorsRequired != nil {
		m.SetTutorsRequired(*event.TutorsRequired)
	}
	if event.WalletsReward != nil {
		m.SetWalletsReward(*event.WalletsReward)
	}
	m.SetStartAt(event.StartAt)
	m.SetEndAt(event.EndAt)
}
