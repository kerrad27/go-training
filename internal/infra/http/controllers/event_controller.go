package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/test_server/internal/domain/event"
)

type EventController struct {
	service *event.Service
}

func NewEventController(s *event.Service) *EventController {
	return &EventController{
		service: s,
	}
}

func (c *EventController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := (*c.service).FindAll()
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindAll(): %s", err)
			}
			return
		}

		err = success(w, events)
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
		}
	}
}

func (c *EventController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}
		event, err := (*c.service).FindOne(id)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}

		err = success(w, event)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
		}
	}
}
func (c *EventController) Add() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := chi.URLParam(r, "title")
		description := chi.URLParam(r, "description")
		year, err := strconv.ParseInt(chi.URLParam(r, "year"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.Add(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Add(): %s", err)
			}
			return
		}

		info, err := (*c.service).Add(title, description, year)
		if err != nil {
			fmt.Printf("EventController.Add(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Add(): %s", err)
			}
			return
		}

		err = success(w, info)
		if err != nil {
			fmt.Printf("EventController.Add(): %s", err)
		}
	}
}

func (c *EventController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.Update(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Update(): %s", err)
			}
			return

		}
		title := chi.URLParam(r, "title")
		description := chi.URLParam(r, "description")
		year, err := strconv.ParseInt(chi.URLParam(r, "year"), 10, 64)

		info, err := (*c.service).Update(id, title, description, year)
		if err != nil {
			fmt.Printf("EventController.Update(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Update(): %s", err)
			}
			return
		}

		err = success(w, info)
		if err != nil {
			fmt.Printf("EventController.UpdateMovie(): %s", err)
		}
	}
}
