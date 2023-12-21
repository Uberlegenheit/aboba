package main

import (
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
	"time"
)

type Article struct {
	ID                uint       `gorm:"primaryKey"`
	Title             string     `gorm:"type:json;not null"`
	Type              string     `gorm:"type:article_type_enum;default:0;not null"`
	Subject           string     `gorm:"type:varchar"`
	Body              string     `gorm:"type:json;not null"`
	Publish           bool       `gorm:"default:false;not null"`
	ConferenceURL     string     `gorm:"type:varchar"`
	GoogleCalendarURL string     `gorm:"type:varchar"`
	Description       string     `gorm:"type:json;not null"`
	Duration          string     `gorm:"type:json"`
	StartsAt          time.Time  `gorm:"type:timestamp"`
	CreatedAt         time.Time  `gorm:"type:timestamp;default:now();not null"`
	UpdatedAt         time.Time  `gorm:"type:timestamp;default:now();not null"`
	DeletedAt         *time.Time `gorm:"type:timestamp"`
	CourseID          int        `gorm:"constraint:FK_94317711db886a7c284f0423293;references:course"`
	Image             string     `gorm:"type:varchar"`
	VideoId           string     `gorm:"type:json"`
	UpdateLinks       string     `gorm:"type:json"`
	LocalizedImage    string     `gorm:"type:json"`
}

type Course struct {
	ID         uint           `gorm:"primaryKey;column:id"`
	Title      string         `gorm:"column:title;not null"`
	CreatedAt  time.Time      `gorm:"column:createdAt;default:now();not null"`
	UpdatedAt  time.Time      `gorm:"column:updatedAt;default:now();not null"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deletedAt"`
	ProductID  uint           `gorm:"column:productId;constraint:FK_a5dc2a0a8b60847ccbe401709ee;references:product"`
	Image      string         `gorm:"column:image"`
	Categories []string       `gorm:"column:categories;type:text[];default:'{}'::text[];not null"`
}

func makeConn() (*gorm.DB, error) {
	s := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), false)
	return gorm.Open(postgres.Open(s), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
}

func main() {
	db, err := makeConn()
	if err != nil {
		log.Fatal(err)
	}

	//articles := make([]Article, 0)
	//err = db.Table("article").
	//	Select("*").
	//	Scan(&articles).Error
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for i := range articles {
	//	articles[i].Body = strings.ReplaceAll(articles[i].Body, "cdn.cogitize.tech", "cryptomannn.s3.eu-central-1.amazonaws.com")
	//
	//	err = db.Table("article").
	//		Where("id = ?", articles[i].ID).
	//		Update("body", articles[i].Body).Error
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	courses := make([]Course, 0)
	err = db.Table("course").
		Select("*").
		Scan(&courses).Error
	if err != nil {
		log.Fatal(err)
	}

	for i := range courses {
		courses[i].Image = strings.ReplaceAll(courses[i].Image, "cdn.cogitize.tech", "api.cryptomannn.com")
		courses[i].Image = strings.ReplaceAll(courses[i].Image, "cryptomannn.s3.eu-central-1.amazonaws.com", "api.cryptomannn.com")

		err = db.Table("article").
			Where("id = ?", courses[i].ID).
			Update("image", courses[i].Image).Error
		if err != nil {
			log.Fatal(err)
		}
	}
}
