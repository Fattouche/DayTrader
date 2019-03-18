import React, { Component } from 'react';
import './App.css';
import SignIn from './authentication/SignIn';
import SignUp from './authentication/SignUp';
import Landing from './landing/Landing';


class App extends Component {
  constructor(props) {
    super(props);
    this.state = { 
      signIn: true,
      signUp: false,
      landing: false 
    };

    this.loginState = this.loginState.bind(this)
    this.landingState = this.landingState.bind(this)
    this.logout = this.logout.bind(this)
  }

  loginState(showSignIn, showSignUp) {
    this.setState({
      signIn: showSignIn,
      signUp: showSignUp,
      landing: false
    })
  }

  landingState() {
    this.setState({
      signIn: false,
      signUp: false,
      landing:true
    })
  }

  logout() {
    this.loginState(true,false)
  }
  
  render() {
    let screen
    if (this.state.signIn){
      screen = <SignIn parentContext={this} handler={[this.loginState, this.landingState]}/>
    }else if (this.state.signUp){
      screen = <SignUp parentContext={this} handler={this.loginState}/>
    }else{
      screen = <Landing parentContext={this} handler={this.logout}/>
    }
    return (
      <div className='App'>{screen}</div>
    );
  }
}

export default App;
