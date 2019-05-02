package repository

import (
	"blog/user"
	"blog/models"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

type mongoUserRepository struct {
	DB *mgo.Session
}

func NewMongoUserRepository(db *mgo.Session) user.Repository {
	return &mongoUserRepository{
		DB: db,
	}
}

func (m *mongoUserRepository) Save(a *models.User) (err error) {
	db := m.DB.Clone()
	defer db.Close()
	if err = db.DB("blog").C("users").Insert(a); err != nil {
		return err
	}
	return nil
}
	
// func (m *mongoUserRepository) Find(a *models.User) (err error) {
// 	db := m.DB.Clone()
// 	defer db.Close()
// 	if err = db.DB("twitter").C("users").
// 		Find(bson.M{"email": a.Email, "password": a.Password}).One(a); err != nil {
// 		if err == mgo.ErrNotFound {
// 			return err
// 		}
// 		return err
// 	}
// 	return nil
// }