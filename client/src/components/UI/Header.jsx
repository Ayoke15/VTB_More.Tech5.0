import './Header.css'; // Import your CSS file for styling
import logo from './VTB_logo_ru.png';

const Header = () => {
  return (
    <header>
        <nav>
          <img src={logo} className='logo'></img>
        </nav>
    </header>
  );
};

export default Header;
