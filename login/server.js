'use strict';
const path = require('path');
const express = require('express');

const PORT = 8080;
const HOST = '0.0.0.0';

const app = express();

app.use(express.static(path.join(__dirname, 'public')));

app.get('/websiteLogin', (req, res) => {
    const body = req.body;
    console.log(JSON.stringify(body));
    res.sendFile(path.join(__dirname + '/views/template.html'));
});

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);