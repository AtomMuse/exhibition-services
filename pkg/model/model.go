package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserID represents user identification data.
type UserID struct {
	UserID    int    `bson:"userId,omitempty" json:"userId,omitempty" validate:"required"`
	FirstName string `bson:"firstName,omitempty" json:"firstName,omitempty"`
	LastName  string `bson:"lastName,omitempty" json:"lastName,omitempty"`
}

// ExhibitionSection represents the structure of an exhibition section.
type ExhibitionSection struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	SectionType  string             `bson:"sectionType,omitempty" json:"sectionType,omitempty" validate:"required"`
	ContentType  string             `bson:"contentType,omitempty" json:"contentType,omitempty" `
	Background   string             `bson:"background,omitempty" json:"background,omitempty"`
	Title        string             `bson:"title,omitempty" json:"title,omitempty"`
	Text         string             `bson:"text,omitempty" json:"text,omitempty"`
	LeftCol      LeftColumn         `bson:"leftCol,omitempty" json:"leftCol,omitempty" `
	RightCol     RightColumn        `bson:"rightCol,omitempty" json:"rightCol,omitempty" `
	Images       []string           `bson:"images,omitempty" json:"images,omitempty" validate:"omitempty,url"`
	ExhibitionID primitive.ObjectID `bson:"exhibitionId" json:"exhibitionIs" validate:"required"`
}

// LeftColumn represents the structure of the left column in an exhibition section.
type LeftColumn struct {
	ContentType      string `bson:"contentType,omitempty" json:"contentType,omitempty" `
	Image            string `bson:"image,omitempty" json:"image,omitempty" validate:"omitempty,url"`
	ImageDescription string `bson:"imageDescription,omitempty" json:"imageDescription,omitempty"`
	Title            string `bson:"title,omitempty" json:"title,omitempty"`
	Text             string `bson:"text,omitempty" json:"text,omitempty"`
}

// RightColumn represents the structure of the right column in an exhibition section.
type RightColumn struct {
	ContentType      string `bson:"contentType,omitempty" json:"contentType,omitempty" `
	Image            string `bson:"image,omitempty" json:"image,omitempty" validate:"omitempty"`
	ImageDescription string `bson:"imageDescription,omitempty" json:"imageDescription,omitempty"`
	Title            string `bson:"title,omitempty" json:"title,omitempty"`
	Text             string `bson:"text,omitempty" json:"text,omitempty"`
}

type Contents struct {
	Title string     `bson:"title,omitempty" json:"title,omitempty"`
	Text  [][]string `bson:"text,omitempty" json:"text,omitempty"`
}

type Details struct {
	Img      string     `bson:"img,omitempty" json:"img,omitempty"`
	Contents []Contents `bson:"contents,omitempty" json:"contents,omitempty"`
}

type CenterItem struct {
	PreviewType string  `bson:"previewType,omitempty" json:"previewType,omitempty"`
	Src         string  `bson:"src,omitempty" json:"src,omitempty"`
	Details     Details `bson:"details,omitempty" json:"details,omitempty"`
}

type LeftRightItem struct {
	PreviewType string  `bson:"previewType,omitempty" json:"previewType,omitempty"`
	Src         string  `bson:"src,omitempty" json:"src,omitempty"`
	Details     Details `bson:"details,omitempty" json:"details,omitempty"`
}

type Room struct {
	MapThumbnail string          `bson:"mapThumbnail,omitempty" json:"mapThumbnail,omitempty"`
	Left         []LeftRightItem `bson:"left,omitempty" json:"left,omitempty"`
	Center       []CenterItem    `bson:"center,omitempty" json:"center,omitempty"`
	Right        []LeftRightItem `bson:"right,omitempty" json:"right,omitempty"`
}

// ResponseExhibition represents the structure of the exhibition data.
type ResponseExhibition struct {
	ID                    primitive.ObjectID  `bson:"_id,omitempty" json:"_id,omitempty" validate:"required"`
	ExhibitionName        string              `bson:"exhibitionName" json:"exhibitionName" validate:"required"`
	ExhibitionDescription string              `bson:"exhibitionDescription,omitempty" json:"exhibitionDescription,omitempty"`
	ThumbnailImg          string              `bson:"thumbnailImg,omitempty" json:"thumbnailImg,omitempty"`
	StartDate             string              `bson:"startDate" json:"startDate"`
	EndDate               string              `bson:"endDate" json:"endDate" validate:"gtfield=StartDate"`
	IsPublic              bool                `bson:"isPublic,omitempty" json:"isPublic,omitempty"`
	ExhibitionCategories  []string            `bson:"exhibitionCategories,omitempty" json:"exhibitionCategories,omitempty"`
	ExhibitionTags        []string            `bson:"exhibitionTags,omitempty" json:"exhibitionTags,omitempty"`
	UserID                UserID              `bson:"userId" json:"userId" validate:"required"`
	LayoutUsed            string              `bson:"layoutUsed,omitempty" json:"layoutUsed,omitempty" validate:"required"`
	ExhibitionSections    []ExhibitionSection `bson:"exhibitionSections,omitempty" json:"exhibitionSections,omitempty" `
	VisitedNumber         int                 `bson:"visitedNumber,omitempty" json:"visitedNumber"`
	Room                  []Room              `bson:"rooms,omitempty" json:"rooms,omitempty"`
	Status                string              `bson:"status" json:"status" validate:"required" error:"status is required"`
}

// RequestGetExhibition represents the structure of the request to get an exhibition.
type ResponseGetExhibitionId struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" validate:"required"`
}

