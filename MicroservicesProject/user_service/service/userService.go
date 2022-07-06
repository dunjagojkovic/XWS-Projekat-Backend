package service

import (
	"crypto/rand"
	"encoding/base64"
	"math"
	"time"
	"userS/model"
	"userS/repository"

	"strings"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func randomBase64String(l int) string {
	buff := make([]byte, int(math.Ceil(float64(l)/float64(1.33333333333))))
	rand.Read(buff)
	str := base64.RawURLEncoding.EncodeToString(buff)
	return str[:l] // strip 1 extra character we get from odd length results
}

type UserService struct {
	store repository.UserStoreI
}

func NewUserService(store repository.UserStoreI) *UserService {
	return &UserService{
		store: store,
	}
}

func (service *UserService) RegisterUser(user *model.User) (*model.User, error) {
	return service.store.RegisterUser(user)
}

func (service *UserService) CheckBlocking(first, second string) bool {
	return service.store.ChechBlocking(first, second)
}

func (service *UserService) Login(username, password string) (string, string, error) {
	isLoged, _ := service.store.Login(username, password)
	if isLoged {

		expirationTime := time.Now().Add(60 * time.Minute)
		claims := &Claims{
			Username: username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			return "", "", err

		}
		key := randomBase64String(10)

		return tokenString, key, nil

	}
	return "", "", nil

}

func (service *UserService) CurrentUser(username string) (model.User, error) {
	return service.store.CurrentUser(username)

}

func (service *UserService) GetUser(id primitive.ObjectID) (model.User, error) {
	return service.store.GetUser(id)

}

func (service *UserService) GetUsers() ([]*model.User, error) {
	return service.store.GetUsers()

}

func (service *UserService) GetPublicUsers() ([]*model.User, error) {
	return service.store.GetPublicUsers()

}

func (service *UserService) GetUsersById(users []string) ([]*model.User, error) {
	return service.store.GetUsersById(users)

}

func (service *UserService) EditUser(user *model.User, work *model.WorkExperience) (*model.User, error) {
	return service.store.EditUser(user, work)

}

func (service *UserService) EditPassword(newPassword, oldPassword, username string) (*model.User, error) {
	return service.store.EditPassword(newPassword, oldPassword, username)

}

func (service *UserService) EditPrivacy(isPublic bool, username string) (*model.User, error) {
	return service.store.EditPrivacy(isPublic, username)

}

func (service *UserService) FilterUsers(searchTerm string) ([]*model.User, error) {
	users, err := service.store.GetPublicUsers()
	var filterUsers []*model.User
	for _, user := range users {
		if strings.Contains(strings.ToLower(user.Username), strings.ToLower(searchTerm)) || strings.Contains(strings.ToLower(user.Name), strings.ToLower(searchTerm)) || strings.Contains(strings.ToLower(user.Surname), strings.ToLower(searchTerm)) {
			filterUsers = append(filterUsers, user)
		}
	}
	if err != nil {
		return nil, err
	}
	return filterUsers, nil

}

func (service *UserService) BlockUser(block *model.Block) (primitive.ObjectID, error) {
	return service.store.BlockUser(block)

}

func (service *UserService) Unblock(block *model.Block) (primitive.ObjectID, error) {
	return service.store.Unblock(block)

}
