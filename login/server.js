'use strict';
const path = require('path');
const express = require('express');
const bodyParser = require('body-parser');
const fs = require('fs');
const sha256 = require('js-sha256');

const PORT = 8080;
const HOST = '0.0.0.0';

const app = express();
app.use(bodyParser.json());
// in latest body-parser use like below.
app.use(bodyParser.urlencoded({extended: true}));

app.use(express.static(path.join(__dirname, 'public')));

const listEventsFiles = (directory, prefix) =>
    new Promise(function (resolve, reject) {
        const eventsDirectory = path.join(__dirname, directory);
        fs.readdir(eventsDirectory, (err, files) => {
            const resolved = files.filter(f => f.startsWith(prefix)).map(f => `${eventsDirectory}/${f}`);
            resolve(resolved);
        });
    });

const getEvents = (directory) =>
    listEventsFiles(directory, 'user_registered')
        .then(filesPaths => filesPaths.map(f => fs.readFileSync(f)).map(f => JSON.parse(f)));

const credentialsCorrect = (user, password) =>
    getEvents('events')
        .then(users => {
            const foundUser = users.find(u => u.name === user && u.password_hash === sha256(password));
            return !!foundUser;
        });

// credentialsCorrect('Milosz', 'wrong').then(q => console.log(q));
// credentialsCorrect('abc', 'abc').then(q => console.log(q));


app.get('/login', (req, res) => {
    res.sendFile(path.join(__dirname + '/views/template.html'));
});

app.post('/do-login', (req, res) => {
    const user = req.body.username;
    const password = req.body.password;
    console.log(req.body);

    return credentialsCorrect(user, password)
        .then(correct => {
            console.log(correct);
            if (correct) {
                res.status(301).redirect('/survey');
            } else {
                // TODO should return 401 and render error page
                res.status(301).redirect('/login');
            }
        })
});

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);