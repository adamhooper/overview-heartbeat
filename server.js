#!/usr/bin/env node

const http = require('http');

const PORT = process.env.PORT || 3000;

const server = http.createServer(function(req, res) {
  res.writeHead(204);
  res.end('');
});

server.listen(PORT, function() {
  console.log("Listening on http://localhost:%s", PORT)
});
