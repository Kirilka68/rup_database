package models

import (
	"encoding/json"
	"time"
)

type Object struct {
	ID                 string    `json:"id"`
	Title              *string   `json:"title,omitempty"`
	Year               *string   `json:"year,omitempty"`
	Type               *string   `json:"type,omitempty"`
	Course             *string   `json:"course,omitempty"`
	StudyForm          *string   `json:"studyForm,omitempty"`
	ExplorationNote    *string   `json:"explorationNote,omitempty"`
	Goals              *string   `json:"goals,omitempty"`
	Prerequisites      *string   `json:"prerequisites,omitempty"`
	Postrequisites     *string   `json:"postrequisites,omitempty"`
	GroupID            *string   `json:"groupId,omitempty"`
	Contacts           *string   `json:"contacts,omitempty"`
	InternetResources  *string   `json:"internetResources,omitempty"`
	TechnicalResources *string   `json:"technicalResources,omitempty"`
	RupType            *string   `json:"rupType,omitempty"`
	Disciplines        *string   `json:"disciplines,omitempty"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type CreateObjectDTO struct {
	Title              json.RawMessage `json:"title"`
	Year               json.RawMessage `json:"year"`
	Type               json.RawMessage `json:"type"`
	Course             json.RawMessage `json:"course"`
	StudyForm          json.RawMessage `json:"studyForm"`
	ExplorationNote    json.RawMessage `json:"explorationNote"`
	Goals              json.RawMessage `json:"goals"`
	Prerequisites      json.RawMessage `json:"prerequisites"`
	Postrequisites     json.RawMessage `json:"postrequisites"`
	GroupID            json.RawMessage `json:"groupId"`
	Contacts           json.RawMessage `json:"contacts"`
	InternetResources  json.RawMessage `json:"internetResources"`
	TechnicalResources json.RawMessage `json:"technicalResources"`
	RupType            json.RawMessage `json:"rupType"`
	Disciplines        json.RawMessage `json:"disciplines"`
}

type UpdateObjectDTO struct {
	ID                 json.RawMessage `json:"id"`
	Title              json.RawMessage `json:"title"`
	Year               json.RawMessage `json:"year"`
	Type               json.RawMessage `json:"type"`
	Course             json.RawMessage `json:"course"`
	StudyForm          json.RawMessage `json:"studyForm"`
	ExplorationNote    json.RawMessage `json:"explorationNote"`
	Goals              json.RawMessage `json:"goals"`
	Prerequisites      json.RawMessage `json:"prerequisites"`
	Postrequisites     json.RawMessage `json:"postrequisites"`
	GroupID            json.RawMessage `json:"groupId"`
	Contacts           json.RawMessage `json:"contacts"`
	InternetResources  json.RawMessage `json:"internetResources"`
	TechnicalResources json.RawMessage `json:"technicalResources"`
	RupType            json.RawMessage `json:"rupType"`
	Disciplines        json.RawMessage `json:"disciplines"`
}

type GetObjectDTO struct {
	ID string `json:"id"`
}

type DeleteObjectDTO struct {
	ID string `json:"id"`
}
