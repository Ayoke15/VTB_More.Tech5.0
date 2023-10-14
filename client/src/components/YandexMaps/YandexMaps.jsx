import { Component } from 'react';
import { offices } from './Offices';

class YandexMaps extends Component {
    componentDidMount() {
      // Инициализация карты после загрузки компонента
      const ymaps = window.ymaps;
  
      if (ymaps) {
        ymaps.ready(() => {
          let map = new ymaps.Map('map', {
            center: [55.751574, 37.573856], // Координаты центра карты
            zoom: 10, // Масштаб карты
            control: [],
          });
  
          // Функция для загрузки данных офисов из JSON файла
          const loadOfficesFromJSON = () => {
            fetch(offices)
              .then((response) => response.json())
              .then((data) => {
                if (Array.isArray(data)) {
                  data.forEach((officeData) => {
                    const { latitude, longitude, salePointName, address, status, openHours, rko, openHoursIndividual, officeType, salePointFormat, suoAvailability, hasRamp } = officeData;
  
                    const placemark = new ymaps.Placemark(
                      [latitude, longitude], // Координаты маркера
                      {
                        hintContent: salePointName, // Подсказка
                        balloonContent: `
                          <div>
                            <h3>${salePointName}</h3>
                            <p>Адрес: ${address}</p>
                            <p>Статус: ${status}</p>
                            <p>Режим работы: ${JSON.stringify(openHours)}</p>
                            <p>РКО: ${rko}</p>
                            <p>Режим работы (индивидуальный): ${JSON.stringify(openHoursIndividual)}</p>
                            <p>Тип офиса: ${officeType}</p>
                            <p>Формат отделения: ${salePointFormat}</p>
                            <p>Наличие СУО: ${suoAvailability}</p>
                            <p>Наличие рампы: ${hasRamp}</p>
                          </div>
                        `, // Информация во всплывающей подсказке
                      }
                    );
  
                    // Добавляем обработчик клика на маркер для отображения всплывающего окна
                    placemark.events.add('click', () => {
                      map.balloon.open([latitude, longitude], placemark.properties.get('balloonContent'));
                    });
  
                    map.geoObjects.add(placemark);
                  });
                }
              })
              .catch((error) => {
                console.error('Ошибка при загрузке данных офисов из JSON файла:', error);
              });
          };
  
          // Вызываем функцию для загрузки данных офисов из JSON файла
          loadOfficesFromJSON();
        });
      }
    }
  
    render() {
      return (
        <div
          id="map"
          style={{ width: '100vw', height: '100vh', position: 'absolute' }}
        ></div>
      );
    }
  }
  

export default YandexMaps;