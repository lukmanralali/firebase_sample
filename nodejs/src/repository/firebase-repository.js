'use strict'
const { connection } = require('./../utility/firebase')
const { firebaseDB } = connection 
const select = (reff) => {
	const ref = firebaseDB.ref(reff)
	ref.once("value", function(snapshot) {
		console.log(snapshot.val());
	});
}

const upsert = (payload, routeKey, id) => {
    const ref = firebaseDB.ref(`${routeKey}`)
    const childRef = ref.child(id)
	childRef.update(payload);
}

module.exports = {
	select,
    upsert
}