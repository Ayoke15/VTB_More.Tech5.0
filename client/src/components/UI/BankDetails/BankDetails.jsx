import './BankDetails.css';

const BankDetails = () => {
    return (
        <div className="bank-container">
            <div className="bank-card">
                <div className="bank-header">
                    <p className="bank-details">ДО «Солнечногорский» Филиала № 7701...<br/>141506, Московская область, г. Солнечногорск, ул. Красная, д. 60</p>
                </div>
                <p className="work-mode">Режим работы отделения</p>
                <p className="timing-details">Для физических лиц пн-пт: 10:00-19:00 сб, вс: выходной<br/>Для юридических лиц пн-чт: 10:00-19:00 пт: 10:00-18:00 сб, вс: выходной</p>
                <p className="metro-details-header">Ближайшая станция метро</p>
                <p className="metro-details">МЦД-1 Белорусско-Савёловский диаметр, станция Долгопрудная</p>
                <p className="accessibility">Доступная среда</p>
                <div className="bank-load">
                    <div className="load-indicator"></div>
                    <p className="load-header">Загруженность отделения на данный момент</p>
                    <p className="load-details">Отделение сейчас сильно загружено, лучше попробовать обратиться в другое отделение</p>
                </div>
                <button className="select-branch-btn">Подобрать отделения</button>
            </div>
        </div>
    );
}

export default BankDetails;