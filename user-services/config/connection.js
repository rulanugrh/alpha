const mongoose = require('mongoose');
const env = require('../helpers/env');
const conn = mongoose.connection;

const dbConn = mongoose.connect(`mongodb://${env.mongodb.host}:${env.mongodb.port}/${env.mongodb.name}`, {
    user: env.mongodb.user,
    pass: env.mongodb.pass,
    autoIndex: true,
}).then(() => {
    console.log(`mongodb running`)
}).catch((err) => {
    console.log(`something error on ${err.message}`)
});


module.exports = dbConn;