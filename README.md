Etherium Contracts Application

Create application

Examples of packages used:
- https://goethereumbook.org/en/
- https://pkg.go.dev/crypto/sha256
- https://pkg.go.dev/crypto/aes

Develop an application
When launching the application:	
	- check for existing accounts	
	- if it is exists - then request a password
	- if it is not exists - then request to create it

Account  == just created wallet
The password must be stored at PostgreSQL database
The password must be stored as a hash


Creating account (wallet) - Issue #1

Optional:
Данные в файле с хэшем нужно шифровать с использованием AES.
Ключ для шифрования может быть захардкоденным

Если пароль сходится с хэшем из зашифр.файла - то этот пароль используется для
разблокировки кошелька


=================================================================
