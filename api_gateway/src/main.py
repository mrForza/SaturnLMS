from fastapi import FastAPI
from routers import auth_router
from routers.profiles import student_profile_router, teacher_profile_router, admin_profile_router

app = FastAPI()

app.include_router(auth_router.router)
app.include_router(student_profile_router.router)
app.include_router(teacher_profile_router.router)
app.include_router(admin_profile_router.router)

print('Lorem ipsum v2.0')
