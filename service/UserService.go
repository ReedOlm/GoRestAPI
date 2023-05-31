package service

import (
    "github.com/ReedOlm/GoRestAPI/model"    
    "github.com/ReedOlm/GoRestAPI/data"
    "strconv"
    "errors" 
)

func FindUserByID(idS string) (*model.User, error, int){
    idN, err := strconv.ParseInt(idS, 10, 64)
    if err != nil{
        return nil, errors.New("Error with user ID"), -1
    }

    for i, u := range data.Users {
        if u.ID == idN {
            return &data.Users[i], nil, i
        }
    }

    return nil, errors.New("Error, user not found"), -1
}

func FindFriends(list []int64) ([]model.User){

    friends:= make([]model.User, 0)
    for i := 0; i < len(list); i++ {
        for j := 0; j < len(data.Users); j++{
            if list[i] == data.Users[j].ID{
                friends = append(friends, data.Users[j])
            }
        }
    }

    return friends
}
