package config

import "time"

type SampleConfig struct {
}

func (s SampleConfig) GeSignedPeriod() time.Duration {
	return 15 * time.Minute
}

func (s SampleConfig) GetSpaceKey() string {
	return "SampleSpaceKey"
}

func (s SampleConfig) GetSpaceSecret() string {
	return "SampleSpaceSecrett"
}

func (s SampleConfig) GetSpaceRegion() string {
	return "SampleSpace"
}

func (s SampleConfig) GetSpaceEndpoint() string {
	return "SampleSpaceEndpoint"
}

func (s SampleConfig) GetSpaceBucket() string {
	return "SampleSpaceBucket"
}

func (s SampleConfig) GetDatabaseDriver() string {
	return "SampleDatabaseDriver"
}

func (s SampleConfig) GetStorageType() string {
	return "SampleStorageType"
}

func (s SampleConfig) GetS3Bucket() string {
	return "SampleS3Bucket"
}

func (s SampleConfig) GetDSN() string {
	return "SampleDSN"
}

func mio() Config {
	return &SampleConfig{}
}
