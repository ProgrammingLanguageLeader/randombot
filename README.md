# RandomBot
This is a Telegram bot that enables you to generate some kind of pseudorandom values. 
For example, it has functionality for virtual "rolling the dice", "flipping coin", choice between set of options, etc.  

## Requirements
- Golang compiler (with 1.13 or higher version)
- Docker CLI
- Redis (you may install it via Docker)

## How to launch locally
- Install a Redis server:
    ```bash
    docker run --name redis-test-instance -p 6379:6379 -d redis
    ```

- Start the Redis server:
    ```bash
    docker start redis-test-instance
    ```

- Create a .env file and define Telegram token parameter:
    ```dotenv
    RANDOM_BOT_TELEGRAM_TOKEN=<your token>
    ```

- Build and launch an application

## How to build

### Building with Docker
Run the following command:
```bash
docker build -t <build tag> .
```
You will get an application image. Then if need to run it, use "docker run" command. For example:
```bash
docker run --name redis-test-instance <build tag> \
    -e RANDOM_BOT_TELEGRAM_TOKEN=<token>
```

### Building with Golang compiler
Run the following command:
```bash
go build
```
You will get an executable file with name "main". 

Launching the application:
```bash
./main
```

## Localization

### Translating new strings
- Find new strings to translate:
    ```bash
    goi18n extract -outdir locale/assets
    goi18n merge -outdir locale/assets locale/assets/active.*.toml
    ```
    You will get "translate.\<language code\>.toml" files with strings that translations need to be updated.
    
- Edit translations files

- Merge the translations into active files:
    ```bash
    goi18n merge -outdir locale/assets locale/assets/active.*.toml locale/assets/translate.*.toml
    ```

- (Optional) Remove "translate.*.toml" files:
    ```bash
    rm locale/assets/translate.*.toml
    ```

### Translating new language
- Generate a file with all strings to translate:
    ```bash
    goi18n extract -outdir locale/assets -sourceLanguage <language code>
    ```

- Edit the translations file

- Merge translations (this command will append hash codes of the strings):
    ```bash
    goi18n merge -outdir locale/assets locale/assets/active.*.toml
    ```
