package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"securechainlog-api/models"
)

func GetAssetLogs(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, asset_id, status, timestamp, anomaly_detected FROM asset_logs")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var logs []models.AssetLog
		for rows.Next() {
			var log models.AssetLog
			if err := rows.Scan(&log.ID, &log.AssetID, &log.Status, &log.Timestamp, &log.AnomalyDetected); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			logs = append(logs, log)
		}
		json.NewEncoder(w).Encode(logs)
	}
}

func CreateAssetLog(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var log models.AssetLog
		if err := json.NewDecoder(r.Body).Decode(&log); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Timestamp = time.Now()

		err := db.QueryRow(`
			INSERT INTO asset_logs (asset_id, status, timestamp, anomaly_detected)
			VALUES ($1, $2, $3, $4)
			RETURNING id
		`, log.AssetID, log.Status, log.Timestamp, log.AnomalyDetected).Scan(&log.ID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(log)
	}
}
