package internal

import "net/http"

// Handler type implements handlers
type Handler struct {
	calendar Calendar
}

// InitRoutes инициализируем хандлер
func (h *Handler) InitRoutes(calendar Calendar) {

	h.calendar = calendar

	//Handle регистрирует и сопостовляет патерн с обработчиком который передаем
	//HandlerFunc это адаптер, позволяющий использовать обычные функции в качестве обработчиков HTTP.
	//Если f — функция с соответствующей сигнатурой, HandlerFunc(f) — это обработчик, вызывающий f.
	http.Handle("/create_event", Logging(http.HandlerFunc(h.CreateEvent)))
	http.Handle("/update_event", Logging(http.HandlerFunc(h.UpdateEvent)))
	http.Handle("/delete_event", Logging(http.HandlerFunc(h.DeleteEvent)))
	http.Handle("/events_for_day", Logging(http.HandlerFunc(h.EventsPerDay)))
	http.Handle("/events_for_week", Logging(http.HandlerFunc(h.EventsPerWeek)))
	http.Handle("/events_for_month", Logging(http.HandlerFunc(h.EventsPerMonth)))
}
