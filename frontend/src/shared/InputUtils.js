export function validateStockSymbol(symbol){
    var re = /^[a-zA-Z]{1,3}$/;
    return re.test(symbol)
}

//Must be a number between 0 and MAX_AMOUNT
export function validatePrice(price){
    // var re = /^\d{1,}$/
    var re = /^\d{1,}\.*\d{0,2}$/
    return (re.test(price) && price < 1000000000000)
}