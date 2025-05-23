// Copyright (c) 2025 Grigoriy Efimov
//
// Licensed under the MIT License. See LICENSE file in the project root for details.

package main

import (
	"embed"
	"encoding/json"

	ink "github.com/CatInBeard/inkview"
)

//go:embed translations
var translations embed.FS

type Translation struct {
	Lang         string
	Translations map[string]string
}

var translationsMap = make(map[string]*Translation)

var currentLang = ink.GetCurrentLang()

func loadTranslation(lang string) (*Translation, error) {
	if translation, ok := translationsMap[lang]; ok {
		return translation, nil
	}
	data, err := translations.ReadFile("translations/" + lang + ".json")
	if err != nil {
		return nil, err
	}
	translation := Translation{Lang: lang}
	err = json.Unmarshal(data, &translation.Translations)
	if err != nil {
		return nil, err
	}
	translationsMap[lang] = &translation
	return &translation, nil
}

func GetTranslation(lang, key string) string {
	translation, err := loadTranslation(lang)
	if err != nil {
		translation, err = loadTranslation("en")
		if err != nil {
			return key
		}
	}
	value, ok := translation.Translations[key]
	if !ok {
		jsonData, _ := json.Marshal(translation.Translations)
		return string(jsonData)
	}
	return value
}

func GetCurrentTranslation(key string) string {
	return GetTranslation(currentLang, key)
}
