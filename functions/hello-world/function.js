exports.helloWorld = function helloWorld(event, callback) {
	console.log(`My Cloud Function: ${event.data.message}`);
	callback();
}
