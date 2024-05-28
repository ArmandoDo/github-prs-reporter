import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import {RepoForm} from './RepoForm'
import { SnackbarProvider, useSnackbar } from 'notistack'

function App() {
  return (
    <SnackbarProvider>
      <div className="App">
        <RepoForm />
      </div>
    </SnackbarProvider>

  );
}

export default App
