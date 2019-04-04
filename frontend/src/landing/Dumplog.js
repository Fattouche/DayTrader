import React, {Component} from 'react';
import { TextField } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import { dumplog } from '../backend_services/Service';
import Button from '@material-ui/core/Button';
import classNames from 'classnames';


const styles = theme => ({
  container: {
    display: 'flex',
    flexDirection: 'column',
    alignContent: 'flex-end',
  },
  item : {
    flexGrow: 1,
  },
  textField: {
    marginLeft: theme.spacing.unit,
    marginRight: theme.spacing.unit,
  },
  dense: {
    marginTop: 16,
  },
  menu: {
    width: 200,
  },
  });
  

class Dumplog extends Component {
    constructor(props) {
      super(props);
      this.state = { 
        userId: props.userId,
        filename: '',
      };

      this.classes = props.classes
      this.handleChange = this.handleChange.bind(this)
      this.buttonPress = this.buttonPress.bind(this)
      this.dumplogCallback = this.dumplogCallback.bind(this)
    }

    handleChange(e){
        this.setState({filename: e.target.value});
    }

    dumplogCallback(err, response){
      if (err) {
        alert(err.message)
      } else {
        console.log(response)
        alert(response.getMessage())
      }
    }

    buttonPress(){
        if(this.state.filename !== ''){
            
            dumplog(
                this.state.userId,
                this.state.filename,
                (err, response) => {this.dumplogCallback(err, response)}
            )
        } else {
            alert("Filename can't be empty")
        }
    }

    render() {
      return(
        <div className={this.classes.container}>
          <TextField
              id="outlined-name"
              label="Filename"
              className={classNames(this.classes.textField, this.classes.item)}
              onChange={this.handleChange}
              value={this.state.filename}
              margin="normal"
              variant="outlined"
              autoComplete='off'
            />
            <Button variant="outlined" color="primary" onClick={this.buttonPress} className={this.classes.item}>
              Dumplog
            </Button>
        </div>);
    }
  }
  
  export default withStyles(styles)(Dumplog);