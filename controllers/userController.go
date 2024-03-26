package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/Desmond51/go-jwt-project/database"
	"github.com/gin-gonic/gin"
	"github.com/Desmond51/go-jwt-project/helpers"
	helper "github.com/Desmond51/go-jwt-project/helpers"
	"github.com/Desmond51/go-jwt-project/models"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword()

func VerifyPassword()

func Signup()

func Login()

func GetUsers()

func GetUser()