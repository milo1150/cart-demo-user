package types

type UserJsonFile struct {
	Users []UserJson
}

type UserJson struct {
	Username string
	Email    string
}
