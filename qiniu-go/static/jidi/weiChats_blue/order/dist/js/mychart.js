var data1 = [
	{
		value: 12345,
		color:"#229ff7"
	},
	{
		value : 2000,
		color : "#fe4b4b"
	}			
]
new Chart(document.getElementById("starbalance").getContext("2d")).Pie(data1);


var data2 = [
	{
		value: 565,
		color:"#229ff7"
	},
	{
		value : 300,
		color : "#fe4b4b"
	}			
]
new Chart(document.getElementById("moneybalance").getContext("2d")).Pie(data2);