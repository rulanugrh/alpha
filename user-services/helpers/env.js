const dotenv = require('dotenv');
dotenv.config();

const dotenv = {
    mongodb: {
        host: process.env.MONGODB_HOST,
        port: process.env.MONGODB_PORT,
        name: process.env.MONGODB_NAME
    },

    app: {
        host: process.env.APP_HOST,
        port: process.env.APP_PORT
    },

    jwtsecret: process.env.JWT_SECRET
}

module.exports = dotenv;