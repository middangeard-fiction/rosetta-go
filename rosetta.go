package rosetta

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/cloudfoundry/jibber_jabber"
)

type message struct {
	Message     string `json:"message"`
	Description string `json:"description,omitempty"`
}

// Translation - localization structure
type translation struct {
	locales      []string
	translations map[string]map[string]message
}

var trans *translation

// InitLocales - initiate locales from the folder
func InitLocales(trPath string) error {
	trans = &translation{translations: make(map[string]map[string]message)}
	return loadTranslations(trPath)
}

// Tr - translate for current locale
func Tr(locale string, trKey string) string {
	trValue, ok := trans.translations[locale][trKey]
	if ok {
		return trValue.Message
	}
	trValue, ok = trans.translations["en"][trKey]
	if ok {
		return trValue.Message
	}
	return trKey
}

// DetectLanguage - parse to find the most preferable language
func DetectLanguage(acceptLanguage string) string {

	langStrs := strings.Split(acceptLanguage, ",")
	for _, langStr := range langStrs {
		lang := strings.Split(strings.Trim(langStr, " "), ";")
		if checkLocale(lang[0]) {
			return lang[0]
		}
	}

	return "en"
}

func GetUILanguage() string {
	userLanguage, err := jibber_jabber.DetectLanguage()
	if err != nil {
		return err.Error()
	}
	return userLanguage
}

// LoadTranslations - load translations files from the folder
func loadTranslations(trPath string) error {
	files, err := filepath.Glob(trPath + "/**/messages.json")
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return errors.New("No translations found")
	}

	for _, file := range files {
		err := loadFileToMap(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadFileToMap(filename string) error {
	var objmap map[string]message

	localName := filepath.Base(filepath.Dir(filename))

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &objmap)
	if err != nil {
		return err
	}
	trans.translations[localName] = objmap
	trans.locales = append(trans.locales, localName)
	return nil
}

func checkLocale(localeName string) bool {
	for _, locale := range trans.locales {
		if locale == localeName {
			return true
		}
	}
	return false
}
