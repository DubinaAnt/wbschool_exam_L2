package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

// CreateEvent создает эвент
// w http.ResponseWriter, r *http.Request передаем чтобы соответсвтовало сигнатуре из хандлер и чтобы взаимодействовать
// с запросом
func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	Logging(r)

	if r.Method != http.MethodPost { //оставляем обработку только пост метода
		w.WriteHeader(http.StatusBadRequest) //пишем в статус ответа 400//такой вызов можно пропустить если все ок
		//и он напишет 200
		w.Write([]byte(`{"error": "bad request"}`))

		return //выходим из метода
	}

	var event Event

	body, err := io.ReadAll(r.Body) //читаем боди, вроде ioutil.ReadAll с 1.16 вызывает io.ReadAll поэтому его использовал
	parsJson := make(map[string]string)
	json.Unmarshal(body, &parsJson) // парсим боди мапу
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Bad request - incorrect body"}`))
		return
	}

	event.UserID, err = strconv.Atoi(parsJson["user_id"]) //парсим айди из строки в инт
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Bad request - incorrect user id"}`))
		return
	}

	event.Date, err = time.Parse("2006-01-02", parsJson["date"]) //парсим по шаблону из строки в дату
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Bad request - incorrect date"}`))
		return
	}

	event.Event = parsJson["event"]
	h.calendar.CreateEvent(event)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"result": "Created"}`))
}

// UpdateEvent обновление календаря
func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	Logging(r)

	if r.Method != http.MethodPost { //оставляем обработку только пост метода
		w.WriteHeader(http.StatusBadRequest) //пишем в статус ответа 400//такой вызов можно пропустить если все ок
		//и он напишет 200
		w.Write([]byte(`{"error": "bad request"}`))

		return //выходим из метода
	}

	var event Event

	body, err := io.ReadAll(r.Body) //читаем боди, вроде ioutil.ReadAll с 1.16 вызывает io.ReadAll поэтому его использовал
	parsJson := make(map[string]string)
	json.Unmarshal(body, &parsJson) // парсим боди мапу
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Bad request - incorrect body"}`))
		return
	}

	event.UserID, err = strconv.Atoi(parsJson["user_id"]) //парсим айди из строки в инт
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Bad request - incorrect user id"}`))
		return
	}

	event.Date, err = time.Parse("2006-01-02", parsJson["date"]) //парсим по шаблону из строки в дату
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Bad request - incorrect date"}`))
		return
	}

	event.Event = parsJson["event"]
	h.calendar.CreateEvent(event)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"result": "Created"}`))
}

// DeleteEvent удаляет эвент по айди
func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	Logging(r)

	if r.Method != http.MethodPost { //оставляем обработку только пост метода
		w.WriteHeader(http.StatusBadRequest) //пишем в статус ответа 400//такой вызов можно пропустить если все ок
		//и он напишет 200
		w.Write([]byte(`{"error": "bad request"}`))

		return //выходим из метода
	}

	var eventId int

	body, err := io.ReadAll(r.Body) //читаем боди, вроде ioutil.ReadAll с 1.16 вызывает io.ReadAll поэтому его использовал
	parsJson := make(map[string]string)
	json.Unmarshal(body, &parsJson) // парсим боди мапу
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Bad request - incorrect body"}`))
		return
	}

	eventId, err = strconv.Atoi(parsJson["user_id"]) //парсим айди из строки в инт
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Bad request - incorrect user id"}`))
		return
	}

	h.calendar.DeleteEvent(eventId)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"result": "Deleted"}`))
}

// EventsPerDay возвращает эвент за день
func (h *Handler) EventsPerDay(w http.ResponseWriter, r *http.Request) {
	Logging(r)

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request"}`))
		return
	}

	param := r.URL.Query()
	dateStr := param.Get("data")

	date, err := time.Parse("2006-01-02", dateStr) //парсим по шаблону из строки в дату
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Bad request - incorrect date"}`))
		return
	}

	events, err := h.calendar.GetEventPerDay(date)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 в случае ошибки логики
		w.Write([]byte(`{"error": "Server BL error"}`))

		return
	}

	if len(events) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Event not found"}`))

		return
	}

	resJson, err := json.Marshal(events)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 в случае ошибки логики
		w.Write([]byte(`{"error": "Server BL error"}`))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resJson)
}

// EventsPerWeek эвенты за неделю
func (h *Handler) EventsPerWeek(w http.ResponseWriter, r *http.Request) {
	Logging(r)

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request"}`))
		return
	}

	param := r.URL.Query()
	dateStr := param.Get("data")

	date, err := time.Parse("2006-01-02", dateStr) //парсим по шаблону из строки в дату
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Bad request - incorrect date"}`))
		return
	}

	events, err := h.calendar.GetEventPerWeek(date)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 в случае ошибки логики
		w.Write([]byte(`{"error": "Server BL error"}`))

		return
	}

	if len(events) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Event not found"}`))

		return
	}

	resJson, err := json.Marshal(events)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 в случае ошибки логики
		w.Write([]byte(`{"error": "Server BL error"}`))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resJson)
}

// EventsPerMonth эвенты за месяц
func (h *Handler) EventsPerMonth(w http.ResponseWriter, r *http.Request) {
	Logging(r)

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request"}`))
		return
	}

	param := r.URL.Query()
	dateStr := param.Get("data")

	date, err := time.Parse("2006-01-02", dateStr) //парсим по шаблону из строки в дату
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Bad request - incorrect date"}`))
		return
	}

	events, err := h.calendar.GetEventPerMonth(date)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 в случае ошибки логики
		w.Write([]byte(`{"error": "Server BL error"}`))

		return
	}

	if len(events) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Event not found"}`))

		return
	}

	resJson, err := json.Marshal(events)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable) //503 в случае ошибки логики
		w.Write([]byte(`{"error": "Server BL error"}`))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resJson)
}
