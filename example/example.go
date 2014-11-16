package main

import (
	"fmt"
	"github.com/amidvidy/wtgo"
)

func main() {
	conn, err := wtgo.Open(".", "create,cache_size=500M")
	if err != nil {
		panic(fmt.Sprintf("Failed to create database: %v", err.Error()))
	}
	session, err := conn.OpenSession("isolation=snapshot")
	if err != nil {
		panic(fmt.Sprintf("Failed to create session: %v", err.Error()))

	}
	err = session.Create("table:access", "key_format=u,value_format=u")
	if err != nil {
		panic(fmt.Sprintf("Failed to create table: %v", err.Error()))
	}
	cursor, err := session.OpenCursor("table:access", "")
	if err != nil {
		panic(fmt.Sprintf("Failed to open cursor: %v", err.Error()))
	}
	cursor.SetKey([]byte("foo"))
	cursor.SetValue([]byte("bar"))
	err = cursor.Insert()
	if err != nil {
		panic(fmt.Sprintf("Failed to insert: %v", err.Error()))
	}
	cursor.Reset()
	for {
		if cursor.Next() != nil {
			break
		}
		key, err := cursor.GetKey()
		if err != nil {
			panic(fmt.Sprintf("Failed to get key: %v", err.Error()))
		}
		val, err := cursor.GetValue()
		if err != nil {
			panic(fmt.Sprintf("Failed to get value: %v", err.Error()))
		}
		fmt.Printf("Got key: %v, Got value: %v\n",
			string(key),
			string(val))
	}
}
