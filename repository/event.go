package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

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
	UpdateEvent(ctx context.Context, event entities.Event, eventID uuid.UUID) error
	PatchEvent(ctx context.Context, event entities.Event, eventID uuid.UUID) error
	DeleteEvent(ctx context.Context, eventID uuid.UUID) error
	Get4Event(ctx context.Context) ([]entities.Event, error)
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
	if err := er.connection.Create(&event).Error; err != nil {
		return entities.Event{}, nil
	}

	// if err := er.connection.Preload("Event")
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
	er.UpdateEvent(ctx, event, like.EventID)

	return nil
}

func (er *eventRepository) UpdateEvent(ctx context.Context, event entities.Event, eventID uuid.UUID) error {
	if err := er.connection.Where("id = ?", eventID).Updates(&event).Error; err != nil {
		return err
	}

	return nil
}

func (er *eventRepository) PatchEvent(ctx context.Context, event entities.Event, eventID uuid.UUID) error {
	if err := er.connection.Where("id = ?", eventID).Updates(&event).Error; err != nil {
		return err
	}

	// Ambil event yang telah diperbarui dari database
	var updatedEvent entities.Event
	if err := er.connection.First(&updatedEvent, eventID).Error; err != nil {
		return err
	}

	fmt.Printf("%+v\n", updatedEvent)

	// Cek apakah event telah kadaluwarsa
	if !updatedEvent.ExpiredDonasi.IsZero() && updatedEvent.ExpiredDonasi.Before(time.Now()) {
		updatedEvent.Is_expired = true
		if err := er.connection.Save(&updatedEvent).Error; err != nil {
			return err
		}
		return errors.New("Event Has Expired")
	}

	if updatedEvent.JumlahDonasi >= updatedEvent.MaxDonasi {
		updatedEvent.Is_target_full = true
	}

	// Tambahkan IsDone dan update lagi
	updatedEvent.TotalBerhasil++
	if err := er.connection.Save(&updatedEvent).Error; err != nil {
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

func (er *eventRepository) Get4Event(ctx context.Context) ([]entities.Event, error) {
	var events []entities.Event
	if err := er.connection.Preload("User").Preload("Likes").Limit(3).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}
