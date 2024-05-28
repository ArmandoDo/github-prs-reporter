import React, { useState } from 'react';
import axios from 'axios';
import { useSnackbar } from 'notistack';

export const RepoForm = () => {
  const [owner, setOwner] = useState('');
  const [repo, setRepo] = useState('');
  const [email, setEmail] = useState('');
  const { enqueueSnackbar } = useSnackbar();


  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.get(
        `http://172.17.0.1:8081/pullrequests?owner=${owner}&repo=${repo}&email=${email}`,
      );
        // console.log('GitHub Repo Info:', response.data);
        if (response.data.Successful){
          enqueueSnackbar(response.data.Successful, { variant: 'success' });
        } else {
          enqueueSnackbar(response.data.Skipped, { variant: 'warning' });
        }

    } catch (error) {
      // console.error('Error fetching repo info:', error);
      enqueueSnackbar(error.response.data.Error, { variant: 'error' });
    }
  };

  return (
    <div>
      <h1>Send Pull Request List</h1>
      <form onSubmit={handleSubmit}>
        <label>
          Owner:
          <input
            type="text"
            required={true}
            value={owner}
            onChange={(e) => setOwner(e.target.value)}
          />
        </label>
        <br />
        <br />
        <label>
          Repo:
          <input
            type="text"
            required={true}
            value={repo}
            onChange={(e) => setRepo(e.target.value)}
          />
        </label>
        <br />
        <br />
        <label>
          Email:
          <input
            type="email"
            required={true}
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </label>
        <br />
        <br />
        <button type="submit">Get Repo Info</button>
      </form>
    </div>
  );
};