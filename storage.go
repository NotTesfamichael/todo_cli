package main

import (
	"encoding/json"
	"os"
)

type Storage [T any] struct {
	FIlename string
}
func NewStorage[T any](fileName string) *Storage[T]{
	return &Storage[T]{FIlename: fileName}
}
func (s *Storage[T]) Save(data T) error{
	fileData, err := json.MarshalIndent(data, "", "")
	if err != nil {
	return err
}
return os.WriteFile(s.FIlename, fileData, 0644)
}
func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FIlename)
	if err!=nil{
		return nil
	}
	return json.Unmarshal(fileData, data)
}