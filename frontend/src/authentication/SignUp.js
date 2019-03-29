
import React, {Component} from 'react';
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
import { createUser } from '../backend_services/Service';

class SignUp extends Component { // Stay consistent and change to class
  constructor(props){
    super(props);
    
    this.state = { 
      email: '',
      password:'',
      confirmPassword:''
    };
    
    this.classes = props.classes
    this.handler = props.handler
    this.signIn = this.signIn.bind(this)
    this.verifyInformation = this.verifyInformation.bind(this)
    this.creatUserCallback = this.createUserCallback.bind(this)
  }
  
  signIn(){
    this.handler(true /* sign in */, false  /* sign up */)
  }

  verifyInformation(){
    if(this.state.email === "" || this.state.password === "" || this.state.confirmPassword === ""){
      console.log("stop fucking around")
      return
    }

    if(this.state.password !== this.state.confirmPassword){
      alert("Passwords do not match")
      return
    }

    createUser(this.state.email,this.state.password,(err, response) => {this.createUserCallback(err,response)})
  }

  createUserCallback(err, response){
    if (err) {
      alert(err.message)
      console.log(err.code);
      console.log(err.message);
    } else {
      console.log(response)
      this.signIn()
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
            Sign Up
          </Typography>
          <form className={this.classes.form}>
            <FormControl margin="normal" required fullWidth>
              <InputLabel htmlFor="email">Email Address</InputLabel>
              <Input id="email" name="email" autoComplete="email" onChange={(e) => {this.setState({email: e.target.value})}} autoFocus />
            </FormControl>
            <FormControl margin="normal" required fullWidth>
              <InputLabel htmlFor="password">Password</InputLabel>
              <Input name="password" type="password" id="password" onChange={(e) => {this.setState({password: e.target.value})}} />
            </FormControl>
            <FormControl margin="normal" required fullWidth>
              <InputLabel htmlFor="confirm-password">Confirm Password</InputLabel>
              <Input name="confirm-password" type="password" id="confirm-password" onChange={(e) => {this.setState({confirmPassword: e.target.value})}} />
            </FormControl>
            <Button
              type="button"
              fullWidth
              variant="contained"
              color="primary"
              className={this.classes.submit}
              onClick={this.verifyInformation}
            >
              Sign up
            </Button>
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
          </form>
        </Paper>
      </main>
    );
}
}

SignUp.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(SignUp);