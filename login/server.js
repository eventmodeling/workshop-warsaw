'use strict';

const express = require('express');

// Constants
const PORT = 8080;
const HOST = '0.0.0.0';

// App
const app = express();
app.get('/', (req, res) => {
    res.send('Hello world root\n');
});

app.get('/hello', (req, res) => {
    res.send('hello\n');
});

app.get('/hello2', (req, res) => {
    res.send('jesteśmy gdzie indziej\n');
});

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);