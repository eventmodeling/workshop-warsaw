'use strict';

const express = require('express');

// Constants
const PORT = 8080;
const HOST = '0.0.0.0';

// App
const app = express();

app.get('/websiteLogin', (req, res) => {
    res.send('login website\n');
});

app.get('/', (req, res) => {
    res.send('We are in root\n');
});

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);