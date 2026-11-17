/*
 * Local Node.js Express Mock Server for Go Web Requests
 * Author: Ravi Ranjan Singh
 * License: MIT
 */

const express = require('express');
const app = express();
const port = process.env.PORT || 8000;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get('/', (req, res) => {
  res.status(200).send("Welcome to Node.js Mock Server authored by Ravi Ranjan Singh");
});

app.get('/get', (req, res) => {
  res.status(200).json({
    message: "Hello from Mock Server",
    author: "Ravi Ranjan Singh",
    timestamp: new Date().toISOString()
  });
});

app.post('/post', (req, res) => {
  res.status(200).json(req.body);
});

app.post('/postform', (req, res) => {
  res.status(200).json(req.body);
});

app.listen(port, () => {
  console.log(`Mock server active at http://localhost:${port}`);
});
