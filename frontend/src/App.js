import React, { Component } from 'react';
import './App.css';
import SignIn from './SignIn';
import SignUp from './SignUp';


class App extends Component {
  constructor(props) {
    super(props);
    this.state = { 
      signIn: true,
      signUp: false 
    };

    this.loginState = this.loginState.bind(this)
  }

  loginState(showSignIn, showSignUp) {
    this.setState({
      signIn: showSignIn,
      signUp: showSignUp
    })
  }

  render() {
    return (
      <div className='App'>
      {
        this.state.signIn ?
          <SignIn parentContext={this} handler={this.loginState}/>
        :
          <SignUp parentContext={this} handler={this.loginState}/>
      }
      </div>
    );
  }
}

export default App;
