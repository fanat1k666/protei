# protei

## Функционал
В рамаках данного приложения, мы имеем возможность по входящему gRPC запросу с
информацией о пользователе по email найти того же пользователя на внешнем HTTP
сервере, и обогатить входящее имя пользователя (ФИО) статусом отсутствия, если оно
есть (в конец ФИО дописать emoji с соответствующим статусом).

## Пример GRPC запроса
{
  "mail": "exapmple@gmail.com"
}
### Пример ожидаемого ответа
Ivanov Ivan Ivanovich☀️

## Конфигурация (необходимо указать свои параметры)

### "crdntls" отвечает за параметры для коннекта к внешнему http, где:
  "address" - адресс
  "port" - порт
  "login" - логин, для авторизации к внешнему апи
  "password" - пароль, для авторизации к внешнему апи

### "server" отвечает за парметры сервера
  "port" - порт

## Для запуска необходимо:
  подтянем зависимости 
  go mod tidy
  ### запустить сервер
  go build protei/cmd/server
  ### запустить клиент
  go build protei/cmd/client

  ### Если у вас при запуске клиента произошла ошибка "could not greet: rpc error: code = DeadlineExceeded desc = context deadline exceeded",
  тогда проверьте конфигурацию, вероятно вы неверно указали значения

  ### Запустить линтер 
  golangci-lint run
