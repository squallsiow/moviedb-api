package database

import (
	"encoding/json"
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
	pwd, err := os.Getwd()
	dbpath := filepath.Join(pwd, os.Getenv("DEFAULT_DATASTORE_FILEPATH"))

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
	fmt.Println("DB Setup Done")

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

func (db *Database) GetMovieByID(id int) (*model.Movie, error) {
	movie := model.Movie{}
	db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte(model.DB_MOVIE))ß
		v := b.Get([]byte(strconv.Itoa(id)))
		fmt.Printf("The answer is: %s\n", v)
		return nil
	})
}

// Sample
/*
func example() {
	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	conf := Config{Height: 186.0, Birthday: time.Now()}
	err = setConfig(db, conf)
	if err != nil {
		log.Fatal(err)
	}
	err = addWeight(db, "85.0", time.Now())
	if err != nil {
		log.Fatal(err)
	}
	err = addEntry(db, 100, "apple", time.Now())
	if err != nil {
		log.Fatal(err)
	}

	err = addEntry(db, 100, "orange", time.Now().AddDate(0, 0, -2))
	if err != nil {
		log.Fatal(err)
	}

	err = db.View(func(tx *bolt.Tx) error {
		conf := tx.Bucket([]byte("DB")).Get([]byte("CONFIG"))
		fmt.Printf("Config: %s\n", conf)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("WEIGHT"))
		b.ForEach(func(k, v []byte) error {
			fmt.Println(string(k), string(v))
			return nil
		})
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	err = db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte("DB")).Bucket([]byte("ENTRIES")).Cursor()
		min := []byte(time.Now().AddDate(0, 0, -7).Format(time.RFC3339))
		max := []byte(time.Now().AddDate(0, 0, 0).Format(time.RFC3339))
		for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
			fmt.Println(string(k), string(v))
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}


func addWeight(db *bolt.DB, weight string, date time.Time) error {
	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte("WEIGHT")).Put([]byte(date.Format(time.RFC3339)), []byte(weight))
		if err != nil {
			return fmt.Errorf("could not insert weight: %v", err)
		}
		return nil
	})
	fmt.Println("Added Weight")
	return err
}

func addEntry(db *bolt.DB, calories int, food string, date time.Time) error {
	entry := Entry{Calories: calories, Food: food}
	entryBytes, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("could not marshal entry json: %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Bucket([]byte("ENTRIES")).Put([]byte(date.Format(time.RFC3339)), entryBytes)
		if err != nil {
			return fmt.Errorf("could not insert entry: %v", err)
		}

		return nil
	})
	fmt.Println("Added Entry")
	return err
}
*/
