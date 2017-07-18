package model
//这里面model用于映射到页面中

type Userinfo struct {
	Nickname string
	User_id int8
	Sex int8
	Avatar string
	Level int8	`xorm:"-"`
	Birthday int
	Hometown string
	Signature string
	Recommender int		//推荐人id
	Graticule_lat string
	Graticule_lng string
	Version int
}

type Userlevel struct {
	User_id int
	Level int
	Empirical int
}
