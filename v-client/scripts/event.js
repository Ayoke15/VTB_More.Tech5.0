// Получаем кнопки "Отделения" и "Банкоматы" по их id
const officesButton = document.getElementById('offices-button');
const atmsButton = document.getElementById('atms-button');

// Получаем элементы с id "offices" и "atms"
const offices = document.getElementById('offices');
const atms = document.getElementById('atms');

// Добавляем EventListener для кнопки "Банкоматы"
atmsButton.addEventListener('click', () => {
  // Скрываем отделения и отображаем банкоматы
  offices.style.display = none;
  atms.style.display = block;
});

// Добавляем EventListener для кнопки "Отделения"
officesButton.addEventListener('click', () => {
  // Скрываем банкоматы и отображаем отделения
  atms.style.display = none;
  offices.style.display = block;
});
