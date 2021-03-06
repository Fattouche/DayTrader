
import React, { Component } from 'react';
import PropTypes from 'prop-types';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import FormControl from '@material-ui/core/FormControl';
import Input from '@material-ui/core/Input';
import InputLabel from '@material-ui/core/InputLabel';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';
import withStyles from '@material-ui/core/styles/withStyles';
import styles from '../styles/SignInStyles';
import {checkCredentials} from '../backend_services/Service';


class SignIn extends Component {
  constructor(props) {
    super(props);
    this.state = { 
      email: '',
      password:''
    };

    this.classes = props.classes
    this.handler = props.handler[0]
    this.landing = props.handler[1]
    this.signUp = this.signUp.bind(this)
    this.signIn = this.signIn.bind(this)
    this.signInCallback = this.signInCallback.bind(this)
  }

  signUp(){
    this.handler(false /* sign in */, true  /* sign up */)
  }
  
  signIn(){
    if(this.state.email === "" || this.state.password === ""){
      console.log("stop fucking around")
      return
    }

    checkCredentials(this.state.email, this.state.password, (err, response) => {this.signInCallback(err, response)});
  }

  signInCallback(err, response){
      if (err) {
        alert(err.message)
        console.log(err.code);
        console.log(err.message);
      } else {
        console.log(response)
        this.landing(response)
      }
  }
  
render(){
  return (
    <main className={this.classes.main}>
      <CssBaseline />
      <Paper className={this.classes.paper}>
        <Avatar className={this.classes.avatar}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign in
        </Typography>
        <form className={this.classes.form}>
          <FormControl margin="normal" required fullWidth>
            <InputLabel htmlFor="email">Email Address</InputLabel>
            <Input id="email" name="email" onChange={(e) => {this.setState({email: e.target.value})}}/>
          </FormControl>
          <FormControl margin="normal" required fullWidth>
            <InputLabel htmlFor="password">Password</InputLabel>
            <Input name="password" type="password" id="password" onChange={(e) => {this.setState({password: e.target.value})}}/>
          </FormControl>
          <Button
            type="button"
            fullWidth
            variant="contained"
            color="primary"
            className={this.classes.submit}
            onClick={this.signIn}
          >
            Sign in
          </Button>
          <Button
            type="button"
            fullWidth
            variant="contained"
            color="primary"
            className={this.classes.submit}
            onClick={this.signUp}
          >
            Sign up
          </Button>
        </form>
      </Paper>
    </main>
  );
}

}

SignIn.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(SignIn);