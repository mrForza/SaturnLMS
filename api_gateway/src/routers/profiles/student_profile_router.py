from fastapi import APIRouter

router = APIRouter(prefix="/student-profiles")

@router.post("/")
async def create_student_profile():
    return {"message": "Create student profile"}

@router.get("/")
async def get_student_profiles():
    return {"message": "List of student profiles"}

@router.patch("/")
async def update_student_profile():
    return {"message": "Update student profile"}

@router.delete("/{profile_id}")
async def delete_student_profile(profile_id: int):
    return {"message": f"Delete student profile with ID {profile_id}"}
