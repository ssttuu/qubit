exports.helloCloudBuilds = function helloCloudBuilds(event, callback) {
	console.log(`My Cloud Function: ${event.data}`);
	callback();
}

