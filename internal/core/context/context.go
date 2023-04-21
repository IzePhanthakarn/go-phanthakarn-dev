package context

import (
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/config"
	"github.com/IzePhanthakarn/go-phanthakarn-dev/internal/core/database"
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

const (
	pathKey            = "path"
	compositeFormDepth = 3
	// UserKey user key
	UserKey = "user"
	// LangKey lang key
	LangKey = "lang"
	// DatabaseKey database key
	DatabaseKey = "database"
	// ParametersKey parameters key
	ParametersKey = "parameters"
	// SessionTokenKey session token key
	SessionTokenKey = "sessionToken"
)

// Context context
type Context struct {
	*fiber.Ctx
}

// New new custom fiber context
func New(c *fiber.Ctx) *Context {
	return &Context{c}
}

// BindValue bind value
func (c *Context) BindValue(i interface{}, validate bool) error {
	switch c.Method() {
	case http.MethodGet:
		_ = c.QueryParser(i)

	default:
		_ = c.BodyParser(i)
	}

	c.PathParser(i, 1)
	c.Locals(ParametersKey, i)
	TrimSpace(i, 1)

	if validate {
		err := c.Validate(i)
		if err != nil {
			return err
		}
	}
	return nil
}

// PathParser parse path param
func (c *Context) PathParser(i interface{}, depth int) {
	formValue := reflect.ValueOf(i)
	if formValue.Kind() == reflect.Ptr {
		formValue = formValue.Elem()
	}
	t := reflect.TypeOf(formValue.Interface())
	for i := 0; i < t.NumField(); i++ {
		fieldName := t.Field(i).Name
		paramValue := formValue.FieldByName(fieldName)
		if paramValue.IsValid() {
			if depth < compositeFormDepth && paramValue.Kind() == reflect.Struct {
				depth++
				c.PathParser(paramValue.Addr().Interface(), depth)
			}
			tag := t.Field(i).Tag.Get(pathKey)
			if tag != "" {
				setValue(paramValue, c.Params(tag))
			}
		}
	}
}

func setValue(paramValue reflect.Value, value string) {
	if paramValue.IsValid() && value != "" {
		switch paramValue.Kind() {
		case reflect.Uint:
			number, _ := strconv.ParseUint(value, 10, 32)
			paramValue.SetUint(number)

		case reflect.String:
			paramValue.SetString(value)

		default:
			number, err := strconv.Atoi(value)
			if err != nil {
				paramValue.SetString(value)
			} else {
				paramValue.SetInt(int64(number))
			}
		}
	}
}

// Validate validate
func (c *Context) Validate(i interface{}) error {
	if err := config.CF.Validator.Struct(i); err != nil {
		return config.RR.CustomMessage(err.Error(), err.Error()).WithLocale(c.Ctx)
	}

	return nil
}

func (c *Context) trimspace(i interface{}) {
	e := reflect.ValueOf(i).Elem()
	for i := 0; i < e.NumField(); i++ {
		if e.Type().Field(i).Type.Kind() != reflect.String {
			continue
		}

		value := e.Field(i).Interface().(string)
		e.Field(i).SetString(strings.TrimSpace(value))
	}
}

// Claims jwt claims
type Claims struct {
	jwt.StandardClaims
	RefreshTokenID uint
	Roles          []int
}

// GetClaims get user claims
func (c *Context) GetClaims() *Claims {
	user := c.Locals(UserKey).(*jwt.Token)
	return user.Claims.(*Claims)
}

// GetUserID get user id claims
func (c *Context) GetUserID() uint {
	token, ok := c.Locals(UserKey).(*jwt.Token)
	if ok {
		cl := token.Claims.(*Claims)
		if cl != nil {
			i, _ := strconv.ParseUint(c.GetClaims().Subject, 10, 64)
			return uint(i)
		}
	}

	return 0
}

// GetAccessToken get access token
func (c *Context) GetAccessToken() string {
	token, ok := c.Locals(UserKey).(*jwt.Token)
	if ok {
		return token.Raw
	}

	return ""
}

// GetSessionToken get session token from header
func (c *Context) GetSessionToken() string {
	token, ok := c.Locals(SessionTokenKey).(string)
	if ok {
		return token
	}

	return ""
}

// GetDatabase get connection database
func (c *Context) GetDatabase() *gorm.DB {
	val := c.Locals(DatabaseKey)
	if val == nil {
		return database.Database
	}

	return val.(*gorm.DB)
}

// TrimSpace trim space
func TrimSpace(i interface{}, depth int) {
	e := reflect.ValueOf(i).Elem()
	for i := 0; i < e.NumField(); i++ {
		if depth < 3 && e.Type().Field(i).Type.Kind() == reflect.Struct {
			depth++
			TrimSpace(e.Field(i).Addr().Interface(), depth)
		}

		if e.Type().Field(i).Type.Kind() != reflect.String {
			continue
		}

		value := e.Field(i).String()
		e.Field(i).SetString(strings.TrimSpace(value))
	}
}

// GetLevel get level
func (c *Context) GetLevel() string {
	return c.Get("Level")
}
