module goblog.com/service/book

go 1.17

replace (
	goblog.com/api => ../../api
	goblog.com/pkg => ../../pkg
)

require (
	github.com/golang-migrate/migrate/v4 v4.15.1
	github.com/google/wire v0.5.0
	github.com/spf13/viper v1.10.1
	go.uber.org/zap v1.21.0
	goblog.com/api v0.0.0-00010101000000-000000000000
	goblog.com/pkg v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
	gorm.io/gorm v1.23.4
)

require (
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/net v0.0.0-20220407224826-aac1ed45d8e3 // indirect
	golang.org/x/sys v0.0.0-20220406163625-3f8b81556e12 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220407144326-9054f6ed7bac // indirect
	gopkg.in/ini.v1 v1.66.4 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.3.3 // indirect
)
