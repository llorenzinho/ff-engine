package models

type FeatureFlag struct {
	Name      string `json:"name" binding:"required"`
	Namespace string `json:"namespace"`
	Enabled   *bool  `json:"enabled" binding:"required"`
	Active    *bool  `json:"active" binding:"required"`
}
