package main

import (
	"log"
	"fmt"
	"encoding/json"	
	
	bolt "github.com/coreos/bbolt"
)

type Command struct {
	field1, field2 string
}

func main() {
	// Create database
	db, err := bolt.Open("test.db", 0600, nil)
	
	if err != nil {
		log.Fatal(err)
	}
	
	defer db.Close()
	
	// Creating bucket
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("TestBucket"))
		
		if err!=nil {
			return fmt.Errorf("Error creating bucket", err)
		}
		
		// Adding entry
		err = b.Put([]byte("key1"), []byte(`{"field1":"1", "field2":"2"}`))
		if err != nil {
			log.Fatal(err)
		}
		
		// Getting entry
		k1 := "key1"
		v := b.Get([]byte(k1))
		
		fmt.Printf("Key:%s Value:%s", k1, v)
		
		// Getting entry
		k2 := "key2"
		v = b.Get([]byte(k2))
		
		fmt.Printf("Key:%s Value:%s", k2, v)
					
		// Using unmarshal
		var cmd Command
		err = json.Unmarshal(v, &cmd)
		
		if err != nil {
			fmt.Println("error unmarshal:", err)
		}
		
		fmt.Printf("%+v", cmd)
		
		return nil	
	})
}
