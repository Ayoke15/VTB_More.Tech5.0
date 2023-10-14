from flask import Flask, request, jsonify
import json
import MLforATM
import MLforOffice

app = Flask(__name__)


@app.route('/update', methods=['POST'])
def update():
    events = request.json
    if not isinstance(events, list):
        return jsonify({"status": "error", "message": "Expected a list of events"}), 400

    for event in events:
        filter_name = event.get('filter')
        success = event.get('success')

        if filter_name is None or success is None:
            return jsonify({"status": "error", "message": "Missing parameters in one of the events"}), 400

        MLforATM.update_filter_data(filter_name, success)

    MLforATM.save_filter_data()
    return jsonify({"status": "success"}), 200


@app.route('/weights', methods=['GET'])
def get_weights():
    with open("filter_data.json", "r") as f:
        filter_data = json.load(f)

    weights = MLforATM.calculate_weights(filter_data)
    return jsonify(weights), 200


@app.route('/update_office', methods=['POST'])
def update_office():
    events = request.json
    if not isinstance(events, list):
        return jsonify({"status": "error", "message": "Expected a list of events"}), 400

    for event in events:
        filter_name = event.get('filter')
        success = event.get('success')

        if filter_name is None or success is None:
            return jsonify({"status": "error", "message": "Missing parameters in one of the events"}), 400

        MLforOffice.update_office_filter_data(filter_name, success)

    MLforOffice.save_office_filter_data()
    return jsonify({"status": "success"}), 200


@app.route('/office_weights', methods=['GET'])
def get_office_weights():
    weights = MLforOffice.calculate_office_weights()
    return jsonify(weights), 200

if __name__ == "__main__":
    app.run(debug=True, port=8080)
