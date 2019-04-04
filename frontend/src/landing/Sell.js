import React, {Component} from 'react';
import { TextField } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import { validateStockSymbol, validatePrice } from '../shared/InputUtils';
import InputAdornment from '@material-ui/core/InputAdornment';
import { sell, commitSell, cancelSell } from '../backend_services/Service';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import classNames from 'classnames';


const styles = theme => ({
  container: {
    display: 'flex',
    flexDirection: 'column',
    alignContent: 'flex-end',
  },
  textField: {
    marginLeft: theme.spacing.unit,
    marginRight: theme.spacing.unit,
  },
  item : {
    flexGrow: 1,
  },
  dense: {
    marginTop: 16,
  },
  menu: {
    width: 200,
  },
  });

class Sell extends Component {
  constructor(props) {
    super(props);
    this.state = { 
      userId: props.userId,
      stock: '',
      amount: '',
      isModalOpen: false,
    };

    this.classes = props.classes
    this.handler = props.handler
    
    this.handleRejectClose = this.handleRejectClose.bind(this)
    this.handleConfirmClose = this.handleConfirmClose.bind(this)

    this.keyPress = this.keyPress.bind(this)
    this.handleStockChange = this.handleStockChange.bind(this)
    this.handleAmountChange = this.handleAmountChange.bind(this)
    this.buttonPress = this.buttonPress.bind(this)
    this.sellCallback = this.sellCallback.bind(this)
    this.commitSellCallback = this.commitSellCallback.bind(this)
    this.cancelSellCallback = this.cancelSellCallback.bind(this)
    }

    sellCallback(err, response){
      if (err) {
        alert(err.message)
      } else {
        console.log(response)
        this.setState({ isModalOpen: true });
      }
    }

    handleRejectClose(){
      //cancel sell
      cancelSell(this.state.userId, (err, response) => {this.cancelSellCallback(err, response)})
      this.setState({ isModalOpen: false });
    }

    handleConfirmClose(){
      //commit sell
      commitSell(this.state.userId, (err, response) => {this.commitSellCallback(err, response)})
      this.setState({ isModalOpen: false });
    }

    handleStockChange(e){
      if(e.target.value === '' || validateStockSymbol(e.target.value)){
        this.setState({
          stock: e.target.value
        });
      }
    }

    handleAmountChange(e){
      if(e.target.value === '' || validatePrice(e.target.value)){  
        this.setState({
          amount: e.target.value
        });
      }
    }

    keyPress(e){
      if(e.key === 'Enter' && this.state.amount !== ''){
        console.log(this.state.amount)
      }
    }

    buttonPress(){
     if(this.state.amount > 0){
        sell(this.state.userId, this.state.stock, this.state.amount,(err, response) => {this.sellCallback(err,response)})
      }else{
        alert("Figure it out you nerd")
      }
    }

    commitSellCallback(err, response){
      if(err){
        alert(err.message)
      } else{
        console.log("SELL CONFIRMED")
        console.log(response)
      }
    }

    cancelSellCallback(err, response){
      if(err){
        alert(err.message)
      } else{
        console.log("SELL CANCELLED")
        console.log(response)
      }
    }

    render() {
        return(
          <div className={this.classes.container}>
          <TextField
              id="outlined-name"
              label="Stock Symbol"
              className={classNames(this.classes.textField, this.classes.item)}
              onChange={this.handleStockChange}
              value={this.state.stock}
              margin="normal"
              variant="outlined"
            />
            <TextField
              id="outlined-name"
              label="Sell Amount"
              className={classNames(this.classes.textField, this.classes.item)}
              onChange={this.handleAmountChange}
              onKeyDown = {this.keyPress}
              InputProps={{
                startAdornment: <InputAdornment position="start">$</InputAdornment>,
              }}
              value={this.state.amount}
              margin="normal"
              variant="outlined"
            />
            <Button variant="outlined" color="primary" onClick={this.buttonPress} className={this.classes.item}>
              Sell
            </Button>
            <Dialog open={this.state.isModalOpen} onClose={() => {this.setState({ isModalOpen: false })}} aria-labelledby="alert-dialog-title" aria-describedby="alert-dialog-description"
            disableBackdropClick
            disableEscapeKeyDown>
              <DialogTitle id="alert-dialog-title">{"Are you sure you want to make this purchase?"}</DialogTitle>
              <DialogContent>
                <DialogContentText id="alert-dialog-description">
                  By clicking confirm, you are agreeing to sell the amount specified
                </DialogContentText>
              </DialogContent>
              <DialogActions>
                <Button onClick={this.handleRejectClose} color="primary">
                  Reject
                </Button>
                <Button onClick={this.handleConfirmClose} color="primary" autoFocus>
                  Confirm
                </Button>
              </DialogActions>
            </Dialog>
          </div>
        )
    }
  }
  
  export default withStyles(styles)(Sell);