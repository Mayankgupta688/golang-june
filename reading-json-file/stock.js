var data = `{
    "code": "200",
    "message": "Success",
    "company": "SBI",
    "pricechange": "-1.0500",
    "pricepercentchange": "-0.2265"
}`

var stockData = JSON.parse(data)

console.log(stockData.pricechange)

alert(stockData.company)