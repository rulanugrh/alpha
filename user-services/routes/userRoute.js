import authToken  from '../middleware/authMiddleware.js';
import express from 'express'
import { registerUser, loginUser, testingToken } from '../controller/userController.js';

const router = express.Router();

router.post('/users/register', registerUser)
router.post('/users/login', loginUser)
router.get('/users', authToken, testingToken)

export default router;