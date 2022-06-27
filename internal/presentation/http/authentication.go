package http

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	dto "github.com/mohammaderm/krad/internal/dto/user"
	"github.com/mohammaderm/krad/internal/service/user"
	"github.com/mohammaderm/krad/log"
	"golang.org/x/crypto/bcrypt"
)

const (
	Secretkey = "mohammad123456789@"
	Issue     = "127.0.1.1"
)

type (
	AuthtHandler struct {
		logger      log.Logger
		userService user.UserServiceContracts
		HandlerHelper
	}
	AuthHandlerContract interface {
		Login(w http.ResponseWriter, r *http.Request)
		Register(w http.ResponseWriter, r *http.Request)
	}
	JwtClaims struct {
		Email string `json:"email"`
		Id    uint   `json:"id"`
		jwt.StandardClaims
	}
)

func NewAuthHanlder(logger log.Logger, userservice user.UserServiceContracts) AuthHandlerContract {
	return &AuthtHandler{
		logger:      logger,
		userService: userservice,
		HandlerHelper: HandlerHelper{
			logger: logger,
		},
	}
}

func (a *AuthtHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userreq dto.CreateUserReq
	err := a.readJSON(w, r, &userreq)
	if err != nil {
		a.errorJSON(w, errors.New("can not parse values"), http.StatusBadRequest)
		return
	}
	hashpass, err := bcrypt.GenerateFromPassword([]byte(userreq.Password), bcrypt.DefaultCost)
	if err != nil {
		a.errorJSON(w, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}
	if err := a.userService.Create_User(r.Context(), dto.CreateUserReq{
		UserName: userreq.UserName,
		Email:    userreq.Email,
		Password: string(hashpass),
	}); err != nil {
		a.errorJSON(w, errors.New("can not create user"), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	payload := jsonResponse{
		Error:   false,
		Message: "user saved succesfully",
		Data:    userreq,
	}
	a.writeJSON(w, http.StatusOK, payload)

}

func (a *AuthtHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user dto.UserLogin
	err := a.readJSON(w, r, &user)
	if err != nil {
		a.errorJSON(w, errors.New("can not parse values"), http.StatusBadRequest)
		return
	}
	founduser, err := a.userService.GetbyEmail_User(r.Context(), dto.GetByEmailReq{Email: user.Email})
	if err != nil {
		a.errorJSON(w, errors.New("invalid credential"), http.StatusForbidden)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(founduser.User.Password), []byte(user.Password))
	if err != nil {
		a.errorJSON(w, errors.New("wrong password"), http.StatusForbidden)
		return
	}
	pairtoken, err := generatepairtoken(founduser.User.Email, founduser.User.Id)
	if err != nil {
		a.errorJSON(w, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	payload := jsonResponse{
		Error:   false,
		Message: "Logged in user",
		Data:    pairtoken,
	}
	a.writeJSON(w, http.StatusAccepted, payload)
}

func generatepairtoken(email string, id uint) (map[string]string, error) {
	// Access_token
	a_claims := JwtClaims{
		Email: email,
		Id:    id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    Issue,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, a_claims)
	a_token, err := jwtToken.SignedString([]byte(Secretkey))
	if err != nil {
		return nil, err
	}
	// Refresh Token
	r_claims := JwtClaims{
		Email: email,
		Id:    id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    Issue,
		},
	}
	r_jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, r_claims)
	r_token, err := r_jwtToken.SignedString([]byte(Secretkey))
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"access_token":  a_token,
		"refresh_token": r_token,
	}, nil

}
