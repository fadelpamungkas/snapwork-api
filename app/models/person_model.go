package models

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	PersonEntity struct {
		Id              primitive.ObjectID      `json:"_id,omitempty"`
		UserId          primitive.ObjectID      `json:"userid,omitempty"`
		Name            string                  `json:"name,omitempty"`
		Birth           string                  `json:"birth,omitempty"`
		Gender          string                  `json:"gender,omitempty"`
		Religion        string                  `json:"religion,omitempty"`
		Marriage        string                  `json:"marriage,omitempty"`
		Hobby           string                  `json:"hobby,omitempty"`
		Telephone       string                  `json:"telephone,omitempty"`
		Email           string                  `json:"email,omitempty"`
		Twitter         string                  `json:"twitter,omitempty"`
		Linkedin        string                  `json:"linkedin,omitempty"`
		Address         string                  `json:"address,omitempty"`
		City            string                  `json:"city,omitempty"`
		Province        string                  `json:"province,omitempty"`
		State           string                  `json:"state,omitempty"`
		About           string                  `json:"about,omitempty"`
		Portfolio       Portfolio               `json:"portfolio,omitempty"`
		Document        Document                `json:"document,omitempty"`
		SelfDevelopment SelfDevelopment         `json:"selfdevelopment,omitempty"`
		Education       Education               `json:"education,omitempty"`
		Notification    []Notification          `json:"notification,omitempty"`
		Applications    []ApplicationUserEntity `json:"applications,omitempty"`
		CreatedAt       string                  `json:"created_at,omitempty"`
		UpdatedAt       string                  `json:"updated_at,omitempty"`
	}
	PersonRequest struct {
		Id        primitive.ObjectID `json:"_id,omitempty"`
		UserId    primitive.ObjectID `json:"userid,omitempty"`
		Name      string             `json:"name,omitempty"`
		Birth     string             `json:"birth,omitempty"`
		Gender    string             `json:"gender,omitempty"`
		Religion  string             `json:"religion,omitempty"`
		Marriage  string             `json:"marriage,omitempty"`
		Hobby     string             `json:"hobby,omitempty"`
		Telephone string             `json:"telephone,omitempty"`
		Email     string             `json:"email,omitempty"`
		Twitter   string             `json:"twitter,omitempty"`
		Linkedin  string             `json:"linkedin,omitempty"`
		Address   string             `json:"address,omitempty"`
		City      string             `json:"city,omitempty"`
		Province  string             `json:"province,omitempty"`
		State     string             `json:"state,omitempty"`
		About     string             `json:"about,omitempty"`
		Education Education          `json:"education,omitempty"`
		UpdatedAt string             `json:"updated_at,omitempty"`
	}
	PersonDocumentRequest struct {
		Id              primitive.ObjectID `json:"_id,omitempty"`
		Avatar          *multipart.Form    `json:"avatar,omitempty"`
		AvatarName      string             `json:"avatarname,omitempty"`
		KTP             *multipart.Form    `json:"ktp,omitempty"`
		KTPName         string             `json:"ktpname,omitempty"`
		Ijazah          *multipart.Form    `json:"ijazah,omitempty"`
		IjazahName      string             `json:"ijazahname,omitempty"`
		SKCK            *multipart.Form    `json:"skck,omitempty"`
		SKCKName        string             `json:"skckname,omitempty"`
		CV              *multipart.Form    `json:"cv,omitempty"`
		CVName          string             `json:"cvname,omitempty"`
		Certificate     *multipart.Form    `json:"certificate,omitempty"`
		CertificateName string             `json:"certificatename,omitempty"`
	}
	PersonResponse struct {
		Status  int        `json:"status"`
		Message string     `json:"message"`
		Data    *fiber.Map `json:"data"`
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
		Id         primitive.ObjectID `json:"_id,omitempty"`
		S1         string             `json:"s1,omitempty"`
		S1Major    string             `json:"s1major,omitempty"`
		S1DateIn   string             `json:"s1datein,omitempty"`
		S1DateOut  string             `json:"s1dateout,omitempty"`
		SMA        string             `json:"sma,omitempty"`
		SMAMajor   string             `json:"smamajor,omitempty"`
		SMADateIn  string             `json:"smadatein,omitempty"`
		SMADateOut string             `json:"smadateout,omitempty"`
	}
	SelfDevelopment struct {
		Id      primitive.ObjectID           `json:"_id,omitempty"`
		Score   string                       `json:"score,omitempty"`
		File    string                       `json:"file,omitempty"`
		Payment SelfDevelopmentPaymentEntity `json:"payment,omitempty"`
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
	SelfDevelopmentPaymentEntity struct {
		Status    string `json:"status,omitempty"`
		Price     int    `json:"price,omitempty"`
		FileProof string `json:"fileproof,omitempty"`
		CreatedAt string `json:"created_at,omitempty"`
	}
	SelfDevelopmentPaymentRequest struct {
		Id        primitive.ObjectID `json:"_id,omitempty"`
		Status    string             `json:"status,omitempty"`
		Price     int                `json:"price,omitempty"`
		FileProof string             `json:"fileproof,omitempty"`
	}
)
