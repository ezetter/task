package db

import (
	"fmt"

	"go.etcd.io/bbolt"
	bolt "go.etcd.io/bbolt"
)

var db *bbolt.DB

func ListTasks() []string {
	var tasks []string
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		c := b.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			tasks = append(tasks, string(k))
		}
		return nil
	})
	return tasks
}

func AddTask(task string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		err := b.Put([]byte(task), []byte("INPROGRESS"))
		return err
	})
}

func RemoveTask(taskNum int) string {
	task := ListTasks()[taskNum-1]

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		err := b.Delete([]byte(task))
		return err
	})
	return task
}

func Init(path string) error {
	var err error
	db, err = bolt.Open(fmt.Sprintf("%s/tasks.db", path), 0600, nil)
	if err != nil {
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Tasks"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	return err
}

func Close() {
	db.Close()
}
