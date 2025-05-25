from fastapi import APIRouter

router = APIRouter(prefix="/lessons")

@router.post("/")
async def create_lesson():
    return {"message": "Create admin profile"}

@router.get("/")
async def get_all_lessons():
    return {"message": "List of admin profiles"}

@router.patch("/")
async def update_lesson():
    return {"message": "Update admin profile"}

@router.delete("/{lesson_id}")
async def delete_lesson(lesson_id: int):
    return {"message": f"Delete lesson with ID {lesson_id}"}