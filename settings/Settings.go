package settings

import (
	"fmt"
	"os"
	"strconv"
	"time"
	constants "wellbe-common/settings/constants"
	commonconstants "wellbe-common/share/commonsettings/constants"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func GetDbSettings() (string, string) {
    dbDriver := os.Getenv(commonconstants.ENV_WELLBE_COMMON_DB_DRIVER)
    domain := os.Getenv(commonconstants.ENV_WELLBE_COMMON_DB_DOMAIN)
    port := os.Getenv(commonconstants.ENV_WELLBE_COMMON_DB_PORT)
    user := os.Getenv(commonconstants.ENV_WELLBE_COMMON_DB_USER)
    password := os.Getenv(commonconstants.ENV_WELLBE_COMMON_DB_PASSWORD)
    dbname := os.Getenv(commonconstants.ENV_WELLBE_COMMON_DB_DBNAME)
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", domain, port, user, password, dbname)

    return dbDriver, dsn
}


func GetApiSettings() (string) {
    apiPort := os.Getenv(commonconstants.ENV_WELLBE_COMMON_PORT)

    return apiPort
}


func GetMailerApiSettings() (string) {
    mailerDomain := os.Getenv(commonconstants.ENV_WELLBE_COMMON_MAILER_DOMAIN)

    return mailerDomain
}


func GetAWSRegionName() (string) {
    regionName := os.Getenv(commonconstants.ENV_WELLBE_COMMON_REGION_NAME)

    return regionName
}

func GetAWSCredentialProfilerNam() (string) {
    profilerName := os.Getenv(commonconstants.ENV_WELLBE_COMMON_CREDENTIAL_PROFILER_NAME)

    return profilerName
}

func GetAWSConfigure() (*aws.Config) {
    local := os.Getenv(commonconstants.ENV_WELLBE_AWS_CONFIGURE_LOCAL)
    if local == commonconstants.ENV_WELLBE_AWS_CONFIGURE_LOCAL_TRUE {
        return &aws.Config{
            Region:      aws.String(GetAWSRegionName()),
            Credentials: credentials.NewSharedCredentials("", GetAWSCredentialProfilerNam()),
        }
    } else {
        return aws.NewConfig().WithRegion(GetAWSRegionName())
    }
}

func GetImageDomainName() (string) {
    userPoolId := os.Getenv(commonconstants.ENV_WELLBE_COMMON_IMAGE_DOMAIN)

    return userPoolId
}

func GetImgixDomainName() (string) {
    imgixDomain := os.Getenv(commonconstants.ENV_WELLBE_COMMON_IMGIX_DOMAIN)

    return imgixDomain
}

func GetRecaptchSecretKey() (string) {
    key := os.Getenv(commonconstants.ENV_WELLBE_COMMON_RECAPTCHA_SECRET_KEY)

    return key
}

func GetExchangerateDomain() (string) {
    domain := os.Getenv(constants.ENV_EXCHANGERATE_DOMAIN)

    return domain
}

func GetExchangerateAccessKey() (string) {
    accesskey := os.Getenv(constants.ENV_EXCHANGERATE_ACCESS_KEY)

    return accesskey
}


func GetS3StaticBucketName() (string) {
    bucketName := os.Getenv(commonconstants.ENV_WELLBE_COMMON_S3_STATIC_BUCKET_NAME)

    return bucketName
}

func GetDbMaxOpenConections() (int) {
    openConnectionsStr := os.Getenv(commonconstants.ENV_WELLBE_COMMON_DB_MAX_OPEN_CONNECTIONS)
    openConnections, err := strconv.Atoi(openConnectionsStr)
    if err != nil {
        return 0
    }

    return openConnections
}

func GetDbMaxIdleConections() (int) {
    idleConnectionsStr := os.Getenv(commonconstants.ENV_WELLBE_COMMON_DB_MAX_IDLE_CONNECTIONS)
    idleConnections, err := strconv.Atoi(idleConnectionsStr)
    if err != nil {
        return 0
    }

    return idleConnections
}

func GetDbMaxLifeTimeMinutes() (time.Duration) {
    lifeTimeMinutesStr := os.Getenv(commonconstants.ENV_WELLBE_COMMON_DB_MAX_LIFETIME_MINUTES)
    lifeTimeMinutes, err := strconv.Atoi(lifeTimeMinutesStr)
    if err != nil {
        return 0
    }

    return time.Duration(lifeTimeMinutes)
}

func GetTranslateStab() (string) {
    key := os.Getenv(commonconstants.ENV_WELLBE_AWS_TRANSLATE_STAB)

    return key
}