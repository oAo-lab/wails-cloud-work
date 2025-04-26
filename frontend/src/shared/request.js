/**
 * 请求接口
 */
import axios from 'axios'

const request = new (class ApiClient {
    constructor(baseURL) {
        this.baseURL = baseURL
    }

    // 登录用户
    async loginUser (userInfo) {
        try {
            const response = await axios.post(`${this.baseURL}/user/login`, userInfo)
            return response.data
        } catch (error) {
            throw error
        }
    }

    // 注册用户
    async registerUser (newUser) {
        try {
            const response = await axios.post(`${this.baseURL}/user/register`, newUser)
            return response.data
        } catch (error) {
            throw error
        }
    }

    // 更新用户
    async updateUser (userID, updatedUser) {
        try {
            const response = await axios.put(`${this.baseURL}/user/${userID}`, updatedUser)
            return response.data
        } catch (error) {
            throw error
        }
    }

    // 删除用户
    async deleteUser (userID) {
        try {
            const response = await axios.delete(`${this.baseURL}/user/${userID}`)
            return response.data
        } catch (error) {
            throw error
        }
    }
})("http://localhost:7001/api/v1")

export default request
