var net = require('net')

var ip = '127.0.0.1'

var client = new net.Socket()

client.connect(6969, ip, function(){
	console.log('Connected')
	add('add more todo items')
})

client.on('data', function(d){
	console.log(d.toString())
})

function add(d){
	client.write('add'+d)
}

function rm(d){
	client.write('rm-'+d)
}

function get(){
	client.write('get')
}
