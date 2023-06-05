package config

import "time"

type Config interface {
	GetDatabaseDriver() string
	GetStorageType() string
	GetS3Bucket() string
	GetDSN() string
	GetSpaceKey() string
	GetSpaceSecret() string
	GetSpaceRegion() string
	GetSpaceEndpoint() string
	GetSpaceBucket() string
	GeSignedPeriod() time.Duration
}
