package main

import (
	"fmt"
	"go-googletranslate/translator"
)

func main() {

	t, err := translator.NewDefaultGoogleTranslate()

	if err != nil {
		panic(err)
	}

	translated, err := t.Translate("Merhaba Dünya!", translator.TURKISH, translator.ENGLISH)
	if err != nil {
		panic(err)
	}

	fmt.Println(translated)

}
