package model

type User struct{
    ID        int64     `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Age       int       `json:"age"`
    Friends   []int64   `json:"friends"`
}
