package s3

type Config interface {
	GetSpaceKey() string
	GetSpaceSecret() string
	GetSpaceRegion() string
	GetSpaceEndpoint() string
	GetSpaceBucket() string
}
type AWS struct {
	SpaceKey      string
	SpaceSecret   string
	SpaceRegion   string
	SpaceEndpoint string
	SpaceBucket   string
}

func NewS3(cfg Config) Storage {
	return AWS{
		SpaceKey:      cfg.GetSpaceKey(),
		SpaceSecret:   cfg.GetSpaceSecret(),
		SpaceRegion:   cfg.GetSpaceRegion(),
		SpaceEndpoint: cfg.GetSpaceEndpoint(),
		SpaceBucket:   cfg.GetSpaceBucket(),
	}
}
