# Flask API для поиска офисов и банкоматов

Данный проект предоставляет API, который позволяет пользователям находить ближайшие офисы и банкоматы на основе заданных фильтров.

## Оглавление

- [Установка](#установка)
- [Запуск](#запуск)
- [API Endpoints](#api-endpoints)
  - [Получение офисов](#получение-офисов)
  - [Получение банкоматов](#получение-банкоматов)
- [Контакты](#контакты)

## Установка

1. Убедитесь, что у вас установлен Python.
2. Установите необходимые библиотеки:
```bash
pip install Flask
```
3. Клонируйте этот репозиторий:
```bash
git clone https://github.com/Ayoke15/VTB_More.Tech5.0
cd VTB_More.Tech5.0
```

## Запуск

Для запуска сервера используйте команду:

```bash
python API.py
```

Сервер будет доступен по адресу `http://localhost:8080`.

## API Endpoints

### Получение офисов

**URL**: `/get_offices`

**Метод**: `POST`

**Параметры запроса**:
- `user_location`: массив из двух элементов, представляющих координаты пользователя (широта и долгота).
- `user_filters`: объект с фильтрами, указанными пользователем.

**Пример запроса**:
```json
{
    "user_location": [55.7558, 37.6176],
    "user_filters": {
    "openHours": [
      {
        "days": "пт",
        "hours": "09:00-17:00"
      }
    ],
    "rko": "есть РКО",
    "officeType": "Да (Зона Привилегия)",
  }
}
```

**Ответ**:
Возвращает топ-5 офисов, отсортированных по релевантности.

### Получение банкоматов

**URL**: `/get_atms`

**Метод**: `POST`

**Параметры запроса**:
- `user_location`: массив из двух элементов, представляющих координаты пользователя (широта и долгота).
- `user_filters`: объект с фильтрами, указанными пользователем.

**Пример запроса**:
```json
{
    "user_location": [55.7558, 37.6176],
    "user_filters": {
    "services": {
      "wheelchair": {
        "serviceCapability": "UNKNOWN",
        "serviceActivity": "UNKNOWN"
      },
      "nfcForBankCards": {
        "serviceCapability": "UNKNOWN",
        "serviceActivity": "UNAVAILABLE"
      },
      "supportsChargeRub": {
        "serviceCapability": "SUPPORTED",
        "serviceActivity": "AVAILABLE"
      },
      "supportsRub": {
        "serviceCapability": "SUPPORTED",
        "serviceActivity": "AVAILABLE"
      }
    }
  }
}
```

**Ответ**:
Возвращает топ-5 банкоматов, отсортированных по релевантности.

## Контакты

Если у вас есть вопросы или предложения, пожалуйста, свяжитесь со мной по [email@example.com](mailto:email@example.com).
