package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginUserInput) (User, error)
	IsEmailAvailable(email CheckEmailInput) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Email = input.Email
	user.Role = "user"

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) LoginUser(input LoginUserInput) (User, error) {

	newUser, err := s.repository.FindUserByEmail(input.Email)
	if err != nil {
		return newUser, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(newUser.PasswordHash), []byte(input.Password))
	if err != nil {
		return newUser, err
	}

	return newUser, nil

}

func (s *service) IsEmailAvailable(email CheckEmailInput) (bool, error) {
	var user User
	user, err := s.repository.FindUserByEmail(email.Email)

	if err != nil {
		return false, err
	}

	if user.Email == email.Email {
		return false, nil
	}

	return true, nil
}
