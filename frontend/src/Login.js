import React, { Component } from 'react';
import MuiThemeProvider from '@material-ui/core/styles/MuiThemeProvider';
import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';


class Login extends Component {
constructor(props){
  super(props);
  this.state={
  username:'',
  password:''
  }
 }

render() {
    return (
      <div>
        <MuiThemeProvider>
          <div>
          <AppBar
             title="Login"
           />
           <TextField
             hintText="Enter your Username"
             floatingLabelText="Username"
             onChange = {(event,newValue) => this.setState({username:newValue})}
             />
           <br/>
             <TextField
               type="password"
               hintText="Enter your Password"
               floatingLabelText="Password"
               onChange = {(event,newValue) => this.setState({password:newValue})}
               />
             <br/>
             <Button label="Submit" primary={true} style={style} onClick={(event) => this.handleClick(event)}>
             Sign in
             </Button>
         </div>
         </MuiThemeProvider>
      </div>
    );
  }

  handleClick(event){
    // TODO(isaacsahle): Setup grpc.
    }
}
const style = {
 margin: 15,
};

export default Login;