Рахмет что проверяете мою ТЗшку!

Как запустить?

В корневой папке проекта:

0. Создать файлы .env в корневой папке проекта и в ./stress-test (оставил .env.dist, можно просто его переименовать в .env)
1. Запуск бэка: docker-compose up --build (зависимости подтянутся из vendor)
2. Запуск нагрузочного скрипта: cd ./stress-test && go run .

Также для удобства есть swagger-документашка, лежит по пути localhost:8080/swagger/index.html

Made by Ashim Zhaksylyk
