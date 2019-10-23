'use strict';
const path = require('path');
const express = require('express');
const bodyParser = require('body-parser');
const fs = require('fs');

const PORT = 8080;
const HOST = '0.0.0.0';

const app = express();
app.use(bodyParser.json());
// in latest body-parser use like below.
app.use(bodyParser.urlencoded({extended: true}));

app.use(express.static(path.join(__dirname, 'public')));

const listEventsFiles = (directory) =>
    new Promise(function (resolve, reject) {
        const eventsDirectory = path.join(__dirname, directory);
        fs.readdir(eventsDirectory, (err, files) => {
            const resolved = files.map(f => `${eventsDirectory}/${f}`);
            resolve(resolved);
        });
    });

const getEvents = (directory) =>
    listEventsFiles(directory)
        .then(filesPaths => {
            const readFiles = filesPaths.map(f => fs.readFileSync(f));
            const readFilesJsons = readFiles.map(f => JSON.parse(f));
            const first = readFilesJsons[0];
        });

// const x = getEvents('events');

app.get('/login', (req, res) => {
    res.sendFile(path.join(__dirname + '/views/template.html'));
});

app.post('/do-login', (req, res) => {
    // const eventsDirectory = path.join(__dirname, 'public');
    // fs.readdir(eventsDirectory, (err, files) => {
    //     console.log(JSON.stringify(files));
    // });
    console.log('redirecting to wp');

    res.status(301).redirect('/survey');
});

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);