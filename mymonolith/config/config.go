package config

import (
	"errors"
	"os"
	"sync"
	"time"

	"github.com/spf13/cast"
)

var (
	//ErrExpiredPassword error text
	ErrExpiredPassword error = errors.New("expired_password")
)

var (
	instance *Configuration
	once     sync.Once
)

//Config ...
func Config() *Configuration {
	once.Do(func() {
		instance = load()
	})

	return instance
}

// Configuration ...
type Configuration struct {
	AppURL                     string
	BucketName                 string
	StorageType                string
	AwsS3Id                    string
	AwsS3Secret                string
	AwsBucketURL               string
	Environment                string
	PostgresHost               string
	PostgresPort               int
	PostgresDatabase           string
	PostgresUser               string
	PostgresPassword           string
	ServerPort                 int
	ServerHost                 string
	LogLevel                   string
	ServiceDir                 string
	AccessTokenDuration        time.Duration
	RefreshTokenDuration       time.Duration
	RefreshPasswdTokenDuration time.Duration

	RedisHost string
	RedisPort int

	CasbinConfigPath    string
	MiddleWareRolesPath string

	// context timeout in seconds
	CtxTimeout        int
	SigninKey         string
	ServerReadTimeout int

	FirebaseWebKey     string
	DomainURIPrefix    string
	AndroidPackageName string
	IosBundleID        string

	JWTSecretKey              string
	JWTSecretKeyExpireMinutes int
	JWTRefreshKey             string
	JWTRefreshKeyExpireHours  int
}

func load() *Configuration {
	var bucketName string
	if os.Getenv("ENVIRONMENT") == "production" {
		bucketName = cast.ToString(getOrReturnDefault("BUCKET_NAME", "createhq-storage-prod"))
	} else {
		bucketName = cast.ToString(getOrReturnDefault("BUCKET_NAME", "createhq-storage"))
	}
	return &Configuration{
		BucketName:          bucketName,
		StorageType:         cast.ToString(getOrReturnDefault("STORAGE_TYPE", "s3")),
		AwsS3Id:             cast.ToString(getOrReturnDefault("AWS_S3_ID", "")),
		AwsS3Secret:         cast.ToString(getOrReturnDefault("AWS_S3_SECRET", "")),
		AwsBucketURL:        cast.ToString(getOrReturnDefault("AWS_BUCKET_URL", "https://"+bucketName+".s3.amazonaws.com/")),
		AppURL:              cast.ToString(getOrReturnDefault("APP_URL", "localhost:8000")),
		ServerHost:          cast.ToString(getOrReturnDefault("SERVER_HOST", "localhost")),
		ServerPort:          cast.ToInt(getOrReturnDefault("SERVER_PORT", "8000")),
		Environment:         cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop")),
		PostgresHost:        cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost")),
		PostgresPort:        cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432)),
		PostgresDatabase:    cast.ToString(getOrReturnDefault("POSTGRES_DB", "mymonolith")),
		PostgresUser:        cast.ToString(getOrReturnDefault("POSTGRES_USER", "sardor")),
		PostgresPassword:    cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "sardor")),
		LogLevel:            cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug")),
		ServiceDir:          cast.ToString(getOrReturnDefault("CURRENT_DIR", "")),
		CtxTimeout:          cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7)),
		RedisHost:           cast.ToString(getOrReturnDefault("REDIS_HOST", "localhost")),
		RedisPort:           cast.ToInt(getOrReturnDefault("REDIS_PORT", 0)),
		CasbinConfigPath:    cast.ToString(getOrReturnDefault("CASBIN_CONFIG_PATH", "./config/rbac_model.conf")),
		MiddleWareRolesPath: cast.ToString(getOrReturnDefault("MIDLEWARE_ROLES_PATH", "./config/models.csv")),
		SigninKey:           cast.ToString(getOrReturnDefault("SIGNIN_KEY", "")),
		ServerReadTimeout:   cast.ToInt(getOrReturnDefault("SERVER_READ_TIMEOUT", "")),

		JWTSecretKey:              cast.ToString(getOrReturnDefault("JWT_SECRET_KEY", "")),
		JWTSecretKeyExpireMinutes: cast.ToInt(getOrReturnDefault("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT", 720)),
		JWTRefreshKey:             cast.ToString(getOrReturnDefault("JWT_REFRESH_KEY", "")),
		JWTRefreshKeyExpireHours:  cast.ToInt(getOrReturnDefault("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT", 24)),

		FirebaseWebKey: cast.ToString(getOrReturnDefault("FIREBASE_WEB_KEY", "")),
	}
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
