import React, {Component} from 'react';
import { withStyles } from '@material-ui/core';
import { TextField } from '@material-ui/core';
import Button from '@material-ui/core/Button';
import { validatePrice } from '../shared/InputUtils';
import { add } from '../backend_services/Service';

const styles = theme => ({
    container: {
      display: 'flex',
      flexWrap: 'wrap',
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

class AddBalance extends Component {
    constructor(props) {
      super(props);
      this.state = {
          userId: props.userId, 
          amount: null
      };
      this.classes = props.classes
      this.buttonPress = this.buttonPress.bind(this)
      this.handleAmountChange = this.handleAmountChange.bind(this)
      this.buyCallback = this.buyCallback.bind(this)
    }

    buttonPress(){
        if(this.state.amount > 0){
            add(this.state.userId, this.state.amount, (err, response)=> {this.buyCallback(err, response)})
        }else {
            console.log("Figure it out you nerd")
        }
    }

    handleAmountChange(e){
        if(e.target.value === '' || validatePrice(e.target.value)){  
          this.setState({
            amount: e.target.value
          });
        }
    }

    buyCallback(err, response){
        if(err){
            alert(err.message)
        }else{
            alert("New balance: " + response.getBalance())
        }
    }

    render() {
        return(
        <div>
        <TextField
            id="outlined-name"
            label="Amount"
            className={this.classes.textField}
            onChange={this.handleAmountChange}
            value={this.state.amount}
            margin="normal"
            variant="outlined"
            autoComplete='off'
          />
          <Button variant="outlined" color="primary" onClick={this.buttonPress}>
          Add
        </Button>
        </div>
          )
    }
  }
  
  export default withStyles(styles)(AddBalance);