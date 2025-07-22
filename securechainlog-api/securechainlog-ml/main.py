from fastapi import FastAPI
from model import load_model
from schema import AssetLogInput
import numpy as np

app = FastAPI()
model = load_model()

# Simulate encoding
status_map = {
    "Operational": 0,
    "Maintenance Required": 1,
    "Delivered": 2,
    "Damaged": 3
}

@app.post("/predict")
def predict_anomaly(log: AssetLogInput):
    status_encoded = status_map.get(log.status, -1)
    if status_encoded == -1:
        return {"error": "Invalid status"}

    features = np.array([[log.asset_id, status_encoded]])
    prediction = model.predict(features)

    return {
        "asset_id": log.asset_id,
        "status": log.status,
        "anomaly_detected": bool(prediction[0])
    }
