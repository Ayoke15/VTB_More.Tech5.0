import json
from scipy.stats import beta

def save_office_filter_data():
    # Сохранение в файл
    with open("office_filter_data.json", "w") as f:
        json.dump(filter_data, f, indent=4)

# Чтение исходных данных из файла
with open("office_filter_data.json", "r") as f:
    filter_data = json.load(f)

# # Чтение новых данных из файла
# with open("office_new_data.json", "r") as f:
#     new_data = json.load(f)

def update_office_filter_data(filter_name, success):
    data = filter_data.get(filter_name)
    if data is not None:
        data['total_count'] += 1
        if success:
            data['success_count'] += 1
        data['alpha'] = data['success_count'] + 1
        data['beta'] = data['total_count'] - data['success_count'] + 1

def calculate_office_weights():
    weights = {}
    for filter_name, data in filter_data.items():
        alpha, beta = data['alpha'], data['beta']
        estimated_prob = alpha / (alpha + beta)
        weights[filter_name] = estimated_prob
    return weights

# # Обновляем данные на основе новых событий
# for sub_array in new_data:
#     for event in sub_array:
#         update_office_filter_data(event['filter'], event['success'])
#
# # Вычисляем новые веса для фильтров
# new_weights = calculate_office_weights()
#
# print("New office weights:", new_weights)
#
# # Сохраняем обновленные данные в файл
# save_office_filter_data()
