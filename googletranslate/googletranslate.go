package googletranslate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/firattamur/go-googletranslate/filecache"
	"github.com/firattamur/go-googletranslate/languages"
)

type GoogleTranslate struct {
	Cache *filecache.FileCache
}

const unofficialAPIURL = "https://translate.googleapis.com/translate_a/single?client=gtx&sl=%s&tl=%s&dt=t&q=%s"

func NewGoogleTranslate(cacheFileName string, cacheMaxSize int) (*GoogleTranslate, error) {
	cache, err := filecache.NewFileCache(cacheFileName, cacheMaxSize)
	if err != nil {
		return nil, err
	}
	return &GoogleTranslate{cache}, nil
}

func (gt *GoogleTranslate) Translate(text string, sourceLanguage languages.LanguageCode, targetLanguage languages.LanguageCode) (string, error) {

	if translatedText, ok := gt.GetFromCache(text, sourceLanguage, targetLanguage); ok {
		return translatedText, nil
	}

	escapedText := url.QueryEscape(text)
	apiURL := fmt.Sprintf(unofficialAPIURL, sourceLanguage.String(), targetLanguage.String(), escapedText)

	response, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var resp []interface{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return "", err
	}

	if len(resp) > 0 {
		if nestedResp, ok := resp[0].([]interface{}); ok && len(nestedResp) > 0 {
			if translatedText, ok := nestedResp[0].([]interface{}); ok && len(translatedText) > 0 {

				err = gt.SetToCache(text, translatedText[0].(string), sourceLanguage, targetLanguage)
				return translatedText[0].(string), nil

			}

		}

	}

	return "", fmt.Errorf("Unable to parse translation response")

}

func (gt *GoogleTranslate) GetFromCache(text string, sourceLanguage languages.LanguageCode, targetLanguage languages.LanguageCode) (string, bool) {
	key := fmt.Sprintf("%s-%s-%s", text, sourceLanguage, targetLanguage)
	return gt.Cache.Get(key)
}

func (gt *GoogleTranslate) SetToCache(text, translatedText string, sourceLanguage languages.LanguageCode, targetLanguage languages.LanguageCode) error {
	key := fmt.Sprintf("%s-%s-%s", text, sourceLanguage, targetLanguage)
	return gt.Cache.Set(key, translatedText)
}
