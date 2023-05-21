## Run
```bash
$ go run main.go message.go 
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
## Generate by message.go(use i18n.Message struct)
The following command will update `i18n/active.en.json`.
```bash
$  goi18n extract -sourceLanguage=ja -format=json --outdir=i18n
```

## Add message
1. Generate `i18n/translate.*.json`
    ```bash
    $ goi18n merge -sourceLanguage=ja -format=json --outdir=i18n  i18n/active.*.json
    ```

2. Translate the contents of `i18n/translate.*.json`.

3. Merge `i18n/translate.*.json` to `i18n/active.*.json`
    ```bash
    $ goi18n merge -sourceLanguage=ja -format=json --outdir=i18n  i18n/active.*.json i18n/translate.*.json
    ```

## Add new language
1. Create `i18n/translate.xx.json`. file (e.g. `i18n/translate.fr.json)
2. Execute the following commands.  
    This will result in the contents of `i18n/active.en.json` being transcribed to `i18n/translate.xx.json`.
    ```bash
    $ goi18n merge -sourceLanguage=ja -format=json --outdir=i18n  i18n/active.ja.json i18n/translate.fr.json
    ```

3. Translate the contents of `i18n/translate.xx.json`.
4. Merge `i18n/translate.xx.json` to `i18n/active.xx.json`
    ```
    $ goi18n merge -sourceLanguage=ja -format=json --outdir=i18n  i18n/active.*.json i18n/translate.*.json
    ```
5. Add a process to read the language file (`i18n/translate.xx.json`) added from the program.  
    For example, the code might look like this
    ```bash
    bundle.LoadMessageFileFS(LocaleFS, "i18n/active.xx.json")
    ```