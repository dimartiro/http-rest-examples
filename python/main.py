from fastapi import FastAPI
from fastapi.responses import JSONResponse
import requests
import uvicorn

app = FastAPI()

URL = "https://dummyjson.com/todo/1"

@app.get("/hello")
def ping():
    return "World!"

@app.get("/performcall")
def call():
    try:
        response = requests.get(URL)
        return response.json()
    except: 
        return JSONResponse(
            status_code=500,
            content={"message": "Error executing request"},
        )

if __name__ == '__main__':
    uvicorn.run(app, host="0.0.0.0", port=8080)