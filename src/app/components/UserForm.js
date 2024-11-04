import React, { useState } from 'react';
import { createUser } from '../services/api'; // Import only if needed

const UserForm = () => {
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    bio: '',
    profile_picture: ''
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await createUser(formData); // Ensure `createUser` is defined in your api.js
      alert('User created successfully!');
      setFormData({ name: '', email: '', bio: '', profile_picture: '' }); // Reset the form after submission
    } catch (error) {
      alert('Failed to create user.');
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        name="name"
        placeholder="Name"
        value={formData.name}
        onChange={handleChange}
      />
      <input
        type="email"
        name="email"
        placeholder="Email"
        value={formData.email}
        onChange={handleChange}
      />
      <textarea
        name="bio"
        placeholder="Bio"
        value={formData.bio}
        onChange={handleChange}
      />
      <input
        type="text"
        name="profile_picture"
        placeholder="Profile Picture URL"
        value={formData.profile_picture}
        onChange={handleChange}
      />
      <button type="submit">Submit</button>
    </form>
  );
};

export default UserForm;
