import logo from './logo.svg';
import './App.css';
import React from 'react';
import Home from './pages/home';
import { Route, Routes } from 'react-router-dom';
import DevPage from './pages/dev';

function App() {
  return (
    <Routes>
      <Route path="/dev" element={<DevPage />} />
      <Route path="/" element={<Home />} />
    </Routes>
  );
}

export default App;
