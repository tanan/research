import axios from 'axios'

const baseDomain = "http://localhost:8080"
const baseURL = `${baseDomain}/v1`

export default axios.create({
  baseURL,
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json'
  }
})
