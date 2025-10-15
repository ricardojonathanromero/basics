package main

/*
| name  	| start | end 	|
| Abi 		| 10    | 100  	|
| Ben 		| 50 	| 70	|
| John		| 60 	| 120	|
| Robert	| 120 	| 300	|


output:
| start | end 	| name  			|
| 10   	| 50  	| Abi   			|
| 50   	| 60  	| Abi, Ben   		|
| 60   	| 70  	| Abi, Ben John   	|
| 70	| 100	| Abi, John			|
| 100	| 120	| John				|
| 120	| 300	| Robert			|
*/

/*
input = [
	{
		name: "Abi"
		start: 10
		end: 100
	},
	{
		name: "Ben"
		start: 50
		end: 70
	},
	{
		name: "John"
		start: 60
		end: 120
	},
	{
		name: "Robert"
		start: 120
		end: 300
	}
]

{
	"1050": ["Abi"]
	"5060": ["Abi", "Ben"]
	"6070": ["Abi", "Ben", "John"]
	"70100": ["Abi", "John"]
	"100120": ["John"]
}
*/

func main() {
}
