import axios from 'axios'

const request = axios.create({
	baseURL: 'http://localhost:8080/api/auth/',
	timeout: 1000
})

export default request

