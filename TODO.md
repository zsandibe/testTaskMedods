# MEDODS

Test task BackDev Используемые технологии:

    Go
    JWT
    MongoDB

Задание:

Написать часть сервиса аутентификации.

Два REST маршрута:

    Первый маршрут выдает пару Access, Refresh токенов для пользователя сидентификатором (GUID) указанным в параметре запроса
    Второй маршрут выполняет Refresh операцию на пару Access, Refreshтокенов

Требования:

Access токен тип JWT, алгоритм SHA512, хранить в базе строго запрещено.

Refresh токен тип произвольный, формат передачи base64, хранится в базеисключительно в виде bcrypt хеша, должен быть защищен от изменения настороне клиента и попыток повторного использования.

Access, Refresh токены обоюдно связаны, Refresh операцию для Access токена можно выполнить только тем Refresh токеном который был выдан вместе с ним.

Результат:

Результат выполнения задания нужно предоставить в виде исходного кода на Github.

P.S. Друзья! Задания, выполненные полностью или частично с использованием chatGPT видно сразу. Если вы не готовы самостоятельно решать это тестовое задание, то пожалуйста, давайте будем ценить время друг друга и даже не будем пытаться :)