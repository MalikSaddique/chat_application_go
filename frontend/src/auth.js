import axios from './utils/axiosInstance';

console.log("setting up apis")
export const signup = (email, password) => {
  return axios.post('/signup', { email, password });
};


export const login = (email, password) => {
  return axios.post('/login', { email, password });
};

// export const send = (receiver_id, message) =>{
//   return axios.post('/send', {receiver_id, message})
// }