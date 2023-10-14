import json
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

# Функция для проверки соответствия офиса выбранным фильтрам
def check_filters(office, user_filters):
    if 'rko' in user_filters and office.get('rko', None) != user_filters['rko']:
        return False
    if 'current_workload' in user_filters and office.get('current_workload') > user_filters['current_workload']:
        return False
    if 'hasRamp' in user_filters and office.get('hasRamp') != user_filters['hasRamp']:
        return False
    if 'officeType' in user_filters and office.get('officeType') != user_filters['officeType']:
        return False
    if 'salePointFormat' in user_filters and office.get('salePointFormat') != user_filters['salePointFormat']:
        return False
    if 'suoAvailability' in user_filters and office.get('suoAvailability') != user_filters['suoAvailability']:
        return False
    if 'kep' in user_filters and office.get('kep') != user_filters['kep']:
        return False
    if 'myBranch' in user_filters and office.get('myBranch') != user_filters['myBranch']:
        return False
    return True

# Функция для вычисления оценки офиса
def evaluate_office(office, weights, distance):
    score = 0

    # Расчет оценки расстояния
    distance_score = 1 / (1 + distance)
    score += weights['distance'] * distance_score

    # Расчет оценки наличия РКО
    rko_score = 1 if office.get('rko', None) == "есть РКО" else 0
    score += weights.get('rko', 0) * rko_score

    # Расчет оценки текущей нагрузки
    workload_score = 1 - (office['current_workload'] / 100)
    score += weights['current_workload'] * workload_score

    # Расчет оценки наличия рампы
    has_ramp_score = 1 if office['hasRamp'] == "Y" else 0
    score += weights['has_ramp'] * has_ramp_score

    # Расчет оценки типа офиса
    office_type_score = 1 if office['officeType'] == "Да (Зона Привилегия)" else 0
    score += weights['office_type'] * office_type_score

    # Расчет оценки для salePointFormat
    sale_point_format_score = 1 if office['salePointFormat'] == "Универсальный" else 0
    score += weights['sale_point_format'] * sale_point_format_score

    # Расчет оценки для SUO availability
    suo_availability_score = 1 if office['suoAvailability'] == "Y" else 0
    score += weights['suo_availability'] * suo_availability_score

    # Расчет оценки для KEP
    kep_score = 1 if office['kep'] else 0
    score += weights['kep'] * kep_score

    # Расчет оценки для MyBranch
    my_branch_score = 1 if office['myBranch'] else 0
    score += weights['my_branch'] * my_branch_score

    return score

# Веса для каждого критерия
office_weights = {
    'distance': 0.4,
    'rko': 0.1,
    'current_workload': 0.1,
    'has_ramp': 0.1,
    'office_type': 0.1,
    'sale_point_format': 0.1,
    'suo_availability': 0.1,
    'kep': 0.1,
    'my_branch': 0.1
}

# user_location = (55.755826, 37.6172999)
#
# # Пример данных о фильтрах (можно заменить на реальные данные)
# user_filters = {
#     'rko': 'есть РКО',
#     'current_workload': 70,
#     'hasRamp': 'Y',
#     'officeType': 'Да (Зона Привилегия)',
#     'salePointFormat': 'Универсальный',
#     'suoAvailability': 'Y',
#     'kep': True,
#     'myBranch': False
# }
#
#
# # Загрузить данные офисов из JSON файла
# with open('mergedOffices.json', 'r',encoding='utf-8') as f:
#     offices = json.load(f)
#
# filtered_and_scored_offices = []
# for office in offices:
#     if check_filters(office, user_filters):
#         # добавить расчет расстояния
#         office_location = (office['latitude'], office['longitude'])
#         distance = haversine_distance(user_location, office_location)
#
#         score = evaluate_office(office, weights, distance)
#         filtered_and_scored_offices.append((office, score))
#
# # Сортировка офисов по оценке в убывающем порядке
# filtered_and_scored_offices.sort(key=lambda x: x[1], reverse=True)
#
# # Вывод первых 5 офисов
# top_5_offices = filtered_and_scored_offices[:5]
# for i, (office, score) in enumerate(top_5_offices):
#    print(f"{i+1}. Офис ID: {office['offices_id']}, Оценка: {score}")