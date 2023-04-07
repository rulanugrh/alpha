const dotenv = require('dotenv');
dotenv.config();

const dotenv = {
    sequelize: {
        host: process.env.MYSQL_HOST,
        port: process.env.MYSQL_PORT,
        name: process.env.MYSQL_NAME,
        user: process.env.MYSQL_USER,
        pass: process.env.MYSQL_PASS
    },

    app: {
        host: process.env.APP_HOST,
        port: process.env.APP_PORT
    },

    jwtsecret: process.env.JWT_SECRET
}

module.exports = dotenv;