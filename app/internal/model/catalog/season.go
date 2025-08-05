package catalog

type Season struct {
	Id        int             `bson:"_id" json:"id"`
	Title     string          `bson:"title" json:"title" binding:"required" form:"fields[title]"`
	SerialId  int             `bson:"serialId" json:"serialId" binding:"required" form:"fields[serial_id]"`
	Sort      int             `bson:"sort" json:"sort" form:"fields[sort]"`
	Moderated bool            `bson:"moderated" json:"moderated" form:"fields[moderated]"`
	CreatedBy int             `bson:"createdBy" json:"created_by" form:"fields[created_by]"`
	episodes  map[int]Episode `bson:"-"`
}

func (s *Season) GetId() int {
	return s.Id
}

func (s *Season) SetId(id int) {
	s.Id = id
}

func NewSeason() *Season {
	return &Season{
		episodes: make(map[int]Episode),
	}
}

type SeasonOption func(*Season)

func WithEpisode(episode *Episode) SeasonOption {
	return func(opts *Season) {
		opts.episodes[episode.Id] = *episode
	}
}
