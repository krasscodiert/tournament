package tournamentserver

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/boltdb/bolt"
)

//DB Struct
type DB struct {
	db         *bolt.DB
	userBucket *bolt.Bucket
}

//New - Default constructor
func (d *DB) New() *DB {

	dbd, err := bolt.Open("my.db", 0600, nil)
	d.db = dbd
	d.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("User"))
		return err
	})

	if err != nil {
		log.Fatal(err)
	}

	return d
}

// Close - closes Db connection
func (d *DB) Close() {
	d.db.Close()
}

func (d *DB) getUser(name string) (login, bool) {
	var User login

	value := d.read("User", name)
	if value == nil {
		return login{}, false
	}
	var buf bytes.Buffer
	buf.Write(value)
	dec := gob.NewDecoder(&buf)
	err := dec.Decode(&User)
	if err != nil {
		log.Fatal("decode error:", err)
	}

	return User, true

}

func (d *DB) saveUser(user *login) error {
	return d.write("User", user.Username, user)
}

func (d *DB) write(bucket string, key string, value interface{}) error {
	return d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		bytes, _ := d.getBytes(value)
		err := b.Put([]byte(key), []byte(bytes))
		return err
	})
}
func (d *DB) read(bucket string, key string) []byte {
	var val []byte
	d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		val = b.Get([]byte(key))
		return nil
	})
	return val
}

func (d *DB) getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
