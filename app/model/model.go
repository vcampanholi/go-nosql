package model

import (
	"time"
)

type (
	Route struct {
		TravelID int       `bson:"travel_id,omitempty"`
		SiteID   string    `bson:"site_id,omitempty"`
		Status   string    `bson:"status,omitempty"`
		CreateAt time.Time `bson:"create_at,omitempty"`
	}
	Create struct {
		TotalRoutes int `json:"total_routes" validate:"required,numeric"`
	}
)
