import express from 'express'
import cors from "cors"
import env from "./helpers/env.js"
import userRoute from './routes/userRoute.js'

const app = express()
app.use(cors());
app.use(express.json());
app.use(userRoute);

app.listen(`${env.app.port}`, ()=> console.log('server app runingg ....'))