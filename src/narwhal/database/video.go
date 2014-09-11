package database

import "log"
import "time"

type Video struct {
	Id int64
	Url string `sql:"not null"`
	Position string `sql:"not null"`
	CreatedAt int64
}

func CreateVideo(pos string, url string) (*Video) {
	video := Video{ Url: url, Position: pos, CreatedAt: time.Now().Unix() }
	err := DbMap.Save(&video)

	if err != nil {
		log.Print("Error: ")
		log.Println(err)
	}

	return &video
}

func GetAllVideos() (*[]Video) {
	var videos []Video
	DbMap.Select("*").Find(&videos)
	return &videos
}
