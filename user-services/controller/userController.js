const User = require('../models/userModel.js')

export const createUser = async(req, res) => {
    try {
        const response = await User.create(req.body)
        res.status(201).json({
            'msg': "success",
            'code': 201,
            'data': {
                'name': response.name,
                'username': response.username,
                'email': response.email,
            }
        })
    } catch (error) {
        res.status(500).json({
            'msg': 'failure cant create user',
            'error': error,
        })
    }
}