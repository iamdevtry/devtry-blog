package common

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type Image struct {
	Id        *uuid.UUID `json:"id" gorm:"column:id;"`
	Url       string     `json:"url" gorm:"column:url;"`
	Width     int        `json:"width" gorm:"column:width;"`
	Height    int        `json:"height" gorm:"column:height;"`
	CloudName string     `json:"cloud_name,omitempty" gorm:"column:cloud_name;"`
	Extension string     `json:"extension,omitempty" gorm:"column:extension;"`
	FileName  string     `json:"file_name,omitempty" gorm:"column:file_name;"`
}

func (Image) TableName() string {
	return "images"
}

func (j *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}

	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

// Value returns json value, implement driver.Valuer interface
func (j *Image) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}

type Images []Image

func (j *Images) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}

	var img []Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

func (j *Images) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}
