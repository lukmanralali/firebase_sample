const connection = require("firebase-admin");
const serviceAccount = require("./credential.json");

connection.initializeApp({
  credential: connection.credential.cert(serviceAccount),
  databaseURL: "https://sample-project-267411.firebaseio.com"
});

const firebaseDB = connection.database()

module.exports = {
	connection,
	firebaseDB
}