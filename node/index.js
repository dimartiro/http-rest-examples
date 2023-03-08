import express from 'express'
import axios from 'axios';
const app = express()
const port = 8080

const URL = "https://dummyjson.com/todo/1"

app.get('/hello', (req, res) => {
  res.send('World!')
})

app.get('/performcall', async (req, res) => {
  try {
    const response = await axios.get(URL);
    res.send(response.data)
  } catch (e) {
    res.status(500).send({ "message": "error getting data from external service" })
  }
})

app.listen(port, () => {
  console.log(`Listening on port ${port}`)
})