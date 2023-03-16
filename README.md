# workflow

1. db without server / apply mock up data, defining Res Model/Data Model struct
2. make service
3. install server / make controller
4. routing test
5. install err handling
6. debug
7. impl token, sign/verify
8. login, signin
9. remove debug / test route

# how jwt works
```go
1. generate token with claims
2. extract token from Bearer with c.GetHeader("Authorization")
3. parse and verfiy token
4. after verifying, use token claims.

claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
            Message: "unable to parse claims",
        })
        return
    }

    claimedUID, ok := claims["user"].(string)
    if !ok {
        c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
            Message: "no user property in claims",
        })
        return
    }
```

# json
[json tags](https://stackoverflow.com/questions/17306358/removing-fields-from-struct-or-hiding-them-in-json-response)
```go
type OkResponse struct {
	User    primitive.M   `json:"user,omitempty"` //omit field
	Users   []primitive.M `json:"users,omitempty"` //omit field
	Message string        `json:"message"`
}

//in bson, handle automatically without tag,
//just write to change field - tagName + omit field
type User struct {
	Email     string
	Password  string    `bson:"omitempty"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	RtToken   string    `bson:"rt_token"`
}
```