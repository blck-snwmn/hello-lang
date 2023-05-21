## Run
```bash
go run main.go
```

```bash
$ curl http://localhost:8080/greet

こんにちわ John.
```

```bash
$ curl http://localhost:8080/greet?lang=ja

こんにちわ John.
```

```bash
$ curl http://localhost:8080/greet?lang=en

Hello John.
```

## Add message
1. Generate `i18n/translate.*.json`
    ```
    $ goi18n merge -sourceLanguage=ja -format=json --outdir=i18n  i18n/active.*.json
    ```

2. Translate the contents of `i18n/translate.*.json`.

3. Merge `i18n/translate.*.json` to `i18n/active.*.json`
    ```
    $ goi18n merge -sourceLanguage=ja -format=json --outdir=i18n  i18n/active.*.json i18n/translate.*.json
    ```



## Add new language
