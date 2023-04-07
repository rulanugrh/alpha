import { Sequelize } from 'sequelize';
import db from '../config/connection.js'

const { DataTypes } = Sequelize;
const User = db.define('user', {
    name: DataTypes.STRING,
    email: DataTypes.STRING,
    username: DataTypes.STRING,
    password: DataTypes.STRING
}, {
    freezeTableName: true
})

export default User;
(async() => {
    await db.sync();
})();

