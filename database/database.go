package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/moviedb-api/model"
)

// Database : type
type Database struct {
	DB *bolt.DB
}

// New : Initialize Database connection
func New() (*Database, error) {
	// pwd, err := os.Getwd()
	// dbpath := filepath.Join(pwd, os.Getenv("DEFAULT_DATASTORE_FILEPATH"))
	dbpath, err := filepath.Abs(os.Getenv("DEFAULT_DATASTORE_FILEPATH"))

	if err != nil {
		return nil, fmt.Errorf("path not valid, %v", err)
	}
	// Create db file
	if _, err := os.Stat(dbpath); os.IsNotExist(err) {
		os.MkdirAll(dbpath, os.ModeDir)
	}

	db, err := bolt.Open(dbpath, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		// Root DB
		root, err := tx.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		// movie DB Table
		_, err = root.CreateBucketIfNotExists([]byte(model.DB_MOVIE))
		if err != nil {
			return fmt.Errorf("could not create movie bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}
	database := Database{DB: db}

	return &database, nil
}

// InsertMovie : Insert movie to DB
func (db *Database) InsertMovie(m model.Movie) error {
	mBytes, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("could not marshal movie json: %v", err)
	}
	err = db.DB.Update(func(tx *bolt.Tx) error {
		key := strconv.Itoa(m.ID)
		err = tx.Bucket([]byte("DB")).Bucket([]byte(model.DB_MOVIE)).Put([]byte(key), []byte(mBytes))
		if err != nil {
			return fmt.Errorf("could not set movie: %v", err)
		}
		return nil
	})
	return err
}

// GetAllMovies : Get all local movies
func (db *Database) GetAllMovies() ([]*model.Movie, error) {
	movies := make([]*model.Movie, 0)
	err := db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte(model.DB_MOVIE))
		b.ForEach(func(k, v []byte) error {
			m := model.Movie{}
			if err := json.Unmarshal(v, &m); err != nil {
				return err
			}
			movies = append(movies, &m)
			return nil
		})
		return nil
	})

	return movies, err
}

// GetMovieByID :
func (db *Database) GetMovieByID(id int) (*model.Movie, error) {
	movie := model.Movie{}
	err := db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte(model.DB_MOVIE))
		v := b.Get([]byte(strconv.Itoa(id)))
		if v == nil {
			return errors.New("No movie ID " + strconv.Itoa(id) + " exist!")
		}
		err := json.Unmarshal(v, &movie)

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &movie, nil
}
