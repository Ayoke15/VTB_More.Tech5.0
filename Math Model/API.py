from flask import Flask, request, jsonify, abort
import json
from MathModelOffice import haversine_distance, check_filters, evaluate_office, office_weights
from MathModelATM import haversine_distance, check_filters, evaluate_atm, atm_weights

app = Flask(__name__)


@app.route('/get_offices', methods=['POST'])
def get_filtered_offices():
    data = request.json
    user_location = tuple(data['user_location'])
    user_filters = data['user_filters']

    with open('mergedOffices.json', 'r', encoding='utf-8') as f:
        offices = json.load(f)

    filtered_and_scored_offices = []
    for office in offices:
        if check_filters(office, user_filters):
            office_location = (office['latitude'], office['longitude'])
            distance = haversine_distance(user_location, office_location)
            score = evaluate_office(office, office_weights, distance)
            print(office_weights)
            print(office)

            filtered_and_scored_offices.append({'id': office['offices_id'], 'score': score})

    if not filtered_and_scored_offices:
        abort(400, 'Bad request: No offices found.')

    filtered_and_scored_offices.sort(key=lambda x: x['score'], reverse=True)
    top_5_offices = filtered_and_scored_offices[:5]

    return jsonify(top_5_offices)

@app.route('/get_atms', methods=['POST'])
def get_filtered_atms():
    data = request.json
    user_location = tuple(data['user_location'])
    user_filters = data['user_filters']

    with open('mergedAtms.json', 'r', encoding='utf-8') as f:
        loaded_atms_data = json.load(f)

    atms = loaded_atms_data['atms']
    atm_scores = []
    for atm in atms:
        score = evaluate_atm(atm, atm_weights, user_location, user_filters)
        if score > 0:
            atm_scores.append({'id': atm['id_atms'], 'score': score})

    if not atm_scores:
        abort(400, 'Bad request: No ATMs found.')

    atm_scores.sort(key=lambda x: x['score'], reverse=True)
    best_atms = atm_scores[:5]

    return jsonify(best_atms)


if __name__ == '__main__':
    app.run(debug=True, port=8080)