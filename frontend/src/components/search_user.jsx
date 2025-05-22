import { useState, useEffect } from 'react';
import axiosInstance from '../utils/axiosInstance';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faSearch } from '@fortawesome/free-solid-svg-icons';


export default function UserSearch({ onUserSelect }) {
  const [searchTerm, setSearchTerm] = useState('');
  const [searchResults, setSearchResults] = useState([]);
  const [showDropdown, setShowDropdown] = useState(false);
  const [selectedUser, setSelectedUser] = useState(null);

  useEffect(() => {
    const delayDebounce = setTimeout(() => {
      if (searchTerm.trim()) {
        handleSearch();
      } else {
        setSearchResults([]);
      }
    }, 300); 

    return () => clearTimeout(delayDebounce);
  }, [searchTerm]);

  const handleSearch = async () => {
    try {
      const res = await axiosInstance.get(`/protected/search`, {
        params: { email: searchTerm },
        headers: {
          Authorization: `Bearer ${localStorage.getItem('accessToken')}`,
        },
      });
      setSearchResults(res.data.users || []);
      setShowDropdown(true);
    } catch (err) {
      console.error('Search error:', err);
    }
  };

  const handleUserClick = (user) => {
    setSelectedUser(user);
    onUserSelect(user.id);
    setSearchResults([]);
    setSearchTerm(''); 
  };

  return (
     <div className="user-search-container">
      <input
        className="user-search-input"
        placeholder="Search by email"
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
      />
      


      {searchResults.length > 0 && (
        <ul className="dropdown-list">
          {searchResults.map(user => (
            <li key={user.id} className="dropdown-item" onClick={() => handleUserClick(user)}>
              ({user.email})
            </li>
          ))}
        </ul>
      )}

      {selectedUser && (
        <div className="selected-user">
        <strong>Selected User:</strong>({selectedUser.email})
      </div>
      
      )}
    </div>
  );
}
