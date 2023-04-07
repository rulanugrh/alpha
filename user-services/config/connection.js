const Sequelize = require('sequelize');
const env = require('../helpers/env');

const db = new Sequelize(`${env.sequelize.name}`, `${env.sequelize.user}`, `${env.sequelize.pass}`, {
    host: `${env.sequelize.host}`,
    dialect: "mysql"
})

const dbConn = db.authenticate().then(() => {
    return true
}).catch(err => {
    console.log(new Error(`cant connect db because ${err}`))
    return false
});

export default db;