package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type EventService interface {
	CreateEvent(ctx context.Context, eventDTO dto.EventCreateDTO) (entities.Event, error)
	GetAllEvent(ctx context.Context) ([]entities.Event, error)
	GetAllEventByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Event, error)
	GetEventByID(ctx context.Context, eventID uuid.UUID) (entities.Event, error)
	LikeEventByEventID(ctx context.Context, userID uuid.UUID, eventID uuid.UUID) error
	UpdateEvent(ctx context.Context, eventDTO dto.EventUpdateDTO, eventID uuid.UUID) error
	PatchEvent(ctx context.Context, eventDTO dto.EventUpdateDTO, eventID uuid.UUID) error
	DeleteEvent(ctx context.Context, eventID uuid.UUID) error
}

type eventService struct {
	eventRepository repository.EventRepository
}

func NewEventService(er repository.EventRepository) EventService {
	return &eventService{
		eventRepository: er,
	}
}

func (es *eventService) CreateEvent(ctx context.Context, eventDTO dto.EventCreateDTO) (entities.Event, error) {
	event := entities.Event{}
	err := smapping.FillStruct(&event, smapping.MapFields(eventDTO))
	if err != nil {
		return entities.Event{}, err
	}
	return es.eventRepository.CreateEvent(ctx, event)
}

func (es *eventService) GetAllEvent(ctx context.Context) ([]entities.Event, error) {
	return es.eventRepository.GetAllEvent(ctx)
}

func (es *eventService) GetAllEventByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Event, error) {
	return es.eventRepository.GetAllEventByUserID(ctx, userID)
}

func (es *eventService) GetEventByID(ctx context.Context, eventID uuid.UUID) (entities.Event, error) {
	return es.eventRepository.GetEventByID(ctx, eventID)
}

func (es *eventService) LikeEventByEventID(ctx context.Context, userID uuid.UUID, eventID uuid.UUID) error {
	return es.eventRepository.LikeEventByEventID(ctx, userID, eventID)
}

func (es *eventService) UpdateEvent(ctx context.Context, eventDTO dto.EventUpdateDTO, eventID uuid.UUID) error {
	event := entities.Event{}
	if err := smapping.FillStruct(&event, smapping.MapFields(eventDTO)); err != nil {
		return nil
	}
	return es.eventRepository.UpdateEvent(ctx, event, eventID)
}

func (es *eventService) PatchEvent(ctx context.Context, eventDTO dto.EventUpdateDTO, eventID uuid.UUID) error {
	event := entities.Event{}
	if err := smapping.FillStruct(&event, smapping.MapFields(eventDTO)); err != nil {
		return nil
	}
	return es.eventRepository.PatchEvent(ctx, event, eventID)
}

func (es *eventService) DeleteEvent(ctx context.Context, eventID uuid.UUID) error {
	return es.eventRepository.DeleteEvent(ctx, eventID)
}
