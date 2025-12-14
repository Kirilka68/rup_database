package repository

import (
	"context"
	"database/sql"
	"rup_database/internal/models"
)

type ObjectRepository interface {
	Create(ctx context.Context, input models.CreateObjectDTO) (*models.Object, error)
	GetByID(ctx context.Context, id string) (*models.Object, error)
	List(ctx context.Context) ([]*models.Object, error)
	Update(ctx context.Context, input models.UpdateObjectDTO) (*models.Object, error)
	Delete(ctx context.Context, id string) error
}

type ObjectRepo struct {
	db *sql.DB
}

func NewObjectRepo(db *sql.DB) *ObjectRepo {
	return &ObjectRepo{db: db}
}

func (r *ObjectRepo) Create(ctx context.Context, input models.CreateObjectDTO) (*models.Object, error) {
	query := `
        INSERT INTO objects (title, year, type, course, study_form, exploration_note, goals, prerequisites, postrequisites, group_id, contacts, internet_resources, technical_resources, rup_type, disciplines)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
        RETURNING id, title, year, type, course, study_form, exploration_note, goals, prerequisites, postrequisites, group_id, contacts, internet_resources, technical_resources, rup_type, disciplines, created_at, updated_at
    `

	obj := &models.Object{}
	var title, year, typ, course, studyForm, explorationNote, goals, prerequisites, postrequisites, groupID, contacts, internetResources, technicalResources, rupType, disciplines interface{}
	if len(input.Title) > 0 {
		title = string(input.Title)
	}
	if len(input.Year) > 0 {
		year = string(input.Year)
	}
	if len(input.Type) > 0 {
		typ = string(input.Type)
	}
	if len(input.Course) > 0 {
		course = string(input.Course)
	}
	if len(input.StudyForm) > 0 {
		studyForm = string(input.StudyForm)
	}
	if len(input.ExplorationNote) > 0 {
		explorationNote = string(input.ExplorationNote)
	}
	if len(input.Goals) > 0 {
		goals = string(input.Goals)
	}
	if len(input.Prerequisites) > 0 {
		prerequisites = string(input.Prerequisites)
	}
	if len(input.Postrequisites) > 0 {
		postrequisites = string(input.Postrequisites)
	}
	if len(input.GroupID) > 0 {
		groupID = string(input.GroupID)
	}
	if len(input.Contacts) > 0 {
		contacts = string(input.Contacts)
	}
	if len(input.InternetResources) > 0 {
		internetResources = string(input.InternetResources)
	}
	if len(input.TechnicalResources) > 0 {
		technicalResources = string(input.TechnicalResources)
	}
	if len(input.RupType) > 0 {
		rupType = string(input.RupType)
	}
	if len(input.Disciplines) > 0 {
		disciplines = string(input.Disciplines)
	}

	err := r.db.QueryRowContext(ctx, query, title, year, typ, course, studyForm, explorationNote, goals, prerequisites, postrequisites, groupID, contacts, internetResources, technicalResources, rupType, disciplines).
		Scan(&obj.ID, &obj.Title, &obj.Year, &obj.Type, &obj.Course, &obj.StudyForm, &obj.ExplorationNote, &obj.Goals, &obj.Prerequisites, &obj.Postrequisites, &obj.GroupID, &obj.Contacts, &obj.InternetResources, &obj.TechnicalResources, &obj.RupType, &obj.Disciplines, &obj.CreatedAt, &obj.UpdatedAt)

	return obj, err
}

func (r *ObjectRepo) GetByID(ctx context.Context, id string) (*models.Object, error) {
	query := `
        SELECT id, title, year, type, course, study_form, exploration_note, goals, prerequisites, postrequisites, group_id, contacts, internet_resources, technical_resources, rup_type, disciplines, created_at, updated_at
        FROM objects
        WHERE id = $1
    `
	obj := &models.Object{}
	err := r.db.QueryRowContext(ctx, query, id).
		Scan(&obj.ID, &obj.Title, &obj.Year, &obj.Type, &obj.Course, &obj.StudyForm, &obj.ExplorationNote, &obj.Goals, &obj.Prerequisites, &obj.Postrequisites, &obj.GroupID, &obj.Contacts, &obj.InternetResources, &obj.TechnicalResources, &obj.RupType, &obj.Disciplines, &obj.CreatedAt, &obj.UpdatedAt)

	return obj, err
}

