const YandexMaps = () => {
    const ymaps = window.ymaps;
      ymaps.ready(() => {
        const map = new ymaps.Map('map', {
          center: [55.751574, 37.573856],
          zoom: 10,
          controls: ['routePanelControl']
        });

        // Дальнейший код инициализации карты

        map.controls.remove('fullscreenControl');
      });
    
  return (
    <div
      id="map"
      style={{
        width: '100vw',
        height: '100vh',
        position: 'absolute',
      }}
    ></div>
  );
};

export default YandexMaps;
