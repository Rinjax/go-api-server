package store

type Faker struct {
	Users []NewUser

}

func NewFaker() *Faker {
	return &Faker{}
}

func (f *Faker) AddUserAccount(u *NewUser) error {
	f.Users = append(f.Users, *u)
	return nil
}