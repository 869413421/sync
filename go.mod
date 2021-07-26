module sync

go 1.16

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/casbin/casbin/v2 v2.31.9
	github.com/casbin/gorm-adapter/v3 v3.3.1
	github.com/denisenkom/go-mssqldb v0.10.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.2
	github.com/go-mysql-org/go-mysql v1.3.0 // indirect
	github.com/go-playground/validator/v10 v10.6.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jackc/pgproto3/v2 v2.1.0 // indirect
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/siddontang/go-mysql v1.3.0
	github.com/stretchr/testify v1.7.0
	github.com/thedevsaddam/govalidator v1.9.10
	github.com/ugorji/go v1.2.6 // indirect
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22 // indirect
	google.golang.org/protobuf v1.27.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	gorm.io/driver/mysql v1.1.1
	gorm.io/driver/postgres v1.1.0 // indirect
	gorm.io/driver/sqlserver v1.0.7 // indirect
	gorm.io/gorm v1.21.11
)

replace github.com/siddontang/go-mysql v1.3.0 => github.com/go-mysql-org/go-mysql v1.3.0

replace github.com/go-mysql-org/go-mysql v1.3.0 => github.com/siddontang/go-mysql v1.3.0
