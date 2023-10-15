import BankBranch from '../BankBranch/BankBranch';
import SearchBar from '../SearchBar/SearchBar';
import SegmentSelector from '../SegmentSelector/SegmentSelector';

import './Menu.css';

const Menu = () => {
  return (
    <div className="menu">
        <SearchBar></SearchBar>
        <SegmentSelector></SegmentSelector>
        <BankBranch></BankBranch>
    </div>
  )
}

export default Menu