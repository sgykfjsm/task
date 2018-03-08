package storage

import (
	"encoding/binary"
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
)

type Task struct {
	SystemID    []byte // id is just for internal usage. Do not update!
	Description string
	Finished    bool
}

type Storage interface {
	Add(string) (int, error)
	FindByTaskNo(int) error
	FindAll()
}

type BoltDBStorage struct {
	*bolt.DB
	BucketName []byte
}

// NewBoltDBStorage generates the new BoltDBStorage object
func NewBoltDBStorage(filePath, bucketName string) (*BoltDBStorage, error) {
	b, err := bolt.Open(filePath, 0644, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	return &BoltDBStorage{
		DB:         b,
		BucketName: []byte(bucketName),
	}, nil
}

// Add adds new task. If succeeded, Add returns the pointer of Task with Nil. If not, Add returns error and the pointer of Task is nil
func (bs *BoltDBStorage) Add(description string) (task *Task, err error) {
	err = bs.Update(func(tx *bolt.Tx) (err error) {
		bucket, err := tx.CreateBucketIfNotExists(bs.BucketName)
		if err != nil {
			return
		}

		id, err := bucket.NextSequence()
		if err != nil {
			return
		}

		task = &Task{
			SystemID:    bs.itob(int(id)),
			Description: description,
		}

		buf, err := json.Marshal(task)
		if err != nil {
			return
		}

		return bucket.Put(task.SystemID, buf)
	})

	if err != nil {
		return
	}

	return
}

// PUT updates Task object
func (bs *BoltDBStorage) Put(t *Task) (err error) {
	err = bs.Update(func(tx *bolt.Tx) (err error) {
		b, err := json.Marshal(t)
		if err != nil {
			return
		}

		if err := tx.Bucket(bs.BucketName).Put(t.SystemID, b); err != nil {
			return err
		}

		return
	})

	return err
}

// Find retrieves a single TODO based on given taskNo.
func (bs *BoltDBStorage) FindByTaskNo(taskNo int) (task *Task, err error) {
	err = bs.View(func(tx *bolt.Tx) (err error) {
		b := tx.Bucket(bs.BucketName)
		if b == nil {
			return
		}

		c := b.Cursor()
		i := 1
		var t Task
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if err := json.Unmarshal(v, &t); err != nil {
				return err
			}
			if t.Finished {
				continue
			}
			if i == taskNo {
				task = &t
				return
			}
			i++
		}

		return
	})

	return
}

// Find retrieves all TODOs.
func (bs *BoltDBStorage) FindAll() (tasks []Task, err error) {
	bs.View(func(tx *bolt.Tx) (err error) {
		data := tx.Bucket(bs.BucketName)
		if data == nil {
			return
		}

		data.ForEach(func(k, v []byte) error {
			var task Task
			if err := json.Unmarshal(v, &task); err != nil {
				return err
			}
			tasks = append(tasks, task)
			return nil
		})

		return
	})

	return
}

func (bs *BoltDBStorage) itob(i int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))

	return b
}

func (bs *BoltDBStorage) btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
