package config

import "time"

type SampleConfig struct {
}

func (s SampleConfig) GeSignedPeriod() time.Duration {
	return 15 * time.Minute
}

func (s SampleConfig) GetSpaceKey() string {
	return "DO00QNH8XPWKQ8LACPYH"
}

func (s SampleConfig) GetSpaceSecret() string {
	return "8fW6lccpP19IjpTRCxS7AENWuhiz6/768ftZ7KDfyws"
}

func (s SampleConfig) GetSpaceRegion() string {
	return "eu-central-1"
}

func (s SampleConfig) GetSpaceEndpoint() string {
	return "https://traderfour.fra1.digitaloceanspaces.com"
}

func (s SampleConfig) GetSpaceBucket() string {
	return "t4"
}

func (s SampleConfig) GetDatabaseDriver() string {
	return "postgres"
}

func (s SampleConfig) GetStorageType() string {
	return "file"
}

func (s SampleConfig) GetS3Bucket() string {
	return ""
}

func (s SampleConfig) GetDSN() string {
	return "host=localhost user=sajjad password=sajjad123 dbname=fs port=5432 sslmode=require TimeZone=UTC"
}

func mio() Config {
	return &SampleConfig{}
}
