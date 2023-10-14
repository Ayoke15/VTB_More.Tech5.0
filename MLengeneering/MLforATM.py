import json
import random
from scipy.stats import beta


def save_filter_data():
    # Преобразуем "плоские" ключи обратно в иерархическую структуру
    services_data = {}
    for key, value in filter_data.items():
        if key.startswith('services_'):
            original_key = key[len('services_'):]
            services_data[original_key] = value

    # Запись в исходный словарь
    filter_data['services'] = services_data
    for key in list(filter_data.keys()):
        if key.startswith('services_'):
            del filter_data[key]

    # Сохранение в файл
    with open("filter_data.json", "w") as f:
        json.dump(filter_data, f, indent=4)


# Чтение исходных данных из файла
with open("filter_data.json", "r") as f:
    filter_data = json.load(f)

# # Преобразование вложенных фильтров в "плоские" ключи
# for key in list(filter_data['services'].keys()):
#     new_key = f'services_{key}'
#     filter_data[new_key] = filter_data['services'][key]
# del filter_data['services']

# # Чтение новых данных из файла
# with open("new_data.json", "r") as f:
#     new_data = json.load(f)


def update_filter_data(filter_name, success):
    data = filter_data.get(filter_name)
    if data is not None:
        data['total_count'] += 1
        if success:
            data['success_count'] += 1
        data['alpha'] = data['success_count'] + 1
        data['beta'] = data['total_count'] - data['success_count'] + 1


def calculate_weights(filter_data):
    weights = {}
    for filter_name, data in filter_data.items():
        if isinstance(data, dict) and 'alpha' in data and 'beta' in data:
            alpha, beta = data['alpha'], data['beta']
            estimated_prob = alpha / (alpha + beta)
            weights[filter_name] = estimated_prob

        # Если у нас есть вложенный словарь, рекурсивно вызываем функцию
        elif isinstance(data, dict):
            weights[filter_name] = calculate_weights(data)

    return weights

# Обновляем данные на основе новых событий
# for sub_array in new_data:  # Добавлен этот цикл для прохода по подмассивам
#     for event in sub_array:
#         update_filter_data(event['filter'], event['success'])
#
# # Вычисляем новые веса для фильтров
# new_weights = calculate_weights()
#
# print("New weights:", new_weights)
#
# # Сохраняем обновленные данные в файл
# save_filter_data()
