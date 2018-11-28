package common

import (
	"fmt"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
)

type glasspoolDB struct {
	boltptr *bolt.DB
}

var Gdb glasspoolDB

// multiple processes cannot open the same database at the same time
func InitDB() {
	// It will be created if it doesn't exist.
	db, err := bolt.Open("glasspool.bolt", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("DEPOSIT"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("WITHDRAW"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("ORDER"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("POSTED"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("INTERNAL"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("WALLET"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	// db.Update(func(tx *bolt.Tx) error {
	// 	_, err := tx.CreateBucketIfNotExists([]byte("TIMESTAMP"))
	// 	if err != nil {
	// 		return fmt.Errorf("create bucket: %s", err)
	// 	}
	// 	return nil
	// })

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("INTERNAL"))
		v := b.Get([]byte("OrderID"))
		if v == nil {
			b.Put([]byte("OrderID"), Itob(1))
		}
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("INTERNAL"))
		v := b.Get([]byte("OrderStatusUpdate"))
		if v == nil {
			b.Put([]byte("OrderStatusUpdate"), Itob(1))
		}
		return nil
	})

	Gdb.boltptr = db
}

// key: ETH0x4814..., value: "http://127.0.0.1:9008/callback"
func (gdb glasspoolDB) DEPOSIT(key, value []byte) error {
	return gdb.boltptr.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DEPOSIT"))
		err := b.Put(key, value)
		return err
	})
}

// key: ETH0x4814..., value: "http://127.0.0.1:9008/callback"
func (gdb glasspoolDB) WITHDRAW(key, value []byte) error {
	return gdb.boltptr.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("WITHDRAW"))
		err := b.Put(key, value)
		return err
	})
}

// key: request, value: response
func (gdb glasspoolDB) POSTED(key, value []byte) error {
	return gdb.boltptr.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("POSTED"))
		err := b.Put(key, value)
		return err
	})
}

// key: id, value: order
func (gdb glasspoolDB) ORDER(key, value []byte) error {
	return gdb.boltptr.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("ORDER"))
		err := b.Put(key, value)
		return err
	})
}

// key: key, value: value
// func (gdb glasspoolDB) INTERNAL(key, value []byte) error {
// 	return gdb.boltptr.Update(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte("INTERNAL"))
// 		err := b.Put(key, value)
// 		return err
// 	})
// }

// key: auto timestamp
// func (gdb glasspoolDB) TIMESTAMP(value []byte) error {
// 	return gdb.boltptr.Update(func(tx *bolt.Tx) error {
// 		key := Itob(uint64(Timestamp()))
// 		b := tx.Bucket([]byte("TIMESTAMP"))
// 		err := b.Put(key, value)
// 		return err
// 	})
// }

func (gdb glasspoolDB) WALLET(key, value []byte) error {
	return gdb.boltptr.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("WALLET"))
		err := b.Put(key, value)
		return err
	})
}

// return id then increase it in database
func (gdb glasspoolDB) OrderID() string {
	var id uint64
	gdb.boltptr.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("INTERNAL"))
		v := b.Get([]byte("OrderID"))
		id = Btoi(v)
		return nil
	})
	gdb.boltptr.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("INTERNAL"))
		err := b.Put([]byte("OrderID"), Itob(id+1))
		return err
	})
	return strconv.FormatUint(id, 10)
}

// get the last id updated in previous batch
func (gdb glasspoolDB) ReadOrderStateUpdate() (id uint64) {
	gdb.boltptr.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("INTERNAL"))
		v := b.Get([]byte("OrderStatusUpdate"))
		id = Btoi(v)
		return nil
	})
	return id
}

// write the last id updated in this batch
func (gdb glasspoolDB) WriteOrderStateUpdate(id uint64) error {
	return gdb.boltptr.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("INTERNAL"))
		err := b.Put([]byte("OrderStatusUpdate"), Itob(id))
		return err
	})
}

// return callback
func (gdb glasspoolDB) Callback(bizType, addr string) (url []byte) {
	gdb.boltptr.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bizType))
		url = b.Get([]byte(addr))
		return nil
	})
	return url
}

// query order with order id
func (gdb glasspoolDB) QueryOrder(orderID string) (orderBytes []byte) {
	gdb.boltptr.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("ORDER"))
		orderBytes = b.Get([]byte(orderID))
		return nil
	})
	return orderBytes
}
