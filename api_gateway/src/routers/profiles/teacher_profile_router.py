from fastapi import APIRouter

router = APIRouter(prefix="/teacher-profiles")

@router.post("/")
async def create_teacher_profile():
    return {"message": "Create teacher profile"}

@router.get("/")
async def get_teacher_profiles():
    return {"message": "List of teacher profiles"}

@router.patch("/")
async def update_teacher_profile():
    return {"message": "Update teacher profile"}

@router.delete("/{profile_id}")
async def delete_teacher_profile(profile_id: int):
    return {"message": f"Delete teacher profile with ID {profile_id}"}
