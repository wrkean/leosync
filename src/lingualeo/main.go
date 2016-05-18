package lingualeo

import (
	"github.com/igrybkov/leosync/src/configuration"
	"github.com/igrybkov/leosync/src/lingualeo/api"
	"log"
)

var leoClient api.Client

func getClient() api.Client {
	if leoClient == (api.Client{}) {
		var err error

		config := configuration.GetConfig()

		leoClient, err = api.NewClient(config.LinguaLeo.Email, config.LinguaLeo.Password)
		if err != nil {
			log.Fatalf("%v \n", err)
		}
	}
	return leoClient
}

func GetTranslations(word string) api.Word {
	translations, errs := getClient().GetTranslations(word)
	if errs != nil {
		log.Fatalf("%v \n", errs)
	}
	return translations
}

func AddWordWithTranslation(word string, translation string) error {
	_, err := getClient().AddWord(word, translation)
	if err != nil {
		log.Fatalf("%v \n", err)
	}
	return err
}

func AddWord(word string) {
	translations := GetTranslations(word)
	if len(translations.Translations) == 0 {
		log.Fatalf("Translation not found for word \"%s\"\n", word)
	}
	translation := translations.Translations[0].Value
	errs := AddWordWithTranslation(word, translation)
	if errs != nil {
		log.Fatalf("Cannot add word: %v", errs)
	}
}
