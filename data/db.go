package data

import "github.com/ReedOlm/GoRestAPI/model"

// In-memory database of users
var Users = []model.User{
    {ID: 0005, Name: "Johnny Appleseed", Email: "john@gmail.com", Age: 30, Friends: []int64{1234, 5678, 9012}},
    {ID: 1234, Name: "Rebecca Peachpit", Email: "Rpp@yahoo.com", Age: 20, Friends: []int64{0005, 2222, 9012}},
    {ID: 5678, Name: "Phillip Strawberryleaf", Email: "Phillberry@hotmail.com", Age: 40, Friends: []int64{0005}},
    {ID: 9012, Name: "Jenna Grapestem", Email: "lovergurl222@protonmail.net", Age: 50, Friends: []int64{0005, 2222, 1234}},
    {ID: 2222, Name: "Shamuel Stone", Email: "shamtheman@stone.org", Age: 107, Friends: []int64{1234, 9012}},
}
