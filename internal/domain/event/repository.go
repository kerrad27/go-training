package event

import (
	"fmt"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

var settings = postgresql.ConnectionURL{
	Database: `golangTest`,
	Host:     `localhost`,
	User:     `postgres`,
	Password: `golang27`,
}

type Repository interface {
	FindAll() ([]Poster, error)
	FindOne(id int64) (*Poster, error)
	Add(title string, description string, year int64) (*Poster, error)
	Update(id int64, title string, description string, year int64) (*Poster, error)
}

const EventsCount int64 = 10

type repository struct {
	// Some internal data
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() ([]Poster, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("FindAll.Open(settings): ", err)
	}
	defer func(sess db.Session) {
		err := sess.Close()
		if err != nil {
			log.Fatal("", err)
		}
	}(sess)

	infoColl := sess.Collection("Poster")

	var infos []Poster

	err = infoColl.Find().All(&infos)
	if err != nil {
		log.Fatal("infoCol: ", err)
	}

	fmt.Printf("Records in the %q collection:\n", infoColl.Name())
	for i := range infos {
		fmt.Printf("Record #%d: %#v\n", i, infos[i])
	}

	return infos, nil
}

func (r *repository) FindOne(id int64) (*Poster, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("FindOne.Open(settings): ", err)
	}
	defer func(sess db.Session) {
		err := sess.Close()
		if err != nil {
			log.Fatal("", err)
		}
	}(sess)

	var info Poster

	err = sess.SQL().
		SelectFrom("Poster").
		Where("id", id).
		One(&info)
	if err != nil {
		log.Fatal("Query: ", err)
	}

	return &Poster{
		ID:          info.ID,
		Title:       info.Title,
		Description: info.Description,
		Year:        info.Year,
	}, nil
}

func (r *repository) Add(title, description string, year int64) (*Poster, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("postgresql.Open: ", err)
	}
	defer sess.Close()
	sess.SQL().InsertInto("Poster").Columns("title", "description", "year").
		Values(title, description, year).Exec()
	if err != nil {
		log.Fatal("sess.Insert:", err)
	}
	return &Poster{
		Title:       title,
		Description: description,
		Year:        year,
	}, nil
}

func (r *repository) Update(id int64, title, description string, year int64) (*Poster, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	db.LC().SetLevel(db.LogLevelDebug)
	defer sess.Close()

	if title != "" {
		_, err := sess.SQL().
			Update("Poster").
			Set("title = ?", title, "description = ?", description, "year = ?", year).
			Where("id = ?", id).
			Exec()
		if err != nil {
			fmt.Printf("sess.SQL: %v. This is expected on the read-only sandbox.\n", err)
		}
	}

	var info Poster

	err = sess.SQL().
		SelectFrom("Poster").
		Where("id", id).
		One(&info)
	if err != nil {
		log.Fatal("sess.SQL: ", err)
	}

	return &Poster{
		ID:          info.ID,
		Title:       info.Title,
		Description: info.Description,
		Year:        info.Year,
	}, nil
}
