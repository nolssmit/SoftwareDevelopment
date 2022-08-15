import fastapi
import uvicorn

app = fastapi.FastAPI()


# we decorate the function.  It's not a regular
#   function but a function on the web
@app.get('/')
def index():
    return {
        'message': "Hello world"
    }

if __name__ == '__main__':
    uvicorn.run(app)

