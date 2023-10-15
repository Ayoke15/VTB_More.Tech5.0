ymaps.ready(function () {
    var myMap = new ymaps.Map('map', {
        center: [55.753994, 37.622093],
        zoom: 9,
    });

}
);

/* ПОСТРОЕНИЕ МАРШРУТА

     // Создайте объект для построения маршрута.
     var multiRoute = new ymaps.multiRouter.MultiRoute({
        referencePoints: [
            begin_route,
            end_route,
        ],
        params: {
            routingMode: 'masstransit',
        },
    });

    // Добавьте построенный маршрут на карту.
    myMap.geoObjects.add(multiRoute);

    // Теперь уберем панель маршрутизации.
    // Удалим элемент управления для панели.
    myMap.controls.remove('routePanelControl');

    // Получим длину кратчайшего маршрута и выведем ее в консоль.
    multiRoute.model.events.add('requestsuccess', function () {
        var activeRoute = multiRoute.getActiveRoute();
        if (activeRoute) {
            var routeLength = activeRoute.properties.get('distance').text;
            console.log('Длина кратчайшего маршрута: ' + routeLength);
        }
    });

*/