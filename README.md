# Файлы для итогового задания
Backend for the task scheduler.
		The Yandex Practicum Go project is a developer from scratch.
	Tasks with an asterisk have been completed:
Step 1: Determining the port from the outside when starting the server via the TODO_PORT environment variable
Step 2: Determining the path to the database file via the TODO_DBFILE environment variable
Step 3: Assign tasks for the specified days of the week and for the specified days of the month
Step 5: Select a task via the search bar
Step 8: The docker container and authentication are created
	To run tests in tests/settings.go requires the following parameters
var Port = 7540
var DBFile = "../scheduler.db"
var FullNextDate = true
var Search = true
var Token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.gW0hFNtJo0_Q9z1n13dR1SCMfvFQltEh3z9weFL6YQ4`

Keys to run the program:
	password (The password for the app, in this case the password should be 1234567890)
		port (Port for launching the web server 7540  )
			dbpath (The path to the database, the database is created in the current directory )

The `tests` directory contains tests to test the API, which should be implemented in the web server.
	The 'web` directory contains frontend files.
	The `db` directory contains a database and a function for creating and verifying the existence of a database
	The 'env` directory contains functions and parameters for working with environment variables
	The `nextdate` directory contains functions for calculating dates to which tasks with the repeat parameter are transferred
	The `server` directory contains server creation and API functions
	The `structures` directory contains exported structures and constants




Бэкенд для планировщика задач.
Проект Яндекс Практикум Go разработчик с нуля.
Сделаны задания со звездочкой:
Шаг 1: Определение извне порт при запуске сервера через переменную окружения TODO_PORT
Шаг 2: Определение пути к файлу базы данных через переменную окружения TODO_DBFILE
Шаг 3: Назначение задач на указанные дни недели и на указанные дни месяца
Шаг 5: Выбор задачи через строку поиска
Шаг 8: Создан докер контейнер и аутентификация

Для запуска тестов в tests/settings.go необходимы следующие параметры
var Port = 7540
var DBFile = "../scheduler.db"
var FullNextDate = true
var Search = true
var Token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.gW0hFNtJo0_Q9z1n13dR1SCMfvFQltEh3z9weFL6YQ4`

Ключи для запуска программы:
	password (Пароль для приложения 1234567890 )
	port (Порт для запуска веб сервера 7540 )
	dbpath (Путь к базе данных, база создается в текущей директории )

В директории `tests` находятся тесты для проверки API, которое должно быть реализовано в веб-сервере.
Директория `web` содержит файлы фронтенда.
Директория `db` содержит БД и функцию создания и проверки существоваяи БД
Директория `env` содержит функции и параметры для работы с переменными окружения
Директория `nextdate` содержит функции для вычисления дат, на которые переносятся задачи с параметром repeat
Директория `server` содержит функции по созданию сервера и API
Директория `structs` содержит экспортируемые структуры и константы



