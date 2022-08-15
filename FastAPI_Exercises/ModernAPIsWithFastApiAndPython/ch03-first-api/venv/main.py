import fastapi
import uvicorn
from typing import Optional

api = fastapi.FastAPI()

@api.get('/')
def index():
    body = "<html>" \
           "<body style='padding: 10px;'>" \
           "<h1>Welcome to API</h1>"  \
           "<div>" \
           "Try it: <a href='/api/calculate?x=7&y=11'>/api/calculate?x=7&y=11</a>"  \
           "</div>" \
           "</body>"  \
    "</html>"

    return fastapi.responses.HTMLResponse(content=body)

@api.get('/api/calculate')
def calculate(x: int, y: int, z: Optional[int] = None):
    if z == 0:
        return fastapi.responses.JSONResponse(content='{"ERROR": z cannot be zero."}', status_code=400)
#        return fastapi.responses.JSONResponse(status_code=400, content=error_dict)


    value = x + y

    if z is not None:
#        value /= z
        value = value/z

    return {
        'x': x,
        'y': y,
        'z': z,
        'value': value
    }


uvicorn.run(api, port=8000, host='127.0.0.1')