// RequestGetExhibition represents the structure of the request to get an exhibition.
type RequestGetExhibition struct {
	ID primitive.ObjectID `json:"-" validate:"required,primitive_object"`
}

// RequestGetExhibitionSection represents the structure of the request to get an exhibition.
type ResponseGetExhibitionSectionId struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" validate:"required"`
}

// RequestCreateExhibition represents the structure of the request to create an exhibition.
type RequestCreateExhibition struct {
	ExhibitionName        string   `bson:"exhibitionName" json:"exhibitionName" validate:"required" error:"ExhibitionName is required"`
	ExhibitionDescription string   `bson:"exhibitionDescription" json:"exhibitionDescription" validate:"required" error:"ExhibitionDescription is required"`
	ThumbnailImg          string   `bson:"thumbnailImg,omitempty" json:"thumbnailImg,omitempty"`
	StartDate             string   `bson:"startDate" json:"startDate" validate:"required" error:"StartDate is required"`
	EndDate               string   `bson:"endDate" json:"endDate" validate:"required" error:"EndDate is required and must be greater than StartDate"`
	IsPublic              bool     `bson:"isPublic" json:"isPublic" validate:"required" error:"IsPublic is required"`
	ExhibitionCategories  []string `bson:"exhibitionCategories,omitempty" json:"exhibitionCategories,omitempty"`
	ExhibitionTags        []string `bson:"exhibitionTags,omitempty" json:"exhibitionTags,omitempty"`
	UserID                UserID   `bson:"userId" json:"userId" validate:"required" error:"UserID is required"`
	LayoutUsed            string   `bson:"layoutUsed,omitempty" json:"layoutUsed,omitempty" validate:"required" error:"LayoutUsed is required"`
	ExhibitionSectionsID  []string `bson:"exhibitionSectionsID,omitempty" json:"exhibitionSectionsID,omitempty"`
	VisitedNumber         int      `bson:"visitedNumber,omitempty" json:"visitedNumber,omitempty"`
	Room                  []Room   `bson:"rooms,omitempty" json:"rooms,omitempty"`
	Status                string   `bson:"status" json:"status" validate:"required" error:"status is required"`
}

type RequestUpdateExhibition struct {
	ExhibitionName        string   `bson:"exhibitionName,omitempty" json:"exhibitionName,omitempty"`
	ExhibitionDescription string   `bson:"exhibitionDescription,omitempty" json:"exhibitionDescription,omitempty"`
	ThumbnailImg          string   `bson:"thumbnailImg,omitempty" json:"thumbnailImg,omitempty"`
	StartDate             string   `bson:"startDate,omitempty" json:"startDate,omitempty"`
	EndDate               string   `bson:"endDate,omitempty" json:"endDate,omitempty"`
	IsPublic              bool     `bson:"isPublic,omitempty" json:"isPublic,omitempty"`
	ExhibitionCategories  []string `bson:"exhibitionCategories,omitempty" json:"exhibitionCategories,omitempty"`
	ExhibitionTags        []string `bson:"exhibitionTags,omitempty" json:"exhibitionTags,omitempty"`
	UserID                UserID   `bson:"userId,omitempty" json:"userId,omitempty"`
	LayoutUsed            string   `bson:"layoutUsed,omitempty" json:"layoutUsed,omitempty"`
	ExhibitionSectionsID  []string `bson:"exhibitionSectionsID,omitempty" json:"exhibitionSectionsID,omitempty"`
	VisitedNumber         int      `bson:"visitedNumber,omitempty" json:"visitedNumber,omitempty"`
	Room                  []Room   `bson:"rooms,omitempty" json:"rooms,omitempty"`
	Status                string   `bson:"status" json:"status" validate:"required" error:"status is required"`
}

type RequestCreateExhibitionSection struct {
	SectionType  string             `bson:"sectionType,omitempty" json:"sectionType,omitempty" validate:"required"`
	ContentType  string             `bson:"contentType,omitempty" json:"contentType,omitempty" `
	Background   string             `bson:"background,omitempty" json:"background,omitempty"`
	Title        string             `bson:"title,omitempty" json:"title,omitempty"`
	Text         string             `bson:"text,omitempty" json:"text,omitempty"`
	LeftCol      LeftColumn         `bson:"leftCol,omitempty" json:"leftCol,omitempty" `
	RightCol     RightColumn        `bson:"rightCol,omitempty" json:"rightCol,omitempty" `
	Images       []string           `bson:"images,omitempty" json:"images,omitempty" validate:"omitempty,url"`
	ExhibitionID primitive.ObjectID `bson:"exhibitionID" json:"exhibitionID" validate:"required"`
}

type ResponseExhibitionSection struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	SectionType  string             `bson:"sectionType,omitempty" json:"sectionType,omitempty" validate:"required"`
	ContentType  string             `bson:"contentType,omitempty" json:"contentType,omitempty"`
	Background   string             `bson:"background,omitempty" json:"background,omitempty"`
	Title        string             `bson:"title,omitempty" json:"title,omitempty" `
	Text         string             `bson:"text,omitempty" json:"text,omitempty"`
	LeftCol      LeftColumn         `bson:"leftCol,omitempty" json:"leftCol,omitempty" `
	RightCol     RightColumn        `bson:"rightCol,omitempty" json:"rightCol,omitempty" `
	Images       []string           `bson:"images,omitempty" json:"images,omitempty" validate:"omitempty,url"`
	ExhibitionID primitive.ObjectID `bson:"exhibitionID" json:"exhibitionID" validate:"required"`
}
