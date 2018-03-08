package storage

import (
	"encoding/binary"
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
		// TODO Write the code to add new task into the bucket
		// Hint1: See https://github.com/boltdb/bolt#using-keyvalue-pairs to learn how to save the data
		// Hint2: You have to create Task object at first. Then, encode it as Json data typed `[]byte`
		//        If you forget how to marshal the object, see https://golang.org/pkg/encoding/json/#example_Marshal
		return
	})

	if err != nil {
		return
	}

	return
}

// PUT updates Task object
func (bs *BoltDBStorage) Put(t *Task) (err error) {
	err = bs.Update(func(tx *bolt.Tx) (err error) {
		// TODO Write the code to update the Task object
		// Hint1: See Hints of `Add` function
		return
	})

	return err
}

// Find retrieves a single TODO based on given taskNo.
func (bs *BoltDBStorage) FindByTaskNo(taskNo int) (task *Task, err error) {
	err = bs.View(func(tx *bolt.Tx) (err error) {
		// TODO Write the code to find the task connected to given `taskNo`
		// Hint1: When you want to iterate the data over keys, use `Cursor`(See https://github.com/boltdb/bolt#iterating-over-keys).
		// Hint2: You can use built-in function `copy(dst, src)` to copy the values of slice
		// Hint3: `FindByTaskNo` is using the given `taskNo` to search the data. This is **NOT** the key of the value.
		//        And `taskNo` is the number of **UN**finished tasks in the list. You can filter the tasks by Task.Finished
		//        `listCommand` in `main` function might be helpful.
		return
	})

	return
}

// Find retrieves all TODOs.
func (bs *BoltDBStorage) FindAll() (tasks []Task, err error) {
	bs.View(func(tx *bolt.Tx) (err error) {
		// TODO Write the code to fetch all data from the bucket and then append the result set into `tasks`
		// Hint1: You might want to iterate over all keys. See https://github.com/boltdb/bolt#foreach
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
