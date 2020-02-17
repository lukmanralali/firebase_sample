package repository

import (
    firebase "firebase_sample/utility/firebase"
    firebaseEntity "firebase_sample/entity/firebase"

    "fmt"
)

type FirebaseRepository struct {
}

type FirebaseRepository interface {
    getData(ref String) map[string]interface
    upsert(payload map[string]UserAttendant, routeKey string, id string)
}

func FirebaseRepositoryHandler() *FirebaseRepository{
	return &FirebaseRepository{
	}
}

func (repo *FirebaseRepository) getData(ref string) map[string]interface {
	ref := client.NewRef(ref)
	var data map[string]interface{}
	if err := ref.Get(ctx, &data); err != nil {
	        log.Fatalln("Error reading from database:", err)
	}
	return data
}

func (repo *FirebaseRepository) upsert(payload map[string]UserAttendant, routeKey string, id string){
	baseRef := client.NewRef(ref)
	childRef := ref.Child(baseRef)
	usersRef.Update(ctx, payload); err != nil {
	        log.Fatalln("Error updating children:", err)
	}
}