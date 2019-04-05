import React, {Component} from 'react';
import { withStyles } from '@material-ui/core';
import { TextField } from '@material-ui/core';
import Button from '@material-ui/core/Button';
import { validatePrice, validateStockSymbol } from '../shared/InputUtils';
import { setSellAmount, setSellTrigger, setBuyAmount, setBuyTrigger, cancelSetBuy, cancelSetSell } from '../backend_services/Service';
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


class Triggers extends Component {
    constructor(props) {
      super(props);
      this.state = {
        userId:props.userId, 
        buyTrigger:{amount:0.00, price:0.00, symbol:""},
        sellTrigger:{amount:0.00, price:0.00, symbol:""},
        buyDisabled:false,
        sellDisabled:false
      };
      this.classes = props.classes
      this.handleNumericalChange = this.handleNumericalChange.bind(this)
      this.handleStockChange = this.handleStockChange.bind(this)
      
      this.handleBuyTrigger = this.handleBuyTrigger.bind(this)
      this.setBuyAmountCallback = this.setBuyAmountCallback.bind(this)
      this.setBuyTriggerCallback = this.setBuyTriggerCallback.bind(this)
      this.cancelBuyTrigger = this.cancelBuyTrigger.bind(this)
      this.cancelSetBuyCallback = this.cancelSetBuyCallback.bind(this)

      this.handleSellTrigger = this.handleSellTrigger.bind(this)
      this.setSellAmountCallback = this.setSellAmountCallback.bind(this)
      this.setSellTriggerCallback = this.setSellTriggerCallback.bind(this)
      this.cancelSellTrigger = this.cancelSellTrigger.bind(this)
      this.cancelSetSellCallback = this.cancelSetSellCallback.bind(this)
    }

    handleStockChange(e, buyOrSell){
      if(e.target.value === '' || validateStockSymbol(e.target.value)){  
        const{buyTrigger, sellTrigger} = this.state
        if(buyOrSell === "buy"){
          buyTrigger.symbol = e.target.value
          this.setState({ buyTrigger })
        }else {
          sellTrigger.symbol = e.target.value
          this.setState({ sellTrigger })
        }
      }
    }

    handleNumericalChange(e, buyOrSell, amountOrPrice){
      if(e.target.value === '' || validatePrice(e.target.value)){  
        const{buyTrigger, sellTrigger} = this.state
        if(buyOrSell === "buy"){
          if(amountOrPrice === "amount"){
            buyTrigger.amount = e.target.value
          }else{
            buyTrigger.price = e.target.value
          }
          this.setState({ buyTrigger })
        }else{
          if(amountOrPrice === "amount"){
            sellTrigger.amount = e.target.value
          }else{
            sellTrigger.price = e.target.value
          }
          this.setState({ sellTrigger })
        }
      }
    }

    handleBuyTrigger(){
      if(this.state.buyTrigger.amount > 0 && this.state.buyTrigger.price > 0 && this.state.buyTrigger.symbol !== ""){
        // Disable input fields since we will be chaining callbacks 
        this.setState({buyDisabled:true})
        setBuyAmount(
           this.state.userId,
           this.state.buyTrigger.symbol,
           this.state.buyTrigger.amount,
           (err, response) => {this.setBuyAmountCallback(err, response)}
          ) 
      }else{
        alert("Figure it out nerd")
      }
    }

    cancelBuyTrigger(){
      if(this.state.buyTrigger.symbol !== ""){
        cancelSetBuy(
          this.state.userId,
          this.state.buyTrigger.symbol,
          (err, response) => {this.cancelSetBuyCallback(err,response)}
        )
      }else {
        alert('Enter a stock symbol')
      }
    }

    cancelSetBuyCallback(err, response){
      if(err){
        alert(err.message)
      }else{
        console.log(response)
      }
    }


    setBuyAmountCallback(err, response){
      if(err){
        this.setState({buyDisabled:false})
        alert(err.message)
      }else{
        console.log(response)
        var balance = response.getBalance()
        setBuyTrigger(
          this.state.userId,
          this.state.buyTrigger.symbol,
          this.state.buyTrigger.price,
          (err, response) => {this.setBuyTriggerCallback(err, response, balance)}
        )
      }
    }

    setBuyTriggerCallback(err, response, newBalance){
      if(err){
        alert(err.message)
      }else{
        alert(response.getMessage() + "\nNew balance: " + newBalance.toFixed(2))
      }
      this.setState({buyDisabled:false})
    }

