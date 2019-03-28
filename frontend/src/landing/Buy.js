import React, {Component} from 'react';
import { List, ListItem, TextField } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import { validateStockSymbol, validatePrice } from '../shared/InputUtils';
import InputAdornment from '@material-ui/core/InputAdornment';
import { buy, commitBuy, cancelBuy } from '../backend_services/Service';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';

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

class Buy extends Component {
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

      this.handleClickOpen = this.handleClickOpen.bind(this)
      this.handleClose = this.handleClose.bind(this)
      this.handleRejectClose = this.handleRejectClose.bind(this)
      this.handleConfirmClose = this.handleConfirmClose.bind(this)
      this.keyPress = this.keyPress.bind(this)
      this.handleStockChange = this.handleStockChange.bind(this)
      this.handlePriceChange = this.handlePriceChange.bind(this)
      this.buttonPress = this.buttonPress.bind(this)
      this.buyCallback = this.buyCallback.bind(this)
      this.commitBuyCallback = this.commitBuyCallback.bind(this)
      this.cancelBuyCallback = this.cancelBuyCallback.bind(this)
    }

    handleClickOpen(){
      this.setState({ isModalOpen: true });
    };
  
    handleClose(){
      this.setState({ isModalOpen: false });
    };

    handleRejectClose(){
      //cancel buy
      cancelBuy(this.state.userId, (err, response) => {this.cancelBuyCallback(err, response)})
      this.handleClose()
    }

    handleConfirmClose(){
      //commit buy
      commitBuy(this.state.userId, (err, response) => {this.commitBuyCallback(err, response)})
      this.handleClose()
    }

    keyPress(e){
      if(e.key == 'Enter' && this.state.amount !== ''){
        console.log(this.state.amount)
      }
    }

    handleStockChange(e){
      if(e.target.value === '' || validateStockSymbol(e.target.value)){
        this.setState({
          stock: e.target.value
        });
      }
    }

    handlePriceChange(e){
      if(e.target.value === '' || validatePrice(e.target.value)){  
        this.setState({
          amount: e.target.value
        });
      }
    }

    buttonPress(){
      buy(this.state.userId, this.state.stock, this.state.amount,(err, response) => {this.buyCallback(err,response)})
    }
    
    buyCallback(err, response){
      if (err) {
        alert(err.message)
      } else {
        //modal dialogue asking: Are you sure you want to make this buy?
        this.handleClickOpen()
      }
    }

    commitBuyCallback(err, response){
      if(err){
        alert(err.message)
      } else{
        console.log("BUY CONFIRMED")
        console.log(response)
      }
    }

    cancelBuyCallback(err, response){
      if(err){
        alert(err.message)
      } else{
        console.log("BUY CANCELLED")
        console.log(response)
      }
    }

    render() {
      return(
        <div>
          <TextField
              id="outlined-name"
              label="Stock Symbol"
              className={this.classes.textField}
              onChange={this.handleStockChange}
              value={this.state.stock}
              margin="normal"
              variant="outlined"
            />
            <TextField
              id="outlined-name"
              label="Purchase Amount"
              className={this.classes.textField}
              onChange={this.handlePriceChange}
              onKeyDown = {this.keyPress}
              InputProps={{
                startAdornment: <InputAdornment position="start">$</InputAdornment>,
              }}
              value={this.state.amount}
              margin="normal"
              variant="outlined"
            />
            <Button variant="outlined" color="primary" onClick={this.buttonPress}>
              Buy
            </Button>
            <Dialog open={this.state.isModalOpen} onClose={this.handleClose} aria-labelledby="alert-dialog-title" aria-describedby="alert-dialog-description">
              <DialogTitle id="alert-dialog-title">{"Are you sure you want to make this purchase?"}</DialogTitle>
              <DialogContent>
                <DialogContentText id="alert-dialog-description">
                  By clicking confirm, you are agreeing to spend the amount specified
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
          </div>);
    }
  }
  
  export default withStyles(styles)(Buy);