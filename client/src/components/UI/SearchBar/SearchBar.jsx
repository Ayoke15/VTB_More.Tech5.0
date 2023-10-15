import { useState } from 'react';

function SearchBar() {
  const [inputValue, setInputValue] = useState('');

  return (
    <div style={{ position: 'absolute', left: '31px', top: '223px' }}>
      <input
        type="text"
        value={inputValue}
        onChange={e => setInputValue(e.target.value)}
        placeholder="Город, район, улица..."
        style={{
            display: 'flex',
            justifyContent: 'center',
          marginTop: '50px',
          width: '351px',
          height: '45px',
          background: '#DCE0EB',
          borderRadius: '10px',
          border: 'none',
          paddingLeft: '10px',
          fontSize: '16px'
        }}
      />
    </div>
  );
}

export default SearchBar;