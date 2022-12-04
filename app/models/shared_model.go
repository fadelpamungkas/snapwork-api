package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Image struct {
		Id   primitive.ObjectID `json:"_id,omitempty"`
		Name string             `json:"name,omitempty"`
		Url  string             `json:"url,omitempty"`
	}
	Document struct {
		Id              primitive.ObjectID `json:"_id,omitempty"`
		Avatar          Image              `json:"avatar,omitempty"`
		AvatarName      string             `json:"avatarname,omitempty"`
		KTP             Image              `json:"ktp,omitempty"`
		KTPName         string             `json:"ktpname,omitempty"`
		Ijazah          Image              `json:"ijazah,omitempty"`
		IjazahName      string             `json:"ijazahname,omitempty"`
		SKCK            Image              `json:"skck,omitempty"`
		SKCKName        string             `json:"skckname,omitempty"`
		CV              Image              `json:"cv,omitempty"`
		CVName          string             `json:"cvname,omitempty"`
		Certificate     Image              `json:"certificate,omitempty"`
		CertificateName string             `json:"certificatename,omitempty"`
	}
	Portfolio struct {
		Id          primitive.ObjectID `json:"_id,omitempty"`
		Link        string             `json:"link,omitempty"`
		Description string             `json:"description,omitempty"`
	}
	Education struct {
		Id       primitive.ObjectID `json:"_id,omitempty"`
		S1       string             `json:"s1,omitempty"`
		S1Major  string             `json:"s1major,omitempty"`
		S1Date   string             `json:"s1date,omitempty"`
		SMA      string             `json:"sma,omitempty"`
		SMAMajor string             `json:"smamajor,omitempty"`
		SMADate  string             `json:"smadate,omitempty"`
	}
	SelfDevelopment struct {
		Id     primitive.ObjectID `json:"_id,omitempty"`
		Score  string             `json:"score,omitempty"`
		Status string             `json:"status,omitempty"`
		File   string             `json:"file,omitempty"`
	}
	Notification struct {
		Id          primitive.ObjectID `json:"_id,omitempty"`
		UserId      primitive.ObjectID `json:"userid,omitempty"`
		Status      string             `json:"status,omitempty"`
		Title       string             `json:"title,omitempty"`
		Description string             `json:"description,omitempty"`
		IsRead      bool               `json:"isread,omitempty"`
		CreatedAt   string             `json:"created_at,omitempty"`
	}
)
