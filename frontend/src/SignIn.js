
import React, { Component } from 'react';
import PropTypes from 'prop-types';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import FormControl from '@material-ui/core/FormControl';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Input from '@material-ui/core/Input';
import InputLabel from '@material-ui/core/InputLabel';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';
import withStyles from '@material-ui/core/styles/withStyles';
import styles from './styles/SignInStyles';
import proto from './proto/credentials_grpc_web_pb';


class SignIn extends Component {
  constructor(props) {
    super(props);
    this.state = { 
      email: '',
      password:'' 
    };

    this.classes = props.classes
    this.handler = props.handler
    this.signUp = this.signUp.bind(this)
    this.signIn = this.signIn.bind(this)
  }

  signUp(){
    this.handler(false /* sign in */, true  /* sign up */)
  }

  // TODO(isaac): grpc-web works, but need to change backend service
  // proto to be able to retrieve users.
  // TODO(isaac): fix weird form submission bug where async call doesn't finish
  signIn(){
    // if(this.state.email === "" || this.state.password === ""){
    //   console.log("stop fucking around")
    //   return
    // }
    
    // var client = new proto.DayTraderClient('http://localhost:80')
    // var command = new proto.command()
    // command.setUserId(this.state.email)
    // command.setPassword(password)
    // var metadata = {'custom-header-1': 'value1'};
    // var call = client.createUser(command, metadata, function(err, response) {
    // if (err) {
    //   console.log(err.code);
    //   console.log(err.message);
    // } else {
    //   console.log(response.getMessage());
    //   // TODO(isaac): show landing page on successful
    // }
    // });

    // call.on('status', function(status) {
    // console.log(status.code);
    // console.log(status.details);
    // console.log(status.metadata);
    // })
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
            <Input id="email" name="email" autoComplete="email" autoFocus onChange={(e) => {this.setState({email: e.target.value})}}/>
          </FormControl>
          <FormControl margin="normal" required fullWidth>
            <InputLabel htmlFor="password">Password</InputLabel>
            <Input name="password" type="password" id="password" autoComplete="current-password" onChange={(e) => {this.setState({password: e.target.value})}}/>
          </FormControl>
          <FormControlLabel
            control={<Checkbox value="remember" color="primary" />}
            label="Remember me"
          />
          <Button
            type="submit"
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