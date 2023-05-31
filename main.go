package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "errors"
    "strconv"
)

type User struct{
    ID        int64     `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Age       int       `json:"age"`
    Friends   []int64   `json:"friends"`
}

// In-memory database of users
var users = []User{
    {ID: 0005, Name: "Johnny Appleseed", Email: "john@gmail.com", Age: 30, Friends: []int64{1234, 5678, 9012}},
    {ID: 1234, Name: "Rebecca Peachpit", Email: "Rpp@yahoo.com", Age: 20, Friends: []int64{0005, 2222, 9012}},
    {ID: 5678, Name: "Phillip Strawberryleaf", Email: "Phillberry@hotmail.com", Age: 40, Friends: []int64{0005}},
    {ID: 9012, Name: "Jenna Grapestem", Email: "lovergurl222@protonmail.net", Age: 50, Friends: []int64{0005, 2222, 1234}},
    {ID: 2222, Name: "Shamuel Stone", Email: "shamtheman@stone.org", Age: 107, Friends: []int64{1234, 9012}},
}

// GET
func getUsers(c *gin.Context){
    c.IndentedJSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context){
    id := c.Param("id")
    user, err, _ := findUserByID(id)

    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found."})
        return
    }

    c.IndentedJSON(http.StatusOK, user)
}

func getFriendsOfUser(c *gin.Context){
    id := c.Param("id")
    user, err, _ := findUserByID(id)
    
    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found."})
        return
    }

    friends := findFriends(user.Friends)
    c.IndentedJSON(http.StatusOK, friends)
}

// POST
func addUser(c *gin.Context){
    var newUser User

    if err := c.BindJSON(&newUser); err != nil {
        return
    }

    users = append(users, newUser)
    c.IndentedJSON(http.StatusCreated, newUser)
}

// TODO Post /users/:id/friends
func addFriendToUser(c *gin.Context){

}

// PUT
func updateUser(c *gin.Context){
    id := c.Param("id")
    user, err, _ := findUserByID(id)

    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found."})
        return
    }
    var updatedUser User
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
func deleteUser(c *gin.Context){
    id := c.Param("id")
    _, err, index := findUserByID(id)

    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found."})
        return
    }

    users = append(users[:index], users[index+1:]...)
    c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Success"})
}

// TODO Delete /users/:id/friends/:friendId
func deleteFriendFromUser(c *gin.Context){
}


func main(){
    router := gin.Default()
    router.GET("/users", getUsers)
    router.GET("/users/:id", getUserByID)
    router.GET("/users/:id/friends", getFriendsOfUser)
    router.POST("/users", addUser)
    router.POST("/users/:id/friends", addFriendToUser)
    router.PUT("/users/:id", updateUser)
    router.DELETE("/users/:id", deleteUser)
    router.DELETE("/users/:id/friends/:friendId", deleteFriendFromUser)
    router.Run("localhost:8080")
}

// Helper Functions
func findUserByID(idS string) (*User, error, int){
    idN, err := strconv.ParseInt(idS, 10, 64)
    if err != nil{
        return nil, errors.New("Error with user ID"), -1
    }

    for i, u := range users {
        if u.ID == idN {
            return &users[i], nil, i
        }
    }

    return nil, errors.New("Error, user not found"), -1
}

func findFriends(list []int64) ([]User){
    
    friends:= make([]User, 0)
    for i := 0; i < len(list); i++ {
        for j := 0; j < len(users); j++{
            if list[i] == users[j].ID{
                friends = append(friends, users[j])
            }
        }
    }

    return friends
}
