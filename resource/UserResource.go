package resource

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "github.com/ReedOlm/GoRestAPI/data"
    "github.com/ReedOlm/GoRestAPI/model"
    "github.com/ReedOlm/GoRestAPI/service"
)

// GET
func GetUsers(c *gin.Context){
    c.IndentedJSON(http.StatusOK, data.Users)
}

func GetUserByID(c *gin.Context){
    id := c.Param("id")
    user, err, _ := service.FindUserByID(id)

    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found."})
        return
    }

    c.IndentedJSON(http.StatusOK, user)
}

func GetFriendsOfUser(c *gin.Context){
    id := c.Param("id")
    user, err, _ := service.FindUserByID(id)

    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found."})
        return
    }

    friends := service.FindFriends(user.Friends)
    c.IndentedJSON(http.StatusOK, friends)
}

// POST
func AddUser(c *gin.Context){
    var newUser model.User

    if err := c.BindJSON(&newUser); err != nil {
        return
    }

    data.Users = append(data.Users, newUser)
    c.IndentedJSON(http.StatusCreated, newUser)
}

func AddFriendToUser(c *gin.Context){
    id := c.Param("id")
    user, err, _ := service.FindUserByID(id)

    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found."})
        return
    }

    var newFriend model.User
    if err = c.BindJSON(&newFriend); err != nil {
        return
    }

    user.Friends = append(user.Friends, newFriend.ID)
    c.IndentedJSON(http.StatusCreated, newFriend.ID)
}

// PUT
func UpdateUser(c *gin.Context){
    id := c.Param("id")
    user, err, _ := service.FindUserByID(id)

    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found."})
        return
    }
    var updatedUser model.User
    if er := c.BindJSON(&updatedUser); er != nil{
        return
    }

    user.Name = updatedUser.Name
    user.Email = updatedUser.Email
    user.Age = updatedUser.Age
    user.Friends = updatedUser.Friends

    c.IndentedJSON(http.StatusOK, gin.H{"message": "Updated user information"})
}

// DELETE
func DeleteUser(c *gin.Context){
    id := c.Param("id")
    _, err, index := service.FindUserByID(id)

    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found."})
        return
    }

    data.Users = append(data.Users[:index], data.Users[index+1:]...)
    c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Success"})
}

func DeleteFriendFromUser(c *gin.Context){
    id := c.Param("id")
    user, err, _ := service.FindUserByID(id)

    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found."})
        return
    }

    friendId := c.Param("friendId")
    friendIdN, er := strconv.ParseInt(friendId, 10, 64)
    if er != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Error with friend's ID"})
        return
    }
    
    index := service.FindFriendIndex(user.Friends, friendIdN)
    if index < 0 {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Friend not found"})
        return
    }

    user.Friends = append(user.Friends[:index], user.Friends[index+1:]...)

    c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Success"})
}
