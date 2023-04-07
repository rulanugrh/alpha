const express = require('express')
const cors = require('cors')
const env = require('./helpers/env')
const userRoute = require('./routes/userRoute')

const app = express()
app.use(cors());
app.use(express.json());
app.use(userRoute);

app.listen(`${env.app.port}`, ()=> console.log('server app runingg ....'))