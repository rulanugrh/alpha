import user from '../models/userModel.js'
import bcrypt from 'bcrypt';
import env from '../helpers/env.js'
import jwt from 'jsonwebtoken'
import amqp from 'amqplib';

export const testingToken = async(req, res) => {
    try {
        jwt.verify(req.token, env.jwtsecret, function(err, data){
            if (err){
                res.sendstatus(403)
            } else {
                res.status(200).json({
                    'msg': 'hello world',
                    'status': 'success'
                })
            }
        })

    } catch (error) {
        res.status(500).json({
            'error': error.message
        })
    }
}
export const registerUser = async(req, res) => {
    try {
        const response = await user.create(req.body)
        let hashedPassword = await bcrypt.hash(red.body.password, 8)
        user.query('INSERT INTO users SET ?', {name: req.body.name, email: req.body.email, username: req.body.username, password: hashedPassword}, (error) => {
            if (error) {
                console.log(error)
            } else {
                res.status(201).json({
                    'msg': "success registered",
                    'code': 201,
                })
            }
        })
        amqp.connect('amqp://localhost', function(error0, connection){
            if (error0) {
                throw error0
            }

            connection.createChannel(function(error1, channel) {
                if (error1) {
                    throw error1
                }

                var queue = "user-created"

                channel.assertQueue(queue, {
                    durable: false
                });

                channel.sendToQueue(queue, Buffer.from(response))
                console.log("[x] Send %s", response);
            })
        }) 
        
    } catch (error) {
        res.status(500).json({
            'msg': 'failure cant create user',
            'error': error,
        })
    }
}

export const loginUser = async(req, res) => {
    const secret = env.jwtsecret
    try {
        const response = await user.findOne({
            where:{
                email: req.body.email
            }
        })
    
        const token = jwt.sign(req.body, secret);
        res.status(200).json({
            'msg': 'success login',
            'token': token
        })
    } catch (error) {
        res.status(500).json({'msg': 'something error', 'error': error.message})
    }
}

export const detailsuser = async(req, res) => {
    try {
        const response = await user.findOne({
            where:{
                id: req.params.id
            }
        })
        res.status(200).json({
            'msg': 'success',
            'code': 200,
            'data': {
                'name': response.name,
                'email': response.email
            }
        })

    } catch (error) {
        res.status(500).json({
            'msg': 'erorr cant find account',
            'code': 500,
            'desc': error.message
        })
    }
}