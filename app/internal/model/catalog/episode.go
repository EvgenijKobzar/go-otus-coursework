package catalog

type Episode struct {
	Id               int     `bson:"_id" json:"id"`
	Title            string  `bson:"title" json:"title" binding:"required" form:"fields[title]"`
	FileId           int     `bson:"fileId" json:"fileId" form:"fields[fileId]"`
	SeasonId         int     `bson:"seasonId" json:"seasonId" binding:"required" form:"fields[seasonId]"`
	SerialId         int     `bson:"serialId" json:"serialId" binding:"required" form:"fields[serialId]"`
	Description      string  `bson:"description" json:"description" form:"fields[description]"`
	Duration         float64 `bson:"duration" json:"duration" form:"fields[duration]"`
	Sort             int     `bson:"sort" json:"sort" form:"fields[sort]"`
	Rating           float64 `bson:"rating" json:"rating" form:"fields[rating]"`
	ProductionPeriod string  `bson:"productionDate" json:"productionDate" form:"fields[productionDate]"`
	Quality          string  `bson:"quality" json:"quality" form:"fields[quality]"`
	Moderated        bool    `bson:"moderated" json:"moderated" form:"fields[moderated]"`
	CreatedBy        int     `bson:"createdBy" json:"created_by" form:"fields[created_by]"`
}

func NewEpisode() *Episode {
	return &Episode{}
}

func (e *Episode) GetId() int {
	return e.Id
}

func (e *Episode) SetId(id int) {
	e.Id = id
}
