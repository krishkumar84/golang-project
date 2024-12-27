package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/krishkumar84/golang-project/pkg/storage"
	"github.com/krishkumar84/golang-project/pkg/types"
	"github.com/krishkumar84/golang-project/pkg/utils/response"
)



func New(storage storage.Storage) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("New User Handler")
        var user types.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if errors.Is(err, io.EOF){ 
			response.WriteJson(w , http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w,http.StatusBadRequest,response.GeneralError(err))
			return
		}

		if err := validator.New().Struct(user); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		lastId, err := storage.CreateUser(
			user.Name,
			user.Email,
			user.Age,
		)

		slog.Info("User created sucessfully",slog.String("userId",fmt.Sprint(lastId)))

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]int64{"id":lastId})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		id := r.PathValue("id")
		slog.Info("Get User by Id",slog.String("id",id))
		intId, err := strconv.ParseInt(id,10,64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		user, err := storage.GetUserById(intId)
		if err != nil {
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(err))
			return
		}
		response.WriteJson(w, http.StatusOK, user)
	}
}

func GetAll(storage storage.Storage) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		users, err := storage.GetAllUsers()
		if err != nil {
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(err))
			return
		}
		response.WriteJson(w, http.StatusOK, users)
	}
}

