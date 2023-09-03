## Setup

* run go mod
    ```
    go mod init immersive-dash-4
    ```

* download echo
    ```
    go get -u github.com/labstack/echo/v4
    ```

* download gorm
    ```
    go get -u gorm.io/gorm
    go get -u gorm.io/driver/mysql
    ```

* download viper (to load .env automatically)
    ```
    go get github.com/spf13/viper
    ```
 * download bcrypt for hashing password
```
golang.org/x/crypto/bcrypt
```

* download validator v10 untuk validasi struct
```
github.com/go-playground/validator/v10
```
* create file `local.env`
    ```
    export DBUSER='root'
    export DBPASS='qwerty123'
    export DBHOST='127.0.0.1'
    export DBPORT='3306'
    export DBNAME='db_loanee_gorm'
    ```

## Task
* tambahkan endpoint untuk CRUD user dan item
* buat repo pengumpulan tugas dengan nama `rest-api-clean-arch`
    ```
    UPDATE PUT /users
    DELETE DELETE /users/:user_id
    POST /items
    GET /items
    GET /items/:item_id
    PUT /items/:item_id
    DELETE /items/:item_id
    ```