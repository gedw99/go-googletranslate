# 🌍 Go Package for Free and Unlimited Google Translate API

This Go package provides a simple and easy-to-use interface for the Google Translate API, allowing you to translate text between different languages. It also includes a cache feature that stores translations locally to improve performance and reduce the number of API requests made.

## 📦 Installation

To install this package, simply run:

```bash
go get github.com/firattamur/go-googletranslate
```

## 🚀 Usage

Here's an example of how to use the package to translate text:

```go

package main

import (
    "fmt"

   translator "github.com/firattamur/go-googletranslate"
)

func main() {

    text := "Merhaba Dünya!"

    googleTranslator, err := translator.NewDefaultGoogleTranslate()

    if err != nil {
        fmt.Printf("Error creating translator: %v", err)
        return
    }

    translatedText, err := googleTranslator.Translate(text, translator.TURKISH, translator.ENGLISH)

    if err != nil {
        fmt.Printf("Error translating text: %v", err)
        return
    }

    fmt.Printf("Original text   : %s\nTranslated text : %s\n", text, translatedText)

    // Output:
    // Original text   : Merhaba Dünya!
    // Translated text : Hello World!

}

```

To improve performance, the package includes a cache feature that stores translations locally. To configure cache file name and max size, you can use the following methods:

```go

package main

import (
    "fmt"

    translator "github.com/firattamur/go-googletranslate"
)

func main() {

    text := "Merhaba Dünya!"

    cacheFileName := "translations.json"
    maxCacheSize := 100

    googleTranslator, err := translator.NewGoogleTranslate(cacheFileName, maxCacheSize)

    translatedText, err := googleTranslator.Translate(text, translator.TURKISH, translator.ENGLISH)

    if err != nil {
        fmt.Printf("Error translating text: %v", err)
        return
    }

    fmt.Printf("Original text   : %s\nTranslated text : %s\n", text, translatedText)

    // Output:
    // Original text   : Merhaba Dünya!
    // Translated text : Hello World!

}

```

## 📄 License

This project is licensed under the MIT License. The terms are as follows:

```text
The MIT License (MIT)

Copyright (c) 2015 SuHun Han

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
