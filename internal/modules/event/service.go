package event

import (
	"context"
	"time"

	"github.com/gmarcha/ent-goswagger-app/internal/ent"
	entEvent "github.com/gmarcha/ent-goswagger-app/internal/ent/event"
	"github.com/google/uuid"
)

// Service holds an ent event client.
type Service struct {
	Event *ent.EventClient
}

// ListEvent returns an event list with category, or an error.
func (s *Service) ListEvent(ctx context.Context, start, end *string) ([]*ent.Event, error) {

	builder := s.Event.Query().WithCategory()
	if start != nil {
		filter, err := time.Parse(time.RFC3339, *start)
		if err != nil {
			return nil, err
		}
		builder.Where(entEvent.StartAtGTE(filter))
	} else if end != nil {
		filter, err := time.Parse(time.RFC3339, *end)
		if err != nil {
			return nil, err
		}
		builder.Where(entEvent.EndAtLTE(filter))
	}
	res, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ListEventWithUsers returns an event list with category and users, or an error.
func (s *Service) ListEventWithUsers(ctx context.Context, start, end *string) ([]*ent.Event, error) {

	builder := s.Event.Query().WithCategory().WithUsers()
	if start != nil {
		filter, err := time.Parse(time.RFC3339, *start)
		if err != nil {
			return nil, err
		}
		builder.Where(entEvent.StartAtGTE(filter))
	} else if end != nil {
		filter, err := time.Parse(time.RFC3339, *end)
		if err != nil {
			return nil, err
		}
		builder.Where(entEvent.EndAtLTE(filter))
	}
	res, err := builder.All(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateEvent returns a new created event or an error.
func (s *Service) CreateEvent(ctx context.Context, event *ent.Event) (*ent.Event, error) {

	builder := s.Event.Create()
	setEvent(builder.Mutation(), event)
	res, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ReadEventByID returns an event or an error.
func (s *Service) ReadEventByID(ctx context.Context, id uuid.UUID) (*ent.Event, error) {

	res, err := s.Event.Query().Where(entEvent.ID(id)).WithCategory().WithUsers().Only(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateEventByID returns an existing updated event or an error.
func (s *Service) UpdateEventByID(ctx context.Context, id uuid.UUID, event *ent.Event) (*ent.Event, error) {

	builder := s.Event.UpdateOneID(id)
	setEvent(builder.Mutation(), event)
	res, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteEventByID returns an error.
func (s *Service) DeleteEventByID(ctx context.Context, id uuid.UUID) error {

	return s.Event.DeleteOneID(id).Exec(ctx)
}

// ListEventUsersByID returns a list of subscribed users to an event or an error.
func (s *Service) ListEventUsersByID(ctx context.Context, id uuid.UUID) ([]*ent.User, error) {

	res, err := s.ReadEventByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return res.Edges.Users, nil
}

// ReadEventTypeByID returns the type of an event or an error.
func (s *Service) ReadEventTypeByID(ctx context.Context, id uuid.UUID) (*ent.EventType, error) {

	event, err := s.ReadEventByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return event.Edges.Category, nil
}

// UpdateEventTypeByID returns an event with type updated or an error.
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
