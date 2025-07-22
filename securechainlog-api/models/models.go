package models

import "time"

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
}

type LoginActivity struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	LoginTime time.Time `json:"login_time"`
	IPAddress string    `json:"ip_address"`
}

type Asset struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	CreatedAt   time.Time `json:"created_at"`
}

type AssetLog struct {
	ID              int       `json:"id"`
	AssetID         int       `json:"asset_id"`
	Status          string    `json:"status"`
	Timestamp       time.Time `json:"timestamp"`
	AnomalyDetected bool      `json:"anomaly_detected"`
}
