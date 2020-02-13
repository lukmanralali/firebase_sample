'use strict'

const { firebaseRepo } = require('./src/repository')
const uuidv1 = require('uuid/v1')
const currentUUID = uuidv1()
function timeout(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}
const delay = 2000 //ms

const insertData = {
	user1:{
		is_attended:true
	},
	user2:{
		is_attended:false
	}
}
const updateData = {
	is_attended:true
}

async function startApp(){
	console.log(`current: ${currentUUID}`)
	// firebaseRepo.select('/')
	console.log('inserting data')
	firebaseRepo.upsert(insertData, '/', currentUUID)
	await timeout(delay)
	console.log('selecting data')
	firebaseRepo.select(`/${currentUUID}`)
	await timeout(delay)
	console.log(`updating data`)
	firebaseRepo.upsert(updateData, `/${currentUUID}`, 'user2')
	await timeout(delay)
	console.log(`select data: ${currentUUID}`)
	firebaseRepo.select(`/${currentUUID}`)
}
startApp()
