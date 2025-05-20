import { useState } from 'react';
import axiosInstance from '../utils/axiosInstance';

export default function UserSearch({ onUserSelect }) {
  const [searchTerm, setSearchTerm] = useState('');
  const [searchResults, setSearchResults] = useState([]);

  const handleSearch = async () => {
    if (!searchTerm.trim()) return;
    try {
      const res = await axiosInstance.get(`/protected/message=${searchTerm}`, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('accessToken')}`
        }
      });
      setSearchResults(res.data);
    } catch (err) {
      console.error('Search error:', err);
    }
  };

  return (
    <div>
      <h3>Search Users</h3>
      <input
        placeholder="Search by email"
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
        style={{ width: '250px', marginRight: '10px' }}
      />
      <button onClick={handleSearch}>Search</button>

      <ul style={{ listStyle: 'none', paddingLeft: 0 }}>
        {searchResults.map(user => (
          <li key={user.id} style={{ marginTop: '8px' }}>
            {user.name} ({user.email}){' '}
            <button onClick={() => onUserSelect(user.id)}>Message</button>
          </li>
        ))}
      </ul>
    </div>
  );
}
