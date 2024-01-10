package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Exhibition represents the structure of the exhibition data.
type Exhibition struct {
	ID                    primitive.ObjectID  `bson:"_id,omitempty" json:"_id,omitempty"`
	ExhibitionName        string              `bson:"exhibitionName,omitempty" json:"exhibitionName,omitempty"`
	ExhibitionDescription string              `bson:"exhibitionDescription,omitempty" json:"exhibitionDescription,omitempty"`
	ThumbnailImg          string              `bson:"thumbnailImg,omitempty" json:"thumbnailImg,omitempty"`
	StartDate             string              `bson:"startDate,omitempty" json:"startDate,omitempty"`
	EndDate               string              `bson:"endDate,omitempty" json:"endDate,omitempty"`
	IsPublic              bool                `bson:"isPublic,omitempty" json:"isPublic,omitempty"`
	ExhibitionCategories  []string            `bson:"exhibitionCategories,omitempty" json:"exhibitionCategories,omitempty"`
	ExhibitionTags        []string            `bson:"exhibitionTags,omitempty" json:"exhibitionTags,omitempty"`
	UserID                UserID              `bson:"userId,omitempty" json:"userId,omitempty"`
	LayoutUsed            string              `bson:"layoutUsed,omitempty" json:"layoutUsed,omitempty"`
	ExhibitionSections    []ExhibitionSection `bson:"exhibitionSections,omitempty" json:"exhibitionSections,omitempty"`
}

// UserID represents user identification data.
type UserID struct {
	UserID    int    `bson:"userId,omitempty" json:"userId,omitempty"`
	FirstName string `bson:"firstName,omitempty" json:"firstName,omitempty"`
	LastName  string `bson:"lastName,omitempty" json:"lastName,omitempty"`
}

// ExhibitionSection represents the structure of an exhibition section.
type ExhibitionSection struct {
	SectionType string      `bson:"sectionType,omitempty" json:"sectionType,omitempty"`
	ContentType string      `bson:"contentType,omitempty" json:"contentType,omitempty"`
	Background  string      `bson:"background,omitempty" json:"background,omitempty"`
	Title       string      `bson:"title,omitempty" json:"title,omitempty"`
	Text        string      `bson:"text,omitempty" json:"text,omitempty"`
	LeftCol     LeftColumn  `bson:"leftCol,omitempty" json:"leftCol,omitempty"`
	RightCol    RightColumn `bson:"rightCol,omitempty" json:"rightCol,omitempty"`
	Images      []string    `bson:"images,omitempty" json:"images,omitempty"`
}

// LeftColumn represents the structure of the left column in an exhibition section.
type LeftColumn struct {
	ContentType      string `bson:"contentType,omitempty" json:"contentType,omitempty"`
	Image            string `bson:"image,omitempty" json:"image,omitempty"`
	ImageDescription string `bson:"imageDescription,omitempty" json:"imageDescription,omitempty"`
}

// RightColumn represents the structure of the right column in an exhibition section.
type RightColumn struct {
	ContentType      string `bson:"contentType,omitempty" json:"contentType,omitempty"`
	Image            string `bson:"image,omitempty" json:"image,omitempty"`
	ImageDescription string `bson:"imageDescription,omitempty" json:"imageDescription,omitempty"`
	Title            string `bson:"title,omitempty" json:"title,omitempty"`
}
