var http = require('http');


const PORT = 8090;
var responses = {};


function testHandler(request, response, error) {
  responses = {"Name":"HelloWorld","Types":["hello","world"]};
  response.end(JSON.stringify(responses));
}


var server = http.createServer(testHandler);

server.listen(PORT, function() {
    console.log(`Listening at http://localhost:${PORT.toString()}`);
});
