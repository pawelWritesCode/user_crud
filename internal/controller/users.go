package controller

import (
	"encoding/xml"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v2"

	"example-user-crud/internal/models"
	"example-user-crud/internal/repository"
)

const (
	formatJSON = "json"
	formatYAML = "yaml"
	formatXML  = "xml"
)

type Error string

type UsersResponse struct {
	XMLName xml.Name `xml:"users"`
	Users   []models.User
}

func CreateUser(ur repository.UsersRepository) func(context echo.Context) error {
	return func(context echo.Context) error {
		format := context.QueryParam("format")
		var u models.User
		if err := context.Bind(&u); err != nil {
			if format == formatJSON {
				return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			}

			if format == formatYAML {
				yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
				context.Response().Header().Set("Content-Type", "application/x-yaml")
				return context.String(http.StatusBadRequest, string(yamlResopnse))
			}

			if format == formatXML {
				return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
			}

			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		user, err := ur.Create(u)
		if err != nil {
			if format == formatJSON {
				return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			}

			if format == formatYAML {
				yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
				context.Response().Header().Set("Content-Type", "application/x-yaml")
				return context.String(http.StatusBadRequest, string(yamlResopnse))
			}

			if format == formatXML {
				return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
			}

			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		if format == formatJSON {
			return context.JSON(http.StatusCreated, user)
		}

		if format == formatYAML {
			yamlResopnse, _ := yaml.Marshal(user)
			context.Response().Header().Set("Content-Type", "application/x-yaml")
			return context.String(http.StatusCreated, string(yamlResopnse))
		}

		if format == formatXML {
			return context.XMLPretty(http.StatusCreated, user, "\t")
		}

		return context.JSON(http.StatusCreated, user)
	}
}

func GetUser(ur repository.UsersRepository) func(context echo.Context) error {
	return func(context echo.Context) error {
		uidString := context.Param("userId")
		uid, err := strconv.Atoi(uidString)
		format := context.QueryParam("format")
		if err != nil {
			if format == formatJSON {
				return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			}

			if format == formatYAML {
				yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
				context.Response().Header().Set("Content-Type", "application/x-yaml")
				return context.String(http.StatusBadRequest, string(yamlResopnse))
			}

			if format == formatXML {
				return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
			}

			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		user, err := ur.FindOne(uid)
		if err != nil {
			if format == formatJSON {
				return context.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			}

			if format == formatYAML {
				yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
				context.Response().Header().Set("Content-Type", "application/x-yaml")
				return context.String(http.StatusNotFound, string(yamlResopnse))
			}

			if format == formatXML {
				return context.XMLPretty(http.StatusNotFound, Error(err.Error()), "\t")
			}

			return context.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		}

		if format == formatJSON {
			return context.JSON(http.StatusOK, user)
		}

		if format == formatYAML {
			yamlResopnse, _ := yaml.Marshal(user)
			context.Response().Header().Set("Content-Type", "application/x-yaml")
			return context.String(http.StatusOK, string(yamlResopnse))
		}

		if format == formatXML {
			return context.XMLPretty(http.StatusOK, user, "\t")
		}

		return context.JSON(http.StatusOK, user)
	}
}

func GetUsers(ur repository.UsersRepository) func(context echo.Context) error {
	return func(context echo.Context) error {
		users := ur.GetAll()

		format := context.QueryParam("format")
		if format == formatJSON {
			return context.JSON(http.StatusOK, users)
		}

		if format == formatYAML {
			yamlResopnse, _ := yaml.Marshal(users)
			context.Response().Header().Set("Content-Type", "application/x-yaml")
			return context.String(http.StatusOK, string(yamlResopnse))
		}

		if format == formatXML {
			return context.XMLPretty(http.StatusOK, UsersResponse{Users: users}, "\t")
		}

		return context.JSON(http.StatusOK, users)
	}
}

func ReplaceUser(ur repository.UsersRepository) func(context echo.Context) error {
	return func(context echo.Context) error {
		format := context.QueryParam("format")
		uidString := context.Param("userId")
		uid, err := strconv.Atoi(uidString)
		if err != nil {
			if format == formatJSON {
				return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			}

			if format == formatYAML {
				yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
				context.Response().Header().Set("Content-Type", "application/x-yaml")
				return context.String(http.StatusBadRequest, string(yamlResopnse))
			}

			if format == formatXML {
				return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
			}

			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		var u models.User
		if err := context.Bind(&u); err != nil {
			if format == formatJSON {
				return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			}

			if format == formatYAML {
				yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
				context.Response().Header().Set("Content-Type", "application/x-yaml")
				return context.String(http.StatusBadRequest, string(yamlResopnse))
			}

			if format == formatXML {
				return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
			}

			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		if err := ur.Replace(uid, u); err != nil {
			if format == formatJSON {
				return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			}

			if format == formatYAML {
				yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
				context.Response().Header().Set("Content-Type", "application/x-yaml")
				return context.String(http.StatusBadRequest, string(yamlResopnse))
			}

			if format == formatXML {
				return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
			}

			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		if format == formatJSON {
			return context.JSON(http.StatusOK, map[string]interface{}{})
		}

		if format == formatYAML {
			context.Response().Header().Set("Content-Type", "application/x-yaml")
			return nil
		}

		if format == formatXML {
			return context.XMLPretty(http.StatusOK, "", "\t")
		}

		return context.JSON(http.StatusOK, map[string]interface{}{})
	}
}

func DeleteUser(ur repository.UsersRepository) func(context echo.Context) error {
	return func(context echo.Context) error {
		format := context.QueryParam("format")
		uidString := context.Param("userId")
		uid, err := strconv.Atoi(uidString)
		if err != nil {
			if format == formatJSON {
				return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			}

			if format == formatYAML {
				yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
				context.Response().Header().Set("Content-Type", "application/x-yaml")
				return context.String(http.StatusBadRequest, string(yamlResopnse))
			}

			if format == formatXML {
				return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
			}

			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		if err := ur.Delete(uid); err != nil {
			if format == formatJSON {
				return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			}

			if format == formatYAML {
				yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
				context.Response().Header().Set("Content-Type", "application/x-yaml")
				return context.String(http.StatusBadRequest, string(yamlResopnse))
			}

			if format == formatXML {
				return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
			}

			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		return context.NoContent(http.StatusNoContent)
	}
}

func ReceiveAvatar(ur repository.UsersRepository) func(context echo.Context) error {
	return func(context echo.Context) error {
		format := context.QueryParam("format")
		uidString := context.Param("userId")
		uid, err := strconv.Atoi(uidString)
		if err != nil {
			if format == formatJSON {
				return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			}

			if format == formatYAML {
				yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
				context.Response().Header().Set("Content-Type", "application/x-yaml")
				return context.String(http.StatusBadRequest, string(yamlResopnse))
			}

			if format == formatXML {
				return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
			}

			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		name := context.FormValue("name")
		if len(name) == 0 {
			if format == formatJSON {
				return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": "name should not be empty string and should contain extension part"})
			}

			if format == formatYAML {
				yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": "name should not be empty string and should contain extension part"})
				context.Response().Header().Set("Content-Type", "application/x-yaml")
				return context.String(http.StatusBadRequest, string(yamlResopnse))
			}

			if format == formatXML {
				return context.XMLPretty(http.StatusBadRequest, Error("name should not be empty string and should contain extension part"), "\t")
			}

			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": "name should not be empty string and should contain extension part"})
		}

		avatar, err := context.FormFile("avatar")
		if err != nil {
			if format == formatJSON {
				return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			}

			if format == formatYAML {
				yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
				context.Response().Header().Set("Content-Type", "application/x-yaml")
				return context.String(http.StatusBadRequest, string(yamlResopnse))
			}

			if format == formatXML {
				return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
			}

			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		src, err := avatar.Open()
		if err != nil {
			if format == formatJSON {
				return context.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "a" + err.Error()})
			}

			if format == formatYAML {
				yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": "a" + err.Error()})
				context.Response().Header().Set("Content-Type", "application/x-yaml")
				return context.String(http.StatusInternalServerError, string(yamlResopnse))
			}

			if format == formatXML {
				return context.XMLPretty(http.StatusInternalServerError, Error(err.Error()), "\t")
			}

			return context.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "a" + err.Error()})
		}

		defer src.Close()

		fileName := uidString + "_" + name

		user, err := ur.FindOne(uid)
		if err != nil {
			if format == formatJSON {
				return context.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			}

			if format == formatYAML {
				yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
				context.Response().Header().Set("Content-Type", "application/x-yaml")
				return context.String(http.StatusNotFound, string(yamlResopnse))
			}

			if format == formatXML {
				return context.XMLPretty(http.StatusNotFound, Error(err.Error()), "\t")
			}

			return context.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		}

		user.Avatar = fileName[len(uidString)+1:]
		err = ur.Replace(uid, user)
		if err != nil {
			if format == formatJSON {
				return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			}

			if format == formatYAML {
				yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
				context.Response().Header().Set("Content-Type", "application/x-yaml")
				return context.String(http.StatusBadRequest, string(yamlResopnse))
			}

			if format == formatXML {
				return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
			}

			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		dst, err := os.Create(path.Join(os.TempDir(), fileName))
		if err != nil {
			if format == formatJSON {
				return context.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "b" + err.Error()})
			}

			if format == formatYAML {
				yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": "b" + err.Error()})
				context.Response().Header().Set("Content-Type", "application/x-yaml")
				return context.String(http.StatusInternalServerError, string(yamlResopnse))
			}

			if format == formatXML {
				return context.XMLPretty(http.StatusInternalServerError, Error(err.Error()), "\t")
			}

			return context.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "b" + err.Error()})
		}

		if _, err = io.Copy(dst, src); err != nil {
			if format == formatJSON {
				return context.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "c" + err.Error()})
			}

			if format == formatYAML {
				yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": "c" + err.Error()})
				context.Response().Header().Set("Content-Type", "application/x-yaml")
				return context.String(http.StatusInternalServerError, string(yamlResopnse))
			}

			if format == formatXML {
				return context.XMLPretty(http.StatusInternalServerError, Error(err.Error()), "\t")
			}

			return context.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "c" + err.Error()})
		}

		return nil
	}
}
