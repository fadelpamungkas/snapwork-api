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
		Id          primitive.ObjectID `json:"_id,omitempty"`
		Avatar      Image              `json:"avatar,omitempty"`
		KTP         Image              `json:"ktp,omitempty"`
		Ijazah      Image              `json:"ijazah,omitempty"`
		SKCK        Image              `json:"skck,omitempty"`
		CV          Image              `json:"cv,omitempty"`
		Certificate Image              `json:"certificate,omitempty"`
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
	Career struct {
		Id            primitive.ObjectID `json:"_id,omitempty"`
		ApplicationId primitive.ObjectID `json:"applicationId,omitempty"`
		Status        string             `json:"name,omitempty"`
		CompanyName   string             `json:"concentration,omitempty"`
		Date          string             `json:"date,omitempty"`
		Position      string             `json:"position,omitempty"`
		Type          string             `json:"type,omitempty"`
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
