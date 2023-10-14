import math

def haversine_distance(coord1, coord2):
    lat1, lon1 = coord1
    lat2, lon2 = coord2
    R = 6371  # радиус Земли в километрах

    dlat = math.radians(lat2 - lat1)
    dlon = math.radians(lon2 - lon1)
    a = math.sin(dlat / 2) ** 2 + math.cos(math.radians(lat1)) * math.cos(math.radians(lat2)) * math.sin(dlon / 2) ** 2
    c = 2 * math.atan2(math.sqrt(a), math.sqrt(1 - a))
    distance = R * c

    return distance

def check_filters(atm, user_filters):
    if user_filters.get('allDay') is not None and atm['allDay'] != user_filters['allDay']:
        return False

    for service, service_info in user_filters.get('services', {}).items():
        atm_service_info = atm['services'].get(service, {})
        for key, value in service_info.items():
            if atm_service_info.get(key) != value:
                return False

    if user_filters.get('cash') is not None and atm['cash'] < user_filters['cash']:
        return False

    return True

def evaluate_atm(atm, weights, user_location, user_filters):
    if not check_filters(atm, user_filters):
        return 0

    score = 0
    distance = haversine_distance(user_location, (atm['latitude'], atm['longitude']))
    distance_score = 1 / (1 + distance)
    score += weights['distance'] * distance_score

    all_day_score = 1 if atm['allDay'] else 0
    score += weights['allDay'] * all_day_score

    services_score = 0
    for service, weight in weights['services'].items():
        if atm['services'][service]['serviceCapability'] == 'SUPPORTED' and atm['services'][service]['serviceActivity'] == 'AVAILABLE':
            services_score += weight
    score += weights['service_total'] * services_score

    cash_score = atm['cash'] / 1000000
    score += weights['cash'] * cash_score

    return score

# Загрузка данных о банкоматах
# with open('mergedAtms.json', 'r', encoding='utf-8') as f:
#     loaded_atms_data = json.load(f)
#
# atms = loaded_atms_data['atms']


# Веса для каждого критерия
atm_weights = {
    'distance': 0.4,
    'allDay': 0.1,
    'service_total': 0.3,
    'services': {
        'wheelchair': 0.05,
        'blind': 0.05,
        'nfcForBankCards': 0.05,
        'qrRead': 0.05,
        'supportsUsd': 0.05,
        'supportsChargeRub': 0.05,
        'supportsEur': 0.05,
        'supportsRub': 0.05,
    },
    'cash': 0.2
}
#add here data fro mergedAtmes.json


# Местоположение пользователя (можно получить через веб-сервис)
# user_location = (55.7558, 37.6173)
# user_filters = {'allDay': True, 'wheelchair': True}  # Пример фильтров
#
# # Вычисление оценки для каждого банкомата
# atm_scores = []
# for atm in atms:
#     score = evaluate_atm(atm, weights, user_location, user_filters)
#     atm_scores.append((atm['id_atms'], score))
#
# # Сортировка и вывод результатов
# atm_scores.sort(key=lambda x: x[1], reverse=True)
# best_atms = [atm_id for atm_id, _ in atm_scores[:5]]
# print("Лучшие банкоматы:", best_atms)