import axios from 'axios'

const apiClient = axios.create({
  baseURL: `http://localhost:4000`,
  withCredentials: false,
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json'
  }
})

export default {
  getArticle(id) {
    return apiClient.get('/v1/articles/' + id)
  },
  postRequest(request) {
    return apiClient.post('/post', request)
  }
}