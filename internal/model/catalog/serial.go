package catalog

type Serial struct {
	Id               int            `bson:"_id" json:"id" example:"1"`
	Title            string         `bson:"title" json:"title" binding:"required" form:"fields[title]" example:"Breaking Bad"`
	FileId           int            `bson:"fileId" json:"fileId" form:"fields[fileId]" example:"0"`
	Description      string         `bson:"description" json:"description" form:"fields[description]" example:"TV series"`
	Rating           float64        `bson:"rating" json:"rating" form:"fields[rating]" example:"9.5"`
	Duration         float64        `bson:"duration" json:"duration" form:"fields[duration]" example:"40"`
	Sort             int            `bson:"sort" json:"sort" form:"fields[sort]" example:"1"`
	ProductionPeriod string         `bson:"productionPeriod" json:"productionPeriod" form:"fields[productionPeriod]" example:"2008-2013"`
	Quality          string         `bson:"quality" json:"quality" form:"fields[quality]" example:"High"`
	seasons          map[int]Season `bson:"-"`
}

func (s *Serial) GetId() int {
	return s.Id
}

func (s *Serial) SetId(id int) {
	s.Id = id
}

type SerialOption func(*Serial)

func WithSeason(season *Season) SerialOption {
	return func(opts *Serial) {
		opts.seasons[season.Id] = *season
	}
}

func NewSerial() *Serial {
	return &Serial{
		seasons: make(map[int]Season),
	}
}
