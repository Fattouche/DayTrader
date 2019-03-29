import React, {Component} from 'react';
import { TextField } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import { getQuote } from '../backend_services/Service';
import { validateStockSymbol } from '../shared/InputUtils';
import InputAdornment from '@material-ui/core/InputAdornment';

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
  

class Browse extends Component {
    constructor(props) {
      super(props);
      this.state = { 
        stock: '',
        price: '',
        userId: props.userId,
      };

      this.classes = props.classes
      this.handler = props.handler
      this.handleChange = this.handleChange.bind(this)
      this.keyPress = this.keyPress.bind(this)
      this.getQuoteCallback = this.getQuoteCallback.bind(this)
      // this.validateStockSymbol = this.validateStockSymbol.bind(this)
    }

    handleChange(e){
      if(e.target.value === '' || validateStockSymbol(e.target.value)){
        this.setState({
          stock: e.target.value
        });
      }
    }

    getQuoteCallback(err, response){
      if (err) {
        alert(err.message)
      } else {
        this.setState({price: parseFloat(response.getPrice().toFixed(3))});
      }
    }

    keyPress(e){
      if(e.key === 'Enter' && this.state.stock !== ''){
        getQuote(this.state.userId, this.state.stock,(err, response) => {this.getQuoteCallback(err,response)})
      }
   };

    render() {
      return(
        <div>
          <TextField
              id="outlined-name"
              label="Stock Symbol"
              className={this.classes.textField}
              onChange={this.handleChange}
              onKeyDown = {this.keyPress}
              value={this.state.stock}
              margin="normal"
              variant="outlined"
              autoComplete='off'
            />
            <TextField
              id="outlined-name"
              label="Stock Price"
              className={this.classes.textField}
              InputProps={{
                startAdornment: <InputAdornment position="start">$</InputAdornment>,
              }}
              value={this.state.price}
              margin="normal"
              variant="outlined"
            />
        </div>);
    }
  }
  
  export default withStyles(styles)(Browse);