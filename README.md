# Тестовое задание для [Theboats](https://theboats.com/)

Текст самого задания: [посмотреть](/task.md)

# Установка
1. `mkdir -p $GOPATH/src/github.com/Madredix/theboats && cd $GOPATH/src/github.com/Madredix/theboats`
2. `git clone https://github.com/Madredix/theboats ./`
3. `curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh` при отсутствии dep
4. `dep ensure`
5. отредактировать config/config.json
6. `go build -o app src/main.go`

# Запуск
* `./app --config=config/config.json` - автоматически накатит миграцию, запустит первое скачивание с гдс, запустит веб-сервер
* `./app --help` - посмотреть доступные команды

# Запуск в Docker
1. `mkdir -p $GOPATH/src/github.com/Madredix/theboats && cd $GOPATH/src/github.com/Madredix/theboats`
2. `git clone https://github.com/Madredix/theboats ./`
3. `docker build -t theboats . && docker run -it theboats`

# HTTP Api
* `http://localhost:2222/api/v1/search?q=bavaria` - поиск яхт (меньше 3х символов искать не будет)
* `http://localhost:2222/api/v1/autocomplete?q=bavaria` - автокомплит (меньше 3х символов искать не будет)


## Не доделано
1. отсутствует front
2. отсутствуют бенчмарк тесты + тестов явно мало (успел сделать только для gds)
3. какая-то проблема с http сервером - ошибка при shutdown
