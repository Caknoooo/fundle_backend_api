package repository

import (
	"context"
	"errors"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, event entities.Event) (entities.Event, error)
	GetAllEvent(ctx context.Context) ([]entities.Event, error)
	GetAllEventByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Event, error)
	GetEventByID(ctx context.Context, eventID uuid.UUID) (entities.Event, error)
	LikeEventByEventID(ctx context.Context, userID uuid.UUID, eventID uuid.UUID) error
	UpdateEvent(ctx context.Context, event entities.Event) error
	PatchEvent(ctx context.Context, event entities.Event, eventID uuid.UUID) error
	DeleteEvent(ctx context.Context, eventID uuid.UUID) error
}

type eventRepository struct {
	connection *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{
		connection: db,
	}
}

func (er *eventRepository) CreateEvent(ctx context.Context, event entities.Event) (entities.Event, error) {
	if err := er.connection.Preload("User").Create(&event).Error; err != nil {
		return entities.Event{}, nil
	}
	return event, nil
}

func (er *eventRepository) GetAllEvent(ctx context.Context) ([]entities.Event, error) {
	var events []entities.Event
	if err := er.connection.Preload("User").Preload("Likes").Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (er *eventRepository) GetAllEventByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Event, error) {
	var events []entities.Event
	if err := er.connection.Preload("User").Preload("Likes").Where("user_id = ?", userID).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (er *eventRepository) GetEventByID(ctx context.Context, eventID uuid.UUID) (entities.Event, error) {
	var event entities.Event
	if err := er.connection.Preload("User").Preload("Likes").Where("id = ?", eventID).Take(&event).Error; err != nil {
		return entities.Event{}, err
	}
	return event, nil
}

func (er *eventRepository) LikeEventByEventID(ctx context.Context, userID uuid.UUID, eventID uuid.UUID) error {
	var like entities.Like
	if err := er.connection.Where("user_id = ? AND event_id = ?", userID, eventID).Find(&like).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	if like.ID != uuid.Nil {
		return errors.New("User sudah melakukan like pada event ini")
	}

	like = entities.Like{
		EventID: eventID,
		UserID:  userID,
	}

	if err := er.connection.Create(&like).Error; err != nil {
		return err
	}

	var event entities.Event
	if err := er.connection.Where("id = ?", eventID).Find(&event).Error; err != nil {
		return err
	}

	event.LikeCount++
	er.UpdateEvent(ctx, event)

	return nil
}

func (er *eventRepository) UpdateEvent(ctx context.Context, event entities.Event) error {
	if err := er.connection.Updates(&event).Error; err != nil {
		return err
	}
	return nil
}

func (er *eventRepository) PatchEvent(ctx context.Context, event entities.Event, eventID uuid.UUID) error {
	if err := er.connection.Where("id = ?", eventID).Updates(&event).Error; err != nil {
		return err
	}
	return nil
}

func (er *eventRepository) DeleteEvent(ctx context.Context, eventID uuid.UUID) error {
	if err := er.connection.Delete(&entities.Event{}, &eventID).Error; err != nil {
		return err
	}
	return nil
}