    handleSellTrigger(){
      if(this.state.sellTrigger.amount > 0 && this.state.sellTrigger.price > 0 && this.state.sellTrigger.symbol !== ""){
        // Disable input fields since we will be chaining callbacks 
        this.setState({sellDisabled:true})
        setSellAmount(
           this.state.userId,
           this.state.sellTrigger.symbol,
           this.state.sellTrigger.amount,
           (err, response) => {this.setSellAmountCallback(err, response)}
          ) 
      }else{
        alert("Figure it out nerd")
      }
    }

    cancelSellTrigger(){
      if(this.state.sellTrigger.symbol !== ""){
        cancelSetSell(
          this.state.userId,
          this.state.sellTrigger.symbol,
          (err, response) => {this.cancelSetSellCallback(err,response)}
        )
      }else {
        alert('Enter a stock symbol')
      }
    }

    cancelSetSellCallback(err, response){
      if(err){
        alert(err.message)
      }else{
        console.log(response)
      }
    }

    setSellAmountCallback(err, response){
      if(err){
        this.setState({sellDisabled:false})
        alert(err.message)
      }else{
        setSellTrigger(
          this.state.userId,
          this.state.sellTrigger.symbol,
          this.state.sellTrigger.price,
          (err, response) => {this.setSellTriggerCallback(err, response)}
        )
      }
    }

    setSellTriggerCallback(err, response){
      if(err){
        alert(err.message)
      }else{
        alert("Successfully set sell trigger.")
      }
      this.setState({sellDisabled:false})
    }

    render() {
        return(<div className={this.classes.container}>
          <TextField
              id="outlined-name"
              label="Symbol"
              className={classNames(this.classes.textField, this.classes.item)}
              onChange={(e) => {this.handleStockChange(e, "buy")}}
              value={this.state.buyTrigger.symbol}
              margin="normal"
              variant="outlined"
              autoComplete='off'
              disabled={this.state.buyDisabled}
            />
          <TextField
              id="outlined-name"
              label="Buy Amount"
              className={classNames(this.classes.textField, this.classes.item)}
              onChange={(e) => {this.handleNumericalChange(e,"buy","amount")}}
              value={this.state.buyTrigger.amount}
              margin="normal"
              variant="outlined"
              autoComplete='off'
              disabled={this.state.buyDisabled}
            />
            <TextField
              id="outlined-name"
              label="Buy Price"
              className={classNames(this.classes.textField, this.classes.item)}
              onChange={(e) => {this.handleNumericalChange(e,"buy","price")}}
              value={this.state.buyTrigger.price}
              margin="normal"
              variant="outlined"
              autoComplete='off'
              disabled={this.state.buyDisabled}
            />
            <Button variant="outlined" color="primary" onClick={() => {this.handleBuyTrigger()}} className={this.classes.item}>
            Set Buy Trigger
          </Button>
          <Button variant="outlined" color="primary" onClick={() => {this.cancelBuyTrigger()}} className={this.classes.item}>
            Cancel Buy Trigger
          </Button>
          <br/>
          <br/>
          <br/>
          <TextField
              id="outlined-name"
              label="Symbol"
              className={classNames(this.classes.textField, this.classes.item)}
              onChange={(e) => {this.handleStockChange(e, "sell")}}
              value={this.state.sellTrigger.symbol}
              margin="normal"
              variant="outlined"
              autoComplete='off'
              disabled={this.state.sellDisabled}
            />
          <TextField
              id="outlined-name"
              label="Sell Amount"
              className={classNames(this.classes.textField, this.classes.item)}
              onChange={(e) => {this.handleNumericalChange(e,"sell","amount")}}
              value={this.state.sellTrigger.amount}
              margin="normal"
              variant="outlined"
              autoComplete='off'
              disabled={this.state.sellDisabled}
            />
            <TextField
              id="outlined-name"
              label="Sell Price"
              className={classNames(this.classes.textField, this.classes.item)}
              onChange={(e) => {this.handleNumericalChange(e,"sell","price")}}
              value={this.state.sellTrigger.price}
              margin="normal"
              variant="outlined"
              autoComplete='off'
              disabled={this.state.sellDisabled}
            />
            <Button variant="outlined" color="primary" onClick={() => {this.handleSellTrigger()}} className={this.classes.item}>
            Set Sell Trigger
          </Button>
          <Button variant="outlined" color="primary" onClick={() => {this.cancelSellTrigger()}} className={this.classes.item}>
            Cancel Sell Trigger
          </Button>
          </div>
          )
    }
  }
  
  export default withStyles(styles)(Triggers);