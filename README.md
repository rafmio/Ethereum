Etherium test
Создать CLI приложение

Используемое (примеры):
- https://goethereumbook.org/en/
- https://pkg.go.dev/crypto/sha256
- https://pkg.go.dev/crypto/aes

Сделать CLI
При запуске приложение:
	- проверяет наличие аккаунтов
	- если есть  - просит пароль
	- если нет 	 - просит его создать

Аккаунт == созданный кошелек
Пароль нужно хранить в файле
Пароль нужно держать в виде хэша, чтоб при получении можно было сравнить с хэшем

Данные в файле с хэшем нужно шифровать с использованием AES.
Ключ для шифрования может быть захардкоденным

Если пароль сходится с хэшем из зашифр.файла - то этот пароль используется для
разблокировки кошелька


===============================================================================
Оригинал текста задания:

Привет, как насчёт прохождения тестового задания?:) вот некоторые вещи, которые
у нас используются:
https://goethereumbook.org/en/
https://pkg.go.dev/crypto/sha256
https://pkg.go.dev/crypto/aes
Через либу go Ethereum можно создавать кошельки, там генерится пароль для них.
В дальнейшем этот пароль можно использовать для доступа к кошельку.
Давай сделаем CLI приложение небольшое, которое когда запускаешь, если нет
аккаунтов, просит его создавать, а если есть - просит пароль.
То есть аккаунт в данном случае это будет созданный кошелек.
Пароль который передается нужно в файле где нибудь держать в виде хэша,
чтобы когда получен пароль, можно было сравнить его с хэшом и понять,
тот он или не тот.
Данные в файле с хэшом надо шифровать с использованием AES, ключ для шифрования
можно использовать какой-нить захардкоденный. Если пароль сходится с хэшом из
зашифрованного файла, то он используется для разблокирования кошелька.
Вопросы/результаты можно отправлять в телеграм на @r4m1_l.
По времени, чем раньше получится, тем лучше:)
P.S. мы в настоящее время рассматриваем только кандидатов в офис.
