import React, { Component } from 'react';
import './App.css';
import SignIn from './SignIn';

class App extends Component {
  render() {
    return (
      <div className='App'><SignIn parentContext={this}/></div>
    );
  }
}

export default App;
