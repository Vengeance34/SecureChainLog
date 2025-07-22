import pandas as pd
import numpy as np
from sklearn.ensemble import RandomForestClassifier
import joblib

def train_model():
    # Sample training data
    data = {
        'asset_id': [1, 2, 1, 2],
        'status': [0, 1, 2, 1],  # Encoded status
        'anomaly': [0, 1, 0, 1]
    }

    df = pd.DataFrame(data)
    X = df[['asset_id', 'status']]
    y = df['anomaly']

    model = RandomForestClassifier()
    model.fit(X, y)

    joblib.dump(model, 'model.joblib')
    print("✅ Model trained and saved.")

def load_model():
    return joblib.load('model.joblib')
