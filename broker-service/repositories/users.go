package repositories

import (
	"broker/data"
	"broker/utils"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func CreateUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) insert(user *data.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) Create(username, password string) (*data.User, error) {
	var user = data.User{}
	id, err := utils.GenerateUUID()
	if err != nil {
		return &user, err
	}

	hashedPassword, err := utils.GeneratePassword(password)
	if err != nil {
		return &user, err
	}

	user.ID = id
	user.Username = username
	user.HashedPassword = hashedPassword
	err = r.insert(&user)

	return &user, err
}

func (r *UserRepo) GetByUsernameAndPassword(username, password string) (*data.User, bool) {
	user := new(data.User)
	emptyUser := data.User{}

	err := r.db.Where("name = ?", username).Find(user).Error
	if err != nil {
		return &emptyUser, false
	}

	ok := utils.ValidatePassword(password, user.HashedPassword)
	if ok {
		return user, true
	}

	return &emptyUser, false

}