func (r *ObjectRepo) List(ctx context.Context) ([]*models.Object, error) {
	query := `
        SELECT id, title, year, type, course, study_form, exploration_note, goals, prerequisites, postrequisites, group_id, contacts, internet_resources, technical_resources, rup_type, disciplines, created_at, updated_at
        FROM objects
        ORDER BY created_at DESC
    `

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := []*models.Object{}
	for rows.Next() {
		obj := &models.Object{}
		err := rows.Scan(&obj.ID, &obj.Title, &obj.Year, &obj.Type, &obj.Course, &obj.StudyForm, &obj.ExplorationNote, &obj.Goals, &obj.Prerequisites, &obj.Postrequisites, &obj.GroupID, &obj.Contacts, &obj.InternetResources, &obj.TechnicalResources, &obj.RupType, &obj.Disciplines, &obj.CreatedAt, &obj.UpdatedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, obj)
	}

	return list, nil
}

func (r *ObjectRepo) Update(ctx context.Context, input models.UpdateObjectDTO) (*models.Object, error) {
	query := `
        UPDATE objects
        SET title = $1, year = $2, type = $3, course = $4, study_form = $5, exploration_note = $6, goals = $7, prerequisites = $8, postrequisites = $9, group_id = $10, contacts = $11, internet_resources = $12, technical_resources = $13, rup_type = $14, disciplines = $15, updated_at = NOW()
        WHERE id = $16
        RETURNING id, title, year, type, course, study_form, exploration_note, goals, prerequisites, postrequisites, group_id, contacts, internet_resources, technical_resources, rup_type, disciplines, created_at, updated_at
    `

	obj := &models.Object{}
	var id, title, year, typ, course, studyForm, explorationNote, goals, prerequisites, postrequisites, groupID, contacts, internetResources, technicalResources, rupType, disciplines interface{}
	if len(input.ID) > 0 {
		id = string(input.ID)
	}
	if len(input.Title) > 0 {
		title = string(input.Title)
	}
	if len(input.Year) > 0 {
		year = string(input.Year)
	}
	if len(input.Type) > 0 {
		typ = string(input.Type)
	}
	if len(input.Course) > 0 {
		course = string(input.Course)
	}
	if len(input.StudyForm) > 0 {
		studyForm = string(input.StudyForm)
	}
	if len(input.ExplorationNote) > 0 {
		explorationNote = string(input.ExplorationNote)
	}
	if len(input.Goals) > 0 {
		goals = string(input.Goals)
	}
	if len(input.Prerequisites) > 0 {
		prerequisites = string(input.Prerequisites)
	}
	if len(input.Postrequisites) > 0 {
		postrequisites = string(input.Postrequisites)
	}
	if len(input.GroupID) > 0 {
		groupID = string(input.GroupID)
	}
	if len(input.Contacts) > 0 {
		contacts = string(input.Contacts)
	}
	if len(input.InternetResources) > 0 {
		internetResources = string(input.InternetResources)
	}
	if len(input.TechnicalResources) > 0 {
		technicalResources = string(input.TechnicalResources)
	}
	if len(input.RupType) > 0 {
		rupType = string(input.RupType)
	}
	if len(input.Disciplines) > 0 {
		disciplines = string(input.Disciplines)
	}

	err := r.db.QueryRowContext(ctx, query, title, year, typ, course, studyForm, explorationNote, goals, prerequisites, postrequisites, groupID, contacts, internetResources, technicalResources, rupType, disciplines, id).
		Scan(&obj.ID, &obj.Title, &obj.Year, &obj.Type, &obj.Course, &obj.StudyForm, &obj.ExplorationNote, &obj.Goals, &obj.Prerequisites, &obj.Postrequisites, &obj.GroupID, &obj.Contacts, &obj.InternetResources, &obj.TechnicalResources, &obj.RupType, &obj.Disciplines, &obj.CreatedAt, &obj.UpdatedAt)

	return obj, err
}

func (r *ObjectRepo) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM objects WHERE id = $1`, id)
	return err
}
