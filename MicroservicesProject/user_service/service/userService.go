package service

import (
	"common/tracer"
	"context"
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

func (service *UserService) RegisterUser(ctx context.Context, user *model.User) (*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE RegisterUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.RegisterUser(ctx, user)
}

func (service *UserService) Login(ctx context.Context, username, password string) (string, string, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE Login")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	isLoged, _ := service.store.Login(ctx, username, password)
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

func (service *UserService) CurrentUser(ctx context.Context, username string) (model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE CurrentUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.CurrentUser(ctx, username)

}

func (service *UserService) GetUser(ctx context.Context, id primitive.ObjectID) (model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.GetUser(ctx, id)

}

func (service *UserService) GetUsers(ctx context.Context) ([]*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetUsers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.GetUsers(ctx)

}

func (service *UserService) GetPublicUsers(ctx context.Context) ([]*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetPublicUsers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.GetPublicUsers(ctx)

}

func (service *UserService) GetUsersById(ctx context.Context, users []string) ([]*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE GetUsersById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.GetUsersById(ctx, users)

}

func (service *UserService) EditUser(ctx context.Context, user *model.User, work *model.WorkExperience) (*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE EditUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.EditUser(ctx, user, work)

}

func (service *UserService) EditPassword(ctx context.Context, newPassword, oldPassword, username string) (*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE EditPassword")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.EditPassword(ctx, newPassword, oldPassword, username)

}

func (service *UserService) EditPrivacy(ctx context.Context, isPublic bool, username string) (*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE EditPrivacy")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.EditPrivacy(ctx, isPublic, username)

}

func (service *UserService) FilterUsers(ctx context.Context, searchTerm string) ([]*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE FilterUsers")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	users, err := service.store.GetPublicUsers(ctx)
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

func (service *UserService) BlockUser(ctx context.Context, block *model.Block) (primitive.ObjectID, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE BlockUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.BlockUser(ctx, block)

}

func (service *UserService) Unblock(ctx context.Context, block *model.Block) (primitive.ObjectID, error) {
	span := tracer.StartSpanFromContext(ctx, "SERVICE Unblock")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.Unblock(ctx, block)

}

func (service *UserService) CheckBlocking(ctx context.Context, first, second string) bool {
	span := tracer.StartSpanFromContext(ctx, "SERVICE CheckBlocking")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.store.CheckBlocking(ctx, first, second)
}
