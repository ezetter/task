package db

import (
	"encoding/json"
	"fmt"
	"strconv"

	"go.etcd.io/bbolt"
	bolt "go.etcd.io/bbolt"
)

var db *bbolt.DB

func ListTasks() []Task {
	var tasks []Task
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var t Task
			json.Unmarshal(v, &t)
			tasks = append(tasks, t)
		}
		return nil
	})
	return tasks
}

type Task struct {
	ID     uint64
	Desc   string
	Status string
}

func AddTask(task string) error {
	return db.Update(func(tx *bolt.Tx) error {
		var id uint64
		b := tx.Bucket([]byte("Tasks"))
		id, err := b.NextSequence()
		if err != nil {
			return err
		}
		t := new(Task)
		t.ID = id
		t.Desc = task
		t.Status = "INPROGRESS"

		if buf, err := json.Marshal(t); err != nil {
			return err
		} else if err := b.Put([]byte(strconv.FormatUint(id, 10)), buf); err != nil {
			return err
		}

		return nil
	})
}

func RemoveTask(taskNum int) string {
	task := ListTasks()[taskNum-1]
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		err := b.Delete([]byte(strconv.FormatUint(task.ID, 10)))
		return err
	})
	return task.Desc
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
