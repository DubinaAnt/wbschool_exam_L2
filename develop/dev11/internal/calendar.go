package internal

import (
	"sync"
	"time"
)

// Calendar интерфейс
type Calendar interface {
	CreateEvent(ev Event)
	UpdateEvent(ev Event)
	DeleteEvent(eventId int)
	GetEventPerDay(date time.Time) ([]Event, error)
	GetEventPerWeek(date time.Time) ([]Event, error)
	GetEventPerMonth(date time.Time) ([]Event, error)
}

// Calendar структура имплементящая интерфейс
type calendar struct {
	events  map[int]Event
	idCount int
	sync.RWMutex
}

// NewCalendar создает календарь
func NewCalendar() Calendar {
	return &calendar{
		events:  make(map[int]Event),
		idCount: 0,
	}
}

// CreateEvent создаем эвент
func (c *calendar) CreateEvent(event Event) {
	c.Lock()
	event.ID = c.idCount
	c.events[event.ID] = event
	c.idCount++
	c.Unlock()
}

// DeleteEvent удаляем эвент
func (c *calendar) DeleteEvent(eventId int) {
	c.Lock()
	delete(c.events, eventId)
	c.Unlock()
}

// UpdateEvent апдейтим эвент
func (c *calendar) UpdateEvent(event Event) {
	c.Lock()
	c.events[event.ID] = event
	c.Unlock()
}

// GetEventPerDay выдает эвенты за день
func (c *calendar) GetEventPerDay(date time.Time) ([]Event, error) {
	result := make([]Event, 0)

	for _, e := range c.events {
		if e.Date == date {
			result = append(result, e)
		}
	}
	return result, nil
}

// GetEventPerWeek выдает эвенты за неделю от даты пришедшей
func (c *calendar) GetEventPerWeek(date time.Time) ([]Event, error) {
	result := make([]Event, 0)

	dateE := date.AddDate(0, 0, 7)

	for _, e := range c.events {
		if e.Date.After(date) && e.Date.Before(dateE) {
			result = append(result, e)
		}
	}
	return result, nil
}

// GetEventPerMonth выдает эвенты за месяц от даты пришедшей
func (c *calendar) GetEventPerMonth(date time.Time) ([]Event, error) {
	result := make([]Event, 0)

	dateE := date.AddDate(0, 1, 0)

	for _, e := range c.events {
		if e.Date.After(date) && e.Date.Before(dateE) {
			result = append(result, e)
		}
	}
	return result, nil
}
