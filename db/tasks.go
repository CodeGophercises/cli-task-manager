package db

import (
	"time"

	"github.com/CodeGophercises/cli-task-manager/utils"
	"github.com/boltdb/bolt"
)

var db *bolt.DB
var bucket = []byte("tasks")

type Task struct {
	Id  int
	Val string
}

// Creates a task and persists in DB. Returns taskId back on succcessful creation.
func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(bucket)
		// No reason for bucket to be nil, ignoring error check
		id64, _ := bkt.NextSequence()
		id = int(id64)
		key := utils.Itob(id)
		err := bkt.Put(key, []byte(task))
		return err
	})

	if err != nil {
		return -1, err
	}

	return id, nil
}

func GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(bucket)
		c := bkt.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Id:  utils.Btoi(k),
				Val: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func Delete(key []byte) error {
	return db.Update(func(tx *bolt.Tx) error {

		bkt := tx.Bucket(bucket)
		return bkt.Delete(key)
	})
}

func InitPath(path string) error {
	// Open DB and create the bucket
	var err error
	db, err = bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}
		return nil
	})
}
