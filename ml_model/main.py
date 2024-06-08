import io
import os
import json
import sys
import keras
import tensorflow as tf

import pandas as pd
from sklearn.preprocessing import StandardScaler

# Suppress TensorFlow logging
os.environ['TF_CPP_MIN_LOG_LEVEL'] = '3'  # Suppresses most TensorFlow logs
tf.get_logger().setLevel('ERROR')  # Suppresses TensorFlow warnings


def preprocess_data(data_frame):
    scaler = StandardScaler()
    df_cleaned = data_frame.apply(pd.to_numeric, errors='coerce').fillna(0)
    processed_data = scaler.fit_transform(df_cleaned)
    return processed_data


def get_first_prediction(prediction):
    first_prediction = prediction.flatten()[0]
    label = first_prediction * 100
    return float("{:.2f}".format(label))


def main():
    data = json.loads(sys.argv[1])

    data_frame = pd.DataFrame(data)

    processed_data = preprocess_data(data_frame)

    model_path = './ml_model/my_model.h5'
    model = keras.models.load_model(model_path)

    prediction = model.predict(processed_data)

    first_prediction_text = get_first_prediction(prediction)
    print(json.dumps({"Prediction": first_prediction_text}, indent=4))


if __name__ == '__main__':
    sys.stdout = io.TextIOWrapper(sys.stdout.detach(), encoding='utf-8')
    sys.stderr = io.TextIOWrapper(sys.stderr.detach(), encoding='utf-8')
    main()
