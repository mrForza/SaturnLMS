from fastapi import APIRouter

router = APIRouter(prefix="/courses")

@router.post("/")
async def create_course():
    return {"message": "Create admin profile"}

@router.get("/")
async def get_all_courses():
    return {"message": "List of admin profiles"}

@router.patch("/")
async def update_course():
    return {"message": "Update admin profile"}

@router.delete("/{course_id}")
async def delete_course(course_id: int):
    return {"message": f"Delete course with ID {course_id}"}