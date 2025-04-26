# -*- encoding: utf-8 -*-
# @File     : img-serve
# @Time     : 2023-11-21 16:33:21
# @Docs     : 验证码-API服务
import hashlib
from datetime import datetime

from fastapi.responses import RedirectResponse
from fastapi.security import OAuth2PasswordBearer
from fastapi import FastAPI, HTTPException, File, UploadFile, Request, Form, Depends, Header

from ddddocr import DdddOcr
from loguru import logger

app = FastAPI()

oauth2_scheme = OAuth2PasswordBearer(tokenUrl="token")

# 模拟数据库中的用户信息
fake_users_db = {
    "admin": {
        "username": "admin",
        "password": "hashedpassword",
    }
}

# 访问频率统计字典
ip_access_count = {}

# 日志配置
log = logger.bind(name="")
log.add("xxx.log", format="{time} | {level} | {message}", rotation="1 week")


# 设置重定向
@app.get("/", include_in_schema=False)
async def redirect_to_home():
    return {"status": 200, "data": "欢迎来到主页 => https://github.com/Fromsko"}


@app.exception_handler(404)
async def redirect_to_login(request: Request, exc: Exception):
    # 重定向到主页面
    return RedirectResponse(url="/")


# @app.middleware("http")
# async def log_host(request: Request, call_next):
#     log.bind(name=request.client.host).info("Request received")

#     if request.url.path.startswith("/api/v1"):

#         ip = request.client.host
#         if ip not in ip_access_count:
#             ip_access_count[ip] = {
#                 "次数": 1,
#                 "初次访问": datetime.now().strftime("%Y年%m月%d日 %H:%M:%S"),
#                 "最后一次": ""
#             }
#         else:
#             ip_access_count[ip]["次数"] += 1

#         ip_access_count[ip]["最后一次"] = datetime.now().strftime(
#             "%Y年%m月%d日 %H:%M:%S"
#         )

#     response = await call_next(request)
#     return response


class VerifyParams:
    def __init__(self) -> None:
        self.token = ""

    async def create_token(self, username: str, password: str) -> str:
        token_data = f"{username}{password}{datetime.now().timestamp()}"
        hashed_token = hashlib.sha256(token_data.encode()).hexdigest()
        self.token = hashed_token
        return hashed_token

    @staticmethod
    async def get_token(authorization: str = Header(...)) -> str:
        if not authorization.startswith("Bearer "):
            raise HTTPException(
                status_code=401,
                detail="The interface is not accessible.",
            )
        return authorization.replace("Bearer ", "")

    @staticmethod
    def msg(code: int = 200, msg: str = "", err: str = ""):
        return {"status": code, "msg": msg, "err": err}


verify = VerifyParams()


@app.post("/api/v1/login")
async def login(username: str = Form(...), password: str = Form(...)):
    if username in fake_users_db and password == "admin":
        return {"status": 200, "token": await verify.create_token(username, password)}
    else:
        return verify.msg(code=401, err="Invalid credentials")


@app.post("/api/v1/verify-code")
async def verify_code(file: UploadFile = File(...), token: str = Depends(verify.get_token)):
    if token != verify.token:
        return verify.msg(code=401, err="Invalid token")

    image_bytes = await file.read()

    try:
        result = DdddOcr(show_ad=False).classification(image_bytes)
        logger.info(f"Verification code recognized: {result}")
        print(result)
        return verify.msg(msg=result)
    except Exception as err:
        logger.error(f"Verification code recognition failed. Error: {err}")
    return verify.msg(code=500, err=f"{err}")

if __name__ == '__main__':
    import uvicorn
    uvicorn.run("main:app", host="0.0.0.0", port=8080, reload=True)
