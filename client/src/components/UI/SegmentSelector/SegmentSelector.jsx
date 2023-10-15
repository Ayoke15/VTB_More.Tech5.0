import './SegmentSelector.css';

const SegmentSelector = () => {
  return (
    <div className="segment">
      <h3>Выбрать сегмент</h3>
      <div className="button-group">
        <button className="segment-button">
          Физические лица
        </button>
        <button className="segment-button">
          Привилегия
        </button>
        <button className="segment-button">
          Юридические лица
        </button>
        <button className="segment-button">
          Прайм
        </button>
    <h3>Уточнить услуги</h3>
        <button className="segment-button">
          Доступно для маломобильных граждан
        </button>
        <button className="segment-button">
          Рядом метро
        </button>
      <button className='search_office'>
        Подобрать отделение
      </button>
      </div>
    </div>
  );
}

export default SegmentSelector;