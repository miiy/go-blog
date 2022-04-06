module goblog.com/web

go 1.17

replace (
	goblog.com/pkg => ../pkg
	goblog.com/service/article => ../service/article
)

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-playground/validator/v10 v10.4.1
	github.com/go-redis/redis/v8 v8.11.4
	github.com/google/wire v0.5.0
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.21.0
	goblog.com/pkg v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20220307211146-efcb8507fb70
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/gorm v1.23.3
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/gin-contrib/cors v1.3.1 // indirect
	github.com/gin-contrib/multitemplate v0.0.0-20220323084503-710510e67c20 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.10.1 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	goblog.com/service/article v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220307203707-22a9840ba4d7 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220322021311-435b647f9ef2 // indirect
	google.golang.org/grpc v1.45.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	gorm.io/driver/mysql v1.3.2 // indirect
)
